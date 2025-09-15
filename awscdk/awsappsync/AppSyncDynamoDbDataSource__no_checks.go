//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func validateAppSyncDynamoDbDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSyncDynamoDbDataSource) validateSetApiParameters(val IApi) error {
	return nil
}

func validateNewAppSyncDynamoDbDataSourceParameters(scope constructs.Construct, id *string, props *AppSyncDynamoDbDataSourceProps) error {
	return nil
}

