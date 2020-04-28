package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Missing password")

		return
	}

	password := args[0]
	h := sha1.New()
	io.WriteString(h, password)
	hash := strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))

	prefix := hash[0:5]
	suffix := hash[5:]

	resp, err := http.Get(fmt.Sprintf("https://api.pwnedpasswords.com/range/%s", prefix))
	if err != nil {
		fmt.Println(err)

		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	found := bytes.Contains(body, []byte(suffix))

	if found == true {
		fmt.Println("Oh no — pwned!")
	} else {
		fmt.Println("Good news — no pwnage found!")
	}
}
