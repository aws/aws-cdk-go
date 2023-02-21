//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GelfLogDriver) validateBindParameters(_scope constructs.Construct, _containerDefinition ContainerDefinition) error {
	return nil
}

func validateGelfLogDriver_AwsLogsParameters(props *AwsLogDriverProps) error {
	return nil
}

func validateNewGelfLogDriverParameters(props *GelfLogDriverProps) error {
	return nil
}

