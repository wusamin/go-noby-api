package main

type Persona string

const (
	Normal         Persona = "0"
	TsundereFemale Persona = "1"
	TsundereMale   Persona = "2"
	God                    = "3"
)

// NobyRequest represents request parameters of noby API.
type NobyRequest struct {
	Appkey  string
	Mail    string
	Pass    string
	Text    string
	Lat     float64
	Lng     float64
	Study   bool
	Persona Persona
	Ending  string
}

// NobyResponse represents response of noby API.
type NobyResponse struct {
	CommandId    string     `json:"commandId"`
	CommandName  string     `json:"commandName"`
	Text         string     `json:"text"`
	Type         string     `json:"type"`
	Mood         float64    `json:"mood"`
	Negaposi     float64    `json:"negaposi"`
	NegaposiList []Negaposi `json:"negaposiList"`
	Emotion      Emotion    `json:"emotion"`
	Word         Word       `json:"word"`
	EmotionList  []Emotion  `json:"emotionList"`
	WordList     []Word     `json:"wordList"`
	Art          string     `json:"art"`
	Org          string     `json:"org"`
	Psn          string     `json:"psn"`
	Loc          string     `json:"loc"`
	Dat          string     `json:"dat"`
	Tim          string     `json:"tim"`
}

// Negaposi represents value of "Negaposi".
type Negaposi struct {
	Word  string  `json:"word"`
	Score float64 `json:"score"`
}

// Emotion represents value of "Emotion".
type Emotion struct {
	Word        string  `json:"word"`
	LikeDislike float64 `json:"likeDislike"`
	JoySad      float64 `json:"joySad"`
	AngerFear   float64 `json:"angerFear"`
}

// Word represents value of "Word".
type Word struct {
	Feature string `json:"feature"`
	Start   string `json:"start"`
	Surface string `json:"surface"`
}
