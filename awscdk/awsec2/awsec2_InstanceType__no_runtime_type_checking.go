//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateInstanceType_OfParameters(instanceClass InstanceClass, instanceSize InstanceSize) error {
	return nil
}

func validateNewInstanceTypeParameters(instanceTypeIdentifier *string) error {
	return nil
}

