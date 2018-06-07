package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.66433,-74.39967,
	}
	m["Bell"] = Vertex{
		11.66433,-14.39967,
	}
	delete(m, "Bell")
	fmt.Println(m)
}
