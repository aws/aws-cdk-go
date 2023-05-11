//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JournaldLogDriver) validateBindParameters(_scope constructs.Construct, _containerDefinition ContainerDefinition) error {
	return nil
}

func validateJournaldLogDriver_AwsLogsParameters(props *AwsLogDriverProps) error {
	return nil
}

func validateNewJournaldLogDriverParameters(props *JournaldLogDriverProps) error {
	return nil
}

