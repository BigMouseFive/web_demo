package demo

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/web_demo/v2/log"
)

func CommonPub(client mqtt.Client, topic string, data interface{}) {
	value, err := json.Marshal(data)
	if err != nil {
		log.Sugar.Errorf("[pub][failed ][%s]%s", topic, "transform data error")
		return
	}

	if client == nil {
		log.Sugar.Errorf("[pub][failed ][%s]%s", topic, "mqtt client is nil")
		return
	}

	token := client.Publish(topic, 0, false, value)
	token.Wait()
	log.Sugar.Infof("[pub][success][%s]%s", topic, value)
}

// PubDemo 请求Demo
func PubDemo(client mqtt.Client, transId string, timestamp string, deviceId string) {
	var data DemoModel
	data.TransId = transId
	data.Timestamp = timestamp
	data.DeviceId = deviceId
	CommonPub(client, Topic["demoTo"]+deviceId, data)
}
