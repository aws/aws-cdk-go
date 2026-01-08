//go:build !no_runtime_type_checking

package awscodestarnotifications

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (i *jsiiProxy_INotificationRule) validateAddTargetParameters(target INotificationRuleTarget) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_INotificationRule) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

