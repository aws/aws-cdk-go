package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

// The details of a daemon task definition.
//
// A daemon task definition is a template that describes the containers that form a daemon. Daemons deploy cross-cutting software agents independently across your Amazon ECS infrastructure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDaemonTaskDefinition := awscdk.Aws_ecs.NewCfnDaemonTaskDefinition(this, jsii.String("MyCfnDaemonTaskDefinition"), &CfnDaemonTaskDefinitionProps{
//   	ContainerDefinitions: []interface{}{
//   		&DaemonContainerDefinitionProperty{
//   			Image: jsii.String("image"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Cpu: jsii.Number(123),
//   			DependsOn: []interface{}{
//   				&ContainerDependencyProperty{
//   					Condition: jsii.String("condition"),
//   					ContainerName: jsii.String("containerName"),
//   				},
//   			},
//   			EntryPoint: []*string{
//   				jsii.String("entryPoint"),
//   			},
//   			Environment: []interface{}{
//   				&KeyValuePairProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			EnvironmentFiles: []interface{}{
//   				&EnvironmentFileProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Essential: jsii.Boolean(false),
//   			FirelensConfiguration: &FirelensConfigurationProperty{
//   				Options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			HealthCheck: &HealthCheckProperty{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Interval: jsii.Number(123),
//   				Retries: jsii.Number(123),
//   				StartPeriod: jsii.Number(123),
//   				Timeout: jsii.Number(123),
//   			},
//   			Interactive: jsii.Boolean(false),
//   			LinuxParameters: &LinuxParametersProperty{
//   				Capabilities: &KernelCapabilitiesProperty{
//   					Add: []*string{
//   						jsii.String("add"),
//   					},
//   					Drop: []*string{
//   						jsii.String("drop"),
//   					},
//   				},
//   				Devices: []interface{}{
//   					&DeviceProperty{
//   						ContainerPath: jsii.String("containerPath"),
//   						HostPath: jsii.String("hostPath"),
//   						Permissions: []*string{
//   							jsii.String("permissions"),
//   						},
//   					},
//   				},
//   				InitProcessEnabled: jsii.Boolean(false),
//   				Tmpfs: []interface{}{
//   					&TmpfsProperty{
//   						Size: jsii.Number(123),
//
//   						// the properties below are optional
//   						ContainerPath: jsii.String("containerPath"),
//   						MountOptions: []*string{
//   							jsii.String("mountOptions"),
//   						},
//   					},
//   				},
//   			},
//   			LogConfiguration: &LogConfigurationProperty{
//   				LogDriver: jsii.String("logDriver"),
//
//   				// the properties below are optional
//   				Options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   				SecretOptions: []interface{}{
//   					&SecretProperty{
//   						Name: jsii.String("name"),
//   						ValueFrom: jsii.String("valueFrom"),
//   					},
//   				},
//   			},
//   			Memory: jsii.Number(123),
//   			MemoryReservation: jsii.Number(123),
//   			MountPoints: []interface{}{
//   				&MountPointProperty{
//   					ContainerPath: jsii.String("containerPath"),
//   					ReadOnly: jsii.Boolean(false),
//   					SourceVolume: jsii.String("sourceVolume"),
//   				},
//   			},
//   			Privileged: jsii.Boolean(false),
//   			PseudoTerminal: jsii.Boolean(false),
//   			ReadonlyRootFilesystem: jsii.Boolean(false),
//   			RepositoryCredentials: &RepositoryCredentialsProperty{
//   				CredentialsParameter: jsii.String("credentialsParameter"),
//   			},
//   			RestartPolicy: &RestartPolicyProperty{
//   				Enabled: jsii.Boolean(false),
//   				IgnoredExitCodes: []interface{}{
//   					jsii.Number(123),
//   				},
//   				RestartAttemptPeriod: jsii.Number(123),
//   			},
//   			Secrets: []interface{}{
//   				&SecretProperty{
//   					Name: jsii.String("name"),
//   					ValueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   			StartTimeout: jsii.Number(123),
//   			StopTimeout: jsii.Number(123),
//   			SystemControls: []interface{}{
//   				&SystemControlProperty{
//   					Namespace: jsii.String("namespace"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Ulimits: []interface{}{
//   				&UlimitProperty{
//   					HardLimit: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					SoftLimit: jsii.Number(123),
//   				},
//   			},
//   			User: jsii.String("user"),
//   			WorkingDirectory: jsii.String("workingDirectory"),
//   		},
//   	},
//   	Cpu: jsii.String("cpu"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Family: jsii.String("family"),
//   	IpcMode: jsii.String("ipcMode"),
//   	Memory: jsii.String("memory"),
//   	PidMode: jsii.String("pidMode"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskRoleArn: jsii.String("taskRoleArn"),
//   	Volumes: []interface{}{
//   		&VolumeProperty{
//   			Host: &HostVolumePropertiesProperty{
//   				SourcePath: jsii.String("sourcePath"),
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html
//
type CfnDaemonTaskDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsecs.IDaemonTaskDefinitionRef
	awscdk.ITaggableV2
	AttrDaemonTaskDefinitionArn() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnPropertyNames() *map[string]*string
	// AWS resource type.
	CfnResourceType() *string
	// A list of container definitions in JSON format that describe the containers that make up the daemon task.
	ContainerDefinitions() interface{}
	SetContainerDefinitions(val interface{})
	// The number of CPU units used by the daemon task.
	Cpu() *string
	SetCpu(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A reference to a DaemonTaskDefinition resource.
	DaemonTaskDefinitionRef() *interfacesawsecs.DaemonTaskDefinitionReference
	Env() *interfaces.ResourceEnvironment
	// The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make Amazon Web Services API calls on your behalf.
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// The name of a family that this daemon task definition is registered to.
	Family() *string
	SetFamily(val *string)
	// The IPC namespace mode for the daemon.
	IpcMode() *string
	SetIpcMode(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The amount of memory (in MiB) used by the daemon task.
	Memory() *string
	SetMemory(val *string)
	// The tree node.
	Node() constructs.Node
	// The PID namespace mode for the daemon.
	PidMode() *string
	SetPidMode(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// The short name or full Amazon Resource Name (ARN) of the IAM role that grants containers in the daemon task permission to call Amazon Web Services APIs on your behalf.
	TaskRoleArn() *string
	SetTaskRoleArn(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The list of data volume definitions for the daemon task.
	Volumes() interface{}
	SetVolumes(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the cross-stack reference strength for this resource.
	//
	// When set, any cross-stack reference to this resource will use the specified
	// strength instead of the global default from the consuming stack's context.
	ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength)
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	CfnPropertyName(cdkPropertyName *string) *string
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnDaemonTaskDefinition
type jsiiProxy_CfnDaemonTaskDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsecsIDaemonTaskDefinitionRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) AttrDaemonTaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDaemonTaskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) ContainerDefinitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) Cpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) DaemonTaskDefinitionRef() *interfacesawsecs.DaemonTaskDefinitionReference {
	var returns *interfacesawsecs.DaemonTaskDefinitionReference
	_jsii_.Get(
		j,
		"daemonTaskDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) IpcMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipcMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) Memory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) PidMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) TaskRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinition) Volumes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECS::DaemonTaskDefinition`.
func NewCfnDaemonTaskDefinition(scope constructs.Construct, id *string, props *CfnDaemonTaskDefinitionProps) CfnDaemonTaskDefinition {
	_init_.Initialize()

	if err := validateNewCfnDaemonTaskDefinitionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDaemonTaskDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnDaemonTaskDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECS::DaemonTaskDefinition`.
func NewCfnDaemonTaskDefinition_Override(c CfnDaemonTaskDefinition, scope constructs.Construct, id *string, props *CfnDaemonTaskDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CfnDaemonTaskDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDaemonTaskDefinition)SetContainerDefinitions(val interface{}) {
	if err := j.validateSetContainerDefinitionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnDaemonTaskDefinition)SetCpu(val *string) {
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_CfnDaemonTaskDefinition)SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDaemonTaskDefinition)SetFamily(val *string) {
	_jsii_.Set(
		j,
		"family",
		val,
	)
}

func (j *jsiiProxy_CfnDaemonTaskDefinition)SetIpcMode(val *string) {
	_jsii_.Set(
		j,
		"ipcMode",
		val,
	)
}

func (j *jsiiProxy_CfnDaemonTaskDefinition)SetMemory(val *string) {
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_CfnDaemonTaskDefinition)SetPidMode(val *string) {
	_jsii_.Set(
		j,
		"pidMode",
		val,
	)
}

func (j *jsiiProxy_CfnDaemonTaskDefinition)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnDaemonTaskDefinition)SetTaskRoleArn(val *string) {
	_jsii_.Set(
		j,
		"taskRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDaemonTaskDefinition)SetVolumes(val interface{}) {
	if err := j.validateSetVolumesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumes",
		val,
	)
}

func CfnDaemonTaskDefinition_ArnForDaemonTaskDefinition(resource interfacesawsecs.IDaemonTaskDefinitionRef) *string {
	_init_.Initialize()

	if err := validateCfnDaemonTaskDefinition_ArnForDaemonTaskDefinitionParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnDaemonTaskDefinition",
		"arnForDaemonTaskDefinition",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnDaemonTaskDefinition.
func CfnDaemonTaskDefinition_IsCfnDaemonTaskDefinition(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDaemonTaskDefinition_IsCfnDaemonTaskDefinitionParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnDaemonTaskDefinition",
		"isCfnDaemonTaskDefinition",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDaemonTaskDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDaemonTaskDefinition_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnDaemonTaskDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDaemonTaskDefinition_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDaemonTaskDefinition_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnDaemonTaskDefinition",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
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
func CfnDaemonTaskDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDaemonTaskDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CfnDaemonTaskDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDaemonTaskDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.CfnDaemonTaskDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) CfnPropertyName(cdkPropertyName *string) *string {
	if err := c.validateCfnPropertyNameParameters(cdkPropertyName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"cfnPropertyName",
		[]interface{}{cdkPropertyName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinition) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

