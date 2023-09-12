package bodytransformer

import (
	"fmt"
	"html"
	"regexp"

	"github.com/beevik/etree"
)

// TransformBody transforms content body in format presentable for external/non-FT consumers of the content
func TransformBody(body string) (string, error) {
	doc := etree.NewDocument()

	err := doc.ReadFromString(body)
	if err != nil {
		return "", fmt.Errorf("failed to parse body as xml: %w", err)
	}

	// Find all tags with name "content" and replace their name with "ft-content", transform element attributes
	for _, el := range doc.FindElements("//content") {
		el.Tag = "ft-content"
		transformElementAttributes(el)
	}

	// Find all tags with name "related" and replace their name with "ft-related", transform element attributes
	for _, el := range doc.FindElements("//related") {
		el.Tag = "ft-related"
		transformElementAttributes(el)
	}

	// Find all tags with name "concept" and replace their name with "ft-concept", transform element attributes
	for _, el := range doc.FindElements("//concept") {
		el.Tag = "ft-concept"
		transformElementAttributes(el)
	}

	scrollableTextExtraction(doc)

	// Remove elements with particular tag names -
	// "pull-quote","promo-box","ft-related","timeline","ft-timeline","table","big-number","img"
	elementsToStrip := []string{
		"pull-quote", "promo-box", "ft-related", "timeline", "ft-timeline", "table", "big-number", "img",
		"experimental",
		"recommended",
	}
	for _, name := range elementsToStrip {
		for _, el := range doc.FindElements(
			"//" + name) {
			p := el.Parent()
			p.RemoveChild(el)
		}
	}

	// Remove blockquote elements with attribute "class" with value "twitter-tweet"
	for _, el := range doc.FindElements("//blockquote[@class='twitter-tweet']") {
		p := el.Parent()
		p.RemoveChild(el)
	}

	// Remove "a" elements with attribute "data-asset-type" with value "video"
	for _, el := range doc.FindElements("//a[@data-asset-type='video']") {
		p := el.Parent()
		p.RemoveChild(el)
	}
	// Remove "a" elements with attribute "data-asset-type" with value "interactive-graphic"
	for _, el := range doc.FindElements("//a[@data-asset-type='interactive-graphic']") {
		p := el.Parent()
		p.RemoveChild(el)
	}

	// Remove elements with tag ft-content that have attribute "type" with value "http://www.ft.com/ontology/content/ImageSet"
	for _, t := range doc.FindElements(
		"//ft-content[@type='http://www.ft.com/ontology/content/ImageSet']") {
		p := t.Parent()
		p.RemoveChild(t)
	}

	// Remove elements with tag ft-content that have attribute "type" with value "http://www.ft.com/ontology/content/MediaResource"
	for _, t := range doc.FindElements("//ft-content[@type='http://www.ft.com/ontology/content/MediaResource']") {
		p := t.Parent()
		p.RemoveChild(t)
	}

	strBody, err := doc.WriteToString()
	if err != nil {
		return "", err
	}

	// Apply specific rules to some combinations of tags
	strBody = transformParagraphElements(strBody)

	// Remove the html escape sequences, if client of the library needs html escaping, this should be their responsibility
	strBody = html.UnescapeString(strBody)

	// Remove empty lines from the output
	strBody = removeEmptyLines(strBody)

	return strBody, nil
}

// transformElementAttributes makes specific transformations to internal ft elements attributes.
// The type and url attributes are added to the end of the attribute list, where url is created based on the values of
// the existing id and type attributes. The id attribute is removed if present.
func transformElementAttributes(contentTag *etree.Element) {
	idAttr := contentTag.RemoveAttr("id")
	typeAttr := contentTag.RemoveAttr("type")
	_ = contentTag.RemoveAttr("url")

	if typeAttr != nil {
		contentTag.CreateAttr("type", typeAttr.Value)
	}
	if idAttr != nil && typeAttr != nil {
		contentTag.CreateAttr("url", getURLAttrValue(idAttr.Value, typeAttr.Value))
	}
}

func getURLAttrValue(uuid string, t string) string {
	typeSubURL := map[string]string{
		"http://www.ft.com/ontology/content/Article":        "content",
		"http://www.ft.com/ontology/content/ImageSet":       "content",
		"http://www.ft.com/ontology/content/MediaResource":  "content",
		"http://www.ft.com/ontology/content/Video":          "content",
		"http://www.ft.com/ontology/company/PublicCompany":  "organisations",
		"http://www.ft.com/ontology/content/ContentPackage": "content",
		"http://www.ft.com/ontology/content/Content":        "content",
		"http://www.ft.com/ontology/content/Image":          "content",
		"http://www.ft.com/ontology/content/DynamicContent": "content",
		"http://www.ft.com/ontology/content/Graphic":        "content",
		"http://www.ft.com/ontology/content/Audio":          "content",
	}
	return fmt.Sprintf("http://api.ft.com/%s/%s", typeSubURL[t], uuid)
}

// transformParagraphElements apply very specific rules to <p> elements
func transformParagraphElements(input string) string {
	reBrTagP := regexp.MustCompile(`(<p>)(\\s|(<br/>))*(</p>)`)
	result := reBrTagP.ReplaceAllString(input, "")

	reEmptyP := regexp.MustCompile(`</p> +<p>`)
	result = reEmptyP.ReplaceAllString(result, "</p><p>")

	reNewLineP := regexp.MustCompile(`</p>(\\r?\\n)+<p>`)
	result = reNewLineP.ReplaceAllString(result, "</p>\n<p>")

	return result
}

// removeEmptyLines removes empty lines, removes empty <p> tags as well
func removeEmptyLines(input string) string {
	reLines := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
	return reLines.ReplaceAllString(input, "")
}

func scrollableTextExtraction(doc *etree.Document) {
	for _, block := range doc.FindElements("//scrollable-block") {
		parent := block.Parent()
		insertIndex := block.Index()
		texts := block.FindElements(".//scrollable-text")
		for _, text := range texts {
			children := text.ChildElements()
			for childIdx, el := range children {
				el.RemoveAttr("theme-style")
				parent.InsertChildAt(insertIndex+childIdx, el)
			}
			insertIndex += len(children)
		}
		parent.RemoveChild(block)
	}
}
