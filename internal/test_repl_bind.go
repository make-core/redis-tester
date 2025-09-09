package internal

import (
	"strconv"

	"github.com/make-core/redis-tester/internal/redis_executable"
	"github.com/make-core/redis-tester/internal/test_cases"
	testerutils_random "github.com/make-core/tester-utils/random"
	"github.com/make-core/tester-utils/test_case_harness"
)

func testReplBindToCustomPort(stageHarness *test_case_harness.TestCaseHarness) error {
	port := testerutils_random.RandomInt(6380, 6390)

	b := redis_executable.NewRedisExecutable(stageHarness)
	if err := b.Run("--port", strconv.Itoa(port)); err != nil {
		return err
	}

	logger := stageHarness.Logger

	bindTestCase := test_cases.BindTestCase{
		Port:    port,
		Retries: 15,
	}

	return bindTestCase.Run(b, logger)
}
