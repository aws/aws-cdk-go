//go:build !no_runtime_type_checking

package awss3

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func validateObjectLockRetention_ComplianceParameters(duration awscdk.Duration) error {
	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

func validateObjectLockRetention_GovernanceParameters(duration awscdk.Duration) error {
	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

