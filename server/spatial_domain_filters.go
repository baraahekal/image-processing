package main

import "image"

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
	// Implement the filter
	return img
}

func apply_gaussian_noise_filter(img image.Image) image.Image {
	// Implement the filter
	return img
}

func apply_uniform_noise_filter(img image.Image) image.Image {
	// Implement the filter
	return img
}
