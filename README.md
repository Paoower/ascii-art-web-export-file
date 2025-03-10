# ASCII Art Web Export File

A web application for ASCII art generation with file export capabilities.

## Description

This project extends the ASCII Art Web application by adding the ability to export the generated ASCII art to various file formats. Users can generate ASCII art and download it as a text file, HTML, or other supported formats.

## Features

- All functionality of the ASCII Art Web application
- Export generated ASCII art to different file formats
- Download options integrated into the web interface
- Format selection for exports

## Usage

### Running the Server

```bash
go run .
```

Then navigate to `http://localhost:8080` in your web browser.

## Implementation Details

- File generation on the server-side
- Content-Disposition headers for file downloads
- MIME type handling for different file formats
- Client-side and server-side validation
- Temporary file management

## Requirements

- Go 1.13 or higher
- Web browser with JavaScript enabled
