# Welcome

This application generates static sites from markdown using Go. To modify, please edit the markdown file in `content/index.md`. The generation is done through a Golang application that parses the markdown, wraps in a base HTML template found at `templates/base.html`, then outputs the result into the `public/`.

Once generated, the application will host the content at localhost on port 8080.

## Usage (Download and Run)

Select the appropriate package for your system.

- For Intel Macs: [ssg-apple-intel-amd64.zip](https://github.com/nodesleep/static-site-generator/raw/refs/heads/main/pkg/ssg-apple-intel-amd64.zip)
- For Apple Silicon Macs: [ssg-apple-silicon-arm64.zip](https://github.com/nodesleep/static-site-generator/raw/refs/heads/main/pkg/ssg-apple-silicon-arm64.zip)
- For Linux: [ssg-linux-amd64.zip](https://github.com/nodesleep/static-site-generator/raw/refs/heads/main/pkg/ssg-linux-amd64.zip)
- For Windows: [ssg-windows-amd64.zip](https://github.com/nodesleep/static-site-generator/raw/refs/heads/main/pkg/ssg-windows-amd64.zip)

Extract the files and modify to your liking. Be sure to execute `chmod +x path/to/executable` if running on Linux or Apple systems.

1. Modify the `index.md` file located in `content/`
2. If desired, modify the HTML file located in `templates/`
3. The application reads any CSS file named `style.css` in the `assets/` directory. If you don't wish to use The Monospace Web, simply replace the CSS with your own.
4. Once finished, run the executable.

## Usage (For Golang Devs)

1. Modify the `index.md` file located in `content/`
2. If desired, modify the HTML file located in `templates/`
3. The application reads any CSS file named `style.css` in the `assets/` directory. If you don't wish to use The Monospace Web, simply replace the CSS with your own.
4. Once finished, simply run the application with `go run .`

## Credits

- This site generator uses Oskar Wickström's ["The Monospace Web"](https://github.com/owickstrom/the-monospace-web) CSS file.
