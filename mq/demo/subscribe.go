package demo

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/web_demo/v2/log"
	"reflect"
	"strings"
)

func CommonSub(client mqtt.Client, topic string, handler mqtt.MessageHandler) {
	f1 := reflect.ValueOf(CommonHandler)
	f2 := reflect.ValueOf(handler)
	if f1.Pointer() == f2.Pointer() {
		handler = nil
	}
	token := client.Subscribe(topic, 0, CommonHandler)
	token.Wait()
	log.Sugar.Infof("mqtt subscribe %s", topic)
	regexTopic := strings.ReplaceAll(topic, "/", "\\/")
	regexTopic = strings.ReplaceAll(regexTopic, "+", "[^/]+")
	regexTopic = strings.ReplaceAll(regexTopic, "#", "(.+)")
	regexTopic = "^" + regexTopic + "$"
	MqttHandlerMap[regexTopic] = handler
}

func Subscribe(client mqtt.Client) {
	CommonSub(client, Topic["demoFrom"]+"#", DemoHandler)
}
