package kafka

import (
	"github.com/Shopify/sarama"
	"fmt"
	//"strconv"
	"config"
	"log"
	"strconv"
)

var Producer sarama.SyncProducer

func init() {
	p, err := newProducer()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	Producer = p
}

func newProducer() (sarama.SyncProducer, error) {
	configs := sarama.NewConfig()
	configs.Producer.MaxMessageBytes = 102400000
	configs.Producer.Return.Successes = true
	configs.Producer.Return.Errors = true
	configs.Producer.RequiredAcks = sarama.WaitForAll
	configs.Producer.Partitioner = sarama.NewRandomPartitioner
	return sarama.NewSyncProducer([]string{config.KafkaServerConfig.Host + ":" + strconv.Itoa(config.KafkaServerConfig.Port)}, configs)
}

func SavePushData(deviceOS string, appName string, list []string) int {

	var arr []*sarama.ProducerMessage
	strLength := 0
	for i := 0; i < len(list); i++ {
		msg := &sarama.ProducerMessage{
			Topic:     config.Topic + "_" + deviceOS + "_" + appName,
			Partition: int32(1),//不指定分区 就会将数据平均的写到所有的分区上去
			Key:       sarama.StringEncoder(""),
			Value:     sarama.StringEncoder(list[i]),
		}
		strLength += len(list[i])
		arr = append(arr, msg)
	}

	err := Producer.SendMessages(arr)
	if err != nil {
		if config.GetEnv() != "production" {
			log.Fatal("send message fail", err)
		} else {
			fmt.Println("send message fail", err)
		}

	}
	list = nil
	return strLength
}
