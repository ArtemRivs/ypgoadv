package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Id         int
	Name       string
	Occupation string
}

func main() {

	u1 := User{1, "John Fruschiante", "Guitarist"}
	json_data, err := json.Marshal(u1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("json_data:", json_data)
	fmt.Println("string json_data:", string(json_data))

	fmt.Println("USER:", os.Getenv("USER"))
	fmt.Println("USER_ABUSER:", os.Getenv("USER_ABUSER"))

}
