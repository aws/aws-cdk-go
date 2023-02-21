//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualGatewayListener) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateVirtualGatewayListener_GrpcParameters(options *GrpcGatewayListenerOptions) error {
	return nil
}

func validateVirtualGatewayListener_HttpParameters(options *HttpGatewayListenerOptions) error {
	return nil
}

func validateVirtualGatewayListener_Http2Parameters(options *Http2GatewayListenerOptions) error {
	return nil
}

