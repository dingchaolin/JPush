package apns2_test

import (
"log"
"fmt"

"github.com/sideshow/apns2"
"github.com/sideshow/apns2/certificate"
	//"path/filepath"
	//"os"
)

func Test() {
	cert, err := certificate.FromP12File("./pem/release_Apns_Cer.p12", "")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "5b6a29435cf1177cfadf586812b800a5d0492882dd13f70841e3cef07d495961"
	notification.Topic = "net.artron.artexpress"
	notification.Payload = []byte(`{"aps":{"alert":"Production Hello!"}}`) // See Payload section below

	client := apns2.NewClient(cert).Development()
	//client := apns2.NewClient(cert).Production()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
