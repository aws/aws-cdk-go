//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"
)

func validateValues_AttributeParameters(attr *string) error {
	if attr == nil {
		return fmt.Errorf("parameter attr is required, but nil was provided")
	}

	return nil
}

