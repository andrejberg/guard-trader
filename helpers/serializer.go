package shared

import (
	"bytes"
	"fmt"
	"log"

	acc "github.com/guard-trader/account"
	"github.com/ugorji/go/codec"
)

var (
	ch codec.CborHandle
)

func DeserializerString(d []byte) string {
	var str string
	dec := codec.NewDecoderBytes(d, &ch)
	if err := dec.Decode(&str); err != nil {
		fmt.Printf("Error on Deserializer: %s\n", err)
	}

	return str
}

func SerializerString(b string) []byte {
	var result bytes.Buffer
	enc := codec.NewEncoder(&result, &ch)
	if err := enc.Encode(b); err != nil {
		fmt.Printf("error on serializer: %s\n", err)
	}
	return result.Bytes()
}

// SerializerAccount serialize acc
func SerializerAccount(account acc.Account) []byte {
	var result bytes.Buffer
	enc := codec.NewEncoder(&result, &ch)
	if err := enc.Encode(account); err != nil {
		log.Printf("error on serialize acc: %s\n", err)
	}
	return result.Bytes()
}

// DeserializerAccount deselialize acc
func DeserializerAccount(b []byte) acc.Account {
	var a acc.Account
	dec := codec.NewDecoderBytes(b, &ch)
	if err := dec.Decode(&a); err != nil {
		log.Printf("error on deserialize acc : %s\n", err)
	}

	return a
}
