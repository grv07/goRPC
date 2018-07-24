package main

import (
	"log"
	"net/rpc"
)

type ToDo struct {
	Title, Status string
}

func main() {
	var reply ToDo
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	client.Call("Task.MakeToDo", ToDo{Title: "test 1", Status: "test 1"}, &reply)
	log.Printf("data: %v", reply)
	client.Call("Task.MakeToDo", ToDo{Title: "test 2", Status: "test 2"}, &reply)
	log.Printf("data: %v", reply)
	reply = ToDo{}
	err = client.Call("Task.GetToDo", "test 2", &reply)
	if err != nil {
		log.Print(err)
	}
	log.Printf("search data: %v", reply)
}
