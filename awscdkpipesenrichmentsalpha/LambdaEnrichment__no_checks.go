//go:build no_runtime_type_checking

package awscdkpipesenrichmentsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaEnrichment) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (l *jsiiProxy_LambdaEnrichment) validateGrantInvokeParameters(pipeRole awsiam.IRole) error {
	return nil
}

func validateNewLambdaEnrichmentParameters(lambda awslambda.IFunction, props *LambdaEnrichmentProps) error {
	return nil
}

