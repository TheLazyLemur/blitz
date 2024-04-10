package protocol

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"testing"
)

func Test_parseRequest(t *testing.T) {
	buf := new(bytes.Buffer)

	cmd := Command{
		Command: "set",
		Key:     "key",
		Value:   "value",
	}

	bytes, _ := json.Marshal(cmd)
	str := string(bytes)

	keyLen := int32(len(str))
	if err := binary.Write(buf, binary.LittleEndian, keyLen); err != nil {
		t.Fatal(err)
	}

	if err := binary.Write(buf, binary.LittleEndian, []byte(str)); err != nil {
		t.Fatal(err)
	}

	actual, err := ParseRequest(buf)
	if err != nil {
		t.Fatal(err)
	}

	if actual != cmd {
		t.Fail()
	}
}
