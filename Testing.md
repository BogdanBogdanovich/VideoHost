type T struct {
NQroYmogp struct {
Publisher struct {
App            string    `json:"app"`
Stream         string    `json:"stream"`
ClientId       string    `json:"clientId"`
ConnectCreated time.Time `json:"connectCreated"`
Bytes          int       `json:"bytes"`
Ip             string    `json:"ip"`
Audio          struct {
Codec      string `json:"codec"`
Profile    string `json:"profile"`
Samplerate int    `json:"samplerate"`
Channels   int    `json:"channels"`
} `json:"audio"`
Video struct {
Codec   string  `json:"codec"`
Width   int     `json:"width"`
Height  int     `json:"height"`
Profile string  `json:"profile"`
Level   float64 `json:"level"`
Fps     int     `json:"fps"`
} `json:"video"`
} `json:"publisher"`
Subscribers []struct {
App            string    `json:"app"`
Stream         string    `json:"stream"`
ClientId       string    `json:"clientId"`
ConnectCreated time.Time `json:"connectCreated"`
Bytes          int       `json:"bytes"`
Ip             string    `json:"ip"`
Protocol       string    `json:"protocol"`
} `json:"subscribers"`
} `json:"nQroYmogp"`
}