package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Available versions for .NET instrumentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   dotnetInstrumentationVersion := applicationsignals_alpha.NewDotnetInstrumentationVersion(jsii.String("imageRepo"), jsii.String("version"), jsii.Number(123))
//
// Experimental.
type DotnetInstrumentationVersion interface {
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

// The jsii proxy struct for DotnetInstrumentationVersion
type jsiiProxy_DotnetInstrumentationVersion struct {
	jsiiProxy_InstrumentationVersion
}

func (j *jsiiProxy_DotnetInstrumentationVersion) ImageRepo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DotnetInstrumentationVersion) MemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DotnetInstrumentationVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewDotnetInstrumentationVersion(imageRepo *string, version *string, memoryLimit *float64) DotnetInstrumentationVersion {
	_init_.Initialize()

	if err := validateNewDotnetInstrumentationVersionParameters(imageRepo, version, memoryLimit); err != nil {
		panic(err)
	}
	j := jsiiProxy_DotnetInstrumentationVersion{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentationVersion",
		[]interface{}{imageRepo, version, memoryLimit},
		&j,
	)

	return &j
}

// Experimental.
func NewDotnetInstrumentationVersion_Override(d DotnetInstrumentationVersion, imageRepo *string, version *string, memoryLimit *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentationVersion",
		[]interface{}{imageRepo, version, memoryLimit},
		d,
	)
}

func DotnetInstrumentationVersion_DEFAULT_MEMORY_LIMIT_MIB() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentationVersion",
		"DEFAULT_MEMORY_LIMIT_MIB",
		&returns,
	)
	return returns
}

func DotnetInstrumentationVersion_IMAGE_REPO() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentationVersion",
		"IMAGE_REPO",
		&returns,
	)
	return returns
}

func DotnetInstrumentationVersion_V1_6_0() DotnetInstrumentationVersion {
	_init_.Initialize()
	var returns DotnetInstrumentationVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentationVersion",
		"V1_6_0",
		&returns,
	)
	return returns
}

func DotnetInstrumentationVersion_V1_6_0_WINDOWS2019() DotnetInstrumentationVersion {
	_init_.Initialize()
	var returns DotnetInstrumentationVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentationVersion",
		"V1_6_0_WINDOWS2019",
		&returns,
	)
	return returns
}

func DotnetInstrumentationVersion_V1_6_0_WINDOWS2022() DotnetInstrumentationVersion {
	_init_.Initialize()
	var returns DotnetInstrumentationVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentationVersion",
		"V1_6_0_WINDOWS2022",
		&returns,
	)
	return returns
}

func DotnetInstrumentationVersion_V1_7_0() DotnetInstrumentationVersion {
	_init_.Initialize()
	var returns DotnetInstrumentationVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentationVersion",
		"V1_7_0",
		&returns,
	)
	return returns
}

func DotnetInstrumentationVersion_V1_7_0_WINDOWS2019() DotnetInstrumentationVersion {
	_init_.Initialize()
	var returns DotnetInstrumentationVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentationVersion",
		"V1_7_0_WINDOWS2019",
		&returns,
	)
	return returns
}

func DotnetInstrumentationVersion_V1_7_0_WINDOWS2022() DotnetInstrumentationVersion {
	_init_.Initialize()
	var returns DotnetInstrumentationVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotnetInstrumentationVersion",
		"V1_7_0_WINDOWS2022",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DotnetInstrumentationVersion) ImageURI() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"imageURI",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DotnetInstrumentationVersion) MemoryLimitMiB() *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"memoryLimitMiB",
		nil, // no parameters
		&returns,
	)

	return returns
}

