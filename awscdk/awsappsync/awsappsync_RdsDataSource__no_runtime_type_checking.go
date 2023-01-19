//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsDataSource) validateCreateFunctionParameters(props *BaseAppsyncFunctionProps) error {
	return nil
}

func (r *jsiiProxy_RdsDataSource) validateCreateResolverParameters(props *BaseResolverProps) error {
	return nil
}

func (r *jsiiProxy_RdsDataSource) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (r *jsiiProxy_RdsDataSource) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateRdsDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_RdsDataSource) validateSetApiParameters(val IGraphqlApi) error {
	return nil
}

func validateNewRdsDataSourceParameters(scope constructs.Construct, id *string, props *RdsDataSourceProps) error {
	return nil
}

