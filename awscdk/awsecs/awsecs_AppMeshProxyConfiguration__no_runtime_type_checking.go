//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppMeshProxyConfiguration) validateBindParameters(_scope awscdk.Construct, _taskDefinition TaskDefinition) error {
	return nil
}

func validateNewAppMeshProxyConfigurationParameters(props *AppMeshProxyConfigurationConfigProps) error {
	return nil
}

