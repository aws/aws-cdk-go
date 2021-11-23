package awsworkspaces

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspaces/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::WorkSpaces::ConnectionAlias`.
//
// TODO: EXAMPLE
//
type CfnConnectionAlias interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAliasId() *string
	AttrAssociations() awscdk.IResolvable
	AttrConnectionAliasState() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConnectionString() *string
	SetConnectionString(val *string)
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnConnectionAlias
type jsiiProxy_CfnConnectionAlias struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConnectionAlias) AttrAliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAliasId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) AttrAssociations() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrAssociations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) AttrConnectionAliasState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConnectionAliasState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) ConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionAlias) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WorkSpaces::ConnectionAlias`.
func NewCfnConnectionAlias(scope constructs.Construct, id *string, props *CfnConnectionAliasProps) CfnConnectionAlias {
	_init_.Initialize()

	j := jsiiProxy_CfnConnectionAlias{}

	_jsii_.Create(
		"aws-cdk-lib.aws_workspaces.CfnConnectionAlias",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WorkSpaces::ConnectionAlias`.
func NewCfnConnectionAlias_Override(c CfnConnectionAlias, scope constructs.Construct, id *string, props *CfnConnectionAliasProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_workspaces.CfnConnectionAlias",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConnectionAlias) SetConnectionString(val *string) {
	_jsii_.Set(
		j,
		"connectionString",
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
func CfnConnectionAlias_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspaces.CfnConnectionAlias",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnConnectionAlias_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspaces.CfnConnectionAlias",
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
func CfnConnectionAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspaces.CfnConnectionAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectionAlias_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_workspaces.CfnConnectionAlias",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnConnectionAlias) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnConnectionAlias) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnConnectionAlias) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnConnectionAlias) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnConnectionAlias) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnConnectionAlias) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnConnectionAlias) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnConnectionAlias) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnConnectionAlias) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnConnectionAlias) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnConnectionAlias) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConnectionAlias) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnConnectionAlias) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnConnectionAlias) ToString() *string {
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
func (c *jsiiProxy_CfnConnectionAlias) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnConnectionAlias_ConnectionAliasAssociationProperty struct {
	// `CfnConnectionAlias.ConnectionAliasAssociationProperty.AssociatedAccountId`.
	AssociatedAccountId *string `json:"associatedAccountId"`
	// `CfnConnectionAlias.ConnectionAliasAssociationProperty.AssociationStatus`.
	AssociationStatus *string `json:"associationStatus"`
	// `CfnConnectionAlias.ConnectionAliasAssociationProperty.ConnectionIdentifier`.
	ConnectionIdentifier *string `json:"connectionIdentifier"`
	// `CfnConnectionAlias.ConnectionAliasAssociationProperty.ResourceId`.
	ResourceId *string `json:"resourceId"`
}

// Properties for defining a `AWS::WorkSpaces::ConnectionAlias`.
//
// TODO: EXAMPLE
//
type CfnConnectionAliasProps struct {
	// `AWS::WorkSpaces::ConnectionAlias.ConnectionString`.
	ConnectionString *string `json:"connectionString"`
	// `AWS::WorkSpaces::ConnectionAlias.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::WorkSpaces::Workspace`.
//
// TODO: EXAMPLE
//
type CfnWorkspace interface {
	awscdk.CfnResource
	awscdk.IInspectable
	BundleId() *string
	SetBundleId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DirectoryId() *string
	SetDirectoryId(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RootVolumeEncryptionEnabled() interface{}
	SetRootVolumeEncryptionEnabled(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	UserName() *string
	SetUserName(val *string)
	UserVolumeEncryptionEnabled() interface{}
	SetUserVolumeEncryptionEnabled(val interface{})
	VolumeEncryptionKey() *string
	SetVolumeEncryptionKey(val *string)
	WorkspaceProperties() interface{}
	SetWorkspaceProperties(val interface{})
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

// The jsii proxy struct for CfnWorkspace
type jsiiProxy_CfnWorkspace struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWorkspace) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) RootVolumeEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootVolumeEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) UserVolumeEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userVolumeEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) VolumeEncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) WorkspaceProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::WorkSpaces::Workspace`.
func NewCfnWorkspace(scope constructs.Construct, id *string, props *CfnWorkspaceProps) CfnWorkspace {
	_init_.Initialize()

	j := jsiiProxy_CfnWorkspace{}

	_jsii_.Create(
		"aws-cdk-lib.aws_workspaces.CfnWorkspace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WorkSpaces::Workspace`.
func NewCfnWorkspace_Override(c CfnWorkspace, scope constructs.Construct, id *string, props *CfnWorkspaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_workspaces.CfnWorkspace",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWorkspace) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace) SetDirectoryId(val *string) {
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace) SetRootVolumeEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"rootVolumeEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace) SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace) SetUserVolumeEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"userVolumeEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace) SetVolumeEncryptionKey(val *string) {
	_jsii_.Set(
		j,
		"volumeEncryptionKey",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace) SetWorkspaceProperties(val interface{}) {
	_jsii_.Set(
		j,
		"workspaceProperties",
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
func CfnWorkspace_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspaces.CfnWorkspace",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnWorkspace_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspaces.CfnWorkspace",
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
func CfnWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspaces.CfnWorkspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkspace_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_workspaces.CfnWorkspace",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnWorkspace) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnWorkspace) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnWorkspace) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnWorkspace) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnWorkspace) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnWorkspace) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnWorkspace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnWorkspace) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnWorkspace) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnWorkspace) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnWorkspace) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWorkspace) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnWorkspace) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnWorkspace) ToString() *string {
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
func (c *jsiiProxy_CfnWorkspace) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnWorkspace_WorkspacePropertiesProperty struct {
	// `CfnWorkspace.WorkspacePropertiesProperty.ComputeTypeName`.
	ComputeTypeName *string `json:"computeTypeName"`
	// `CfnWorkspace.WorkspacePropertiesProperty.RootVolumeSizeGib`.
	RootVolumeSizeGib *float64 `json:"rootVolumeSizeGib"`
	// `CfnWorkspace.WorkspacePropertiesProperty.RunningMode`.
	RunningMode *string `json:"runningMode"`
	// `CfnWorkspace.WorkspacePropertiesProperty.RunningModeAutoStopTimeoutInMinutes`.
	RunningModeAutoStopTimeoutInMinutes *float64 `json:"runningModeAutoStopTimeoutInMinutes"`
	// `CfnWorkspace.WorkspacePropertiesProperty.UserVolumeSizeGib`.
	UserVolumeSizeGib *float64 `json:"userVolumeSizeGib"`
}

// Properties for defining a `AWS::WorkSpaces::Workspace`.
//
// TODO: EXAMPLE
//
type CfnWorkspaceProps struct {
	// `AWS::WorkSpaces::Workspace.BundleId`.
	BundleId *string `json:"bundleId"`
	// `AWS::WorkSpaces::Workspace.DirectoryId`.
	DirectoryId *string `json:"directoryId"`
	// `AWS::WorkSpaces::Workspace.RootVolumeEncryptionEnabled`.
	RootVolumeEncryptionEnabled interface{} `json:"rootVolumeEncryptionEnabled"`
	// `AWS::WorkSpaces::Workspace.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::WorkSpaces::Workspace.UserName`.
	UserName *string `json:"userName"`
	// `AWS::WorkSpaces::Workspace.UserVolumeEncryptionEnabled`.
	UserVolumeEncryptionEnabled interface{} `json:"userVolumeEncryptionEnabled"`
	// `AWS::WorkSpaces::Workspace.VolumeEncryptionKey`.
	VolumeEncryptionKey *string `json:"volumeEncryptionKey"`
	// `AWS::WorkSpaces::Workspace.WorkspaceProperties`.
	WorkspaceProperties interface{} `json:"workspaceProperties"`
}

