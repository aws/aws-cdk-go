//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func (i *jsiiProxy_IResource) validateApplyRemovalPolicyParameters(policy RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

