package awsgamelift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::GameLift::Alias`.
//
// The `AWS::GameLift::Alias` resource creates an alias for an Amazon GameLift (GameLift) fleet destination. There are two types of routing strategies for aliases: simple and terminal. A simple alias points to an active fleet. A terminal alias displays a message instead of routing players to an active fleet. For example, a terminal alias might display a URL link that directs players to an upgrade site. You can use aliases to define destinations in a game session queue or when requesting new game sessions.
//
// TODO: EXAMPLE
//
type CfnAlias interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAliasId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	RoutingStrategy() interface{}
	SetRoutingStrategy(val interface{})
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnAlias
type jsiiProxy_CfnAlias struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAlias) AttrAliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAliasId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) RoutingStrategy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::GameLift::Alias`.
func NewCfnAlias(scope awscdk.Construct, id *string, props *CfnAliasProps) CfnAlias {
	_init_.Initialize()

	j := jsiiProxy_CfnAlias{}

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnAlias",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::Alias`.
func NewCfnAlias_Override(c CfnAlias, scope awscdk.Construct, id *string, props *CfnAliasProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnAlias",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAlias) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnAlias) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnAlias) SetRoutingStrategy(val interface{}) {
	_jsii_.Set(
		j,
		"routingStrategy",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnAlias_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnAlias",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAlias_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnAlias",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAlias_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_gamelift.CfnAlias",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnAlias) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnAlias) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnAlias) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnAlias) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnAlias) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnAlias) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CfnAlias) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnAlias) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnAlias) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnAlias) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnAlias) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnAlias) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnAlias) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnAlias) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnAlias) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAlias) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnAlias) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnAlias) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnAlias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnAlias) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnAlias) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The routing configuration for a fleet alias.
//
// TODO: EXAMPLE
//
type CfnAlias_RoutingStrategyProperty struct {
	// A type of routing strategy.
	//
	// Possible routing types include the following:
	//
	// - *SIMPLE* - The alias resolves to one specific fleet. Use this type when routing to active fleets.
	// - *TERMINAL* - The alias does not resolve to a fleet but instead can be used to display a message to the user. A terminal alias throws a `TerminalRoutingStrategyException` with the message that you specified in the `Message` property.
	Type *string `json:"type" yaml:"type"`
	// A unique identifier for a fleet that the alias points to.
	//
	// If you specify `SIMPLE` for the `Type` property, you must specify this property.
	FleetId *string `json:"fleetId" yaml:"fleetId"`
	// The message text to be used with a terminal routing strategy.
	//
	// If you specify `TERMINAL` for the `Type` property, you must specify this property.
	Message *string `json:"message" yaml:"message"`
}

// Properties for defining a `CfnAlias`.
//
// TODO: EXAMPLE
//
type CfnAliasProps struct {
	// A descriptive label that is associated with an alias.
	//
	// Alias names do not need to be unique.
	Name *string `json:"name" yaml:"name"`
	// The routing configuration, including routing type and fleet target, for the alias.
	RoutingStrategy interface{} `json:"routingStrategy" yaml:"routingStrategy"`
	// A human-readable description of the alias.
	Description *string `json:"description" yaml:"description"`
}

// A CloudFormation `AWS::GameLift::Build`.
//
// The `AWS::GameLift::Build` resource creates a game server build that is installed and run on instances in an Amazon GameLift fleet. This resource points to an Amazon S3 location that contains a zip file with all of the components of the game server build.
//
// TODO: EXAMPLE
//
type CfnBuild interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	OperatingSystem() *string
	SetOperatingSystem(val *string)
	Ref() *string
	Stack() awscdk.Stack
	StorageLocation() interface{}
	SetStorageLocation(val interface{})
	UpdatedProperites() *map[string]interface{}
	Version() *string
	SetVersion(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnBuild
type jsiiProxy_CfnBuild struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBuild) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBuild) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBuild) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBuild) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBuild) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBuild) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBuild) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBuild) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBuild) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBuild) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBuild) StorageLocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBuild) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBuild) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new `AWS::GameLift::Build`.
func NewCfnBuild(scope awscdk.Construct, id *string, props *CfnBuildProps) CfnBuild {
	_init_.Initialize()

	j := jsiiProxy_CfnBuild{}

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnBuild",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::Build`.
func NewCfnBuild_Override(c CfnBuild, scope awscdk.Construct, id *string, props *CfnBuildProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnBuild",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBuild) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnBuild) SetOperatingSystem(val *string) {
	_jsii_.Set(
		j,
		"operatingSystem",
		val,
	)
}

func (j *jsiiProxy_CfnBuild) SetStorageLocation(val interface{}) {
	_jsii_.Set(
		j,
		"storageLocation",
		val,
	)
}

func (j *jsiiProxy_CfnBuild) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnBuild_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnBuild",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnBuild_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnBuild",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnBuild_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnBuild",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBuild_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_gamelift.CfnBuild",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnBuild) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnBuild) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnBuild) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnBuild) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnBuild) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnBuild) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CfnBuild) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnBuild) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnBuild) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnBuild) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnBuild) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnBuild) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnBuild) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnBuild) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnBuild) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBuild) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnBuild) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnBuild) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnBuild) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnBuild) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnBuild) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The location in Amazon S3 where build or script files are stored for access by Amazon GameLift.
//
// TODO: EXAMPLE
//
type CfnBuild_S3LocationProperty struct {
	// An Amazon S3 bucket identifier. This is the name of the S3 bucket.
	//
	// > GameLift currently does not support uploading from Amazon S3 buckets with names that contain a dot (.).
	Bucket *string `json:"bucket" yaml:"bucket"`
	// The name of the zip file that contains the build files or script files.
	Key *string `json:"key" yaml:"key"`
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) for an IAM role that allows Amazon Web Services to access the S3 bucket.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The version of the file, if object versioning is turned on for the bucket.
	//
	// Amazon GameLift uses this information when retrieving files from your S3 bucket. To retrieve a specific version of the file, provide an object version. To retrieve the latest version of the file, do not set this parameter.
	ObjectVersion *string `json:"objectVersion" yaml:"objectVersion"`
}

// Properties for defining a `CfnBuild`.
//
// TODO: EXAMPLE
//
type CfnBuildProps struct {
	// A descriptive label that is associated with a build.
	//
	// Build names do not need to be unique.
	Name *string `json:"name" yaml:"name"`
	// The operating system that the game server binaries are built to run on.
	//
	// This value determines the type of fleet resources that you can use for this build. If your game build contains multiple executables, they all must run on the same operating system. If an operating system is not specified when creating a build, Amazon GameLift uses the default value (WINDOWS_2012). This value cannot be changed later.
	OperatingSystem *string `json:"operatingSystem" yaml:"operatingSystem"`
	// Information indicating where your game build files are stored.
	//
	// Use this parameter only when creating a build with files stored in an Amazon S3 bucket that you own. The storage location must specify an Amazon S3 bucket name and key. The location must also specify a role ARN that you set up to allow Amazon Web Services to access your Amazon S3 bucket. The S3 bucket and your new build must be in the same Region.
	//
	// If a `StorageLocation` is specified, the size of your file can be found in your Amazon S3 bucket. Amazon Web Services will report a `SizeOnDisk` of 0.
	StorageLocation interface{} `json:"storageLocation" yaml:"storageLocation"`
	// Version information that is associated with this build.
	//
	// Version strings do not need to be unique.
	Version *string `json:"version" yaml:"version"`
}

// A CloudFormation `AWS::GameLift::Fleet`.
//
// The `AWS::GameLift::Fleet` resource creates an Amazon GameLift (GameLift) fleet to host game servers. A fleet is a set of EC2 instances, each of which can host multiple game sessions.
//
// TODO: EXAMPLE
//
type CfnFleet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrFleetId() *string
	BuildId() *string
	SetBuildId(val *string)
	CertificateConfiguration() interface{}
	SetCertificateConfiguration(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DesiredEc2Instances() *float64
	SetDesiredEc2Instances(val *float64)
	Ec2InboundPermissions() interface{}
	SetEc2InboundPermissions(val interface{})
	Ec2InstanceType() *string
	SetEc2InstanceType(val *string)
	FleetType() *string
	SetFleetType(val *string)
	InstanceRoleArn() *string
	SetInstanceRoleArn(val *string)
	Locations() interface{}
	SetLocations(val interface{})
	LogicalId() *string
	MaxSize() *float64
	SetMaxSize(val *float64)
	MetricGroups() *[]*string
	SetMetricGroups(val *[]*string)
	MinSize() *float64
	SetMinSize(val *float64)
	Name() *string
	SetName(val *string)
	NewGameSessionProtectionPolicy() *string
	SetNewGameSessionProtectionPolicy(val *string)
	Node() awscdk.ConstructNode
	PeerVpcAwsAccountId() *string
	SetPeerVpcAwsAccountId(val *string)
	PeerVpcId() *string
	SetPeerVpcId(val *string)
	Ref() *string
	ResourceCreationLimitPolicy() interface{}
	SetResourceCreationLimitPolicy(val interface{})
	RuntimeConfiguration() interface{}
	SetRuntimeConfiguration(val interface{})
	ScriptId() *string
	SetScriptId(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnFleet
type jsiiProxy_CfnFleet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFleet) AttrFleetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFleetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) BuildId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) CertificateConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) DesiredEc2Instances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredEc2Instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Ec2InboundPermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InboundPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Ec2InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) FleetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) InstanceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Locations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) MetricGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) NewGameSessionProtectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) PeerVpcAwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVpcAwsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) PeerVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) ResourceCreationLimitPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceCreationLimitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) RuntimeConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) ScriptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::GameLift::Fleet`.
func NewCfnFleet(scope awscdk.Construct, id *string, props *CfnFleetProps) CfnFleet {
	_init_.Initialize()

	j := jsiiProxy_CfnFleet{}

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnFleet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::Fleet`.
func NewCfnFleet_Override(c CfnFleet, scope awscdk.Construct, id *string, props *CfnFleetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnFleet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFleet) SetBuildId(val *string) {
	_jsii_.Set(
		j,
		"buildId",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetCertificateConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"certificateConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetDesiredEc2Instances(val *float64) {
	_jsii_.Set(
		j,
		"desiredEc2Instances",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetEc2InboundPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"ec2InboundPermissions",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetEc2InstanceType(val *string) {
	_jsii_.Set(
		j,
		"ec2InstanceType",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetFleetType(val *string) {
	_jsii_.Set(
		j,
		"fleetType",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetInstanceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"instanceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetLocations(val interface{}) {
	_jsii_.Set(
		j,
		"locations",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetMaxSize(val *float64) {
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetMetricGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"metricGroups",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetMinSize(val *float64) {
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetNewGameSessionProtectionPolicy(val *string) {
	_jsii_.Set(
		j,
		"newGameSessionProtectionPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetPeerVpcAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"peerVpcAwsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetPeerVpcId(val *string) {
	_jsii_.Set(
		j,
		"peerVpcId",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetResourceCreationLimitPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"resourceCreationLimitPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetRuntimeConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"runtimeConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetScriptId(val *string) {
	_jsii_.Set(
		j,
		"scriptId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnFleet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnFleet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFleet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnFleet",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFleet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_gamelift.CfnFleet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnFleet) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnFleet) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnFleet) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnFleet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnFleet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnFleet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CfnFleet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnFleet) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnFleet) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnFleet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnFleet) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnFleet) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnFleet) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnFleet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnFleet) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFleet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnFleet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnFleet) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnFleet) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnFleet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Determines whether a TLS/SSL certificate is generated for a fleet.
//
// This feature must be enabled when creating the fleet. All instances in a fleet share the same certificate. The certificate can be retrieved by calling the [GameLift Server SDK](https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk.html) operation `GetInstanceCertificate` .
//
// TODO: EXAMPLE
//
type CfnFleet_CertificateConfigurationProperty struct {
	// Indicates whether a TLS/SSL certificate is generated for a fleet.
	//
	// Valid values include:
	//
	// - *GENERATED* - Generate a TLS/SSL certificate for this fleet.
	// - *DISABLED* - (default) Do not generate a TLS/SSL certificate for this fleet.
	CertificateType *string `json:"certificateType" yaml:"certificateType"`
}

// A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an instance in a fleet.
//
// New game sessions are assigned an IP address/port number combination, which must fall into the fleet's allowed ranges. Fleets with custom game builds must have permissions explicitly set. For Realtime Servers fleets, GameLift automatically opens two port ranges, one for TCP messaging and one for UDP.
//
// TODO: EXAMPLE
//
type CfnFleet_IpPermissionProperty struct {
	// A starting value for a range of allowed port numbers.
	//
	// For fleets using Windows and Linux builds, only ports 1026-60000 are valid.
	FromPort *float64 `json:"fromPort" yaml:"fromPort"`
	// A range of allowed IP addresses.
	//
	// This value must be expressed in CIDR notation. Example: " `000.000.000.000/[subnet mask]` " or optionally the shortened version " `0.0.0.0/[subnet mask]` ".
	IpRange *string `json:"ipRange" yaml:"ipRange"`
	// The network communication protocol used by the fleet.
	Protocol *string `json:"protocol" yaml:"protocol"`
	// An ending value for a range of allowed port numbers.
	//
	// Port numbers are end-inclusive. This value must be higher than `FromPort` .
	//
	// For fleets using Windows and Linux builds, only ports 1026-60000 are valid.
	ToPort *float64 `json:"toPort" yaml:"toPort"`
}

// Current resource capacity settings in a specified fleet or location.
//
// The location value might refer to a fleet's remote location or its home Region.
//
// *Related actions*
//
// [DescribeFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetCapacity.html) | [DescribeFleetLocationCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html) | [UpdateFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateFleetCapacity.html)
//
// TODO: EXAMPLE
//
type CfnFleet_LocationCapacityProperty struct {
	// The number of Amazon EC2 instances you want to maintain in the specified fleet location.
	//
	// This value must fall between the minimum and maximum size limits.
	DesiredEc2Instances *float64 `json:"desiredEc2Instances" yaml:"desiredEc2Instances"`
	// The maximum number of instances that are allowed in the specified fleet location.
	//
	// If this parameter is not set, the default is 1.
	MaxSize *float64 `json:"maxSize" yaml:"maxSize"`
	// The minimum number of instances that are allowed in the specified fleet location.
	//
	// If this parameter is not set, the default is 0.
	MinSize *float64 `json:"minSize" yaml:"minSize"`
}

// A remote location where a multi-location fleet can deploy EC2 instances for game hosting.
//
// *Related actions*
//
// [CreateFleet](https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateFleet.html)
//
// TODO: EXAMPLE
//
type CfnFleet_LocationConfigurationProperty struct {
	// An AWS Region code, such as `us-west-2` .
	Location *string `json:"location" yaml:"location"`
	// Current resource capacity settings in a specified fleet or location.
	//
	// The location value might refer to a fleet's remote location or its home Region.
	//
	// *Related actions*
	//
	// [DescribeFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetCapacity.html) | [DescribeFleetLocationCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html) | [UpdateFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateFleetCapacity.html)
	LocationCapacity interface{} `json:"locationCapacity" yaml:"locationCapacity"`
}

// A policy that limits the number of game sessions a player can create on the same fleet.
//
// This optional policy gives game owners control over how players can consume available game server resources. A resource creation policy makes the following statement: "An individual player can create a maximum number of new game sessions within a specified time period".
//
// The policy is evaluated when a player tries to create a new game session. For example, assume you have a policy of 10 new game sessions and a time period of 60 minutes. On receiving a `CreateGameSession` request, Amazon GameLift checks that the player (identified by `CreatorId` ) has created fewer than 10 game sessions in the past 60 minutes.
//
// TODO: EXAMPLE
//
type CfnFleet_ResourceCreationLimitPolicyProperty struct {
	// The maximum number of game sessions that an individual can create during the policy period.
	NewGameSessionsPerCreator *float64 `json:"newGameSessionsPerCreator" yaml:"newGameSessionsPerCreator"`
	// The time span used in evaluating the resource creation limit policy.
	PolicyPeriodInMinutes *float64 `json:"policyPeriodInMinutes" yaml:"policyPeriodInMinutes"`
}

// A collection of server process configurations that describe the set of processes to run on each instance in a fleet.
//
// Server processes run either an executable in a custom game build or a Realtime Servers script. GameLift launches the configured processes, manages their life cycle, and replaces them as needed. Each instance checks regularly for an updated runtime configuration.
//
// A GameLift instance is limited to 50 processes running concurrently. To calculate the total number of processes in a runtime configuration, add the values of the `ConcurrentExecutions` parameter for each ServerProcess. Learn more about [Running Multiple Processes on a Fleet](https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-multiprocess.html) .
//
// TODO: EXAMPLE
//
type CfnFleet_RuntimeConfigurationProperty struct {
	// The maximum amount of time (in seconds) allowed to launch a new game session and have it report ready to host players.
	//
	// During this time, the game session is in status `ACTIVATING` . If the game session does not become active before the timeout, it is ended and the game session status is changed to `TERMINATED` .
	GameSessionActivationTimeoutSeconds *float64 `json:"gameSessionActivationTimeoutSeconds" yaml:"gameSessionActivationTimeoutSeconds"`
	// The number of game sessions in status `ACTIVATING` to allow on an instance.
	//
	// This setting limits the instance resources that can be used for new game activations at any one time.
	MaxConcurrentGameSessionActivations *float64 `json:"maxConcurrentGameSessionActivations" yaml:"maxConcurrentGameSessionActivations"`
	// A collection of server process configurations that identify what server processes to run on each instance in a fleet.
	ServerProcesses interface{} `json:"serverProcesses" yaml:"serverProcesses"`
}

// A set of instructions for launching server processes on each instance in a fleet.
//
// Server processes run either an executable in a custom game build or a Realtime Servers script.
//
// TODO: EXAMPLE
//
type CfnFleet_ServerProcessProperty struct {
	// The number of server processes using this configuration that run concurrently on each instance.
	ConcurrentExecutions *float64 `json:"concurrentExecutions" yaml:"concurrentExecutions"`
	// The location of a game build executable or the Realtime script file that contains the `Init()` function.
	//
	// Game builds and Realtime scripts are installed on instances at the root:
	//
	// - Windows (custom game builds only): `C:\game` . Example: " `C:\game\MyGame\server.exe` "
	// - Linux: `/local/game` . Examples: " `/local/game/MyGame/server.exe` " or " `/local/game/MyRealtimeScript.js` "
	LaunchPath *string `json:"launchPath" yaml:"launchPath"`
	// An optional list of parameters to pass to the server executable or Realtime script on launch.
	Parameters *string `json:"parameters" yaml:"parameters"`
}

// Properties for defining a `CfnFleet`.
//
// TODO: EXAMPLE
//
type CfnFleetProps struct {
	// A unique identifier for a build to be deployed on the new fleet.
	//
	// If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a `READY` status. This fleet setting cannot be changed once the fleet is created.
	BuildId *string `json:"buildId" yaml:"buildId"`
	// Prompts GameLift to generate a TLS/SSL certificate for the fleet.
	//
	// TLS certificates are used for encrypting traffic between game clients and the game servers that are running on GameLift. By default, the `CertificateConfiguration` is set to `DISABLED` . This property cannot be changed after the fleet is created.
	//
	// Note: This feature requires the AWS Certificate Manager (ACM) service, which is not available in all AWS regions. When working in a region that does not support this feature, a fleet creation request with certificate generation fails with a 4xx error.
	CertificateConfiguration interface{} `json:"certificateConfiguration" yaml:"certificateConfiguration"`
	// A human-readable description of the fleet.
	Description *string `json:"description" yaml:"description"`
	// The number of EC2 instances that you want this fleet to host.
	//
	// When creating a new fleet, GameLift automatically sets this value to "1" and initiates a single instance. Once the fleet is active, update this value to trigger GameLift to add or remove instances from the fleet.
	DesiredEc2Instances *float64 `json:"desiredEc2Instances" yaml:"desiredEc2Instances"`
	// The allowed IP address ranges and port settings that allow inbound traffic to access game sessions on this fleet.
	//
	// If the fleet is hosting a custom game build, this property must be set before players can connect to game sessions. For Realtime Servers fleets, GameLift automatically sets TCP and UDP ranges.
	Ec2InboundPermissions interface{} `json:"ec2InboundPermissions" yaml:"ec2InboundPermissions"`
	// The GameLift-supported Amazon EC2 instance type to use for all fleet instances.
	//
	// Instance type determines the computing resources that will be used to host your game servers, including CPU, memory, storage, and networking capacity. See [Amazon Elastic Compute Cloud Instance Types](https://docs.aws.amazon.com/ec2/instance-types/) for detailed descriptions of Amazon EC2 instance types.
	Ec2InstanceType *string `json:"ec2InstanceType" yaml:"ec2InstanceType"`
	// Indicates whether to use On-Demand or Spot instances for this fleet.
	//
	// By default, this property is set to `ON_DEMAND` . Learn more about when to use [On-Demand versus Spot Instances](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-ec2-instances.html#gamelift-ec2-instances-spot) . This property cannot be changed after the fleet is created.
	FleetType *string `json:"fleetType" yaml:"fleetType"`
	// A unique identifier for an IAM role that manages access to your AWS services.
	//
	// With an instance role ARN set, any application that runs on an instance in this fleet can assume the role, including install scripts, server processes, and daemons (background processes). Create a role or look up a role's ARN by using the [IAM dashboard](https://docs.aws.amazon.com/iam/) in the AWS Management Console . Learn more about using on-box credentials for your game servers at [Access external resources from a game server](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html) . This property cannot be changed after the fleet is created.
	InstanceRoleArn *string `json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// A set of remote locations to deploy additional instances to and manage as part of the fleet.
	//
	// This parameter can only be used when creating fleets in AWS Regions that support multiple locations. You can add any GameLift-supported AWS Region as a remote location, in the form of an AWS Region code such as `us-west-2` . To create a fleet with instances in the home Region only, omit this parameter.
	Locations interface{} `json:"locations" yaml:"locations"`
	// The maximum number of instances that are allowed in the specified fleet location.
	//
	// If this parameter is not set, the default is 1.
	MaxSize *float64 `json:"maxSize" yaml:"maxSize"`
	// The name of an AWS CloudWatch metric group to add this fleet to.
	//
	// A metric group is used to aggregate the metrics for multiple fleets. You can specify an existing metric group name or set a new name to create a new metric group. A fleet can be included in only one metric group at a time.
	MetricGroups *[]*string `json:"metricGroups" yaml:"metricGroups"`
	// The minimum number of instances that are allowed in the specified fleet location.
	//
	// If this parameter is not set, the default is 0.
	MinSize *float64 `json:"minSize" yaml:"minSize"`
	// A descriptive label that is associated with a fleet.
	//
	// Fleet names do not need to be unique.
	Name *string `json:"name" yaml:"name"`
	// The status of termination protection for active game sessions on the fleet.
	//
	// By default, this property is set to `NoProtection` .
	//
	// - *NoProtection* - Game sessions can be terminated during active gameplay as a result of a scale-down event.
	// - *FullProtection* - Game sessions in `ACTIVE` status cannot be terminated during a scale-down event.
	NewGameSessionProtectionPolicy *string `json:"newGameSessionProtectionPolicy" yaml:"newGameSessionProtectionPolicy"`
	// Used when peering your GameLift fleet with a VPC, the unique identifier for the AWS account that owns the VPC.
	//
	// You can find your account ID in the AWS Management Console under account settings.
	PeerVpcAwsAccountId *string `json:"peerVpcAwsAccountId" yaml:"peerVpcAwsAccountId"`
	// A unique identifier for a VPC with resources to be accessed by your GameLift fleet.
	//
	// The VPC must be in the same Region as your fleet. To look up a VPC ID, use the [VPC Dashboard](https://docs.aws.amazon.com/vpc/) in the AWS Management Console . Learn more about VPC peering in [VPC Peering with GameLift Fleets](https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html) .
	PeerVpcId *string `json:"peerVpcId" yaml:"peerVpcId"`
	// A policy that limits the number of game sessions that an individual player can create on instances in this fleet within a specified span of time.
	ResourceCreationLimitPolicy interface{} `json:"resourceCreationLimitPolicy" yaml:"resourceCreationLimitPolicy"`
	// Instructions for how to launch and maintain server processes on instances in the fleet.
	//
	// The runtime configuration defines one or more server process configurations, each identifying a build executable or Realtime script file and the number of processes of that type to run concurrently.
	//
	// > The `RuntimeConfiguration` parameter is required unless the fleet is being configured using the older parameters `ServerLaunchPath` and `ServerLaunchParameters` , which are still supported for backward compatibility.
	RuntimeConfiguration interface{} `json:"runtimeConfiguration" yaml:"runtimeConfiguration"`
	// The unique identifier for a Realtime configuration script to be deployed on fleet instances.
	//
	// You can use either the script ID or ARN. Scripts must be uploaded to GameLift prior to creating the fleet. This fleet property cannot be changed later.
	//
	// > You can't use the `!Ref` command to reference a script created with a CloudFormation template for the fleet property `ScriptId` . Instead, use `Fn::GetAtt Script.Arn` or `Fn::GetAtt Script.Id` to retrieve either of these properties as input for `ScriptId` . Alternatively, enter a `ScriptId` string manually.
	ScriptId *string `json:"scriptId" yaml:"scriptId"`
}

// A CloudFormation `AWS::GameLift::GameServerGroup`.
//
// *This operation is used with the Amazon GameLift FleetIQ solution and game server groups.*
//
// Creates a GameLift FleetIQ game server group for managing game hosting on a collection of Amazon EC2 instances for game hosting. This operation creates the game server group, creates an Auto Scaling group in your AWS account , and establishes a link between the two groups. You can view the status of your game server groups in the GameLift console. Game server group metrics and events are emitted to Amazon CloudWatch.
//
// Before creating a new game server group, you must have the following:
//
// - An Amazon EC2 launch template that specifies how to launch Amazon EC2 instances with your game server build. For more information, see [Launching an Instance from a Launch Template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html) in the *Amazon EC2 User Guide* .
// - An IAM role that extends limited access to your AWS account to allow GameLift FleetIQ to create and interact with the Auto Scaling group. For more information, see [Create IAM roles for cross-service interaction](https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-iam-permissions-roles.html) in the *GameLift FleetIQ Developer Guide* .
//
// To create a new game server group, specify a unique group name, IAM role and Amazon EC2 launch template, and provide a list of instance types that can be used in the group. You must also set initial maximum and minimum limits on the group's instance count. You can optionally set an Auto Scaling policy with target tracking based on a GameLift FleetIQ metric.
//
// Once the game server group and corresponding Auto Scaling group are created, you have full access to change the Auto Scaling group's configuration as needed. Several properties that are set when creating a game server group, including maximum/minimum size and auto-scaling policy settings, must be updated directly in the Auto Scaling group. Keep in mind that some Auto Scaling group properties are periodically updated by GameLift FleetIQ as part of its balancing activities to optimize for availability and cost.
//
// *Learn more*
//
// [GameLift FleetIQ Guide](https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
//
// TODO: EXAMPLE
//
type CfnGameServerGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAutoScalingGroupArn() *string
	AttrGameServerGroupArn() *string
	AutoScalingPolicy() interface{}
	SetAutoScalingPolicy(val interface{})
	BalancingStrategy() *string
	SetBalancingStrategy(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DeleteOption() *string
	SetDeleteOption(val *string)
	GameServerGroupName() *string
	SetGameServerGroupName(val *string)
	GameServerProtectionPolicy() *string
	SetGameServerProtectionPolicy(val *string)
	InstanceDefinitions() interface{}
	SetInstanceDefinitions(val interface{})
	LaunchTemplate() interface{}
	SetLaunchTemplate(val interface{})
	LogicalId() *string
	MaxSize() *float64
	SetMaxSize(val *float64)
	MinSize() *float64
	SetMinSize(val *float64)
	Node() awscdk.ConstructNode
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VpcSubnets() *[]*string
	SetVpcSubnets(val *[]*string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnGameServerGroup
type jsiiProxy_CfnGameServerGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGameServerGroup) AttrAutoScalingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAutoScalingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) AttrGameServerGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGameServerGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) AutoScalingPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) BalancingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"balancingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) DeleteOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) GameServerGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) GameServerProtectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerProtectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) InstanceDefinitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) LaunchTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameServerGroup) VpcSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSubnets",
		&returns,
	)
	return returns
}


// Create a new `AWS::GameLift::GameServerGroup`.
func NewCfnGameServerGroup(scope awscdk.Construct, id *string, props *CfnGameServerGroupProps) CfnGameServerGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnGameServerGroup{}

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnGameServerGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::GameServerGroup`.
func NewCfnGameServerGroup_Override(c CfnGameServerGroup, scope awscdk.Construct, id *string, props *CfnGameServerGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnGameServerGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGameServerGroup) SetAutoScalingPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"autoScalingPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnGameServerGroup) SetBalancingStrategy(val *string) {
	_jsii_.Set(
		j,
		"balancingStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnGameServerGroup) SetDeleteOption(val *string) {
	_jsii_.Set(
		j,
		"deleteOption",
		val,
	)
}

func (j *jsiiProxy_CfnGameServerGroup) SetGameServerGroupName(val *string) {
	_jsii_.Set(
		j,
		"gameServerGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnGameServerGroup) SetGameServerProtectionPolicy(val *string) {
	_jsii_.Set(
		j,
		"gameServerProtectionPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnGameServerGroup) SetInstanceDefinitions(val interface{}) {
	_jsii_.Set(
		j,
		"instanceDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnGameServerGroup) SetLaunchTemplate(val interface{}) {
	_jsii_.Set(
		j,
		"launchTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnGameServerGroup) SetMaxSize(val *float64) {
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_CfnGameServerGroup) SetMinSize(val *float64) {
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_CfnGameServerGroup) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnGameServerGroup) SetVpcSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcSubnets",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnGameServerGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnGameServerGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnGameServerGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnGameServerGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnGameServerGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnGameServerGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGameServerGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_gamelift.CfnGameServerGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnGameServerGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnGameServerGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// *This data type is used with the GameLift FleetIQ and game server groups.*.
//
// Configuration settings for intelligent automatic scaling that uses target tracking. After the Auto Scaling group is created, all updates to Auto Scaling policies, including changing this policy and adding or removing other policies, is done directly on the Auto Scaling group.
//
// TODO: EXAMPLE
//
type CfnGameServerGroup_AutoScalingPolicyProperty struct {
	// Settings for a target-based scaling policy applied to Auto Scaling group.
	//
	// These settings are used to create a target-based policy that tracks the GameLift FleetIQ metric `PercentUtilizedGameServers` and specifies a target value for the metric. As player usage changes, the policy triggers to adjust the game server group capacity so that the metric returns to the target value.
	TargetTrackingConfiguration interface{} `json:"targetTrackingConfiguration" yaml:"targetTrackingConfiguration"`
	// Length of time, in seconds, it takes for a new instance to start new game server processes and register with GameLift FleetIQ.
	//
	// Specifying a warm-up time can be useful, particularly with game servers that take a long time to start up, because it avoids prematurely starting new instances.
	EstimatedInstanceWarmup *float64 `json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
}

// *This data type is used with the Amazon GameLift FleetIQ and game server groups.*.
//
// An allowed instance type for a `GameServerGroup` . All game server groups must have at least two instance types defined for it. GameLift FleetIQ periodically evaluates each defined instance type for viability. It then updates the Auto Scaling group with the list of viable instance types.
//
// TODO: EXAMPLE
//
type CfnGameServerGroup_InstanceDefinitionProperty struct {
	// An Amazon EC2 instance type designation.
	InstanceType *string `json:"instanceType" yaml:"instanceType"`
	// Instance weighting that indicates how much this instance type contributes to the total capacity of a game server group.
	//
	// Instance weights are used by GameLift FleetIQ to calculate the instance type's cost per unit hour and better identify the most cost-effective options. For detailed information on weighting instance capacity, see [Instance Weighting](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-weighting.html) in the *Amazon Elastic Compute Cloud Auto Scaling User Guide* . Default value is "1".
	WeightedCapacity *string `json:"weightedCapacity" yaml:"weightedCapacity"`
}

// *This data type is used with the GameLift FleetIQ and game server groups.*.
//
// An Amazon EC2 launch template that contains configuration settings and game server code to be deployed to all instances in a game server group. The launch template is specified when creating a new game server group with `GameServerGroup` .
//
// TODO: EXAMPLE
//
type CfnGameServerGroup_LaunchTemplateProperty struct {
	// A unique identifier for an existing Amazon EC2 launch template.
	LaunchTemplateId *string `json:"launchTemplateId" yaml:"launchTemplateId"`
	// A readable identifier for an existing Amazon EC2 launch template.
	LaunchTemplateName *string `json:"launchTemplateName" yaml:"launchTemplateName"`
	// The version of the Amazon EC2 launch template to use.
	//
	// If no version is specified, the default version will be used. With Amazon EC2, you can specify a default version for a launch template. If none is set, the default is the first version created.
	Version *string `json:"version" yaml:"version"`
}

// *This data type is used with the Amazon GameLift FleetIQ and game server groups.*.
//
// Settings for a target-based scaling policy as part of a `GameServerGroupAutoScalingPolicy` . These settings are used to create a target-based policy that tracks the GameLift FleetIQ metric `"PercentUtilizedGameServers"` and specifies a target value for the metric. As player usage changes, the policy triggers to adjust the game server group capacity so that the metric returns to the target value.
//
// TODO: EXAMPLE
//
type CfnGameServerGroup_TargetTrackingConfigurationProperty struct {
	// Desired value to use with a game server group target-based scaling policy.
	TargetValue *float64 `json:"targetValue" yaml:"targetValue"`
}

// Properties for defining a `CfnGameServerGroup`.
//
// TODO: EXAMPLE
//
type CfnGameServerGroupProps struct {
	// A developer-defined identifier for the game server group.
	//
	// The name is unique for each Region in each AWS account.
	GameServerGroupName *string `json:"gameServerGroupName" yaml:"gameServerGroupName"`
	// The set of Amazon EC2 instance types that GameLift FleetIQ can use when balancing and automatically scaling instances in the corresponding Auto Scaling group.
	InstanceDefinitions interface{} `json:"instanceDefinitions" yaml:"instanceDefinitions"`
	// The Amazon EC2 launch template that contains configuration settings and game server code to be deployed to all instances in the game server group.
	//
	// You can specify the template using either the template name or ID. For help with creating a launch template, see [Creating a Launch Template for an Auto Scaling Group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html) in the *Amazon Elastic Compute Cloud Auto Scaling User Guide* . After the Auto Scaling group is created, update this value directly in the Auto Scaling group using the AWS console or APIs.
	//
	// > If you specify network interfaces in your launch template, you must explicitly set the property `AssociatePublicIpAddress` to "true". If no network interface is specified in the launch template, GameLift FleetIQ uses your account's default VPC.
	LaunchTemplate interface{} `json:"launchTemplate" yaml:"launchTemplate"`
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) for an IAM role that allows Amazon Web Services to access your Amazon EC2 Auto Scaling groups.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Configuration settings to define a scaling policy for the Auto Scaling group that is optimized for game hosting.
	//
	// The scaling policy uses the metric `"PercentUtilizedGameServers"` to maintain a buffer of idle game servers that can immediately accommodate new games and players. After the Auto Scaling group is created, update this value directly in the Auto Scaling group using the AWS console or APIs.
	AutoScalingPolicy interface{} `json:"autoScalingPolicy" yaml:"autoScalingPolicy"`
	// Indicates how GameLift FleetIQ balances the use of Spot Instances and On-Demand Instances in the game server group.
	//
	// Method options include the following:
	//
	// - `SPOT_ONLY` - Only Spot Instances are used in the game server group. If Spot Instances are unavailable or not viable for game hosting, the game server group provides no hosting capacity until Spot Instances can again be used. Until then, no new instances are started, and the existing nonviable Spot Instances are terminated (after current gameplay ends) and are not replaced.
	// - `SPOT_PREFERRED` - (default value) Spot Instances are used whenever available in the game server group. If Spot Instances are unavailable, the game server group continues to provide hosting capacity by falling back to On-Demand Instances. Existing nonviable Spot Instances are terminated (after current gameplay ends) and are replaced with new On-Demand Instances.
	// - `ON_DEMAND_ONLY` - Only On-Demand Instances are used in the game server group. No Spot Instances are used, even when available, while this balancing strategy is in force.
	BalancingStrategy *string `json:"balancingStrategy" yaml:"balancingStrategy"`
	// The type of delete to perform.
	//
	// To delete a game server group, specify the `DeleteOption` . Options include the following:
	//
	// - `SAFE_DELETE` – (default) Terminates the game server group and Amazon EC2 Auto Scaling group only when it has no game servers that are in `UTILIZED` status.
	// - `FORCE_DELETE` – Terminates the game server group, including all active game servers regardless of their utilization status, and the Amazon EC2 Auto Scaling group.
	// - `RETAIN` – Does a safe delete of the game server group but retains the Amazon EC2 Auto Scaling group as is.
	DeleteOption *string `json:"deleteOption" yaml:"deleteOption"`
	// A flag that indicates whether instances in the game server group are protected from early termination.
	//
	// Unprotected instances that have active game servers running might be terminated during a scale-down event, causing players to be dropped from the game. Protected instances cannot be terminated while there are active game servers running except in the event of a forced game server group deletion (see ). An exception to this is with Spot Instances, which can be terminated by AWS regardless of protection status.
	GameServerProtectionPolicy *string `json:"gameServerProtectionPolicy" yaml:"gameServerProtectionPolicy"`
	// The maximum number of instances allowed in the Amazon EC2 Auto Scaling group.
	//
	// During automatic scaling events, GameLift FleetIQ and EC2 do not scale up the group above this maximum. After the Auto Scaling group is created, update this value directly in the Auto Scaling group using the AWS console or APIs.
	MaxSize *float64 `json:"maxSize" yaml:"maxSize"`
	// The minimum number of instances allowed in the Amazon EC2 Auto Scaling group.
	//
	// During automatic scaling events, GameLift FleetIQ and Amazon EC2 do not scale down the group below this minimum. In production, this value should be set to at least 1. After the Auto Scaling group is created, update this value directly in the Auto Scaling group using the AWS console or APIs.
	MinSize *float64 `json:"minSize" yaml:"minSize"`
	// A list of labels to assign to the new game server group resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management, and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* . Once the resource is created, you can use TagResource, UntagResource, and ListTagsForResource to add, remove, and view tags, respectively. The maximum tag limit may be lower than stated. See the AWS General Reference for actual tagging limits.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// A list of virtual private cloud (VPC) subnets to use with instances in the game server group.
	//
	// By default, all GameLift FleetIQ-supported Availability Zones are used. You can use this parameter to specify VPCs that you've set up. This property cannot be updated after the game server group is created, and the corresponding Auto Scaling group will always use the property value that is set with this request, even if the Auto Scaling group is updated directly.
	VpcSubnets *[]*string `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// A CloudFormation `AWS::GameLift::GameSessionQueue`.
//
// The `AWS::GameLift::GameSessionQueue` resource creates a placement queue that processes requests for new game sessions. A queue uses FleetIQ algorithms to determine the best placement locations and find an available game server, then prompts the game server to start a new game session. Queues can have destinations (GameLift fleets or aliases), which determine where the queue can place new game sessions. A queue can have destinations with varied fleet type (Spot and On-Demand), instance type, and AWS Region .
//
// TODO: EXAMPLE
//
type CfnGameSessionQueue interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	CustomEventData() *string
	SetCustomEventData(val *string)
	Destinations() interface{}
	SetDestinations(val interface{})
	FilterConfiguration() interface{}
	SetFilterConfiguration(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	NotificationTarget() *string
	SetNotificationTarget(val *string)
	PlayerLatencyPolicies() interface{}
	SetPlayerLatencyPolicies(val interface{})
	PriorityConfiguration() interface{}
	SetPriorityConfiguration(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	TimeoutInSeconds() *float64
	SetTimeoutInSeconds(val *float64)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnGameSessionQueue
type jsiiProxy_CfnGameSessionQueue struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGameSessionQueue) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) CustomEventData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEventData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) Destinations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) FilterConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) NotificationTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) PlayerLatencyPolicies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"playerLatencyPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) PriorityConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"priorityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueue) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::GameLift::GameSessionQueue`.
func NewCfnGameSessionQueue(scope awscdk.Construct, id *string, props *CfnGameSessionQueueProps) CfnGameSessionQueue {
	_init_.Initialize()

	j := jsiiProxy_CfnGameSessionQueue{}

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnGameSessionQueue",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::GameSessionQueue`.
func NewCfnGameSessionQueue_Override(c CfnGameSessionQueue, scope awscdk.Construct, id *string, props *CfnGameSessionQueueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnGameSessionQueue",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGameSessionQueue) SetCustomEventData(val *string) {
	_jsii_.Set(
		j,
		"customEventData",
		val,
	)
}

func (j *jsiiProxy_CfnGameSessionQueue) SetDestinations(val interface{}) {
	_jsii_.Set(
		j,
		"destinations",
		val,
	)
}

func (j *jsiiProxy_CfnGameSessionQueue) SetFilterConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"filterConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnGameSessionQueue) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnGameSessionQueue) SetNotificationTarget(val *string) {
	_jsii_.Set(
		j,
		"notificationTarget",
		val,
	)
}

func (j *jsiiProxy_CfnGameSessionQueue) SetPlayerLatencyPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"playerLatencyPolicies",
		val,
	)
}

func (j *jsiiProxy_CfnGameSessionQueue) SetPriorityConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"priorityConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnGameSessionQueue) SetTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnGameSessionQueue_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnGameSessionQueue",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnGameSessionQueue_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnGameSessionQueue",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnGameSessionQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnGameSessionQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGameSessionQueue_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_gamelift.CfnGameSessionQueue",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnGameSessionQueue) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnGameSessionQueue) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A fleet or alias designated in a game session queue.
//
// Queues fulfill requests for new game sessions by placing a new game session on any of the queue's destinations.
//
// TODO: EXAMPLE
//
type CfnGameSessionQueue_DestinationProperty struct {
	// The Amazon Resource Name (ARN) that is assigned to fleet or fleet alias.
	//
	// ARNs, which include a fleet ID or alias ID and a Region name, provide a unique identifier across all Regions.
	DestinationArn *string `json:"destinationArn" yaml:"destinationArn"`
}

// A list of fleet locations where a game session queue can place new game sessions.
//
// You can use a filter to temporarily turn off placements for specific locations. For queues that have multi-location fleets, you can use a filter configuration allow placement with some, but not all of these locations.
//
// TODO: EXAMPLE
//
type CfnGameSessionQueue_FilterConfigurationProperty struct {
	// A list of locations to allow game session placement in, in the form of AWS Region codes such as `us-west-2` .
	AllowedLocations *[]*string `json:"allowedLocations" yaml:"allowedLocations"`
}

// The queue setting that determines the highest latency allowed for individual players when placing a game session.
//
// When a latency policy is in force, a game session cannot be placed with any fleet in a Region where a player reports latency higher than the cap. Latency policies are only enforced when the placement request contains player latency information.
//
// TODO: EXAMPLE
//
type CfnGameSessionQueue_PlayerLatencyPolicyProperty struct {
	// The maximum latency value that is allowed for any player, in milliseconds.
	//
	// All policies must have a value set for this property.
	MaximumIndividualPlayerLatencyMilliseconds *float64 `json:"maximumIndividualPlayerLatencyMilliseconds" yaml:"maximumIndividualPlayerLatencyMilliseconds"`
	// The length of time, in seconds, that the policy is enforced while placing a new game session.
	//
	// A null value for this property means that the policy is enforced until the queue times out.
	PolicyDurationSeconds *float64 `json:"policyDurationSeconds" yaml:"policyDurationSeconds"`
}

// Custom prioritization settings for use by a game session queue when placing new game sessions with available game servers.
//
// When defined, this configuration replaces the default FleetIQ prioritization process, which is as follows:
//
// - If player latency data is included in a game session request, destinations and locations are prioritized first based on lowest average latency (1), then on lowest hosting cost (2), then on destination list order (3), and finally on location (alphabetical) (4). This approach ensures that the queue's top priority is to place game sessions where average player latency is lowest, and--if latency is the same--where the hosting cost is less, etc.
// - If player latency data is not included, destinations and locations are prioritized first on destination list order (1), and then on location (alphabetical) (2). This approach ensures that the queue's top priority is to place game sessions on the first destination fleet listed. If that fleet has multiple locations, the game session is placed on the first location (when listed alphabetically).
//
// Changing the priority order will affect how game sessions are placed.
//
// TODO: EXAMPLE
//
type CfnGameSessionQueue_PriorityConfigurationProperty struct {
	// The prioritization order to use for fleet locations, when the `PriorityOrder` property includes `LOCATION` .
	//
	// Locations are identified by AWS Region codes such as `us-west-2` . Each location can only be listed once.
	LocationOrder *[]*string `json:"locationOrder" yaml:"locationOrder"`
	// The recommended sequence to use when prioritizing where to place new game sessions.
	//
	// Each type can only be listed once.
	//
	// - `LATENCY` -- FleetIQ prioritizes locations where the average player latency (provided in each game session request) is lowest.
	// - `COST` -- FleetIQ prioritizes destinations with the lowest current hosting costs. Cost is evaluated based on the location, instance type, and fleet type (Spot or On-Demand) for each destination in the queue.
	// - `DESTINATION` -- FleetIQ prioritizes based on the order that destinations are listed in the queue configuration.
	// - `LOCATION` -- FleetIQ prioritizes based on the provided order of locations, as defined in `LocationOrder` .
	PriorityOrder *[]*string `json:"priorityOrder" yaml:"priorityOrder"`
}

// Properties for defining a `CfnGameSessionQueue`.
//
// TODO: EXAMPLE
//
type CfnGameSessionQueueProps struct {
	// A descriptive label that is associated with game session queue.
	//
	// Queue names must be unique within each Region.
	Name *string `json:"name" yaml:"name"`
	// Information to be added to all events that are related to this game session queue.
	CustomEventData *string `json:"customEventData" yaml:"customEventData"`
	// A list of fleets and/or fleet aliases that can be used to fulfill game session placement requests in the queue.
	//
	// Destinations are identified by either a fleet ARN or a fleet alias ARN, and are listed in order of placement preference.
	Destinations interface{} `json:"destinations" yaml:"destinations"`
	// A list of locations where a queue is allowed to place new game sessions.
	//
	// Locations are specified in the form of AWS Region codes, such as `us-west-2` . If this parameter is not set, game sessions can be placed in any queue location.
	FilterConfiguration interface{} `json:"filterConfiguration" yaml:"filterConfiguration"`
	// An SNS topic ARN that is set up to receive game session placement notifications.
	//
	// See [Setting up notifications for game session placement](https://docs.aws.amazon.com/gamelift/latest/developerguide/queue-notification.html) .
	NotificationTarget *string `json:"notificationTarget" yaml:"notificationTarget"`
	// A set of policies that act as a sliding cap on player latency.
	//
	// FleetIQ works to deliver low latency for most players in a game session. These policies ensure that no individual player can be placed into a game with unreasonably high latency. Use multiple policies to gradually relax latency requirements a step at a time. Multiple policies are applied based on their maximum allowed latency, starting with the lowest value.
	PlayerLatencyPolicies interface{} `json:"playerLatencyPolicies" yaml:"playerLatencyPolicies"`
	// Custom settings to use when prioritizing destinations and locations for game session placements.
	//
	// This configuration replaces the FleetIQ default prioritization process. Priority types that are not explicitly named will be automatically applied at the end of the prioritization process.
	PriorityConfiguration interface{} `json:"priorityConfiguration" yaml:"priorityConfiguration"`
	// A list of labels to assign to the new game session queue resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources are useful for resource management, access management and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* . Once the resource is created, you can use TagResource, UntagResource, and ListTagsForResource to add, remove, and view tags. The maximum tag limit may be lower than stated. See the AWS General Reference for actual tagging limits.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The maximum time, in seconds, that a new game session placement request remains in the queue.
	//
	// When a request exceeds this time, the game session placement changes to a `TIMED_OUT` status.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

// A CloudFormation `AWS::GameLift::MatchmakingConfiguration`.
//
// The `AWS::GameLift::MatchmakingConfiguration` resource defines a new matchmaking configuration for use with FlexMatch. Whether you're using FlexMatch with GameLift hosting or as a standalone matchmaking service, the matchmaking configuration sets out rules for matching players and forming teams. If you're using GameLift hosting, it also defines how to start game sessions for each match. Your matchmaking system can use multiple configurations to handle different game scenarios. All matchmaking requests identify the matchmaking configuration to use and provide player attributes that are consistent with that configuration.
//
// TODO: EXAMPLE
//
type CfnMatchmakingConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AcceptanceRequired() interface{}
	SetAcceptanceRequired(val interface{})
	AcceptanceTimeoutSeconds() *float64
	SetAcceptanceTimeoutSeconds(val *float64)
	AdditionalPlayerCount() *float64
	SetAdditionalPlayerCount(val *float64)
	AttrArn() *string
	AttrName() *string
	BackfillMode() *string
	SetBackfillMode(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	CustomEventData() *string
	SetCustomEventData(val *string)
	Description() *string
	SetDescription(val *string)
	FlexMatchMode() *string
	SetFlexMatchMode(val *string)
	GameProperties() interface{}
	SetGameProperties(val interface{})
	GameSessionData() *string
	SetGameSessionData(val *string)
	GameSessionQueueArns() *[]*string
	SetGameSessionQueueArns(val *[]*string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	NotificationTarget() *string
	SetNotificationTarget(val *string)
	Ref() *string
	RequestTimeoutSeconds() *float64
	SetRequestTimeoutSeconds(val *float64)
	RuleSetName() *string
	SetRuleSetName(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnMatchmakingConfiguration
type jsiiProxy_CfnMatchmakingConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) AcceptanceRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptanceRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) AcceptanceTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"acceptanceTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) AdditionalPlayerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalPlayerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) BackfillMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backfillMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) CustomEventData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEventData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) FlexMatchMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flexMatchMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) GameProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gameProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) GameSessionData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameSessionData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) GameSessionQueueArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gameSessionQueueArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) NotificationTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) RequestTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) RuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::GameLift::MatchmakingConfiguration`.
func NewCfnMatchmakingConfiguration(scope awscdk.Construct, id *string, props *CfnMatchmakingConfigurationProps) CfnMatchmakingConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnMatchmakingConfiguration{}

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnMatchmakingConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::MatchmakingConfiguration`.
func NewCfnMatchmakingConfiguration_Override(c CfnMatchmakingConfiguration, scope awscdk.Construct, id *string, props *CfnMatchmakingConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnMatchmakingConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetAcceptanceRequired(val interface{}) {
	_jsii_.Set(
		j,
		"acceptanceRequired",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetAcceptanceTimeoutSeconds(val *float64) {
	_jsii_.Set(
		j,
		"acceptanceTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetAdditionalPlayerCount(val *float64) {
	_jsii_.Set(
		j,
		"additionalPlayerCount",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetBackfillMode(val *string) {
	_jsii_.Set(
		j,
		"backfillMode",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetCustomEventData(val *string) {
	_jsii_.Set(
		j,
		"customEventData",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetFlexMatchMode(val *string) {
	_jsii_.Set(
		j,
		"flexMatchMode",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetGameProperties(val interface{}) {
	_jsii_.Set(
		j,
		"gameProperties",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetGameSessionData(val *string) {
	_jsii_.Set(
		j,
		"gameSessionData",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetGameSessionQueueArns(val *[]*string) {
	_jsii_.Set(
		j,
		"gameSessionQueueArns",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetNotificationTarget(val *string) {
	_jsii_.Set(
		j,
		"notificationTarget",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetRequestTimeoutSeconds(val *float64) {
	_jsii_.Set(
		j,
		"requestTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingConfiguration) SetRuleSetName(val *string) {
	_jsii_.Set(
		j,
		"ruleSetName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnMatchmakingConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnMatchmakingConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnMatchmakingConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnMatchmakingConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnMatchmakingConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnMatchmakingConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMatchmakingConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_gamelift.CfnMatchmakingConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnMatchmakingConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnMatchmakingConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Set of key-value pairs that contain information about a game session.
//
// When included in a game session request, these properties communicate details to be used when setting up the new game session. For example, a game property might specify a game mode, level, or map. Game properties are passed to the game server process when initiating a new game session. For more information, see the [GameLift Developer Guide](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-client-api.html#gamelift-sdk-client-api-create) .
//
// TODO: EXAMPLE
//
type CfnMatchmakingConfiguration_GamePropertyProperty struct {
	// The game property identifier.
	Key *string `json:"key" yaml:"key"`
	// The game property value.
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnMatchmakingConfiguration`.
//
// TODO: EXAMPLE
//
type CfnMatchmakingConfigurationProps struct {
	// A flag that determines whether a match that was created with this configuration must be accepted by the matched players.
	//
	// To require acceptance, set to `TRUE` . With this option enabled, matchmaking tickets use the status `REQUIRES_ACCEPTANCE` to indicate when a completed potential match is waiting for player acceptance.
	AcceptanceRequired interface{} `json:"acceptanceRequired" yaml:"acceptanceRequired"`
	// A unique identifier for the matchmaking configuration.
	//
	// This name is used to identify the configuration associated with a matchmaking request or ticket.
	Name *string `json:"name" yaml:"name"`
	// The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.
	//
	// Requests that fail due to timing out can be resubmitted as needed.
	RequestTimeoutSeconds *float64 `json:"requestTimeoutSeconds" yaml:"requestTimeoutSeconds"`
	// A unique identifier for the matchmaking rule set to use with this configuration.
	//
	// You can use either the rule set name or ARN value. A matchmaking configuration can only use rule sets that are defined in the same Region.
	RuleSetName *string `json:"ruleSetName" yaml:"ruleSetName"`
	// The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.
	AcceptanceTimeoutSeconds *float64 `json:"acceptanceTimeoutSeconds" yaml:"acceptanceTimeoutSeconds"`
	// The number of player slots in a match to keep open for future players.
	//
	// For example, if the configuration's rule set specifies a match for a single 12-person team, and the additional player count is set to 2, only 10 players are selected for the match. This parameter is not used if `FlexMatchMode` is set to `STANDALONE` .
	AdditionalPlayerCount *float64 `json:"additionalPlayerCount" yaml:"additionalPlayerCount"`
	// The method used to backfill game sessions that are created with this matchmaking configuration.
	//
	// Specify `MANUAL` when your game manages backfill requests manually or does not use the match backfill feature. Specify `AUTOMATIC` to have GameLift create a `StartMatchBackfill` request whenever a game session has one or more open slots. Learn more about manual and automatic backfill in [Backfill Existing Games with FlexMatch](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-backfill.html) . Automatic backfill is not available when `FlexMatchMode` is set to `STANDALONE` .
	BackfillMode *string `json:"backfillMode" yaml:"backfillMode"`
	// Information to add to all events related to the matchmaking configuration.
	CustomEventData *string `json:"customEventData" yaml:"customEventData"`
	// A descriptive label that is associated with matchmaking configuration.
	Description *string `json:"description" yaml:"description"`
	// Indicates whether this matchmaking configuration is being used with GameLift hosting or as a standalone matchmaking solution.
	//
	// - *STANDALONE* - FlexMatch forms matches and returns match information, including players and team assignments, in a [MatchmakingSucceeded](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-events.html#match-events-matchmakingsucceeded) event.
	// - *WITH_QUEUE* - FlexMatch forms matches and uses the specified GameLift queue to start a game session for the match.
	FlexMatchMode *string `json:"flexMatchMode" yaml:"flexMatchMode"`
	// A set of custom properties for a game session, formatted as key-value pairs.
	//
	// These properties are passed to a game server process with a request to start a new game session. See [Start a Game Session](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession) . This parameter is not used if `FlexMatchMode` is set to `STANDALONE` .
	GameProperties interface{} `json:"gameProperties" yaml:"gameProperties"`
	// A set of custom game session properties, formatted as a single string value.
	//
	// This data is passed to a game server process with a request to start a new game session. See [Start a Game Session](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession) . This parameter is not used if `FlexMatchMode` is set to `STANDALONE` .
	GameSessionData *string `json:"gameSessionData" yaml:"gameSessionData"`
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) that is assigned to a GameLift game session queue resource and uniquely identifies it. ARNs are unique across all Regions. Format is `arn:aws:gamelift:<region>::gamesessionqueue/<queue name>` . Queues can be located in any Region. Queues are used to start new GameLift-hosted game sessions for matches that are created with this matchmaking configuration. If `FlexMatchMode` is set to `STANDALONE` , do not set this parameter.
	GameSessionQueueArns *[]*string `json:"gameSessionQueueArns" yaml:"gameSessionQueueArns"`
	// An SNS topic ARN that is set up to receive matchmaking notifications.
	//
	// See [Setting up notifications for matchmaking](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html) for more information.
	NotificationTarget *string `json:"notificationTarget" yaml:"notificationTarget"`
	// A list of labels to assign to the new matchmaking configuration resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources are useful for resource management, access management and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* . Once the resource is created, you can use TagResource, UntagResource, and ListTagsForResource to add, remove, and view tags. The maximum tag limit may be lower than stated. See the AWS General Reference for actual tagging limits.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::GameLift::MatchmakingRuleSet`.
//
// Creates a new rule set for FlexMatch matchmaking. A rule set describes the type of match to create, such as the number and size of teams. It also sets the parameters for acceptable player matches, such as minimum skill level or character type.
//
// To create a matchmaking rule set, provide unique rule set name and the rule set body in JSON format. Rule sets must be defined in the same Region as the matchmaking configuration they are used with.
//
// Since matchmaking rule sets cannot be edited, it is a good idea to check the rule set syntax.
//
// *Learn more*
//
// - [Build a rule set](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-rulesets.html)
// - [Design a matchmaker](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-configuration.html)
// - [Matchmaking with FlexMatch](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-intro.html)
//
// TODO: EXAMPLE
//
type CfnMatchmakingRuleSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	RuleSetBody() *string
	SetRuleSetBody(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnMatchmakingRuleSet
type jsiiProxy_CfnMatchmakingRuleSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) RuleSetBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::GameLift::MatchmakingRuleSet`.
func NewCfnMatchmakingRuleSet(scope awscdk.Construct, id *string, props *CfnMatchmakingRuleSetProps) CfnMatchmakingRuleSet {
	_init_.Initialize()

	j := jsiiProxy_CfnMatchmakingRuleSet{}

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnMatchmakingRuleSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::MatchmakingRuleSet`.
func NewCfnMatchmakingRuleSet_Override(c CfnMatchmakingRuleSet, scope awscdk.Construct, id *string, props *CfnMatchmakingRuleSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnMatchmakingRuleSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnMatchmakingRuleSet) SetRuleSetBody(val *string) {
	_jsii_.Set(
		j,
		"ruleSetBody",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnMatchmakingRuleSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnMatchmakingRuleSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnMatchmakingRuleSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnMatchmakingRuleSet",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnMatchmakingRuleSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnMatchmakingRuleSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMatchmakingRuleSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_gamelift.CfnMatchmakingRuleSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnMatchmakingRuleSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnMatchmakingRuleSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnMatchmakingRuleSet`.
//
// TODO: EXAMPLE
//
type CfnMatchmakingRuleSetProps struct {
	// A unique identifier for the matchmaking rule set.
	//
	// A matchmaking configuration identifies the rule set it uses by this name value. Note that the rule set name is different from the optional `name` field in the rule set body.
	Name *string `json:"name" yaml:"name"`
	// A collection of matchmaking rules, formatted as a JSON string.
	//
	// Comments are not allowed in JSON, but most elements support a description field.
	RuleSetBody *string `json:"ruleSetBody" yaml:"ruleSetBody"`
	// A list of labels to assign to the new matchmaking rule set resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources are useful for resource management, access management and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* . Once the resource is created, you can use TagResource, UntagResource, and ListTagsForResource to add, remove, and view tags. The maximum tag limit may be lower than stated. See the AWS General Reference for actual tagging limits.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::GameLift::Script`.
//
// The `AWS::GameLift::Script` resource creates a new script record for your Realtime Servers script. Realtime scripts are JavaScript that provide configuration settings and optional custom game logic for your game. The script is deployed when you create a Realtime Servers fleet to host your game sessions. Script logic is executed during an active game session.
//
// TODO: EXAMPLE
//
type CfnScript interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	Stack() awscdk.Stack
	StorageLocation() interface{}
	SetStorageLocation(val interface{})
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	Version() *string
	SetVersion(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnScript
type jsiiProxy_CfnScript struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScript) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) StorageLocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScript) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new `AWS::GameLift::Script`.
func NewCfnScript(scope awscdk.Construct, id *string, props *CfnScriptProps) CfnScript {
	_init_.Initialize()

	j := jsiiProxy_CfnScript{}

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnScript",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::Script`.
func NewCfnScript_Override(c CfnScript, scope awscdk.Construct, id *string, props *CfnScriptProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_gamelift.CfnScript",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScript) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnScript) SetStorageLocation(val interface{}) {
	_jsii_.Set(
		j,
		"storageLocation",
		val,
	)
}

func (j *jsiiProxy_CfnScript) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnScript_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnScript",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnScript_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnScript",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnScript_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_gamelift.CfnScript",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScript_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_gamelift.CfnScript",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnScript) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnScript) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnScript) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnScript) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnScript) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnScript) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CfnScript) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnScript) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnScript) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnScript) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnScript) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnScript) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnScript) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnScript) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnScript) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnScript) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnScript) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnScript) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnScript) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnScript) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnScript) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The location in Amazon S3 where build or script files can be stored for access by Amazon GameLift.
//
// TODO: EXAMPLE
//
type CfnScript_S3LocationProperty struct {
	// An Amazon S3 bucket identifier. This is the name of the S3 bucket.
	//
	// > GameLift currently does not support uploading from Amazon S3 buckets with names that contain a dot (.).
	Bucket *string `json:"bucket" yaml:"bucket"`
	// The name of the zip file that contains the build files or script files.
	Key *string `json:"key" yaml:"key"`
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) for an IAM role that allows Amazon Web Services to access the S3 bucket.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The version of the file, if object versioning is turned on for the bucket.
	//
	// Amazon Web Services uses this information when retrieving files from an S3 bucket that you own. Use this parameter to specify a specific version of the file. If not set, the latest version of the file is retrieved.
	ObjectVersion *string `json:"objectVersion" yaml:"objectVersion"`
}

// Properties for defining a `CfnScript`.
//
// TODO: EXAMPLE
//
type CfnScriptProps struct {
	// The location of the Amazon S3 bucket where a zipped file containing your Realtime scripts is stored.
	//
	// The storage location must specify the Amazon S3 bucket name, the zip file name (the "key"), and a role ARN that allows Amazon Web Services to access the Amazon S3 storage location. The S3 bucket must be in the same Region where you want to create a new script. By default, Amazon Web Services uploads the latest version of the zip file; if you have S3 object versioning turned on, you can use the `ObjectVersion` parameter to specify an earlier version.
	StorageLocation interface{} `json:"storageLocation" yaml:"storageLocation"`
	// A descriptive label that is associated with a script.
	//
	// Script names do not need to be unique.
	Name *string `json:"name" yaml:"name"`
	// A list of labels to assign to the new script resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources are useful for resource management, access management and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* . Once the resource is created, you can use TagResource, UntagResource, and ListTagsForResource to add, remove, and view tags. The maximum tag limit may be lower than stated. See the AWS General Reference for actual tagging limits.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The version that is associated with a build or script.
	//
	// Version strings do not need to be unique.
	Version *string `json:"version" yaml:"version"`
}

