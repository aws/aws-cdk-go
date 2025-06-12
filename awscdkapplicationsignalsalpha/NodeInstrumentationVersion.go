package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Available versions for Node.js instrumentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   nodeInstrumentationVersion := applicationsignals_alpha.NewNodeInstrumentationVersion(jsii.String("imageRepo"), jsii.String("version"), jsii.Number(123))
//
// Experimental.
type NodeInstrumentationVersion interface {
	InstrumentationVersion
	// Experimental.
	ImageRepo() *string
	// Experimental.
	MemoryLimit() *float64
	// Experimental.
	Version() *string
	// Get the image URI for the instrumentation version.
	//
	// Returns: The image URI.
	// Experimental.
	ImageURI() *string
	// Get the memory limit in MiB for the instrumentation version.
	//
	// Returns: The memory limit.
	// Experimental.
	MemoryLimitMiB() *float64
}

// The jsii proxy struct for NodeInstrumentationVersion
type jsiiProxy_NodeInstrumentationVersion struct {
	jsiiProxy_InstrumentationVersion
}

func (j *jsiiProxy_NodeInstrumentationVersion) ImageRepo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeInstrumentationVersion) MemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeInstrumentationVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewNodeInstrumentationVersion(imageRepo *string, version *string, memoryLimit *float64) NodeInstrumentationVersion {
	_init_.Initialize()

	if err := validateNewNodeInstrumentationVersionParameters(imageRepo, version, memoryLimit); err != nil {
		panic(err)
	}
	j := jsiiProxy_NodeInstrumentationVersion{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInstrumentationVersion",
		[]interface{}{imageRepo, version, memoryLimit},
		&j,
	)

	return &j
}

// Experimental.
func NewNodeInstrumentationVersion_Override(n NodeInstrumentationVersion, imageRepo *string, version *string, memoryLimit *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInstrumentationVersion",
		[]interface{}{imageRepo, version, memoryLimit},
		n,
	)
}

func NodeInstrumentationVersion_DEFAULT_MEMORY_LIMIT_MIB() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInstrumentationVersion",
		"DEFAULT_MEMORY_LIMIT_MIB",
		&returns,
	)
	return returns
}

func NodeInstrumentationVersion_IMAGE_REPO() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInstrumentationVersion",
		"IMAGE_REPO",
		&returns,
	)
	return returns
}

func NodeInstrumentationVersion_V0_5_0() NodeInstrumentationVersion {
	_init_.Initialize()
	var returns NodeInstrumentationVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInstrumentationVersion",
		"V0_5_0",
		&returns,
	)
	return returns
}

func NodeInstrumentationVersion_V0_6_0() NodeInstrumentationVersion {
	_init_.Initialize()
	var returns NodeInstrumentationVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInstrumentationVersion",
		"V0_6_0",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NodeInstrumentationVersion) ImageURI() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"imageURI",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeInstrumentationVersion) MemoryLimitMiB() *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"memoryLimitMiB",
		nil, // no parameters
		&returns,
	)

	return returns
}

