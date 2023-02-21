package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

// A container definition is used in a task definition to describe the containers that are launched as part of a task.
//
// Example:
//   var taskDefinition taskDefinition
//   var cluster cluster
//
//
//   // Add a container to the task definition
//   specificContainer := taskDefinition.addContainer(jsii.String("Container"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("/aws/aws-example-app")),
//   	memoryLimitMiB: jsii.Number(2048),
//   })
//
//   // Add a port mapping
//   specificContainer.addPortMappings(&portMapping{
//   	containerPort: jsii.Number(7600),
//   	protocol: ecs.protocol_TCP,
//   })
//
//   ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	cloudMapOptions: &cloudMapOptions{
//   		// Create SRV records - useful for bridge networking
//   		dnsRecordType: cloudmap.dnsRecordType_SRV,
//   		// Targets port TCP port 7600 `specificContainer`
//   		container: specificContainer,
//   		containerPort: jsii.Number(7600),
//   	},
//   })
//
// Experimental.
type ContainerDefinition interface {
	awscdk.Construct
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

// The jsii proxy struct for ContainerDefinition
type jsiiProxy_ContainerDefinition struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_ContainerDefinition) ContainerDependencies() *[]*ContainerDependency {
	var returns *[]*ContainerDependency
	_jsii_.Get(
		j,
		"containerDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) ContainerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) EnvironmentFiles() *[]*EnvironmentFileConfig {
	var returns *[]*EnvironmentFileConfig
	_jsii_.Get(
		j,
		"environmentFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) Essential() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) IngressPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) LinuxParameters() LinuxParameters {
	var returns LinuxParameters
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) LogDriverConfig() *LogDriverConfig {
	var returns *LogDriverConfig
	_jsii_.Get(
		j,
		"logDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) MemoryLimitSpecified() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"memoryLimitSpecified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) MountPoints() *[]*MountPoint {
	var returns *[]*MountPoint
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) PortMappings() *[]*PortMapping {
	var returns *[]*PortMapping
	_jsii_.Get(
		j,
		"portMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) ReferencesSecretJsonField() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"referencesSecretJsonField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) TaskDefinition() TaskDefinition {
	var returns TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) Ulimits() *[]*Ulimit {
	var returns *[]*Ulimit
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) VolumesFrom() *[]*VolumeFrom {
	var returns *[]*VolumeFrom
	_jsii_.Get(
		j,
		"volumesFrom",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ContainerDefinition class.
// Experimental.
func NewContainerDefinition(scope constructs.Construct, id *string, props *ContainerDefinitionProps) ContainerDefinition {
	_init_.Initialize()

	if err := validateNewContainerDefinitionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerDefinition{}

	_jsii_.Create(
		"monocdk.aws_ecs.ContainerDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ContainerDefinition class.
// Experimental.
func NewContainerDefinition_Override(c ContainerDefinition, scope constructs.Construct, id *string, props *ContainerDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.ContainerDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func ContainerDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.ContainerDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerDefinition) AddContainerDependencies(containerDependencies ...*ContainerDependency) {
	if err := c.validateAddContainerDependenciesParameters(&containerDependencies); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range containerDependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addContainerDependencies",
		args,
	)
}

func (c *jsiiProxy_ContainerDefinition) AddEnvironment(name *string, value *string) {
	if err := c.validateAddEnvironmentParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addEnvironment",
		[]interface{}{name, value},
	)
}

func (c *jsiiProxy_ContainerDefinition) AddInferenceAcceleratorResource(inferenceAcceleratorResources ...*string) {
	args := []interface{}{}
	for _, a := range inferenceAcceleratorResources {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addInferenceAcceleratorResource",
		args,
	)
}

func (c *jsiiProxy_ContainerDefinition) AddLink(container ContainerDefinition, alias *string) {
	if err := c.validateAddLinkParameters(container); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addLink",
		[]interface{}{container, alias},
	)
}

func (c *jsiiProxy_ContainerDefinition) AddMountPoints(mountPoints ...*MountPoint) {
	if err := c.validateAddMountPointsParameters(&mountPoints); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range mountPoints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addMountPoints",
		args,
	)
}

func (c *jsiiProxy_ContainerDefinition) AddPortMappings(portMappings ...*PortMapping) {
	if err := c.validateAddPortMappingsParameters(&portMappings); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range portMappings {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addPortMappings",
		args,
	)
}

func (c *jsiiProxy_ContainerDefinition) AddScratch(scratch *ScratchSpace) {
	if err := c.validateAddScratchParameters(scratch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addScratch",
		[]interface{}{scratch},
	)
}

func (c *jsiiProxy_ContainerDefinition) AddToExecutionPolicy(statement awsiam.PolicyStatement) {
	if err := c.validateAddToExecutionPolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addToExecutionPolicy",
		[]interface{}{statement},
	)
}

func (c *jsiiProxy_ContainerDefinition) AddUlimits(ulimits ...*Ulimit) {
	if err := c.validateAddUlimitsParameters(&ulimits); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range ulimits {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addUlimits",
		args,
	)
}

func (c *jsiiProxy_ContainerDefinition) AddVolumesFrom(volumesFrom ...*VolumeFrom) {
	if err := c.validateAddVolumesFromParameters(&volumesFrom); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range volumesFrom {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addVolumesFrom",
		args,
	)
}

func (c *jsiiProxy_ContainerDefinition) FindPortMapping(containerPort *float64, protocol Protocol) *PortMapping {
	if err := c.validateFindPortMappingParameters(containerPort, protocol); err != nil {
		panic(err)
	}
	var returns *PortMapping

	_jsii_.Invoke(
		c,
		"findPortMapping",
		[]interface{}{containerPort, protocol},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerDefinition) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerDefinition) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_ContainerDefinition) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerDefinition) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerDefinition) RenderContainerDefinition(_taskDefinition TaskDefinition) *CfnTaskDefinition_ContainerDefinitionProperty {
	var returns *CfnTaskDefinition_ContainerDefinitionProperty

	_jsii_.Invoke(
		c,
		"renderContainerDefinition",
		[]interface{}{_taskDefinition},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerDefinition) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_ContainerDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerDefinition) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

