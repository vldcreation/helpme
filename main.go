package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"

	"github.com/spf13/cobra"
)

var (
	pkgFlag  string
	saveFlag bool
	runFlag  bool
	langFlag string
)

func main() {
	var rootCmd = &cobra.Command{
		GroupID: "helpme",
		Use:     "helpme [function name]",
		Short:   "Search and generate code examples",
		Args:    cobra.ExactArgs(1),
		Run:     run,
	}

	rootCmd.PersistentFlags().StringVarP(&langFlag, "lang", "l", "", "Language to search (go/javascript)")
	rootCmd.PersistentFlags().StringVarP(&pkgFlag, "pkg", "p", "", "Package name (optional)")
	rootCmd.PersistentFlags().BoolVarP(&saveFlag, "save", "s", false, "Save example to a file")
	rootCmd.PersistentFlags().BoolVarP(&runFlag, "run", "r", false, "Run the saved example file")

	rootCmd.MarkPersistentFlagRequired("lang")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	funcName := args[0]

	var docURL string
	switch langFlag {
	case "go", "golang":
		langFlag = "go"
		if pkgFlag != "" {
			docURL = fmt.Sprintf("https://pkg.go.dev/%s#%s", pkgFlag, funcName)
		} else {
			docURL = fmt.Sprintf("https://pkg.go.dev/search?q=%s", funcName)
		}
	case "javascript", "js":
		langFlag = "javascript"
		docURL = fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/%s", funcName)
	default:
		fmt.Printf("Unsupported language: %s\n", langFlag)
		os.Exit(1)
	}

	fmt.Printf("Documentation URL: %s\n", docURL)

	if saveFlag {
		fileName := fmt.Sprintf("example_%s_%s.%s", pkgFlag, funcName, getFileExtension(langFlag))
		filePath := filepath.Join("examples", langFlag, fileName)

		// Create language directory if it doesn't exist
		os.MkdirAll(filepath.Dir(filePath), 0755)

		// Generate example code based on language and function
		exampleCode := generateExample(langFlag, pkgFlag, funcName)

		// Save example code to file
		err := os.WriteFile(filePath, []byte(exampleCode), 0644)
		if err != nil {
			fmt.Printf("Error saving example: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Example saved to: %s\n", filePath)

		if runFlag {
			fmt.Printf("Running example file: %s\n", filePath)
			if err := runExample(langFlag, filePath); err != nil {
				fmt.Printf("Error running example: %v\n", err)
				os.Exit(1)
			}
		}
	}
}

func runExample(lang, filePath string) error {
	var cmd *exec.Cmd
	switch lang {
	case "go":
		cmd = exec.Command("go", "run", filePath)
	case "javascript":
		cmd = exec.Command("node", filePath)
	default:
		return fmt.Errorf("unsupported language for running examples: %s", lang)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getFileExtension(lang string) string {
	switch lang {
	case "go":
		return "go"
	case "javascript":
		return "js"
	default:
		return ""
	}
}

func generateExample(lang, pkg, funcName string) string {
	switch lang {
	case "go":
		// Fetch example from pkg.go.dev
		resp, err := http.Get(fmt.Sprintf("https://pkg.go.dev/%s@go1.23.5#example-%s", pkg, funcName))
		if err != nil {
			fmt.Printf("Error fetching example: %v\n", err)
			return fmt.Sprintf(`package main

import (
	"fmt"
	"%s"
)

func main() {
	// TODO: Add example code for %s.%s
	// Visit: https://pkg.go.dev/%s#%s for documentation
}
`, pkg, pkg, funcName, pkg, funcName)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			return ""
		}

		// Parse HTML content
		doc, err := html.Parse(bytes.NewReader(body))
		if err != nil {
			fmt.Printf("Error parsing HTML: %v\n", err)
			return ""
		}

		// Find example code in the HTML
		var exampleCode string
		var findExample func(*html.Node)
		findExample = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "details" {
				for _, attr := range n.Attr {
					if attr.Key == "id" && strings.Contains(attr.Val, fmt.Sprintf("example-%s", funcName)) {
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

		// Fallback to template if no example found
		return fmt.Sprintf(`package main

import (
	"fmt"
	"%s"
)

func main() {
	// TODO: Add example code for %s.%s
	// Visit: https://pkg.go.dev/%s#%s for documentation
}
`, pkg, pkg, funcName, pkg, funcName)

	case "javascript":
		if funcName == "join" {
			return `// Example usage of Array.prototype.join()
// Visit: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/join

const elements = ['Fire', 'Air', 'Water'];
console.log(elements.join());
// Expected output: "Fire,Air,Water"

console.log(elements.join(''));
// Expected output: "FireAirWater"

console.log(elements.join('-'));
// Expected output: "Fire-Air-Water"`
		}
		return fmt.Sprintf(`// Example usage of %s
// Visit: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/%s

// TODO: Add example code
`, funcName, funcName)

	default:
		return ""
	}
}
