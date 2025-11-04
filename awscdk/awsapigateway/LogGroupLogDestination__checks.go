//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

func (l *jsiiProxy_LogGroupLogDestination) validateBindParameters(stage IStageRef) error {
	if stage == nil {
		return fmt.Errorf("parameter stage is required, but nil was provided")
	}

	return nil
}

func validateNewLogGroupLogDestinationParameters(logGroup awslogs.ILogGroup) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	return nil
}

