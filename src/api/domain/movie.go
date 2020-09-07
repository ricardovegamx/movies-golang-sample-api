package domain

type Movie struct {
	Id          uint64   `json:"id"`
	Title       string   `json:"title" binding:"-"`
	ReleaseDate string   `json:"release_date"`
	Categories  []string `json:"categories"`
	Runtime     int16    `json:"runtime"`
	UserScore   float32  `json:"user_score"`
	Overview    string   `json:"overview"`
	Director    string   `json:"director"`
}
