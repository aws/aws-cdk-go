//go:build no_runtime_type_checking

package awsekslegacy

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksOptimizedImage) validateGetImageParameters(scope awscdk.Construct) error {
	return nil
}

func validateNewEksOptimizedImageParameters(props *EksOptimizedImageProps) error {
	return nil
}

