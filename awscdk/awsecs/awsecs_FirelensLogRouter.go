package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

// Firelens log router.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var containerImage containerImage
//   var duration duration
//   var environmentFile environmentFile
//   var linuxParameters linuxParameters
//   var logDriver logDriver
//   var secret secret
//   var taskDefinition taskDefinition
//
//   firelensLogRouter := awscdk.Aws_ecs.NewFirelensLogRouter(this, jsii.String("MyFirelensLogRouter"), &firelensLogRouterProps{
//   	firelensConfig: &firelensConfig{
//   		type: awscdk.*Aws_ecs.firelensLogRouterType_FLUENTBIT,
//
//   		// the properties below are optional
//   		options: &firelensOptions{
//   			configFileType: awscdk.*Aws_ecs.firelensConfigFileType_S3,
//   			configFileValue: jsii.String("configFileValue"),
//   			enableECSLogMetadata: jsii.Boolean(false),
//   		},
//   	},
//   	image: containerImage,
//   	taskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	containerName: jsii.String("containerName"),
//   	cpu: jsii.Number(123),
//   	disableNetworking: jsii.Boolean(false),
//   	dnsSearchDomains: []*string{
//   		jsii.String("dnsSearchDomains"),
//   	},
//   	dnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	dockerLabels: map[string]*string{
//   		"dockerLabelsKey": jsii.String("dockerLabels"),
//   	},
//   	dockerSecurityOptions: []*string{
//   		jsii.String("dockerSecurityOptions"),
//   	},
//   	entryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	environmentFiles: []*environmentFile{
//   		environmentFile,
//   	},
//   	essential: jsii.Boolean(false),
//   	extraHosts: map[string]*string{
//   		"extraHostsKey": jsii.String("extraHosts"),
//   	},
//   	gpuCount: jsii.Number(123),
//   	healthCheck: &healthCheck{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//
//   		// the properties below are optional
//   		interval: duration,
//   		retries: jsii.Number(123),
//   		startPeriod: duration,
//   		timeout: duration,
//   	},
//   	hostname: jsii.String("hostname"),
//   	inferenceAcceleratorResources: []*string{
//   		jsii.String("inferenceAcceleratorResources"),
//   	},
//   	linuxParameters: linuxParameters,
//   	logging: logDriver,
//   	memoryLimitMiB: jsii.Number(123),
//   	memoryReservationMiB: jsii.Number(123),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			hostPort: jsii.Number(123),
//   			protocol: awscdk.*Aws_ecs.protocol_TCP,
//   		},
//   	},
//   	privileged: jsii.Boolean(false),
//   	readonlyRootFilesystem: jsii.Boolean(false),
//   	secrets: map[string]*secret{
//   		"secretsKey": secret,
//   	},
//   	startTimeout: duration,
//   	stopTimeout: duration,
//   	systemControls: []systemControl{
//   		&systemControl{
//   			namespace: jsii.String("namespace"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	user: jsii.String("user"),
//   	workingDirectory: jsii.String("workingDirectory"),
//   })
//
// Experimental.
type FirelensLogRouter interface {
	ContainerDefinition
	// An array dependencies defined for container startup and shutdown.
	// Experimental.
	ContainerDependencies() *[]*ContainerDependency
	// The name of this container.
	// Experimental.
	ContainerName() *string
	// The port the container will listen on.
	// Experimental.
	ContainerPort() *float64
	// The environment files for this container.
	// Experimental.
	EnvironmentFiles() *[]*EnvironmentFileConfig
	// Specifies whether the container will be marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container
	// fails or stops for any reason, all other containers that are part of the task are
	// stopped. If the essential parameter of a container is marked as false, then its
	// failure does not affect the rest of the containers in a task.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	// Experimental.
	Essential() *bool
	// Firelens configuration.
	// Experimental.
	FirelensConfig() *FirelensConfig
	// The name of the image referenced by this container.
	// Experimental.
	ImageName() *string
	// The inbound rules associated with the security group the task or service will use.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	// Experimental.
	IngressPort() *float64
	// The Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	// Experimental.
	LinuxParameters() LinuxParameters
	// The log configuration specification for the container.
	// Experimental.
	LogDriverConfig() *LogDriverConfig
	// Whether there was at least one memory limit specified in this definition.
	// Experimental.
	MemoryLimitSpecified() *bool
	// The mount points for data volumes in your container.
	// Experimental.
	MountPoints() *[]*MountPoint
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The list of port mappings for the container.
	//
	// Port mappings allow containers to access ports
	// on the host container instance to send or receive traffic.
	// Experimental.
	PortMappings() *[]*PortMapping
	// Whether this container definition references a specific JSON field of a secret stored in Secrets Manager.
	// Experimental.
	ReferencesSecretJsonField() *bool
	// The name of the task definition that includes this container definition.
	// Experimental.
	TaskDefinition() TaskDefinition
	// An array of ulimits to set in the container.
	// Experimental.
	Ulimits() *[]*Ulimit
	// The data volumes to mount from another container in the same task definition.
	// Experimental.
	VolumesFrom() *[]*VolumeFrom
	// This method adds one or more container dependencies to the container.
	// Experimental.
	AddContainerDependencies(containerDependencies ...*ContainerDependency)
	// This method adds an environment variable to the container.
	// Experimental.
	AddEnvironment(name *string, value *string)
	// This method adds one or more resources to the container.
	// Experimental.
	AddInferenceAcceleratorResource(inferenceAcceleratorResources ...*string)
	// This method adds a link which allows containers to communicate with each other without the need for port mappings.
	//
	// This parameter is only supported if the task definition is using the bridge network mode.
	// Warning: The --link flag is a legacy feature of Docker. It may eventually be removed.
	// Experimental.
	AddLink(container ContainerDefinition, alias *string)
	// This method adds one or more mount points for data volumes to the container.
	// Experimental.
	AddMountPoints(mountPoints ...*MountPoint)
	// This method adds one or more port mappings to the container.
	// Experimental.
	AddPortMappings(portMappings ...*PortMapping)
	// This method mounts temporary disk space to the container.
	//
	// This adds the correct container mountPoint and task definition volume.
	// Experimental.
	AddScratch(scratch *ScratchSpace)
	// This method adds the specified statement to the IAM task execution policy in the task definition.
	// Experimental.
	AddToExecutionPolicy(statement awsiam.PolicyStatement)
	// This method adds one or more ulimits to the container.
	// Experimental.
	AddUlimits(ulimits ...*Ulimit)
	// This method adds one or more volumes to the container.
	// Experimental.
	AddVolumesFrom(volumesFrom ...*VolumeFrom)
	// Returns the host port for the requested container port if it exists.
	// Experimental.
	FindPortMapping(containerPort *float64, protocol Protocol) *PortMapping
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Render this container definition to a CloudFormation object.
	// Experimental.
	RenderContainerDefinition(_taskDefinition TaskDefinition) *CfnTaskDefinition_ContainerDefinitionProperty
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for FirelensLogRouter
type jsiiProxy_FirelensLogRouter struct {
	jsiiProxy_ContainerDefinition
}

func (j *jsiiProxy_FirelensLogRouter) ContainerDependencies() *[]*ContainerDependency {
	var returns *[]*ContainerDependency
	_jsii_.Get(
		j,
		"containerDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) ContainerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) EnvironmentFiles() *[]*EnvironmentFileConfig {
	var returns *[]*EnvironmentFileConfig
	_jsii_.Get(
		j,
		"environmentFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) Essential() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) FirelensConfig() *FirelensConfig {
	var returns *FirelensConfig
	_jsii_.Get(
		j,
		"firelensConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) IngressPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) LinuxParameters() LinuxParameters {
	var returns LinuxParameters
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) LogDriverConfig() *LogDriverConfig {
	var returns *LogDriverConfig
	_jsii_.Get(
		j,
		"logDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) MemoryLimitSpecified() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"memoryLimitSpecified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) MountPoints() *[]*MountPoint {
	var returns *[]*MountPoint
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) PortMappings() *[]*PortMapping {
	var returns *[]*PortMapping
	_jsii_.Get(
		j,
		"portMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) ReferencesSecretJsonField() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"referencesSecretJsonField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) TaskDefinition() TaskDefinition {
	var returns TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) Ulimits() *[]*Ulimit {
	var returns *[]*Ulimit
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) VolumesFrom() *[]*VolumeFrom {
	var returns *[]*VolumeFrom
	_jsii_.Get(
		j,
		"volumesFrom",
		&returns,
	)
	return returns
}


// Constructs a new instance of the FirelensLogRouter class.
// Experimental.
func NewFirelensLogRouter(scope constructs.Construct, id *string, props *FirelensLogRouterProps) FirelensLogRouter {
	_init_.Initialize()

	j := jsiiProxy_FirelensLogRouter{}

	_jsii_.Create(
		"monocdk.aws_ecs.FirelensLogRouter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the FirelensLogRouter class.
// Experimental.
func NewFirelensLogRouter_Override(f FirelensLogRouter, scope constructs.Construct, id *string, props *FirelensLogRouterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.FirelensLogRouter",
		[]interface{}{scope, id, props},
		f,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func FirelensLogRouter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.FirelensLogRouter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirelensLogRouter) AddContainerDependencies(containerDependencies ...*ContainerDependency) {
	args := []interface{}{}
	for _, a := range containerDependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addContainerDependencies",
		args,
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddEnvironment(name *string, value *string) {
	_jsii_.InvokeVoid(
		f,
		"addEnvironment",
		[]interface{}{name, value},
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddInferenceAcceleratorResource(inferenceAcceleratorResources ...*string) {
	args := []interface{}{}
	for _, a := range inferenceAcceleratorResources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addInferenceAcceleratorResource",
		args,
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddLink(container ContainerDefinition, alias *string) {
	_jsii_.InvokeVoid(
		f,
		"addLink",
		[]interface{}{container, alias},
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddMountPoints(mountPoints ...*MountPoint) {
	args := []interface{}{}
	for _, a := range mountPoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addMountPoints",
		args,
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddPortMappings(portMappings ...*PortMapping) {
	args := []interface{}{}
	for _, a := range portMappings {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addPortMappings",
		args,
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddScratch(scratch *ScratchSpace) {
	_jsii_.InvokeVoid(
		f,
		"addScratch",
		[]interface{}{scratch},
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddToExecutionPolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		f,
		"addToExecutionPolicy",
		[]interface{}{statement},
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddUlimits(ulimits ...*Ulimit) {
	args := []interface{}{}
	for _, a := range ulimits {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addUlimits",
		args,
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddVolumesFrom(volumesFrom ...*VolumeFrom) {
	args := []interface{}{}
	for _, a := range volumesFrom {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addVolumesFrom",
		args,
	)
}

func (f *jsiiProxy_FirelensLogRouter) FindPortMapping(containerPort *float64, protocol Protocol) *PortMapping {
	var returns *PortMapping

	_jsii_.Invoke(
		f,
		"findPortMapping",
		[]interface{}{containerPort, protocol},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirelensLogRouter) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirelensLogRouter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FirelensLogRouter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirelensLogRouter) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirelensLogRouter) RenderContainerDefinition(_taskDefinition TaskDefinition) *CfnTaskDefinition_ContainerDefinitionProperty {
	var returns *CfnTaskDefinition_ContainerDefinitionProperty

	_jsii_.Invoke(
		f,
		"renderContainerDefinition",
		[]interface{}{_taskDefinition},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirelensLogRouter) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FirelensLogRouter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirelensLogRouter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

