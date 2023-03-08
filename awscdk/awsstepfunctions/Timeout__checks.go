//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func validateTimeout_AtParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateTimeout_DurationParameters(duration awscdk.Duration) error {
	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

