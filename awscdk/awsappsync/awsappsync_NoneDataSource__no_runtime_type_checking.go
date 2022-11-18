//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NoneDataSource) validateCreateFunctionParameters(props *BaseAppsyncFunctionProps) error {
	return nil
}

func (n *jsiiProxy_NoneDataSource) validateCreateResolverParameters(props *BaseResolverProps) error {
	return nil
}

func (n *jsiiProxy_NoneDataSource) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (n *jsiiProxy_NoneDataSource) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateNoneDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_NoneDataSource) validateSetApiParameters(val IGraphqlApi) error {
	return nil
}

func validateNewNoneDataSourceParameters(scope constructs.Construct, id *string, props *NoneDataSourceProps) error {
	return nil
}

