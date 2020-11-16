package service

import (
	"online-tests/delivery/command"
	"online-tests/delivery/response"
	"online-tests/domain"
	"online-tests/domain/model"
)

type testsService struct {
	testsRepo domain.TestsRepository
}

func (es *testsService) AddTest(cmd *command.AddTestCmd) error {
	var (
		err  error
		test = model.NewTestModel(cmd)
	)

	es.testsRepo.Create(&test)

	return err
}

func (es *testsService) GetTest(testID uint) (*response.GetTestResp, error) {
	var err error
	return nil, err
}

func NewTestService(er domain.TestsRepository) domain.TestsService {
	es := &testsService{
		testsRepo: er,
	}

	return es
}
