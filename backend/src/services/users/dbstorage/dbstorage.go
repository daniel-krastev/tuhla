package dbstorage

import(
	"tuhla/common/db"
)

type DBStorage struct{
	db *db.DBConnection
}

func New(db *db.DBConnection) *DBStorage {
	return &DBStorage{db}
}

func (s *DBStorage) CreateUser(requestID, userID string) (string, error) {
	return "", nil
}
