package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for metrics.
type IMetric interface {
	// Inspect the details of the metric object.
	ToMetricConfig() *MetricConfig
	// Any warnings related to this metric.
	//
	// Should be attached to the consuming construct.
	Warnings() *[]*string
}

// The jsii proxy for IMetric
type jsiiProxy_IMetric struct {
	_ byte // padding
}

func (i *jsiiProxy_IMetric) ToMetricConfig() *MetricConfig {
	var returns *MetricConfig

	_jsii_.Invoke(
		i,
		"toMetricConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IMetric) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

