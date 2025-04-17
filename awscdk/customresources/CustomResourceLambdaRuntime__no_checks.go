//go:build no_runtime_type_checking

package customresources

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomResourceLambdaRuntime) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func validateNewCustomResourceLambdaRuntimeParameters(lambdaRuntime awslambda.Runtime) error {
	return nil
}

