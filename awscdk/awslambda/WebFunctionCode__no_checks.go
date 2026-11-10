//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebFunctionCode) validateBindParameters(scope constructs.Construct, options *WebFunctionCodeBindOptions) error {
	return nil
}

func (w *jsiiProxy_WebFunctionCode) validateBindToResourceParameters(_resource awscdk.CfnResource) error {
	return nil
}

func validateWebFunctionCode_FromAssetToManagedBucketParameters(path *string, options *WebFunctionManagedBucketAssetOptions) error {
	return nil
}

func validateWebFunctionCode_FromAssetToProvidedBucketParameters(bucket awss3.IBucket, path *string, options *WebFunctionProvidedBucketAssetOptions) error {
	return nil
}

func validateWebFunctionCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

