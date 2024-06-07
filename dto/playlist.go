package dto

type CreatePlaylistRequest struct {
	Title      string   `json:"title" validate:"required"`
	AlbumCover string   `json:"album_cover" validate:"required"`
	TrackIDs   []string `json:"track_ids"`
}

type UpdatePlaylistRequest struct {
	Title      *string `json:"title"`
	AlbumCover *string `json:"album_cover"`
}

type UpdatePlaylistTrackRequest struct {
	TrackID  string `json:"id" validate:"required"`
	IsDelete *bool   `json:"is_delete" validate:"required"`
}
