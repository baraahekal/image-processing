package main

import (
	_ "fmt"
	"image"
	"image/color"
	_ "image/color"
	_ "image/jpeg" // or "image/png" depending on your image type
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

func apply_unsharp_masking_filter(img image.Image) image.Image {
	// Implement the filter
	return img
}

func apply_roberts_filter(img image.Image) image.Image {
	// Implement the filter
	return img
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

func apply_salt_pepper_filter(img image.Image) image.Image {

	return img
}

func apply_gaussian_noise_filter(img image.Image) image.Image {

	return img
}

func apply_uniform_noise_filter(img image.Image) image.Image {
	// Implement the filter
	return img
}
