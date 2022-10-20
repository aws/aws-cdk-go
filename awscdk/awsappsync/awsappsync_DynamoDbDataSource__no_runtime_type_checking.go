//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamoDbDataSource) validateCreateFunctionParameters(props *BaseAppsyncFunctionProps) error {
	return nil
}

func (d *jsiiProxy_DynamoDbDataSource) validateCreateResolverParameters(props *BaseResolverProps) error {
	return nil
}

func (d *jsiiProxy_DynamoDbDataSource) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_DynamoDbDataSource) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateDynamoDbDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamoDbDataSource) validateSetApiParameters(val IGraphqlApi) error {
	return nil
}

func validateNewDynamoDbDataSourceParameters(scope constructs.Construct, id *string, props *DynamoDbDataSourceProps) error {
	return nil
}

