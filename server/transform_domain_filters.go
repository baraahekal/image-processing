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

// Perform 1D FFT
func fft1D(a []complex128) {
	n := len(a)
	if n <= 1 {
		return
	}

	// Divide
	odd := make([]complex128, n/2)
	even := make([]complex128, n/2)
	for i := range even {
		even[i] = a[i*2]
		odd[i] = a[i*2+1]
	}

	// Conquer
	fft1D(even)
	fft1D(odd)

	// Combine
	for k := 0; k < n/2; k++ {
		t := cmplx.Exp(-2i*math.Pi*complex(float64(k)/float64(n), 0)) * odd[k]
		a[k] = even[k] + t
		a[k+n/2] = even[k] - t
	}
}

// Perform 2D FFT
func fft2D(input [][]float64) [][]complex128 {
	height := len(input)
	width := len(input[0])
	output := make([][]complex128, height)
	for i := range output {
		output[i] = make([]complex128, width)
		for j := range output[i] {
			output[i][j] = complex(input[i][j], 0)
		}
	}

	// Perform FFT on each row
	for i := 0; i < height; i++ {
		fft1D(output[i])
	}

	// Perform FFT on each column
	column := make([]complex128, height)
	for j := 0; j < width; j++ {
		for i := 0; i < height; i++ {
			column[i] = output[i][j]
		}
		fft1D(column)
		for i := 0; i < height; i++ {
			output[i][j] = column[i]
		}
	}

	return output
}

// Shift the zero-frequency component to the center of the spectrum
func shiftDFT(input [][]complex128) [][]complex128 {
	height := len(input)
	width := len(input[0])
	output := make([][]complex128, height)
	for i := range output {
		output[i] = make([]complex128, width)
	}

	for u := 0; u < height; u++ {
		for v := 0; v < width; v++ {
			output[u][v] = input[(u+height/2)%height][(v+width/2)%width]
		}
	}
	return output
}

// Compute the logarithmic spectrum for visualization
func computeSpectrum(dftShifted [][]complex128, k float64) [][]float64 {
	height := len(dftShifted)
	width := len(dftShifted[0])
	spectrum := make([][]float64, height)
	for i := range spectrum {
		spectrum[i] = make([]float64, width)
	}

	for u := 0; u < height; u++ {
		for v := 0; v < width; v++ {
			spectrum[u][v] = k * math.Log(1+cmplx.Abs(dftShifted[u][v]))
		}
	}
	return spectrum
}

// Normalize the spectrum to the range [0, 255] for visualization
func normalizeSpectrum(spectrum [][]float64) [][]float64 {
	height := len(spectrum)
	width := len(spectrum[0])
	minVal := spectrum[0][0]
	maxVal := spectrum[0][0]

	// Find the min and max values
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if spectrum[y][x] < minVal {
				minVal = spectrum[y][x]
			}
			if spectrum[y][x] > maxVal {
				maxVal = spectrum[y][x]
			}
		}
	}

	// Normalize the spectrum
	normalized := make([][]float64, height)
	for y := 0; y < height; y++ {
		normalized[y] = make([]float64, width)
		for x := 0; x < width; x++ {
			normalized[y][x] = 255 * (spectrum[y][x] - minVal) / (maxVal - minVal)
		}
	}
	return normalized
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

// Apply the Fourier Transform filter
func apply_fourier_transform_filter(img image.Image) image.Image {
	// Step 1: Convert to grayscale
	grayImage := toGrayscale(img)
	// Step 2: Perform 2D FFT
	fftResult := fft2D(grayImage)
	// Step 3: Shift the zero-frequency component to the center
	fftShifted := shiftDFT(fftResult)
	// Step 4: Compute the spectrum for visualization
	spectrum := computeSpectrum(fftShifted, 20)
	// Step 5: Normalize the spectrum
	normalizedSpectrum := normalizeSpectrum(spectrum)
	// Step 6: Convert the normalized spectrum to a grayscale image
	filteredImage := floatArrayToGrayImage(normalizedSpectrum)

	return filteredImage
}

func apply_interpolation_filter(img image.Image) image.Image {
	// Implement the filter
	// عدل عليه وظبطه وخليه لل3 واتاكد انه صح عشان مش لاقي حاجة اتاكد منها معلش
	return img
}
