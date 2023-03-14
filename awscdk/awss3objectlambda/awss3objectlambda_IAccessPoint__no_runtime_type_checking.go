//go:build no_runtime_type_checking

package awss3objectlambda

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAccessPoint) validateVirtualHostedUrlForObjectParameters(options *awss3.VirtualHostedStyleUrlOptions) error {
	return nil
}

