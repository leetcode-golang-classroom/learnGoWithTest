package hello

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

var languageMap = map[string]string{
	"":        englishPrefix,
	"French":  frenchPrefix,
	"Spanish": spanishPrefix,
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return languageMap[language] + name
}
