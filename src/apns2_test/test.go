package apns2_test

import (
"log"
"fmt"

"github.com/sideshow/apns2"
"github.com/sideshow/apns2/certificate"
"github.com/sideshow/apns2/token"
	//"path/filepath"
	//"os"
	//"github.com/sideshow/apns2/payload"
)

func Test() {
	cert, err := certificate.FromP12File("./artexpress_inhouse.p12", "")
	if err != nil {
		log.Fatal("Cert Error###########################:", err)
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "e29594ddaf98261819379b539f9e76ef3659568844f18c690b7305b4611ec4a6"
	notification.Topic = "com.artron.artexpress"
	notification.Payload = []byte(`{"aps":{"badge":0,"alert":{"title":"dcl标题", "body":"dcl内容"}}}`) // See Payload section below
	//notification.Payload = payload.NewPayload().AlertTitle(p.Title).AlertBody(p.Context).Badge(p.Badge).Sound(p.Sound)

	client := apns2.NewClient(cert).Development()
	//client := apns2.NewClient(cert).Production()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error==------------------:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}


func TestP8(){
	authKey, err := token.AuthKeyFromFile("./release.p8")
	if err != nil {
		log.Fatal("token error:", err)
	}

	token := &token.Token{
		AuthKey: authKey,
		// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
		KeyID:   "W8VUK89K53",
		// TeamID from developer account (View Account -> Membership)
		TeamID:  "HAV44L48AA",
	}
	notification := &apns2.Notification{}
	notification.DeviceToken = "5b6a29435cf1177cfadf586812b800a5d0492882dd13f70841e3cef07d495961"
	notification.Topic = "net.artron.artexpress"
	notification.Payload = []byte(`{"aps":{"badge":0,"alert":{"title":"dcl标题", "body":"dcl内容"}}}`) // See Payload section below
	//notification.Payload = paylo
	//client := apns2.NewTokenClient(token).Development()
	client := apns2.NewTokenClient(token)
	res, err := client.Push(notification)
	fmt.Println( res, err )
}
