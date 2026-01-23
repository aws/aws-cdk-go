//go:build !no_runtime_type_checking

package awscloudwatchactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_Ec2Action) validateBindParameters(scope constructs.Construct, alarm awscloudwatch.IAlarm) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if alarm == nil {
		return fmt.Errorf("parameter alarm is required, but nil was provided")
	}

	return nil
}

func validateNewEc2ActionParameters(instanceAction Ec2InstanceAction) error {
	if instanceAction == "" {
		return fmt.Errorf("parameter instanceAction is required, but nil was provided")
	}

	return nil
}

