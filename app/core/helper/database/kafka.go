package database

import (
	"go-boilerplate/app/core/helper/logger"
	"go-boilerplate/config"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

type Producer struct {
	MostOnceProducer sarama.AsyncProducer
	AsyncProducer    sarama.AsyncProducer
	SyncProducer     sarama.SyncProducer
	Signal           chan os.Signal
}

func NewKafkaProducer(config *config.Config) *Producer {
	conf := sarama.NewConfig()
	conf.Version = sarama.V0_11_0_0
	conf.Producer.Retry.Max = 3
	conf.Producer.RequiredAcks = sarama.WaitForLocal
	conf.Producer.Compression = sarama.CompressionNone // CHECK 관련 여쭤보기
	conf.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	async, err := sarama.NewAsyncProducer(config.Kafka.Broker, conf)
	if err != nil {
		logger.Zap.Fatalf("NewKafkaProducer NewAsyncProducer error %v", err)
	}

	syncConf := sarama.NewConfig()
	syncConf.Producer.Return.Successes = true
	syncConf.Version = sarama.V0_11_0_0
	syncConf.Producer.Retry.Max = 3
	syncConf.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	sync, err := sarama.NewSyncProducer(config.Kafka.Broker, syncConf)
	if err != nil {
		logger.Zap.Fatalf("NewKafkaProducer NewSyncProducer error %v", err)
	}

	mostOnceConf := sarama.NewConfig()
	mostOnceConf.Version = sarama.V0_11_0_0
	mostOnceConf.Producer.Return.Successes = false
	mostOnceConf.Producer.Return.Errors = false
	mostOnceConf.Producer.RequiredAcks = sarama.NoResponse
	mostOnce, err := sarama.NewAsyncProducer(config.Kafka.Broker, mostOnceConf)
	if err != nil {
		logger.Zap.Fatalf("NewKafkaProducer MostOnceProducer error %v", err)
	}

	producer := &Producer{
		MostOnceProducer: mostOnce,
		AsyncProducer:    async,
		SyncProducer:     sync,
		Signal:           make(chan os.Signal, 1),
	}

	signal.Notify(producer.Signal, os.Interrupt)
	signal.Notify(producer.Signal, os.Kill)

	go func() {
		select {
		case err := <-producer.MostOnceProducer.Errors():
			logger.Zap.Errorf("kafka write error: %+v", err)
		case err := <-producer.AsyncProducer.Errors():
			logger.Zap.Errorf("kafka write error: %+v", err)
		case osSignal := <-producer.Signal:
			if err := producer.AsyncProducer.Close(); err != nil {
				logger.Zap.Errorf("kafka close error: %+v", err)
			}
			logger.Zap.Errorf("os signal occurrence: %+v", osSignal)

			if err := producer.MostOnceProducer.Close(); err != nil {
				logger.Zap.Errorf("kafka close error: %+v", err)
			}
			logger.Zap.Errorf("os signal occurrence: %+v", osSignal)
			break
		}
	}()

	return producer
}
