package rpc_test

import (
	"educationalsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncoding(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
    if expected != actual {
        t.Fatalf("Expected: %s, Actual: %s", expected, actual)
    }
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
    method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
    if err != nil {
        t.Fatal(err)
    }

    contentLength := len(content)

    if contentLength != 15 {
        t.Fatalf("Expected: 15, Got %d", contentLength)
    }

    if method != "hi" {
        t.Fatalf("Expected: \"hi\", Got: %s", method)
    }


}
