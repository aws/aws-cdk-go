package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Node-specific OpenTelemetry instrumentation configurations.
//
// Contains constants for Node.js runtime settings and options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   nodeInstrumentation := applicationsignals_alpha.NewNodeInstrumentation()
//
// Experimental.
type NodeInstrumentation interface {
}

// The jsii proxy struct for NodeInstrumentation
type jsiiProxy_NodeInstrumentation struct {
	_ byte // padding
}

// Experimental.
func NewNodeInstrumentation() NodeInstrumentation {
	_init_.Initialize()

	j := jsiiProxy_NodeInstrumentation{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInstrumentation",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewNodeInstrumentation_Override(n NodeInstrumentation) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInstrumentation",
		nil, // no parameters
		n,
	)
}

func NodeInstrumentation_NODE_OPTIONS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInstrumentation",
		"NODE_OPTIONS",
		&returns,
	)
	return returns
}

