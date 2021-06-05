package files

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/disintegration/imaging"
)

var Count = len(GetFiles("static/photos"))

var (
	thumbs = "static/thumbnails/"
	photos = "static/photos/"
)

const RFC822 = "Mon, 02 Jan 2006 15:04:05 -0700"

func crop(i image.Image) image.Image {
	x := i.Bounds().Size().X / 2
	y := i.Bounds().Size().Y / 2
	rect := image.Rect(x-50, y-50, x+50, y+50)
	return imaging.Crop(i, rect)
}

func SavePhoto(f *multipart.File, h *multipart.FileHeader) {
	path := photos + fmt.Sprintf("%d.png", Count)
	if newfile, err := os.Create(path); err == nil {
		upload, _ := h.Open()
		orientation := imaging.AutoOrientation(true)
		if img, err := imaging.Decode(upload, orientation); err == nil {
			img = imaging.Fit(img, 1000, 1000, imaging.Lanczos)
			png.Encode(newfile, img)
			GenerateThumbnail(path, Count)
			Count += 1
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

func GetFileDates(path string) (dates []string) {
	if files, err := ioutil.ReadDir(path); err == nil {
		for _, file := range files {
			date := file.ModTime().Format(RFC822)
			fmt.Println(date)
			dates = append(dates, date)
		}
	}
	return dates
}

func GetFiles(path string) (names []string) {
	if files, err := ioutil.ReadDir(path); err == nil {
		for _, file := range files {
			names = append(names, file.Name())
		}
	}
	return names
}
