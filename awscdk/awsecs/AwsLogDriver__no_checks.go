//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsLogDriver) validateBindParameters(scope constructs.Construct, containerDefinition ContainerDefinition) error {
	return nil
}

func validateAwsLogDriver_AwsLogsParameters(props *AwsLogDriverProps) error {
	return nil
}

func validateNewAwsLogDriverParameters(props *AwsLogDriverProps) error {
	return nil
}

