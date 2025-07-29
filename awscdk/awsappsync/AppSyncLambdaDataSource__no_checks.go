//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func validateAppSyncLambdaDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSyncLambdaDataSource) validateSetApiParameters(val IApi) error {
	return nil
}

func validateNewAppSyncLambdaDataSourceParameters(scope constructs.Construct, id *string, props *AppSyncLambdaDataSourceProps) error {
	return nil
}

