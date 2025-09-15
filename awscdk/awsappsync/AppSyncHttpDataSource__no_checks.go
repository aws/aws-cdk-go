//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func validateAppSyncHttpDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSyncHttpDataSource) validateSetApiParameters(val IApi) error {
	return nil
}

func validateNewAppSyncHttpDataSourceParameters(scope constructs.Construct, id *string, props *AppSyncHttpDataSourceProps) error {
	return nil
}

