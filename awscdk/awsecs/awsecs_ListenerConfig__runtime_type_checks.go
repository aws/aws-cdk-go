//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

func (l *jsiiProxy_ListenerConfig) validateAddTargetsParameters(id *string, target *LoadBalancerTargetOptions, service BaseService) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(target, func() string { return "parameter target" }); err != nil {
		return err
	}

	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

func validateListenerConfig_ApplicationListenerParameters(listener awselasticloadbalancingv2.ApplicationListener, props *awselasticloadbalancingv2.AddApplicationTargetsProps) error {
	if listener == nil {
		return fmt.Errorf("parameter listener is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateListenerConfig_NetworkListenerParameters(listener awselasticloadbalancingv2.NetworkListener, props *awselasticloadbalancingv2.AddNetworkTargetsProps) error {
	if listener == nil {
		return fmt.Errorf("parameter listener is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

