package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Linux-specific options that are applied to the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var size Size
//
//   linuxParameters := awscdk.Aws_batch.NewLinuxParameters(this, jsii.String("MyLinuxParameters"), &LinuxParametersProps{
//   	InitProcessEnabled: jsii.Boolean(false),
//   	MaxSwap: size,
//   	SharedMemorySize: size,
//   	Swappiness: jsii.Number(123),
//   })
//
type LinuxParameters interface {
	constructs.Construct
	// Device mounts.
	Devices() *[]*Device
	// Whether the init process is enabled.
	InitProcessEnabled() *bool
	// The max swap memory.
	MaxSwap() awscdk.Size
	// The tree node.
	Node() constructs.Node
	// The shared memory size (in MiB).
	//
	// Not valid for Fargate launch type.
	SharedMemorySize() awscdk.Size
	// The swappiness behavior.
	Swappiness() *float64
	// TmpFs mounts.
	Tmpfs() *[]*Tmpfs
	// Adds one or more host devices to a container.
	AddDevices(device ...*Device)
	// Specifies the container path, mount options, and size (in MiB) of the tmpfs mount for a container.
	//
	// Only works with EC2 launch type.
	AddTmpfs(tmpfs ...*Tmpfs)
	// Renders the Linux parameters to the Batch version of this resource, which does not have 'capabilities' and requires tmpfs.containerPath to be defined.
	RenderLinuxParameters() *CfnJobDefinition_LinuxParametersProperty
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for LinuxParameters
type jsiiProxy_LinuxParameters struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LinuxParameters) Devices() *[]*Device {
	var returns *[]*Device
	_jsii_.Get(
		j,
		"devices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxParameters) InitProcessEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"initProcessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxParameters) MaxSwap() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"maxSwap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxParameters) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxParameters) SharedMemorySize() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"sharedMemorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxParameters) Swappiness() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swappiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxParameters) Tmpfs() *[]*Tmpfs {
	var returns *[]*Tmpfs
	_jsii_.Get(
		j,
		"tmpfs",
		&returns,
	)
	return returns
}


// Constructs a new instance of the LinuxParameters class.
func NewLinuxParameters(scope constructs.Construct, id *string, props *LinuxParametersProps) LinuxParameters {
	_init_.Initialize()

	if err := validateNewLinuxParametersParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxParameters{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.LinuxParameters",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the LinuxParameters class.
func NewLinuxParameters_Override(l LinuxParameters, scope constructs.Construct, id *string, props *LinuxParametersProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.LinuxParameters",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func LinuxParameters_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxParameters_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.LinuxParameters",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxParameters) AddDevices(device ...*Device) {
	if err := l.validateAddDevicesParameters(&device); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range device {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"addDevices",
		args,
	)
}

func (l *jsiiProxy_LinuxParameters) AddTmpfs(tmpfs ...*Tmpfs) {
	if err := l.validateAddTmpfsParameters(&tmpfs); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range tmpfs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"addTmpfs",
		args,
	)
}

func (l *jsiiProxy_LinuxParameters) RenderLinuxParameters() *CfnJobDefinition_LinuxParametersProperty {
	var returns *CfnJobDefinition_LinuxParametersProperty

	_jsii_.Invoke(
		l,
		"renderLinuxParameters",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxParameters) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

