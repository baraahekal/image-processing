package main

import (
	"image"
	"image/color"
	"math"
	"math/cmplx"
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
func apply_histogram_specification_filter(img image.Image) image.Image {
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
func toGrayscale(img image.Image) [][]float64 {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	gray := make([][]float64, height)
	for y := 0; y < height; y++ {
		gray[y] = make([]float64, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray[y][x] = 0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8)
		}
	}
	return gray
}

// Perform the 2D DFT
func dft2D(input [][]float64) [][]complex128 {
	height := len(input)
	width := len(input[0])
	output := make([][]complex128, height)
	for u := range output {
		output[u] = make([]complex128, width)
	}

	for u := 0; u < height; u++ {
		for v := 0; v < width; v++ {
			var sum complex128
			for x := 0; x < height; x++ {
				for y := 0; y < width; y++ {
					angle := 2.0 * math.Pi * (float64(u*x)/float64(height) + float64(v*y)/float64(width))
					sum += complex(input[x][y], 0) * cmplx.Exp(-1i*complex(angle, 0))
				}
			}
			output[u][v] = sum
		}
	}
	return output
}

// Perform the inverse 2D DFT
func idft2D(input [][]complex128) [][]float64 {
	height := len(input)
	width := len(input[0])
	output := make([][]float64, height)
	for x := range output {
		output[x] = make([]float64, width)
	}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			var sum complex128
			for u := 0; u < height; u++ {
				for v := 0; v < width; v++ {
					angle := 2.0 * math.Pi * (float64(u*x)/float64(height) + float64(v*y)/float64(width))
					sum += input[u][v] * cmplx.Exp(1i*complex(angle, 0))
				}
			}
			output[x][y] = real(sum) / float64(height*width)
		}
	}
	return output
}

// Apply a low-pass filter in the frequency domain
func applyLowPassFilter(fft [][]complex128, cutoff float64) {
	height := len(fft)
	width := len(fft[0])
	for u := 0; u < height; u++ {
		for v := 0; v < width; v++ {
			distance := math.Sqrt(math.Pow(float64(u-height/2), 2) + math.Pow(float64(v-width/2), 2))
			if distance > cutoff {
				fft[u][v] = 0
			}
		}
	}
}

// Convert 2D float array to grayscale image
func floatArrayToGrayImage(input [][]float64) *image.Gray {
	height := len(input)
	width := len(input[0])
	img := image.NewGray(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			value := uint8(math.Min(math.Max(input[y][x], 0), 255))
			img.SetGray(x, y, color.Gray{Y: value})
		}
	}
	return img
}

func apply_fourier_transform_filter(img image.Image) image.Image {
	// Step 1: Convert to grayscale
	grayImage := toGrayscale(img)
	println("1")
	// Step 2: Perform 2D DFT
	fftResult := dft2D(grayImage)
	println("2")

	// Step 3: Apply a low-pass filter
	cutoff := 30.0 // Adjust cutoff frequency as needed
	applyLowPassFilter(fftResult, cutoff)
	println("3")

	// Step 4: Perform inverse 2D DFT
	filteredImageArray := idft2D(fftResult)
	println("4")

	// Step 5: Convert back to image
	filteredImage := floatArrayToGrayImage(filteredImageArray)
	println("5")

	return filteredImage
}

func apply_interpolation_filter(img image.Image) image.Image {
	// Implement the filter
	// عدل عليه وظبطه وخليه لل3 واتاكد انه صح عشان مش لاقي حاجة اتاكد منها معلش
	return img
}
