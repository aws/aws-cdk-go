//go:build no_runtime_type_checking

package awscloudfrontorigins

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RestApiOrigin) validateBindParameters(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) error {
	return nil
}

func validateNewRestApiOriginParameters(restApi awsapigateway.RestApi, props *RestApiOriginProps) error {
	return nil
}

