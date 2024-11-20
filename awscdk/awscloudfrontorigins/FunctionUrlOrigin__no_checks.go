//go:build no_runtime_type_checking

package awscloudfrontorigins

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionUrlOrigin) validateBindParameters(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) error {
	return nil
}

func validateFunctionUrlOrigin_WithOriginAccessControlParameters(lambdaFunctionUrl awslambda.IFunctionUrl, props *FunctionUrlOriginWithOACProps) error {
	return nil
}

func validateNewFunctionUrlOriginParameters(lambdaFunctionUrl awslambda.IFunctionUrl, props *FunctionUrlOriginProps) error {
	return nil
}

