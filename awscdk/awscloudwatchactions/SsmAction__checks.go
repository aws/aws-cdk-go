//go:build !no_runtime_type_checking

package awscloudwatchactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SsmAction) validateBindParameters(scope constructs.Construct, alarm awscloudwatch.IAlarm) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if alarm == nil {
		return fmt.Errorf("parameter alarm is required, but nil was provided")
	}

	return nil
}

func validateNewSsmActionParameters(severity OpsItemSeverity) error {
	if severity == "" {
		return fmt.Errorf("parameter severity is required, but nil was provided")
	}

	return nil
}

