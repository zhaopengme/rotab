package simple_rpc

import (
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

// fPtr 指向函数原型
func (c *Client) callRPC(rpcName string, fPtr interface{}) {

	fn := reflect.ValueOf(fPtr).Elem()

	f := func(args []reflect.Value) []reflect.Value {

		// 处理输入参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}

		// 编码数据
		cliSession := NewSession(c.conn)
		reqRPC := RPCData{Name: rpcName, Args: inArgs}
		b, err := encode(reqRPC)
		if err != nil {
			panic(err)
		}

		err = cliSession.Write(b)
		if err != nil {
			panic(err)
		}

		respBytes, err := cliSession.Read()
		if err != nil {
			panic(err)
		}

		respRPC, err := decode(respBytes)
		if err != nil {
			panic(err)
		}

		outArgs := make([]reflect.Value, 0, len(respRPC.Args))
		for i, arg := range respRPC.Args {
			// 必须进行 nil 转换
			if arg == nil {
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}
	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)
}
