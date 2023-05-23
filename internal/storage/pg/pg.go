package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	errNotUnique         = errors.New("not unique original url")
	errNotUniqueShortUrl = errors.New("not unique short url")
	errNowSuchUrl        = errors.New("unknown url")
)

type Postgres struct {
	pg string
}

func NewPostgres() *Postgres {
	return &Postgres{
		pg: "postrgres",
	}
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment")
	}

	db, err := sql.Open("postgres", os.Getenv("PG_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func (p *Postgres) SaveData(og, short string) (string, error) {
	if err := p.uniqueCheck(og, short); err != nil {
		return "", err
	}

	db := createConnection()
	defer db.Close()

	preparedStatement := `INSERT INTO urls (originalurl, shorturl) VALUES ($1, $2) RETURNING shorturl`

	var ret string

	err := db.QueryRow(preparedStatement, og, short).Scan(&ret)
	if err != nil {
		return "", err
	}

	return ret, nil
}

func (p *Postgres) GetData(short string) (string, error) {
	db := createConnection()
	defer db.Close()

	preparedStatement := `SELECT originalurl FROM urls WHERE shorturl=$1`

	var ret string

	err := db.QueryRow(preparedStatement, short).Scan(&ret)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errNowSuchUrl
		}
		return "", err
	}

	return ret, nil
}

func (p *Postgres) uniqueCheck(og, short string) error {
	db := createConnection()
	defer db.Close()

	preparedStatement := `SELECT id FROM urls WHERE originalurl=$1`
	var id int

	err := db.QueryRow(preparedStatement, og).Scan(&id)
	if err != sql.ErrNoRows {
		return err
	}

	preparedStatement = `SELECT id FROM urls WHERE shorturl=$1`

	err = db.QueryRow(preparedStatement, short).Scan(&id)
	if err != sql.ErrNoRows {
		return err
	}

	return nil
}
