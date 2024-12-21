package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const example = `{'telemetry.sdk.version': '1.29.0', 'deployment.environment': 'development', 'telemetry.sdk.name': 'opentelemetry', 'service.version': '0.0.1', 'http.user_agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', 'http.url': 'http://localhost:1111/_next/image?url=%2Fimages%2Fhero.webp&w=1080&q=75', 'session.id': '38f8224be670d12db0196e04047abddd', 'location.href': 'http://localhost:1111/', 'service.name': 'realworld-web', 'telemetry.sdk.language': 'webjs', 'os.name': 'macOS', 'os.version': '10.15.7', 'city': '{"a": 1, "b": 2}', "tree":"sky", "color":"{"red":"true", "blue":false,"green": 1, "pink":"{"aa":"111", "bb": 22, "cc": "{"dd":"ee", "ff":"[{"dd":"sad","kk":3243},{"hh":"asdf","dddd":false}]", "gg":"["{"aa":"bb"}","{"cc":true}"]"}"}"}"}`

type cleaner string

func (c cleaner) clean() cleaner {
	return cleaner(strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						string(c), `'`, `"`),
					`"{`, `{`),
				`}"`, `}`),
			`"[`, `[`),
		`]"`, `]`),
	)
}

type parsedMap map[string]interface{}

func (p parsedMap) String() string {
	v, _ := json.MarshalIndent(p, "", "   ")
	return string(v)
}

func main() {
	fmt.Println("put the dirty json format:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			return
		}
		fmt.Println("=== result ======================================")
		t := parsedMap(make(map[string]interface{}))
		if err := json.Unmarshal([]byte(cleaner(input).clean()), &t); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(t)
		return
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
}
