//go:build !no_runtime_type_checking

package awsevents

import (
	"fmt"
)

func (i *jsiiProxy_IRuleTarget) validateBindParameters(rule IRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

