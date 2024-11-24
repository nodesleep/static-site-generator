---
title: "Static Site Generator Example"
---

# Welcome

This site was generated entirely from markdown. To modify, please edit the markdown file in `content/index.md`. The generation is done through a Golang application that parses the markdown, wraps in a base HTML template found at `templates/base.html`, then outputs the result into the `public/`.

## Usage

1. Modify the `index.md` file located in `content/`
2. If desired, modify the HTML file located in `templates/`
3. The application reads any CSS file named `style.css` in the `assets/` directory. If you don't wish to use The Monospace Web, simply replace the CSS with your own.
4. Once finished, simply run the application with `go run .`

## Credits

- This site generator uses Oskar Wickström's ["The Monospace Web"](https://github.com/owickstrom/the-monospace-web) CSS file.
