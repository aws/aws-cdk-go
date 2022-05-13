package awsmediapackage

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackage/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MediaPackage::Asset`.
//
// Creates an asset to ingest VOD content.
//
// After it's created, the asset starts ingesting content and generates playback URLs for the packaging configurations associated with it. When ingest is complete, downstream devices use the appropriate URL to request VOD content from AWS Elemental MediaPackage .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAsset := awscdk.Aws_mediapackage.NewCfnAsset(this, jsii.String("MyCfnAsset"), &cfnAssetProps{
//   	id: jsii.String("id"),
//   	packagingGroupId: jsii.String("packagingGroupId"),
//   	sourceArn: jsii.String("sourceArn"),
//   	sourceRoleArn: jsii.String("sourceRoleArn"),
//
//   	// the properties below are optional
//   	resourceId: jsii.String("resourceId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnAsset interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) for the asset.
	//
	// You can get this from the response to any request to
	// the asset.
	AttrArn() *string
	// The time that the asset was initially submitted for ingest.
	AttrCreatedAt() *string
	AttrEgressEndpoints() awscdk.IResolvable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Unique identifier that you assign to the asset.
	Id() *string
	SetId(val *string)
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
	// The tree node.
	Node() constructs.Node
	// The ID of the packaging group associated with this asset.
	PackagingGroupId() *string
	SetPackagingGroupId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Unique identifier for this asset, as it's configured in the key provider service.
	ResourceId() *string
	SetResourceId(val *string)
	// The ARN for the source content in Amazon S3.
	SourceArn() *string
	SetSourceArn(val *string)
	// The ARN for the IAM role that provides AWS Elemental MediaPackage access to the Amazon S3 bucket where the source content is stored.
	//
	// Valid format: arn:aws:iam::{accountID}:role/{name}.
	SourceRoleArn() *string
	SetSourceRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to assign to the asset.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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
}

// The jsii proxy struct for CfnAsset
type jsiiProxy_CfnAsset struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAsset) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) AttrEgressEndpoints() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrEgressEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) PackagingGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packagingGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) SourceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAsset) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackage::Asset`.
func NewCfnAsset(scope constructs.Construct, id *string, props *CfnAssetProps) CfnAsset {
	_init_.Initialize()

	j := jsiiProxy_CfnAsset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackage.CfnAsset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackage::Asset`.
func NewCfnAsset_Override(c CfnAsset, scope constructs.Construct, id *string, props *CfnAssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackage.CfnAsset",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAsset) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CfnAsset) SetPackagingGroupId(val *string) {
	_jsii_.Set(
		j,
		"packagingGroupId",
		val,
	)
}

func (j *jsiiProxy_CfnAsset) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_CfnAsset) SetSourceArn(val *string) {
	_jsii_.Set(
		j,
		"sourceArn",
		val,
	)
}

func (j *jsiiProxy_CfnAsset) SetSourceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"sourceRoleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAsset_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnAsset",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAsset_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnAsset",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAsset_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediapackage.CfnAsset",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAsset) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAsset) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAsset) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAsset) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAsset) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAsset) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAsset) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAsset) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAsset) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAsset) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAsset) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAsset) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAsset) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAsset) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The playback endpoint for a packaging configuration on an asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   egressEndpointProperty := &egressEndpointProperty{
//   	packagingConfigurationId: jsii.String("packagingConfigurationId"),
//   	url: jsii.String("url"),
//   }
//
type CfnAsset_EgressEndpointProperty struct {
	// The ID of a packaging configuration that's applied to this asset.
	PackagingConfigurationId *string `field:"required" json:"packagingConfigurationId" yaml:"packagingConfigurationId"`
	// The URL that's used to request content from this endpoint.
	Url *string `field:"required" json:"url" yaml:"url"`
}

// Properties for defining a `CfnAsset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssetProps := &cfnAssetProps{
//   	id: jsii.String("id"),
//   	packagingGroupId: jsii.String("packagingGroupId"),
//   	sourceArn: jsii.String("sourceArn"),
//   	sourceRoleArn: jsii.String("sourceRoleArn"),
//
//   	// the properties below are optional
//   	resourceId: jsii.String("resourceId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAssetProps struct {
	// Unique identifier that you assign to the asset.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ID of the packaging group associated with this asset.
	PackagingGroupId *string `field:"required" json:"packagingGroupId" yaml:"packagingGroupId"`
	// The ARN for the source content in Amazon S3.
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// The ARN for the IAM role that provides AWS Elemental MediaPackage access to the Amazon S3 bucket where the source content is stored.
	//
	// Valid format: arn:aws:iam::{accountID}:role/{name}.
	SourceRoleArn *string `field:"required" json:"sourceRoleArn" yaml:"sourceRoleArn"`
	// Unique identifier for this asset, as it's configured in the key provider service.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The tags to assign to the asset.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::MediaPackage::Channel`.
//
// Creates a channel to receive content.
//
// After it's created, a channel provides static input URLs. These URLs remain the same throughout the lifetime of the channel, regardless of any failures or upgrades that might occur. Use these URLs to configure the outputs of your upstream encoder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannel := awscdk.Aws_mediapackage.NewCfnChannel(this, jsii.String("MyCfnChannel"), &cfnChannelProps{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	egressAccessLogs: &logConfigurationProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	ingressAccessLogs: &logConfigurationProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The channel's unique system-generated resource name, based on the AWS record.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Any descriptive information that you want to add to the channel for future identification purposes.
	Description() *string
	SetDescription(val *string)
	// Configures egress access logs.
	EgressAccessLogs() interface{}
	SetEgressAccessLogs(val interface{})
	// Unique identifier that you assign to the channel.
	Id() *string
	SetId(val *string)
	// Configures ingress access logs.
	IngressAccessLogs() interface{}
	SetIngressAccessLogs(val interface{})
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
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to assign to the channel.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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
}

// The jsii proxy struct for CfnChannel
type jsiiProxy_CfnChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnChannel) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) EgressAccessLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressAccessLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) IngressAccessLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressAccessLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackage::Channel`.
func NewCfnChannel(scope constructs.Construct, id *string, props *CfnChannelProps) CfnChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackage.CfnChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackage::Channel`.
func NewCfnChannel_Override(c CfnChannel, scope constructs.Construct, id *string, props *CfnChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackage.CfnChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnChannel) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetEgressAccessLogs(val interface{}) {
	_jsii_.Set(
		j,
		"egressAccessLogs",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetIngressAccessLogs(val interface{}) {
	_jsii_.Set(
		j,
		"ingressAccessLogs",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnChannel",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediapackage.CfnChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The access log configuration parameters for your channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &logConfigurationProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   }
//
type CfnChannel_LogConfigurationProperty struct {
	// Sets a custom Amazon CloudWatch log group name.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
}

// Properties for defining a `CfnChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelProps := &cfnChannelProps{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	egressAccessLogs: &logConfigurationProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	ingressAccessLogs: &logConfigurationProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnChannelProps struct {
	// Unique identifier that you assign to the channel.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Any descriptive information that you want to add to the channel for future identification purposes.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configures egress access logs.
	EgressAccessLogs interface{} `field:"optional" json:"egressAccessLogs" yaml:"egressAccessLogs"`
	// Configures ingress access logs.
	IngressAccessLogs interface{} `field:"optional" json:"ingressAccessLogs" yaml:"ingressAccessLogs"`
	// The tags to assign to the channel.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::MediaPackage::OriginEndpoint`.
//
// Create an endpoint on an AWS Elemental MediaPackage channel.
//
// An endpoint represents a single delivery point of a channel, and defines content output handling through various components, such as packaging protocols, DRM and encryption integration, and more.
//
// After it's created, an endpoint provides a fixed public URL. This URL remains the same throughout the lifetime of the endpoint, regardless of any failures or upgrades that might occur. Integrate the URL with a downstream CDN (such as Amazon CloudFront) or playback device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOriginEndpoint := awscdk.Aws_mediapackage.NewCfnOriginEndpoint(this, jsii.String("MyCfnOriginEndpoint"), &cfnOriginEndpointProps{
//   	channelId: jsii.String("channelId"),
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	authorization: &authorizationProperty{
//   		cdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   		secretsRoleArn: jsii.String("secretsRoleArn"),
//   	},
//   	cmafPackage: &cmafPackageProperty{
//   		encryption: &cmafEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			constantInitializationVector: jsii.String("constantInitializationVector"),
//   			keyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		hlsManifests: []interface{}{
//   			&hlsManifestProperty{
//   				id: jsii.String("id"),
//
//   				// the properties below are optional
//   				adMarkers: jsii.String("adMarkers"),
//   				adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   				adTriggers: []*string{
//   					jsii.String("adTriggers"),
//   				},
//   				includeIframeOnlyStream: jsii.Boolean(false),
//   				manifestName: jsii.String("manifestName"),
//   				playlistType: jsii.String("playlistType"),
//   				playlistWindowSeconds: jsii.Number(123),
//   				programDateTimeIntervalSeconds: jsii.Number(123),
//   				url: jsii.String("url"),
//   			},
//   		},
//   		segmentDurationSeconds: jsii.Number(123),
//   		segmentPrefix: jsii.String("segmentPrefix"),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   	},
//   	dashPackage: &dashPackageProperty{
//   		adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   		adTriggers: []*string{
//   			jsii.String("adTriggers"),
//   		},
//   		encryption: &dashEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			keyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		manifestLayout: jsii.String("manifestLayout"),
//   		manifestWindowSeconds: jsii.Number(123),
//   		minBufferTimeSeconds: jsii.Number(123),
//   		minUpdatePeriodSeconds: jsii.Number(123),
//   		periodTriggers: []*string{
//   			jsii.String("periodTriggers"),
//   		},
//   		profile: jsii.String("profile"),
//   		segmentDurationSeconds: jsii.Number(123),
//   		segmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   		suggestedPresentationDelaySeconds: jsii.Number(123),
//   		utcTiming: jsii.String("utcTiming"),
//   		utcTimingUri: jsii.String("utcTimingUri"),
//   	},
//   	description: jsii.String("description"),
//   	hlsPackage: &hlsPackageProperty{
//   		adMarkers: jsii.String("adMarkers"),
//   		adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   		adTriggers: []*string{
//   			jsii.String("adTriggers"),
//   		},
//   		encryption: &hlsEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			constantInitializationVector: jsii.String("constantInitializationVector"),
//   			encryptionMethod: jsii.String("encryptionMethod"),
//   			keyRotationIntervalSeconds: jsii.Number(123),
//   			repeatExtXKey: jsii.Boolean(false),
//   		},
//   		includeIframeOnlyStream: jsii.Boolean(false),
//   		playlistType: jsii.String("playlistType"),
//   		playlistWindowSeconds: jsii.Number(123),
//   		programDateTimeIntervalSeconds: jsii.Number(123),
//   		segmentDurationSeconds: jsii.Number(123),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   		useAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	manifestName: jsii.String("manifestName"),
//   	mssPackage: &mssPackageProperty{
//   		encryption: &mssEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//   		},
//   		manifestWindowSeconds: jsii.Number(123),
//   		segmentDurationSeconds: jsii.Number(123),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   	},
//   	origination: jsii.String("origination"),
//   	startoverWindowSeconds: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeDelaySeconds: jsii.Number(123),
//   	whitelist: []*string{
//   		jsii.String("whitelist"),
//   	},
//   })
//
type CfnOriginEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The endpoint's unique system-generated resource name, based on the AWS record.
	AttrArn() *string
	// URL for the key provider’s key retrieval API endpoint.
	//
	// Must start with https://.
	AttrUrl() *string
	// Parameters for CDN authorization.
	Authorization() interface{}
	SetAuthorization(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The ID of the channel associated with this endpoint.
	ChannelId() *string
	SetChannelId(val *string)
	// Parameters for Common Media Application Format (CMAF) packaging.
	CmafPackage() interface{}
	SetCmafPackage(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Parameters for DASH packaging.
	DashPackage() interface{}
	SetDashPackage(val interface{})
	// Any descriptive information that you want to add to the endpoint for future identification purposes.
	Description() *string
	SetDescription(val *string)
	// Parameters for Apple HLS packaging.
	HlsPackage() interface{}
	SetHlsPackage(val interface{})
	// The manifest ID is required and must be unique within the OriginEndpoint.
	//
	// The ID can't be changed after the endpoint is created.
	Id() *string
	SetId(val *string)
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
	// A short string that's appended to the end of the endpoint URL to create a unique path to this endpoint.
	ManifestName() *string
	SetManifestName(val *string)
	// Parameters for Microsoft Smooth Streaming packaging.
	MssPackage() interface{}
	SetMssPackage(val interface{})
	// The tree node.
	Node() constructs.Node
	// Controls video origination from this endpoint.
	//
	// - `ALLOW` - enables this endpoint to serve content to requesting devices.
	// - `DENY` - prevents this endpoint from serving content. Denying origination is helpful for harvesting live-to-VOD assets. For more information about harvesting and origination, see [Live-to-VOD Requirements](https://docs.aws.amazon.com/mediapackage/latest/ug/ltov-reqmts.html) .
	Origination() *string
	SetOrigination(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Maximum duration (seconds) of content to retain for startover playback.
	//
	// Omit this attribute or enter `0` to indicate that startover playback is disabled for this endpoint.
	StartoverWindowSeconds() *float64
	SetStartoverWindowSeconds(val *float64)
	// The tags to assign to the endpoint.
	Tags() awscdk.TagManager
	// Minimum duration (seconds) of delay to enforce on the playback of live content.
	//
	// Omit this attribute or enter `0` to indicate that there is no time delay in effect for this endpoint.
	TimeDelaySeconds() *float64
	SetTimeDelaySeconds(val *float64)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The IP addresses that can access this endpoint.
	Whitelist() *[]*string
	SetWhitelist(val *[]*string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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
}

// The jsii proxy struct for CfnOriginEndpoint
type jsiiProxy_CfnOriginEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnOriginEndpoint) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) AttrUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Authorization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) ChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CmafPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cmafPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) DashPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) HlsPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) ManifestName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) MssPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mssPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Origination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"origination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) StartoverWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startoverWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) TimeDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpoint) Whitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelist",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackage::OriginEndpoint`.
func NewCfnOriginEndpoint(scope constructs.Construct, id *string, props *CfnOriginEndpointProps) CfnOriginEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CfnOriginEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackage.CfnOriginEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackage::OriginEndpoint`.
func NewCfnOriginEndpoint_Override(c CfnOriginEndpoint, scope constructs.Construct, id *string, props *CfnOriginEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackage.CfnOriginEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetAuthorization(val interface{}) {
	_jsii_.Set(
		j,
		"authorization",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetChannelId(val *string) {
	_jsii_.Set(
		j,
		"channelId",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetCmafPackage(val interface{}) {
	_jsii_.Set(
		j,
		"cmafPackage",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetDashPackage(val interface{}) {
	_jsii_.Set(
		j,
		"dashPackage",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetHlsPackage(val interface{}) {
	_jsii_.Set(
		j,
		"hlsPackage",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetManifestName(val *string) {
	_jsii_.Set(
		j,
		"manifestName",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetMssPackage(val interface{}) {
	_jsii_.Set(
		j,
		"mssPackage",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetOrigination(val *string) {
	_jsii_.Set(
		j,
		"origination",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetStartoverWindowSeconds(val *float64) {
	_jsii_.Set(
		j,
		"startoverWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetTimeDelaySeconds(val *float64) {
	_jsii_.Set(
		j,
		"timeDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_CfnOriginEndpoint) SetWhitelist(val *[]*string) {
	_jsii_.Set(
		j,
		"whitelist",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnOriginEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnOriginEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnOriginEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnOriginEndpoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnOriginEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnOriginEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOriginEndpoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediapackage.CfnOriginEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnOriginEndpoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOriginEndpoint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Parameters for enabling CDN authorization on the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationProperty := &authorizationProperty{
//   	cdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   	secretsRoleArn: jsii.String("secretsRoleArn"),
//   }
//
type CfnOriginEndpoint_AuthorizationProperty struct {
	// The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that your Content Distribution Network (CDN) uses for authorization to access your endpoint.
	CdnIdentifierSecret *string `field:"required" json:"cdnIdentifierSecret" yaml:"cdnIdentifierSecret"`
	// The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager .
	SecretsRoleArn *string `field:"required" json:"secretsRoleArn" yaml:"secretsRoleArn"`
}

// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cmafEncryptionProperty := &cmafEncryptionProperty{
//   	spekeKeyProvider: &spekeKeyProviderProperty{
//   		resourceId: jsii.String("resourceId"),
//   		roleArn: jsii.String("roleArn"),
//   		systemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		url: jsii.String("url"),
//
//   		// the properties below are optional
//   		certificateArn: jsii.String("certificateArn"),
//   		encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   			presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   			presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	constantInitializationVector: jsii.String("constantInitializationVector"),
//   	keyRotationIntervalSeconds: jsii.Number(123),
//   }
//
type CfnOriginEndpoint_CmafEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// An optional 128-bit, 16-byte hex value represented by a 32-character string, used in conjunction with the key for encrypting blocks.
	//
	// If you don't specify a value, then MediaPackage creates the constant initialization vector (IV).
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// Number of seconds before AWS Elemental MediaPackage rotates to a new key.
	//
	// By default, rotation is set to 60 seconds. Set to `0` to disable key rotation.
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
}

// Parameters for Common Media Application Format (CMAF) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cmafPackageProperty := &cmafPackageProperty{
//   	encryption: &cmafEncryptionProperty{
//   		spekeKeyProvider: &spekeKeyProviderProperty{
//   			resourceId: jsii.String("resourceId"),
//   			roleArn: jsii.String("roleArn"),
//   			systemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			url: jsii.String("url"),
//
//   			// the properties below are optional
//   			certificateArn: jsii.String("certificateArn"),
//   			encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   				presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   				presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		constantInitializationVector: jsii.String("constantInitializationVector"),
//   		keyRotationIntervalSeconds: jsii.Number(123),
//   	},
//   	hlsManifests: []interface{}{
//   		&hlsManifestProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			adMarkers: jsii.String("adMarkers"),
//   			adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   			adTriggers: []*string{
//   				jsii.String("adTriggers"),
//   			},
//   			includeIframeOnlyStream: jsii.Boolean(false),
//   			manifestName: jsii.String("manifestName"),
//   			playlistType: jsii.String("playlistType"),
//   			playlistWindowSeconds: jsii.Number(123),
//   			programDateTimeIntervalSeconds: jsii.Number(123),
//   			url: jsii.String("url"),
//   		},
//   	},
//   	segmentDurationSeconds: jsii.Number(123),
//   	segmentPrefix: jsii.String("segmentPrefix"),
//   	streamSelection: &streamSelectionProperty{
//   		maxVideoBitsPerSecond: jsii.Number(123),
//   		minVideoBitsPerSecond: jsii.Number(123),
//   		streamOrder: jsii.String("streamOrder"),
//   	},
//   }
//
type CfnOriginEndpoint_CmafPackageProperty struct {
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// A list of HLS manifest configurations that are available from this endpoint.
	HlsManifests interface{} `field:"optional" json:"hlsManifests" yaml:"hlsManifests"`
	// Duration (in seconds) of each segment.
	//
	// Actual segments are rounded to the nearest multiple of the source segment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// An optional custom string that is prepended to the name of each segment.
	//
	// If not specified, the segment prefix defaults to the ChannelId.
	SegmentPrefix *string `field:"optional" json:"segmentPrefix" yaml:"segmentPrefix"`
	// Limitations for outputs from the endpoint, based on the video bitrate.
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashEncryptionProperty := &dashEncryptionProperty{
//   	spekeKeyProvider: &spekeKeyProviderProperty{
//   		resourceId: jsii.String("resourceId"),
//   		roleArn: jsii.String("roleArn"),
//   		systemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		url: jsii.String("url"),
//
//   		// the properties below are optional
//   		certificateArn: jsii.String("certificateArn"),
//   		encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   			presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   			presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	keyRotationIntervalSeconds: jsii.Number(123),
//   }
//
type CfnOriginEndpoint_DashEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// Number of seconds before AWS Elemental MediaPackage rotates to a new key.
	//
	// By default, rotation is set to 60 seconds. Set to `0` to disable key rotation.
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
}

// Parameters for DASH packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashPackageProperty := &dashPackageProperty{
//   	adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   	adTriggers: []*string{
//   		jsii.String("adTriggers"),
//   	},
//   	encryption: &dashEncryptionProperty{
//   		spekeKeyProvider: &spekeKeyProviderProperty{
//   			resourceId: jsii.String("resourceId"),
//   			roleArn: jsii.String("roleArn"),
//   			systemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			url: jsii.String("url"),
//
//   			// the properties below are optional
//   			certificateArn: jsii.String("certificateArn"),
//   			encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   				presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   				presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		keyRotationIntervalSeconds: jsii.Number(123),
//   	},
//   	manifestLayout: jsii.String("manifestLayout"),
//   	manifestWindowSeconds: jsii.Number(123),
//   	minBufferTimeSeconds: jsii.Number(123),
//   	minUpdatePeriodSeconds: jsii.Number(123),
//   	periodTriggers: []*string{
//   		jsii.String("periodTriggers"),
//   	},
//   	profile: jsii.String("profile"),
//   	segmentDurationSeconds: jsii.Number(123),
//   	segmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   	streamSelection: &streamSelectionProperty{
//   		maxVideoBitsPerSecond: jsii.Number(123),
//   		minVideoBitsPerSecond: jsii.Number(123),
//   		streamOrder: jsii.String("streamOrder"),
//   	},
//   	suggestedPresentationDelaySeconds: jsii.Number(123),
//   	utcTiming: jsii.String("utcTiming"),
//   	utcTimingUri: jsii.String("utcTimingUri"),
//   }
//
type CfnOriginEndpoint_DashPackageProperty struct {
	// The flags on SCTE-35 segmentation descriptors that have to be present for MediaPackage to insert ad markers in the output manifest.
	//
	// For information about SCTE-35 in MediaPackage, see [SCTE-35 Message Options in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/scte.html) .
	AdsOnDeliveryRestrictions *string `field:"optional" json:"adsOnDeliveryRestrictions" yaml:"adsOnDeliveryRestrictions"`
	// Specifies the SCTE-35 message types that MediaPackage treats as ad markers in the output manifest.
	//
	// Valid values:
	//
	// - `BREAK`
	// - `DISTRIBUTOR_ADVERTISEMENT`
	// - `DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY` .
	// - `DISTRIBUTOR_PLACEMENT_OPPORTUNITY` .
	// - `PROVIDER_ADVERTISEMENT` .
	// - `PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY` .
	// - `PROVIDER_PLACEMENT_OPPORTUNITY` .
	// - `SPLICE_INSERT` .
	AdTriggers *[]*string `field:"optional" json:"adTriggers" yaml:"adTriggers"`
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// Determines the position of some tags in the manifest.
	//
	// Options:
	//
	// - `FULL` - elements like `SegmentTemplate` and `ContentProtection` are included in each `Representation` .
	// - `COMPACT` - duplicate elements are combined and presented at the `AdaptationSet` level.
	ManifestLayout *string `field:"optional" json:"manifestLayout" yaml:"manifestLayout"`
	// Time window (in seconds) contained in each manifest.
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// Minimum amount of content (measured in seconds) that a player must keep available in the buffer.
	MinBufferTimeSeconds *float64 `field:"optional" json:"minBufferTimeSeconds" yaml:"minBufferTimeSeconds"`
	// Minimum amount of time (in seconds) that the player should wait before requesting updates to the manifest.
	MinUpdatePeriodSeconds *float64 `field:"optional" json:"minUpdatePeriodSeconds" yaml:"minUpdatePeriodSeconds"`
	// Controls whether MediaPackage produces single-period or multi-period DASH manifests.
	//
	// For more information about periods, see [Multi-period DASH in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/multi-period.html) .
	//
	// Valid values:
	//
	// - `ADS` - MediaPackage will produce multi-period DASH manifests. Periods are created based on the SCTE-35 ad markers present in the input manifest.
	// - *No value* - MediaPackage will produce single-period DASH manifests. This is the default setting.
	PeriodTriggers *[]*string `field:"optional" json:"periodTriggers" yaml:"periodTriggers"`
	// DASH profile for the output, such as HbbTV.
	//
	// Valid values:
	//
	// - `NONE` - the output doesn't use a DASH profile.
	// - `HBBTV_1_5` - the output is HbbTV-compliant.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// Determines the type of variable used in the `media` URL of the `SegmentTemplate` tag in the manifest.
	//
	// Also specifies if segment timeline information is included in `SegmentTimeline` or `SegmentTemplate` .
	//
	// - `NUMBER_WITH_TIMELINE` - The `$Number$` variable is used in the `media` URL. The value of this variable is the sequential number of the segment. A full `SegmentTimeline` object is presented in each `SegmentTemplate` .
	// - `NUMBER_WITH_DURATION` - The `$Number$` variable is used in the `media` URL and a `duration` attribute is added to the segment template. The `SegmentTimeline` object is removed from the representation.
	// - `TIME_WITH_TIMELINE` - The `$Time$` variable is used in the `media` URL. The value of this variable is the timestamp of when the segment starts. A full `SegmentTimeline` object is presented in each `SegmentTemplate` .
	SegmentTemplateFormat *string `field:"optional" json:"segmentTemplateFormat" yaml:"segmentTemplateFormat"`
	// Limitations for outputs from the endpoint, based on the video bitrate.
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
	// Amount of time (in seconds) that the player should be from the live point at the end of the manifest.
	SuggestedPresentationDelaySeconds *float64 `field:"optional" json:"suggestedPresentationDelaySeconds" yaml:"suggestedPresentationDelaySeconds"`
	// Determines the type of UTC timing included in the DASH Media Presentation Description (MPD).
	UtcTiming *string `field:"optional" json:"utcTiming" yaml:"utcTiming"`
	// Specifies the value attribute of the UTC timing field when utcTiming is set to HTTP-ISO or HTTP-HEAD.
	UtcTimingUri *string `field:"optional" json:"utcTimingUri" yaml:"utcTimingUri"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionContractConfigurationProperty := &encryptionContractConfigurationProperty{
//   	presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   	presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   }
//
type CfnOriginEndpoint_EncryptionContractConfigurationProperty struct {
	// `CfnOriginEndpoint.EncryptionContractConfigurationProperty.PresetSpeke20Audio`.
	PresetSpeke20Audio *string `field:"required" json:"presetSpeke20Audio" yaml:"presetSpeke20Audio"`
	// `CfnOriginEndpoint.EncryptionContractConfigurationProperty.PresetSpeke20Video`.
	PresetSpeke20Video *string `field:"required" json:"presetSpeke20Video" yaml:"presetSpeke20Video"`
}

// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsEncryptionProperty := &hlsEncryptionProperty{
//   	spekeKeyProvider: &spekeKeyProviderProperty{
//   		resourceId: jsii.String("resourceId"),
//   		roleArn: jsii.String("roleArn"),
//   		systemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		url: jsii.String("url"),
//
//   		// the properties below are optional
//   		certificateArn: jsii.String("certificateArn"),
//   		encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   			presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   			presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	constantInitializationVector: jsii.String("constantInitializationVector"),
//   	encryptionMethod: jsii.String("encryptionMethod"),
//   	keyRotationIntervalSeconds: jsii.Number(123),
//   	repeatExtXKey: jsii.Boolean(false),
//   }
//
type CfnOriginEndpoint_HlsEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// A 128-bit, 16-byte hex value represented by a 32-character string, used with the key for encrypting blocks.
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// HLS encryption type.
	EncryptionMethod *string `field:"optional" json:"encryptionMethod" yaml:"encryptionMethod"`
	// Number of seconds before AWS Elemental MediaPackage rotates to a new key.
	//
	// By default, rotation is set to 60 seconds. Set to `0` to disable key rotation.
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
	// Repeat the `EXT-X-KEY` directive for every media segment.
	//
	// This might result in an increase in client requests to the DRM server.
	RepeatExtXKey interface{} `field:"optional" json:"repeatExtXKey" yaml:"repeatExtXKey"`
}

// An HTTP Live Streaming (HLS) manifest configuration on a CMAF endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsManifestProperty := &hlsManifestProperty{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	adMarkers: jsii.String("adMarkers"),
//   	adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   	adTriggers: []*string{
//   		jsii.String("adTriggers"),
//   	},
//   	includeIframeOnlyStream: jsii.Boolean(false),
//   	manifestName: jsii.String("manifestName"),
//   	playlistType: jsii.String("playlistType"),
//   	playlistWindowSeconds: jsii.Number(123),
//   	programDateTimeIntervalSeconds: jsii.Number(123),
//   	url: jsii.String("url"),
//   }
//
type CfnOriginEndpoint_HlsManifestProperty struct {
	// The manifest ID is required and must be unique within the OriginEndpoint.
	//
	// The ID can't be changed after the endpoint is created.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Controls how ad markers are included in the packaged endpoint.
	//
	// Valid values are `none` , `passthrough` , or `scte35_enhanced` .
	//
	// - `NONE` - omits all SCTE-35 ad markers from the output.
	// - `PASSTHROUGH` - creates a copy in the output of the SCTE-35 ad markers (comments) taken directly from the input manifest.
	// - `SCTE35_ENHANCED` - generates ad markers and blackout tags in the output based on the SCTE-35 messages from the input manifest.
	AdMarkers *string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// The flags on SCTE-35 segmentation descriptors that have to be present for MediaPackage to insert ad markers in the output manifest.
	//
	// For information about SCTE-35 in MediaPackage, see [SCTE-35 Message Options in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/scte.html) .
	AdsOnDeliveryRestrictions *string `field:"optional" json:"adsOnDeliveryRestrictions" yaml:"adsOnDeliveryRestrictions"`
	// Specifies the SCTE-35 message types that MediaPackage treats as ad markers in the output manifest.
	//
	// Valid values:
	//
	// - `BREAK`
	// - `DISTRIBUTOR_ADVERTISEMENT`
	// - `DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY`
	// - `DISTRIBUTOR_PLACEMENT_OPPORTUNITY`
	// - `PROVIDER_ADVERTISEMENT`
	// - `PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY`
	// - `PROVIDER_PLACEMENT_OPPORTUNITY`
	// - `SPLICE_INSERT`.
	AdTriggers *[]*string `field:"optional" json:"adTriggers" yaml:"adTriggers"`
	// Applies to stream sets with a single video track only.
	//
	// When true, the stream set includes an additional I-frame only stream, along with the other tracks. If false, this extra stream is not included.
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// A short string that's appended to the end of the endpoint URL to create a unique path to this endpoint.
	//
	// The manifestName on the HLSManifest object overrides the manifestName that you provided on the originEndpoint object.
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// When specified as either `event` or `vod` , a corresponding `EXT-X-PLAYLIST-TYPE` entry is included in the media playlist.
	//
	// Indicates if the playlist is live-to-VOD content.
	PlaylistType *string `field:"optional" json:"playlistType" yaml:"playlistType"`
	// Time window (in seconds) contained in each parent manifest.
	PlaylistWindowSeconds *float64 `field:"optional" json:"playlistWindowSeconds" yaml:"playlistWindowSeconds"`
	// Inserts `EXT-X-PROGRAM-DATE-TIME` tags in the output manifest at the interval that you specify.
	//
	// Additionally, ID3Timed metadata messages are generated every 5 seconds starting when the content was ingested.
	//
	// Irrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.
	//
	// Omit this attribute or enter `0` to indicate that the `EXT-X-PROGRAM-DATE-TIME` tags are not included in the manifest.
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// The URL that's used to request this manifest from this endpoint.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

// Parameters for Apple HLS packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsPackageProperty := &hlsPackageProperty{
//   	adMarkers: jsii.String("adMarkers"),
//   	adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   	adTriggers: []*string{
//   		jsii.String("adTriggers"),
//   	},
//   	encryption: &hlsEncryptionProperty{
//   		spekeKeyProvider: &spekeKeyProviderProperty{
//   			resourceId: jsii.String("resourceId"),
//   			roleArn: jsii.String("roleArn"),
//   			systemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			url: jsii.String("url"),
//
//   			// the properties below are optional
//   			certificateArn: jsii.String("certificateArn"),
//   			encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   				presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   				presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		constantInitializationVector: jsii.String("constantInitializationVector"),
//   		encryptionMethod: jsii.String("encryptionMethod"),
//   		keyRotationIntervalSeconds: jsii.Number(123),
//   		repeatExtXKey: jsii.Boolean(false),
//   	},
//   	includeIframeOnlyStream: jsii.Boolean(false),
//   	playlistType: jsii.String("playlistType"),
//   	playlistWindowSeconds: jsii.Number(123),
//   	programDateTimeIntervalSeconds: jsii.Number(123),
//   	segmentDurationSeconds: jsii.Number(123),
//   	streamSelection: &streamSelectionProperty{
//   		maxVideoBitsPerSecond: jsii.Number(123),
//   		minVideoBitsPerSecond: jsii.Number(123),
//   		streamOrder: jsii.String("streamOrder"),
//   	},
//   	useAudioRenditionGroup: jsii.Boolean(false),
//   }
//
type CfnOriginEndpoint_HlsPackageProperty struct {
	// Controls how ad markers are included in the packaged endpoint.
	//
	// Valid values are `none` , `passthrough` , or `scte35_enhanced` .
	//
	// - `NONE` - omits all SCTE-35 ad markers from the output.
	// - `PASSTHROUGH` - creates a copy in the output of the SCTE-35 ad markers (comments) taken directly from the input manifest.
	// - `SCTE35_ENHANCED` - generates ad markers and blackout tags in the output based on the SCTE-35 messages from the input manifest.
	AdMarkers *string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// The flags on SCTE-35 segmentation descriptors that have to be present for MediaPackage to insert ad markers in the output manifest.
	//
	// For information about SCTE-35 in MediaPackage, see [SCTE-35 Message Options in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/scte.html) .
	AdsOnDeliveryRestrictions *string `field:"optional" json:"adsOnDeliveryRestrictions" yaml:"adsOnDeliveryRestrictions"`
	// Specifies the SCTE-35 message types that MediaPackage treats as ad markers in the output manifest.
	//
	// Valid values:
	//
	// - `BREAK`
	// - `DISTRIBUTOR_ADVERTISEMENT`
	// - `DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY`
	// - `DISTRIBUTOR_PLACEMENT_OPPORTUNITY`
	// - `PROVIDER_ADVERTISEMENT`
	// - `PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY`
	// - `PROVIDER_PLACEMENT_OPPORTUNITY`
	// - `SPLICE_INSERT`.
	AdTriggers *[]*string `field:"optional" json:"adTriggers" yaml:"adTriggers"`
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// Only applies to stream sets with a single video track.
	//
	// When true, the stream set includes an additional I-frame only stream, along with the other tracks. If false, this extra stream is not included.
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// When specified as either `event` or `vod` , a corresponding `EXT-X-PLAYLIST-TYPE` entry is included in the media playlist.
	//
	// Indicates if the playlist is live-to-VOD content.
	PlaylistType *string `field:"optional" json:"playlistType" yaml:"playlistType"`
	// Time window (in seconds) contained in each parent manifest.
	PlaylistWindowSeconds *float64 `field:"optional" json:"playlistWindowSeconds" yaml:"playlistWindowSeconds"`
	// Inserts `EXT-X-PROGRAM-DATE-TIME` tags in the output manifest at the interval that you specify.
	//
	// Additionally, ID3Timed metadata messages are generated every 5 seconds starting when the content was ingested.
	//
	// Irrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.
	//
	// Omit this attribute or enter `0` to indicate that the `EXT-X-PROGRAM-DATE-TIME` tags are not included in the manifest.
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// Limitations for outputs from the endpoint, based on the video bitrate.
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
	// When true, AWS Elemental MediaPackage bundles all audio tracks in a rendition group.
	//
	// All other tracks in the stream can be used with any audio rendition from the group.
	UseAudioRenditionGroup interface{} `field:"optional" json:"useAudioRenditionGroup" yaml:"useAudioRenditionGroup"`
}

// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mssEncryptionProperty := &mssEncryptionProperty{
//   	spekeKeyProvider: &spekeKeyProviderProperty{
//   		resourceId: jsii.String("resourceId"),
//   		roleArn: jsii.String("roleArn"),
//   		systemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		url: jsii.String("url"),
//
//   		// the properties below are optional
//   		certificateArn: jsii.String("certificateArn"),
//   		encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   			presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   			presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   		},
//   	},
//   }
//
type CfnOriginEndpoint_MssEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}

// Parameters for Microsoft Smooth Streaming packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mssPackageProperty := &mssPackageProperty{
//   	encryption: &mssEncryptionProperty{
//   		spekeKeyProvider: &spekeKeyProviderProperty{
//   			resourceId: jsii.String("resourceId"),
//   			roleArn: jsii.String("roleArn"),
//   			systemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			url: jsii.String("url"),
//
//   			// the properties below are optional
//   			certificateArn: jsii.String("certificateArn"),
//   			encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   				presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   				presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   			},
//   		},
//   	},
//   	manifestWindowSeconds: jsii.Number(123),
//   	segmentDurationSeconds: jsii.Number(123),
//   	streamSelection: &streamSelectionProperty{
//   		maxVideoBitsPerSecond: jsii.Number(123),
//   		minVideoBitsPerSecond: jsii.Number(123),
//   		streamOrder: jsii.String("streamOrder"),
//   	},
//   }
//
type CfnOriginEndpoint_MssPackageProperty struct {
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// Time window (in seconds) contained in each manifest.
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// Limitations for outputs from the endpoint, based on the video bitrate.
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

// Keyprovider settings for DRM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spekeKeyProviderProperty := &spekeKeyProviderProperty{
//   	resourceId: jsii.String("resourceId"),
//   	roleArn: jsii.String("roleArn"),
//   	systemIds: []*string{
//   		jsii.String("systemIds"),
//   	},
//   	url: jsii.String("url"),
//
//   	// the properties below are optional
//   	certificateArn: jsii.String("certificateArn"),
//   	encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   		presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   		presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   	},
//   }
//
type CfnOriginEndpoint_SpekeKeyProviderProperty struct {
	// Unique identifier for this endpoint, as it is configured in the key provider service.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The ARN for the IAM role that's granted by the key provider to provide access to the key provider API.
	//
	// This role must have a trust policy that allows AWS Elemental MediaPackage to assume the role, and it must have a sufficient permissions policy to allow access to the specific key retrieval URL. Valid format: arn:aws:iam::{accountID}:role/{name}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// List of unique identifiers for the DRM systems to use, as defined in the CPIX specification.
	SystemIds *[]*string `field:"required" json:"systemIds" yaml:"systemIds"`
	// URL for the key provider’s key retrieval API endpoint.
	//
	// Must start with https://.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The Amazon Resource Name (ARN) for the certificate that you imported to AWS Certificate Manager to add content key encryption to this endpoint.
	//
	// For this feature to work, your DRM key provider must support content key encryption.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// `CfnOriginEndpoint.SpekeKeyProviderProperty.EncryptionContractConfiguration`.
	EncryptionContractConfiguration interface{} `field:"optional" json:"encryptionContractConfiguration" yaml:"encryptionContractConfiguration"`
}

// Limitations for outputs from the endpoint, based on the video bitrate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamSelectionProperty := &streamSelectionProperty{
//   	maxVideoBitsPerSecond: jsii.Number(123),
//   	minVideoBitsPerSecond: jsii.Number(123),
//   	streamOrder: jsii.String("streamOrder"),
//   }
//
type CfnOriginEndpoint_StreamSelectionProperty struct {
	// The upper limit of the bitrates that this endpoint serves.
	//
	// If the video track exceeds this threshold, then AWS Elemental MediaPackage excludes it from output. If you don't specify a value, it defaults to 2147483647 bits per second.
	MaxVideoBitsPerSecond *float64 `field:"optional" json:"maxVideoBitsPerSecond" yaml:"maxVideoBitsPerSecond"`
	// The lower limit of the bitrates that this endpoint serves.
	//
	// If the video track is below this threshold, then AWS Elemental MediaPackage excludes it from output. If you don't specify a value, it defaults to 0 bits per second.
	MinVideoBitsPerSecond *float64 `field:"optional" json:"minVideoBitsPerSecond" yaml:"minVideoBitsPerSecond"`
	// Order in which the different video bitrates are presented to the player.
	//
	// Valid values: `ORIGINAL` , `VIDEO_BITRATE_ASCENDING` , `VIDEO_BITRATE_DESCENDING` .
	StreamOrder *string `field:"optional" json:"streamOrder" yaml:"streamOrder"`
}

// Properties for defining a `CfnOriginEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOriginEndpointProps := &cfnOriginEndpointProps{
//   	channelId: jsii.String("channelId"),
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	authorization: &authorizationProperty{
//   		cdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   		secretsRoleArn: jsii.String("secretsRoleArn"),
//   	},
//   	cmafPackage: &cmafPackageProperty{
//   		encryption: &cmafEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			constantInitializationVector: jsii.String("constantInitializationVector"),
//   			keyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		hlsManifests: []interface{}{
//   			&hlsManifestProperty{
//   				id: jsii.String("id"),
//
//   				// the properties below are optional
//   				adMarkers: jsii.String("adMarkers"),
//   				adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   				adTriggers: []*string{
//   					jsii.String("adTriggers"),
//   				},
//   				includeIframeOnlyStream: jsii.Boolean(false),
//   				manifestName: jsii.String("manifestName"),
//   				playlistType: jsii.String("playlistType"),
//   				playlistWindowSeconds: jsii.Number(123),
//   				programDateTimeIntervalSeconds: jsii.Number(123),
//   				url: jsii.String("url"),
//   			},
//   		},
//   		segmentDurationSeconds: jsii.Number(123),
//   		segmentPrefix: jsii.String("segmentPrefix"),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   	},
//   	dashPackage: &dashPackageProperty{
//   		adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   		adTriggers: []*string{
//   			jsii.String("adTriggers"),
//   		},
//   		encryption: &dashEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			keyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		manifestLayout: jsii.String("manifestLayout"),
//   		manifestWindowSeconds: jsii.Number(123),
//   		minBufferTimeSeconds: jsii.Number(123),
//   		minUpdatePeriodSeconds: jsii.Number(123),
//   		periodTriggers: []*string{
//   			jsii.String("periodTriggers"),
//   		},
//   		profile: jsii.String("profile"),
//   		segmentDurationSeconds: jsii.Number(123),
//   		segmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   		suggestedPresentationDelaySeconds: jsii.Number(123),
//   		utcTiming: jsii.String("utcTiming"),
//   		utcTimingUri: jsii.String("utcTimingUri"),
//   	},
//   	description: jsii.String("description"),
//   	hlsPackage: &hlsPackageProperty{
//   		adMarkers: jsii.String("adMarkers"),
//   		adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   		adTriggers: []*string{
//   			jsii.String("adTriggers"),
//   		},
//   		encryption: &hlsEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			constantInitializationVector: jsii.String("constantInitializationVector"),
//   			encryptionMethod: jsii.String("encryptionMethod"),
//   			keyRotationIntervalSeconds: jsii.Number(123),
//   			repeatExtXKey: jsii.Boolean(false),
//   		},
//   		includeIframeOnlyStream: jsii.Boolean(false),
//   		playlistType: jsii.String("playlistType"),
//   		playlistWindowSeconds: jsii.Number(123),
//   		programDateTimeIntervalSeconds: jsii.Number(123),
//   		segmentDurationSeconds: jsii.Number(123),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   		useAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	manifestName: jsii.String("manifestName"),
//   	mssPackage: &mssPackageProperty{
//   		encryption: &mssEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//   		},
//   		manifestWindowSeconds: jsii.Number(123),
//   		segmentDurationSeconds: jsii.Number(123),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   	},
//   	origination: jsii.String("origination"),
//   	startoverWindowSeconds: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeDelaySeconds: jsii.Number(123),
//   	whitelist: []*string{
//   		jsii.String("whitelist"),
//   	},
//   }
//
type CfnOriginEndpointProps struct {
	// The ID of the channel associated with this endpoint.
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// The manifest ID is required and must be unique within the OriginEndpoint.
	//
	// The ID can't be changed after the endpoint is created.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Parameters for CDN authorization.
	Authorization interface{} `field:"optional" json:"authorization" yaml:"authorization"`
	// Parameters for Common Media Application Format (CMAF) packaging.
	CmafPackage interface{} `field:"optional" json:"cmafPackage" yaml:"cmafPackage"`
	// Parameters for DASH packaging.
	DashPackage interface{} `field:"optional" json:"dashPackage" yaml:"dashPackage"`
	// Any descriptive information that you want to add to the endpoint for future identification purposes.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Parameters for Apple HLS packaging.
	HlsPackage interface{} `field:"optional" json:"hlsPackage" yaml:"hlsPackage"`
	// A short string that's appended to the end of the endpoint URL to create a unique path to this endpoint.
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// Parameters for Microsoft Smooth Streaming packaging.
	MssPackage interface{} `field:"optional" json:"mssPackage" yaml:"mssPackage"`
	// Controls video origination from this endpoint.
	//
	// - `ALLOW` - enables this endpoint to serve content to requesting devices.
	// - `DENY` - prevents this endpoint from serving content. Denying origination is helpful for harvesting live-to-VOD assets. For more information about harvesting and origination, see [Live-to-VOD Requirements](https://docs.aws.amazon.com/mediapackage/latest/ug/ltov-reqmts.html) .
	Origination *string `field:"optional" json:"origination" yaml:"origination"`
	// Maximum duration (seconds) of content to retain for startover playback.
	//
	// Omit this attribute or enter `0` to indicate that startover playback is disabled for this endpoint.
	StartoverWindowSeconds *float64 `field:"optional" json:"startoverWindowSeconds" yaml:"startoverWindowSeconds"`
	// The tags to assign to the endpoint.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Minimum duration (seconds) of delay to enforce on the playback of live content.
	//
	// Omit this attribute or enter `0` to indicate that there is no time delay in effect for this endpoint.
	TimeDelaySeconds *float64 `field:"optional" json:"timeDelaySeconds" yaml:"timeDelaySeconds"`
	// The IP addresses that can access this endpoint.
	Whitelist *[]*string `field:"optional" json:"whitelist" yaml:"whitelist"`
}

// A CloudFormation `AWS::MediaPackage::PackagingConfiguration`.
//
// Creates a packaging configuration in a packaging group.
//
// The packaging configuration represents a single delivery point for an asset. It determines the format and setting for the egressing content. Specify only one package format per configuration, such as `HlsPackage` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackagingConfiguration := awscdk.Aws_mediapackage.NewCfnPackagingConfiguration(this, jsii.String("MyCfnPackagingConfiguration"), &cfnPackagingConfigurationProps{
//   	id: jsii.String("id"),
//   	packagingGroupId: jsii.String("packagingGroupId"),
//
//   	// the properties below are optional
//   	cmafPackage: &cmafPackageProperty{
//   		hlsManifests: []interface{}{
//   			&hlsManifestProperty{
//   				adMarkers: jsii.String("adMarkers"),
//   				includeIframeOnlyStream: jsii.Boolean(false),
//   				manifestName: jsii.String("manifestName"),
//   				programDateTimeIntervalSeconds: jsii.Number(123),
//   				repeatExtXKey: jsii.Boolean(false),
//   				streamSelection: &streamSelectionProperty{
//   					maxVideoBitsPerSecond: jsii.Number(123),
//   					minVideoBitsPerSecond: jsii.Number(123),
//   					streamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		encryption: &cmafEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//   			},
//   		},
//   		includeEncoderConfigurationInSegments: jsii.Boolean(false),
//   		segmentDurationSeconds: jsii.Number(123),
//   	},
//   	dashPackage: &dashPackageProperty{
//   		dashManifests: []interface{}{
//   			&dashManifestProperty{
//   				manifestLayout: jsii.String("manifestLayout"),
//   				manifestName: jsii.String("manifestName"),
//   				minBufferTimeSeconds: jsii.Number(123),
//   				profile: jsii.String("profile"),
//   				streamSelection: &streamSelectionProperty{
//   					maxVideoBitsPerSecond: jsii.Number(123),
//   					minVideoBitsPerSecond: jsii.Number(123),
//   					streamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		encryption: &dashEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//   			},
//   		},
//   		includeEncoderConfigurationInSegments: jsii.Boolean(false),
//   		periodTriggers: []*string{
//   			jsii.String("periodTriggers"),
//   		},
//   		segmentDurationSeconds: jsii.Number(123),
//   		segmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   	},
//   	hlsPackage: &hlsPackageProperty{
//   		hlsManifests: []interface{}{
//   			&hlsManifestProperty{
//   				adMarkers: jsii.String("adMarkers"),
//   				includeIframeOnlyStream: jsii.Boolean(false),
//   				manifestName: jsii.String("manifestName"),
//   				programDateTimeIntervalSeconds: jsii.Number(123),
//   				repeatExtXKey: jsii.Boolean(false),
//   				streamSelection: &streamSelectionProperty{
//   					maxVideoBitsPerSecond: jsii.Number(123),
//   					minVideoBitsPerSecond: jsii.Number(123),
//   					streamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		encryption: &hlsEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//   			},
//
//   			// the properties below are optional
//   			constantInitializationVector: jsii.String("constantInitializationVector"),
//   			encryptionMethod: jsii.String("encryptionMethod"),
//   		},
//   		segmentDurationSeconds: jsii.Number(123),
//   		useAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	mssPackage: &mssPackageProperty{
//   		mssManifests: []interface{}{
//   			&mssManifestProperty{
//   				manifestName: jsii.String("manifestName"),
//   				streamSelection: &streamSelectionProperty{
//   					maxVideoBitsPerSecond: jsii.Number(123),
//   					minVideoBitsPerSecond: jsii.Number(123),
//   					streamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		encryption: &mssEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//   			},
//   		},
//   		segmentDurationSeconds: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPackagingConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) for the packaging configuration.
	//
	// You can get this from the response to any request to the packaging configuration.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Parameters for CMAF packaging.
	CmafPackage() interface{}
	SetCmafPackage(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Parameters for DASH-ISO packaging.
	DashPackage() interface{}
	SetDashPackage(val interface{})
	// Parameters for Apple HLS packaging.
	HlsPackage() interface{}
	SetHlsPackage(val interface{})
	// Unique identifier that you assign to the packaging configuration.
	Id() *string
	SetId(val *string)
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
	// Parameters for Microsoft Smooth Streaming packaging.
	MssPackage() interface{}
	SetMssPackage(val interface{})
	// The tree node.
	Node() constructs.Node
	// The ID of the packaging group associated with this packaging configuration.
	PackagingGroupId() *string
	SetPackagingGroupId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to assign to the packaging configuration.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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
}

// The jsii proxy struct for CfnPackagingConfiguration
type jsiiProxy_CfnPackagingConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPackagingConfiguration) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CmafPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cmafPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) DashPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) HlsPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) MssPackage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mssPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) PackagingGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packagingGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackage::PackagingConfiguration`.
func NewCfnPackagingConfiguration(scope constructs.Construct, id *string, props *CfnPackagingConfigurationProps) CfnPackagingConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnPackagingConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackage::PackagingConfiguration`.
func NewCfnPackagingConfiguration_Override(c CfnPackagingConfiguration, scope constructs.Construct, id *string, props *CfnPackagingConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration) SetCmafPackage(val interface{}) {
	_jsii_.Set(
		j,
		"cmafPackage",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration) SetDashPackage(val interface{}) {
	_jsii_.Set(
		j,
		"dashPackage",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration) SetHlsPackage(val interface{}) {
	_jsii_.Set(
		j,
		"hlsPackage",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration) SetMssPackage(val interface{}) {
	_jsii_.Set(
		j,
		"mssPackage",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingConfiguration) SetPackagingGroupId(val *string) {
	_jsii_.Set(
		j,
		"packagingGroupId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPackagingConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPackagingConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnPackagingConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPackagingConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPackagingConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cmafEncryptionProperty := &cmafEncryptionProperty{
//   	spekeKeyProvider: &spekeKeyProviderProperty{
//   		roleArn: jsii.String("roleArn"),
//   		systemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		url: jsii.String("url"),
//   	},
//   }
//
type CfnPackagingConfiguration_CmafEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}

// Parameters for a packaging configuration that uses Common Media Application Format (CMAF) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cmafPackageProperty := &cmafPackageProperty{
//   	hlsManifests: []interface{}{
//   		&hlsManifestProperty{
//   			adMarkers: jsii.String("adMarkers"),
//   			includeIframeOnlyStream: jsii.Boolean(false),
//   			manifestName: jsii.String("manifestName"),
//   			programDateTimeIntervalSeconds: jsii.Number(123),
//   			repeatExtXKey: jsii.Boolean(false),
//   			streamSelection: &streamSelectionProperty{
//   				maxVideoBitsPerSecond: jsii.Number(123),
//   				minVideoBitsPerSecond: jsii.Number(123),
//   				streamOrder: jsii.String("streamOrder"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	encryption: &cmafEncryptionProperty{
//   		spekeKeyProvider: &spekeKeyProviderProperty{
//   			roleArn: jsii.String("roleArn"),
//   			systemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			url: jsii.String("url"),
//   		},
//   	},
//   	includeEncoderConfigurationInSegments: jsii.Boolean(false),
//   	segmentDurationSeconds: jsii.Number(123),
//   }
//
type CfnPackagingConfiguration_CmafPackageProperty struct {
	// A list of HLS manifest configurations that are available from this endpoint.
	HlsManifests interface{} `field:"required" json:"hlsManifests" yaml:"hlsManifests"`
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// When includeEncoderConfigurationInSegments is set to true, MediaPackage places your encoder's Sequence Parameter Set (SPS), Picture Parameter Set (PPS), and Video Parameter Set (VPS) metadata in every video segment instead of in the init fragment.
	//
	// This lets you use different SPS/PPS/VPS settings for your assets during content playback.
	IncludeEncoderConfigurationInSegments interface{} `field:"optional" json:"includeEncoderConfigurationInSegments" yaml:"includeEncoderConfigurationInSegments"`
	// Duration (in seconds) of each segment.
	//
	// Actual segments are rounded to the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
}

// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashEncryptionProperty := &dashEncryptionProperty{
//   	spekeKeyProvider: &spekeKeyProviderProperty{
//   		roleArn: jsii.String("roleArn"),
//   		systemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		url: jsii.String("url"),
//   	},
//   }
//
type CfnPackagingConfiguration_DashEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}

// Parameters for a DASH manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashManifestProperty := &dashManifestProperty{
//   	manifestLayout: jsii.String("manifestLayout"),
//   	manifestName: jsii.String("manifestName"),
//   	minBufferTimeSeconds: jsii.Number(123),
//   	profile: jsii.String("profile"),
//   	streamSelection: &streamSelectionProperty{
//   		maxVideoBitsPerSecond: jsii.Number(123),
//   		minVideoBitsPerSecond: jsii.Number(123),
//   		streamOrder: jsii.String("streamOrder"),
//   	},
//   }
//
type CfnPackagingConfiguration_DashManifestProperty struct {
	// Determines the position of some tags in the Media Presentation Description (MPD).
	//
	// When set to `FULL` , elements like `SegmentTemplate` and `ContentProtection` are included in each `Representation` . When set to `COMPACT` , duplicate elements are combined and presented at the AdaptationSet level.
	ManifestLayout *string `field:"optional" json:"manifestLayout" yaml:"manifestLayout"`
	// A short string that's appended to the end of the endpoint URL to create a unique path to this packaging configuration.
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// Minimum amount of content (measured in seconds) that a player must keep available in the buffer.
	MinBufferTimeSeconds *float64 `field:"optional" json:"minBufferTimeSeconds" yaml:"minBufferTimeSeconds"`
	// The DASH profile type.
	//
	// When set to `HBBTV_1_5` , the content is compliant with HbbTV 1.5.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Limitations for outputs from the endpoint, based on the video bitrate.
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

// Parameters for a packaging configuration that uses Dynamic Adaptive Streaming over HTTP (DASH) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashPackageProperty := &dashPackageProperty{
//   	dashManifests: []interface{}{
//   		&dashManifestProperty{
//   			manifestLayout: jsii.String("manifestLayout"),
//   			manifestName: jsii.String("manifestName"),
//   			minBufferTimeSeconds: jsii.Number(123),
//   			profile: jsii.String("profile"),
//   			streamSelection: &streamSelectionProperty{
//   				maxVideoBitsPerSecond: jsii.Number(123),
//   				minVideoBitsPerSecond: jsii.Number(123),
//   				streamOrder: jsii.String("streamOrder"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	encryption: &dashEncryptionProperty{
//   		spekeKeyProvider: &spekeKeyProviderProperty{
//   			roleArn: jsii.String("roleArn"),
//   			systemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			url: jsii.String("url"),
//   		},
//   	},
//   	includeEncoderConfigurationInSegments: jsii.Boolean(false),
//   	periodTriggers: []*string{
//   		jsii.String("periodTriggers"),
//   	},
//   	segmentDurationSeconds: jsii.Number(123),
//   	segmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   }
//
type CfnPackagingConfiguration_DashPackageProperty struct {
	// A list of DASH manifest configurations that are available from this endpoint.
	DashManifests interface{} `field:"required" json:"dashManifests" yaml:"dashManifests"`
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// When includeEncoderConfigurationInSegments is set to true, MediaPackage places your encoder's Sequence Parameter Set (SPS), Picture Parameter Set (PPS), and Video Parameter Set (VPS) metadata in every video segment instead of in the init fragment.
	//
	// This lets you use different SPS/PPS/VPS settings for your assets during content playback.
	IncludeEncoderConfigurationInSegments interface{} `field:"optional" json:"includeEncoderConfigurationInSegments" yaml:"includeEncoderConfigurationInSegments"`
	// Controls whether MediaPackage produces single-period or multi-period DASH manifests.
	//
	// For more information about periods, see [Multi-period DASH in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/multi-period.html) .
	//
	// Valid values:
	//
	// - `ADS` - MediaPackage will produce multi-period DASH manifests. Periods are created based on the SCTE-35 ad markers present in the input manifest.
	// - *No value* - MediaPackage will produce single-period DASH manifests. This is the default setting.
	PeriodTriggers *[]*string `field:"optional" json:"periodTriggers" yaml:"periodTriggers"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source segment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// Determines the type of SegmentTemplate included in the Media Presentation Description (MPD).
	//
	// When set to `NUMBER_WITH_TIMELINE` , a full timeline is presented in each SegmentTemplate, with $Number$ media URLs. When set to `TIME_WITH_TIMELINE` , a full timeline is presented in each SegmentTemplate, with $Time$ media URLs. When set to `NUMBER_WITH_DURATION` , only a duration is included in each SegmentTemplate, with $Number$ media URLs.
	SegmentTemplateFormat *string `field:"optional" json:"segmentTemplateFormat" yaml:"segmentTemplateFormat"`
}

// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsEncryptionProperty := &hlsEncryptionProperty{
//   	spekeKeyProvider: &spekeKeyProviderProperty{
//   		roleArn: jsii.String("roleArn"),
//   		systemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		url: jsii.String("url"),
//   	},
//
//   	// the properties below are optional
//   	constantInitializationVector: jsii.String("constantInitializationVector"),
//   	encryptionMethod: jsii.String("encryptionMethod"),
//   }
//
type CfnPackagingConfiguration_HlsEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// A 128-bit, 16-byte hex value represented by a 32-character string, used with the key for encrypting blocks.
	//
	// If you don't specify a constant initialization vector (IV), MediaPackage periodically rotates the IV.
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// HLS encryption type.
	EncryptionMethod *string `field:"optional" json:"encryptionMethod" yaml:"encryptionMethod"`
}

// Parameters for an HLS manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsManifestProperty := &hlsManifestProperty{
//   	adMarkers: jsii.String("adMarkers"),
//   	includeIframeOnlyStream: jsii.Boolean(false),
//   	manifestName: jsii.String("manifestName"),
//   	programDateTimeIntervalSeconds: jsii.Number(123),
//   	repeatExtXKey: jsii.Boolean(false),
//   	streamSelection: &streamSelectionProperty{
//   		maxVideoBitsPerSecond: jsii.Number(123),
//   		minVideoBitsPerSecond: jsii.Number(123),
//   		streamOrder: jsii.String("streamOrder"),
//   	},
//   }
//
type CfnPackagingConfiguration_HlsManifestProperty struct {
	// This setting controls ad markers in the packaged content.
	//
	// `NONE` omits SCTE-35 ad markers from the output. `PASSTHROUGH` copies SCTE-35 ad markers from the source content to the output. `SCTE35_ENHANCED` generates ad markers and blackout tags in the output, based on SCTE-35 messages in the source content.
	AdMarkers *string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// Applies to stream sets with a single video track only.
	//
	// When enabled, the output includes an additional I-frame only stream, along with the other tracks.
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// A short string that's appended to the end of the endpoint URL to create a unique path to this packaging configuration.
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// Inserts `EXT-X-PROGRAM-DATE-TIME` tags in the output manifest at the interval that you specify.
	//
	// Additionally, ID3Timed metadata messages are generated every 5 seconds starting when the content was ingested.
	//
	// Irrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.
	//
	// Omit this attribute or enter `0` to indicate that the `EXT-X-PROGRAM-DATE-TIME` tags are not included in the manifest.
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// Repeat the `EXT-X-KEY` directive for every media segment.
	//
	// This might result in an increase in client requests to the DRM server.
	RepeatExtXKey interface{} `field:"optional" json:"repeatExtXKey" yaml:"repeatExtXKey"`
	// Video bitrate limitations for outputs from this packaging configuration.
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

// Parameters for a packaging configuration that uses HTTP Live Streaming (HLS) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsPackageProperty := &hlsPackageProperty{
//   	hlsManifests: []interface{}{
//   		&hlsManifestProperty{
//   			adMarkers: jsii.String("adMarkers"),
//   			includeIframeOnlyStream: jsii.Boolean(false),
//   			manifestName: jsii.String("manifestName"),
//   			programDateTimeIntervalSeconds: jsii.Number(123),
//   			repeatExtXKey: jsii.Boolean(false),
//   			streamSelection: &streamSelectionProperty{
//   				maxVideoBitsPerSecond: jsii.Number(123),
//   				minVideoBitsPerSecond: jsii.Number(123),
//   				streamOrder: jsii.String("streamOrder"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	encryption: &hlsEncryptionProperty{
//   		spekeKeyProvider: &spekeKeyProviderProperty{
//   			roleArn: jsii.String("roleArn"),
//   			systemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			url: jsii.String("url"),
//   		},
//
//   		// the properties below are optional
//   		constantInitializationVector: jsii.String("constantInitializationVector"),
//   		encryptionMethod: jsii.String("encryptionMethod"),
//   	},
//   	segmentDurationSeconds: jsii.Number(123),
//   	useAudioRenditionGroup: jsii.Boolean(false),
//   }
//
type CfnPackagingConfiguration_HlsPackageProperty struct {
	// A list of HLS manifest configurations that are available from this endpoint.
	HlsManifests interface{} `field:"required" json:"hlsManifests" yaml:"hlsManifests"`
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// When true, AWS Elemental MediaPackage bundles all audio tracks in a rendition group.
	//
	// All other tracks in the stream can be used with any audio rendition from the group.
	UseAudioRenditionGroup interface{} `field:"optional" json:"useAudioRenditionGroup" yaml:"useAudioRenditionGroup"`
}

// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mssEncryptionProperty := &mssEncryptionProperty{
//   	spekeKeyProvider: &spekeKeyProviderProperty{
//   		roleArn: jsii.String("roleArn"),
//   		systemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		url: jsii.String("url"),
//   	},
//   }
//
type CfnPackagingConfiguration_MssEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}

// Parameters for a Microsoft Smooth manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mssManifestProperty := &mssManifestProperty{
//   	manifestName: jsii.String("manifestName"),
//   	streamSelection: &streamSelectionProperty{
//   		maxVideoBitsPerSecond: jsii.Number(123),
//   		minVideoBitsPerSecond: jsii.Number(123),
//   		streamOrder: jsii.String("streamOrder"),
//   	},
//   }
//
type CfnPackagingConfiguration_MssManifestProperty struct {
	// A short string that's appended to the end of the endpoint URL to create a unique path to this packaging configuration.
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// Video bitrate limitations for outputs from this packaging configuration.
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

// Parameters for a packaging configuration that uses Microsoft Smooth Streaming (MSS) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mssPackageProperty := &mssPackageProperty{
//   	mssManifests: []interface{}{
//   		&mssManifestProperty{
//   			manifestName: jsii.String("manifestName"),
//   			streamSelection: &streamSelectionProperty{
//   				maxVideoBitsPerSecond: jsii.Number(123),
//   				minVideoBitsPerSecond: jsii.Number(123),
//   				streamOrder: jsii.String("streamOrder"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	encryption: &mssEncryptionProperty{
//   		spekeKeyProvider: &spekeKeyProviderProperty{
//   			roleArn: jsii.String("roleArn"),
//   			systemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			url: jsii.String("url"),
//   		},
//   	},
//   	segmentDurationSeconds: jsii.Number(123),
//   }
//
type CfnPackagingConfiguration_MssPackageProperty struct {
	// A list of Microsoft Smooth manifest configurations that are available from this endpoint.
	MssManifests interface{} `field:"required" json:"mssManifests" yaml:"mssManifests"`
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments are rounded to the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
}

// A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that provides encryption keys.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spekeKeyProviderProperty := &spekeKeyProviderProperty{
//   	roleArn: jsii.String("roleArn"),
//   	systemIds: []*string{
//   		jsii.String("systemIds"),
//   	},
//   	url: jsii.String("url"),
//   }
//
type CfnPackagingConfiguration_SpekeKeyProviderProperty struct {
	// The ARN for the IAM role that's granted by the key provider to provide access to the key provider API.
	//
	// Valid format: arn:aws:iam::{accountID}:role/{name}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// List of unique identifiers for the DRM systems to use, as defined in the CPIX specification.
	SystemIds *[]*string `field:"required" json:"systemIds" yaml:"systemIds"`
	// URL for the key provider's key retrieval API endpoint.
	//
	// Must start with https://.
	Url *string `field:"required" json:"url" yaml:"url"`
}

// Limitations for outputs from the endpoint, based on the video bitrate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamSelectionProperty := &streamSelectionProperty{
//   	maxVideoBitsPerSecond: jsii.Number(123),
//   	minVideoBitsPerSecond: jsii.Number(123),
//   	streamOrder: jsii.String("streamOrder"),
//   }
//
type CfnPackagingConfiguration_StreamSelectionProperty struct {
	// The upper limit of the bitrates that this endpoint serves.
	//
	// If the video track exceeds this threshold, then AWS Elemental MediaPackage excludes it from output. If you don't specify a value, it defaults to 2147483647 bits per second.
	MaxVideoBitsPerSecond *float64 `field:"optional" json:"maxVideoBitsPerSecond" yaml:"maxVideoBitsPerSecond"`
	// The lower limit of the bitrates that this endpoint serves.
	//
	// If the video track is below this threshold, then AWS Elemental MediaPackage excludes it from output. If you don't specify a value, it defaults to 0 bits per second.
	MinVideoBitsPerSecond *float64 `field:"optional" json:"minVideoBitsPerSecond" yaml:"minVideoBitsPerSecond"`
	// Order in which the different video bitrates are presented to the player.
	//
	// Valid values: `ORIGINAL` , `VIDEO_BITRATE_ASCENDING` , `VIDEO_BITRATE_DESCENDING` .
	StreamOrder *string `field:"optional" json:"streamOrder" yaml:"streamOrder"`
}

// Properties for defining a `CfnPackagingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackagingConfigurationProps := &cfnPackagingConfigurationProps{
//   	id: jsii.String("id"),
//   	packagingGroupId: jsii.String("packagingGroupId"),
//
//   	// the properties below are optional
//   	cmafPackage: &cmafPackageProperty{
//   		hlsManifests: []interface{}{
//   			&hlsManifestProperty{
//   				adMarkers: jsii.String("adMarkers"),
//   				includeIframeOnlyStream: jsii.Boolean(false),
//   				manifestName: jsii.String("manifestName"),
//   				programDateTimeIntervalSeconds: jsii.Number(123),
//   				repeatExtXKey: jsii.Boolean(false),
//   				streamSelection: &streamSelectionProperty{
//   					maxVideoBitsPerSecond: jsii.Number(123),
//   					minVideoBitsPerSecond: jsii.Number(123),
//   					streamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		encryption: &cmafEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//   			},
//   		},
//   		includeEncoderConfigurationInSegments: jsii.Boolean(false),
//   		segmentDurationSeconds: jsii.Number(123),
//   	},
//   	dashPackage: &dashPackageProperty{
//   		dashManifests: []interface{}{
//   			&dashManifestProperty{
//   				manifestLayout: jsii.String("manifestLayout"),
//   				manifestName: jsii.String("manifestName"),
//   				minBufferTimeSeconds: jsii.Number(123),
//   				profile: jsii.String("profile"),
//   				streamSelection: &streamSelectionProperty{
//   					maxVideoBitsPerSecond: jsii.Number(123),
//   					minVideoBitsPerSecond: jsii.Number(123),
//   					streamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		encryption: &dashEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//   			},
//   		},
//   		includeEncoderConfigurationInSegments: jsii.Boolean(false),
//   		periodTriggers: []*string{
//   			jsii.String("periodTriggers"),
//   		},
//   		segmentDurationSeconds: jsii.Number(123),
//   		segmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   	},
//   	hlsPackage: &hlsPackageProperty{
//   		hlsManifests: []interface{}{
//   			&hlsManifestProperty{
//   				adMarkers: jsii.String("adMarkers"),
//   				includeIframeOnlyStream: jsii.Boolean(false),
//   				manifestName: jsii.String("manifestName"),
//   				programDateTimeIntervalSeconds: jsii.Number(123),
//   				repeatExtXKey: jsii.Boolean(false),
//   				streamSelection: &streamSelectionProperty{
//   					maxVideoBitsPerSecond: jsii.Number(123),
//   					minVideoBitsPerSecond: jsii.Number(123),
//   					streamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		encryption: &hlsEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//   			},
//
//   			// the properties below are optional
//   			constantInitializationVector: jsii.String("constantInitializationVector"),
//   			encryptionMethod: jsii.String("encryptionMethod"),
//   		},
//   		segmentDurationSeconds: jsii.Number(123),
//   		useAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	mssPackage: &mssPackageProperty{
//   		mssManifests: []interface{}{
//   			&mssManifestProperty{
//   				manifestName: jsii.String("manifestName"),
//   				streamSelection: &streamSelectionProperty{
//   					maxVideoBitsPerSecond: jsii.Number(123),
//   					minVideoBitsPerSecond: jsii.Number(123),
//   					streamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		encryption: &mssEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//   			},
//   		},
//   		segmentDurationSeconds: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPackagingConfigurationProps struct {
	// Unique identifier that you assign to the packaging configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ID of the packaging group associated with this packaging configuration.
	PackagingGroupId *string `field:"required" json:"packagingGroupId" yaml:"packagingGroupId"`
	// Parameters for CMAF packaging.
	CmafPackage interface{} `field:"optional" json:"cmafPackage" yaml:"cmafPackage"`
	// Parameters for DASH-ISO packaging.
	DashPackage interface{} `field:"optional" json:"dashPackage" yaml:"dashPackage"`
	// Parameters for Apple HLS packaging.
	HlsPackage interface{} `field:"optional" json:"hlsPackage" yaml:"hlsPackage"`
	// Parameters for Microsoft Smooth Streaming packaging.
	MssPackage interface{} `field:"optional" json:"mssPackage" yaml:"mssPackage"`
	// The tags to assign to the packaging configuration.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::MediaPackage::PackagingGroup`.
//
// Creates a packaging group.
//
// The packaging group holds one or more packaging configurations. When you create an asset, you specify the packaging group associated with the asset. The asset has playback endpoints for each packaging configuration within the group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackagingGroup := awscdk.Aws_mediapackage.NewCfnPackagingGroup(this, jsii.String("MyCfnPackagingGroup"), &cfnPackagingGroupProps{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	authorization: &authorizationProperty{
//   		cdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   		secretsRoleArn: jsii.String("secretsRoleArn"),
//   	},
//   	egressAccessLogs: &logConfigurationProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPackagingGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) for the packaging group.
	//
	// You can get this from the response to any request to the packaging group.
	AttrArn() *string
	// The fully qualified domain name for the assets in the PackagingGroup.
	AttrDomainName() *string
	// Parameters for CDN authorization.
	Authorization() interface{}
	SetAuthorization(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The configuration parameters for egress access logging.
	EgressAccessLogs() interface{}
	SetEgressAccessLogs(val interface{})
	// Unique identifier that you assign to the packaging group.
	Id() *string
	SetId(val *string)
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
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to assign to the packaging group.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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
}

// The jsii proxy struct for CfnPackagingGroup
type jsiiProxy_CfnPackagingGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPackagingGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) AttrDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) Authorization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) EgressAccessLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressAccessLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPackagingGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaPackage::PackagingGroup`.
func NewCfnPackagingGroup(scope constructs.Construct, id *string, props *CfnPackagingGroupProps) CfnPackagingGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnPackagingGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaPackage::PackagingGroup`.
func NewCfnPackagingGroup_Override(c CfnPackagingGroup, scope constructs.Construct, id *string, props *CfnPackagingGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPackagingGroup) SetAuthorization(val interface{}) {
	_jsii_.Set(
		j,
		"authorization",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingGroup) SetEgressAccessLogs(val interface{}) {
	_jsii_.Set(
		j,
		"egressAccessLogs",
		val,
	)
}

func (j *jsiiProxy_CfnPackagingGroup) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPackagingGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPackagingGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnPackagingGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPackagingGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediapackage.CfnPackagingGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPackagingGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPackagingGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPackagingGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPackagingGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPackagingGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPackagingGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPackagingGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPackagingGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPackagingGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPackagingGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPackagingGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Parameters for enabling CDN authorization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationProperty := &authorizationProperty{
//   	cdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   	secretsRoleArn: jsii.String("secretsRoleArn"),
//   }
//
type CfnPackagingGroup_AuthorizationProperty struct {
	// The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that is used for CDN authorization.
	CdnIdentifierSecret *string `field:"required" json:"cdnIdentifierSecret" yaml:"cdnIdentifierSecret"`
	// The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager .
	SecretsRoleArn *string `field:"required" json:"secretsRoleArn" yaml:"secretsRoleArn"`
}

// Sets a custom Amazon CloudWatch log group name for egress logs.
//
// If a log group name isn't specified, the default name is used: /aws/MediaPackage/EgressAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &logConfigurationProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   }
//
type CfnPackagingGroup_LogConfigurationProperty struct {
	// Sets a custom Amazon CloudWatch log group name for egress logs.
	//
	// If a log group name isn't specified, the default name is used: /aws/MediaPackage/EgressAccessLogs.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
}

// Properties for defining a `CfnPackagingGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackagingGroupProps := &cfnPackagingGroupProps{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	authorization: &authorizationProperty{
//   		cdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   		secretsRoleArn: jsii.String("secretsRoleArn"),
//   	},
//   	egressAccessLogs: &logConfigurationProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPackagingGroupProps struct {
	// Unique identifier that you assign to the packaging group.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Parameters for CDN authorization.
	Authorization interface{} `field:"optional" json:"authorization" yaml:"authorization"`
	// The configuration parameters for egress access logging.
	EgressAccessLogs interface{} `field:"optional" json:"egressAccessLogs" yaml:"egressAccessLogs"`
	// The tags to assign to the packaging group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

