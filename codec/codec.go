package codec

import (
	"io"
)

// Header 定义一个Header报头
type Header struct {
	ServiceMethod string // 消息方法
	Seq           int64  // 序列号
	Error         string // 错误信息
}

// Codec 定义消息编解码的接口
type Codec interface {
	io.Closer // 关闭
	ReadHead(header *Header) error
	ReadBody(body interface{}) error
	Write(header *Header, body interface{}) error
}

// NewCodecFunc 建造者模式
// 构造器
type NewCodecFunc func(conn io.ReadWriteCloser) Codec

var NewCodecFuncMap map[Type]NewCodecFunc

type Type string

var (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
