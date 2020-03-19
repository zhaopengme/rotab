package simple_rpc

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

func TestRPC(t *testing.T) {

	gob.Register(User{})

	addr := "0.0.0.0:2333"
	srv := NewServer(addr)
	srv.Register("queryUser", queryUser)
	go srv.Run()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Error(err)
	}
	cli := NewClient(conn)

	var query func(int) (User, error)
	cli.callRPC("queryUser", &query)

	u, err := query(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}

type User struct {
	Name string
	Age  int
}

func queryUser(uid int) (User, error) {
	userDB := make(map[int]User)
	userDB[0] = User{"Dennis", 70}
	userDB[1] = User{"Ken", 75}
	userDB[2] = User{"Rob", 62}
	if u, ok := userDB[uid]; ok {
		return u, nil
	}
	return User{}, fmt.Errorf("id %d not in user db", uid)
}
