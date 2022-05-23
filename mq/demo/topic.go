package demo

var (
	Topic map[string]string
)

func InitTopic() {
	Topic = map[string]string{}
	Topic["demoTo"] = "to/v1/demo/"
	Topic["demoFrom"] = "from/v1/demo/"
}
