package translations

import (
	"fmt"
	"os"
	"regexp"
	"github.com/pborman/indent"
	"bytes"
	"embed"
)

func Output(fs embed.FS, filenames []string) {
	pattern := regexp.MustCompile(`\{\{ \.Translator\.String "(.*)" \}\}`)
	keys := []string{}
	for _, filename := range filenames {
		contents, err := fs.ReadFile(fmt.Sprintf("%s.html", filename))
		if err != nil { 
			fmt.Printf("translations.go: err trying to open %s; skipping\n", filename)
			continue 
		}
		matches := pattern.FindAllStringSubmatch(string(contents), -1)
		for _, m := range matches {
			keys = append(keys, m[1])
		}
	}
	header := `// THIS IS GENERATED, PLS NO EDITS
package translations
import (
	"eyeneighteenn/tokibundle"
)
func Translations(t tokibundle.Reader) {
	// LOOP
`
	footer := `}`
	body := new(bytes.Buffer)
	indentWriter := indent.New(body, "\t")

	for _, key := range keys {
		fmt.Fprintf(indentWriter, "t.String(\"%s\")\n", key)
	}

	out := fmt.Sprintf("%s%s%s", header, body.String(), footer)
	os.WriteFile("./translations/generated-translations.go", []byte(out), 0666)
}
