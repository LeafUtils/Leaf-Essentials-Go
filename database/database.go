package database

import (
	"encoding/json"
	"time"
)

// Document structure
type Document struct {
	ID   int64                  `json:"id"`
	Data map[string]interface{} `json:"data"`
}

// Database structure
type Database struct {
	documents map[int64]Document
}

// NewDatabase creates a new instance of the Database
func NewDatabase() *Database {
	return &Database{
		documents: make(map[int64]Document),
	}
}

// findFirst finds the first document matching the criteria
func (db *Database) FindFirst(criteria map[string]interface{}) *Document {
	for _, doc := range db.documents {
		if matchesCriteria(doc.Data, criteria) {
			return &doc
		}
	}
	return nil
}

// findDocuments returns all documents matching the criteria
func (db *Database) FindDocuments(criteria map[string]interface{}) []Document {
	var docs []Document
	for _, doc := range db.documents {
		if matchesCriteria(doc.Data, criteria) {
			docs = append(docs, doc)
		}
	}
	return docs
}

// matchesCriteria checks if a document matches all criteria
func matchesCriteria(docData, criteria map[string]interface{}) bool {
	if criteria == nil || len(criteria) == 0 {
		return true
	}
	for key, value := range criteria {
		if docValue, ok := docData[key]; !ok || docValue != value {
			return false
		}
	}
	return true
}

// insertDocument inserts a new document into the database
func (db *Database) InsertDocument(data map[string]interface{}) Document {
	// Create a new document with current timestamp as ID
	doc := Document{
		ID:   time.Now().UnixNano() / int64(time.Millisecond),
		Data: data,
	}

	// Insert the document into the map
	db.documents[doc.ID] = doc
	return doc
}

// getByID retrieves a document by its ID
func (db *Database) GetByID(id int64) *Document {
	doc, exists := db.documents[id]
	if !exists {
		return nil
	}
	return &doc
}

// deleteByID deletes a document by its ID
func (db *Database) DeleteByID(id int64) {
	delete(db.documents, id)
}

// overwriteByID overwrites an existing document with a new one
func (db *Database) OverwriteByID(id int64, newData map[string]interface{}) *Document {
	if _, exists := db.documents[id]; !exists {
		return nil
	}

	doc := Document{
		ID:   id,
		Data: newData,
	}
	db.documents[id] = doc
	return &doc
}

// exportToJSON exports the database content to a JSON string
func (db *Database) ExportToJSON() (string, error) {
	docs := db.FindDocuments(nil) // Fetch all documents without criteria
	jsonData, err := json.Marshal(docs)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// import data from JSON
func (db *Database) ImportFromJSON(jsonStr string) error {
	var docs []Document
	if err := json.Unmarshal([]byte(jsonStr), &docs); err != nil {
		return err
	}

	for _, doc := range docs {
		db.documents[doc.ID] = doc
	}
	return nil
}
