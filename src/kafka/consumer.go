package kafka
//
//import (
//	"github.com/Shopify/sarama"
//	"sync"
//	"log"
//	//"fmt"
//	//"cron/job"
//	//"time"
//	"strconv"
//	"config"
//	//"time"
//)
//
//var consumer sarama.Consumer
//var consumerErr error
//
//var partitionConsumerMap map[string]bool
//var partitionMu sync.Mutex
//
//func init()  {
//
//	partitionConsumerMap = make(map[string]bool)
//    consumer,consumerErr = sarama.NewConsumer([]string{config.KafkaServerConfig.Host + ":" + strconv.Itoa(config.KafkaServerConfig.Port)},nil)
//    if consumerErr != nil{
//    	log.Fatal("consumerErr===",consumerErr)
//	}
//}
//
//func Close()  {
//	consumer.Close()
//}
//
//func GetNewMessage(deviceOS string, appName string) {
//	//var offset = 0
//	//var _partition = 0
//	//partition, partitionErr := consumer.ConsumePartition(config.Topic+"_" + deviceOS + "_" + appName, _partition, offset)
//	//
//	//if partitionErr != nil {
//	//	log.Fatal("partitionErr===", partitionErr, config.Topic+"_" + deviceOS + "_" + appName,)
//	//}
//	//
//	//for{
//	//	if !cronJob.ForceStop{
//	//		if  len(partition.Messages()) > 0 {
//	//			*<-partition.Messages()
//	//			continue
//	//		}else{  //kafka没有新数据的时候才暂停防止for死循环，有数据通过通道阻塞
//	//			pauseTime := time.NewTimer(time.Millisecond * 1)
//	//			<-pauseTime.C
//	//		}
//	//	} else {
//	//		partition.Close()
//	//		break
//	//	}
//	//
//	//}
//
//
//
//}
