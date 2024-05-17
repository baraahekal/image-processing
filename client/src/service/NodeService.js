export const NodeService = {
    getTreeNodes: async () => {
        return [
            {
                "key": "0",
                "label": "Spatial Domain Filters",
                "icon": "pi pi-folder",
                "children": [
                    {   "key": "0-0",
                        "label": "Smoothing Spatial filters",
                        "icon": "pi pi-folder",
                        "children": [
                            { "key": "0-0-0", "label": "Median filter", "icon": "pi pi-folder", },
                            { "key": "0-0-1", "label": "Adaptive filters" ,
                                "icon": "pi pi-folder",
                                "children": [
                                    { "key": "0-0-1-0", "label": "Median", "icon": "pi pi-folder", },
                                    { "key": "0-0-1-1", "label": "Min", "icon": "pi pi-folder", },
                                    { "key": "0-0-1-2", "label": "Max", "icon": "pi pi-folder", },
                                ]
                            },
                            { "key": "0-0-2", "label": "Averaging filter","icon": "pi pi-folder", },
                            { "key": "0-0-3", "label": "Gaussian filter", "icon": "pi pi-folder", },
                        ]
                    },
                    {   "key": "0-1", "label": "Sharpening Spatial filters" ,
                        "icon": "pi pi-folder",
                        "children": [
                            { "key": "0-1-0", "label": "Laplacian filter", "icon": "pi pi-folder", },
                            { "key": "0-1-1", "label": "Unsharp Masking", "icon": "pi pi-folder", },
                            { "key": "0-1-2", "label": "Roberts Cross-Gradient Operators", "icon": "pi pi-folder", },
                            { "key": "0-1-3", "label": "Sobel filter", "icon": "pi pi-folder", },
                        ]
                    },
                    {   "key": "0-2",
                        "label": "Noise filters",
                        "icon": "pi pi-folder",
                        "children": [
                            { "key": "0-2-0", "label": "Salt and Pepper Noise", "icon": "pi pi-folder",},
                            { "key": "0-2-1", "label": "Gaussian Noise", "icon": "pi pi-folder", },
                            { "key": "0-2-2", "label": "Uniform Noise", "icon": "pi pi-folder", },
                        ]
                    },
                ]
            },
            {
                "key": "1",
                "icon": "pi pi-folder",
                "label": "Transform Domain filters",
                "children": [
                    { "key": "1-0", "label": "Histogram Equalization", "icon": "pi pi-folder", },
                    { "key": "1-1", "label": "Histogram Specification", "icon": "pi pi-folder", },
                    { "key": "1-2", "label": "Fourier transform", "icon": "pi pi-folder", },
                    { "key": "1-3", "label": "Interpolation", "icon": "pi pi-folder", },
                ]
            },
            {
                "key": "2",
                "icon": "pi pi-folder",
                "label": "Compression Techniques",
                "children": [
                    { "key": "2-0", "label": "Huffman coding", "icon": "pi pi-folder", },
                ]
            }
        ];
    }
};