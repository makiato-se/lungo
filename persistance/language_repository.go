package persistance

import "go.mongodb.org/mongo-driver/mongo"

type LanguageRepository struct {
	database *mongo.Database
}

type Word struct {
	NativeWord string
	TranslatedWord string
}

func NewLanguageRepository(context *DbContext) LanguageRepository {
	database := context.Client.Database("languages")
	return LanguageRepository{
		database: database,
	}
}

func (repo *LanguageRepository) GetWordsForLanguage(language string) []Word {
	return []Word{
		{"Hola!", "Hello"},
	}
}