//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsOptimizedAmi) validateGetImageParameters(scope awscdk.Construct) error {
	return nil
}

func validateNewEcsOptimizedAmiParameters(props *EcsOptimizedAmiProps) error {
	return nil
}

