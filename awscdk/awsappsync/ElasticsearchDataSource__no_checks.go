//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ElasticsearchDataSource) validateCreateFunctionParameters(props *BaseAppsyncFunctionProps) error {
	return nil
}

func (e *jsiiProxy_ElasticsearchDataSource) validateCreateResolverParameters(props *BaseResolverProps) error {
	return nil
}

func (e *jsiiProxy_ElasticsearchDataSource) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (e *jsiiProxy_ElasticsearchDataSource) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateElasticsearchDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ElasticsearchDataSource) validateSetApiParameters(val IGraphqlApi) error {
	return nil
}

func validateNewElasticsearchDataSourceParameters(scope constructs.Construct, id *string, props *ElasticsearchDataSourceProps) error {
	return nil
}

