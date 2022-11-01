//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awscloudwatchactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SsmIncidentAction) validateBindParameters(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _alarm == nil {
		return fmt.Errorf("parameter _alarm is required, but nil was provided")
	}

	return nil
}

func validateNewSsmIncidentActionParameters(responsePlanName *string) error {
	if responsePlanName == nil {
		return fmt.Errorf("parameter responsePlanName is required, but nil was provided")
	}

	return nil
}

