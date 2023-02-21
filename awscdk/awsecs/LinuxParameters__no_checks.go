//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxParameters) validateAddDevicesParameters(device *[]*Device) error {
	return nil
}

func (l *jsiiProxy_LinuxParameters) validateAddTmpfsParameters(tmpfs *[]*Tmpfs) error {
	return nil
}

func validateLinuxParameters_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLinuxParametersParameters(scope constructs.Construct, id *string, props *LinuxParametersProps) error {
	return nil
}

