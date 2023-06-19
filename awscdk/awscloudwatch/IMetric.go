package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for metrics.
// Experimental.
type IMetric interface {
	// Turn this metric object into an alarm configuration.
	// Deprecated: Use `toMetricConfig()` instead.
	ToAlarmConfig() *MetricAlarmConfig
	// Turn this metric object into a graph configuration.
	// Deprecated: Use `toMetricConfig()` instead.
	ToGraphConfig() *MetricGraphConfig
	// Inspect the details of the metric object.
	// Experimental.
	ToMetricConfig() *MetricConfig
	// Any warnings related to this metric.
	//
	// Should be attached to the consuming construct.
	// Experimental.
	Warnings() *[]*string
}

// The jsii proxy for IMetric
type jsiiProxy_IMetric struct {
	_ byte // padding
}

func (i *jsiiProxy_IMetric) ToAlarmConfig() *MetricAlarmConfig {
	var returns *MetricAlarmConfig

	_jsii_.Invoke(
		i,
		"toAlarmConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMetric) ToGraphConfig() *MetricGraphConfig {
	var returns *MetricGraphConfig

	_jsii_.Invoke(
		i,
		"toGraphConfig",
		nil, // no parameters
		&returns,
	)

	return returns
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

