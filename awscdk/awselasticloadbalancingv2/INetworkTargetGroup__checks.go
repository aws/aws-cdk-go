//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"
)

func (i *jsiiProxy_INetworkTargetGroup) validateRegisterListenerParameters(listener INetworkListener) error {
	if listener == nil {
		return fmt.Errorf("parameter listener is required, but nil was provided")
	}

	return nil
}

