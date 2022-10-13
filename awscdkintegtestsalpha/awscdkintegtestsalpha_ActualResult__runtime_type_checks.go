//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func validateActualResult_FromAwsApiCallParameters(query IApiCall, attribute *string) error {
	if query == nil {
		return fmt.Errorf("parameter query is required, but nil was provided")
	}

	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func validateActualResult_FromCustomResourceParameters(customResource awscdk.CustomResource, attribute *string) error {
	if customResource == nil {
		return fmt.Errorf("parameter customResource is required, but nil was provided")
	}

	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ActualResult) validateSetResultParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

