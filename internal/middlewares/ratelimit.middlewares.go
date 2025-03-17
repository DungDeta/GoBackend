package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
	"myproject/global"
)

type RateLimiter struct {
	globalRateLimiter         *limiter.Limiter
	publicAPIRateLimiter      *limiter.Limiter
	userPrivateAPIRateLimiter *limiter.Limiter
}

func NewRateLimiter() *RateLimiter {
	rateLimiter := &RateLimiter{
		globalRateLimiter:         rateLimiter("1000-S"),
		publicAPIRateLimiter:      rateLimiter("80-M"),
		userPrivateAPIRateLimiter: rateLimiter("50-M"),
	}
	return rateLimiter
}
func rateLimiter(interval string) *limiter.Limiter {
	store, err := redisStore.NewStoreWithOptions(
		global.Rdb, limiter.StoreOptions{
			Prefix:          "limiter",
			MaxRetry:        3,
			CleanUpInterval: time.Hour,
		},
	)
	if err != nil {
		panic(err)
	}
	rate, err := limiter.NewRateFromFormatted(interval)
	if err != nil {
		panic(err)
	}
	return limiter.New(store, rate)
}

func (rl *RateLimiter) GlobalRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "global"
		log.Println("global--->")
		limitContext, err := rl.globalRateLimiter.Get(c, key)
		if err != nil {
			fmt.Println("Failed to check rate limit GLOBAL", err)
			c.Next()
			return
		}
		if limitContext.Reached {
			log.Printf("Rate limit breached GLOBAL %s", key)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached GLOBAL, try later"})
			return
		}
		c.Next()
	}

}
func (rl *RateLimiter) PublicAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {

		urlPath := c.Request.URL.Path // urlPath: /ping/80 80 >
		rateLimitPath := rl.filterLimitUrlPath(urlPath)
		if rateLimitPath != nil {
			log.Println("Client Ip--->", c.ClientIP())

			key := fmt.Sprintf("%s-%s", "111-222-333-44", urlPath)
			limitContext, err := rateLimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limit", err)
				c.Next()
				return
			}
			if limitContext.Reached {
				log.Printf("Rate limit breached %s", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached, try later"})
				return
			}
		}

		c.Next()
	}
}

func (rl *RateLimiter) UserAndPrivateRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {

		urlPath := c.Request.URL.Path
		rateLimitPath := rl.filterLimitUrlPath(urlPath)
		if rateLimitPath != nil {
			userId := 1001 // Private đi theo userId
			key := fmt.Sprintf("%d-%s", userId, urlPath)
			limitContext, err := rateLimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limit", err)
				c.Next()
				return
			}
			if limitContext.Reached {
				log.Printf("Rate limit breached %s", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached, try later"})
				return
			}
		}

		c.Next()

	}
}
func (rl *RateLimiter) filterLimitUrlPath(urlPath string) *limiter.Limiter {
	if urlPath == "/v1/2024/user/login" || urlPath == "/ping/80" {
		return rl.publicAPIRateLimiter
	} else if urlPath == "/v1/2024/user/info" || urlPath == "/ping/50" {
		return rl.userPrivateAPIRateLimiter
	} else {
		return rl.globalRateLimiter
	}
}
