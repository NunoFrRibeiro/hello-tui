package constants

type Language struct {
	Name        string
	Code        string
	Translation string
}

var AvailableLanguages = []Language{
	{Name: "English", Code: "en", Translation: "Hello, World"},
  {Name: "French", Code: "fr", Translation: "Bonjour, le monde!"},
	{Name: "Spanish", Code: "es", Translation: "¡Hola, Mundo!"},
	{Name: "German", Code: "de", Translation: "Hallo, Welt!"},
	{Name: "Portuguese", Code: "pt", Translation: "Olá, Mundo!"},
	{Name: "Italian", Code: "it", Translation: "Ciao, Mondo!"},
	{Name: "Japanese", Code: "ja", Translation: "こんにちは、世界！"},
	{Name: "Chinese", Code: "zh", Translation: "你好，世界！"},
}
