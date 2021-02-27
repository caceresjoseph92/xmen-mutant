package increasing

type PersonCounterService struct{}

func NewPersonCounterService() PersonCounterService {
	return PersonCounterService{}
}

func (s PersonCounterService) Increase(id string) error {
	return nil
}
