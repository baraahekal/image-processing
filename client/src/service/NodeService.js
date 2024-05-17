export const NodeService = {
    getTreeNodes: async () => {
        return [
            {
                "key": "0",
                "label": "Spatial Domain Filters",
                "children": [
                    { "key": "0-0",
                        "label": "Smoothing Spatial filters",
                        "children": [
                            { "key": "0-0-0", "label": "Median filter" },
                            { "key": "0-0-1", "label": "Adaptive filters" ,
                                "children": [
                                    { "key": "0-0-1-0", "label": "Median" },
                                    { "key": "0-0-1-1", "label": "Min" },
                                    { "key": "0-0-1-2", "label": "Max" },
                                ]
                            },
                            { "key": "0-0-2", "label": "Averaging filter" },
                            { "key": "0-0-3", "label": "Gaussian filter" },
                        ]
                    },
                    { "key": "0-1", "label": "Sharpening Spatial filters" ,
                        "children": [
                            { "key": "0-1-0", "label": "Laplacian filter" },
                            { "key": "0-1-2", "label": "Unsharp Masking" },
                            { "key": "0-1-3", "label": "Roberts Cross-Gradient Operators" },
                            { "key": "0-1-1", "label": "Sobel filter" },
                        ]
                    },
                    { "key": "0-2", "label": "Noise filters", "children": [
                            { "key": "0-2-0", "label": "Salt and Pepper Noise" },
                            { "key": "0-2-1", "label": "Gaussian Noise" },
                            { "key": "0-2-2", "label": "Uniform Noise" },
                        ]
                    },
                ]
            },
            {
                "key": "1",
                "label": "Transform /Frequency Domain filters",
                "children": [
                    { "key": "1-0", "label": "Histogram Equalization" },
                    { "key": "1-1", "label": "Histogram Specification" },
                    { "key": "1-2", "label": "Fourier transform" },
                    { "key": "1-3", "label": "Interpolation" },
                ]
            },
            {
                "key": "2",
                "label": "Compression Techniques",
                "children": [
                    { "key": "2-0", "label": "Huffman coding" },
                ]
            }
        ];
    }
};