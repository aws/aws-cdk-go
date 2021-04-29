package awsgamelift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::GameLift::Alias`.
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnAlias) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnAlias(scope constructs.Construct, id *string, props *CfnAliasProps) CfnAlias {
	_init_.Initialize()

	j := jsiiProxy_CfnAlias{}

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnAlias",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::Alias`.
func NewCfnAlias_Override(c CfnAlias, scope constructs.Construct, id *string, props *CfnAliasProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnAlias",
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
		"aws-cdk-lib.aws_gamelift.CfnAlias",
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
		"aws-cdk-lib.aws_gamelift.CfnAlias",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnAlias",
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
		"aws-cdk-lib.aws_gamelift.CfnAlias",
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnAlias) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Experimental.
func (c *jsiiProxy_CfnAlias) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnAlias_RoutingStrategyProperty struct {
	// `CfnAlias.RoutingStrategyProperty.Type`.
	Type *string `json:"type"`
	// `CfnAlias.RoutingStrategyProperty.FleetId`.
	FleetId *string `json:"fleetId"`
	// `CfnAlias.RoutingStrategyProperty.Message`.
	Message *string `json:"message"`
}

// Properties for defining a `AWS::GameLift::Alias`.
type CfnAliasProps struct {
	// `AWS::GameLift::Alias.Name`.
	Name *string `json:"name"`
	// `AWS::GameLift::Alias.RoutingStrategy`.
	RoutingStrategy interface{} `json:"routingStrategy"`
	// `AWS::GameLift::Alias.Description`.
	Description *string `json:"description"`
}

// A CloudFormation `AWS::GameLift::Build`.
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnBuild) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnBuild(scope constructs.Construct, id *string, props *CfnBuildProps) CfnBuild {
	_init_.Initialize()

	j := jsiiProxy_CfnBuild{}

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnBuild",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::Build`.
func NewCfnBuild_Override(c CfnBuild, scope constructs.Construct, id *string, props *CfnBuildProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnBuild",
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
		"aws-cdk-lib.aws_gamelift.CfnBuild",
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
		"aws-cdk-lib.aws_gamelift.CfnBuild",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnBuild_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnBuild",
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
		"aws-cdk-lib.aws_gamelift.CfnBuild",
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnBuild) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Experimental.
func (c *jsiiProxy_CfnBuild) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnBuild_S3LocationProperty struct {
	// `CfnBuild.S3LocationProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnBuild.S3LocationProperty.Key`.
	Key *string `json:"key"`
	// `CfnBuild.S3LocationProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnBuild.S3LocationProperty.ObjectVersion`.
	ObjectVersion *string `json:"objectVersion"`
}

// Properties for defining a `AWS::GameLift::Build`.
type CfnBuildProps struct {
	// `AWS::GameLift::Build.Name`.
	Name *string `json:"name"`
	// `AWS::GameLift::Build.OperatingSystem`.
	OperatingSystem *string `json:"operatingSystem"`
	// `AWS::GameLift::Build.StorageLocation`.
	StorageLocation interface{} `json:"storageLocation"`
	// `AWS::GameLift::Build.Version`.
	Version *string `json:"version"`
}

// A CloudFormation `AWS::GameLift::Fleet`.
type CfnFleet interface {
	awscdk.CfnResource
	awscdk.IInspectable
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
	LogicalId() *string
	LogPaths() *[]*string
	SetLogPaths(val *[]*string)
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
	Node() constructs.Node
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
	ServerLaunchParameters() *string
	SetServerLaunchParameters(val *string)
	ServerLaunchPath() *string
	SetServerLaunchPath(val *string)
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnFleet
type jsiiProxy_CfnFleet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
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

func (j *jsiiProxy_CfnFleet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) LogPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logPaths",
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

func (j *jsiiProxy_CfnFleet) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnFleet) ServerLaunchParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverLaunchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) ServerLaunchPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverLaunchPath",
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
func NewCfnFleet(scope constructs.Construct, id *string, props *CfnFleetProps) CfnFleet {
	_init_.Initialize()

	j := jsiiProxy_CfnFleet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnFleet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::Fleet`.
func NewCfnFleet_Override(c CfnFleet, scope constructs.Construct, id *string, props *CfnFleetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnFleet",
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

func (j *jsiiProxy_CfnFleet) SetLogPaths(val *[]*string) {
	_jsii_.Set(
		j,
		"logPaths",
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

func (j *jsiiProxy_CfnFleet) SetServerLaunchParameters(val *string) {
	_jsii_.Set(
		j,
		"serverLaunchParameters",
		val,
	)
}

func (j *jsiiProxy_CfnFleet) SetServerLaunchPath(val *string) {
	_jsii_.Set(
		j,
		"serverLaunchPath",
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
		"aws-cdk-lib.aws_gamelift.CfnFleet",
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
		"aws-cdk-lib.aws_gamelift.CfnFleet",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnFleet",
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
		"aws-cdk-lib.aws_gamelift.CfnFleet",
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnFleet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Experimental.
func (c *jsiiProxy_CfnFleet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnFleet_CertificateConfigurationProperty struct {
	// `CfnFleet.CertificateConfigurationProperty.CertificateType`.
	CertificateType *string `json:"certificateType"`
}

type CfnFleet_IpPermissionProperty struct {
	// `CfnFleet.IpPermissionProperty.FromPort`.
	FromPort *float64 `json:"fromPort"`
	// `CfnFleet.IpPermissionProperty.IpRange`.
	IpRange *string `json:"ipRange"`
	// `CfnFleet.IpPermissionProperty.Protocol`.
	Protocol *string `json:"protocol"`
	// `CfnFleet.IpPermissionProperty.ToPort`.
	ToPort *float64 `json:"toPort"`
}

type CfnFleet_ResourceCreationLimitPolicyProperty struct {
	// `CfnFleet.ResourceCreationLimitPolicyProperty.NewGameSessionsPerCreator`.
	NewGameSessionsPerCreator *float64 `json:"newGameSessionsPerCreator"`
	// `CfnFleet.ResourceCreationLimitPolicyProperty.PolicyPeriodInMinutes`.
	PolicyPeriodInMinutes *float64 `json:"policyPeriodInMinutes"`
}

type CfnFleet_RuntimeConfigurationProperty struct {
	// `CfnFleet.RuntimeConfigurationProperty.GameSessionActivationTimeoutSeconds`.
	GameSessionActivationTimeoutSeconds *float64 `json:"gameSessionActivationTimeoutSeconds"`
	// `CfnFleet.RuntimeConfigurationProperty.MaxConcurrentGameSessionActivations`.
	MaxConcurrentGameSessionActivations *float64 `json:"maxConcurrentGameSessionActivations"`
	// `CfnFleet.RuntimeConfigurationProperty.ServerProcesses`.
	ServerProcesses interface{} `json:"serverProcesses"`
}

type CfnFleet_ServerProcessProperty struct {
	// `CfnFleet.ServerProcessProperty.ConcurrentExecutions`.
	ConcurrentExecutions *float64 `json:"concurrentExecutions"`
	// `CfnFleet.ServerProcessProperty.LaunchPath`.
	LaunchPath *string `json:"launchPath"`
	// `CfnFleet.ServerProcessProperty.Parameters`.
	Parameters *string `json:"parameters"`
}

// Properties for defining a `AWS::GameLift::Fleet`.
type CfnFleetProps struct {
	// `AWS::GameLift::Fleet.EC2InstanceType`.
	Ec2InstanceType *string `json:"ec2InstanceType"`
	// `AWS::GameLift::Fleet.Name`.
	Name *string `json:"name"`
	// `AWS::GameLift::Fleet.BuildId`.
	BuildId *string `json:"buildId"`
	// `AWS::GameLift::Fleet.CertificateConfiguration`.
	CertificateConfiguration interface{} `json:"certificateConfiguration"`
	// `AWS::GameLift::Fleet.Description`.
	Description *string `json:"description"`
	// `AWS::GameLift::Fleet.DesiredEC2Instances`.
	DesiredEc2Instances *float64 `json:"desiredEc2Instances"`
	// `AWS::GameLift::Fleet.EC2InboundPermissions`.
	Ec2InboundPermissions interface{} `json:"ec2InboundPermissions"`
	// `AWS::GameLift::Fleet.FleetType`.
	FleetType *string `json:"fleetType"`
	// `AWS::GameLift::Fleet.InstanceRoleARN`.
	InstanceRoleArn *string `json:"instanceRoleArn"`
	// `AWS::GameLift::Fleet.LogPaths`.
	LogPaths *[]*string `json:"logPaths"`
	// `AWS::GameLift::Fleet.MaxSize`.
	MaxSize *float64 `json:"maxSize"`
	// `AWS::GameLift::Fleet.MetricGroups`.
	MetricGroups *[]*string `json:"metricGroups"`
	// `AWS::GameLift::Fleet.MinSize`.
	MinSize *float64 `json:"minSize"`
	// `AWS::GameLift::Fleet.NewGameSessionProtectionPolicy`.
	NewGameSessionProtectionPolicy *string `json:"newGameSessionProtectionPolicy"`
	// `AWS::GameLift::Fleet.PeerVpcAwsAccountId`.
	PeerVpcAwsAccountId *string `json:"peerVpcAwsAccountId"`
	// `AWS::GameLift::Fleet.PeerVpcId`.
	PeerVpcId *string `json:"peerVpcId"`
	// `AWS::GameLift::Fleet.ResourceCreationLimitPolicy`.
	ResourceCreationLimitPolicy interface{} `json:"resourceCreationLimitPolicy"`
	// `AWS::GameLift::Fleet.RuntimeConfiguration`.
	RuntimeConfiguration interface{} `json:"runtimeConfiguration"`
	// `AWS::GameLift::Fleet.ScriptId`.
	ScriptId *string `json:"scriptId"`
	// `AWS::GameLift::Fleet.ServerLaunchParameters`.
	ServerLaunchParameters *string `json:"serverLaunchParameters"`
	// `AWS::GameLift::Fleet.ServerLaunchPath`.
	ServerLaunchPath *string `json:"serverLaunchPath"`
}

// A CloudFormation `AWS::GameLift::GameServerGroup`.
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnGameServerGroup) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnGameServerGroup(scope constructs.Construct, id *string, props *CfnGameServerGroupProps) CfnGameServerGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnGameServerGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnGameServerGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::GameServerGroup`.
func NewCfnGameServerGroup_Override(c CfnGameServerGroup, scope constructs.Construct, id *string, props *CfnGameServerGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnGameServerGroup",
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
		"aws-cdk-lib.aws_gamelift.CfnGameServerGroup",
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
		"aws-cdk-lib.aws_gamelift.CfnGameServerGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnGameServerGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnGameServerGroup",
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
		"aws-cdk-lib.aws_gamelift.CfnGameServerGroup",
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Experimental.
func (c *jsiiProxy_CfnGameServerGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnGameServerGroup_AutoScalingPolicyProperty struct {
	// `CfnGameServerGroup.AutoScalingPolicyProperty.TargetTrackingConfiguration`.
	TargetTrackingConfiguration interface{} `json:"targetTrackingConfiguration"`
	// `CfnGameServerGroup.AutoScalingPolicyProperty.EstimatedInstanceWarmup`.
	EstimatedInstanceWarmup *float64 `json:"estimatedInstanceWarmup"`
}

type CfnGameServerGroup_InstanceDefinitionProperty struct {
	// `CfnGameServerGroup.InstanceDefinitionProperty.InstanceType`.
	InstanceType *string `json:"instanceType"`
	// `CfnGameServerGroup.InstanceDefinitionProperty.WeightedCapacity`.
	WeightedCapacity *string `json:"weightedCapacity"`
}

type CfnGameServerGroup_LaunchTemplateProperty struct {
	// `CfnGameServerGroup.LaunchTemplateProperty.LaunchTemplateId`.
	LaunchTemplateId *string `json:"launchTemplateId"`
	// `CfnGameServerGroup.LaunchTemplateProperty.LaunchTemplateName`.
	LaunchTemplateName *string `json:"launchTemplateName"`
	// `CfnGameServerGroup.LaunchTemplateProperty.Version`.
	Version *string `json:"version"`
}

type CfnGameServerGroup_TargetTrackingConfigurationProperty struct {
	// `CfnGameServerGroup.TargetTrackingConfigurationProperty.TargetValue`.
	TargetValue *float64 `json:"targetValue"`
}

// Properties for defining a `AWS::GameLift::GameServerGroup`.
type CfnGameServerGroupProps struct {
	// `AWS::GameLift::GameServerGroup.GameServerGroupName`.
	GameServerGroupName *string `json:"gameServerGroupName"`
	// `AWS::GameLift::GameServerGroup.InstanceDefinitions`.
	InstanceDefinitions interface{} `json:"instanceDefinitions"`
	// `AWS::GameLift::GameServerGroup.LaunchTemplate`.
	LaunchTemplate interface{} `json:"launchTemplate"`
	// `AWS::GameLift::GameServerGroup.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::GameLift::GameServerGroup.AutoScalingPolicy`.
	AutoScalingPolicy interface{} `json:"autoScalingPolicy"`
	// `AWS::GameLift::GameServerGroup.BalancingStrategy`.
	BalancingStrategy *string `json:"balancingStrategy"`
	// `AWS::GameLift::GameServerGroup.DeleteOption`.
	DeleteOption *string `json:"deleteOption"`
	// `AWS::GameLift::GameServerGroup.GameServerProtectionPolicy`.
	GameServerProtectionPolicy *string `json:"gameServerProtectionPolicy"`
	// `AWS::GameLift::GameServerGroup.MaxSize`.
	MaxSize *float64 `json:"maxSize"`
	// `AWS::GameLift::GameServerGroup.MinSize`.
	MinSize *float64 `json:"minSize"`
	// `AWS::GameLift::GameServerGroup.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::GameLift::GameServerGroup.VpcSubnets`.
	VpcSubnets *[]*string `json:"vpcSubnets"`
}

// A CloudFormation `AWS::GameLift::GameSessionQueue`.
type CfnGameSessionQueue interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Destinations() interface{}
	SetDestinations(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	PlayerLatencyPolicies() interface{}
	SetPlayerLatencyPolicies(val interface{})
	Ref() *string
	Stack() awscdk.Stack
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnGameSessionQueue) Destinations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinations",
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

func (j *jsiiProxy_CfnGameSessionQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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
func NewCfnGameSessionQueue(scope constructs.Construct, id *string, props *CfnGameSessionQueueProps) CfnGameSessionQueue {
	_init_.Initialize()

	j := jsiiProxy_CfnGameSessionQueue{}

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnGameSessionQueue",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::GameSessionQueue`.
func NewCfnGameSessionQueue_Override(c CfnGameSessionQueue, scope constructs.Construct, id *string, props *CfnGameSessionQueueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnGameSessionQueue",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGameSessionQueue) SetDestinations(val interface{}) {
	_jsii_.Set(
		j,
		"destinations",
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

func (j *jsiiProxy_CfnGameSessionQueue) SetPlayerLatencyPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"playerLatencyPolicies",
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
		"aws-cdk-lib.aws_gamelift.CfnGameSessionQueue",
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
		"aws-cdk-lib.aws_gamelift.CfnGameSessionQueue",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnGameSessionQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnGameSessionQueue",
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
		"aws-cdk-lib.aws_gamelift.CfnGameSessionQueue",
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Experimental.
func (c *jsiiProxy_CfnGameSessionQueue) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnGameSessionQueue_DestinationProperty struct {
	// `CfnGameSessionQueue.DestinationProperty.DestinationArn`.
	DestinationArn *string `json:"destinationArn"`
}

type CfnGameSessionQueue_PlayerLatencyPolicyProperty struct {
	// `CfnGameSessionQueue.PlayerLatencyPolicyProperty.MaximumIndividualPlayerLatencyMilliseconds`.
	MaximumIndividualPlayerLatencyMilliseconds *float64 `json:"maximumIndividualPlayerLatencyMilliseconds"`
	// `CfnGameSessionQueue.PlayerLatencyPolicyProperty.PolicyDurationSeconds`.
	PolicyDurationSeconds *float64 `json:"policyDurationSeconds"`
}

// Properties for defining a `AWS::GameLift::GameSessionQueue`.
type CfnGameSessionQueueProps struct {
	// `AWS::GameLift::GameSessionQueue.Name`.
	Name *string `json:"name"`
	// `AWS::GameLift::GameSessionQueue.Destinations`.
	Destinations interface{} `json:"destinations"`
	// `AWS::GameLift::GameSessionQueue.PlayerLatencyPolicies`.
	PlayerLatencyPolicies interface{} `json:"playerLatencyPolicies"`
	// `AWS::GameLift::GameSessionQueue.TimeoutInSeconds`.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds"`
}

// A CloudFormation `AWS::GameLift::MatchmakingConfiguration`.
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
	Node() constructs.Node
	NotificationTarget() *string
	SetNotificationTarget(val *string)
	Ref() *string
	RequestTimeoutSeconds() *float64
	SetRequestTimeoutSeconds(val *float64)
	RuleSetName() *string
	SetRuleSetName(val *string)
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnMatchmakingConfiguration) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnMatchmakingConfiguration(scope constructs.Construct, id *string, props *CfnMatchmakingConfigurationProps) CfnMatchmakingConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnMatchmakingConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnMatchmakingConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::MatchmakingConfiguration`.
func NewCfnMatchmakingConfiguration_Override(c CfnMatchmakingConfiguration, scope constructs.Construct, id *string, props *CfnMatchmakingConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnMatchmakingConfiguration",
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
		"aws-cdk-lib.aws_gamelift.CfnMatchmakingConfiguration",
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
		"aws-cdk-lib.aws_gamelift.CfnMatchmakingConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnMatchmakingConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnMatchmakingConfiguration",
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
		"aws-cdk-lib.aws_gamelift.CfnMatchmakingConfiguration",
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Experimental.
func (c *jsiiProxy_CfnMatchmakingConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnMatchmakingConfiguration_GamePropertyProperty struct {
	// `CfnMatchmakingConfiguration.GamePropertyProperty.Key`.
	Key *string `json:"key"`
	// `CfnMatchmakingConfiguration.GamePropertyProperty.Value`.
	Value *string `json:"value"`
}

// Properties for defining a `AWS::GameLift::MatchmakingConfiguration`.
type CfnMatchmakingConfigurationProps struct {
	// `AWS::GameLift::MatchmakingConfiguration.AcceptanceRequired`.
	AcceptanceRequired interface{} `json:"acceptanceRequired"`
	// `AWS::GameLift::MatchmakingConfiguration.Name`.
	Name *string `json:"name"`
	// `AWS::GameLift::MatchmakingConfiguration.RequestTimeoutSeconds`.
	RequestTimeoutSeconds *float64 `json:"requestTimeoutSeconds"`
	// `AWS::GameLift::MatchmakingConfiguration.RuleSetName`.
	RuleSetName *string `json:"ruleSetName"`
	// `AWS::GameLift::MatchmakingConfiguration.AcceptanceTimeoutSeconds`.
	AcceptanceTimeoutSeconds *float64 `json:"acceptanceTimeoutSeconds"`
	// `AWS::GameLift::MatchmakingConfiguration.AdditionalPlayerCount`.
	AdditionalPlayerCount *float64 `json:"additionalPlayerCount"`
	// `AWS::GameLift::MatchmakingConfiguration.BackfillMode`.
	BackfillMode *string `json:"backfillMode"`
	// `AWS::GameLift::MatchmakingConfiguration.CustomEventData`.
	CustomEventData *string `json:"customEventData"`
	// `AWS::GameLift::MatchmakingConfiguration.Description`.
	Description *string `json:"description"`
	// `AWS::GameLift::MatchmakingConfiguration.FlexMatchMode`.
	FlexMatchMode *string `json:"flexMatchMode"`
	// `AWS::GameLift::MatchmakingConfiguration.GameProperties`.
	GameProperties interface{} `json:"gameProperties"`
	// `AWS::GameLift::MatchmakingConfiguration.GameSessionData`.
	GameSessionData *string `json:"gameSessionData"`
	// `AWS::GameLift::MatchmakingConfiguration.GameSessionQueueArns`.
	GameSessionQueueArns *[]*string `json:"gameSessionQueueArns"`
	// `AWS::GameLift::MatchmakingConfiguration.NotificationTarget`.
	NotificationTarget *string `json:"notificationTarget"`
}

// A CloudFormation `AWS::GameLift::MatchmakingRuleSet`.
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
	Node() constructs.Node
	Ref() *string
	RuleSetBody() *string
	SetRuleSetBody(val *string)
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnMatchmakingRuleSet) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnMatchmakingRuleSet(scope constructs.Construct, id *string, props *CfnMatchmakingRuleSetProps) CfnMatchmakingRuleSet {
	_init_.Initialize()

	j := jsiiProxy_CfnMatchmakingRuleSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnMatchmakingRuleSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::MatchmakingRuleSet`.
func NewCfnMatchmakingRuleSet_Override(c CfnMatchmakingRuleSet, scope constructs.Construct, id *string, props *CfnMatchmakingRuleSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnMatchmakingRuleSet",
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
		"aws-cdk-lib.aws_gamelift.CfnMatchmakingRuleSet",
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
		"aws-cdk-lib.aws_gamelift.CfnMatchmakingRuleSet",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnMatchmakingRuleSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnMatchmakingRuleSet",
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
		"aws-cdk-lib.aws_gamelift.CfnMatchmakingRuleSet",
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Experimental.
func (c *jsiiProxy_CfnMatchmakingRuleSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::GameLift::MatchmakingRuleSet`.
type CfnMatchmakingRuleSetProps struct {
	// `AWS::GameLift::MatchmakingRuleSet.Name`.
	Name *string `json:"name"`
	// `AWS::GameLift::MatchmakingRuleSet.RuleSetBody`.
	RuleSetBody *string `json:"ruleSetBody"`
}

// A CloudFormation `AWS::GameLift::Script`.
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnScript) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnScript(scope constructs.Construct, id *string, props *CfnScriptProps) CfnScript {
	_init_.Initialize()

	j := jsiiProxy_CfnScript{}

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnScript",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::Script`.
func NewCfnScript_Override(c CfnScript, scope constructs.Construct, id *string, props *CfnScriptProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnScript",
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
		"aws-cdk-lib.aws_gamelift.CfnScript",
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
		"aws-cdk-lib.aws_gamelift.CfnScript",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnScript_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_gamelift.CfnScript",
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
		"aws-cdk-lib.aws_gamelift.CfnScript",
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnScript) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Experimental.
func (c *jsiiProxy_CfnScript) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnScript_S3LocationProperty struct {
	// `CfnScript.S3LocationProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnScript.S3LocationProperty.Key`.
	Key *string `json:"key"`
	// `CfnScript.S3LocationProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnScript.S3LocationProperty.ObjectVersion`.
	ObjectVersion *string `json:"objectVersion"`
}

// Properties for defining a `AWS::GameLift::Script`.
type CfnScriptProps struct {
	// `AWS::GameLift::Script.StorageLocation`.
	StorageLocation interface{} `json:"storageLocation"`
	// `AWS::GameLift::Script.Name`.
	Name *string `json:"name"`
	// `AWS::GameLift::Script.Version`.
	Version *string `json:"version"`
}

