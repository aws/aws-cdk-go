//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Runtime) validateRuntimeEqualsParameters(other Runtime) error {
	return nil
}

func validateNewRuntimeParameters(name *string, props *LambdaRuntimeProps) error {
	return nil
}

