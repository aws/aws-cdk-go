//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (t *jsiiProxy_TargetGroupBase) validateAddLoadBalancerTargetParameters(props *LoadBalancerTargetProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TargetGroupBase) validateConfigureHealthCheckParameters(healthCheck *HealthCheck) error {
	if healthCheck == nil {
		return fmt.Errorf("parameter healthCheck is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(healthCheck, func() string { return "parameter healthCheck" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TargetGroupBase) validateSetAttributeParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateTargetGroupBase_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TargetGroupBase) validateSetHealthCheckParameters(val *HealthCheck) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func validateNewTargetGroupBaseParameters(scope constructs.Construct, id *string, baseProps *BaseTargetGroupProps, additionalProps interface{}) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if baseProps == nil {
		return fmt.Errorf("parameter baseProps is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(baseProps, func() string { return "parameter baseProps" }); err != nil {
		return err
	}

	if additionalProps == nil {
		return fmt.Errorf("parameter additionalProps is required, but nil was provided")
	}

	return nil
}

