//go:build !no_runtime_type_checking

package awscloudwatchactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SsmIncidentAction) validateBindParameters(scope constructs.Construct, alarm interfacesawscloudwatch.IAlarmRef) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if alarm == nil {
		return fmt.Errorf("parameter alarm is required, but nil was provided")
	}

	return nil
}

func validateNewSsmIncidentActionParameters(responsePlanName *string) error {
	if responsePlanName == nil {
		return fmt.Errorf("parameter responsePlanName is required, but nil was provided")
	}

	return nil
}

