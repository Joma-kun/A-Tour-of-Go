package main

import "fmt"

type Vertex struct {
	Lat, Lonf float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -7439967,
	}
	fmt.Println(m["Bell Labs"])
}
