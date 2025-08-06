package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Python-specific OpenTelemetry instrumentation configurations.
//
// Contains constants for Python distribution, configurator, and path settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   pythonInstrumentation := applicationsignals_alpha.NewPythonInstrumentation()
//
// Experimental.
type PythonInstrumentation interface {
}

// The jsii proxy struct for PythonInstrumentation
type jsiiProxy_PythonInstrumentation struct {
	_ byte // padding
}

// Experimental.
func NewPythonInstrumentation() PythonInstrumentation {
	_init_.Initialize()

	j := jsiiProxy_PythonInstrumentation{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInstrumentation",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewPythonInstrumentation_Override(p PythonInstrumentation) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInstrumentation",
		nil, // no parameters
		p,
	)
}

func PythonInstrumentation_OTEL_PYTHON_CONFIGURATOR() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInstrumentation",
		"OTEL_PYTHON_CONFIGURATOR",
		&returns,
	)
	return returns
}

func PythonInstrumentation_OTEL_PYTHON_CONFIGURATOR_AWS_CONFIGURATOR() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInstrumentation",
		"OTEL_PYTHON_CONFIGURATOR_AWS_CONFIGURATOR",
		&returns,
	)
	return returns
}

func PythonInstrumentation_OTEL_PYTHON_DISTRO() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInstrumentation",
		"OTEL_PYTHON_DISTRO",
		&returns,
	)
	return returns
}

func PythonInstrumentation_OTEL_PYTHON_DISTRO_AWS_DISTRO() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInstrumentation",
		"OTEL_PYTHON_DISTRO_AWS_DISTRO",
		&returns,
	)
	return returns
}

func PythonInstrumentation_PYTHONPATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInstrumentation",
		"PYTHONPATH",
		&returns,
	)
	return returns
}

