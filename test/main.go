package main

import (
	"fatfox-single-sign-on/utils"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	res, err := utils.GenerateJWT(uuid.New(), time.Now().Add(time.Hour*24*30))
	if err != nil {
		panic(err)
		return
	}

	c, err := utils.ParseJWT(res)
	if err != nil {
		// handle error
		panic(err)
		return
	}

	fmt.Println(c)
}
