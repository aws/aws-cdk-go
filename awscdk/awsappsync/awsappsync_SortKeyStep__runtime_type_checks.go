//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsappsync

import (
	"fmt"
)

func (s *jsiiProxy_SortKeyStep) validateIsParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewSortKeyStepParameters(pkey Assign, skey *string) error {
	if pkey == nil {
		return fmt.Errorf("parameter pkey is required, but nil was provided")
	}

	if skey == nil {
		return fmt.Errorf("parameter skey is required, but nil was provided")
	}

	return nil
}

