package awscdksagemakeralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Represents an instance production variant that has been associated with an endpoint.
// Experimental.
type IEndpointInstanceProductionVariant interface {
	// Enable autoscaling for SageMaker Endpoint production variant.
	// Experimental.
	AutoScaleInstanceCount(scalingProps *awsapplicationautoscaling.EnableScalingProps) ScalableInstanceCount
	// Return the given named metric for Endpoint.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	Metric(namespace *string, metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for CPU utilization.
	// Default: - average over 5 minutes.
	//
	// Experimental.
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for disk utilization.
	// Default: - average over 5 minutes.
	//
	// Experimental.
	MetricDiskUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for GPU memory utilization.
	// Default: - average over 5 minutes.
	//
	// Experimental.
	MetricGpuMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for GPU utilization.
	// Default: - average over 5 minutes.
	//
	// Experimental.
	MetricGpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of invocations by HTTP response code.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricInvocationResponseCode(responseCode InvocationHttpResponseCode, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of invocations.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of invocations per instance.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricInvocationsPerInstance(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for memory utilization.
	// Default: - average over 5 minutes.
	//
	// Experimental.
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for model latency.
	// Default: - average over 5 minutes.
	//
	// Experimental.
	MetricModelLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for overhead latency.
	// Default: - average over 5 minutes.
	//
	// Experimental.
	MetricOverheadLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The name of the production variant.
	// Experimental.
	VariantName() *string
}

// The jsii proxy for IEndpointInstanceProductionVariant
type jsiiProxy_IEndpointInstanceProductionVariant struct {
	_ byte // padding
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) AutoScaleInstanceCount(scalingProps *awsapplicationautoscaling.EnableScalingProps) ScalableInstanceCount {
	if err := i.validateAutoScaleInstanceCountParameters(scalingProps); err != nil {
		panic(err)
	}
	var returns ScalableInstanceCount

	_jsii_.Invoke(
		i,
		"autoScaleInstanceCount",
		[]interface{}{scalingProps},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) Metric(namespace *string, metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(namespace, metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{namespace, metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) MetricDiskUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDiskUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDiskUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) MetricGpuMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricGpuMemoryUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricGpuMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) MetricGpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricGpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricGpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) MetricInvocationResponseCode(responseCode InvocationHttpResponseCode, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationResponseCodeParameters(responseCode, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocationResponseCode",
		[]interface{}{responseCode, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) MetricInvocationsPerInstance(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationsPerInstanceParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocationsPerInstance",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricMemoryUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) MetricModelLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricModelLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricModelLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEndpointInstanceProductionVariant) MetricOverheadLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricOverheadLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricOverheadLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IEndpointInstanceProductionVariant) VariantName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantName",
		&returns,
	)
	return returns
}

