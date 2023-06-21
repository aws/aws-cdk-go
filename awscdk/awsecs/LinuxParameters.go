package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Linux-specific options that are applied to the container.
//
// Example:
//   var taskDefinition taskDefinition
//
//
//   taskDefinition.AddContainer(jsii.String("container"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryLimitMiB: jsii.Number(1024),
//   	LinuxParameters: ecs.NewLinuxParameters(this, jsii.String("LinuxParameters"), &LinuxParametersProps{
//   		InitProcessEnabled: jsii.Boolean(true),
//   		SharedMemorySize: jsii.Number(1024),
//   		MaxSwap: awscdk.Size_Mebibytes(jsii.Number(5000)),
//   		Swappiness: jsii.Number(90),
//   	}),
//   })
//
type LinuxParameters interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Adds one or more Linux capabilities to the Docker configuration of a container.
	//
	// Tasks launched on Fargate only support adding the 'SYS_PTRACE' kernel capability.
	AddCapabilities(cap ...Capability)
	// Adds one or more host devices to a container.
	AddDevices(device ...*Device)
	// Specifies the container path, mount options, and size (in MiB) of the tmpfs mount for a container.
	//
	// Only works with EC2 launch type.
	AddTmpfs(tmpfs ...*Tmpfs)
	// Removes one or more Linux capabilities to the Docker configuration of a container.
	DropCapabilities(cap ...Capability)
	// Renders the Linux parameters to a CloudFormation object.
	RenderLinuxParameters() *CfnTaskDefinition_LinuxParametersProperty
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for LinuxParameters
type jsiiProxy_LinuxParameters struct {
	internal.Type__constructsConstruct
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


// Constructs a new instance of the LinuxParameters class.
func NewLinuxParameters(scope constructs.Construct, id *string, props *LinuxParametersProps) LinuxParameters {
	_init_.Initialize()

	if err := validateNewLinuxParametersParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxParameters{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.LinuxParameters",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the LinuxParameters class.
func NewLinuxParameters_Override(l LinuxParameters, scope constructs.Construct, id *string, props *LinuxParametersProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.LinuxParameters",
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
		"aws-cdk-lib.aws_ecs.LinuxParameters",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxParameters) AddCapabilities(cap ...Capability) {
	args := []interface{}{}
	for _, a := range cap {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"addCapabilities",
		args,
	)
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

func (l *jsiiProxy_LinuxParameters) DropCapabilities(cap ...Capability) {
	args := []interface{}{}
	for _, a := range cap {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"dropCapabilities",
		args,
	)
}

func (l *jsiiProxy_LinuxParameters) RenderLinuxParameters() *CfnTaskDefinition_LinuxParametersProperty {
	var returns *CfnTaskDefinition_LinuxParametersProperty

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

