//go:build no_runtime_type_checking

package customresources

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomResourceConfig) validateAddLambdaRuntimeParameters(lambdaRuntime awslambda.Runtime) error {
	return nil
}

func (c *jsiiProxy_CustomResourceConfig) validateAddLogRetentionLifetimeParameters(retention awslogs.RetentionDays) error {
	return nil
}

func (c *jsiiProxy_CustomResourceConfig) validateAddRemovalPolicyParameters(removalPolicy awscdk.RemovalPolicy) error {
	return nil
}

func validateCustomResourceConfig_OfParameters(scope constructs.IConstruct) error {
	return nil
}

