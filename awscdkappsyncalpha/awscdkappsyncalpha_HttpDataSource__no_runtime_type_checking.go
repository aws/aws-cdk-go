//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpDataSource) validateCreateFunctionParameters(props *BaseAppsyncFunctionProps) error {
	return nil
}

func (h *jsiiProxy_HttpDataSource) validateCreateResolverParameters(props *BaseResolverProps) error {
	return nil
}

func validateHttpDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_HttpDataSource) validateSetApiParameters(val IGraphqlApi) error {
	return nil
}

func validateNewHttpDataSourceParameters(scope constructs.Construct, id *string, props *HttpDataSourceProps) error {
	return nil
}

