package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)



/* */ 
type Person struct {
	firstname string 
	lastname string
	age int
}

func main() {
	fmt.Printf("hello, world\n")

	person := Person{
		firstname: "fletch",
		lastname: "uknowit",
	}

	fmt.Printf(person.firstname)
	fmt.Printf(" is %v age", person.age)
	fmt.Printf("\n%v", getIt())


	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close() 

	log.Println(string(body))

	// resp, err := http.Get("https://google.com")
	
	// if (err != nil) {
	// 	log.Fatalln("bad stuff")
	// }

	// defer resp.Body.Close()
	// log.Println(resp)

	// body, err = ioutil.ReadAll(resp.Body)

	// if (err != nil) {
	// 	log.Fatalln(err)
	// }

	log.Println("interesting")
}


func getIt() int { return 1101 }

