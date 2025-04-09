//go:build !no_runtime_type_checking

package awscdkpipesalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

func (c *jsiiProxy_CloudwatchLogsLogDestination) validateBindParameters(_pipe IPipe) error {
	if _pipe == nil {
		return fmt.Errorf("parameter _pipe is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudwatchLogsLogDestination) validateGrantPushParameters(pipeRole awsiam.IRole) error {
	if pipeRole == nil {
		return fmt.Errorf("parameter pipeRole is required, but nil was provided")
	}

	return nil
}

func validateNewCloudwatchLogsLogDestinationParameters(logGroup awslogs.ILogGroup) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	return nil
}

