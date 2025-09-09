package internal

import (
	"github.com/make-core/redis-tester/internal/instrumented_resp_connection"
	"github.com/make-core/redis-tester/internal/redis_executable"
	"github.com/make-core/redis-tester/internal/resp_assertions"
	"github.com/make-core/redis-tester/internal/test_cases"
	"github.com/make-core/tester-utils/random"
	"github.com/make-core/tester-utils/test_case_harness"
)

// Tests 'ECHO'
func testEcho(stageHarness *test_case_harness.TestCaseHarness) error {
	b := redis_executable.NewRedisExecutable(stageHarness)
	if err := b.Run(); err != nil {
		return err
	}

	logger := stageHarness.Logger

	client, err := instrumented_resp_connection.NewFromAddr(logger, "localhost:6379", "client")
	if err != nil {
		return err
	}

	randomWord := random.RandomWord()

	commandTestCase := test_cases.SendCommandTestCase{
		Command:   "echo",
		Args:      []string{randomWord},
		Assertion: resp_assertions.NewStringAssertion(randomWord),
	}

	if err := commandTestCase.Run(client, logger); err != nil {
		logFriendlyError(logger, err)
		return err
	}

	client.Close()

	return nil
}
