package simple_rpc

import (
	"bytes"
	"encoding/gob"
)

type RPCData struct {
	Name string
	Args []interface{}
}

func encode(data RPCData) ([]byte, error) {
	var buf bytes.Buffer
	bufEnc := gob.NewEncoder(&buf)
	if err := bufEnc.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(b []byte) (RPCData, error) {
	buf := bytes.NewBuffer(b)
	bufDec := gob.NewDecoder(buf)
	var data RPCData
	if err := bufDec.Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}
