package postgres

import (
	"database/sql"
	"errors"
	"external-api/internal/config"
	"external-api/internal/models"
	"fmt"
	_ "github.com/lib/pq" // init postgres driver
)

type InfoSonger interface {
	GetSongDetails(group, name string) (*models.SongDetails, error)
}

type Database struct {
	db *sql.DB
}

func New(cfg config.Database) (InfoSonger, error) {
	const op = "internal.database.postgres.New"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: open: %w", op, err)
	}

	return &Database{db: db}, nil
}

func (d *Database) GetSongDetails(group, name string) (*models.SongDetails, error) {
	const op = "internal.database.postgres.GetSongDetails"
	query := "SELECT release_date, text, link FROM songs_info WHERE group_name = $1 AND name = $2"
	var song models.SongDetails

	err := d.db.QueryRow(query, group, name).
		Scan(&song.ReleaseDate, &song.Text, &song.Link)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: no song found", op)
		}
		return nil, fmt.Errorf("%s: query row scan %w", op, err)
	}
	return &song, nil
}
