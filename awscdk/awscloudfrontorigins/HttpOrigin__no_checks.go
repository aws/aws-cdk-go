//go:build no_runtime_type_checking

package awscloudfrontorigins

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpOrigin) validateBindParameters(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) error {
	return nil
}

func validateNewHttpOriginParameters(domainName *string, props *HttpOriginProps) error {
	return nil
}

