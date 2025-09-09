package internal

import (
	"github.com/make-core/redis-tester/internal/redis_executable"
	"github.com/make-core/redis-tester/internal/test_cases"
	"github.com/make-core/tester-utils/test_case_harness"
)

func testBindToPort(stageHarness *test_case_harness.TestCaseHarness) error {
	b := redis_executable.NewRedisExecutable(stageHarness)
	if err := b.Run(); err != nil {
		return err
	}

	logger := stageHarness.Logger

	bindTestCase := test_cases.BindTestCase{
		Port:    6379,
		Retries: 15,
	}

	return bindTestCase.Run(b, logger)
}
