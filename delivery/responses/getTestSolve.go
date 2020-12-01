package responses

import (
	"online-tests/delivery/models/testsolvemodel"
	"online-tests/domain/domainmodel"
)

type TestSolveModel struct {
	Test *testsolvemodel.Test `json:"test"`
}

func NewTestSolveModelResp(test *domainmodel.Test) *TestSolveModel {
	return &TestSolveModel{
		Test: NewTestSolveModel(test),
	}
}

func NewTestSolveModel(domainTest *domainmodel.Test) *testsolvemodel.Test {
	return &testsolvemodel.Test{
		ID:        domainTest.ID,
		Title:     domainTest.Title,
		Questions: *newTestSolveQuestionsArray(&domainTest.Questions, domainTest.NumTestOfQuestions, domainTest.ID),
	}
}

func newTestSolveQuestion(q *domainmodel.Question, tID uint64) *testsolvemodel.Question {
	return &testsolvemodel.Question{
		ID:       q.ID,
		Question: q.Question,
		Answers:  *newTestSolveAnswersArray(&q.Answers, q.ID),
	}
}

func newTestSolveQuestionsArray(qArr *[]domainmodel.Question, numOfQuestions uint, tID uint64) *[]testsolvemodel.Question {
	questions := []testsolvemodel.Question{}
	for _, q := range *qArr {
		questions = append(questions, *newTestSolveQuestion(&q, tID))
	}

	return &questions
}

func newTestSolveAnswer(a *domainmodel.Answer, qID uint64) *testsolvemodel.Answer {
	return &testsolvemodel.Answer{
		ID:     a.ID,
		Answer: a.Answer,
	}
}

func newTestSolveAnswersArray(aArr *[]domainmodel.Answer, qID uint64) *[]testsolvemodel.Answer {
	answers := []testsolvemodel.Answer{}
	for _, a := range *aArr {
		answers = append(answers, *newTestSolveAnswer(&a, qID))
	}

	return &answers
}
