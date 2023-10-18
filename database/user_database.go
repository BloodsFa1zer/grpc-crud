package database

import (
	"app4/config"
	"app4/hash"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"time"
)

type User struct {
	ID        int64  `db:"ID"`
	Nickname  string `db:"nick_name"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Password  string `db:"password"`
	Role      string `db:"role"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	DeletedAt string `db:"deleted_at"`
}

type UserDatabase struct {
	Connection *sql.DB
}

func NewUserDatabase() *UserDatabase {
	cfg := config.LoadENV("config/.env")
	cfg.ParseENV()

	connStr := "user=" + cfg.UserDBName + " password=" + cfg.UserDBPassword + " dbname=" + cfg.DBName + " sslmode=disable"
	db, err := sql.Open(cfg.DriverDBName, connStr)
	if err != nil {
		log.Warn().Err(err).Msg("can`t connect to database")
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Warn().Err(err).Msg("failed to ping the database")
		return nil
	}
	log.Info().Msg("successfully connected to the database.")

	return &UserDatabase{Connection: db}
}

func (db *UserDatabase) FindByID(ID int64) (*User, error) {
	sqlSelect := `SELECT * FROM Users WHERE ID = ? AND deleted_at == 'NULL'`
	var num sql.NullInt64
	var selectedUser User

	row := db.Connection.QueryRow(sqlSelect, ID)
	err := row.Scan(&selectedUser.ID, &selectedUser.Nickname, &selectedUser.FirstName,
		&selectedUser.LastName, &selectedUser.Password, &selectedUser.CreatedAt,
		&selectedUser.UpdatedAt, &selectedUser.DeletedAt, &selectedUser.Role, &num)
	// selectedUser.Rating = num.Int64
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		log.Warn().Err(err).Msg(" can`t find user")
		return nil, err
	}

	return &selectedUser, nil
}

func (db *UserDatabase) InsertUser(user User) (int64, error) {
	formattedTime := time.Now().Format("2006.01.02 15:04")

	sqlInsert := "INSERT INTO Users (nick_name, first_name, last_name, password, created_at) VALUES (?, ?, ?, ?, ?)"

	hashedPassword := hash.Hash(user.Password)

	result, err := db.Connection.Exec(sqlInsert, user.Nickname, user.FirstName, user.LastName, hashedPassword, formattedTime)
	if err != nil {
		log.Warn().Err(err).Msg(" can`t insert user")
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		log.Warn().Err(err).Msg(" can`t find userID")
		return 0, err
	}

	return userID, nil
}

func (db *UserDatabase) UpdateUser(ID int64, user User) (int64, error) {

	hashedPassword := hash.Hash(user.Password)

	formattedTime := time.Now().Format("2006.01.02 15:04")
	sqlUpdate := "UPDATE Users SET nick_name = ?, first_name = ?, last_name = ?, password = ?, updated_at = ? WHERE ID = ? AND deleted_at == 'NULL'"

	result, err := db.Connection.Exec(sqlUpdate, user.Nickname, user.FirstName, user.LastName, hashedPassword, formattedTime, ID)
	if err != nil {
		log.Warn().Err(err).Msg(" can`t update user`s data")
		return 0, err
	}

	affectedRow, err := result.RowsAffected()
	if err != nil {
		log.Warn().Err(err).Msg(" error getting affected rows")
		return 0, err
	}

	if affectedRow == 0 {
		log.Warn().Msg(" no rows affected")
		return 0, sql.ErrNoRows
	}

	return ID, nil
}

func (db *UserDatabase) FindUsers() (*[]User, error) {

	sqlSelect := "SELECT * FROM Users WHERE deleted_at == 'NULL'"
	rows, err := db.Connection.Query(sqlSelect)
	if err != nil {
		log.Warn().Err(err).Msg(" can`t find users")
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var singleUser User
		err := rows.Scan(&singleUser.ID, &singleUser.Nickname, &singleUser.FirstName,
			&singleUser.LastName, &singleUser.Password, &singleUser.CreatedAt,
			&singleUser.UpdatedAt, &singleUser.DeletedAt, &singleUser.Role)
		if err != nil {
			return nil, err
		}

		users = append(users, singleUser)
	}
	return &users, nil
}

func (db *UserDatabase) DeleteUserByID(ID int64) error {
	formattedTime := time.Now().Format("2006.01.02 15:04")
	sqlSoftDelete := "UPDATE Users SET (deleted_at) = (?) WHERE ID = ? AND deleted_at == 'NULL'"

	result, err := db.Connection.Exec(sqlSoftDelete, formattedTime, ID)
	if err != nil {
		log.Warn().Err(err).Msg(" can`t delete user`s data")
		return err
	}

	affectedRow, err := result.RowsAffected()
	if err != nil {
		log.Warn().Err(err).Msg(" error getting affected rows")
		return err
	}

	if affectedRow == 0 {
		log.Warn().Msg(" no rows affected")
		return sql.ErrNoRows
	}

	return nil
}
