//go:build !no_runtime_type_checking

package awssecretsmanager

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_SecretRotationApplication) validateApplicationArnForPartitionParameters(partition *string) error {
	if partition == nil {
		return fmt.Errorf("parameter partition is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecretRotationApplication) validateSemanticVersionForPartitionParameters(partition *string) error {
	if partition == nil {
		return fmt.Errorf("parameter partition is required, but nil was provided")
	}

	return nil
}

func validateNewSecretRotationApplicationParameters(applicationName *string, awsSemanticVersion *string, options *SecretRotationApplicationOptions) error {
	if applicationName == nil {
		return fmt.Errorf("parameter applicationName is required, but nil was provided")
	}

	if awsSemanticVersion == nil {
		return fmt.Errorf("parameter awsSemanticVersion is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

