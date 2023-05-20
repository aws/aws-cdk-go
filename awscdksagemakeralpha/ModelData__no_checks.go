//go:build no_runtime_type_checking

package awscdksagemakeralpha

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ModelData) validateBindParameters(scope constructs.Construct, model IModel) error {
	return nil
}

func validateModelData_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateModelData_FromBucketParameters(bucket awss3.IBucket, objectKey *string) error {
	return nil
}

