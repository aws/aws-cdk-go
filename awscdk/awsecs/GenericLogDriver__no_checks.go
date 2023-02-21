//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GenericLogDriver) validateBindParameters(_scope constructs.Construct, _containerDefinition ContainerDefinition) error {
	return nil
}

func validateGenericLogDriver_AwsLogsParameters(props *AwsLogDriverProps) error {
	return nil
}

func validateNewGenericLogDriverParameters(props *GenericLogDriverProps) error {
	return nil
}

