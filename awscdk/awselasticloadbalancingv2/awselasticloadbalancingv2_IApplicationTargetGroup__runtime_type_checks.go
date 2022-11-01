//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

func (i *jsiiProxy_IApplicationTargetGroup) validateRegisterConnectableParameters(connectable awsec2.IConnectable) error {
	if connectable == nil {
		return fmt.Errorf("parameter connectable is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IApplicationTargetGroup) validateRegisterListenerParameters(listener IApplicationListener) error {
	if listener == nil {
		return fmt.Errorf("parameter listener is required, but nil was provided")
	}

	return nil
}

