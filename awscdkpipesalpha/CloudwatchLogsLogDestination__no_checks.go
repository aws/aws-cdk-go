//go:build no_runtime_type_checking

package awscdkpipesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudwatchLogsLogDestination) validateBindParameters(_pipe IPipe) error {
	return nil
}

func (c *jsiiProxy_CloudwatchLogsLogDestination) validateGrantPushParameters(pipeRole awsiam.IRole) error {
	return nil
}

func validateNewCloudwatchLogsLogDestinationParameters(logGroup awslogs.ILogGroup) error {
	return nil
}

