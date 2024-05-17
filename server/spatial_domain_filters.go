package main

import (
	_ "fmt"
	"image"
	"image/color"
	_ "image/color"
	_ "image/jpeg" // or "image/png" depending on your image type
	"math"
	"math/rand"
	_ "os"
	"sort"
)

// TEST FUNCTION
func apply_median_filter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)

	kernelSize := 9
	h := kernelSize / 2
	w := kernelSize / 2

	for y := h; y < height-h; y++ {
		for x := w; x < width-w; x++ {
			var rValues, gValues, bValues, aValues []uint32

			for ky := -h; ky < h; ky++ {
				for kx := -w; kx < w; kx++ {
					px := x + kx
					py := y + ky

					r, g, b, a := img.At(px, py).RGBA()
					rValues = append(rValues, r)
					gValues = append(gValues, g)
					bValues = append(bValues, b)
					aValues = append(aValues, a)
				}
			}

			// Sort the values to find the median
			sort.Slice(rValues, func(i, j int) bool { return rValues[i] < rValues[j] })
			sort.Slice(gValues, func(i, j int) bool { return gValues[i] < gValues[j] })
			sort.Slice(bValues, func(i, j int) bool { return bValues[i] < bValues[j] })
			sort.Slice(aValues, func(i, j int) bool { return aValues[i] < aValues[j] })

			medianIndex := len(rValues) / 2
			medianR := rValues[medianIndex]
			medianG := gValues[medianIndex]
			medianB := bValues[medianIndex]
			medianA := aValues[medianIndex]

			newImg.Set(x, y, color.RGBA64{uint16(medianR), uint16(medianG), uint16(medianB), uint16(medianA)})
		}
	}

	return newImg
}
func apply_min_filter(img image.Image) image.Image {

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)

	kernelHeight := 3
	kernelWidth := 3
	h := kernelHeight / 2
	w := kernelWidth / 2
	for y := h; y < height-h; y++ {
		for x := w; x < width-w; x++ {
			minR, minG, minB, minA := uint32(65535), uint32(65535), uint32(65535), uint32(65535)
			for ky := 0; ky < kernelHeight; ky++ {
				for kx := 0; kx < kernelWidth; kx++ {
					r, g, b, a := img.At(x+w-kx, y+h-ky).RGBA()
					minR = min(minR, r)
					minG = min(minG, g)
					minB = min(minB, b)
					minA = min(minA, a)
				}
			}
			newImg.Set(x, y, color.RGBA64{uint16(minR), uint16(minG), uint16(minB), uint16(minA)})
		}
	}

	return newImg
}
func apply_max_filter(img image.Image) image.Image {

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)

	kernelHeight := 3
	kernelWidth := 3
	h := kernelHeight / 2
	w := kernelWidth / 2
	for y := h; y < height-h; y++ {
		for x := w; x < width-w; x++ {
			maxR, maxG, maxB, maxA := uint32(0), uint32(0), uint32(0), uint32(0)
			for ky := 0; ky < kernelHeight; ky++ {
				for kx := 0; kx < kernelWidth; kx++ {
					r, g, b, a := img.At(x+w-kx, y+h-ky).RGBA()
					maxR = max(maxR, r)
					maxG = max(maxG, g)
					maxB = max(maxB, b)
					maxA = max(maxA, a)
				}
			}
			newImg.Set(x, y, color.RGBA64{uint16(maxR), uint16(maxG), uint16(maxB), uint16(maxA)})
		}
	}

	return newImg
}
func apply_averaging_filter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)

	kernel := [][]float64{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	kernelSum := 9.0
	kernelHeight := len(kernel)
	kernelWidth := len(kernel[0])
	h := kernelHeight / 2
	w := kernelWidth / 2
	for y := h; y < height-h; y++ {
		for x := w; x < width-w; x++ {
			var sumR, sumG, sumB, sumA float64
			for ky := 0; ky < kernelHeight; ky++ {
				for kx := 0; kx < kernelWidth; kx++ {
					r, g, b, a := img.At(x+w-kx, y+h-ky).RGBA()
					sumR += float64(r) * kernel[ky][kx]
					sumG += float64(g) * kernel[ky][kx]
					sumB += float64(b) * kernel[ky][kx]
					sumA += float64(a) * kernel[ky][kx]
				}
			}

			sumR /= kernelSum
			sumG /= kernelSum
			sumB /= kernelSum
			sumA /= kernelSum

			newImg.Set(x, y, color.RGBA64{uint16(sumR), uint16(sumG), uint16(sumB), uint16(sumA)})
		}
	}

	return newImg
}
func apply_gaussian_filter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)

	kernel := [][]float64{
		{1, 4, 7, 4, 1},
		{4, 16, 26, 16, 4},
		{7, 26, 41, 26, 7},
		{4, 16, 26, 16, 4},
		{1, 4, 7, 4, 1},
	}
	kernelSum := 273.0
	kernelHeight := len(kernel)
	kernelWidth := len(kernel[0])
	h := kernelHeight / 2
	w := kernelWidth / 2
	for y := h; y < height-h; y++ {
		for x := w; x < width-w; x++ {
			var sumR, sumG, sumB, sumA float64
			for ky := 0; ky < kernelHeight; ky++ {
				for kx := 0; kx < kernelWidth; kx++ {
					r, g, b, a := img.At(x+w-kx, y+h-ky).RGBA()
					sumR += float64(r) * kernel[ky][kx]
					sumG += float64(g) * kernel[ky][kx]
					sumB += float64(b) * kernel[ky][kx]
					sumA += float64(a) * kernel[ky][kx]
				}
			}

			sumR /= kernelSum
			sumG /= kernelSum
			sumB /= kernelSum
			sumA /= kernelSum

			sumR = clamp(sumR, 0, 65535)
			sumG = clamp(sumG, 0, 65535)
			sumB = clamp(sumB, 0, 65535)
			sumA = clamp(sumA, 0, 65535)

			newImg.Set(x, y, color.RGBA64{uint16(sumR), uint16(sumG), uint16(sumB), uint16(sumA)})
		}
	}

	return newImg
}

func apply_laplacian_filter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)
	kernel := [][]float64{
		{0, 1, 0},
		{1, -4, 1},
		{0, 1, 0},
	}
	//kernel := [][]float64{
	//	{1, 1, 1},
	//	{1, -8, 1},
	//	{1, 1, 1},
	//}
	kernelHeight := len(kernel)
	kernelWidth := len(kernel[0])
	h := kernelHeight / 2
	w := kernelWidth / 2
	for y := h; y < height-h; y++ {
		for x := w; x < width-w; x++ {
			var sumR, sumG, sumB, sumA float64
			for ky := 0; ky < kernelHeight; ky++ {
				for kx := 0; kx < kernelWidth; kx++ {
					r, g, b, a := img.At(x+w-kx, y+h-ky).RGBA()
					sumR += float64(r) * kernel[ky][kx]
					sumG += float64(g) * kernel[ky][kx]
					sumB += float64(b) * kernel[ky][kx]
					sumA += float64(a) * kernel[ky][kx]
				}
			}
			newImg.Set(x, y, color.RGBA64{uint16(sumR), uint16(sumG), uint16(sumB), uint16(sumA)})
		}
	}
	return newImg
}

func apply_unsharp_mask_filter(img image.Image) image.Image {
	// Convert the image to grayscale
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	grayImg := image.NewGray(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			grayImg.Set(x, y, gray)
		}
	}

	// Blur the image using a simple box blur
	blurredImg := image.NewGray(bounds)
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			var sum uint32
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					gray := grayImg.GrayAt(x+dx, y+dy)
					sum += uint32(gray.Y)
				}
			}
			blurredImg.Set(x, y, color.Gray{Y: uint8(sum / 9)})
		}
	}

	// Create the mask by subtracting the blurred image from the original image
	mask := image.NewGray(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			gray := grayImg.GrayAt(x, y)
			blur := blurredImg.GrayAt(x, y)
			mask.Set(x, y, color.Gray{Y: uint8(math.Abs(float64(gray.Y) - float64(blur.Y)))})
		}
	}

	// Amplify the mask
	amplifiedMask := image.NewGray(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			gray := mask.GrayAt(x, y)
			amplifiedMask.Set(x, y, color.Gray{Y: uint8(float64(gray.Y) * 3)})
		}
	}

	// Add the amplified mask to the original image
	unsharpImg := image.NewGray(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			gray := grayImg.GrayAt(x, y)
			mask := amplifiedMask.GrayAt(x, y)
			unsharpImg.Set(x, y, color.Gray{Y: uint8(math.Min(float64(gray.Y)+float64(mask.Y), 255))})
		}
	}

	return unsharpImg
}

func apply_roberts_filter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)

	//	 x= {1, 0},  y= {0, 1},
	//	    {0,-1},     {-1, 0},

	for y := 0; y < height-1; y++ {
		for x := 0; x < width-1; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r2, g2, b2, _ := img.At(x+1, y+1).RGBA()
			r3, g3, b3, _ := img.At(x+1, y).RGBA()
			r4, g4, b4, _ := img.At(x, y+1).RGBA()

			gxR := r1 - r2
			gyR := r3 - r4
			gxG := g1 - g2
			gyG := g3 - g4
			gxB := b1 - b2
			gyB := b3 - b4

			//Euclidan distance forumla
			magnitudeR := uint8(math.Sqrt(float64((gxR*gxR)+(gyR*gyR))) / 256)
			magnitudeG := uint8(math.Sqrt(float64((gxG*gxG)+(gyG*gyG))) / 256)
			magnitudeB := uint8(math.Sqrt(float64((gxB*gxB)+(gyB*gyB))) / 256)
			newImg.Set(x, y, color.RGBA{magnitudeR, magnitudeG, magnitudeB, 255})
		}
	}
	return newImg
}

func apply_sobel_filter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)

	//This is for sobel in X axis
	kernel := [][]float64{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	// For the Y axis
	//kernel := [][]float64{
	//	{-1, -2, -1},
	//	{0, 0, 0},
	//	{1, 2, 1},
	//}

	kernelHeight := len(kernel)
	kernelWidth := len(kernel[0])
	h := kernelHeight / 2
	w := kernelWidth / 2
	for y := h; y < height-h; y++ {
		for x := w; x < width-w; x++ {
			var sumR, sumG, sumB, sumA float64
			for ky := 0; ky < kernelHeight; ky++ {
				for kx := 0; kx < kernelWidth; kx++ {
					r, g, b, a := img.At(x-w+kx, y-h+ky).RGBA()
					sumR += float64(r) * kernel[ky][kx]
					sumG += float64(g) * kernel[ky][kx]
					sumB += float64(b) * kernel[ky][kx]
					sumA += float64(a) * kernel[ky][kx]
				}
			}

			newImg.Set(x, y, color.RGBA64{uint16(sumR), uint16(sumG), uint16(sumB), uint16(sumA)})
		}
	}

	return newImg
}

func apply_salt_pepper_filter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	total_pixels := width * height
	output_img := image.NewRGBA(image.Rect(0, 0, width, height))
	// Copy the original image to the new image
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			output_img.Set(x, y, img.At(x, y))
		}
	}
	// Add white noise (salt)
	number_of_pixels := rand.Intn(total_pixels + 1)
	for i := 0; i < number_of_pixels; i++ {
		y_coord := rand.Intn(height)
		x_coord := rand.Intn(width)
		output_img.Set(x_coord, y_coord, color.RGBA{255, 255, 255, 255})
	}
	// Add black noise (pepper)
	number_of_pixels = rand.Intn(total_pixels + 1)
	for i := 0; i < number_of_pixels; i++ {
		y_coord := rand.Intn(height)
		x_coord := rand.Intn(width)
		output_img.Set(x_coord, y_coord, color.RGBA{0, 0, 0, 255})
	}
	return output_img
}

func apply_gaussian_noise_filter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)

	mean := 0.0
	stdDev := 25.0 // Adjust as needed

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			r8 := float64(r >> 8)
			g8 := float64(g >> 8)
			b8 := float64(b >> 8)

			noisyR := addGaussianNoise(r8, mean, stdDev)
			noisyG := addGaussianNoise(g8, mean, stdDev)
			noisyB := addGaussianNoise(b8, mean, stdDev)

			noisyR = clamp(noisyR, 0, 255)
			noisyG = clamp(noisyG, 0, 255)
			noisyB = clamp(noisyB, 0, 255)

			newImg.Set(x, y, color.RGBA{uint8(noisyR), uint8(noisyG), uint8(noisyB), 255})
		}
	}

	return newImg
}

func addGaussianNoise(value, mean, stdDev float64) float64 {
	noise := rand.NormFloat64()*stdDev + mean
	return value + noise
}
func apply_uniform_noise_filter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)

	q := 0.0
	t := 0.2

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, alpha := img.At(x, y).RGBA()

			// Convert to float64 and normalize to [0,1]
			rFloat := float64(r) / 65535.0
			gFloat := float64(g) / 65535.0
			bFloat := float64(b) / 65535.0

			// Generate uniform noise
			noise := q + rand.Float64()*(t-q)

			// Add noise to image
			rFloat = rFloat + noise
			gFloat = gFloat + noise
			bFloat = bFloat + noise

			// Clip values to [0,1]
			rFloat = clamp(rFloat, 0, 1)
			gFloat = clamp(gFloat, 0, 1)
			bFloat = clamp(bFloat, 0, 1)

			// Convert back to uint16 and set pixel
			newImg.Set(x, y, color.RGBA64{R: uint16(rFloat * 65535), G: uint16(gFloat * 65535), B: uint16(bFloat * 65535), A: uint16(alpha)})
		}
	}

	return newImg
}
