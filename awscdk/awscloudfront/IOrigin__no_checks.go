//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IOrigin) validateBindParameters(scope constructs.Construct, options *OriginBindOptions) error {
	return nil
}

