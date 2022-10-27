//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func validateInitCommandWaitDuration_OfParameters(duration awscdk.Duration) error {
	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

