//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualNodeListener) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateVirtualNodeListener_GrpcParameters(props *GrpcVirtualNodeListenerOptions) error {
	return nil
}

func validateVirtualNodeListener_HttpParameters(props *HttpVirtualNodeListenerOptions) error {
	return nil
}

func validateVirtualNodeListener_Http2Parameters(props *Http2VirtualNodeListenerOptions) error {
	return nil
}

func validateVirtualNodeListener_TcpParameters(props *TcpVirtualNodeListenerOptions) error {
	return nil
}

