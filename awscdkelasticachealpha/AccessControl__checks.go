//go:build !no_runtime_type_checking

package awscdkelasticachealpha

import (
	"fmt"
)

func validateAccessControl_FromAccessStringParameters(accessString *string) error {
	if accessString == nil {
		return fmt.Errorf("parameter accessString is required, but nil was provided")
	}

	return nil
}

