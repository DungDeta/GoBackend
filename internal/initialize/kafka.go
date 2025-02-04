package initialize

import (
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"myproject/global"
)

var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP("localhost:19092"),
		Topic:    "otp-auth-topic",
		Balancer: &kafka.LeastBytes{},
	}
}
func CloseKafka() {
	err := global.KafkaWriter.Close()
	if err != nil {
		global.Logger.Error("Close kafka writer error", zap.Error(err))
		return
	}
}
