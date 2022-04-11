package database

import _ "fmt"

type Stream struct {
	StreamKey string `json:"stream_key"`
}

var Streams = []Stream{
	{StreamKey: "111"},
	{StreamKey: "123"},
}
