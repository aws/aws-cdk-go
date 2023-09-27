//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OptimalInstanceType) validateSameInstanceClassAsParameters(other awsec2.InstanceType) error {
	return nil
}

func validateOptimalInstanceType_OfParameters(instanceClass awsec2.InstanceClass, instanceSize awsec2.InstanceSize) error {
	return nil
}

