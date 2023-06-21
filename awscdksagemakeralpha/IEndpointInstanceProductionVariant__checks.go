//go:build !no_runtime_type_checking

package awscdksagemakeralpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

func (i *jsiiProxy_IEndpointInstanceProductionVariant) validateAutoScaleInstanceCountParameters(scalingProps *awsapplicationautoscaling.EnableScalingProps) error {
	if scalingProps == nil {
		return fmt.Errorf("parameter scalingProps is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(scalingProps, func() string { return "parameter scalingProps" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) validateMetricParameters(namespace *string, metricName *string, props *awscloudwatch.MetricOptions) error {
	if namespace == nil {
		return fmt.Errorf("parameter namespace is required, but nil was provided")
	}

	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) validateMetricCpuUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) validateMetricDiskUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) validateMetricGpuMemoryUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) validateMetricGpuUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) validateMetricInvocationResponseCodeParameters(responseCode InvocationHttpResponseCode, props *awscloudwatch.MetricOptions) error {
	if responseCode == "" {
		return fmt.Errorf("parameter responseCode is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) validateMetricInvocationsPerInstanceParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) validateMetricMemoryUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) validateMetricModelLatencyParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) validateMetricOverheadLatencyParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

