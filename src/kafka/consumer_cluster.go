package kafka

import (
	cluster "github.com/bsm/sarama-cluster"
	//"sync"
	"log"
	//"fmt"
	//"cron/job"
	//"time"
	//"strconv"
	"config"
	//"time"
	//"time"
	"fmt"
	"os"
	"github.com/Shopify/sarama"
)

var consumer_c *cluster.Consumer
var topic = "PushApiQueue__ios_express"// groupID 需要协调offset   同一个group中 公用同一个offset
// 一个parition只能被一个consumer消费。 如果多个consumer消费同一个partition 需要协调offset
// 多个consumer 消费同一个 topic 中的不同partition


func init()  {
	config := cluster.NewConfig()
	config.Group.Mode = cluster.ConsumerModePartitions
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// init consumer
	brokers := []string{"192.168.64.185:9092"}
	topics := []string{topic}

	consumer_c,consumerErr = cluster.NewConsumer(brokers,topic,topics,config)
    if consumerErr != nil{
    	log.Fatal("consumerErr===",consumerErr)
	}
}

func Close_c()  {
	consumer_c.Close()
}

func GetNewMessage_c(deviceOS string, appName string, _partition int32) {
	//var offset int64 = 0
	partition, ok := <-consumer_c.Partitions()

	if !ok {
		log.Fatal("partitionErr===", ok, config.Topic+"_" + deviceOS + "_" + appName,)
	}


	go func(pc cluster.PartitionConsumer) {
		for msg := range pc.Messages() {
			fmt.Fprintf(os.Stdout, "携程1============%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
			consumer_c.MarkOffset(msg, "")	// mark message as processed
		}
	}(partition)

	go func(pc cluster.PartitionConsumer) {
		for msg := range pc.Messages() {
			fmt.Fprintf(os.Stdout, "携程2============%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
			consumer_c.MarkOffset(msg, "")	// mark message as processed
		}
	}(partition)

	go func(pc cluster.PartitionConsumer) {
		for msg := range pc.Messages() {
			fmt.Fprintf(os.Stdout, "携程3============%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
			consumer_c.MarkOffset(msg, "")	// mark message as processed
		}
	}(partition)

	go func(pc cluster.PartitionConsumer) {
		for msg := range pc.Messages() {
			fmt.Fprintf(os.Stdout, "携程4============%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
			consumer_c.MarkOffset(msg, "")	// mark message as processed
		}
	}(partition)

	go func(pc cluster.PartitionConsumer) {
		for msg := range pc.Messages() {
			fmt.Fprintf(os.Stdout, "携程5============%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
			consumer_c.MarkOffset(msg, "")	// mark message as processed
		}
	}(partition)
	go func(pc cluster.PartitionConsumer) {
		for msg := range pc.Messages() {
			fmt.Fprintf(os.Stdout, "携程6============%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
			consumer_c.MarkOffset(msg, "")	// mark message as processed
		}
	}(partition)


}
