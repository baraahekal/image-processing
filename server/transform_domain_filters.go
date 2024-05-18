package main

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/draw"
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

func apply_nearest_neighbour_filter(inputImage image.Image, newWidth, newHeight int) image.Image {
	// Create a new RGBA image with the desired dimensions
	outputImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// Get the bounds of the input image
	bounds := inputImage.Bounds()

	// Calculate the scaling factors
	scaleX := float64(bounds.Dx()) / float64(newWidth)
	scaleY := float64(bounds.Dy()) / float64(newHeight)

	// Perform nearest neighbor interpolation
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			// Calculate the corresponding pixel in the original image
			px := int(float64(x) * scaleX)
			py := int(float64(y) * scaleY)

			// Get the color of the nearest pixel in the original image
			color := inputImage.At(px, py)

			// Set the color of the pixel in the output image
			outputImage.Set(x, y, color)
		}
	}

	return outputImage
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
func applyBilinearFilter(img image.Image, newWidth, newHeight int) image.Image {
	oldBounds := img.Bounds()
	newImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	xRatio := float64(oldBounds.Dx()) / float64(newWidth)
	yRatio := float64(oldBounds.Dy()) / float64(newHeight)

	for newY := 0; newY < newHeight; newY++ {
		for newX := 0; newX < newWidth; newX++ {
			x := float64(newX) * xRatio
			y := float64(newY) * yRatio
			x1 := int(x)
			y1 := int(y)
			x2 := min(oldBounds.Max.X-1, x1+1)
			y2 := min(oldBounds.Max.Y-1, y1+1)

			// 4 nearest neighbors
			a := img.At(x1, y1)
			b := img.At(x1, y2)
			c := img.At(x2, y1)
			d := img.At(x2, y2)

			// Interpolate in X direction
			col1 := interpolate(a, c, x-float64(x1))
			col2 := interpolate(b, d, x-float64(x1))

			// Interpolate in Y direction
			final := interpolate(col1, col2, y-float64(y1))

			newImage.Set(newX, newY, final)
		}
	}

	return newImage
}

func interpolate(c1, c2 color.Color, t float64) color.Color {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()

	return color.RGBA64{
		R: uint16(float64(r1)*(1-t) + float64(r2)*t),
		G: uint16(float64(g1)*(1-t) + float64(g2)*t),
		B: uint16(float64(b1)*(1-t) + float64(b2)*t),
		A: uint16(float64(a1)*(1-t) + float64(a2)*t),
	}
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

func calculate_color_histogram(img image.Image) ([256]int, [256]int, [256]int) {
	var histogramR, histogramG, histogramB [256]int
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			color := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			histogramR[color.R]++
			histogramG[color.G]++
			histogramB[color.B]++
		}
	}

	return histogramR, histogramG, histogramB
}

func applyBicubicFilter(img image.Image, newWidth, newHeight int) image.Image {
	oldBounds := img.Bounds()
	newImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	xRatio := float64(oldBounds.Dx()) / float64(newWidth)
	yRatio := float64(oldBounds.Dy()) / float64(newHeight)

	for newY := 0; newY < newHeight; newY++ {
		for newX := 0; newX < newWidth; newX++ {
			x := float64(newX) * xRatio
			y := float64(newY) * yRatio
			x1 := int(x)
			y1 := int(y)

			// Get the 16 nearest pixels
			pixels := make([]color.Color, 16)
			for i := -1; i <= 2; i++ {
				for j := -1; j <= 2; j++ {
					pixels[(i+1)*4+(j+1)] = img.At(clasmp(x1+i, 0, oldBounds.Max.X-1), clasmp(y1+j, 0, oldBounds.Max.Y-1))
				}
			}

			// Calculate the weights for the 16 pixels
			weights := make([]float64, 16)
			for i := 0; i < 16; i++ {
				dx := math.Abs(x - float64(x1+i%4))
				dy := math.Abs(y - float64(y1+i/4))
				weights[i] = bicubicWeight(dx) * bicubicWeight(dy)
			}

			// Calculate the new color for the pixel
			var r, g, b, a float64
			for i, pixel := range pixels {
				ri, gi, bi, ai := pixel.RGBA()
				r += weights[i] * float64(ri)
				g += weights[i] * float64(gi)
				b += weights[i] * float64(bi)
				a += weights[i] * float64(ai)
			}

			newImage.Set(newX, newY, color.RGBA{
				R: uint8(r / 0x101),
				G: uint8(g / 0x101),
				B: uint8(b / 0x101),
				A: uint8(a / 0x101),
			})
		}
	}

	return newImage
}

func clasmp(x, min, max int) int {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func bicubicWeight(x float64) float64 {
	// The coefficients 0.5 and -0.5 can be adjusted to change the sharpness of the interpolation
	if x < 1.0 {
		return 1.5*x*x*x - 2.5*x*x + 1.0
	} else if x < 2.0 {
		return -0.5*x*x*x + 2.5*x*x - 4.0*x + 2.0
	} else {
		return 0.0
	}
}

// cumulativeDistribution calculates the CDF of a grayscale image
func cumulativeDistribution(img image.Image) []float64 {
	hist := imaging.Histogram(img)
	cdf := make([]float64, 256)
	cdf[0] = hist[0]

	for i := 1; i < 256; i++ {
		cdf[i] = cdf[i-1] + hist[i]

	}
	for i := range cdf {
		cdf[i] /= cdf[255]
	}
	return cdf
}

func apply_histogram_equalization_filter(img image.Image) image.Image {
	histogramR, histogramG, histogramB := calculate_color_histogram(img)
	bounds := img.Bounds()
	totalPixels := float64((bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y))

	var cdfR, cdfG, cdfB [256]float64
	cdfR[0] = float64(histogramR[0]) / totalPixels
	cdfG[0] = float64(histogramG[0]) / totalPixels
	cdfB[0] = float64(histogramB[0]) / totalPixels

	for i := 1; i < 256; i++ {
		cdfR[i] = cdfR[i-1] + float64(histogramR[i])/totalPixels
		cdfG[i] = cdfG[i-1] + float64(histogramG[i])/totalPixels
		cdfB[i] = cdfB[i-1] + float64(histogramB[i])/totalPixels
	}

	equalizedImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			clr := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			newR := uint8(math.Round(cdfR[clr.R] * 255))
			newG := uint8(math.Round(cdfG[clr.G] * 255))
			newB := uint8(math.Round(cdfB[clr.B] * 255))
			equalizedImg.Set(x, y, color.RGBA{R: newR, G: newG, B: newB, A: 255})
		}
	}

	return equalizedImg
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
	templatePath := "asset/Sunset.jpeg"
	templateImg, err := readImage(templatePath)

	file, err := os.Open("asset/Sunset.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decode the image
	templateImg, _, err = image.Decode(file)
	if err != nil {
		panic(err)

	}

	println("A7AAAAA")
	inputCDF := cumulativeDistribution(img)
	templateCDF := cumulativeDistribution(templateImg)

	mapping := make([]uint8, 256)
	for i := range mapping {
		diff := make([]float64, 256)
		for j := range diff {
			diff[j] = abs(inputCDF[i] - templateCDF[j])
		}
		mapping[i] = uint8(argmin(diff))
	}

	bounds := img.Bounds()
	outputImg := image.NewRGBA(bounds)
	draw.Draw(outputImg, bounds, img, image.Point{}, draw.Src)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := outputImg.At(x, y).RGBA()
			outputImg.Set(x, y, color.RGBA{
				R: mapping[r>>8],
				G: mapping[g>>8],
				B: mapping[b>>8],
				A: uint8(a >> 8),
			})
		}
	}

	return outputImg

}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func argmin(x []float64) int {
	min := x[0]
	minIndex := 0
	for i, v := range x {
		if v < min {
			min = v
			minIndex = i
		}
	}
	return minIndex
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
