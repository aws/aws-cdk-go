//go:build !no_runtime_type_checking

package awscodestarnotifications

import (
	"fmt"
)

func (i *jsiiProxy_INotificationRule) validateAddTargetParameters(target INotificationRuleTarget) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

