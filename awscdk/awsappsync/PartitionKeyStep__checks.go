//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"
)

func (p *jsiiProxy_PartitionKeyStep) validateIsParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPartitionKeyStepParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

