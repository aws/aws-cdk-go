package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// OpenTelemetry metrics exporter configurations.
//
// Contains constants for configuring metrics export behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   metricsExporting := applicationsignals_alpha.NewMetricsExporting()
//
// Experimental.
type MetricsExporting interface {
}

// The jsii proxy struct for MetricsExporting
type jsiiProxy_MetricsExporting struct {
	_ byte // padding
}

// Experimental.
func NewMetricsExporting() MetricsExporting {
	_init_.Initialize()

	j := jsiiProxy_MetricsExporting{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.MetricsExporting",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMetricsExporting_Override(m MetricsExporting) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.MetricsExporting",
		nil, // no parameters
		m,
	)
}

func MetricsExporting_OTEL_METRICS_EXPORTER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.MetricsExporting",
		"OTEL_METRICS_EXPORTER",
		&returns,
	)
	return returns
}

func MetricsExporting_OTEL_METRICS_EXPORTER_NONE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.MetricsExporting",
		"OTEL_METRICS_EXPORTER_NONE",
		&returns,
	)
	return returns
}

func MetricsExporting_OTEL_METRICS_EXPORTER_OTLP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.MetricsExporting",
		"OTEL_METRICS_EXPORTER_OTLP",
		&returns,
	)
	return returns
}

