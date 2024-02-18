// models/translation.go

package model

// Translation represents a translation entity.
type Translation struct {
	Key         string `json:"key"`
	Lang        string `json:"lang"`
	Translation string `json:"translation"`
}

// TranslationTable represents the structure of the translation table in the database.
const TranslationTable = `
CREATE TABLE IF NOT EXISTS translations (
    key TEXT NOT NULL,
    lang TEXT NOT NULL,
    translation TEXT NOT NULL,
    CONSTRAINT pk_translation PRIMARY KEY (key, lang)
);
`
