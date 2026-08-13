//go:build no_runtime_type_checking

package awskinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpEndpoint) validateBindParameters(scope constructs.Construct, options *DestinationBindOptions) error {
	return nil
}

func validateNewHttpEndpointParameters(props *HttpEndpointProps) error {
	return nil
}

