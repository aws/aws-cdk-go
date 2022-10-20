//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_Backend) validateBindParameters(_scope awscdk.Construct) error {
	return nil
}

func validateBackend_VirtualServiceParameters(virtualService IVirtualService, props *VirtualServiceBackendOptions) error {
	return nil
}

