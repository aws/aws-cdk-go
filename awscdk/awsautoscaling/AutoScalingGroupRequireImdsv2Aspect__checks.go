//go:build !no_runtime_type_checking

package awsautoscaling

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AutoScalingGroupRequireImdsv2Aspect) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoScalingGroupRequireImdsv2Aspect) validateWarnParameters(node constructs.IConstruct, message *string) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

