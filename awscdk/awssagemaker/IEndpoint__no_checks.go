//go:build no_runtime_type_checking

package awssagemaker

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IEndpoint) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

