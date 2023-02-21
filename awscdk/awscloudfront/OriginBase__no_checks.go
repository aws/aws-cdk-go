//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OriginBase) validateBindParameters(_scope constructs.Construct, options *OriginBindOptions) error {
	return nil
}

func validateNewOriginBaseParameters(domainName *string, props *OriginProps) error {
	return nil
}

