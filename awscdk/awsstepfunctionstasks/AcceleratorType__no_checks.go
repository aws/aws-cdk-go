//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func validateAcceleratorType_OfParameters(acceleratorClass AcceleratorClass, instanceSize awsec2.InstanceSize) error {
	return nil
}

func validateNewAcceleratorTypeParameters(instanceTypeIdentifier *string) error {
	return nil
}

