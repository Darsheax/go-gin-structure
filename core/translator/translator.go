package translator

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Translator struct {
	Bundle *i18n.Bundle
}

func (t *Translator) New() {

	bundle := i18n.NewBundle(language.French)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("active.en.toml")

	t.Bundle = bundle
}

func (t *Translator) Localizer(c *gin.Context) *i18n.Localizer {
	lang := c.Request.FormValue("lang")
	accept := c.Request.Header.Get("Accept-Language")
	return i18n.NewLocalizer(t.Bundle, lang, accept)
}
