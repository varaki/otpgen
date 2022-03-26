package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
    clipb "github.com/atotto/clipboard"
)

func main() {

    var secret string
    var clipboard bool

    flag.StringVar(&secret, "secret", "", "token secret")
    flag.BoolVar(&clipboard, "clipboard", false, "copy generated code to clipboard")
    flag.Parse()

    code, err := generate(secret)
    if err != nil {
        panic(err)
    }

    if clipboard {
        err := clipb.WriteAll(code)
        if err != nil {
            panic(err)
        }
    }
}

func generate(secret string) (string, error) {
	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}

	fmt.Println(code)
	return code, err
}
