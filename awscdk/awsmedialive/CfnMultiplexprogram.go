package awsmedialive

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::MediaLive::Multiplexprogram.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMultiplexprogram := awscdk.Aws_medialive.NewCfnMultiplexprogram(this, jsii.String("MyCfnMultiplexprogram"), &CfnMultiplexprogramProps{
//   	MultiplexId: jsii.String("multiplexId"),
//   	MultiplexProgramSettings: &MultiplexProgramSettingsProperty{
//   		ProgramNumber: jsii.Number(123),
//
//   		// the properties below are optional
//   		PreferredChannelPipeline: jsii.String("preferredChannelPipeline"),
//   		ServiceDescriptor: &MultiplexProgramServiceDescriptorProperty{
//   			ProviderName: jsii.String("providerName"),
//   			ServiceName: jsii.String("serviceName"),
//   		},
//   		VideoSettings: &MultiplexVideoSettingsProperty{
//   			ConstantBitrate: jsii.Number(123),
//   			StatmuxSettings: &MultiplexStatmuxVideoSettingsProperty{
//   				MaximumBitrate: jsii.Number(123),
//   				MinimumBitrate: jsii.Number(123),
//   				Priority: jsii.Number(123),
//   			},
//   		},
//   	},
//   	PacketIdentifiersMap: &MultiplexProgramPacketIdentifiersMapProperty{
//   		AudioPids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		DvbSubPids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		DvbTeletextPid: jsii.Number(123),
//   		EtvPlatformPid: jsii.Number(123),
//   		EtvSignalPid: jsii.Number(123),
//   		KlvDataPids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		PcrPid: jsii.Number(123),
//   		PmtPid: jsii.Number(123),
//   		PrivateMetadataPid: jsii.Number(123),
//   		Scte27Pids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		Scte35Pid: jsii.Number(123),
//   		TimedMetadataPid: jsii.Number(123),
//   		VideoPid: jsii.Number(123),
//   	},
//   	PipelineDetails: []interface{}{
//   		&MultiplexProgramPipelineDetailProperty{
//   			ActiveChannelPipeline: jsii.String("activeChannelPipeline"),
//   			PipelineId: jsii.String("pipelineId"),
//   		},
//   	},
//   	PreferredChannelPipeline: jsii.String("preferredChannelPipeline"),
//   	ProgramName: jsii.String("programName"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplexprogram.html
//
type CfnMultiplexprogram interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsmedialive.IMultiplexprogramRef
	// The unique ID of the channel.
	AttrChannelId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	Env() *interfaces.ResourceEnvironment
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
	// The unique id of the multiplex.
	MultiplexId() *string
	SetMultiplexId(val *string)
	// A reference to a Multiplexprogram resource.
	MultiplexprogramRef() *interfacesawsmedialive.MultiplexprogramReference
	// Multiplex Program settings configuration.
	MultiplexProgramSettings() interface{}
	SetMultiplexProgramSettings(val interface{})
	// The tree node.
	Node() constructs.Node
	// Packet identifiers map for a given Multiplex program.
	PacketIdentifiersMap() interface{}
	SetPacketIdentifiersMap(val interface{})
	// Contains information about the current sources for the specified program in the specified multiplex.
	PipelineDetails() interface{}
	SetPipelineDetails(val interface{})
	// Indicates which pipeline is preferred by the multiplex for program ingest.
	PreferredChannelPipeline() *string
	SetPreferredChannelPipeline(val *string)
	// The name of the multiplex program.
	ProgramName() *string
	SetProgramName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnMultiplexprogram
type jsiiProxy_CfnMultiplexprogram struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsmedialiveIMultiplexprogramRef
}

func (j *jsiiProxy_CfnMultiplexprogram) AttrChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrChannelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) MultiplexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiplexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) MultiplexprogramRef() *interfacesawsmedialive.MultiplexprogramReference {
	var returns *interfacesawsmedialive.MultiplexprogramReference
	_jsii_.Get(
		j,
		"multiplexprogramRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) MultiplexProgramSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiplexProgramSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) PacketIdentifiersMap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"packetIdentifiersMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) PipelineDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) PreferredChannelPipeline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChannelPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) ProgramName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiplexprogram) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaLive::Multiplexprogram`.
func NewCfnMultiplexprogram(scope constructs.Construct, id *string, props *CfnMultiplexprogramProps) CfnMultiplexprogram {
	_init_.Initialize()

	if err := validateNewCfnMultiplexprogramParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMultiplexprogram{}

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaLive::Multiplexprogram`.
func NewCfnMultiplexprogram_Override(c CfnMultiplexprogram, scope constructs.Construct, id *string, props *CfnMultiplexprogramProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMultiplexprogram)SetMultiplexId(val *string) {
	_jsii_.Set(
		j,
		"multiplexId",
		val,
	)
}

func (j *jsiiProxy_CfnMultiplexprogram)SetMultiplexProgramSettings(val interface{}) {
	if err := j.validateSetMultiplexProgramSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiplexProgramSettings",
		val,
	)
}

func (j *jsiiProxy_CfnMultiplexprogram)SetPacketIdentifiersMap(val interface{}) {
	if err := j.validateSetPacketIdentifiersMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packetIdentifiersMap",
		val,
	)
}

func (j *jsiiProxy_CfnMultiplexprogram)SetPipelineDetails(val interface{}) {
	if err := j.validateSetPipelineDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineDetails",
		val,
	)
}

func (j *jsiiProxy_CfnMultiplexprogram)SetPreferredChannelPipeline(val *string) {
	_jsii_.Set(
		j,
		"preferredChannelPipeline",
		val,
	)
}

func (j *jsiiProxy_CfnMultiplexprogram)SetProgramName(val *string) {
	_jsii_.Set(
		j,
		"programName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMultiplexprogram_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMultiplexprogram_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnMultiplexprogram.
func CfnMultiplexprogram_IsCfnMultiplexprogram(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMultiplexprogram_IsCfnMultiplexprogramParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram",
		"isCfnMultiplexprogram",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnMultiplexprogram_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMultiplexprogram_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram",
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
func CfnMultiplexprogram_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMultiplexprogram_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMultiplexprogram_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMultiplexprogram) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnMultiplexprogram) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnMultiplexprogram) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiplexprogram) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiplexprogram) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnMultiplexprogram) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiplexprogram) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiplexprogram) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnMultiplexprogram) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

