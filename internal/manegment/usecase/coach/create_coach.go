package coach

type InputDTO struct{}

type CreateCoach struct {
	Input InputDTO
}

func NewCreateCoach(input InputDTO) (*CreateCoach, error) {
	return &CreateCoach{
		Input: input,
	}, nil
}
