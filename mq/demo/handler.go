package demo

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/web_demo/v2/log"
	"regexp"
)

var (
	MqttHandlerMap map[string]mqtt.MessageHandler
)

func InitHandlerMap() {
	MqttHandlerMap = map[string]mqtt.MessageHandler{}
}

// CommonHandler 公共回调
func CommonHandler(client mqtt.Client, msg mqtt.Message) {
	for topic, handler := range MqttHandlerMap {
		if match, _ := regexp.MatchString(topic, msg.Topic()); match {
			log.Sugar.Infof("[sub][match][%s]%s", msg.Topic(), msg.Payload())
			handler(client, msg)
			return
		}
	}
	log.Sugar.Warnf("[sub][miss ][%s][%s]", msg.Topic(), msg.Payload())
}

func DemoHandler(client mqtt.Client, msg mqtt.Message) {

}
