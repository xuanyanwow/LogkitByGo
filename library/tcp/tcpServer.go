package tcp

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/os/glog"
	"reflect"
)

type Tcp struct {
}

type Message struct {
	Action string
	Data   interface{}
	Sign   string
}

var (
	userConnMap map[string]*gtcp.Conn
	connUserMap map[*gtcp.Conn]string
	// 方法映射 增加功能需要增加映射
	functions map[string]interface{}
)

func (t Tcp) SetFunctionMap(m map[string]interface{}) {
	functions = m
}

func (t Tcp) BindUId(c *gtcp.Conn, uId string) {
	userConnMap[uId] = c
	connUserMap[c] = uId
}

func (t Tcp) GetConnByUId(uId string) *gtcp.Conn {
	conn, ok := userConnMap[uId]
	if !ok {
		return nil
	}
	return conn
}

func (t Tcp) GetUIdByConn(conn *gtcp.Conn) string {
	UId, ok := connUserMap[conn]
	if !ok {
		return ""
	}
	return UId
}

func Encode(str string) string {
	//strLen := len(str)
	return hex.EncodeToString([]byte(str))
}

// 解析tcp内容触发控制器
func (t Tcp) Parse(c *gtcp.Conn, data string) {
	if j, err := gjson.DecodeToJson([]byte(data)); err != nil {
		glog.Error(err)
	} else {
		m := Message{}
		toErr := j.ToStruct(&m)
		if toErr != nil {
			fmt.Println("解析错误")
			return
		}
		// 校验签名

		// 第一个参数是返回值 这里不需要用到
		r, err := Call(functions, m.Action, m.Data)
		if err != nil {
			fmt.Println("method invoke error:", err)
		}

		if len(r) > 0 && r[0].String() != "" {
			c.SendPkg([]byte(r[0].String()))
		}
	}
}

func Call(m map[string]interface{}, name string, params ...interface{}) ([]reflect.Value, error) {
	if _, ok := m[name]; !ok {
		return nil, errors.New("function not exist")
	}
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("the number of input params not match!")
	}
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	return f.Call(in), nil
}
