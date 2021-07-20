package main

const english_hello_prefix = "Hello, "
const chinese_hello_prefix = "哈摟, "

func Hello(name string, language string) string {
	var hello_prefix string

	//Check name is empty string or not
	if name == "" {
		name = "word"
	}

	//Determine hello_prefix language
	switch language {
	case "English":
		hello_prefix = english_hello_prefix
	case "Chinese":
		hello_prefix = chinese_hello_prefix
	default:
		hello_prefix = english_hello_prefix
		//Default language is English.
	}

	return hello_prefix + name
}

func main() {

}
