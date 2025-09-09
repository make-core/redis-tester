package test_cases

import (
	"strconv"

	"github.com/make-core/redis-tester/internal/instrumented_resp_connection"
	"github.com/make-core/redis-tester/internal/resp_assertions"
	"github.com/make-core/tester-utils/logger"
)

type WaitTestCase struct {
	Replicas              int
	TimeoutInMilliseconds int
	ExpectedMessage       int
}

func (t WaitTestCase) Run(client *instrumented_resp_connection.InstrumentedRespConnection, logger *logger.Logger) error {
	commandTest := SendCommandTestCase{
		Command:   "WAIT",
		Args:      []string{strconv.Itoa(t.Replicas), strconv.Itoa(t.TimeoutInMilliseconds)},
		Assertion: resp_assertions.NewIntegerAssertion(t.ExpectedMessage),
	}

	return commandTest.Run(client, logger)
}
