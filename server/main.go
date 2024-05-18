package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/image_processing", imageHandler)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
func imageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		fmt.Fprintf(w, "Hello from the server!")
		return
	}

	if r.Method == "GET" {
		fmt.Fprintf(w, "Hello from the server!")
		return
	}

	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		var m map[string]interface{}
		json.Unmarshal(body, &m)
		base64data := m["data"].(string)
		base64data = strings.Split(base64data, ",")[1]

		imageData, err := base64.StdEncoding.DecodeString(base64data)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Could not decode base64 data", http.StatusBadRequest)
			return
		}

		img, _, err := image.Decode(bytes.NewReader(imageData))

		// Extract the filter from the request body
		filter := m["filter"].(string)

		// Apply the corresponding image processing operation based on the filter
		switch filter {
		case "0-0-0":
			img = apply_median_filter(img)
		case "0-0-1-0":
			img = apply_median_filter(img)
		case "0-0-1-1":
			img = apply_min_filter(img)
		case "0-0-1-2":
			img = apply_max_filter(img)
		case "0-0-2":
			img = apply_averaging_filter(img)
		case "0-0-3":
			img = apply_gaussian_filter(img)
		case "0-1-0":
			img = apply_laplacian_filter(img)
		case "0-1-1":
			img = apply_unsharp_mask_filter(img)
		case "0-1-2":
			img = apply_roberts_filter(img)
		case "0-1-3":
			img = apply_sobel_filter(img)
		case "0-2-0":
			img = apply_salt_pepper_filter(img)
		case "0-2-1":
			img = apply_gaussian_noise_filter(img)
		case "0-2-2":
			img = apply_uniform_noise_filter(img)
		case "1-0":
			img = apply_histogram_equalization_filter(img)
		case "1-1":
			img = apply_histogram_specification_filter(img)
		case "1-2":
			img = apply_fourier_transform_filter(img)
		case "1-3-0":
			img = apply_nearest_neighbour_filter(img, 500, 500)
		case "1-3-1":
			img = applyBilinearFilter(img, 500, 500)
		case "1-3-2":
			img = applyBicubicFilter(img, 500, 500)
		case "2-0":
			img = apply_huffman_coding(img)

		default:
			http.Error(w, "Unsupported filter", http.StatusBadRequest)
			return
		}

		// Encode the image and send it back to the client
		contentType := r.Header.Get("Content-Type")
		if strings.Contains(contentType, "image/jpeg") {
			err = jpeg.Encode(w, img, nil)
		} else if strings.Contains(contentType, "image/png") {
			err = png.Encode(w, img)
		} else {
			http.Error(w, "Unsupported image format", http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, "Could not encode image", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
