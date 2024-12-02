package mocks

import "github.com/stretchr/testify/mock"

type MockLogger struct {
	mock.Mock
}

func (l *MockLogger) InitLogger()                                 {}
func (l *MockLogger) Info(args ...interface{})                    {}
func (l *MockLogger) Infof(template string, args ...interface{})  {}
func (l *MockLogger) Error(args ...interface{})                   {}
func (l *MockLogger) Errorf(template string, args ...interface{}) {}
func (l *MockLogger) Fatal(args ...interface{})                   {}
func (l *MockLogger) Fatalf(template string, args ...interface{}) {}
