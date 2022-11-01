//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpDataSource) validateCreateFunctionParameters(props *BaseAppsyncFunctionProps) error {
	return nil
}

func (h *jsiiProxy_HttpDataSource) validateCreateResolverParameters(props *BaseResolverProps) error {
	return nil
}

func (h *jsiiProxy_HttpDataSource) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (h *jsiiProxy_HttpDataSource) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
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

