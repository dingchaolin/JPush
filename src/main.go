package main

import (
	"apns2_test"
	_ "kafka"
	//"kafka"
	//"strconv"
	"jpush"
	//"sync"
	//"time"
	//"fmt"
	//"runtime"
	//"log"
	"fmt"
	"strconv"
	"kafka"
)

const (
	appKey = "ae118ce7c7a0adda4b996ba2"
	secret = "5d214769184ae5f71dd0717d"
	Num_Partitions = 10
)

func JPush(){
	jpush.Push(appKey, secret)
}
func iosPush(){
	apns2_test.Test()
}

func iosPushP8(){
	apns2_test.TestP8()
}
/*
生产
 */
func produce( index int){
	fmt.Println("开始写入")
	data := []string{}
	for i := index*1000; i < (index+1)*2000; i ++ {
		data = append(data, strconv.Itoa(i))

	}
	kafka.SavePushData("anzhuo", "art", data )
	fmt.Println("写入完成!")
}
//
///*
//消费
// */
//func consume(){
//	for i := 0; i < 10; i ++ {
//		go func(i int) {
//			kafka.GetNewMessage("ios", "express", int32(i)  )
//		}(i)
//	}
//}
//
//func consume_c(){
//	for i := 0; i < 1; i ++ {
//		go kafka.GetNewMessage_c("ios", "express", int32(i)  )
//
//	}
//}
//
//func GetNum(){//返回所有的分区id
//	nums, err := kafka.GetPartitionsNums("ios", "express")
//	if err != nil {
//		log.Fatal( err )
//	}
//	fmt.Println( "nums=========================",nums )
//}

func main() {
	//GetNum()
	////runtime.GOMAXPROCS(4000)
	////produce(0)
	//consume()
	////consume_c()
	//
	//ret, _:= kafka.GetPartitionsNums("ifdsafdsafdsaos", "1fdsafdsafdsafdsa23")
	//fmt.Println( "------------",ret )
	//go func(){
	//	//runtime.GOMAXPROCS(*numCores)
	//	for {
	//		fmt.Println( "----------------------------",runtime.NumGoroutine())
	//		time.Sleep(time.Second*2)
	//	}
	//}()
	//
	//time.Sleep( time.Second * 300 )
	////produce(3)
	//
	////time.Sleep( time.Second * 300 )

	//iosPush()
	//iosPushP8()
	produce(0)

}
