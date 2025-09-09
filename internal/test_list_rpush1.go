package internal

import (
	"github.com/make-core/redis-tester/internal/instrumented_resp_connection"
	"github.com/make-core/redis-tester/internal/redis_executable"
	"github.com/make-core/redis-tester/internal/resp_assertions"
	"github.com/make-core/redis-tester/internal/test_cases"
	testerutils_random "github.com/make-core/tester-utils/random"
	"github.com/make-core/tester-utils/test_case_harness"
)

func testListRpush1(stageHarness *test_case_harness.TestCaseHarness) error {
	b := redis_executable.NewRedisExecutable(stageHarness)
	if err := b.Run(); err != nil {
		return err
	}

	logger := stageHarness.Logger
	client, err := instrumented_resp_connection.NewFromAddr(logger, "localhost:6379", "client")
	if err != nil {
		logFriendlyError(logger, err)
		return err
	}
	defer client.Close()

	keyAndValue := testerutils_random.RandomWords(2)

	testcase := test_cases.SendCommandTestCase{
		Command:   "RPUSH",
		Args:      keyAndValue,
		Assertion: resp_assertions.NewIntegerAssertion(1),
	}

	return testcase.Run(client, logger)
}
