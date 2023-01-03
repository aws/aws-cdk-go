//go:build no_runtime_type_checking

// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerImage) validateBindParameters(scope constructs.Construct, model Model) error {
	return nil
}

func validateContainerImage_FromAssetParameters(directory *string, options *awsecrassets.DockerImageAssetOptions) error {
	return nil
}

func validateContainerImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

