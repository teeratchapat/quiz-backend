package choice

type Service struct {
	Repo *Repository
}

func (s *Service) GetChoices(questionID string) ([]Choice, error) {
	return s.Repo.GetByQuestionID(questionID)
}

func (s *Service) CreateChoice(c Choice) error {
	return s.Repo.Create(c)
}
