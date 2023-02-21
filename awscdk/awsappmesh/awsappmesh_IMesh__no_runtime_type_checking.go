//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IMesh) validateAddVirtualGatewayParameters(id *string, props *VirtualGatewayBaseProps) error {
	return nil
}

func (i *jsiiProxy_IMesh) validateAddVirtualNodeParameters(id *string, props *VirtualNodeBaseProps) error {
	return nil
}

func (i *jsiiProxy_IMesh) validateAddVirtualRouterParameters(id *string, props *VirtualRouterBaseProps) error {
	return nil
}

