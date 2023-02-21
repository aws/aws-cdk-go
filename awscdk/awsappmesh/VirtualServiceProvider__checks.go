//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (v *jsiiProxy_VirtualServiceProvider) validateBindParameters(_construct constructs.Construct) error {
	if _construct == nil {
		return fmt.Errorf("parameter _construct is required, but nil was provided")
	}

	return nil
}

func validateVirtualServiceProvider_NoneParameters(mesh IMesh) error {
	if mesh == nil {
		return fmt.Errorf("parameter mesh is required, but nil was provided")
	}

	return nil
}

func validateVirtualServiceProvider_VirtualNodeParameters(virtualNode IVirtualNode) error {
	if virtualNode == nil {
		return fmt.Errorf("parameter virtualNode is required, but nil was provided")
	}

	return nil
}

func validateVirtualServiceProvider_VirtualRouterParameters(virtualRouter IVirtualRouter) error {
	if virtualRouter == nil {
		return fmt.Errorf("parameter virtualRouter is required, but nil was provided")
	}

	return nil
}

