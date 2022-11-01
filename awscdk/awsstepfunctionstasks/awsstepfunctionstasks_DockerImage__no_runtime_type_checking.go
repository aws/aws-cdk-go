//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DockerImage) validateBindParameters(task ISageMakerTask) error {
	return nil
}

func validateDockerImage_FromAssetParameters(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps) error {
	return nil
}

func validateDockerImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

func validateDockerImage_FromJsonExpressionParameters(expression *string) error {
	return nil
}

func validateDockerImage_FromRegistryParameters(imageUri *string) error {
	return nil
}

