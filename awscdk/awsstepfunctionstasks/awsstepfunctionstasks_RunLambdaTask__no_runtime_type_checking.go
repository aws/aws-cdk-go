//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RunLambdaTask) validateBindParameters(_task awsstepfunctions.Task) error {
	return nil
}

func validateNewRunLambdaTaskParameters(lambdaFunction awslambda.IFunction, props *RunLambdaTaskProps) error {
	return nil
}

