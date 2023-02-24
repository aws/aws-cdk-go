//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualServiceProvider) validateBindParameters(_construct constructs.Construct) error {
	return nil
}

func validateVirtualServiceProvider_NoneParameters(mesh IMesh) error {
	return nil
}

func validateVirtualServiceProvider_VirtualNodeParameters(virtualNode IVirtualNode) error {
	return nil
}

func validateVirtualServiceProvider_VirtualRouterParameters(virtualRouter IVirtualRouter) error {
	return nil
}

