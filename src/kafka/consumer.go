package kafka

import (
	"github.com/Shopify/sarama"
	//"sync"
	"log"
	//"fmt"
	//"cron/job"
	//"time"
	"strconv"
	"config"
	//"time"
	//"time"
	"fmt"
)

var consumer sarama.Consumer
var consumerErr error


func init()  {

    consumer,consumerErr = sarama.NewConsumer([]string{config.KafkaServerConfig.Host + ":" + strconv.Itoa(config.KafkaServerConfig.Port)},nil)
    if consumerErr != nil{
    	log.Fatal("consumerErr===",consumerErr)
	}
}

func Close()  {
	consumer.Close()
}

func GetNewMessage(deviceOS string, appName string, _partition int32) {

	var offset int64 = 0
	partition, partitionErr := consumer.ConsumePartition(config.Topic+"_" + deviceOS + "_" + appName, _partition, offset)

	if partitionErr != nil {
		log.Fatal("partitionErr===", partitionErr, config.Topic+"_" + deviceOS + "_" + appName,)
	}
	fmt.Println( fmt.Sprintf("partition = %d, length = %d, %v, length = %d,", _partition, len(partition.Messages()), *<-partition.Messages(), len(partition.Messages())))

	//for msg := range partition.Messages() {
	//	fmt.Println(msg.Partition)
	//}
	for{

		if  len(partition.Messages()) > 0 {
			message := *<-partition.Messages()
			fmt.Println( fmt.Sprintf("partition = %d, val = %v, current offset = %d", _partition, string(message.Value), message.Offset))
		}else{
			break //这里是个大坑
			//partition.Close()
		}
	}



}
