// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A container orchestrated by ECS that uses Fargate resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var containerImage containerImage
//   var ecsVolume ecsVolume
//   var linuxParameters linuxParameters
//   var logDriver logDriver
//   var role role
//   var secret secret
//   var size size
//
//   ecsFargateContainerDefinition := batch_alpha.NewEcsFargateContainerDefinition(this, jsii.String("MyEcsFargateContainerDefinition"), &EcsFargateContainerDefinitionProps{
//   	Cpu: jsii.Number(123),
//   	Image: containerImage,
//   	Memory: size,
//
//   	// the properties below are optional
//   	AssignPublicIp: jsii.Boolean(false),
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	ExecutionRole: role,
//   	FargatePlatformVersion: awscdk.Aws_ecs.FargatePlatformVersion_LATEST,
//   	JobRole: role,
//   	LinuxParameters: linuxParameters,
//   	Logging: logDriver,
//   	ReadonlyRootFilesystem: jsii.Boolean(false),
//   	Secrets: []iSecret{
//   		secret,
//   	},
//   	User: jsii.String("user"),
//   	Volumes: []*ecsVolume{
//   		ecsVolume,
//   	},
//   })
//
// Experimental.
type EcsFargateContainerDefinition interface {
	constructs.Construct
	IEcsContainerDefinition
	IEcsFargateContainerDefinition
	// Indicates whether the job has a public IP address.
	//
	// For a job that's running on Fargate resources in a private subnet to send outbound traffic to the internet
	// (for example, to pull container images), the private subnet requires a NAT gateway be attached to route requests to the internet.
	// Experimental.
	AssignPublicIp() *bool
	// The command that's passed to the container.
	// Experimental.
	Command() *[]*string
	// The number of vCPUs reserved for the container.
	//
	// Each vCPU is equivalent to 1,024 CPU shares.
	// For containers running on EC2 resources, you must specify at least one vCPU.
	// Experimental.
	Cpu() *float64
	// The environment variables to pass to a container.
	//
	// Cannot start with `AWS_BATCH`.
	// We don't recommend using plaintext environment variables for sensitive information, such as credential data.
	// Experimental.
	Environment() *map[string]*string
	// The role used by Amazon ECS container and AWS Fargate agents to make AWS API calls on your behalf.
	// Experimental.
	ExecutionRole() awsiam.IRole
	// Which version of Fargate to use when running this container.
	// Experimental.
	FargatePlatformVersion() awsecs.FargatePlatformVersion
	// The image that this container will run.
	// Experimental.
	Image() awsecs.ContainerImage
	// The role that the container can assume.
	// Experimental.
	JobRole() awsiam.IRole
	// Linux-specific modifications that are applied to the container, such as details for device mappings.
	// Experimental.
	LinuxParameters() LinuxParameters
	// The configuration of the log driver.
	// Experimental.
	LogDriverConfig() *awsecs.LogDriverConfig
	// The memory hard limit present to the container.
	//
	// If your container attempts to exceed the memory specified, the container is terminated.
	// You must specify at least 4 MiB of memory for a job.
	// Experimental.
	Memory() awscdk.Size
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Gives the container readonly access to its root filesystem.
	// Experimental.
	ReadonlyRootFilesystem() *bool
	// The secrets for the container.
	//
	// Can be referenced in your job definition.
	// Experimental.
	Secrets() *[]awssecretsmanager.ISecret
	// The user name to use inside the container.
	// Experimental.
	User() *string
	// The volumes to mount to this container.
	//
	// Automatically added to the job definition.
	// Experimental.
	Volumes() *[]EcsVolume
	// Add a Volume to this container.
	// Experimental.
	AddVolume(volume EcsVolume)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsFargateContainerDefinition
type jsiiProxy_EcsFargateContainerDefinition struct {
	internal.Type__constructsConstruct
	jsiiProxy_IEcsContainerDefinition
	jsiiProxy_IEcsFargateContainerDefinition
}

func (j *jsiiProxy_EcsFargateContainerDefinition) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) FargatePlatformVersion() awsecs.FargatePlatformVersion {
	var returns awsecs.FargatePlatformVersion
	_jsii_.Get(
		j,
		"fargatePlatformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Image() awsecs.ContainerImage {
	var returns awsecs.ContainerImage
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) JobRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"jobRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) LinuxParameters() LinuxParameters {
	var returns LinuxParameters
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) LogDriverConfig() *awsecs.LogDriverConfig {
	var returns *awsecs.LogDriverConfig
	_jsii_.Get(
		j,
		"logDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Memory() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) ReadonlyRootFilesystem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Secrets() *[]awssecretsmanager.ISecret {
	var returns *[]awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Volumes() *[]EcsVolume {
	var returns *[]EcsVolume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


// Experimental.
func NewEcsFargateContainerDefinition(scope constructs.Construct, id *string, props *EcsFargateContainerDefinitionProps) EcsFargateContainerDefinition {
	_init_.Initialize()

	if err := validateNewEcsFargateContainerDefinitionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsFargateContainerDefinition{}

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.EcsFargateContainerDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcsFargateContainerDefinition_Override(e EcsFargateContainerDefinition, scope constructs.Construct, id *string, props *EcsFargateContainerDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.EcsFargateContainerDefinition",
		[]interface{}{scope, id, props},
		e,
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
// Experimental.
func EcsFargateContainerDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsFargateContainerDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.EcsFargateContainerDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsFargateContainerDefinition) AddVolume(volume EcsVolume) {
	if err := e.validateAddVolumeParameters(volume); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addVolume",
		[]interface{}{volume},
	)
}

func (e *jsiiProxy_EcsFargateContainerDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

