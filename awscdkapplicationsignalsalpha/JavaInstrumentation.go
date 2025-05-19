package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Java-specific OpenTelemetry instrumentation configurations.
//
// Contains constants for Java agent setup and tool options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   javaInstrumentation := applicationsignals_alpha.NewJavaInstrumentation()
//
// Experimental.
type JavaInstrumentation interface {
}

// The jsii proxy struct for JavaInstrumentation
type jsiiProxy_JavaInstrumentation struct {
	_ byte // padding
}

// Experimental.
func NewJavaInstrumentation() JavaInstrumentation {
	_init_.Initialize()

	j := jsiiProxy_JavaInstrumentation{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInstrumentation",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewJavaInstrumentation_Override(j JavaInstrumentation) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInstrumentation",
		nil, // no parameters
		j,
	)
}

func JavaInstrumentation_JAVA_TOOL_OPTIONS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInstrumentation",
		"JAVA_TOOL_OPTIONS",
		&returns,
	)
	return returns
}

