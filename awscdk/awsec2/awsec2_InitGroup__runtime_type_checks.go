//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"
)

func validateInitGroup_FromNameParameters(groupName *string) error {
	if groupName == nil {
		return fmt.Errorf("parameter groupName is required, but nil was provided")
	}

	return nil
}

func validateNewInitGroupParameters(groupName *string) error {
	if groupName == nil {
		return fmt.Errorf("parameter groupName is required, but nil was provided")
	}

	return nil
}

