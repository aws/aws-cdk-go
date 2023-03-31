//go:build !no_runtime_type_checking

package awsevents

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func validateHttpParameter_FromSecretParameters(value awscdk.SecretValue) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateHttpParameter_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

