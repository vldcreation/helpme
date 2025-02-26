package pkg

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	tmpl "vldcreation/github.com/helpme/templates/go/pkg"

	"golang.org/x/net/html"
)

// go generator
type goGenerator struct {
	l Language
}

func (g *goGenerator) Generate() error {
	sampleCode := g.generateExample()

	if g.l.f.save {
		filePath, err := g.writeExample([]byte(sampleCode))
		if err != nil {
			return err
		}

		fmt.Printf("Example code saved to: %s\n", filePath)

		if g.l.f.execute {
			fmt.Printf("Running example code...\n")
			cmd := exec.Command("go", "run", g.getSavePath())
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			fmt.Printf("Output:\n")
			if err := cmd.Run(); err != nil {
				fmt.Printf("Error running example code: %v\n", err)
			}
		}
	}

	return nil
}

func (g *goGenerator) getDocumentUrl() string {
	if g.l.pkg == "" {
		return fmt.Sprintf("%ssearch?q=%s", docBaseUrl["go"], g.l.funcName)
	}

	return fmt.Sprintf(exampleCodeBaseUrl["go"], g.l.pkg, g.l.funcName)
}

func (g *goGenerator) writeExample(b []byte) (string, error) {
	filePath := g.getSavePath()

	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	if err := os.WriteFile(filePath, b, 0644); err != nil {
		return "", fmt.Errorf("failed to write file: %v", err)
	}

	return filePath, nil
}

func (g *goGenerator) getFileExtension() string {
	return "go"
}

func (g *goGenerator) getSavePath() string {
	_, b, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(b), "../../../")
	fileName := fmt.Sprintf("example_%s_%s.%s", g.l.pkg, g.l.funcName, g.getFileExtension())
	filePath := filepath.Join("examples", g.l.lang, fileName)
	filePath = filepath.Join(rootPath, filePath)
	return filePath
}

func (g *goGenerator) generateExample() string {
	fmt.Printf("Documentation URL: %s\n", g.getDocumentUrl())
	resp, err := http.Get(g.getDocumentUrl())
	if err != nil {
		slog.Error("failed to get example code", "error", err)
		return g.generateDefaultTemplate()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read response body", "error", err)
		return g.generateDefaultTemplate()
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		slog.Error("failed to parse html", "error", err)
		return g.generateDefaultTemplate()
	}

	if exampleCode := g.extractExampleCode(doc); exampleCode != "" {
		return exampleCode
	}

	slog.Info("failed to extract example code")
	return g.generateDefaultTemplate()
}

func (g *goGenerator) extractExampleCode(doc *html.Node) string {
	var exampleCode string
	var findExample func(*html.Node)
	findExample = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "details" {
			for _, attr := range n.Attr {
				if attr.Key == "id" && strings.Contains(attr.Val, fmt.Sprintf("example-%s", g.l.funcName)) {
					// Extraxt child with tag div.Docmentation-exampleDetailsBody
					findChild := func(n *html.Node) *html.Node {
						if n.Type == html.ElementNode && n.Data == "div" {
							for _, attr := range n.Attr {
								if attr.Key == "class" && strings.Contains(attr.Val, "Documentation-exampleDetailsBody") {
									return n
								}
							}
						}
						return nil
					}
					copyN := &html.Node{}
					*copyN = *n

					for c := copyN.FirstChild; c != nil; c = c.NextSibling {
						if child := findChild(c); child != nil {
							n = child
							// if  child exists, extra only child that has tag pre.Documentation-exampleCode
							for c := n.FirstChild; c != nil; c = c.NextSibling {
								if c.Type == html.ElementNode && c.Data == "pre" {
									for _, attr := range c.Attr {
										if attr.Key == "class" && strings.Contains(attr.Val, "Documentation-exampleCode") {
											n = c
											break
										}
									}
								}
							}
							break
						}
					}

					// Extract only

					var b strings.Builder
					var extractText func(*html.Node)
					extractText = func(n *html.Node) {
						if n.Type == html.TextNode {
							b.WriteString(n.Data)
						}
						for c := n.FirstChild; c != nil; c = c.NextSibling {
							extractText(c)
						}
					}
					extractText(n)
					exampleCode = b.String()
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findExample(c)
		}
	}
	findExample(doc)

	if exampleCode != "" {
		return exampleCode
	}

	return ""
}

func (g *goGenerator) generateDefaultTemplate() string {
	return fmt.Sprintf(string(tmpl.DefaultPackage), g.l.pkg, g.l.funcName, g.l.pkg, g.l.funcName)
}
