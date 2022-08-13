package i18n

import (
	"wa/global"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func HelloPerson(lang string, name string) string {
	// Create a new localizer.
	var localizer = i18n.NewLocalizer(global.Bundle, lang)

	// Set title message.
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "HelloPerson", // set translation ID
			Other: "Hello *{{.Name}}* \n" +
				"This is the digital assistant of Kottouf's procurement manager", // set default translation
		},
		TemplateData: map[string]string{
			"Name": name,
		},
		PluralCount: nil,
	})
}

var WhoIsThis = func(lang string, name string) string {
	// Create a new localizer.
	var localizer = i18n.NewLocalizer(global.Bundle, lang)

	return localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "WhoIsThis",    // set translation ID
			Other: "Who are you?", // set default translation
		},
		TemplateData: map[string]string{
			"Name": name,
		},
	})

	//var content strings.Builder
	//	content.WriteString(fmt.Sprintf("مرحبا *%v* \n", name))
	//content.WriteString(helloPerson)
}
