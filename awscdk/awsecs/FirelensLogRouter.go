package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Firelens log router.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var appProtocol appProtocol
//   var containerImage containerImage
//   var credentialSpec credentialSpec
//   var environmentFile environmentFile
//   var linuxParameters linuxParameters
//   var logDriver logDriver
//   var secret secret
//   var taskDefinition taskDefinition
//
//   firelensLogRouter := awscdk.Aws_ecs.NewFirelensLogRouter(this, jsii.String("MyFirelensLogRouter"), &FirelensLogRouterProps{
//   	FirelensConfig: &FirelensConfig{
//   		Type: awscdk.*Aws_ecs.FirelensLogRouterType_FLUENTBIT,
//
//   		// the properties below are optional
//   		Options: &FirelensOptions{
//   			ConfigFileType: awscdk.*Aws_ecs.FirelensConfigFileType_S3,
//   			ConfigFileValue: jsii.String("configFileValue"),
//   			EnableECSLogMetadata: jsii.Boolean(false),
//   		},
//   	},
//   	Image: containerImage,
//   	TaskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	ContainerName: jsii.String("containerName"),
//   	Cpu: jsii.Number(123),
//   	CredentialSpecs: []*credentialSpec{
//   		credentialSpec,
//   	},
//   	DisableNetworking: jsii.Boolean(false),
//   	DnsSearchDomains: []*string{
//   		jsii.String("dnsSearchDomains"),
//   	},
//   	DnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	DockerLabels: map[string]*string{
//   		"dockerLabelsKey": jsii.String("dockerLabels"),
//   	},
//   	DockerSecurityOptions: []*string{
//   		jsii.String("dockerSecurityOptions"),
//   	},
//   	EntryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	EnvironmentFiles: []*environmentFile{
//   		environmentFile,
//   	},
//   	Essential: jsii.Boolean(false),
//   	ExtraHosts: map[string]*string{
//   		"extraHostsKey": jsii.String("extraHosts"),
//   	},
//   	GpuCount: jsii.Number(123),
//   	HealthCheck: &HealthCheck{
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//
//   		// the properties below are optional
//   		Interval: cdk.Duration_Minutes(jsii.Number(30)),
//   		Retries: jsii.Number(123),
//   		StartPeriod: cdk.Duration_*Minutes(jsii.Number(30)),
//   		Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	},
//   	Hostname: jsii.String("hostname"),
//   	InferenceAcceleratorResources: []*string{
//   		jsii.String("inferenceAcceleratorResources"),
//   	},
//   	Interactive: jsii.Boolean(false),
//   	LinuxParameters: linuxParameters,
//   	Logging: logDriver,
//   	MemoryLimitMiB: jsii.Number(123),
//   	MemoryReservationMiB: jsii.Number(123),
//   	PortMappings: []portMapping{
//   		&portMapping{
//   			ContainerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			AppProtocol: appProtocol,
//   			ContainerPortRange: jsii.String("containerPortRange"),
//   			HostPort: jsii.Number(123),
//   			Name: jsii.String("name"),
//   			Protocol: awscdk.*Aws_ecs.Protocol_TCP,
//   		},
//   	},
//   	Privileged: jsii.Boolean(false),
//   	PseudoTerminal: jsii.Boolean(false),
//   	ReadonlyRootFilesystem: jsii.Boolean(false),
//   	Secrets: map[string]*secret{
//   		"secretsKey": secret,
//   	},
//   	StartTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	StopTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	SystemControls: []systemControl{
//   		&systemControl{
//   			Namespace: jsii.String("namespace"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Ulimits: []ulimit{
//   		&ulimit{
//   			HardLimit: jsii.Number(123),
//   			Name: awscdk.*Aws_ecs.UlimitName_CORE,
//   			SoftLimit: jsii.Number(123),
//   		},
//   	},
//   	User: jsii.String("user"),
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   })
//
type FirelensLogRouter interface {
	ContainerDefinition
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
	// Firelens configuration.
	FirelensConfig() *FirelensConfig
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

func (j *jsiiProxy_FirelensLogRouter) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirelensLogRouter) CredentialSpecs() *[]*CredentialSpecConfig {
	var returns *[]*CredentialSpecConfig
	_jsii_.Get(
		j,
		"credentialSpecs",
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

func (j *jsiiProxy_FirelensLogRouter) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_FirelensLogRouter) PseudoTerminal() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"pseudoTerminal",
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
func NewFirelensLogRouter(scope constructs.Construct, id *string, props *FirelensLogRouterProps) FirelensLogRouter {
	_init_.Initialize()

	if err := validateNewFirelensLogRouterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirelensLogRouter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FirelensLogRouter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the FirelensLogRouter class.
func NewFirelensLogRouter_Override(f FirelensLogRouter, scope constructs.Construct, id *string, props *FirelensLogRouterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FirelensLogRouter",
		[]interface{}{scope, id, props},
		f,
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
func FirelensLogRouter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFirelensLogRouter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FirelensLogRouter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FirelensLogRouter_CONTAINER_PORT_USE_RANGE() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.FirelensLogRouter",
		"CONTAINER_PORT_USE_RANGE",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FirelensLogRouter) AddContainerDependencies(containerDependencies ...*ContainerDependency) {
	if err := f.validateAddContainerDependenciesParameters(&containerDependencies); err != nil {
		panic(err)
	}
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
	if err := f.validateAddEnvironmentParameters(name, value); err != nil {
		panic(err)
	}
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
	if err := f.validateAddLinkParameters(container); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addLink",
		[]interface{}{container, alias},
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddMountPoints(mountPoints ...*MountPoint) {
	if err := f.validateAddMountPointsParameters(&mountPoints); err != nil {
		panic(err)
	}
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
	if err := f.validateAddPortMappingsParameters(&portMappings); err != nil {
		panic(err)
	}
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
	if err := f.validateAddScratchParameters(scratch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addScratch",
		[]interface{}{scratch},
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddSecret(name *string, secret Secret) {
	if err := f.validateAddSecretParameters(name, secret); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addSecret",
		[]interface{}{name, secret},
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddToExecutionPolicy(statement awsiam.PolicyStatement) {
	if err := f.validateAddToExecutionPolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addToExecutionPolicy",
		[]interface{}{statement},
	)
}

func (f *jsiiProxy_FirelensLogRouter) AddUlimits(ulimits ...*Ulimit) {
	if err := f.validateAddUlimitsParameters(&ulimits); err != nil {
		panic(err)
	}
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
	if err := f.validateAddVolumesFromParameters(&volumesFrom); err != nil {
		panic(err)
	}
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
	if err := f.validateFindPortMappingParameters(containerPort, protocol); err != nil {
		panic(err)
	}
	var returns *PortMapping

	_jsii_.Invoke(
		f,
		"findPortMapping",
		[]interface{}{containerPort, protocol},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirelensLogRouter) FindPortMappingByName(name *string) *PortMapping {
	if err := f.validateFindPortMappingByNameParameters(name); err != nil {
		panic(err)
	}
	var returns *PortMapping

	_jsii_.Invoke(
		f,
		"findPortMappingByName",
		[]interface{}{name},
		&returns,
	)

	return returns
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

