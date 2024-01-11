package main

//"net/http"

func main() {

	server := newAPIServer(":3000")

	server.Run()
}
