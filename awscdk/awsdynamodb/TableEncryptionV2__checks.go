//go:build !no_runtime_type_checking

package awsdynamodb

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

func validateTableEncryptionV2_CustomerManagedKeyParameters(tableKey awskms.IKey) error {
	if tableKey == nil {
		return fmt.Errorf("parameter tableKey is required, but nil was provided")
	}

	return nil
}

