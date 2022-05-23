package demo

type DemoModel struct {
	TransId   string `json:"transId"`   // 消息的唯一标识，统一使用 UUID
	DeviceId  string `json:"deviceId"`  // 设备唯一标识，与设备注册创建设备 ID 对应
	Timestamp string `json:"timestamp"` // 时间戳，消息上报的时间，UTC 时间（格式：yyyyMMdd'T'HHmmssSSS'Z'），如：20161219T114920178Z
}
