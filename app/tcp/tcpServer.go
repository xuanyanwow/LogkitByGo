package tcp

import (
	"encoding/hex"
	"github.com/gogf/gf/net/gtcp"
)

type Tcp struct {

}

var (
	userConnMap map[string]*gtcp.Conn
	connUserMap map[*gtcp.Conn]string
)

func (t Tcp) BindUId(c *gtcp.Conn, uId string){
	userConnMap[uId] = c
	connUserMap[c]   = uId
}

func (t Tcp) GetConnByUId(uId string) *gtcp.Conn {
	conn,ok := userConnMap[uId]
	if !ok {
		return nil
	}
	return conn
}

func (t Tcp) GetUIdByConn (conn *gtcp.Conn) string {
	UId,ok := connUserMap[conn]
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
func (t Tcp) Parse(data string) {

}