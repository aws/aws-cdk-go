//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awslogs"
)

func (l *jsiiProxy_LogGroupLogDestination) validateBindParameters(_stage IStage) error {
	if _stage == nil {
		return fmt.Errorf("parameter _stage is required, but nil was provided")
	}

	return nil
}

func validateNewLogGroupLogDestinationParameters(logGroup awslogs.ILogGroup) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	return nil
}

