package main

import (
	"image"
	"image/color"
	"math"
)

/*
img := image.NewRGBA(image.Rect(0, 0, 4, 4))

	img.Set(0, 0, color.RGBA{255, 0, 0, 255})
	img.Set(1, 1, color.RGBA{0, 255, 0, 255})
	img.Set(2, 2, color.RGBA{0, 0, 255, 255})
	img.Set(3, 3, color.RGBA{255, 255, 255, 255})

	filteredImg := applyNearestNeighborFilter(img)
	بتتنده كده
*/
func applyNearestNeighborFilter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	outputImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			origX := int(math.Round(float64(x) / 2.0))
			origY := int(math.Round(float64(y) / 2.0))
			nearestColor := img.At(origX, origY)
			outputImg.Set(x, y, nearestColor)
		}
	}
	return outputImg
}

/*
img := image.NewRGBA(image.Rect(0, 0, 4, 4))

		img.Set(0, 0, color.RGBA{255, 0, 0, 255})
		img.Set(1, 1, color.RGBA{0, 255, 0, 255})
		img.Set(2, 2, color.RGBA{0, 0, 255, 255})
		img.Set(3, 3, color.RGBA{255, 255, 255, 255})

		filteredImg := applyBilinearFilter(img)
	 بتتنده كده
*/
func applyBilinearFilter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	outputImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			origX := float64(x) / 2.0
			origY := float64(y) / 2.0
			x1 := int(math.Floor(origX))
			y1 := int(math.Floor(origY))
			x2 := x1 + 1
			y2 := y1 + 1
			if x2 >= bounds.Max.X {
				x2 = bounds.Max.X - 1
			}
			if y2 >= bounds.Max.Y {
				y2 = bounds.Max.Y - 1
			}
			fx := origX - float64(x1)
			fy := origY - float64(y1)
			c00 := img.At(x1, y1)
			c01 := img.At(x1, y2)
			c10 := img.At(x2, y1)
			c11 := img.At(x2, y2)
			r := bilinearInterpolation(c00, c01, c10, c11, fx, fy)
			outputImg.Set(x, y, r)
		}
	}
	return outputImg
}
func bilinearInterpolation(c00, c01, c10, c11 color.Color, fx, fy float64) color.Color {
	r00, g00, b00, a00 := c00.RGBA()
	r01, g01, b01, a01 := c01.RGBA()
	r10, g10, b10, a10 := c10.RGBA()
	r11, g11, b11, a11 := c11.RGBA()
	r := uint8(float64(r00)*(1-fx)*(1-fy) + float64(r01)*(1-fx)*fy + float64(r10)*fx*(1-fy) + float64(r11)*fx*fy)
	g := uint8(float64(g00)*(1-fx)*(1-fy) + float64(g01)*(1-fx)*fy + float64(g10)*fx*(1-fy) + float64(g11)*fx*fy)
	b := uint8(float64(b00)*(1-fx)*(1-fy) + float64(b01)*(1-fx)*fy + float64(b10)*fx*(1-fy) + float64(b11)*fx*fy)
	a := uint8(float64(a00)*(1-fx)*(1-fy) + float64(a01)*(1-fx)*fy + float64(a10)*fx*(1-fy) + float64(a11)*fx*fy)
	return color.RGBA{r, g, b, a}
}

/*
img := image.NewRGBA(image.Rect(0, 0, 4, 4))

	img.Set(0, 0, color.RGBA{255, 0, 0, 255})
	img.Set(1, 1, color.RGBA{0, 255, 0, 255})
	img.Set(2, 2, color.RGBA{0, 0, 255, 255})
	img.Set(3, 3, color.RGBA{255, 255, 255, 255})
	filteredImg := applyBicubicFilter(img)

بتتنده كده
*/
func applyBicubicFilter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	outputImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			origX := float64(x) / 2.0
			origY := float64(y) / 2.0
			x1 := int(math.Floor(origX)) - 1
			y1 := int(math.Floor(origY)) - 1
			x2 := x1 + 1
			y2 := y1 + 1
			if x1 < 0 {
				x1 = 0
			}
			if y1 < 0 {
				y1 = 0
			}
			if x2 >= bounds.Max.X {
				x2 = bounds.Max.X - 1
			}
			if y2 >= bounds.Max.Y {
				y2 = bounds.Max.Y - 1
			}
			fx := origX - float64(x1)
			fy := origY - float64(y1)
			var pixels [4][4]color.RGBA
			for j := 0; j < 4; j++ {
				for i := 0; i < 4; i++ {
					pixels[i][j] = img.At(x1+i, y1+j).(color.RGBA)
				}
			}
			r := bicubicInterpolation(pixels, fx, fy)
			outputImg.Set(x, y, r)
		}
	}
	return outputImg
}
func bicubicInterpolation(pixels [4][4]color.RGBA, fx, fy float64) color.RGBA {
	var r, g, b, a float64
	for j := 0; j < 4; j++ {
		var p [4]float64
		for i := 0; i < 4; i++ {
			p[i] = cubicHermite(pixels[0][i].RGBA(), pixels[1][i].RGBA(), pixels[2][i].RGBA(), pixels[3][i].RGBA(), fx)
		}
		h := cubicHermite(p[0], p[1], p[2], p[3], fy)
		r += h * float64(pixels[j][0].R)
		g += h * float64(pixels[j][0].G)
		b += h * float64(pixels[j][0].B)
		a += h * float64(pixels[j][0].A)
	}
	r = math.Min(math.Max(r, 0), 255)
	g = math.Min(math.Max(g, 0), 255)
	b = math.Min(math.Max(b, 0), 255)
	a = math.Min(math.Max(a, 0), 255)
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}
func cubicHermite(n0, n1, n2, n3 float64, t float64) float64 {
	return n1 + 0.5*t*(n2-n0+(2.0*n0-5.0*n1+4.0*n2-n3+t*(3.0*(n1-n2)+n3-n0)))
}
func apply_histogram_equalization_filter(img image.Image) image.Image {
	// Implement the filter

	return img
}
func calculateHistogram(img image.Image) [256]uint64 {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	histogram := [256]uint64{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			pixelValue := int(r >> 8)
			histogram[pixelValue]++
		}
	}
	return histogram
}
func calculateCDF(histogram [256]uint64) [256]uint64 {
	var cdf [256]uint64
	cdf[0] = histogram[0]
	for i := 1; i < 256; i++ {
		cdf[i] = cdf[i-1] + histogram[i]
	}
	return cdf
}
func histogramSpecification(img image.Image) image.Image {
	inputGray := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			pixelValue := uint8(r >> 8)
			inputGray.SetGray(x, y, color.Gray{pixelValue})
		}
	}
	referenceGray := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			pixelValue := uint8(r >> 8)
			referenceGray.SetGray(x, y, color.Gray{pixelValue})
		}
	}
	//referenceHistogram := calculateHistogram(referenceGray)
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	imageHistogram := calculateHistogram(inputGray)
	imageCDF := calculateCDF(imageHistogram)
	//referenceCDF := calculateCDF(referenceHistogram)
	mapping := [256]uint8{}
	for i := 0; i < 256; i++ {
		mapping[i] = uint8((255 * (imageCDF[i] - imageCDF[0])) / (imageCDF[255] - imageCDF[0]))
	}
	outputImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, _, _, a := img.At(x, y).RGBA()
			pixelValue := mapping[int(r>>8)]
			outputImg.Set(x, y, color.RGBA{pixelValue, pixelValue, pixelValue, uint8(a >> 8)})
		}
	}
	return outputImg
}
func apply_histogram_specification_filter(img image.Image) image.Image {
	// Implement the filter
	return img
}

func apply_fourier_transform_filter(img image.Image) image.Image {
	// Implement the filter
	return img
}

func apply_interpolation_filter(img image.Image) image.Image {
	// Implement the filter
	// عدل عليه وظبطه وخليه لل3 واتاكد انه صح عشان مش لاقي حاجة اتاكد منها معلش
	return img
}
