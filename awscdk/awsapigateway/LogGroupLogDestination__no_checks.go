//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogGroupLogDestination) validateBindParameters(_stage IStageRef) error {
	return nil
}

func validateNewLogGroupLogDestinationParameters(logGroup awslogs.ILogGroup) error {
	return nil
}

