package question

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) GetAll() ([]Question, error) {
	rows, err := r.DB.Query("SELECT id, question_text FROM questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []Question
	for rows.Next() {
		var q Question
		if err := rows.Scan(&q.ID, &q.QuestionText); err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func (r *Repository) GetByID(id string) (*Question, error) {
	row := r.DB.QueryRow("SELECT id, question_text FROM questions WHERE id = $1", id)
	var q Question
	if err := row.Scan(&q.ID, &q.QuestionText); err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *Repository) Create(questionText string) error {
	_, err := r.DB.Exec(
		"INSERT INTO questions (question_text) VALUES ($1)",
		questionText,
	)
	return err
}

func (r *Repository) CreateReturningID(text string) (string, error) {
	var id string
	err := r.DB.QueryRow(
		`INSERT INTO questions (question_text) VALUES ($1) RETURNING id`,
		text,
	).Scan(&id)

	return id, err
}

func (r *Repository) DeleteChoicesByQuestionID(id string) error {
	_, err := r.DB.Exec("DELETE FROM choices WHERE question_id = $1", id)
	return err
}

func (r *Repository) DeleteQuestion(id string) error {
	_, err := r.DB.Exec("DELETE FROM questions WHERE id = $1", id)
	return err
}
