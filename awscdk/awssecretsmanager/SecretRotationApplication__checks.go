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

func validateNewSecretRotationApplicationParameters(applicationId *string, semanticVersion *string, options *SecretRotationApplicationOptions) error {
	if applicationId == nil {
		return fmt.Errorf("parameter applicationId is required, but nil was provided")
	}

	if semanticVersion == nil {
		return fmt.Errorf("parameter semanticVersion is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

