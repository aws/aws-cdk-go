//go:build no_runtime_type_checking

package awss3assets

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Asset) validateAddResourceMetadataParameters(resource awscdk.CfnResource, resourceProperty *string) error {
	return nil
}

func (a *jsiiProxy_Asset) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateAsset_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAssetParameters(scope constructs.Construct, id *string, props *AssetProps) error {
	return nil
}

