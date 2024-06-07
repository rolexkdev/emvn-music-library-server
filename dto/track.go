package dto

type CreateTrackRequest struct {
	Name        string `json:"name" validate:"required"`
	Title       string `json:"title" validate:"required"`
	ArtistID    string `json:"artist_id" validate:"required"`
	Album       string `json:"album" validate:"required"`
	Genre       string `json:"genre" validate:"required"`
	ReleaseDate int64  `json:"release_date" validate:"required"`
	Duration    int64  `json:"duration" validate:"required"`
	FileURL     string `json:"file_url" validate:"required"`
}

type UpdateTrackRequest struct {
	Name        *string `json:"name"`
	Title       *string `json:"title"`
	ArtistID    *string `json:"artist_id"`
	Album       *string `json:"album"`
	Genre       *string `json:"genre"`
	FileURL     *string `json:"file_url"`
}
