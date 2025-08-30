package translations

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"github.com/pborman/indent"
	"bytes"
	"embed"
)

func Output(fs embed.FS, filenames []string) {
	pattern := regexp.MustCompile(`\{\{ \.Translator\.String "(.*)"\s?(.*)?\s*\}\}`)
	keys := []string{}
	additional := make(map[string]string)
	for _, filename := range filenames {
		contents, err := fs.ReadFile(fmt.Sprintf("%s.html", filename))
		if err != nil { 
			fmt.Printf("translations.go: err trying to open %s; skipping\n", filename)
			continue 
		}
		matches := pattern.FindAllStringSubmatch(string(contents), -1)
		for _, m := range matches {
			key := m[1]
			data := m[2]
			keys = append(keys, key)
			// additional data
			if len(data) > 0 {
				additional[key] = data
			}
		}
	}
	header := `// THIS IS GENERATED, PLS NO EDITS
package translations
import (
	"eyeneighteenn/tokibundle"
	"time"
)
func Translations(t tokibundle.Reader) {
	time.Now()
	// LOOP
`
	footer := `}`
	body := new(bytes.Buffer)
	indentWriter := indent.New(body, "\t")

	for _, key := range keys {
		if data, has := additional[key]; has {
			parts := strings.Split(data, " ")
			var processed []string
			for _, part := range parts {
				if part == ".Time" {
					processed = append(processed, "time.Now()")
				}
				if part == ".NeutralName" {
					processed = append(processed, `tokibundle.String{Value: "Neutralo", Gender: tokibundle.GenderNeutral}`)
				}
				if part == ".Action" {
					processed = append(processed, `"GenericAction"`)
				}
			}
			fmt.Fprintf(indentWriter, "t.String(\"%s\", %s)\n", key, strings.Join(processed, ", "))
		} else {
			fmt.Fprintf(indentWriter, "t.String(\"%s\")\n", key)
		}
	}

	out := fmt.Sprintf("%s%s%s", header, body.String(), footer)
	os.WriteFile("./translations/generated-translations.go", []byte(out), 0666)
}
