package avatar

import (
	"crypto/md5"
	"image"
	"image/color"
	"image/draw"
)

type Identicon struct {
	Name     string
	hash     []byte
	grid     []grid
	pixelMap []pixelMap
	Image    *image.RGBA
	color    colorRGB
}

type colorRGB struct {
	r uint8
	g uint8
	b uint8
}

type grid struct {
	value int
	index int
}

type pixelMap struct {
	topLeft     axis
	bottomRight axis
}

type axis struct {
	x int
	y int
}

func NewIdenticon(name string) *Identicon {
	identicon := new(Identicon)
	identicon.Name = name
	identicon.hash = identicon.hashInput()
	identicon.grid = identicon.buildGrid()
	identicon.grid = identicon.filterOddSquares()
	identicon.pixelMap = identicon.buildPixelMap()
	identicon.color = colorRGB{r: identicon.hash[0], g: identicon.hash[1], b: identicon.hash[2]}
	identicon.Image = identicon.drawImage()

	return identicon
}

func (i *Identicon) hashInput() []byte {
	hash := md5.Sum([]byte(i.Name))
	return hash[:]
}

func (i *Identicon) buildGrid() []grid {
	slices := chunkSlice(i.hash, 3)
	mirroredSlices := mirrorSlices(slices)

	return flattenSlice(mirroredSlices)
}

func (i *Identicon) filterOddSquares() []grid {
	var result []grid
	for _, g := range i.grid {
		if g.value%2 == 0 {
			result = append(result, grid{value: g.value, index: g.index})
		}
	}

	return result
}

func (i *Identicon) buildPixelMap() []pixelMap {
	var pixel_map []pixelMap
	for _, g := range i.grid {
		horizontal := (g.index % 5) * 50
		vertical := (g.index / 5) * 50
		pixel_map = append(
			pixel_map,
			pixelMap{
				topLeft:     axis{x: horizontal, y: vertical},
				bottomRight: axis{x: horizontal + 50, y: vertical + 50},
			},
		)
	}

	return pixel_map
}

func (i *Identicon) drawImage() *image.RGBA {

	col := color.RGBA{i.color.r, i.color.g, i.color.b, 255}

	myimage := image.NewRGBA(image.Rect(0, 0, 250, 250))
	draw.Draw(myimage, myimage.Bounds(), image.White, image.Point{}, draw.Src)

	for _, p := range i.pixelMap {

		draw.Draw(myimage, image.Rect(p.topLeft.x, p.topLeft.y, p.bottomRight.x, p.bottomRight.y),
			&image.Uniform{col}, image.Point{}, draw.Src)
	}

	return myimage
}

func flattenSlice(slice [][]byte) []grid {

	var flattenSlice []grid
	cont := 0
	for _, s1 := range slice {
		for i := 0; i < len(s1); i++ {
			flattenSlice = append(flattenSlice, grid{value: int(s1[i]), index: cont})
			cont++

		}
	}
	return flattenSlice
}

func mirrorSlices(slices [][]byte) [][]byte {
	var mirroredSlices [][]byte

	for _, s := range slices {
		newSlice := make([]byte, 5)
		copy(newSlice, s)
		newSlice[3] = s[1]
		newSlice[4] = s[0]
		mirroredSlices = append(mirroredSlices, newSlice)
	}
	return mirroredSlices
}

func chunkSlice(slice []byte, chunkSize int) [][]byte {
	var chunks [][]byte

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		if end%chunkSize == 0 {
			chunks = append(chunks, slice[i:end])
		}

	}

	return chunks
}
