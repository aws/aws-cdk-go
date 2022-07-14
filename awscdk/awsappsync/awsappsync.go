package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AppSync::ApiCache`.
//
// The `AWS::AppSync::ApiCache` resource represents the input of a `CreateApiCache` operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiCache := awscdk.Aws_appsync.NewCfnApiCache(this, jsii.String("MyCfnApiCache"), &cfnApiCacheProps{
//   	apiCachingBehavior: jsii.String("apiCachingBehavior"),
//   	apiId: jsii.String("apiId"),
//   	ttl: jsii.Number(123),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	atRestEncryptionEnabled: jsii.Boolean(false),
//   	transitEncryptionEnabled: jsii.Boolean(false),
//   })
//
type CfnApiCache interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Caching behavior.
	//
	// - *FULL_REQUEST_CACHING* : All requests are fully cached.
	// - *PER_RESOLVER_CACHING* : Individual resolvers that you specify are cached.
	ApiCachingBehavior() *string
	SetApiCachingBehavior(val *string)
	// The GraphQL API ID.
	ApiId() *string
	SetApiId(val *string)
	// At-rest encryption flag for cache.
	//
	// You cannot update this setting after creation.
	AtRestEncryptionEnabled() interface{}
	SetAtRestEncryptionEnabled(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// Transit encryption flag when connecting to cache.
	//
	// You cannot update this setting after creation.
	TransitEncryptionEnabled() interface{}
	SetTransitEncryptionEnabled(val interface{})
	// TTL in seconds for cache entries.
	//
	// Valid values are 1–3,600 seconds.
	Ttl() *float64
	SetTtl(val *float64)
	// The cache instance type. Valid values are.
	//
	// - `SMALL`
	// - `MEDIUM`
	// - `LARGE`
	// - `XLARGE`
	// - `LARGE_2X`
	// - `LARGE_4X`
	// - `LARGE_8X` (not available in all regions)
	// - `LARGE_12X`
	//
	// Historically, instance types were identified by an EC2-style value. As of July 2020, this is deprecated, and the generic identifiers above should be used.
	//
	// The following legacy instance types are available, but their use is discouraged:
	//
	// - *T2_SMALL* : A t2.small instance type.
	// - *T2_MEDIUM* : A t2.medium instance type.
	// - *R4_LARGE* : A r4.large instance type.
	// - *R4_XLARGE* : A r4.xlarge instance type.
	// - *R4_2XLARGE* : A r4.2xlarge instance type.
	// - *R4_4XLARGE* : A r4.4xlarge instance type.
	// - *R4_8XLARGE* : A r4.8xlarge instance type.
	Type() *string
	SetType(val *string)
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

// The jsii proxy struct for CfnApiCache
type jsiiProxy_CfnApiCache struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApiCache) ApiCachingBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiCachingBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) AtRestEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atRestEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) TransitEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::ApiCache`.
func NewCfnApiCache(scope constructs.Construct, id *string, props *CfnApiCacheProps) CfnApiCache {
	_init_.Initialize()

	j := jsiiProxy_CfnApiCache{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnApiCache",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::ApiCache`.
func NewCfnApiCache_Override(c CfnApiCache, scope constructs.Construct, id *string, props *CfnApiCacheProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnApiCache",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApiCache) SetApiCachingBehavior(val *string) {
	_jsii_.Set(
		j,
		"apiCachingBehavior",
		val,
	)
}

func (j *jsiiProxy_CfnApiCache) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnApiCache) SetAtRestEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"atRestEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApiCache) SetTransitEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"transitEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApiCache) SetTtl(val *float64) {
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_CfnApiCache) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnApiCache_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnApiCache",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnApiCache_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnApiCache",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnApiCache_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnApiCache",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApiCache_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.CfnApiCache",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApiCache) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApiCache) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApiCache) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApiCache) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApiCache) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApiCache) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApiCache) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApiCache) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiCache) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiCache) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApiCache) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApiCache) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiCache) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiCache) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiCache) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnApiCache`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiCacheProps := &cfnApiCacheProps{
//   	apiCachingBehavior: jsii.String("apiCachingBehavior"),
//   	apiId: jsii.String("apiId"),
//   	ttl: jsii.Number(123),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	atRestEncryptionEnabled: jsii.Boolean(false),
//   	transitEncryptionEnabled: jsii.Boolean(false),
//   }
//
type CfnApiCacheProps struct {
	// Caching behavior.
	//
	// - *FULL_REQUEST_CACHING* : All requests are fully cached.
	// - *PER_RESOLVER_CACHING* : Individual resolvers that you specify are cached.
	ApiCachingBehavior *string `field:"required" json:"apiCachingBehavior" yaml:"apiCachingBehavior"`
	// The GraphQL API ID.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// TTL in seconds for cache entries.
	//
	// Valid values are 1–3,600 seconds.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
	// The cache instance type. Valid values are.
	//
	// - `SMALL`
	// - `MEDIUM`
	// - `LARGE`
	// - `XLARGE`
	// - `LARGE_2X`
	// - `LARGE_4X`
	// - `LARGE_8X` (not available in all regions)
	// - `LARGE_12X`
	//
	// Historically, instance types were identified by an EC2-style value. As of July 2020, this is deprecated, and the generic identifiers above should be used.
	//
	// The following legacy instance types are available, but their use is discouraged:
	//
	// - *T2_SMALL* : A t2.small instance type.
	// - *T2_MEDIUM* : A t2.medium instance type.
	// - *R4_LARGE* : A r4.large instance type.
	// - *R4_XLARGE* : A r4.xlarge instance type.
	// - *R4_2XLARGE* : A r4.2xlarge instance type.
	// - *R4_4XLARGE* : A r4.4xlarge instance type.
	// - *R4_8XLARGE* : A r4.8xlarge instance type.
	Type *string `field:"required" json:"type" yaml:"type"`
	// At-rest encryption flag for cache.
	//
	// You cannot update this setting after creation.
	AtRestEncryptionEnabled interface{} `field:"optional" json:"atRestEncryptionEnabled" yaml:"atRestEncryptionEnabled"`
	// Transit encryption flag when connecting to cache.
	//
	// You cannot update this setting after creation.
	TransitEncryptionEnabled interface{} `field:"optional" json:"transitEncryptionEnabled" yaml:"transitEncryptionEnabled"`
}

// A CloudFormation `AWS::AppSync::ApiKey`.
//
// The `AWS::AppSync::ApiKey` resource creates a unique key that you can distribute to clients who are executing GraphQL operations with AWS AppSync that require an API key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiKey := awscdk.Aws_appsync.NewCfnApiKey(this, jsii.String("MyCfnApiKey"), &cfnApiKeyProps{
//   	apiId: jsii.String("apiId"),
//
//   	// the properties below are optional
//   	apiKeyId: jsii.String("apiKeyId"),
//   	description: jsii.String("description"),
//   	expires: jsii.Number(123),
//   })
//
type CfnApiKey interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Unique AWS AppSync GraphQL API ID for this API key.
	ApiId() *string
	SetApiId(val *string)
	// The API key ID.
	ApiKeyId() *string
	SetApiKeyId(val *string)
	// The API key.
	AttrApiKey() *string
	// The Amazon Resource Name (ARN) of the API key, such as `arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/apikey/apikeya1bzhi` .
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
	// Unique description of your API key.
	Description() *string
	SetDescription(val *string)
	// The time after which the API key expires.
	//
	// The date is represented as seconds since the epoch, rounded down to the nearest hour.
	Expires() *float64
	SetExpires(val *float64)
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

// The jsii proxy struct for CfnApiKey
type jsiiProxy_CfnApiKey struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApiKey) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) ApiKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) AttrApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Expires() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::ApiKey`.
func NewCfnApiKey(scope constructs.Construct, id *string, props *CfnApiKeyProps) CfnApiKey {
	_init_.Initialize()

	j := jsiiProxy_CfnApiKey{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnApiKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::ApiKey`.
func NewCfnApiKey_Override(c CfnApiKey, scope constructs.Construct, id *string, props *CfnApiKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnApiKey",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApiKey) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnApiKey) SetApiKeyId(val *string) {
	_jsii_.Set(
		j,
		"apiKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnApiKey) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnApiKey) SetExpires(val *float64) {
	_jsii_.Set(
		j,
		"expires",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnApiKey_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnApiKey",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnApiKey_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnApiKey",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnApiKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnApiKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApiKey_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.CfnApiKey",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApiKey) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApiKey) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApiKey) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApiKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApiKey) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApiKey) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApiKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApiKey) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiKey) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiKey) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApiKey) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApiKey) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiKey) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiKey) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnApiKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiKeyProps := &cfnApiKeyProps{
//   	apiId: jsii.String("apiId"),
//
//   	// the properties below are optional
//   	apiKeyId: jsii.String("apiKeyId"),
//   	description: jsii.String("description"),
//   	expires: jsii.Number(123),
//   }
//
type CfnApiKeyProps struct {
	// Unique AWS AppSync GraphQL API ID for this API key.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The API key ID.
	ApiKeyId *string `field:"optional" json:"apiKeyId" yaml:"apiKeyId"`
	// Unique description of your API key.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The time after which the API key expires.
	//
	// The date is represented as seconds since the epoch, rounded down to the nearest hour.
	Expires *float64 `field:"optional" json:"expires" yaml:"expires"`
}

// A CloudFormation `AWS::AppSync::DataSource`.
//
// The `AWS::AppSync::DataSource` resource creates data sources for resolvers in AWS AppSync to connect to, such as Amazon DynamoDB , AWS Lambda , and Amazon OpenSearch Service . Resolvers use these data sources to fetch data when clients make GraphQL calls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSource := awscdk.Aws_appsync.NewCfnDataSource(this, jsii.String("MyCfnDataSource"), &cfnDataSourceProps{
//   	apiId: jsii.String("apiId"),
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	dynamoDbConfig: &dynamoDBConfigProperty{
//   		awsRegion: jsii.String("awsRegion"),
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		deltaSyncConfig: &deltaSyncConfigProperty{
//   			baseTableTtl: jsii.String("baseTableTtl"),
//   			deltaSyncTableName: jsii.String("deltaSyncTableName"),
//   			deltaSyncTableTtl: jsii.String("deltaSyncTableTtl"),
//   		},
//   		useCallerCredentials: jsii.Boolean(false),
//   		versioned: jsii.Boolean(false),
//   	},
//   	elasticsearchConfig: &elasticsearchConfigProperty{
//   		awsRegion: jsii.String("awsRegion"),
//   		endpoint: jsii.String("endpoint"),
//   	},
//   	httpConfig: &httpConfigProperty{
//   		endpoint: jsii.String("endpoint"),
//
//   		// the properties below are optional
//   		authorizationConfig: &authorizationConfigProperty{
//   			authorizationType: jsii.String("authorizationType"),
//
//   			// the properties below are optional
//   			awsIamConfig: &awsIamConfigProperty{
//   				signingRegion: jsii.String("signingRegion"),
//   				signingServiceName: jsii.String("signingServiceName"),
//   			},
//   		},
//   	},
//   	lambdaConfig: &lambdaConfigProperty{
//   		lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   	},
//   	openSearchServiceConfig: &openSearchServiceConfigProperty{
//   		awsRegion: jsii.String("awsRegion"),
//   		endpoint: jsii.String("endpoint"),
//   	},
//   	relationalDatabaseConfig: &relationalDatabaseConfigProperty{
//   		relationalDatabaseSourceType: jsii.String("relationalDatabaseSourceType"),
//
//   		// the properties below are optional
//   		rdsHttpEndpointConfig: &rdsHttpEndpointConfigProperty{
//   			awsRegion: jsii.String("awsRegion"),
//   			awsSecretStoreArn: jsii.String("awsSecretStoreArn"),
//   			dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//
//   			// the properties below are optional
//   			databaseName: jsii.String("databaseName"),
//   			schema: jsii.String("schema"),
//   		},
//   	},
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//   })
//
type CfnDataSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Unique AWS AppSync GraphQL API identifier where this data source will be created.
	ApiId() *string
	SetApiId(val *string)
	// The Amazon Resource Name (ARN) of the API key, such as `arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/datasources/datasourcename` .
	AttrDataSourceArn() *string
	// Friendly name for you to identify your AWS AppSync data source after creation.
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the data source.
	Description() *string
	SetDescription(val *string)
	// AWS Region and TableName for an Amazon DynamoDB table in your account.
	DynamoDbConfig() interface{}
	SetDynamoDbConfig(val interface{})
	// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
	//
	// As of September 2021, Amazon Elasticsearch Service is Amazon OpenSearch Service . This property is deprecated. For new data sources, use *OpenSearchServiceConfig* to specify an OpenSearch Service data source.
	ElasticsearchConfig() interface{}
	SetElasticsearchConfig(val interface{})
	// Endpoints for an HTTP data source.
	HttpConfig() interface{}
	SetHttpConfig(val interface{})
	// An ARN of a Lambda function in valid ARN format.
	//
	// This can be the ARN of a Lambda function that exists in the current account or in another account.
	LambdaConfig() interface{}
	SetLambdaConfig(val interface{})
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
	// Friendly name for you to identify your AppSync data source after creation.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
	OpenSearchServiceConfig() interface{}
	SetOpenSearchServiceConfig(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Relational Database configuration of the relational database data source.
	RelationalDatabaseConfig() interface{}
	SetRelationalDatabaseConfig(val interface{})
	// The AWS Identity and Access Management service role ARN for the data source.
	//
	// The system assumes this role when accessing the data source.
	//
	// Required if `Type` is specified as `AWS_LAMBDA` , `AMAZON_DYNAMODB` , `AMAZON_ELASTICSEARCH` , or `AMAZON_OPENSEARCH_SERVICE` .
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The type of the data source.
	//
	// - *AWS_LAMBDA* : The data source is an AWS Lambda function.
	// - *AMAZON_DYNAMODB* : The data source is an Amazon DynamoDB table.
	// - *AMAZON_ELASTICSEARCH* : The data source is an Amazon OpenSearch Service domain.
	// - *AMAZON_OPENSEARCH_SERVICE* : The data source is an Amazon OpenSearch Service domain.
	// - *NONE* : There is no data source. This type is used when you wish to invoke a GraphQL operation without connecting to a data source, such as performing data transformation with resolvers or triggering a subscription to be invoked from a mutation.
	// - *HTTP* : The data source is an HTTP endpoint.
	// - *RELATIONAL_DATABASE* : The data source is a relational database.
	Type() *string
	SetType(val *string)
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

// The jsii proxy struct for CfnDataSource
type jsiiProxy_CfnDataSource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataSource) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrDataSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDataSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DynamoDbConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) ElasticsearchConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) HttpConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) LambdaConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) OpenSearchServiceConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openSearchServiceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) RelationalDatabaseConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relationalDatabaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::DataSource`.
func NewCfnDataSource(scope constructs.Construct, id *string, props *CfnDataSourceProps) CfnDataSource {
	_init_.Initialize()

	j := jsiiProxy_CfnDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::DataSource`.
func NewCfnDataSource_Override(c CfnDataSource, scope constructs.Construct, id *string, props *CfnDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnDataSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataSource) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDynamoDbConfig(val interface{}) {
	_jsii_.Set(
		j,
		"dynamoDbConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetElasticsearchConfig(val interface{}) {
	_jsii_.Set(
		j,
		"elasticsearchConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetHttpConfig(val interface{}) {
	_jsii_.Set(
		j,
		"httpConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetLambdaConfig(val interface{}) {
	_jsii_.Set(
		j,
		"lambdaConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetOpenSearchServiceConfig(val interface{}) {
	_jsii_.Set(
		j,
		"openSearchServiceConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetRelationalDatabaseConfig(val interface{}) {
	_jsii_.Set(
		j,
		"relationalDatabaseConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataSource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnDataSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataSource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnDataSource",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.CfnDataSource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataSource) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataSource) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataSource) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `AuthorizationConfig` property type specifies the authorization type and configuration for an AWS AppSync http data source.
//
// `AuthorizationConfig` is a property of the [AWS AppSync DataSource HttpConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-httpconfig.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationConfigProperty := &authorizationConfigProperty{
//   	authorizationType: jsii.String("authorizationType"),
//
//   	// the properties below are optional
//   	awsIamConfig: &awsIamConfigProperty{
//   		signingRegion: jsii.String("signingRegion"),
//   		signingServiceName: jsii.String("signingServiceName"),
//   	},
//   }
//
type CfnDataSource_AuthorizationConfigProperty struct {
	// The authorization type that the HTTP endpoint requires.
	//
	// - *AWS_IAM* : The authorization type is Signature Version 4 (SigV4).
	AuthorizationType *string `field:"required" json:"authorizationType" yaml:"authorizationType"`
	// The AWS Identity and Access Management settings.
	AwsIamConfig interface{} `field:"optional" json:"awsIamConfig" yaml:"awsIamConfig"`
}

// Use the `AwsIamConfig` property type to specify `AwsIamConfig` for a AWS AppSync authorizaton.
//
// `AwsIamConfig` is a property of the [AWS AppSync DataSource AuthorizationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-httpconfig-authorizationconfig.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsIamConfigProperty := &awsIamConfigProperty{
//   	signingRegion: jsii.String("signingRegion"),
//   	signingServiceName: jsii.String("signingServiceName"),
//   }
//
type CfnDataSource_AwsIamConfigProperty struct {
	// The signing Region for AWS Identity and Access Management authorization.
	SigningRegion *string `field:"optional" json:"signingRegion" yaml:"signingRegion"`
	// The signing service name for AWS Identity and Access Management authorization.
	SigningServiceName *string `field:"optional" json:"signingServiceName" yaml:"signingServiceName"`
}

// Describes a Delta Sync configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deltaSyncConfigProperty := &deltaSyncConfigProperty{
//   	baseTableTtl: jsii.String("baseTableTtl"),
//   	deltaSyncTableName: jsii.String("deltaSyncTableName"),
//   	deltaSyncTableTtl: jsii.String("deltaSyncTableTtl"),
//   }
//
type CfnDataSource_DeltaSyncConfigProperty struct {
	// The number of minutes that an Item is stored in the data source.
	BaseTableTtl *string `field:"required" json:"baseTableTtl" yaml:"baseTableTtl"`
	// The Delta Sync table name.
	DeltaSyncTableName *string `field:"required" json:"deltaSyncTableName" yaml:"deltaSyncTableName"`
	// The number of minutes that a Delta Sync log entry is stored in the Delta Sync table.
	DeltaSyncTableTtl *string `field:"required" json:"deltaSyncTableTtl" yaml:"deltaSyncTableTtl"`
}

// The `DynamoDBConfig` property type specifies the `AwsRegion` and `TableName` for an Amazon DynamoDB table in your account for an AWS AppSync data source.
//
// `DynamoDBConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDBConfigProperty := &dynamoDBConfigProperty{
//   	awsRegion: jsii.String("awsRegion"),
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	deltaSyncConfig: &deltaSyncConfigProperty{
//   		baseTableTtl: jsii.String("baseTableTtl"),
//   		deltaSyncTableName: jsii.String("deltaSyncTableName"),
//   		deltaSyncTableTtl: jsii.String("deltaSyncTableTtl"),
//   	},
//   	useCallerCredentials: jsii.Boolean(false),
//   	versioned: jsii.Boolean(false),
//   }
//
type CfnDataSource_DynamoDBConfigProperty struct {
	// The AWS Region.
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// The table name.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The `DeltaSyncConfig` for a versioned datasource.
	DeltaSyncConfig interface{} `field:"optional" json:"deltaSyncConfig" yaml:"deltaSyncConfig"`
	// Set to `TRUE` to use AWS Identity and Access Management with this data source.
	UseCallerCredentials interface{} `field:"optional" json:"useCallerCredentials" yaml:"useCallerCredentials"`
	// Set to TRUE to use Conflict Detection and Resolution with this data source.
	Versioned interface{} `field:"optional" json:"versioned" yaml:"versioned"`
}

// The `ElasticsearchConfig` property type specifies the `AwsRegion` and `Endpoints` for an Amazon OpenSearch Service domain in your account for an AWS AppSync data source.
//
// ElasticsearchConfig is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// As of September 2021, Amazon Elasticsearch Service is Amazon OpenSearch Service . This property is deprecated. For new data sources, use *OpenSearchServiceConfig* to specify an OpenSearch Service data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticsearchConfigProperty := &elasticsearchConfigProperty{
//   	awsRegion: jsii.String("awsRegion"),
//   	endpoint: jsii.String("endpoint"),
//   }
//
type CfnDataSource_ElasticsearchConfigProperty struct {
	// The AWS Region.
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// The endpoint.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
}

// Use the `HttpConfig` property type to specify `HttpConfig` for an AWS AppSync data source.
//
// `HttpConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpConfigProperty := &httpConfigProperty{
//   	endpoint: jsii.String("endpoint"),
//
//   	// the properties below are optional
//   	authorizationConfig: &authorizationConfigProperty{
//   		authorizationType: jsii.String("authorizationType"),
//
//   		// the properties below are optional
//   		awsIamConfig: &awsIamConfigProperty{
//   			signingRegion: jsii.String("signingRegion"),
//   			signingServiceName: jsii.String("signingServiceName"),
//   		},
//   	},
//   }
//
type CfnDataSource_HttpConfigProperty struct {
	// The endpoint.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The authorization configuration.
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

// The `LambdaConfig` property type specifies the Lambda function ARN for an AWS AppSync data source.
//
// `LambdaConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConfigProperty := &lambdaConfigProperty{
//   	lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   }
//
type CfnDataSource_LambdaConfigProperty struct {
	// The ARN for the Lambda function.
	LambdaFunctionArn *string `field:"required" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
}

// The `OpenSearchServiceConfig` property type specifies the `AwsRegion` and `Endpoints` for an Amazon OpenSearch Service domain in your account for an AWS AppSync data source.
//
// `OpenSearchServiceConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openSearchServiceConfigProperty := &openSearchServiceConfigProperty{
//   	awsRegion: jsii.String("awsRegion"),
//   	endpoint: jsii.String("endpoint"),
//   }
//
type CfnDataSource_OpenSearchServiceConfigProperty struct {
	// The AWS Region.
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// The endpoint.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
}

// Use the `RdsHttpEndpointConfig` property type to specify the `RdsHttpEndpoint` for an AWS AppSync relational database.
//
// `RdsHttpEndpointConfig` is a property of the [AWS AppSync DataSource RelationalDatabaseConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-relationaldatabaseconfig.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rdsHttpEndpointConfigProperty := &rdsHttpEndpointConfigProperty{
//   	awsRegion: jsii.String("awsRegion"),
//   	awsSecretStoreArn: jsii.String("awsSecretStoreArn"),
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//
//   	// the properties below are optional
//   	databaseName: jsii.String("databaseName"),
//   	schema: jsii.String("schema"),
//   }
//
type CfnDataSource_RdsHttpEndpointConfigProperty struct {
	// AWS Region for RDS HTTP endpoint.
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// The ARN for database credentials stored in AWS Secrets Manager .
	AwsSecretStoreArn *string `field:"required" json:"awsSecretStoreArn" yaml:"awsSecretStoreArn"`
	// Amazon RDS cluster Amazon Resource Name (ARN).
	DbClusterIdentifier *string `field:"required" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Logical database name.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Logical schema name.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

// Use the `RelationalDatabaseConfig` property type to specify `RelationalDatabaseConfig` for an AWS AppSync data source.
//
// `RelationalDatabaseConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relationalDatabaseConfigProperty := &relationalDatabaseConfigProperty{
//   	relationalDatabaseSourceType: jsii.String("relationalDatabaseSourceType"),
//
//   	// the properties below are optional
//   	rdsHttpEndpointConfig: &rdsHttpEndpointConfigProperty{
//   		awsRegion: jsii.String("awsRegion"),
//   		awsSecretStoreArn: jsii.String("awsSecretStoreArn"),
//   		dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//
//   		// the properties below are optional
//   		databaseName: jsii.String("databaseName"),
//   		schema: jsii.String("schema"),
//   	},
//   }
//
type CfnDataSource_RelationalDatabaseConfigProperty struct {
	// The type of relational data source.
	RelationalDatabaseSourceType *string `field:"required" json:"relationalDatabaseSourceType" yaml:"relationalDatabaseSourceType"`
	// Information about the Amazon RDS resource.
	RdsHttpEndpointConfig interface{} `field:"optional" json:"rdsHttpEndpointConfig" yaml:"rdsHttpEndpointConfig"`
}

// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSourceProps := &cfnDataSourceProps{
//   	apiId: jsii.String("apiId"),
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	dynamoDbConfig: &dynamoDBConfigProperty{
//   		awsRegion: jsii.String("awsRegion"),
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		deltaSyncConfig: &deltaSyncConfigProperty{
//   			baseTableTtl: jsii.String("baseTableTtl"),
//   			deltaSyncTableName: jsii.String("deltaSyncTableName"),
//   			deltaSyncTableTtl: jsii.String("deltaSyncTableTtl"),
//   		},
//   		useCallerCredentials: jsii.Boolean(false),
//   		versioned: jsii.Boolean(false),
//   	},
//   	elasticsearchConfig: &elasticsearchConfigProperty{
//   		awsRegion: jsii.String("awsRegion"),
//   		endpoint: jsii.String("endpoint"),
//   	},
//   	httpConfig: &httpConfigProperty{
//   		endpoint: jsii.String("endpoint"),
//
//   		// the properties below are optional
//   		authorizationConfig: &authorizationConfigProperty{
//   			authorizationType: jsii.String("authorizationType"),
//
//   			// the properties below are optional
//   			awsIamConfig: &awsIamConfigProperty{
//   				signingRegion: jsii.String("signingRegion"),
//   				signingServiceName: jsii.String("signingServiceName"),
//   			},
//   		},
//   	},
//   	lambdaConfig: &lambdaConfigProperty{
//   		lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   	},
//   	openSearchServiceConfig: &openSearchServiceConfigProperty{
//   		awsRegion: jsii.String("awsRegion"),
//   		endpoint: jsii.String("endpoint"),
//   	},
//   	relationalDatabaseConfig: &relationalDatabaseConfigProperty{
//   		relationalDatabaseSourceType: jsii.String("relationalDatabaseSourceType"),
//
//   		// the properties below are optional
//   		rdsHttpEndpointConfig: &rdsHttpEndpointConfigProperty{
//   			awsRegion: jsii.String("awsRegion"),
//   			awsSecretStoreArn: jsii.String("awsSecretStoreArn"),
//   			dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//
//   			// the properties below are optional
//   			databaseName: jsii.String("databaseName"),
//   			schema: jsii.String("schema"),
//   		},
//   	},
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//   }
//
type CfnDataSourceProps struct {
	// Unique AWS AppSync GraphQL API identifier where this data source will be created.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Friendly name for you to identify your AppSync data source after creation.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the data source.
	//
	// - *AWS_LAMBDA* : The data source is an AWS Lambda function.
	// - *AMAZON_DYNAMODB* : The data source is an Amazon DynamoDB table.
	// - *AMAZON_ELASTICSEARCH* : The data source is an Amazon OpenSearch Service domain.
	// - *AMAZON_OPENSEARCH_SERVICE* : The data source is an Amazon OpenSearch Service domain.
	// - *NONE* : There is no data source. This type is used when you wish to invoke a GraphQL operation without connecting to a data source, such as performing data transformation with resolvers or triggering a subscription to be invoked from a mutation.
	// - *HTTP* : The data source is an HTTP endpoint.
	// - *RELATIONAL_DATABASE* : The data source is a relational database.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the data source.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// AWS Region and TableName for an Amazon DynamoDB table in your account.
	DynamoDbConfig interface{} `field:"optional" json:"dynamoDbConfig" yaml:"dynamoDbConfig"`
	// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
	//
	// As of September 2021, Amazon Elasticsearch Service is Amazon OpenSearch Service . This property is deprecated. For new data sources, use *OpenSearchServiceConfig* to specify an OpenSearch Service data source.
	ElasticsearchConfig interface{} `field:"optional" json:"elasticsearchConfig" yaml:"elasticsearchConfig"`
	// Endpoints for an HTTP data source.
	HttpConfig interface{} `field:"optional" json:"httpConfig" yaml:"httpConfig"`
	// An ARN of a Lambda function in valid ARN format.
	//
	// This can be the ARN of a Lambda function that exists in the current account or in another account.
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
	OpenSearchServiceConfig interface{} `field:"optional" json:"openSearchServiceConfig" yaml:"openSearchServiceConfig"`
	// Relational Database configuration of the relational database data source.
	RelationalDatabaseConfig interface{} `field:"optional" json:"relationalDatabaseConfig" yaml:"relationalDatabaseConfig"`
	// The AWS Identity and Access Management service role ARN for the data source.
	//
	// The system assumes this role when accessing the data source.
	//
	// Required if `Type` is specified as `AWS_LAMBDA` , `AMAZON_DYNAMODB` , `AMAZON_ELASTICSEARCH` , or `AMAZON_OPENSEARCH_SERVICE` .
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
}

// A CloudFormation `AWS::AppSync::DomainName`.
//
// The `AWS::AppSync::DomainName` resource creates a `DomainNameConfig` object to configure a custom domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainName := awscdk.Aws_appsync.NewCfnDomainName(this, jsii.String("MyCfnDomainName"), &cfnDomainNameProps{
//   	certificateArn: jsii.String("certificateArn"),
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   })
//
type CfnDomainName interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The domain name provided by AWS AppSync .
	AttrAppSyncDomainName() *string
	AttrDomainName() *string
	// The ID of your Amazon Route 53 hosted zone.
	AttrHostedZoneId() *string
	// The Amazon Resource Name (ARN) of the certificate.
	//
	// This will be an AWS Certificate Manager certificate.
	CertificateArn() *string
	SetCertificateArn(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The decription for your domain name.
	Description() *string
	SetDescription(val *string)
	// The domain name.
	DomainName() *string
	SetDomainName(val *string)
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

// The jsii proxy struct for CfnDomainName
type jsiiProxy_CfnDomainName struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDomainName) AttrAppSyncDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppSyncDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) AttrDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) AttrHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::DomainName`.
func NewCfnDomainName(scope constructs.Construct, id *string, props *CfnDomainNameProps) CfnDomainName {
	_init_.Initialize()

	j := jsiiProxy_CfnDomainName{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnDomainName",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::DomainName`.
func NewCfnDomainName_Override(c CfnDomainName, scope constructs.Construct, id *string, props *CfnDomainNameProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnDomainName",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomainName) SetCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_CfnDomainName) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDomainName) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDomainName_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnDomainName",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDomainName_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnDomainName",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnDomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnDomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainName_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.CfnDomainName",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomainName) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDomainName) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomainName) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDomainName) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDomainName) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDomainName) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDomainName) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDomainName) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainName) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainName) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDomainName) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDomainName) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainName) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainName) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A CloudFormation `AWS::AppSync::DomainNameApiAssociation`.
//
// The `AWS::AppSync::DomainNameApiAssociation` resource represents the mapping of your custom domain name to the assigned API URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainNameApiAssociation := awscdk.Aws_appsync.NewCfnDomainNameApiAssociation(this, jsii.String("MyCfnDomainNameApiAssociation"), &cfnDomainNameApiAssociationProps{
//   	apiId: jsii.String("apiId"),
//   	domainName: jsii.String("domainName"),
//   })
//
type CfnDomainNameApiAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The API ID.
	ApiId() *string
	SetApiId(val *string)
	AttrApiAssociationIdentifier() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The domain name.
	DomainName() *string
	SetDomainName(val *string)
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

// The jsii proxy struct for CfnDomainNameApiAssociation
type jsiiProxy_CfnDomainNameApiAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) AttrApiAssociationIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrApiAssociationIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::DomainNameApiAssociation`.
func NewCfnDomainNameApiAssociation(scope constructs.Construct, id *string, props *CfnDomainNameApiAssociationProps) CfnDomainNameApiAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnDomainNameApiAssociation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnDomainNameApiAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::DomainNameApiAssociation`.
func NewCfnDomainNameApiAssociation_Override(c CfnDomainNameApiAssociation, scope constructs.Construct, id *string, props *CfnDomainNameApiAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnDomainNameApiAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnDomainNameApiAssociation) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDomainNameApiAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnDomainNameApiAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDomainNameApiAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnDomainNameApiAssociation",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnDomainNameApiAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnDomainNameApiAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainNameApiAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.CfnDomainNameApiAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainNameApiAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDomainNameApiAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainNameApiAssociationProps := &cfnDomainNameApiAssociationProps{
//   	apiId: jsii.String("apiId"),
//   	domainName: jsii.String("domainName"),
//   }
//
type CfnDomainNameApiAssociationProps struct {
	// The API ID.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The domain name.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

// Properties for defining a `CfnDomainName`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainNameProps := &cfnDomainNameProps{
//   	certificateArn: jsii.String("certificateArn"),
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnDomainNameProps struct {
	// The Amazon Resource Name (ARN) of the certificate.
	//
	// This will be an AWS Certificate Manager certificate.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// The domain name.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The decription for your domain name.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

// A CloudFormation `AWS::AppSync::FunctionConfiguration`.
//
// The `AWS::AppSync::FunctionConfiguration` resource defines the functions in GraphQL APIs to perform certain operations. You can use pipeline resolvers to attach functions. For more information, see [Pipeline Resolvers](https://docs.aws.amazon.com/appsync/latest/devguide/pipeline-resolvers.html) in the *AWS AppSync Developer Guide* .
//
// > When you submit an update, AWS CloudFormation updates resources based on differences between what you submit and the stack's current template. To cause this resource to be updated you must change a property value for this resource in the AWS CloudFormation template. Changing the Amazon S3 file content without changing a property value will not result in an update operation.
// >
// > See [Update Behaviors of Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html) in the *AWS CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFunctionConfiguration := awscdk.Aws_appsync.NewCfnFunctionConfiguration(this, jsii.String("MyCfnFunctionConfiguration"), &cfnFunctionConfigurationProps{
//   	apiId: jsii.String("apiId"),
//   	dataSourceName: jsii.String("dataSourceName"),
//   	functionVersion: jsii.String("functionVersion"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	maxBatchSize: jsii.Number(123),
//   	requestMappingTemplate: jsii.String("requestMappingTemplate"),
//   	requestMappingTemplateS3Location: jsii.String("requestMappingTemplateS3Location"),
//   	responseMappingTemplate: jsii.String("responseMappingTemplate"),
//   	responseMappingTemplateS3Location: jsii.String("responseMappingTemplateS3Location"),
//   	syncConfig: &syncConfigProperty{
//   		conflictDetection: jsii.String("conflictDetection"),
//
//   		// the properties below are optional
//   		conflictHandler: jsii.String("conflictHandler"),
//   		lambdaConflictHandlerConfig: &lambdaConflictHandlerConfigProperty{
//   			lambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   		},
//   	},
//   })
//
type CfnFunctionConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The AWS AppSync GraphQL API that you want to attach using this function.
	ApiId() *string
	SetApiId(val *string)
	// The name of data source this function will attach.
	AttrDataSourceName() *string
	// ARN of the function, such as `arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/functions/functionId` .
	AttrFunctionArn() *string
	// The unique ID of this function.
	AttrFunctionId() *string
	// The name of the function.
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of data source this function will attach.
	DataSourceName() *string
	SetDataSourceName(val *string)
	// The `Function` description.
	Description() *string
	SetDescription(val *string)
	// The version of the request mapping template.
	//
	// Currently, only the 2018-05-29 version of the template is supported.
	FunctionVersion() *string
	SetFunctionVersion(val *string)
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
	// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a `BatchInvoke` operation.
	MaxBatchSize() *float64
	SetMaxBatchSize(val *float64)
	// The name of the function.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The `Function` request mapping template.
	//
	// Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate() *string
	SetRequestMappingTemplate(val *string)
	// Describes a Sync configuration for a resolver.
	//
	// Contains information on which Conflict Detection, as well as Resolution strategy, should be performed when the resolver is invoked.
	RequestMappingTemplateS3Location() *string
	SetRequestMappingTemplateS3Location(val *string)
	// The `Function` response mapping template.
	ResponseMappingTemplate() *string
	SetResponseMappingTemplate(val *string)
	// The location of a response mapping template in an Amazon S3 bucket.
	//
	// Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
	ResponseMappingTemplateS3Location() *string
	SetResponseMappingTemplateS3Location(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Describes a Sync configuration for a resolver.
	//
	// Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
	SyncConfig() interface{}
	SetSyncConfig(val interface{})
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

// The jsii proxy struct for CfnFunctionConfiguration
type jsiiProxy_CfnFunctionConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFunctionConfiguration) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) AttrDataSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDataSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) AttrFunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFunctionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) AttrFunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFunctionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) DataSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) FunctionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) MaxBatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) RequestMappingTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) RequestMappingTemplateS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplateS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) ResponseMappingTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) ResponseMappingTemplateS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplateS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) SyncConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::FunctionConfiguration`.
func NewCfnFunctionConfiguration(scope constructs.Construct, id *string, props *CfnFunctionConfigurationProps) CfnFunctionConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnFunctionConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnFunctionConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::FunctionConfiguration`.
func NewCfnFunctionConfiguration_Override(c CfnFunctionConfiguration, scope constructs.Construct, id *string, props *CfnFunctionConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnFunctionConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetDataSourceName(val *string) {
	_jsii_.Set(
		j,
		"dataSourceName",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetFunctionVersion(val *string) {
	_jsii_.Set(
		j,
		"functionVersion",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetMaxBatchSize(val *float64) {
	_jsii_.Set(
		j,
		"maxBatchSize",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetRequestMappingTemplate(val *string) {
	_jsii_.Set(
		j,
		"requestMappingTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetRequestMappingTemplateS3Location(val *string) {
	_jsii_.Set(
		j,
		"requestMappingTemplateS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetResponseMappingTemplate(val *string) {
	_jsii_.Set(
		j,
		"responseMappingTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetResponseMappingTemplateS3Location(val *string) {
	_jsii_.Set(
		j,
		"responseMappingTemplateS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetSyncConfig(val interface{}) {
	_jsii_.Set(
		j,
		"syncConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFunctionConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnFunctionConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFunctionConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnFunctionConfiguration",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnFunctionConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnFunctionConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFunctionConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.CfnFunctionConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFunctionConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFunctionConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFunctionConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFunctionConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFunctionConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFunctionConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFunctionConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFunctionConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFunctionConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFunctionConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunctionConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `LambdaConflictHandlerConfig` object when configuring `LAMBDA` as the Conflict Handler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConflictHandlerConfigProperty := &lambdaConflictHandlerConfigProperty{
//   	lambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   }
//
type CfnFunctionConfiguration_LambdaConflictHandlerConfigProperty struct {
	// The Amazon Resource Name (ARN) for the Lambda function to use as the Conflict Handler.
	LambdaConflictHandlerArn *string `field:"optional" json:"lambdaConflictHandlerArn" yaml:"lambdaConflictHandlerArn"`
}

// Describes a Sync configuration for a resolver.
//
// Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syncConfigProperty := &syncConfigProperty{
//   	conflictDetection: jsii.String("conflictDetection"),
//
//   	// the properties below are optional
//   	conflictHandler: jsii.String("conflictHandler"),
//   	lambdaConflictHandlerConfig: &lambdaConflictHandlerConfigProperty{
//   		lambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   	},
//   }
//
type CfnFunctionConfiguration_SyncConfigProperty struct {
	// The Conflict Detection strategy to use.
	//
	// - *VERSION* : Detect conflicts based on object versions for this resolver.
	// - *NONE* : Do not detect conflicts when invoking this resolver.
	ConflictDetection *string `field:"required" json:"conflictDetection" yaml:"conflictDetection"`
	// The Conflict Resolution strategy to perform in the event of a conflict.
	//
	// - *OPTIMISTIC_CONCURRENCY* : Resolve conflicts by rejecting mutations when versions don't match the latest version at the server.
	// - *AUTOMERGE* : Resolve conflicts with the Automerge conflict resolution strategy.
	// - *LAMBDA* : Resolve conflicts with an AWS Lambda function supplied in the `LambdaConflictHandlerConfig` .
	ConflictHandler *string `field:"optional" json:"conflictHandler" yaml:"conflictHandler"`
	// The `LambdaConflictHandlerConfig` when configuring `LAMBDA` as the Conflict Handler.
	LambdaConflictHandlerConfig interface{} `field:"optional" json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

// Properties for defining a `CfnFunctionConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFunctionConfigurationProps := &cfnFunctionConfigurationProps{
//   	apiId: jsii.String("apiId"),
//   	dataSourceName: jsii.String("dataSourceName"),
//   	functionVersion: jsii.String("functionVersion"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	maxBatchSize: jsii.Number(123),
//   	requestMappingTemplate: jsii.String("requestMappingTemplate"),
//   	requestMappingTemplateS3Location: jsii.String("requestMappingTemplateS3Location"),
//   	responseMappingTemplate: jsii.String("responseMappingTemplate"),
//   	responseMappingTemplateS3Location: jsii.String("responseMappingTemplateS3Location"),
//   	syncConfig: &syncConfigProperty{
//   		conflictDetection: jsii.String("conflictDetection"),
//
//   		// the properties below are optional
//   		conflictHandler: jsii.String("conflictHandler"),
//   		lambdaConflictHandlerConfig: &lambdaConflictHandlerConfigProperty{
//   			lambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   		},
//   	},
//   }
//
type CfnFunctionConfigurationProps struct {
	// The AWS AppSync GraphQL API that you want to attach using this function.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The name of data source this function will attach.
	DataSourceName *string `field:"required" json:"dataSourceName" yaml:"dataSourceName"`
	// The version of the request mapping template.
	//
	// Currently, only the 2018-05-29 version of the template is supported.
	FunctionVersion *string `field:"required" json:"functionVersion" yaml:"functionVersion"`
	// The name of the function.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The `Function` description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a `BatchInvoke` operation.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// The `Function` request mapping template.
	//
	// Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate *string `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// Describes a Sync configuration for a resolver.
	//
	// Contains information on which Conflict Detection, as well as Resolution strategy, should be performed when the resolver is invoked.
	RequestMappingTemplateS3Location *string `field:"optional" json:"requestMappingTemplateS3Location" yaml:"requestMappingTemplateS3Location"`
	// The `Function` response mapping template.
	ResponseMappingTemplate *string `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The location of a response mapping template in an Amazon S3 bucket.
	//
	// Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
	ResponseMappingTemplateS3Location *string `field:"optional" json:"responseMappingTemplateS3Location" yaml:"responseMappingTemplateS3Location"`
	// Describes a Sync configuration for a resolver.
	//
	// Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
	SyncConfig interface{} `field:"optional" json:"syncConfig" yaml:"syncConfig"`
}

// A CloudFormation `AWS::AppSync::GraphQLApi`.
//
// The `AWS::AppSync::GraphQLApi` resource creates a new AWS AppSync GraphQL API. This is the top-level construct for your application. For more information, see [Quick Start](https://docs.aws.amazon.com/appsync/latest/devguide/quickstart.html) in the *AWS AppSync Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphQLApi := awscdk.Aws_appsync.NewCfnGraphQLApi(this, jsii.String("MyCfnGraphQLApi"), &cfnGraphQLApiProps{
//   	authenticationType: jsii.String("authenticationType"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	additionalAuthenticationProviders: []interface{}{
//   		&additionalAuthenticationProviderProperty{
//   			authenticationType: jsii.String("authenticationType"),
//
//   			// the properties below are optional
//   			lambdaAuthorizerConfig: &lambdaAuthorizerConfigProperty{
//   				authorizerResultTtlInSeconds: jsii.Number(123),
//   				authorizerUri: jsii.String("authorizerUri"),
//   				identityValidationExpression: jsii.String("identityValidationExpression"),
//   			},
//   			openIdConnectConfig: &openIDConnectConfigProperty{
//   				authTtl: jsii.Number(123),
//   				clientId: jsii.String("clientId"),
//   				iatTtl: jsii.Number(123),
//   				issuer: jsii.String("issuer"),
//   			},
//   			userPoolConfig: &cognitoUserPoolConfigProperty{
//   				appIdClientRegex: jsii.String("appIdClientRegex"),
//   				awsRegion: jsii.String("awsRegion"),
//   				userPoolId: jsii.String("userPoolId"),
//   			},
//   		},
//   	},
//   	lambdaAuthorizerConfig: &lambdaAuthorizerConfigProperty{
//   		authorizerResultTtlInSeconds: jsii.Number(123),
//   		authorizerUri: jsii.String("authorizerUri"),
//   		identityValidationExpression: jsii.String("identityValidationExpression"),
//   	},
//   	logConfig: &logConfigProperty{
//   		cloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   		excludeVerboseContent: jsii.Boolean(false),
//   		fieldLogLevel: jsii.String("fieldLogLevel"),
//   	},
//   	openIdConnectConfig: &openIDConnectConfigProperty{
//   		authTtl: jsii.Number(123),
//   		clientId: jsii.String("clientId"),
//   		iatTtl: jsii.Number(123),
//   		issuer: jsii.String("issuer"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userPoolConfig: &userPoolConfigProperty{
//   		appIdClientRegex: jsii.String("appIdClientRegex"),
//   		awsRegion: jsii.String("awsRegion"),
//   		defaultAction: jsii.String("defaultAction"),
//   		userPoolId: jsii.String("userPoolId"),
//   	},
//   	xrayEnabled: jsii.Boolean(false),
//   })
//
type CfnGraphQLApi interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A list of additional authentication providers for the `GraphqlApi` API.
	AdditionalAuthenticationProviders() interface{}
	SetAdditionalAuthenticationProviders(val interface{})
	// Unique AWS AppSync GraphQL API identifier.
	AttrApiId() *string
	// The Amazon Resource Name (ARN) of the API key, such as `arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid` .
	AttrArn() *string
	// The Endpoint URL of your GraphQL API.
	AttrGraphQlUrl() *string
	// Security configuration for your GraphQL API.
	//
	// For allowed values (such as `API_KEY` , `AWS_IAM` , `AMAZON_COGNITO_USER_POOLS` , `OPENID_CONNECT` , or `AWS_LAMBDA` ), see [Security](https://docs.aws.amazon.com/appsync/latest/devguide/security.html) in the *AWS AppSync Developer Guide* .
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A `LambdaAuthorizerConfig` holds configuration on how to authorize AWS AppSync API access when using the `AWS_LAMBDA` authorizer mode.
	//
	// Be aware that an AWS AppSync API may have only one Lambda authorizer configured at a time.
	LambdaAuthorizerConfig() interface{}
	SetLambdaAuthorizerConfig(val interface{})
	// The Amazon CloudWatch Logs configuration.
	LogConfig() interface{}
	SetLogConfig(val interface{})
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
	// The API name.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The OpenID Connect configuration.
	OpenIdConnectConfig() interface{}
	SetOpenIdConnectConfig(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An arbitrary set of tags (key-value pairs) for this GraphQL API.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Optional authorization configuration for using Amazon Cognito user pools with your GraphQL endpoint.
	UserPoolConfig() interface{}
	SetUserPoolConfig(val interface{})
	// A flag indicating whether to use AWS X-Ray tracing for this `GraphqlApi` .
	XrayEnabled() interface{}
	SetXrayEnabled(val interface{})
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

// The jsii proxy struct for CfnGraphQLApi
type jsiiProxy_CfnGraphQLApi struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGraphQLApi) AdditionalAuthenticationProviders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalAuthenticationProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrGraphQlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGraphQlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) LambdaAuthorizerConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) LogConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) OpenIdConnectConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openIdConnectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) UserPoolConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) XrayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xrayEnabled",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::GraphQLApi`.
func NewCfnGraphQLApi(scope constructs.Construct, id *string, props *CfnGraphQLApiProps) CfnGraphQLApi {
	_init_.Initialize()

	j := jsiiProxy_CfnGraphQLApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::GraphQLApi`.
func NewCfnGraphQLApi_Override(c CfnGraphQLApi, scope constructs.Construct, id *string, props *CfnGraphQLApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetAdditionalAuthenticationProviders(val interface{}) {
	_jsii_.Set(
		j,
		"additionalAuthenticationProviders",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetAuthenticationType(val *string) {
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetLambdaAuthorizerConfig(val interface{}) {
	_jsii_.Set(
		j,
		"lambdaAuthorizerConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetLogConfig(val interface{}) {
	_jsii_.Set(
		j,
		"logConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetOpenIdConnectConfig(val interface{}) {
	_jsii_.Set(
		j,
		"openIdConnectConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetUserPoolConfig(val interface{}) {
	_jsii_.Set(
		j,
		"userPoolConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetXrayEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"xrayEnabled",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnGraphQLApi_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnGraphQLApi_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnGraphQLApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGraphQLApi_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.CfnGraphQLApi",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGraphQLApi) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLApi) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLApi) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGraphQLApi) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLApi) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLApi) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes an additional authentication provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   additionalAuthenticationProviderProperty := &additionalAuthenticationProviderProperty{
//   	authenticationType: jsii.String("authenticationType"),
//
//   	// the properties below are optional
//   	lambdaAuthorizerConfig: &lambdaAuthorizerConfigProperty{
//   		authorizerResultTtlInSeconds: jsii.Number(123),
//   		authorizerUri: jsii.String("authorizerUri"),
//   		identityValidationExpression: jsii.String("identityValidationExpression"),
//   	},
//   	openIdConnectConfig: &openIDConnectConfigProperty{
//   		authTtl: jsii.Number(123),
//   		clientId: jsii.String("clientId"),
//   		iatTtl: jsii.Number(123),
//   		issuer: jsii.String("issuer"),
//   	},
//   	userPoolConfig: &cognitoUserPoolConfigProperty{
//   		appIdClientRegex: jsii.String("appIdClientRegex"),
//   		awsRegion: jsii.String("awsRegion"),
//   		userPoolId: jsii.String("userPoolId"),
//   	},
//   }
//
type CfnGraphQLApi_AdditionalAuthenticationProviderProperty struct {
	// The authentication type for API key, AWS Identity and Access Management , OIDC, Amazon Cognito user pools , or AWS Lambda .
	//
	// Valid Values: `API_KEY` | `AWS_IAM` | `OPENID_CONNECT` | `AMAZON_COGNITO_USER_POOLS` | `AWS_LAMBDA`.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// Configuration for AWS Lambda function authorization.
	LambdaAuthorizerConfig interface{} `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// The OIDC configuration.
	OpenIdConnectConfig interface{} `field:"optional" json:"openIdConnectConfig" yaml:"openIdConnectConfig"`
	// The Amazon Cognito user pool configuration.
	UserPoolConfig interface{} `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
}

// Describes an Amazon Cognito user pool configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoUserPoolConfigProperty := &cognitoUserPoolConfigProperty{
//   	appIdClientRegex: jsii.String("appIdClientRegex"),
//   	awsRegion: jsii.String("awsRegion"),
//   	userPoolId: jsii.String("userPoolId"),
//   }
//
type CfnGraphQLApi_CognitoUserPoolConfigProperty struct {
	// A regular expression for validating the incoming Amazon Cognito user pool app client ID.
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// The AWS Region in which the user pool was created.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The user pool ID.
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}

// Configuration for AWS Lambda function authorization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaAuthorizerConfigProperty := &lambdaAuthorizerConfigProperty{
//   	authorizerResultTtlInSeconds: jsii.Number(123),
//   	authorizerUri: jsii.String("authorizerUri"),
//   	identityValidationExpression: jsii.String("identityValidationExpression"),
//   }
//
type CfnGraphQLApi_LambdaAuthorizerConfigProperty struct {
	// The number of seconds a response should be cached for.
	//
	// The default is 5 minutes (300 seconds). The Lambda function can override this by returning a `ttlOverride` key in its response. A value of 0 disables caching of responses.
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// The ARN of the Lambda function to be called for authorization.
	//
	// This may be a standard Lambda ARN, a version ARN ( `.../v3` ) or alias ARN.
	//
	// *Note* : This Lambda function must have the following resource-based policy assigned to it. When configuring Lambda authorizers in the console, this is done for you. To do so with the AWS CLI , run the following:
	//
	// `aws lambda add-permission --function-name "arn:aws:lambda:us-east-2:111122223333:function:my-function" --statement-id "appsync" --principal appsync.amazonaws.com --action lambda:InvokeFunction`
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
	// A regular expression for validation of tokens before the Lambda function is called.
	IdentityValidationExpression *string `field:"optional" json:"identityValidationExpression" yaml:"identityValidationExpression"`
}

// The `LogConfig` property type specifies the logging configuration when writing GraphQL operations and tracing to Amazon CloudWatch for an AWS AppSync GraphQL API.
//
// `LogConfig` is a property of the [AWS::AppSync::GraphQLApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigProperty := &logConfigProperty{
//   	cloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   	excludeVerboseContent: jsii.Boolean(false),
//   	fieldLogLevel: jsii.String("fieldLogLevel"),
//   }
//
type CfnGraphQLApi_LogConfigProperty struct {
	// The service role that AWS AppSync will assume to publish to Amazon CloudWatch Logs in your account.
	CloudWatchLogsRoleArn *string `field:"optional" json:"cloudWatchLogsRoleArn" yaml:"cloudWatchLogsRoleArn"`
	// Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging level.
	ExcludeVerboseContent interface{} `field:"optional" json:"excludeVerboseContent" yaml:"excludeVerboseContent"`
	// The field logging level. Values can be NONE, ERROR, or ALL.
	//
	// - *NONE* : No field-level logs are captured.
	// - *ERROR* : Logs the following information only for the fields that are in error:
	//
	// - The error section in the server response.
	// - Field-level errors.
	// - The generated request/response functions that got resolved for error fields.
	// - *ALL* : The following information is logged for all fields in the query:
	//
	// - Field-level tracing information.
	// - The generated request/response functions that got resolved for each field.
	FieldLogLevel *string `field:"optional" json:"fieldLogLevel" yaml:"fieldLogLevel"`
}

// The `OpenIDConnectConfig` property type specifies the optional authorization configuration for using an OpenID Connect compliant service with your GraphQL endpoint for an AWS AppSync GraphQL API.
//
// `OpenIDConnectConfig` is a property of the [AWS::AppSync::GraphQLApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openIDConnectConfigProperty := &openIDConnectConfigProperty{
//   	authTtl: jsii.Number(123),
//   	clientId: jsii.String("clientId"),
//   	iatTtl: jsii.Number(123),
//   	issuer: jsii.String("issuer"),
//   }
//
type CfnGraphQLApi_OpenIDConnectConfigProperty struct {
	// The number of milliseconds that a token is valid after being authenticated.
	AuthTtl *float64 `field:"optional" json:"authTtl" yaml:"authTtl"`
	// The client identifier of the Relying party at the OpenID identity provider.
	//
	// This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so that AWS AppSync can validate against multiple client identifiers at a time.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The number of milliseconds that a token is valid after it's issued to a user.
	IatTtl *float64 `field:"optional" json:"iatTtl" yaml:"iatTtl"`
	// The issuer for the OIDC configuration.
	//
	// The issuer returned by discovery must exactly match the value of `iss` in the ID token.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
}

// The `UserPoolConfig` property type specifies the optional authorization configuration for using Amazon Cognito user pools with your GraphQL endpoint for an AWS AppSync GraphQL API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolConfigProperty := &userPoolConfigProperty{
//   	appIdClientRegex: jsii.String("appIdClientRegex"),
//   	awsRegion: jsii.String("awsRegion"),
//   	defaultAction: jsii.String("defaultAction"),
//   	userPoolId: jsii.String("userPoolId"),
//   }
//
type CfnGraphQLApi_UserPoolConfigProperty struct {
	// A regular expression for validating the incoming Amazon Cognito user pool app client ID.
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// The AWS Region in which the user pool was created.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The action that you want your GraphQL API to take when a request that uses Amazon Cognito user pool authentication doesn't match the Amazon Cognito user pool configuration.
	//
	// When specifying Amazon Cognito user pools as the default authentication, you must set the value for `DefaultAction` to `ALLOW` if specifying `AdditionalAuthenticationProviders` .
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// The user pool ID.
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}

// Properties for defining a `CfnGraphQLApi`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphQLApiProps := &cfnGraphQLApiProps{
//   	authenticationType: jsii.String("authenticationType"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	additionalAuthenticationProviders: []interface{}{
//   		&additionalAuthenticationProviderProperty{
//   			authenticationType: jsii.String("authenticationType"),
//
//   			// the properties below are optional
//   			lambdaAuthorizerConfig: &lambdaAuthorizerConfigProperty{
//   				authorizerResultTtlInSeconds: jsii.Number(123),
//   				authorizerUri: jsii.String("authorizerUri"),
//   				identityValidationExpression: jsii.String("identityValidationExpression"),
//   			},
//   			openIdConnectConfig: &openIDConnectConfigProperty{
//   				authTtl: jsii.Number(123),
//   				clientId: jsii.String("clientId"),
//   				iatTtl: jsii.Number(123),
//   				issuer: jsii.String("issuer"),
//   			},
//   			userPoolConfig: &cognitoUserPoolConfigProperty{
//   				appIdClientRegex: jsii.String("appIdClientRegex"),
//   				awsRegion: jsii.String("awsRegion"),
//   				userPoolId: jsii.String("userPoolId"),
//   			},
//   		},
//   	},
//   	lambdaAuthorizerConfig: &lambdaAuthorizerConfigProperty{
//   		authorizerResultTtlInSeconds: jsii.Number(123),
//   		authorizerUri: jsii.String("authorizerUri"),
//   		identityValidationExpression: jsii.String("identityValidationExpression"),
//   	},
//   	logConfig: &logConfigProperty{
//   		cloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   		excludeVerboseContent: jsii.Boolean(false),
//   		fieldLogLevel: jsii.String("fieldLogLevel"),
//   	},
//   	openIdConnectConfig: &openIDConnectConfigProperty{
//   		authTtl: jsii.Number(123),
//   		clientId: jsii.String("clientId"),
//   		iatTtl: jsii.Number(123),
//   		issuer: jsii.String("issuer"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userPoolConfig: &userPoolConfigProperty{
//   		appIdClientRegex: jsii.String("appIdClientRegex"),
//   		awsRegion: jsii.String("awsRegion"),
//   		defaultAction: jsii.String("defaultAction"),
//   		userPoolId: jsii.String("userPoolId"),
//   	},
//   	xrayEnabled: jsii.Boolean(false),
//   }
//
type CfnGraphQLApiProps struct {
	// Security configuration for your GraphQL API.
	//
	// For allowed values (such as `API_KEY` , `AWS_IAM` , `AMAZON_COGNITO_USER_POOLS` , `OPENID_CONNECT` , or `AWS_LAMBDA` ), see [Security](https://docs.aws.amazon.com/appsync/latest/devguide/security.html) in the *AWS AppSync Developer Guide* .
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The API name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of additional authentication providers for the `GraphqlApi` API.
	AdditionalAuthenticationProviders interface{} `field:"optional" json:"additionalAuthenticationProviders" yaml:"additionalAuthenticationProviders"`
	// A `LambdaAuthorizerConfig` holds configuration on how to authorize AWS AppSync API access when using the `AWS_LAMBDA` authorizer mode.
	//
	// Be aware that an AWS AppSync API may have only one Lambda authorizer configured at a time.
	LambdaAuthorizerConfig interface{} `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// The Amazon CloudWatch Logs configuration.
	LogConfig interface{} `field:"optional" json:"logConfig" yaml:"logConfig"`
	// The OpenID Connect configuration.
	OpenIdConnectConfig interface{} `field:"optional" json:"openIdConnectConfig" yaml:"openIdConnectConfig"`
	// An arbitrary set of tags (key-value pairs) for this GraphQL API.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Optional authorization configuration for using Amazon Cognito user pools with your GraphQL endpoint.
	UserPoolConfig interface{} `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
	// A flag indicating whether to use AWS X-Ray tracing for this `GraphqlApi` .
	XrayEnabled interface{} `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

// A CloudFormation `AWS::AppSync::GraphQLSchema`.
//
// The `AWS::AppSync::GraphQLSchema` resource is used for your AWS AppSync GraphQL schema that controls the data model for your API. Schema files are text written in Schema Definition Language (SDL) format. For more information about schema authoring, see [Designing a GraphQL API](https://docs.aws.amazon.com/appsync/latest/devguide/designing-a-graphql-api.html) in the *AWS AppSync Developer Guide* .
//
// > When you submit an update, AWS CloudFormation updates resources based on differences between what you submit and the stack's current template. To cause this resource to be updated you must change a property value for this resource in the CloudFormation template. Changing the Amazon S3 file content without changing a property value will not result in an update operation.
// >
// > See [Update Behaviors of Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html) in the *AWS CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphQLSchema := awscdk.Aws_appsync.NewCfnGraphQLSchema(this, jsii.String("MyCfnGraphQLSchema"), &cfnGraphQLSchemaProps{
//   	apiId: jsii.String("apiId"),
//
//   	// the properties below are optional
//   	definition: jsii.String("definition"),
//   	definitionS3Location: jsii.String("definitionS3Location"),
//   })
//
type CfnGraphQLSchema interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The AWS AppSync GraphQL API identifier to which you want to apply this schema.
	ApiId() *string
	SetApiId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The text representation of a GraphQL schema in SDL format.
	//
	// For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref) .
	Definition() *string
	SetDefinition(val *string)
	// The location of a GraphQL schema file in an Amazon S3 bucket.
	//
	// Use this if you want to provision with the schema living in Amazon S3 rather than embedding it in your CloudFormation template.
	DefinitionS3Location() *string
	SetDefinitionS3Location(val *string)
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

// The jsii proxy struct for CfnGraphQLSchema
type jsiiProxy_CfnGraphQLSchema struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGraphQLSchema) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) Definition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) DefinitionS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::GraphQLSchema`.
func NewCfnGraphQLSchema(scope constructs.Construct, id *string, props *CfnGraphQLSchemaProps) CfnGraphQLSchema {
	_init_.Initialize()

	j := jsiiProxy_CfnGraphQLSchema{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnGraphQLSchema",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::GraphQLSchema`.
func NewCfnGraphQLSchema_Override(c CfnGraphQLSchema, scope constructs.Construct, id *string, props *CfnGraphQLSchemaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnGraphQLSchema",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGraphQLSchema) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLSchema) SetDefinition(val *string) {
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLSchema) SetDefinitionS3Location(val *string) {
	_jsii_.Set(
		j,
		"definitionS3Location",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnGraphQLSchema_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnGraphQLSchema",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnGraphQLSchema_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnGraphQLSchema",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnGraphQLSchema_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnGraphQLSchema",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGraphQLSchema_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.CfnGraphQLSchema",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGraphQLSchema) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGraphQLSchema) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGraphQLSchema) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGraphQLSchema) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGraphQLSchema) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGraphQLSchema) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGraphQLSchema) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnGraphQLSchema) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLSchema) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLSchema) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnGraphQLSchema) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGraphQLSchema) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLSchema) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLSchema) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLSchema) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnGraphQLSchema`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphQLSchemaProps := &cfnGraphQLSchemaProps{
//   	apiId: jsii.String("apiId"),
//
//   	// the properties below are optional
//   	definition: jsii.String("definition"),
//   	definitionS3Location: jsii.String("definitionS3Location"),
//   }
//
type CfnGraphQLSchemaProps struct {
	// The AWS AppSync GraphQL API identifier to which you want to apply this schema.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The text representation of a GraphQL schema in SDL format.
	//
	// For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref) .
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// The location of a GraphQL schema file in an Amazon S3 bucket.
	//
	// Use this if you want to provision with the schema living in Amazon S3 rather than embedding it in your CloudFormation template.
	DefinitionS3Location *string `field:"optional" json:"definitionS3Location" yaml:"definitionS3Location"`
}

// A CloudFormation `AWS::AppSync::Resolver`.
//
// The `AWS::AppSync::Resolver` resource defines the logical GraphQL resolver that you attach to fields in a schema. Request and response templates for resolvers are written in Apache Velocity Template Language (VTL) format. For more information about resolvers, see [Resolver Mapping Template Reference](https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference.html) .
//
// > When you submit an update, AWS CloudFormation updates resources based on differences between what you submit and the stack's current template. To cause this resource to be updated you must change a property value for this resource in the CloudFormation template. Changing the Amazon S3 file content without changing a property value will not result in an update operation.
// >
// > See [Update Behaviors of Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html) in the *AWS CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolver := awscdk.Aws_appsync.NewCfnResolver(this, jsii.String("MyCfnResolver"), &cfnResolverProps{
//   	apiId: jsii.String("apiId"),
//   	fieldName: jsii.String("fieldName"),
//   	typeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	cachingConfig: &cachingConfigProperty{
//   		ttl: jsii.Number(123),
//
//   		// the properties below are optional
//   		cachingKeys: []*string{
//   			jsii.String("cachingKeys"),
//   		},
//   	},
//   	dataSourceName: jsii.String("dataSourceName"),
//   	kind: jsii.String("kind"),
//   	maxBatchSize: jsii.Number(123),
//   	pipelineConfig: &pipelineConfigProperty{
//   		functions: []*string{
//   			jsii.String("functions"),
//   		},
//   	},
//   	requestMappingTemplate: jsii.String("requestMappingTemplate"),
//   	requestMappingTemplateS3Location: jsii.String("requestMappingTemplateS3Location"),
//   	responseMappingTemplate: jsii.String("responseMappingTemplate"),
//   	responseMappingTemplateS3Location: jsii.String("responseMappingTemplateS3Location"),
//   	syncConfig: &syncConfigProperty{
//   		conflictDetection: jsii.String("conflictDetection"),
//
//   		// the properties below are optional
//   		conflictHandler: jsii.String("conflictHandler"),
//   		lambdaConflictHandlerConfig: &lambdaConflictHandlerConfigProperty{
//   			lambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   		},
//   	},
//   })
//
type CfnResolver interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The AWS AppSync GraphQL API to which you want to attach this resolver.
	ApiId() *string
	SetApiId(val *string)
	// The GraphQL field on a type that invokes the resolver.
	AttrFieldName() *string
	// ARN of the resolver, such as `arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/types/typename/resolvers/resolvername` .
	AttrResolverArn() *string
	// The GraphQL type that invokes this resolver.
	AttrTypeName() *string
	// The caching configuration for the resolver.
	CachingConfig() interface{}
	SetCachingConfig(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The resolver data source name.
	DataSourceName() *string
	SetDataSourceName(val *string)
	// The GraphQL field on a type that invokes the resolver.
	FieldName() *string
	SetFieldName(val *string)
	// The resolver type.
	//
	// - *UNIT* : A UNIT resolver type. A UNIT resolver is the default resolver type. You can use a UNIT resolver to run a GraphQL query against a single data source.
	// - *PIPELINE* : A PIPELINE resolver type. You can use a PIPELINE resolver to invoke a series of `Function` objects in a serial manner. You can use a pipeline resolver to run a GraphQL query against multiple data sources.
	Kind() *string
	SetKind(val *string)
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
	// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a `BatchInvoke` operation.
	MaxBatchSize() *float64
	SetMaxBatchSize(val *float64)
	// The tree node.
	Node() constructs.Node
	// Functions linked with the pipeline resolver.
	PipelineConfig() interface{}
	SetPipelineConfig(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The request mapping template.
	//
	// Request mapping templates are optional when using a Lambda data source. For all other data sources, a request mapping template is required.
	RequestMappingTemplate() *string
	SetRequestMappingTemplate(val *string)
	// The location of a request mapping template in an Amazon S3 bucket.
	//
	// Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
	RequestMappingTemplateS3Location() *string
	SetRequestMappingTemplateS3Location(val *string)
	// The response mapping template.
	ResponseMappingTemplate() *string
	SetResponseMappingTemplate(val *string)
	// The location of a response mapping template in an Amazon S3 bucket.
	//
	// Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
	ResponseMappingTemplateS3Location() *string
	SetResponseMappingTemplateS3Location(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The `SyncConfig` for a resolver attached to a versioned data source.
	SyncConfig() interface{}
	SetSyncConfig(val interface{})
	// The GraphQL type that invokes this resolver.
	TypeName() *string
	SetTypeName(val *string)
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

// The jsii proxy struct for CfnResolver
type jsiiProxy_CfnResolver struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResolver) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) AttrFieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) AttrResolverArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResolverArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) AttrTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) CachingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) DataSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) FieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) MaxBatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) PipelineConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) RequestMappingTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) RequestMappingTemplateS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplateS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) ResponseMappingTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) ResponseMappingTemplateS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplateS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) SyncConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::Resolver`.
func NewCfnResolver(scope constructs.Construct, id *string, props *CfnResolverProps) CfnResolver {
	_init_.Initialize()

	j := jsiiProxy_CfnResolver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnResolver",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::Resolver`.
func NewCfnResolver_Override(c CfnResolver, scope constructs.Construct, id *string, props *CfnResolverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.CfnResolver",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResolver) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetCachingConfig(val interface{}) {
	_jsii_.Set(
		j,
		"cachingConfig",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetDataSourceName(val *string) {
	_jsii_.Set(
		j,
		"dataSourceName",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetFieldName(val *string) {
	_jsii_.Set(
		j,
		"fieldName",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetKind(val *string) {
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetMaxBatchSize(val *float64) {
	_jsii_.Set(
		j,
		"maxBatchSize",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetPipelineConfig(val interface{}) {
	_jsii_.Set(
		j,
		"pipelineConfig",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetRequestMappingTemplate(val *string) {
	_jsii_.Set(
		j,
		"requestMappingTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetRequestMappingTemplateS3Location(val *string) {
	_jsii_.Set(
		j,
		"requestMappingTemplateS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetResponseMappingTemplate(val *string) {
	_jsii_.Set(
		j,
		"responseMappingTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetResponseMappingTemplateS3Location(val *string) {
	_jsii_.Set(
		j,
		"responseMappingTemplateS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetSyncConfig(val interface{}) {
	_jsii_.Set(
		j,
		"syncConfig",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResolver_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnResolver",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResolver_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnResolver",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnResolver_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.CfnResolver",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolver_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.CfnResolver",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResolver) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResolver) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResolver) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResolver) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResolver) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResolver) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResolver) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResolver) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolver) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolver) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResolver) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolver) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolver) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolver) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolver) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The caching configuration for a resolver that has caching activated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cachingConfigProperty := &cachingConfigProperty{
//   	ttl: jsii.Number(123),
//
//   	// the properties below are optional
//   	cachingKeys: []*string{
//   		jsii.String("cachingKeys"),
//   	},
//   }
//
type CfnResolver_CachingConfigProperty struct {
	// The TTL in seconds for a resolver that has caching activated.
	//
	// Valid values are 1–3,600 seconds.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
	// The caching keys for a resolver that has caching activated.
	//
	// Valid values are entries from the `$context.arguments` , `$context.source` , and `$context.identity` maps.
	CachingKeys *[]*string `field:"optional" json:"cachingKeys" yaml:"cachingKeys"`
}

// The `LambdaConflictHandlerConfig` when configuring LAMBDA as the Conflict Handler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConflictHandlerConfigProperty := &lambdaConflictHandlerConfigProperty{
//   	lambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   }
//
type CfnResolver_LambdaConflictHandlerConfigProperty struct {
	// The Amazon Resource Name (ARN) for the Lambda function to use as the Conflict Handler.
	LambdaConflictHandlerArn *string `field:"optional" json:"lambdaConflictHandlerArn" yaml:"lambdaConflictHandlerArn"`
}

// Use the `PipelineConfig` property type to specify `PipelineConfig` for an AWS AppSync resolver.
//
// `PipelineConfig` is a property of the [AWS::AppSync::Resolver](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineConfigProperty := &pipelineConfigProperty{
//   	functions: []*string{
//   		jsii.String("functions"),
//   	},
//   }
//
type CfnResolver_PipelineConfigProperty struct {
	// A list of `Function` objects.
	Functions *[]*string `field:"optional" json:"functions" yaml:"functions"`
}

// Describes a Sync configuration for a resolver.
//
// Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syncConfigProperty := &syncConfigProperty{
//   	conflictDetection: jsii.String("conflictDetection"),
//
//   	// the properties below are optional
//   	conflictHandler: jsii.String("conflictHandler"),
//   	lambdaConflictHandlerConfig: &lambdaConflictHandlerConfigProperty{
//   		lambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   	},
//   }
//
type CfnResolver_SyncConfigProperty struct {
	// The Conflict Detection strategy to use.
	//
	// - *VERSION* : Detect conflicts based on object versions for this resolver.
	// - *NONE* : Do not detect conflicts when invoking this resolver.
	ConflictDetection *string `field:"required" json:"conflictDetection" yaml:"conflictDetection"`
	// The Conflict Resolution strategy to perform in the event of a conflict.
	//
	// - *OPTIMISTIC_CONCURRENCY* : Resolve conflicts by rejecting mutations when versions don't match the latest version at the server.
	// - *AUTOMERGE* : Resolve conflicts with the Automerge conflict resolution strategy.
	// - *LAMBDA* : Resolve conflicts with an AWS Lambda function supplied in the `LambdaConflictHandlerConfig` .
	ConflictHandler *string `field:"optional" json:"conflictHandler" yaml:"conflictHandler"`
	// The `LambdaConflictHandlerConfig` when configuring `LAMBDA` as the Conflict Handler.
	LambdaConflictHandlerConfig interface{} `field:"optional" json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

// Properties for defining a `CfnResolver`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverProps := &cfnResolverProps{
//   	apiId: jsii.String("apiId"),
//   	fieldName: jsii.String("fieldName"),
//   	typeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	cachingConfig: &cachingConfigProperty{
//   		ttl: jsii.Number(123),
//
//   		// the properties below are optional
//   		cachingKeys: []*string{
//   			jsii.String("cachingKeys"),
//   		},
//   	},
//   	dataSourceName: jsii.String("dataSourceName"),
//   	kind: jsii.String("kind"),
//   	maxBatchSize: jsii.Number(123),
//   	pipelineConfig: &pipelineConfigProperty{
//   		functions: []*string{
//   			jsii.String("functions"),
//   		},
//   	},
//   	requestMappingTemplate: jsii.String("requestMappingTemplate"),
//   	requestMappingTemplateS3Location: jsii.String("requestMappingTemplateS3Location"),
//   	responseMappingTemplate: jsii.String("responseMappingTemplate"),
//   	responseMappingTemplateS3Location: jsii.String("responseMappingTemplateS3Location"),
//   	syncConfig: &syncConfigProperty{
//   		conflictDetection: jsii.String("conflictDetection"),
//
//   		// the properties below are optional
//   		conflictHandler: jsii.String("conflictHandler"),
//   		lambdaConflictHandlerConfig: &lambdaConflictHandlerConfigProperty{
//   			lambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   		},
//   	},
//   }
//
type CfnResolverProps struct {
	// The AWS AppSync GraphQL API to which you want to attach this resolver.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The GraphQL field on a type that invokes the resolver.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// The GraphQL type that invokes this resolver.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The caching configuration for the resolver.
	CachingConfig interface{} `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// The resolver data source name.
	DataSourceName *string `field:"optional" json:"dataSourceName" yaml:"dataSourceName"`
	// The resolver type.
	//
	// - *UNIT* : A UNIT resolver type. A UNIT resolver is the default resolver type. You can use a UNIT resolver to run a GraphQL query against a single data source.
	// - *PIPELINE* : A PIPELINE resolver type. You can use a PIPELINE resolver to invoke a series of `Function` objects in a serial manner. You can use a pipeline resolver to run a GraphQL query against multiple data sources.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a `BatchInvoke` operation.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// Functions linked with the pipeline resolver.
	PipelineConfig interface{} `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// The request mapping template.
	//
	// Request mapping templates are optional when using a Lambda data source. For all other data sources, a request mapping template is required.
	RequestMappingTemplate *string `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// The location of a request mapping template in an Amazon S3 bucket.
	//
	// Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
	RequestMappingTemplateS3Location *string `field:"optional" json:"requestMappingTemplateS3Location" yaml:"requestMappingTemplateS3Location"`
	// The response mapping template.
	ResponseMappingTemplate *string `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The location of a response mapping template in an Amazon S3 bucket.
	//
	// Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
	ResponseMappingTemplateS3Location *string `field:"optional" json:"responseMappingTemplateS3Location" yaml:"responseMappingTemplateS3Location"`
	// The `SyncConfig` for a resolver attached to a versioned data source.
	SyncConfig interface{} `field:"optional" json:"syncConfig" yaml:"syncConfig"`
}

