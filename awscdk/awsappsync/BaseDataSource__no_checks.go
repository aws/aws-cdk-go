//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BaseDataSource) validateCreateFunctionParameters(id *string, props *BaseAppsyncFunctionProps) error {
	return nil
}

func (b *jsiiProxy_BaseDataSource) validateCreateResolverParameters(id *string, props *BaseResolverProps) error {
	return nil
}

func validateBaseDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_BaseDataSource) validateSetApiParameters(val IGraphqlApi) error {
	return nil
}

func validateNewBaseDataSourceParameters(scope constructs.Construct, id *string, props *BackedDataSourceProps, extended *ExtendedDataSourceProps) error {
	return nil
}

