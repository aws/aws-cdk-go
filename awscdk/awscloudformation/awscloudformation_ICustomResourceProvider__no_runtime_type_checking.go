//go:build no_runtime_type_checking

package awscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICustomResourceProvider) validateBindParameters(scope awscdk.Construct) error {
	return nil
}

