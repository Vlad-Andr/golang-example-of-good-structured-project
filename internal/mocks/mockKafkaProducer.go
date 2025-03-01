package mocks

import "github.com/stretchr/testify/mock"

type MockKafkaProducer struct {
	mock.Mock
}

func (m *MockKafkaProducer) Produce(key string, value []byte) error {
	args := m.Called(key, value)
	return args.Error(0)
}

func (m *MockKafkaProducer) Close() error {
	args := m.Called()
	return args.Error(0)
}
