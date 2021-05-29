package files

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"os"
)

var (
	thumbs = "static/thumbnails/"
	photos = "static/photos/"
	count  = len(GetFiles("static/photos"))
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func crop(i image.Image) image.Image {
	x := i.Bounds().Size().X / 2
	y := i.Bounds().Size().Y / 2
	rect := image.Rect(x-50, y-50, x+50, y+50)
	return i.(SubImager).SubImage(rect)
}

func SavePhoto(f *multipart.File, h *multipart.FileHeader, rotate bool) {
	path := photos + fmt.Sprintf("%d.png", count)
	if newfile, err := os.Create(path); err == nil {
		upload, _ := h.Open()
		if img, _, err := image.Decode(upload); err == nil {
			if rotate {
				img = imaging.Rotate270(img)
			}
			img = imaging.Fit(img, 1000, 1000, imaging.Lanczos)
			png.Encode(newfile, img)
			GenerateThumbnail(path, count)
		}
	}
}

func GenerateThumbnail(path string, index int) {
	if f, err := os.Open(path); err == nil {
		img, _, err := image.Decode(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		path := thumbs + fmt.Sprintf("%d.png", index)
		f, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		cropped := crop(img)
		png.Encode(f, cropped)
	}
}

func GenerateThumbnails() {
	for index, file := range GetFiles("static/photos") {
		thumb := thumbs + fmt.Sprintf("%d.png", index)
		photo := photos + file
		if _, err := os.Stat(thumb); os.IsNotExist(err) {
			fmt.Println("Generating thumbnail for", file)
			GenerateThumbnail(photo, index)
		}
	}
}

func GetFiles(path string) (names []string) {
	if files, err := ioutil.ReadDir(path); err == nil {
		for _, file := range files {
			names = append(names, file.Name())
		}
	}
	return names
}
