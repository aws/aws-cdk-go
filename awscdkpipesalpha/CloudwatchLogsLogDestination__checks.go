//go:build !no_runtime_type_checking

package awscdkpipesalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

func (c *jsiiProxy_CloudwatchLogsLogDestination) validateBindParameters(pipe IPipe) error {
	if pipe == nil {
		return fmt.Errorf("parameter pipe is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudwatchLogsLogDestination) validateGrantPushParameters(grantee awsiam.IRole) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateNewCloudwatchLogsLogDestinationParameters(logGroup awslogs.ILogGroup) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	return nil
}

