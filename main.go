package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v2"
)

type FrontMatter struct {
	Title string `yaml:"title"`
}

type Page struct {
	Title   string
	Content template.HTML
}

func main() {
	inputDir := "content"
	outputDir := "public"
	assetsDir := "assets"
	templateFile := "templates/base.html"

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	files, err := filepath.Glob(filepath.Join(inputDir, "*.md"))
	if err != nil {
		log.Fatalf("Failed to read input directory: %v", err)
	}

	for _, file := range files {
		log.Printf("Processing file: %s", file)

		htmlContent, frontMatter, err := processMarkdown(file)
		if err != nil {
			log.Fatalf("Error processing file %s: %v", file, err)
		}

		outputFile := filepath.Join(outputDir, strings.TrimSuffix(filepath.Base(file), ".md")+".html")
		err = renderTemplate(templateFile, outputFile, Page{
			Title:   frontMatter.Title,
			Content: htmlContent,
		})
		if err != nil {
			log.Fatalf("Error rendering template for file %s: %v", file, err)
		}
		log.Printf("Generated: %s", outputFile)
	}

	err = copyAssets(assetsDir, outputDir)
	if err != nil {
		log.Fatalf("Error copying assets: %v", err)
	}

	log.Println("Serving files at http://localhost:8080")
	http.Handle("/", http.FileServer(http.Dir(outputDir)))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}


func processMarkdown(file string) (template.HTML, FrontMatter, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return "", FrontMatter{}, fmt.Errorf("failed to read file %s: %v", file, err)
	}

	parts := splitFrontMatter(string(content))
	if len(parts[0]) == 0 {
		return "", FrontMatter{}, fmt.Errorf("missing or invalid frontmatter in file %s", file)
	}

	var frontMatter FrontMatter
	err = yaml.Unmarshal([]byte(parts[0]), &frontMatter)
	if err != nil {
		return "", FrontMatter{}, fmt.Errorf("error parsing frontmatter in file %s: %v", file, err)
	}

	var buf bytes.Buffer
	err = goldmark.New().Convert([]byte(parts[1]), &buf)
	if err != nil {
		return "", FrontMatter{}, fmt.Errorf("error converting Markdown in file %s: %v", file, err)
	}

	return template.HTML(buf.String()), frontMatter, nil
}


func splitFrontMatter(content string) []string {
	lines := strings.Split(content, "\n")
	var frontmatter, markdown []string
	inFrontMatter := false

	for _, line := range lines {
		if strings.TrimSpace(line) == "---" {
			inFrontMatter = !inFrontMatter
			continue
		}

		if inFrontMatter {
			frontmatter = append(frontmatter, line)
		} else {
			markdown = append(markdown, line)
		}
	}

	return []string{strings.Join(frontmatter, "\n"), strings.Join(markdown, "\n")}
}

func renderTemplate(templateFile, outputFile string, page Page) error {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer f.Close()

	return tmpl.Execute(f, page)
}

func copyAssets(inputDir, outputDir string) error {
	entries, err := os.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("failed to read assets directory: %v", err)
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(inputDir, entry.Name())
		destPath := filepath.Join(outputDir, entry.Name())

		// Copy the file
		input, err := os.ReadFile(sourcePath)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %v", sourcePath, err)
		}

		err = os.WriteFile(destPath, input, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to write file %s: %v", destPath, err)
		}
	}
	return nil
}

