package global

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
	"myproject/pkg/logger"
	"myproject/pkg/setting"
)

var (
	Config      setting.Config
	Logger      *logger.LoggerZap
	Mdb         *gorm.DB
	Rdb         *redis.Client
	Mdbc        *sql.DB
	KafkaWriter *kafka.Writer
)
