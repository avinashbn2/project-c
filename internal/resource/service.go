package resource

import (
	"fmt"
	"log"
	"time"
)

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo}
}

func (s *service) get(qp *QueryWithUser) (*ResourceResponse, error) {

	data, err := s.repo.get(qp)

	if err != nil {
		log.Fatal("service", err)
		return nil, err

	}
	fmt.Println("SERVIUCE")
	return data, nil
}

func (s *service) getByUserLikes(qp *QueryWithUser) (*ResourceResponse, error) {

	data, err := s.repo.getByUserLikes(qp)

	if err != nil {
		log.Fatal("service", err)
		return nil, err

	}
	fmt.Println("SERVIUCE")
	return data, nil
}

func (s *service) getByTrending(qp *QueryWithUser) (*ResourceResponse, error) {

	data, err := s.repo.getByTrending(qp)

	if err != nil {
		log.Fatal("service", err)
		return nil, err

	}
	fmt.Println("SERVIUCE")
	return data, nil
}

func (s *service) getByID(id int) (*Resources, error) {

	data, err := s.repo.getByID(id)

	if err != nil {
		log.Fatal("service", err)
		return nil, err

	}

	return data, nil
}

func (s *service) create(resource *ResourceItem) (*ResourceItem, error) {

	resource.CreatedAt = time.Now().UTC()
	resource.UpdatedAt = time.Now().UTC()

	data, err := s.repo.create(resource)
	if err != nil {
		log.Fatal("service", err)
		return nil, err
	}

	return data, nil
}

func (s *service) update(resource *ResourceItem) (*ResourceItem, error) {

	resource.UpdatedAt = time.Now().UTC()
	data, err := s.repo.update(resource)
	if err != nil {
		log.Fatal("service", err)
		return nil, err
	}
	return data, nil
}
