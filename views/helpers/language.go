package helpers

import (
	"context"
	"strings"
	"time"

	emojiflag "github.com/jayco/go-emoji-flag"
	"github.com/sersh88/timeago"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

var englishTag = display.English.Languages()

type LanguageInformation struct {
	Code        string
	EnglishName string
	LocalName   string
	Flag        string
}

func Language(ctx context.Context) string {
	return translator(ctx).Code().String()
}

func SupportedLanguages(ctx context.Context) []language.Tag {
	return []language.Tag{
		language.Dutch,
		language.English,
		language.French,
		language.German,
		language.Indonesian,
		language.Italian,
		language.Norwegian,
		language.Persian,
		language.Russian,
	}
}

func ToLanguageInformation(code string) LanguageInformation {
	cc := code
	if strings.Contains(cc, "-") {
		cc = strings.Split(cc, "-")[1]
	}

	if cc == "en" {
		cc = "us"
	}

	l := LanguageInformation{
		Code: code,
		Flag: emojiflag.GetFlag(cc),
	}

	if l.Flag == "" {
		l.Flag = "👽"
	}

	localTag := language.MustParse(code)
	l.LocalName = display.Self.Name(localTag)
	l.EnglishName = englishTag.Name(localTag)

	return l
}

func RelativeDate(ctx context.Context, t *time.Time) string {
	return timeago.New(*t).WithLocale(Language(ctx)).Format()
}