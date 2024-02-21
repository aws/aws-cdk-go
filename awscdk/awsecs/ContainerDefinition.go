package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// A container definition is used in a task definition to describe the containers that are launched as part of a task.
//
// Example:
//   var taskDefinition taskDefinition
//   var cluster cluster
//
//
//   // Add a container to the task definition
//   specificContainer := taskDefinition.AddContainer(jsii.String("Container"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("/aws/aws-example-app")),
//   	MemoryLimitMiB: jsii.Number(2048),
//   })
//
//   // Add a port mapping
//   specificContainer.AddPortMappings(&PortMapping{
//   	ContainerPort: jsii.Number(7600),
//   	Protocol: ecs.Protocol_TCP,
//   })
//
//   ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	CloudMapOptions: &CloudMapOptions{
//   		// Create SRV records - useful for bridge networking
//   		DnsRecordType: cloudmap.DnsRecordType_SRV,
//   		// Targets port TCP port 7600 `specificContainer`
//   		Container: specificContainer,
//   		ContainerPort: jsii.Number(7600),
//   	},
//   })
//
type ContainerDefinition interface {
	constructs.Construct
	// An array dependencies defined for container startup and shutdown.
	ContainerDependencies() *[]*ContainerDependency
	// The name of this container.
	ContainerName() *string
	// The port the container will listen on.
	ContainerPort() *float64
	// The number of cpu units reserved for the container.
	Cpu() *float64
	// The crdential specifications for this container.
	CredentialSpecs() *[]*CredentialSpecConfig
	// The environment files for this container.
	EnvironmentFiles() *[]*EnvironmentFileConfig
	// Specifies whether the container will be marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container
	// fails or stops for any reason, all other containers that are part of the task are
	// stopped. If the essential parameter of a container is marked as false, then its
	// failure does not affect the rest of the containers in a task.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	Essential() *bool
	// The name of the image referenced by this container.
	ImageName() *string
	// The inbound rules associated with the security group the task or service will use.
	//
	// This property is only used for tasks that use the awsvpc network mode.
	IngressPort() *float64
	// The Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	LinuxParameters() LinuxParameters
	// The log configuration specification for the container.
	LogDriverConfig() *LogDriverConfig
	// Whether there was at least one memory limit specified in this definition.
	MemoryLimitSpecified() *bool
	// The mount points for data volumes in your container.
	MountPoints() *[]*MountPoint
	// The tree node.
	Node() constructs.Node
	// The list of port mappings for the container.
	//
	// Port mappings allow containers to access ports
	// on the host container instance to send or receive traffic.
	PortMappings() *[]*PortMapping
	// Specifies whether a TTY must be allocated for this container.
	PseudoTerminal() *bool
	// Whether this container definition references a specific JSON field of a secret stored in Secrets Manager.
	ReferencesSecretJsonField() *bool
	// The name of the task definition that includes this container definition.
	TaskDefinition() TaskDefinition
	// An array of ulimits to set in the container.
	Ulimits() *[]*Ulimit
	// The data volumes to mount from another container in the same task definition.
	VolumesFrom() *[]*VolumeFrom
	// This method adds one or more container dependencies to the container.
	AddContainerDependencies(containerDependencies ...*ContainerDependency)
	// This method adds an environment variable to the container.
	AddEnvironment(name *string, value *string)
	// This method adds one or more resources to the container.
	AddInferenceAcceleratorResource(inferenceAcceleratorResources ...*string)
	// This method adds a link which allows containers to communicate with each other without the need for port mappings.
	//
	// This parameter is only supported if the task definition is using the bridge network mode.
	// Warning: The --link flag is a legacy feature of Docker. It may eventually be removed.
	AddLink(container ContainerDefinition, alias *string)
	// This method adds one or more mount points for data volumes to the container.
	AddMountPoints(mountPoints ...*MountPoint)
	// This method adds one or more port mappings to the container.
	AddPortMappings(portMappings ...*PortMapping)
	// This method mounts temporary disk space to the container.
	//
	// This adds the correct container mountPoint and task definition volume.
	AddScratch(scratch *ScratchSpace)
	// This method adds a secret as environment variable to the container.
	AddSecret(name *string, secret Secret)
	// This method adds the specified statement to the IAM task execution policy in the task definition.
	AddToExecutionPolicy(statement awsiam.PolicyStatement)
	// This method adds one or more ulimits to the container.
	AddUlimits(ulimits ...*Ulimit)
	// This method adds one or more volumes to the container.
	AddVolumesFrom(volumesFrom ...*VolumeFrom)
	// Returns the host port for the requested container port if it exists.
	FindPortMapping(containerPort *float64, protocol Protocol) *PortMapping
	// Returns the port mapping with the given name, if it exists.
	FindPortMappingByName(name *string) *PortMapping
	// Render this container definition to a CloudFormation object.
	RenderContainerDefinition(_taskDefinition TaskDefinition) *CfnTaskDefinition_ContainerDefinitionProperty
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ContainerDefinition
type jsiiProxy_ContainerDefinition struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_ContainerDefinition) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerDefinition) CredentialSpecs() *[]*CredentialSpecConfig {
	var returns *[]*CredentialSpecConfig
	_jsii_.Get(
		j,
		"credentialSpecs",
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

func (j *jsiiProxy_ContainerDefinition) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_ContainerDefinition) PseudoTerminal() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"pseudoTerminal",
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
func NewContainerDefinition(scope constructs.Construct, id *string, props *ContainerDefinitionProps) ContainerDefinition {
	_init_.Initialize()

	if err := validateNewContainerDefinitionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ContainerDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ContainerDefinition class.
func NewContainerDefinition_Override(c ContainerDefinition, scope constructs.Construct, id *string, props *ContainerDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ContainerDefinition",
		[]interface{}{scope, id, props},
		c,
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
func ContainerDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ContainerDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContainerDefinition_CONTAINER_PORT_USE_RANGE() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.ContainerDefinition",
		"CONTAINER_PORT_USE_RANGE",
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

func (c *jsiiProxy_ContainerDefinition) AddSecret(name *string, secret Secret) {
	if err := c.validateAddSecretParameters(name, secret); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addSecret",
		[]interface{}{name, secret},
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

func (c *jsiiProxy_ContainerDefinition) FindPortMappingByName(name *string) *PortMapping {
	if err := c.validateFindPortMappingByNameParameters(name); err != nil {
		panic(err)
	}
	var returns *PortMapping

	_jsii_.Invoke(
		c,
		"findPortMappingByName",
		[]interface{}{name},
		&returns,
	)

	return returns
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

