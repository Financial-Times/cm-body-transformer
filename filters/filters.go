package filters

import (
	"html"
	"regexp"
	"strings"
)

type Filter func(string) string

func Apply(text string, transformers ...Filter) string {
	current := text
	for _, transformer := range transformers {
		current = transformer(current)
	}
	return current
}

func DefaultContentFilters() []Filter {
	return []Filter{
		RemovePullQuoteTag,
		RemoveWebPullQuoteTag,
		RemoveTableTag,
		RemovePromoBoxTag,
		RemoveWebInlinePictureTag,
		RemoveHTMLEntity,
		RemoveGenericTags,
		strings.TrimSpace,
		DedupSpaces,
	}
}

// ReplaceMatchedText replaces every substring in the src string that matches the provided regex with the provided repl
func ReplaceMatchedText(regex, src, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(src, repl)
}

// RemoveMatchedText it substitutes every substring in src that matches provided regexpr with a whitespace.
// This is to avoid megring words.
// In the process it deduplicates all whitespaces. See DedupSpaces for more info.
func RemoveMatchedText(regex, src string) string {
	result := ReplaceMatchedText(regex, src, " ")
	return strings.TrimSpace(DedupSpaces(result))
}

// DeleteMatchedText removes every substring in src that matches a regexpr
// This filter can merge words together. To avoid that use RemoveMatchedText.
func DeleteMatchedText(regex, src string) string {
	return ReplaceMatchedText(regex, src, "")
}

// DedupSpaces squashes long chains of whitespaces to a single whitespace (the last one in the chain).
func DedupSpaces(src string) string {
	return ReplaceMatchedText(`(\s)+`, src, "$1")
}

func RemovePullQuoteTag(input string) string {
	return DeleteMatchedText(`(?s)<pull-quote.*?</pull-quote>`, input)
}

func RemoveWebPullQuoteTag(input string) string {
	return DeleteMatchedText(`(?s)<web-pull-quote.*?</web-pull-quote>`, input)
}

func RemoveTableTag(input string) string {
	return DeleteMatchedText(`(?s)<table.*?</table>`, input)
}

func RemovePromoBoxTag(input string) string {
	return DeleteMatchedText(`(?s)<promo-box.*?</promo-box>`, input)
}

func RemoveWebInlinePictureTag(input string) string {
	return DeleteMatchedText(`(?s)<web-inline-picture.*?</web-inline-picture>`, input)
}

func RemoveHTMLEntity(input string) string {
	text := RemoveMatchedText(`&nbsp;`, input)
	return html.UnescapeString(text)
}

func RemoveGenericTags(input string) string {
	return RemoveMatchedText(`<[^>]*>`, input)
}
