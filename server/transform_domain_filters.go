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
