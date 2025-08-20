//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_ApplicationTargetGroup) validateAddLoadBalancerTargetParameters(props *LoadBalancerTargetProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateConfigureHealthCheckParameters(healthCheck *HealthCheck) error {
	if healthCheck == nil {
		return fmt.Errorf("parameter healthCheck is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(healthCheck, func() string { return "parameter healthCheck" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateEnableCookieStickinessParameters(duration awscdk.Duration) error {
	if duration == nil {
		return fmt.Errorf("parameter duration is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricHealthyHostCountParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricHttpCodeTargetParameters(code HttpCodeTarget, props *awscloudwatch.MetricOptions) error {
	if code == "" {
		return fmt.Errorf("parameter code is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricIpv6RequestCountParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricRequestCountParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricRequestCountPerTargetParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricTargetConnectionErrorCountParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricTargetResponseTimeParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricTargetTLSNegotiationErrorCountParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateMetricUnhealthyHostCountParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateRegisterConnectableParameters(connectable awsec2.IConnectable) error {
	if connectable == nil {
		return fmt.Errorf("parameter connectable is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateRegisterListenerParameters(listener IApplicationListener) error {
	if listener == nil {
		return fmt.Errorf("parameter listener is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationTargetGroup) validateSetAttributeParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateApplicationTargetGroup_FromTargetGroupAttributesParameters(scope constructs.Construct, id *string, attrs *TargetGroupAttributes) error {
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

func validateApplicationTargetGroup_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApplicationTargetGroup) validateSetHealthCheckParameters(val *HealthCheck) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func validateNewApplicationTargetGroupParameters(scope constructs.Construct, id *string, props *ApplicationTargetGroupProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

