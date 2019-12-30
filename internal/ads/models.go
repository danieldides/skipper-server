package ads

import "database/sql"

// Ad contains information about the individual ads
type Ad struct {
	ID        int `json:"id" db:"id"`
	VideoID   int `json:"video_id" db:"video_id"`
	Beginning int `json:"beginning" db:"beginning"`
	Duration  int `json:"duration" db:"duration"`
	Score     int `json:"score" db:"score"`
}

// GetAdsByVideoID gets a Video by a given ID
func GetAdsByVideoID(db *sql.DB, id string) ([]Ad, error) {
	var ads []Ad

	q := `SELECT id, video_id, beginning, duration, score FROM ads WHERE video_id = ?`
	rows, err := db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ad Ad
		err = rows.Scan(&ad.ID, &ad.VideoID, &ad.Beginning,
			&ad.Duration, &ad.Score)
		if err != nil {
			return nil, err
		}
		ads = append(ads, ad)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ads, nil
}
