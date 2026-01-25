package main

import "log"

func main() {
	cfg := config{
		port: ":8080",
		db: dbConfig{},
	}

	api := api{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		log.Fatalln("Server failed to start: ", err)
	}
}
