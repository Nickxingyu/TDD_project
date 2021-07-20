package main

const english_hello_prefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "word"
	}
	return english_hello_prefix + name
}

func main() {

}
