//go:build !no_runtime_type_checking

package awsscheduler

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func validateTimeWindow_FlexibleParameters(maxWindow awscdk.Duration) error {
	if maxWindow == nil {
		return fmt.Errorf("parameter maxWindow is required, but nil was provided")
	}

	return nil
}

