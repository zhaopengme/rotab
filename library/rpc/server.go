package simple_rpc

import (
	"fmt"
	"net"
	"reflect"
)

type Server struct {
	addr  string
	funcs map[string]reflect.Value
}

func NewServer(addr string) *Server {
	return &Server{addr: addr, funcs: make(map[string]reflect.Value)}
}

// 等待调用
func (s *Server) Run() {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Printf("listen on %s err: %v", s.addr, err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("accept err: %v", err)
			return
		}
		srvSession := NewSession(conn)

		// 读取 RPC 调用数据
		b, err := srvSession.Read()
		if err != nil {
			fmt.Printf("read err: %v", err)
			return
		}

		// 解码 RPC 调用数据
		rpcData, err := decode(b)
		if err != nil {
			fmt.Printf("decode err: %v", err)
			return
		}

		f, ok := s.funcs[rpcData.Name]
		if !ok {
			fmt.Printf("func %s not exists", rpcData.Name)
			return
		}

		// 构造函数的参数
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}

		// 执行调用
		out := f.Call(inArgs)

		outArgs := make([]interface{}, 0, len(out))
		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}

		// 包装数据返回给客户端
		respRPCData := RPCData{rpcData.Name, outArgs}
		respBytes, err := encode(respRPCData)
		if err != nil {
			fmt.Printf("encode err: %v", err)
			return
		}

		err = srvSession.Write(respBytes)
		if err != nil {
			fmt.Printf("session write err: %v", err)
			return
		}
	}
}

func (s *Server) Register(rpcName string, f interface{}) {
	if _, ok := s.funcs[rpcName]; ok {
		return
	}
	fVal := reflect.ValueOf(f)
	s.funcs[rpcName] = fVal
}
