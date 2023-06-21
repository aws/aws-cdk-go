//go:build no_runtime_type_checking

package awscdkclilibalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsCdkCli) validateDeployParameters(options *DeployOptions) error {
	return nil
}

func (a *jsiiProxy_AwsCdkCli) validateDestroyParameters(options *DestroyOptions) error {
	return nil
}

func (a *jsiiProxy_AwsCdkCli) validateListParameters(options *ListOptions) error {
	return nil
}

func (a *jsiiProxy_AwsCdkCli) validateSynthParameters(options *SynthOptions) error {
	return nil
}

func validateAwsCdkCli_FromCdkAppDirectoryParameters(props *CdkAppDirectoryProps) error {
	return nil
}

func validateAwsCdkCli_FromCloudAssemblyDirectoryProducerParameters(producer ICloudAssemblyDirectoryProducer) error {
	return nil
}

