package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Contains all metrics for a Target Group of a Network Load Balancer.
type INetworkTargetGroupMetrics interface {
	// Return the given named metric for this Network Target Group.
	// Default: Average over 5 minutes.
	//
	Custom(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of targets that are considered healthy.
	// Default: Average over 5 minutes.
	//
	HealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of targets that are considered unhealthy.
	// Default: Average over 5 minutes.
	//
	UnHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
}

// The jsii proxy for INetworkTargetGroupMetrics
type jsiiProxy_INetworkTargetGroupMetrics struct {
	_ byte // padding
}

func (i *jsiiProxy_INetworkTargetGroupMetrics) Custom(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateCustomParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"custom",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INetworkTargetGroupMetrics) HealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateHealthyHostCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"healthyHostCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INetworkTargetGroupMetrics) UnHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateUnHealthyHostCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"unHealthyHostCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

