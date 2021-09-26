package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func main() {
	fmt.Print("Enter Password: ")
	in, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Missing password")

		return
	}

	h := sha1.New()
	io.WriteString(h, string(in))
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
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	found := bytes.Contains(body, []byte(suffix))

	if found {
		fmt.Println("Oh no — pwned!")
	} else {
		fmt.Println("Good news — no pwnage found!")
	}
}
