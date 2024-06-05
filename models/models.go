package models

import (
	"buah/config"
	"buah/entities"
)

// Read data
func GetAll() []entities.Buah {
	query, err := config.DB.Query("SELECT * FROM buah")

	if err != nil {
		panic(err)
	}

	defer query.Close()

	var buah []entities.Buah

	for query.Next() {
		var fruits entities.Buah

		if err := query.Scan(&fruits.Id, &fruits.Nama); err != nil {
			panic(err)
		}

		buah = append(buah, fruits)
	}
	return buah
}

//Create data

func Create(buah entities.Buah) bool {
	res, err := config.DB.Exec("INSERT INTO buah (Nama) VALUES (?)", buah.Nama)

	if err != nil {
		panic(err)
	}

	lastinsertid, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastinsertid > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM buah WHERE Id = ?", id)

	return err
}
