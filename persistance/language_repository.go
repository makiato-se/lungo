package persistance

type LanguageRepository struct {

}

type Word struct {
	NativeWord string
	TranslatedWord string
}

func (repo *LanguageRepository) GetWordsForLanguage(language string) []Word {
	return []Word{
		{"Hola!", "Hello"},
	}
}