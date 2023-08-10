//go:build no_runtime_type_checking

package awscdkappconfigalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IApplication) validateAddEnvironmentParameters(id *string, options *EnvironmentOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateAddExistingEnvironmentParameters(environment IEnvironment) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateAddHostedConfigurationParameters(id *string, options *HostedConfigurationOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateAddSourcedConfigurationParameters(id *string, options *SourcedConfigurationOptions) error {
	return nil
}

