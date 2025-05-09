package question

type Question struct {
	ID           string `json:"id"`
	QuestionText string `json:"question_text"`
	CreatedAt    string `json:"created_at"`
}
