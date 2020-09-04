package validate

import (
	"regexp"

	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var v = validator.New()

// 中文
var zh = zhongwen.New()
var uni = ut.New(zh, zh)
var trans, _ = uni.GetTranslator("zh")

func init() {
	zh_translations.RegisterDefaultTranslations(v, trans)
	// 注册手机号码验证
	v.RegisterValidation("phone", Phone)
	// 注册翻译
	v.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
		return ut.Add("phone", "手机号码格式错误!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone", fe.Field())
		return t
	})
}

// Validate Validate
func Validate(i interface{}) string {
	err := v.Struct(i)
	if err != nil {
		errMap := err.(validator.ValidationErrors).Translate(trans)
		for _, v := range errMap {
			return v
		}
	}

	return ""
}

// Phone check phone
func Phone(fl validator.FieldLevel) bool {
	matchPhone := regexp.MustCompile(`^\d{11}$`)
	return matchPhone.MatchString(fl.Field().String())
}
