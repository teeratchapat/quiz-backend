package choice

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) GetByQuestionID(questionID string) ([]Choice, error) {
	rows, err := r.DB.Query("SELECT id, question_id, choice_text, is_correct FROM choices WHERE question_id = $1", questionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var choices []Choice
	for rows.Next() {
		var c Choice
		if err := rows.Scan(&c.ID, &c.QuestionID, &c.ChoiceText, &c.IsCorrect); err != nil {
			return nil, err
		}
		choices = append(choices, c)
	}
	return choices, nil
}

func (r *Repository) Create(choice Choice) error {
	_, err := r.DB.Exec("INSERT INTO choices (question_id, choice_text, is_correct) VALUES ($1, $2, $3)",
		choice.QuestionID, choice.ChoiceText, choice.IsCorrect)
	return err
}
