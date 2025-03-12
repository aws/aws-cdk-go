//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (i *jsiiProxy_IAssumeRolePrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	if document == nil {
		return fmt.Errorf("parameter document is required, but nil was provided")
	}

	return nil
}

