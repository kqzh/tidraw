package pixel

import (
	"fmt"
	"log"
	"tidraw/pkg/model"
	"time"

	"github.com/disintegration/imaging"
)

const drawPeriod = time.Minute

func DrawPicture(file string) error {
	defer time.Sleep(10 * time.Second)

	src, err := imaging.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open image: %v", err)
	}

	// Resize the cropped image to height = 128px preserving the aspect ratio.
	src = imaging.Resize(src, 0, model.PictureHeight, imaging.Lanczos)

	// Create a grayscale version of the image with higher contrast and sharpness.
	img := imaging.Grayscale(src)
	img = imaging.AdjustContrast(img, 20)
	img = imaging.Sharpen(img, 2)

	drawTicker := time.NewTicker(drawPeriod)
	defer drawTicker.Stop()

	for i := 0; i < img.Bounds().Dx(); i++ {
		<-drawTicker.C
		log.Println("drawing success", i)
		query := ""
		for j := 0; j < model.PictureHeight; j++ {
			r, _, _, _ := img.At(i, model.PictureHeight-j-1).RGBA()
			gray := uint8(r >> 8)
			query += fmt.Sprintf(model.SelectTableSql, j, gray)
		}
		_, err = model.DB.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}
