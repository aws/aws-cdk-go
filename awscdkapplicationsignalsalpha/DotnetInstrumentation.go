package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// .NET-specific OpenTelemetry instrumentation configurations. Contains constants for .NET runtime settings, profiler configurations, and paths for both Linux and Windows environments.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   dotnetInstrumentation := applicationsignals_alpha.NewDotnetInstrumentation()
//
// Experimental.
type DotnetInstrumentation interface {
}

// The jsii proxy struct for DotnetInstrumentation
type jsiiProxy_DotnetInstrumentation struct {
	_ byte // padding
}

// Experimental.
func NewDotnetInstrumentation() DotnetInstrumentation {
	_init_.Initialize()

	j := jsiiProxy_DotnetInstrumentation{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDotnetInstrumentation_Override(d DotnetInstrumentation) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		nil, // no parameters
		d,
	)
}

func DotnetInstrumentation_CORECLR_ENABLE_PROFILING() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"CORECLR_ENABLE_PROFILING",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_CORECLR_ENABLE_PROFILING_DISABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"CORECLR_ENABLE_PROFILING_DISABLED",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_CORECLR_ENABLE_PROFILING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"CORECLR_ENABLE_PROFILING_ENABLED",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_CORECLR_PROFILER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"CORECLR_PROFILER",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_CORECLR_PROFILER_OTEL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"CORECLR_PROFILER_OTEL",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_CORECLR_PROFILER_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"CORECLR_PROFILER_PATH",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_DOTNET_ADDITIONAL_DEPS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"DOTNET_ADDITIONAL_DEPS",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_DOTNET_SHARED_STORE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"DOTNET_SHARED_STORE",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_DOTNET_STARTUP_HOOKS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"DOTNET_STARTUP_HOOKS",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_OTEL_DOTNET_AUTO_HOME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"OTEL_DOTNET_AUTO_HOME",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_OTEL_DOTNET_AUTO_PLUGINS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"OTEL_DOTNET_AUTO_PLUGINS",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_OTEL_DOTNET_AUTO_PLUGINS_ADOT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"OTEL_DOTNET_AUTO_PLUGINS_ADOT",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_OTEL_DOTNET_CONFIGURATOR() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"OTEL_DOTNET_CONFIGURATOR",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_OTEL_DOTNET_CONFIGURATOR_AWS_CONFIGURATOR() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"OTEL_DOTNET_CONFIGURATOR_AWS_CONFIGURATOR",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_OTEL_DOTNET_DISTRO() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"OTEL_DOTNET_DISTRO",
		&returns,
	)
	return returns
}

func DotnetInstrumentation_OTEL_DOTNET_DISTRO_AWS_DISTRO() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentation",
		"OTEL_DOTNET_DISTRO_AWS_DISTRO",
		&returns,
	)
	return returns
}

