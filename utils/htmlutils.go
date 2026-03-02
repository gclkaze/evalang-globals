package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func ParseHTMLByXPATHAndGetOne(content string, xpath string) (value string, err error) {
	if strings.Contains(content, `\"`) {
		content = strings.ReplaceAll(content, `\"`, "\"")
	}
	doc, err := htmlquery.Parse(strings.NewReader(content))
	if err != nil {
		return "", err
	}
	list, err := htmlquery.QueryAll(doc, xpath)
	if err != nil {
		return "", err
	}

	if len(list) > 1 {
		return "", fmt.Errorf("found more than one results")
	}

	it := list[0]
	var b bytes.Buffer
	err = html.Render(&b, it)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func GetRandomDoubleInRange(min float64, max float64) float64 {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	r := min + rand.Float64()*(max-min)
	return r
}

func ExtractOneFromRegex(content string, regex string) (value string, err error) {
	if strings.Contains(content, `\"`) {
		content = strings.ReplaceAll(content, `\"`, "\"")
	}
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(content)
	if len(match) == 0 {
		return "", nil
	}
	if len(match) > 1 {
		return "", fmt.Errorf("found more than one results")
	}
	return match[0], nil
}

func ExtractFromRegex(content string, regex string) (value []string, err error) {
	if strings.Contains(content, `\"`) {
		content = strings.ReplaceAll(content, `\"`, "\"")
	}
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(content)
	return match, nil
}

func ParseHTMLByXPATH(content string, xpath string) (value []string, err error) {
	if strings.Contains(content, `\"`) {
		content = strings.ReplaceAll(content, `\"`, "\"")
	}
	var out []string
	doc, err := htmlquery.Parse(strings.NewReader(content))
	if err != nil {
		return out, err
	}
	list, err := htmlquery.QueryAll(doc, xpath)
	if err != nil {
		return out, err
	}

	for i := 0; i < len(list); i++ {
		it := list[i]
		var b bytes.Buffer
		err := html.Render(&b, it)
		if err != nil {
			return out, err
		}
		out = append(out, b.String())
	}

	return out, nil
}

/*func main() {
	// HTML string to parse
	htmlStr := `<html><head><title>Example</title></head><body><h1>Hello, World!</h1></body></html>`

	// Parse the HTML string
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Traverse the HTML document and create a JSON object
	jsonObj := make(map[string]interface{})
	traverseHTMLNode(doc, jsonObj)

	// Print the JSON object
	jsonBytes, err := json.MarshalIndent(jsonObj, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonBytes))
}*/

func GetJSONFromHtml(h string) (j map[string]interface{}, err error) {
	doc, err := html.Parse(strings.NewReader(h))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Traverse the HTML document and create a JSON object
	jsonObj := make(map[string]interface{})
	traverseHTMLNode(doc, jsonObj)

	// Print the JSON object
	/*	jsonBytes, err := json.MarshalIndent(jsonObj, "", "  ")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		fmt.Println(string(jsonBytes))*/
	return jsonObj, nil
}

func traverseHTMLNode(n *html.Node, jsonObj map[string]interface{}) {
	if n.Type == html.ElementNode {
		// Add the element name and attributes to the JSON object
		jsonObj[n.Data] = make(map[string]interface{})
		for _, attr := range n.Attr {
			jsonObj[n.Data].(map[string]interface{})[attr.Key] = attr.Val
		}
	} else if n.Type == html.TextNode {
		// Add the text content to the JSON object
		jsonObj["text"] = n.Data
	}

	a := make([]map[string]interface{}, 0)
	// Recursively traverse the children of the node
	i := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ch := make(map[string]interface{})
		traverseHTMLNode(c, ch)
		a = append(a, ch)

		i += 1
	}

	jsonObj["children"] = a
}
