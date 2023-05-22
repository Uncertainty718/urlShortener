package storage

type Storage interface {
	SaveData(og, short string) error
	GetData(short string) (string, error)
}