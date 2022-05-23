package mq

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/web_demo/v2/config"
	"github.com/web_demo/v2/log"
	"github.com/web_demo/v2/mq/demo"
	"time"
)

var (
	MqttCli mqtt.Client
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Sugar.Infof("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Sugar.Infof("Connected")
	demo.Subscribe(client)
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Sugar.Infof("Connect lost: %v", err)
}

func CreateMqttClient() {
	log.Sugar.Infow("CreateMqtt ", "mqtt_address", config.Config.GetString("mqtt.address"))

	// 初始化话题
	demo.InitTopic()

	// 初始化回调字典
	demo.InitHandlerMap()

	// 创建client
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.Config.GetString("mqtt.address"))
	opts.SetClientID("saas_mqtt_client")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(4 * time.Second)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	MqttCli = mqtt.NewClient(opts)
	if token := MqttCli.Connect(); token.Wait() && token.Error() != nil {
		log.Sugar.Fatal(token.Error())
	}
}
