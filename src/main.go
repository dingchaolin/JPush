package main

import (
	//"jpush"
	//"apns2_test"
	_ "kafka"
	"kafka"
	"strconv"
)

//const (
//	appKey = "ae118ce7c7a0adda4b996ba2"
//	secret = "5d214769184ae5f71dd0717d"
//)

func main() {

	//jpush.Push(appKey, secret)
	//apns2_test.Test()
	data := []string{}
	for i := 0; i < 10; i ++ {
		data = append(data, strconv.Itoa(i))

	}
	kafka.SavePushData("ios", "express", data )

}
