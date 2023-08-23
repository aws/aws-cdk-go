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
	// Default: - None.
	//
	// Deprecated: - use warningsV2.
	Warnings() *[]*string
	// Any warnings related to this metric.
	//
	// Should be attached to the consuming construct.
	// Default: - None.
	//
	WarningsV2() *map[string]*string
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

func (j *jsiiProxy_IMetric) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

