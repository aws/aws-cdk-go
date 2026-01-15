//go:build no_runtime_type_checking

package awsapigatewayv2

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogGroupLogDestination) validateBindParameters(stage IStage) error {
	return nil
}

func validateNewLogGroupLogDestinationParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

