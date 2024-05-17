package main

import (
	"container/heap"
	"fmt"
	"image"
	"image/color"
)

// Huffman Tree Node
type HuffmanNode struct {
	Value       uint8
	Frequency   int
	Left, Right *HuffmanNode
}

// Priority Queue Implementation for Huffman Tree Nodes
type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Frequency < pq[j].Frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*HuffmanNode)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func buildHuffmanTree(freq map[uint8]int) *HuffmanNode {
	pq := make(PriorityQueue, len(freq))
	i := 0
	for value, frequency := range freq {
		pq[i] = &HuffmanNode{Value: value, Frequency: frequency}
		i++
	}
	heap.Init(&pq)

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*HuffmanNode)
		right := heap.Pop(&pq).(*HuffmanNode)
		merged := &HuffmanNode{
			Value:     0,
			Frequency: left.Frequency + right.Frequency,
			Left:      left,
			Right:     right,
		}
		heap.Push(&pq, merged)
	}

	return heap.Pop(&pq).(*HuffmanNode)
}

func buildHuffmanCodes(node *HuffmanNode, prefix string, codes map[uint8]string) {
	if node.Left == nil && node.Right == nil {
		codes[node.Value] = prefix
		return
	}
	if node.Left != nil {
		buildHuffmanCodes(node.Left, prefix+"0", codes)
	}
	if node.Right != nil {
		buildHuffmanCodes(node.Right, prefix+"1", codes)
	}
}

func huffmanEncode(data []uint8) (string, map[uint8]string) {
	// calculate the frequency of each value
	freq := make(map[uint8]int)
	for _, value := range data {
		freq[value]++
	}

	// build the Huffman Tree
	huffmanTree := buildHuffmanTree(freq)

	// generate the Huffman Codes
	codes := make(map[uint8]string)
	buildHuffmanCodes(huffmanTree, "", codes)

	// encode the data
	var encodedData string
	for _, value := range data {
		encodedData += codes[value]
	}

	return encodedData, codes
}

func apply_huffman_coding(img image.Image) image.Image {
	// convert the image to grayscale
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var pixelValues []uint8
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			pixelValues = append(pixelValues, gray.Y)
		}
	}

	// encode the pixel values using Huffman coding
	encodedData, codes := huffmanEncode(pixelValues)

	// printing huffman codes
	fmt.Println("Huffman Codes:", codes)
	fmt.Println("Encoded Data Length:", len(encodedData))

	return img // return original image
}
