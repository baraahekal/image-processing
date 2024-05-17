package main

import (
	_ "fmt"
	"image"
	"image/color"
	_ "image/color"
	_ "image/jpeg" // or "image/png" depending on your image type
	_ "os"
)

// TEST FUNCTION
func apply_median_filter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			newImg.Set(width-x-1, y, img.At(x, y))
		}
	}

	return newImg
	return img
}
func apply_min_filter(img image.Image) image.Image {
	// Implement the filter
	return img
}

func apply_max_filter(img image.Image) image.Image {
	// Implement the filter
	return img
}

func apply_averaging_filter(img image.Image) image.Image {
	// Implement the filter
	return img
}

func apply_gaussian_filter(img image.Image) image.Image {
	// Implement the filter
	return img
}

func apply_laplacian_filter(img image.Image) image.Image {
	// Implement the filter
	return img
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
	// Implement the filter
	return img
}

func apply_salt_pepper_filter(img image.Image) image.Image {

	return img
}

func apply_gaussian_noise_filter(img image.Image) image.Image {
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

func apply_uniform_noise_filter(img image.Image) image.Image {
	// Implement the filter
	return img
}
