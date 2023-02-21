//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyslogLogDriver) validateBindParameters(_scope constructs.Construct, _containerDefinition ContainerDefinition) error {
	return nil
}

func validateSyslogLogDriver_AwsLogsParameters(props *AwsLogDriverProps) error {
	return nil
}

func validateNewSyslogLogDriverParameters(props *SyslogLogDriverProps) error {
	return nil
}

