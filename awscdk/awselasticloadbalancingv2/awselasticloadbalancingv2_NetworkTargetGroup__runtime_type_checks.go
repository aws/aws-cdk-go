//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

func (n *jsiiProxy_NetworkTargetGroup) validateAddLoadBalancerTargetParameters(props *LoadBalancerTargetProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NetworkTargetGroup) validateConfigureHealthCheckParameters(healthCheck *HealthCheck) error {
	if healthCheck == nil {
		return fmt.Errorf("parameter healthCheck is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(healthCheck, func() string { return "parameter healthCheck" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NetworkTargetGroup) validateMetricHealthyHostCountParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NetworkTargetGroup) validateMetricUnHealthyHostCountParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NetworkTargetGroup) validateRegisterListenerParameters(listener INetworkListener) error {
	if listener == nil {
		return fmt.Errorf("parameter listener is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkTargetGroup) validateSetAttributeParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateNetworkTargetGroup_FromTargetGroupAttributesParameters(scope constructs.Construct, id *string, attrs *TargetGroupAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validateNetworkTargetGroup_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkTargetGroup) validateSetHealthCheckParameters(val *HealthCheck) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func validateNewNetworkTargetGroupParameters(scope constructs.Construct, id *string, props *NetworkTargetGroupProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

