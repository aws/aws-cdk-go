//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func validateAppSyncBaseDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSyncBaseDataSource) validateSetApiParameters(val IApi) error {
	return nil
}

func validateNewAppSyncBaseDataSourceParameters(scope constructs.Construct, id *string, props *AppSyncBackedDataSourceProps, extended *AppSyncExtendedDataSourceProps) error {
	return nil
}

