package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// OpenTelemetry logs exporter configurations.
//
// Contains constants for configuring log export behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   logsExporting := applicationsignals_alpha.NewLogsExporting()
//
// Experimental.
type LogsExporting interface {
}

// The jsii proxy struct for LogsExporting
type jsiiProxy_LogsExporting struct {
	_ byte // padding
}

// Experimental.
func NewLogsExporting() LogsExporting {
	_init_.Initialize()

	j := jsiiProxy_LogsExporting{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.LogsExporting",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewLogsExporting_Override(l LogsExporting) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.LogsExporting",
		nil, // no parameters
		l,
	)
}

func LogsExporting_OTEL_LOGS_EXPORTER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.LogsExporting",
		"OTEL_LOGS_EXPORTER",
		&returns,
	)
	return returns
}

func LogsExporting_OTEL_LOGS_EXPORTER_NONE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.LogsExporting",
		"OTEL_LOGS_EXPORTER_NONE",
		&returns,
	)
	return returns
}

func LogsExporting_OTEL_LOGS_EXPORTER_OTLP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.LogsExporting",
		"OTEL_LOGS_EXPORTER_OTLP",
		&returns,
	)
	return returns
}

