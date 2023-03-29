//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogDriver) validateBindParameters(scope constructs.Construct, containerDefinition ContainerDefinition) error {
	return nil
}

func validateLogDriver_AwsLogsParameters(props *AwsLogDriverProps) error {
	return nil
}

