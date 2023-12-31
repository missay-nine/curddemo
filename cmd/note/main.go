package main

import (
	notedemo "curddemo/kitex_gen/notedemo/noteservice"
	"log"
)

func main() {
	svr := notedemo.NewServer(new(NoteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
