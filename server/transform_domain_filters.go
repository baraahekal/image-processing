package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"os"
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
func readImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
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
		var pR, pG, pB, pA [4]float64
		for i := 0; i < 4; i++ {
			r0, g0, b0, a0 := pixels[0][i].RGBA()
			r1, g1, b1, a1 := pixels[1][i].RGBA()
			r2, g2, b2, a2 := pixels[2][i].RGBA()
			r3, g3, b3, a3 := pixels[3][i].RGBA()
			pR[i] = cubicHermite(r0, r1, r2, r3, fx)
			pG[i] = cubicHermite(g0, g1, g2, g3, fx)
			pB[i] = cubicHermite(b0, b1, b2, b3, fx)
			pA[i] = cubicHermite(a0, a1, a2, a3, fx)
		}
		r += cubicHermite(uint32(pR[0]), uint32(pR[1]), uint32(pR[2]), uint32(pR[3]), fy)
		g += cubicHermite(uint32(pG[0]), uint32(pG[1]), uint32(pG[2]), uint32(pG[3]), fy)
		b += cubicHermite(uint32(pB[0]), uint32(pB[1]), uint32(pB[2]), uint32(pB[3]), fy)
		a += cubicHermite(uint32(pA[0]), uint32(pA[1]), uint32(pA[2]), uint32(pA[3]), fy)
	}
	r = math.Min(math.Max(r/257, 0), 255) // Convert from uint32 to uint8 and clamp
	g = math.Min(math.Max(g/257, 0), 255)
	b = math.Min(math.Max(b/257, 0), 255)
	a = math.Min(math.Max(a/257, 0), 255)
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}
func cubicHermite(v0, v1, v2, v3 uint32, t float64) float64 {
	P0 := float64(v1)
	P1 := float64(v2)
	M0 := 0.5 * (float64(v2) - float64(v0))
	M1 := 0.5 * (float64(v3) - float64(v1))
	t2 := t * t
	t3 := t2 * t
	return (2*t3-3*t2+1)*P0 + (t3-2*t2+t)*M0 + (-2*t3+3*t2)*P1 + (t3-t2)*M1
}
func histeqChannel(img *image.RGBA, channelIndex int) *image.Gray {
	bounds := img.Bounds()
	hist := make([]int, 256)
	cdf := make([]int, 256)

	// Compute histogram
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := img.RGBAAt(x, y)
			value := uint8(0)
			switch channelIndex {
			case 0:
				value = c.R
			case 1:
				value = c.G
			case 2:
				value = c.B
			}
			hist[value]++
		}
	}
	// Compute CDF
	sum := 0
	for i := 0; i < 256; i++ {
		sum += hist[i]
		cdf[i] = sum
	}
	// Normalize CDF
	max := float64(bounds.Dx() * bounds.Dy())
	cdfMin := cdf[0]
	for i := 0; i < 256; i++ {
		cdf[i] = int(float64(cdf[i]-cdfMin) * 255 / (max - float64(cdfMin)))
	}
	// Apply equalization
	equalized := image.NewGray(bounds)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := img.RGBAAt(x, y)
			value := uint8(0)
			switch channelIndex {
			case 0:
				value = uint8(cdf[c.R])
			case 1:
				value = uint8(cdf[c.G])
			case 2:
				value = uint8(cdf[c.B])
			}
			equalized.SetGray(x, y, color.Gray{Y: value})
		}
	}
	return equalized
}
func splitChannels(img image.Image) (r, g, b *image.Gray) {
	bounds := img.Bounds()
	r = image.NewGray(bounds)
	g = image.NewGray(bounds)
	b = image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			col := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			r.Set(x, y, color.Gray{Y: col.R})
			g.Set(x, y, color.Gray{Y: col.G})
			b.Set(x, y, color.Gray{Y: col.B})
		}
	}

	return r, g, b
}

// cumulativeDistribution calculates the CDF of a grayscale image
func cumulativeDistribution(img *image.Gray) ([]float64, []int) {
	hist := make([]int, 256)
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			hist[img.GrayAt(x, y).Y]++
		}
	}

	cdf := make([]float64, 256)
	cumulativeSum := 0
	totalPixels := float64(bounds.Dx() * bounds.Dy())
	for i := 0; i < 256; i++ {
		cumulativeSum += hist[i]
		cdf[i] = float64(cumulativeSum) / totalPixels
	}

	return cdf, hist
}

// adjustCDF adjusts the CDF by inserting zeros and ones as needed
func adjustCDF(cdf []float64, hist []int) []float64 {
	adjustedCDF := make([]float64, 256)
	for i := 0; i < hist[0]; i++ {
		adjustedCDF[i] = 0
	}
	copy(adjustedCDF[hist[0]:], cdf)
	for i := len(cdf); i < 256; i++ {
		adjustedCDF[i] = 1
	}
	return adjustedCDF
}

// histogramMatching performs histogram matching on the input image using the template CDF
func histogramMatching(img, template *image.Gray) *image.Gray {
	cdfInput, binsInput := cumulativeDistribution(img)
	cdfTemplate, binsTemplate := cumulativeDistribution(template)

	cdfInput = adjustCDF(cdfInput, binsInput)
	cdfTemplate = adjustCDF(cdfTemplate, binsTemplate)

	bounds := img.Bounds()
	resultImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.GrayAt(x, y).Y
			newPixel := uint8(math.Round(interp(cdfInput[originalPixel], cdfTemplate) * 255))
			resultImg.SetGray(x, y, color.Gray{Y: newPixel})
		}
	}
	return resultImg
}

// interp performs linear interpolation
func interp(value float64, cdfTemplate []float64) float64 {
	for i := 0; i < len(cdfTemplate)-1; i++ {
		if value <= cdfTemplate[i+1] {
			t := (value - cdfTemplate[i]) / (cdfTemplate[i+1] - cdfTemplate[i])
			return float64(i) + t
		}
	}
	return float64(len(cdfTemplate) - 1)
}
func apply_histogram_equalization_filter(img image.Image) image.Image {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	// Perform histogram equalization on each channel
	r := histeqChannel(rgba, 0)
	g := histeqChannel(rgba, 1)
	b := histeqChannel(rgba, 2)

	// Create a new image with the equalized channels
	equalized := image.NewRGBA(bounds)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rColor := color.RGBAModel.Convert(r.At(x, y)).(color.RGBA)
			gColor := color.RGBAModel.Convert(g.At(x, y)).(color.RGBA)
			bColor := color.RGBAModel.Convert(b.At(x, y)).(color.RGBA)
			equalized.Set(x, y, color.RGBA{R: rColor.R, G: gColor.G, B: bColor.B, A: 255})
		}
	}
	return equalized
}
func apply_histogram_specification_filter(img image.Image) image.Image {
	templatePath := "asset/Sunset.jpeg"
	templateImg, err := readImage(templatePath)
	if err != nil {
		fmt.Println("Error reading template image:", err)
		return nil
	}
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	// Split the input and template images into R, G, B channels
	print("انا عديت ")
	inputR, inputG, inputB := splitChannels(img)
	templateR, templateG, templateB := splitChannels(templateImg)
	print("انا عديت السبليت")
	// Perform histogram matching for each channel
	outputR := histogramMatching(inputR, templateR)
	outputG := histogramMatching(inputG, templateG)
	outputB := histogramMatching(inputB, templateB)
	print("انا عديت الماتش")
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r := outputR.GrayAt(x, y).Y
			g := outputG.GrayAt(x, y).Y
			b := outputB.GrayAt(x, y).Y
			newImg.Set(x, y, color.RGBA{
				R: r,
				G: g,
				B: b,
				A: 255,
			})
		}
	}
	print("انا عديت مش عارف")
	return newImg
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
