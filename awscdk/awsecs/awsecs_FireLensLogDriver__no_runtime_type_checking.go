//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FireLensLogDriver) validateBindParameters(_scope constructs.Construct, _containerDefinition ContainerDefinition) error {
	return nil
}

func validateFireLensLogDriver_AwsLogsParameters(props *AwsLogDriverProps) error {
	return nil
}

func validateNewFireLensLogDriverParameters(props *FireLensLogDriverProps) error {
	return nil
}

