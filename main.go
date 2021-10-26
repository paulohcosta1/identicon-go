package main

import (
	"bufio"
	"fmt"
	"os"
	"paulo/identicon/avatar"
	"paulo/identicon/files"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digit your name: ")
	name, _ := reader.ReadString('\n')
	identicon := avatar.NewIdenticon(name)
	files.SaveImagePNG(identicon.Name, identicon.Image)

}
