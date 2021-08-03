package filters

import (
	"html"
	"regexp"
	"strings"
)

var (
	nbspRegex                = regexp.MustCompile(`&nbsp;`)
	pullTagRegex             = regexp.MustCompile(`(?s)<pull-quote.*?</pull-quote>`)
	webPullTagRegex          = regexp.MustCompile(`(?s)<web-pull-quote.*?</web-pull-quote>`)
	tableTagRegex            = regexp.MustCompile(`(?s)<table.*?</table>`)
	promoBoxTagRegex         = regexp.MustCompile(`(?s)<promo-box.*?</promo-box>`)
	webInlinePictureTagRegex = regexp.MustCompile(`(?s)<web-inline-picture.*?</web-inline-picture>`)
	tagRegex                 = regexp.MustCompile(`<[^>]*>`)
	duplicateWhiteSpaceRegex = regexp.MustCompile(`\s+`)
)

type TextTransformer func(string) string

func TransformText(text string, transformers ...TextTransformer) string {
	current := text
	for _, transformer := range transformers {
		current = transformer(current)
	}
	return current
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

func PullTagTransformer(input string) string {
	return pullTagRegex.ReplaceAllString(input, "")
}

func WebPullTagTransformer(input string) string {
	return webPullTagRegex.ReplaceAllString(input, "")
}

func TableTagTransformer(input string) string {
	return tableTagRegex.ReplaceAllString(input, "")
}

func PromoBoxTagTransformer(input string) string {
	return promoBoxTagRegex.ReplaceAllString(input, "")
}

func WebInlinePictureTagTransformer(input string) string {
	return webInlinePictureTagRegex.ReplaceAllString(input, "")
}

func HtmlEntityTransformer(input string) string {
	text := nbspRegex.ReplaceAllString(input, " ")
	return html.UnescapeString(text)
}

func TagsRemover(input string) string {
	result := tagRegex.ReplaceAllString(input, " ")
	return OuterSpaceTrimmer(DuplicateWhiteSpaceRemover(result))
}

func OuterSpaceTrimmer(input string) string {
	return strings.TrimSpace(input)
}

func DuplicateWhiteSpaceRemover(input string) string {
	return duplicateWhiteSpaceRegex.ReplaceAllString(input, " ")
}

func DefaultValueTransformer(input string) string {
	if input == "" {
		return "."
	}
	return input
}
