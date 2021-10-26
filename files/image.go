package files

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func SaveImagePNG(name string, image *image.RGBA) {
	filename := strings.TrimSpace(name) + ".png"
	pwd, err := os.Getwd()

	if err != nil {
		panic(err.Error())
	}
	myfile, err := os.Create(filepath.Join(fmt.Sprintf("%s/images/", pwd), filepath.Base(filename)))
	if err != nil {
		panic(err.Error())
	}
	defer myfile.Close()

	png.Encode(myfile, image)
}
