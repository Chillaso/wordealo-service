package model

type Letter struct {
	Letter   string       `json:"letter" binding:"required"`
	Position map[int]bool `json:"position"`
}

// Create a struct to hold the request body
type Wordealo struct {
	GameMode   *int     `json:"gameMode,omitempty"`
	WordLength int      `json:"wordLength" binding:"required"`
	Correct    []Letter `json:"correct"`
	Incorrect  []Letter `json:"incorrect"`
}
