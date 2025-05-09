package question

type Service struct {
	Repo *Repository
}

func (s *Service) GetQuestions() ([]Question, error) {
	return s.Repo.GetAll()
}

func (s *Service) GetQuestionByID(id string) (*Question, error) {
	return s.Repo.GetByID(id)
}

func (s *Service) CreateQuestion(text string) error {
	return s.Repo.Create(text)
}

func (s *Service) CreateQuestionReturningID(text string) (string, error) {
	return s.Repo.CreateReturningID(text)
}

func (s *Service) DeleteQuestion(id string) error {
	// ลบ choices ก่อน
	if err := s.Repo.DeleteChoicesByQuestionID(id); err != nil {
		return err
	}
	return s.Repo.DeleteQuestion(id)
}
