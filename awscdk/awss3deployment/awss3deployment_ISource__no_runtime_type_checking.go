//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awss3deployment

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ISource) validateBindParameters(scope awscdk.Construct, context *DeploymentSourceContext) error {
	return nil
}

