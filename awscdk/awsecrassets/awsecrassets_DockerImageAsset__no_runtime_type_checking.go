//go:build no_runtime_type_checking

package awsecrassets

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DockerImageAsset) validateAddResourceMetadataParameters(resource awscdk.CfnResource, resourceProperty *string) error {
	return nil
}

func validateDockerImageAsset_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DockerImageAsset) validateSetImageUriParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DockerImageAsset) validateSetRepositoryParameters(val awsecr.IRepository) error {
	return nil
}

func validateNewDockerImageAssetParameters(scope constructs.Construct, id *string, props *DockerImageAssetProps) error {
	return nil
}

