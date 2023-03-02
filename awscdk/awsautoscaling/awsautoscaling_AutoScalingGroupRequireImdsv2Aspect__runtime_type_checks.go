//go:build !no_runtime_type_checking

package awsautoscaling

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (a *jsiiProxy_AutoScalingGroupRequireImdsv2Aspect) validateVisitParameters(node awscdk.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoScalingGroupRequireImdsv2Aspect) validateWarnParameters(node awscdk.IConstruct, message *string) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

