package choice

type Choice struct {
	ID         string `json:"id"`
	QuestionID string `json:"question_id"`
	ChoiceText string `json:"choice_text"`
	IsCorrect  bool   `json:"is_correct"`
}
