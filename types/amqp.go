package types

type AMQPConfig struct {
	URL          string `json:"url"`
	Exchange     string `json:"exchange"`
	Queue        string `json:"queue"`
	Kind         string `json:"kind"`
	Key          string `json:"key"`
	Durable      bool   `json:"durable"`
	AutoDelete   bool   `json:"auto_delete"`
	DeliveryMode uint   `json:"delivery_mode"`
}
