package orm

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseURL struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Hash string `gorm:"uniqueIndex" json:"hash"`
	URL  string `json:"url"`
}

type RecordIP struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	IP   string `json:"ip"`
	Hash string `json:"hash"`
}

type ShortURL struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type DB struct {
	db *gorm.DB
}

func (db *DB) Init() {
	db.db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	err := db.db.AutoMigrate(&DatabaseURL{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = db.db.AutoMigrate(&RecordIP{})

	if err != nil {
		log.Fatal("failed to connect database")
	}
}

func (db *DB) Create(hash string, url string) error {
	// check url existed
	var dbURL DatabaseURL
	tx := db.db.Where("Hash = ?", hash).First(&dbURL)
	if tx.Error == nil {
		return nil
	}

	// create new url and increase id
	tx = db.db.Create(&DatabaseURL{Hash: hash, URL: url})
	// tx = db.db.Create(&DatabaseURL{Hash: hash, URL: url})
	return tx.Error
}

func (db *DB) Find(hash string) (string, error) {
	var dbURL DatabaseURL
	tx := db.db.Where("Hash = ?", hash).First(&dbURL)
	return dbURL.URL, tx.Error
}

func (db *DB) FindAll(host string) ([]ShortURL, error) {
	var dbURL []DatabaseURL
	tx := db.db.Find(&dbURL)

	// convert to short url
	var shortURL []ShortURL
	for _, url := range dbURL {
		short := host + "/" + url.Hash
		shortURL = append(shortURL, ShortURL{ShortURL: short, OriginalURL: url.URL})
	}

	return shortURL, tx.Error
}

func (db *DB) CreateRecordIP(ip string, hash string) error {
	// create new url and increase id
	tx := db.db.Create(&RecordIP{IP: ip, Hash: hash})
	return tx.Error
}

func (db *DB) FindAllRecordIP() ([]RecordIP, error) {
	var dbURL []RecordIP
	tx := db.db.Find(&dbURL)

	// convert to short url
	var shortURL []RecordIP
	for _, url := range dbURL {
		shortURL = append(shortURL, RecordIP{IP: url.IP, Hash: url.Hash})
	}

	return shortURL, tx.Error
}
