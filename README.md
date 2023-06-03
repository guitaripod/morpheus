# Morpheus 

"Morpheus" is a command-line utility written in Go that, akin to its namesake from the Matrix, can transform the reality of your images. This tool allows you to process an image by reading the image's file path from your clipboard and writing the processed image's path back to your clipboard. This is especially useful for bulk editing or processing images without the need for manual input or navigation.

## Requirements

- Go 1.16 or later
- Image file path must be available on the clipboard (Note: The clipboard should contain the image file path, not the image file itself)

## Installation

1. Clone this repository using the following command:
```bash
git clone https://github.com/yourusername/morpheus.git
```
2. Navigate to the project's directory:
```bash
cd morpheus
```
3. Build the project with the Go command:
```bash
go build
```

## Usage

1. Ensure that the file path of the image you want to process is copied to your clipboard. Please note that Morpheus works with the path to the image, not the image file itself.
2. Run the Morpheus binary from the command line:
```bash
./morpheus
```
3. Morpheus will read the image file path from your clipboard, process the image, and write the path of the output image back to your clipboard. It will also print a success message with the output image's path:
```
Image saved to /path/to/output/image.png
```
4. If there is an error at any stage, Morpheus will print an error message and exit.

## Contributing

We welcome contributions! Please see our contributing guidelines for more details.

## License

This project is licensed under the terms of the MIT license.
