//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentLifecycleLambdaTarget) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewDeploymentLifecycleLambdaTargetParameters(handler awslambda.IFunction, id *string, props *DeploymentLifecycleLambdaTargetProps) error {
	return nil
}

