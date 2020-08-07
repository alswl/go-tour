package main

func main() {
	// arguments evaluated immediately,
	words := "world"
	defer println(words)
	words = "changed"

	println("hello")
}
