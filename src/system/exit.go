package system

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"cron/job"
	"sync/atomic"
	"time"
	"kafka"
	"mysql"
)

func Sig() {
	savePid()
	fmt.Println("check signal")
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGUSR2, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	fmt.Println("get signal :", s)
	fmt.Println("开始进程收尾工作")

	for _,job := range job.CronJobInfo.CronJobList{
		job.ForceStop = true
	}
	stopTime1 := time.NewTimer(time.Millisecond * 1000)
	<-stopTime1.C

	for {
		httpRequestNotFinishNum := atomic.LoadInt64(&job.HttpRequestNotFinishNum)
		if httpRequestNotFinishNum > 0 {
			fmt.Println("httpRequestNotFinishNum===", httpRequestNotFinishNum)
			stopTime := time.NewTimer(time.Millisecond * 2000)
			<-stopTime.C
			continue
		}
		break
	}

	stopTime2 := time.NewTimer(time.Millisecond * 1000)
	<-stopTime2.C

	kafka.Close()
	for _,job := range job.CronJobInfo.CronJobList{
		// 更新offset
		mysql.UpsertOffset(job.Topic, job.Partition, job.KafkaOffset)
	}


	fmt.Println("结束进程收尾工作")
	removePid()
	os.Exit(0)
}
