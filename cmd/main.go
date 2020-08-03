package main

import (
	"flag"
	"log"
	"tidraw/pkg/model"
	"tidraw/pkg/pixel"
)

func main() {
	host := flag.String("host", "localhost", "Specifies the host to connect tiDB")
	port := flag.String("port", "4000", "Specifies the port to connect tiDB")
	file := flag.String("file", "assets/butterfly.jpg", "Specifies the image path for drawing on keyviz")
	flag.Parse()

	err := model.InitDB(*host, *port)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		err = pixel.DrawPicture(*file)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
