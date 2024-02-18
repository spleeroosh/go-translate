// database/translation.go

package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/spleeroosh/go-translate/model"
)

// TranslationStore represents a store for translations in the database.
type TranslationStore struct {
	DB *sql.DB
}

// NewTranslationStore initializes a new TranslationStore.
func NewTranslationStore(db *sql.DB) *TranslationStore {
	return &TranslationStore{DB: db}
}

// GetAllTranslations retrieves all translations from the database.
func (ts *TranslationStore) GetAllTranslations(ctx context.Context) ([]model.Translation, error) {
	rows, err := ts.DB.QueryContext(ctx, "SELECT key, lang, translation FROM translations")
	if err != nil {
		log.Println("Error retrieving translations:", err)
		return nil, err
	}
	defer rows.Close()

	var translations []model.Translation
	for rows.Next() {
		var translation model.Translation
		if err := rows.Scan(&translation.Key, &translation.Lang, &translation.Translation); err != nil {
			log.Println("Error scanning translation row:", err)
			continue
		}
		translations = append(translations, translation)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over translation rows:", err)
		return nil, err
	}

	return translations, nil
}

// AddTranslation adds a new translation to the database.
func (ts *TranslationStore) AddTranslation(ctx context.Context, translation model.Translation) error {
	_, err := ts.DB.ExecContext(ctx, "INSERT INTO translations (key, lang, translation) VALUES ($1, $2, $3)", translation.Key, translation.Lang, translation.Translation)
	if err != nil {
		log.Println("Error adding translation:", err)
		return err
	}
	return nil
}
