// if google auth fails please check the route of GOOGLE_APPLICATION_CREDENTIALS env variable and main.go/110 according to your folder structure

package main

import (
	"bytes"
	"os"

	// "context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	// "net/http"
)

type ReadCloser struct {
	*bytes.Reader
	Body io.ReadCloser
}

func (rc *ReadCloser) WriteTo(w io.Writer) (n int64, err error) {
	return rc.Reader.WriteTo(w)
}

func (rc ReadCloser) Close() error {
	return nil
}

func (rc *ReadCloser) UnmarshalBinary(b []byte) error {
	// rc.Body = b
	rc.Reader = bytes.NewReader(b)
	fmt.Println("======= unmarshalled and sent to gob: =========")
	fmt.Println(string(b))
	return nil
}

func (rc *ReadCloser) MarshalBinary() ([]byte, error) {
	if rc.Body != nil {
		b, err := ioutil.ReadAll(rc.Body)
		rc.Body.Close()
		rc.Reader = bytes.NewReader(b)
		return b, err
	}
	return nil, nil
}

func main() {
	gob.Register(ReadCloser{})
	gob.Register(elliptic.P256())
	gob.Register(ecdsa.PublicKey{})
	gob.Register(rsa.PublicKey{})

	dat, err := os.ReadFile("./data.txt")
	if err != nil {
		panic(err)
	}

	obj := &http.Response{}
	dec := gob.NewDecoder(bytes.NewBuffer(dat))
	err = dec.Decode(obj)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(obj.Body)
	obj.Body.Close()
	fmt.Println("========= received from gob ========")
	fmt.Println(string(b))
}
