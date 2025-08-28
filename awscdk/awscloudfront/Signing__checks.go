//go:build !no_runtime_type_checking

package awscloudfront

import (
	"fmt"
)

func validateNewSigningParameters(protocol SigningProtocol, behavior SigningBehavior) error {
	if protocol == "" {
		return fmt.Errorf("parameter protocol is required, but nil was provided")
	}

	if behavior == "" {
		return fmt.Errorf("parameter behavior is required, but nil was provided")
	}

	return nil
}

