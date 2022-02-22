package api

// CookService contains the methods of the cook service
type CookService interface {
	New(cook NewCookRequest) error
	Get(eventMessageID string) (NewCookRequest, error)
}

// CookRepository is what lets our service do db operations without knowing anything about the implementation
type CookRepository interface {
	CreateCook(NewCookRequest) error
	GetCook(eventMessageID string) (NewCookRequest, error)
}

type cookService struct {
	storage CookRepository
}

func NewCookService(cookRepo CookRepository) CookService {
	return &cookService{
		storage: cookRepo,
	}
}

func (c *cookService) New(cook NewCookRequest) error {
	c.storage.CreateCook(cook)
	return nil
}

func (c *cookService) Get(eventMessageID string) (NewCookRequest, error) {
	return c.storage.GetCook(eventMessageID)
}
