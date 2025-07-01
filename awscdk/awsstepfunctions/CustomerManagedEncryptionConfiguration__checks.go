//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

func (j *jsiiProxy_CustomerManagedEncryptionConfiguration) validateSetTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCustomerManagedEncryptionConfigurationParameters(kmsKey awskms.IKey) error {
	if kmsKey == nil {
		return fmt.Errorf("parameter kmsKey is required, but nil was provided")
	}

	return nil
}

