package mop_user

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
)

func getLastInsertedID(db *gorm.DB) (int, error) {
	lastInsertedID := 0
	lastInsertIDQuery := `SELECT LAST_INSERT_ID()`

	if err := db.Debug().Raw(lastInsertIDQuery).Scan(&lastInsertedID).Error; err != nil {
		log.Printf("error occurred while fetching last insert ID: %v\n", err)
		return 0, ErrInternal
	}

	return lastInsertedID, nil
}

func isUniqueConstraintViolation(err error) bool {
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == 1062
}
