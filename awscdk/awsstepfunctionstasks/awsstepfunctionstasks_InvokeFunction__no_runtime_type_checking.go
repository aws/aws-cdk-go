//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InvokeFunction) validateBindParameters(_task awsstepfunctions.Task) error {
	return nil
}

func validateNewInvokeFunctionParameters(lambdaFunction awslambda.IFunction, props *InvokeFunctionProps) error {
	return nil
}

