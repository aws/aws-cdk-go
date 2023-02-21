//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LookupMachineImage) validateGetImageParameters(scope constructs.Construct) error {
	return nil
}

func validateNewLookupMachineImageParameters(props *LookupMachineImageProps) error {
	return nil
}

