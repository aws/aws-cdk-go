//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsappsync

import (
	"fmt"
)

func validatePrimaryKey_PartitionParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateNewPrimaryKeyParameters(pkey Assign) error {
	if pkey == nil {
		return fmt.Errorf("parameter pkey is required, but nil was provided")
	}

	return nil
}

