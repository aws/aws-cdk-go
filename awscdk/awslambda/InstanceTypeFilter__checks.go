//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func validateInstanceTypeFilter_AllowParameters(instanceTypes *[]awsec2.InstanceType) error {
	if instanceTypes == nil {
		return fmt.Errorf("parameter instanceTypes is required, but nil was provided")
	}

	return nil
}

func validateInstanceTypeFilter_ExcludeParameters(instanceTypes *[]awsec2.InstanceType) error {
	if instanceTypes == nil {
		return fmt.Errorf("parameter instanceTypes is required, but nil was provided")
	}

	return nil
}

