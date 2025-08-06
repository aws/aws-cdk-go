//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func validateAppSyncBackedDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSyncBackedDataSource) validateSetApiParameters(val IApi) error {
	return nil
}

func validateNewAppSyncBackedDataSourceParameters(scope constructs.Construct, id *string, props *AppSyncBackedDataSourceProps, extended *AppSyncExtendedDataSourceProps) error {
	return nil
}

