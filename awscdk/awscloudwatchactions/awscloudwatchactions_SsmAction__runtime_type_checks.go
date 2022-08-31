//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awscloudwatchactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
)

func (s *jsiiProxy_SsmAction) validateBindParameters(_scope awscdk.Construct, _alarm awscloudwatch.IAlarm) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _alarm == nil {
		return fmt.Errorf("parameter _alarm is required, but nil was provided")
	}

	return nil
}

func validateNewSsmActionParameters(severity OpsItemSeverity) error {
	if severity == "" {
		return fmt.Errorf("parameter severity is required, but nil was provided")
	}

	return nil
}

