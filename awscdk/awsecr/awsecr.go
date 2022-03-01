package awsecr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

// Authorization token to access private ECR repositories in the current environment via Docker CLI.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry_auth.html
//
type AuthorizationToken interface {
}

// The jsii proxy struct for AuthorizationToken
type jsiiProxy_AuthorizationToken struct {
	_ byte // padding
}

// Grant access to retrieve an authorization token.
func AuthorizationToken_GrantRead(grantee awsiam.IGrantable) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_ecr.AuthorizationToken",
		"grantRead",
		[]interface{}{grantee},
	)
}

// A CloudFormation `AWS::ECR::PublicRepository`.
//
// The `AWS::ECR::PublicRepository` resource specifies an Amazon Elastic Container Registry Public (Amazon ECR Public) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see [Amazon ECR public repositories](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repositories.html) in the *Amazon ECR Public User Guide* .
//
// TODO: EXAMPLE
//
type CfnPublicRepository interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RepositoryCatalogData() interface{}
	SetRepositoryCatalogData(val interface{})
	RepositoryName() *string
	SetRepositoryName(val *string)
	RepositoryPolicyText() interface{}
	SetRepositoryPolicyText(val interface{})
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

// The jsii proxy struct for CfnPublicRepository
type jsiiProxy_CfnPublicRepository struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPublicRepository) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) RepositoryCatalogData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryCatalogData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) RepositoryPolicyText() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryPolicyText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicRepository) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECR::PublicRepository`.
func NewCfnPublicRepository(scope constructs.Construct, id *string, props *CfnPublicRepositoryProps) CfnPublicRepository {
	_init_.Initialize()

	j := jsiiProxy_CfnPublicRepository{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.CfnPublicRepository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECR::PublicRepository`.
func NewCfnPublicRepository_Override(c CfnPublicRepository, scope constructs.Construct, id *string, props *CfnPublicRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.CfnPublicRepository",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPublicRepository) SetRepositoryCatalogData(val interface{}) {
	_jsii_.Set(
		j,
		"repositoryCatalogData",
		val,
	)
}

func (j *jsiiProxy_CfnPublicRepository) SetRepositoryName(val *string) {
	_jsii_.Set(
		j,
		"repositoryName",
		val,
	)
}

func (j *jsiiProxy_CfnPublicRepository) SetRepositoryPolicyText(val interface{}) {
	_jsii_.Set(
		j,
		"repositoryPolicyText",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPublicRepository_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnPublicRepository",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPublicRepository_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnPublicRepository",
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
func CfnPublicRepository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnPublicRepository",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPublicRepository_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr.CfnPublicRepository",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnPublicRepository) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnPublicRepository) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnPublicRepository) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnPublicRepository) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnPublicRepository) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnPublicRepository) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnPublicRepository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnPublicRepository) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnPublicRepository) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnPublicRepository) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnPublicRepository) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPublicRepository) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnPublicRepository) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnPublicRepository) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicRepository) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPublicRepository`.
//
// TODO: EXAMPLE
//
type CfnPublicRepositoryProps struct {
	// The details about the repository that are publicly visible in the Amazon ECR Public Gallery.
	//
	// For more information, see [Amazon ECR Public repository catalog data](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repository-catalog-data.html) in the *Amazon ECR Public User Guide* .
	RepositoryCatalogData interface{} `json:"repositoryCatalogData" yaml:"repositoryCatalogData"`
	// The name to use for the public repository.
	//
	// The repository name may be specified on its own (such as `nginx-web-app` ) or it can be prepended with a namespace to group the repository into a category (such as `project-a/nginx-web-app` ). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
	// The JSON repository policy text to apply to the public repository.
	//
	// For more information, see [Amazon ECR Public repository policies](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repository-policies.html) in the *Amazon ECR Public User Guide* .
	RepositoryPolicyText interface{} `json:"repositoryPolicyText" yaml:"repositoryPolicyText"`
	// An array of key-value pairs to apply to this resource.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ECR::PullThroughCacheRule`.
//
// Creates a pull through cache rule. A pull through cache rule provides a way to cache images from an external public registry in your Amazon ECR private registry.
//
// TODO: EXAMPLE
//
type CfnPullThroughCacheRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EcrRepositoryPrefix() *string
	SetEcrRepositoryPrefix(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	UpstreamRegistryUrl() *string
	SetUpstreamRegistryUrl(val *string)
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

// The jsii proxy struct for CfnPullThroughCacheRule
type jsiiProxy_CfnPullThroughCacheRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPullThroughCacheRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPullThroughCacheRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPullThroughCacheRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPullThroughCacheRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPullThroughCacheRule) EcrRepositoryPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecrRepositoryPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPullThroughCacheRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPullThroughCacheRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPullThroughCacheRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPullThroughCacheRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPullThroughCacheRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPullThroughCacheRule) UpstreamRegistryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upstreamRegistryUrl",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECR::PullThroughCacheRule`.
func NewCfnPullThroughCacheRule(scope constructs.Construct, id *string, props *CfnPullThroughCacheRuleProps) CfnPullThroughCacheRule {
	_init_.Initialize()

	j := jsiiProxy_CfnPullThroughCacheRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.CfnPullThroughCacheRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECR::PullThroughCacheRule`.
func NewCfnPullThroughCacheRule_Override(c CfnPullThroughCacheRule, scope constructs.Construct, id *string, props *CfnPullThroughCacheRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.CfnPullThroughCacheRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPullThroughCacheRule) SetEcrRepositoryPrefix(val *string) {
	_jsii_.Set(
		j,
		"ecrRepositoryPrefix",
		val,
	)
}

func (j *jsiiProxy_CfnPullThroughCacheRule) SetUpstreamRegistryUrl(val *string) {
	_jsii_.Set(
		j,
		"upstreamRegistryUrl",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPullThroughCacheRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnPullThroughCacheRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPullThroughCacheRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnPullThroughCacheRule",
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
func CfnPullThroughCacheRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnPullThroughCacheRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPullThroughCacheRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr.CfnPullThroughCacheRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnPullThroughCacheRule) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnPullThroughCacheRule) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnPullThroughCacheRule) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnPullThroughCacheRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnPullThroughCacheRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnPullThroughCacheRule) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnPullThroughCacheRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnPullThroughCacheRule) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnPullThroughCacheRule) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnPullThroughCacheRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnPullThroughCacheRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPullThroughCacheRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnPullThroughCacheRule) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnPullThroughCacheRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPullThroughCacheRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPullThroughCacheRule`.
//
// TODO: EXAMPLE
//
type CfnPullThroughCacheRuleProps struct {
	// The Amazon ECR repository prefix associated with the pull through cache rule.
	EcrRepositoryPrefix *string `json:"ecrRepositoryPrefix" yaml:"ecrRepositoryPrefix"`
	// The upstream registry URL associated with the pull through cache rule.
	UpstreamRegistryUrl *string `json:"upstreamRegistryUrl" yaml:"upstreamRegistryUrl"`
}

// A CloudFormation `AWS::ECR::RegistryPolicy`.
//
// The `AWS::ECR::RegistryPolicy` resource creates or updates the permissions policy for a private registry.
//
// A private registry policy is used to specify permissions for another AWS account and is used when configuring cross-account replication. For more information, see [Registry permissions](https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html) in the *Amazon Elastic Container Registry User Guide* .
//
// TODO: EXAMPLE
//
type CfnRegistryPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrRegistryId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	PolicyText() interface{}
	SetPolicyText(val interface{})
	Ref() *string
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

// The jsii proxy struct for CfnRegistryPolicy
type jsiiProxy_CfnRegistryPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRegistryPolicy) AttrRegistryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRegistryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistryPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistryPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistryPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistryPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistryPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistryPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistryPolicy) PolicyText() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistryPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistryPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistryPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECR::RegistryPolicy`.
func NewCfnRegistryPolicy(scope constructs.Construct, id *string, props *CfnRegistryPolicyProps) CfnRegistryPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnRegistryPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.CfnRegistryPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECR::RegistryPolicy`.
func NewCfnRegistryPolicy_Override(c CfnRegistryPolicy, scope constructs.Construct, id *string, props *CfnRegistryPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.CfnRegistryPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRegistryPolicy) SetPolicyText(val interface{}) {
	_jsii_.Set(
		j,
		"policyText",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRegistryPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnRegistryPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRegistryPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnRegistryPolicy",
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
func CfnRegistryPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnRegistryPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRegistryPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr.CfnRegistryPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRegistryPolicy) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnRegistryPolicy) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnRegistryPolicy) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnRegistryPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRegistryPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRegistryPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnRegistryPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnRegistryPolicy) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnRegistryPolicy) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnRegistryPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRegistryPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRegistryPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnRegistryPolicy) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnRegistryPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRegistryPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnRegistryPolicy`.
//
// TODO: EXAMPLE
//
type CfnRegistryPolicyProps struct {
	// The JSON policy text for your registry.
	PolicyText interface{} `json:"policyText" yaml:"policyText"`
}

// A CloudFormation `AWS::ECR::ReplicationConfiguration`.
//
// The `AWS::ECR::ReplicationConfiguration` resource creates or updates the replication configuration for a private registry. The first time a replication configuration is applied to a private registry, a service-linked IAM role is created in your account for the replication process. For more information, see [Using Service-Linked Roles for Amazon ECR](https://docs.aws.amazon.com/AmazonECR/latest/userguide/using-service-linked-roles.html) in the *Amazon Elastic Container Registry User Guide* .
//
// > When configuring cross-account replication, the destination account must grant the source account permission to replicate. This permission is controlled using a private registry permissions policy. For more information, see `AWS::ECR::RegistryPolicy` .
//
// TODO: EXAMPLE
//
type CfnReplicationConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrRegistryId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ReplicationConfiguration() interface{}
	SetReplicationConfiguration(val interface{})
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

// The jsii proxy struct for CfnReplicationConfiguration
type jsiiProxy_CfnReplicationConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnReplicationConfiguration) AttrRegistryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRegistryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationConfiguration) ReplicationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECR::ReplicationConfiguration`.
func NewCfnReplicationConfiguration(scope constructs.Construct, id *string, props *CfnReplicationConfigurationProps) CfnReplicationConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnReplicationConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.CfnReplicationConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECR::ReplicationConfiguration`.
func NewCfnReplicationConfiguration_Override(c CfnReplicationConfiguration, scope constructs.Construct, id *string, props *CfnReplicationConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.CfnReplicationConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReplicationConfiguration) SetReplicationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"replicationConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnReplicationConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnReplicationConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnReplicationConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnReplicationConfiguration",
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
func CfnReplicationConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnReplicationConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicationConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr.CfnReplicationConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnReplicationConfiguration) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnReplicationConfiguration) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnReplicationConfiguration) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnReplicationConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnReplicationConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnReplicationConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnReplicationConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnReplicationConfiguration) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnReplicationConfiguration) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnReplicationConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnReplicationConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReplicationConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnReplicationConfiguration) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnReplicationConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The replication configuration for a registry.
//
// TODO: EXAMPLE
//
type CfnReplicationConfiguration_ReplicationConfigurationProperty struct {
	// An array of objects representing the replication destinations and repository filters for a replication configuration.
	Rules interface{} `json:"rules" yaml:"rules"`
}

// An array of objects representing the destination for a replication rule.
//
// TODO: EXAMPLE
//
type CfnReplicationConfiguration_ReplicationDestinationProperty struct {
	// The Region to replicate to.
	Region *string `json:"region" yaml:"region"`
	// The AWS account ID of the Amazon ECR private registry to replicate to.
	//
	// When configuring cross-Region replication within your own registry, specify your own account ID.
	RegistryId *string `json:"registryId" yaml:"registryId"`
}

// An array of objects representing the replication destinations and repository filters for a replication configuration.
//
// TODO: EXAMPLE
//
type CfnReplicationConfiguration_ReplicationRuleProperty struct {
	// An array of objects representing the destination for a replication rule.
	Destinations interface{} `json:"destinations" yaml:"destinations"`
	// An array of objects representing the filters for a replication rule.
	//
	// Specifying a repository filter for a replication rule provides a method for controlling which repositories in a private registry are replicated.
	RepositoryFilters interface{} `json:"repositoryFilters" yaml:"repositoryFilters"`
}

// The filter settings used with image replication.
//
// Specifying a repository filter to a replication rule provides a method for controlling which repositories in a private registry are replicated. If no repository filter is specified, all images in the repository are replicated.
//
// TODO: EXAMPLE
//
type CfnReplicationConfiguration_RepositoryFilterProperty struct {
	// The repository filter details.
	//
	// When the `PREFIX_MATCH` filter type is specified, this value is required and should be the repository name prefix to configure replication for.
	Filter *string `json:"filter" yaml:"filter"`
	// The repository filter type.
	//
	// The only supported value is `PREFIX_MATCH` , which is a repository name prefix specified with the `filter` parameter.
	FilterType *string `json:"filterType" yaml:"filterType"`
}

// Properties for defining a `CfnReplicationConfiguration`.
//
// TODO: EXAMPLE
//
type CfnReplicationConfigurationProps struct {
	// The replication configuration for a registry.
	ReplicationConfiguration interface{} `json:"replicationConfiguration" yaml:"replicationConfiguration"`
}

// A CloudFormation `AWS::ECR::Repository`.
//
// The `AWS::ECR::Repository` resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see [Amazon ECR private repositories](https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html) in the *Amazon ECR User Guide* .
//
// TODO: EXAMPLE
//
type CfnRepository interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrRepositoryUri() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EncryptionConfiguration() interface{}
	SetEncryptionConfiguration(val interface{})
	ImageScanningConfiguration() interface{}
	SetImageScanningConfiguration(val interface{})
	ImageTagMutability() *string
	SetImageTagMutability(val *string)
	LifecyclePolicy() interface{}
	SetLifecyclePolicy(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RepositoryName() *string
	SetRepositoryName(val *string)
	RepositoryPolicyText() interface{}
	SetRepositoryPolicyText(val interface{})
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

// The jsii proxy struct for CfnRepository
type jsiiProxy_CfnRepository struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRepository) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) AttrRepositoryUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRepositoryUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) EncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) ImageScanningConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageScanningConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) ImageTagMutability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTagMutability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) LifecyclePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecyclePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) RepositoryPolicyText() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryPolicyText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECR::Repository`.
func NewCfnRepository(scope constructs.Construct, id *string, props *CfnRepositoryProps) CfnRepository {
	_init_.Initialize()

	j := jsiiProxy_CfnRepository{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.CfnRepository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECR::Repository`.
func NewCfnRepository_Override(c CfnRepository, scope constructs.Construct, id *string, props *CfnRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.CfnRepository",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRepository) SetEncryptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnRepository) SetImageScanningConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"imageScanningConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnRepository) SetImageTagMutability(val *string) {
	_jsii_.Set(
		j,
		"imageTagMutability",
		val,
	)
}

func (j *jsiiProxy_CfnRepository) SetLifecyclePolicy(val interface{}) {
	_jsii_.Set(
		j,
		"lifecyclePolicy",
		val,
	)
}

func (j *jsiiProxy_CfnRepository) SetRepositoryName(val *string) {
	_jsii_.Set(
		j,
		"repositoryName",
		val,
	)
}

func (j *jsiiProxy_CfnRepository) SetRepositoryPolicyText(val interface{}) {
	_jsii_.Set(
		j,
		"repositoryPolicyText",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRepository_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnRepository",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRepository_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnRepository",
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
func CfnRepository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.CfnRepository",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRepository_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr.CfnRepository",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRepository) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnRepository) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnRepository) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnRepository) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRepository) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRepository) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnRepository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnRepository) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnRepository) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnRepository) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRepository) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRepository) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnRepository) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnRepository) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRepository) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.
//
// By default, when no encryption configuration is set or the `AES256` encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part.
//
// For more control over the encryption of the contents of your repository, you can use server-side encryption with AWS Key Management Service key stored in AWS Key Management Service ( AWS KMS ) to encrypt your images. For more information, see [Amazon ECR encryption at rest](https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html) in the *Amazon Elastic Container Registry User Guide* .
//
// TODO: EXAMPLE
//
type CfnRepository_EncryptionConfigurationProperty struct {
	// The encryption type to use.
	//
	// If you use the `KMS` encryption type, the contents of the repository will be encrypted using server-side encryption with AWS Key Management Service key stored in AWS KMS . When you use AWS KMS to encrypt your data, you can either use the default AWS managed AWS KMS key for Amazon ECR, or specify your own AWS KMS key, which you already created. For more information, see [Protecting data using server-side encryption with an AWS KMS key stored in AWS Key Management Service (SSE-KMS)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide* .
	//
	// If you use the `AES256` encryption type, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts the images in the repository using an AES-256 encryption algorithm. For more information, see [Protecting data using server-side encryption with Amazon S3-managed encryption keys (SSE-S3)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide* .
	EncryptionType *string `json:"encryptionType" yaml:"encryptionType"`
	// If you use the `KMS` encryption type, specify the AWS KMS key to use for encryption.
	//
	// The alias, key ID, or full ARN of the AWS KMS key can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed AWS KMS key for Amazon ECR will be used.
	KmsKey *string `json:"kmsKey" yaml:"kmsKey"`
}

// The image scanning configuration for a repository.
//
// TODO: EXAMPLE
//
type CfnRepository_ImageScanningConfigurationProperty struct {
	// The setting that determines whether images are scanned after being pushed to a repository.
	//
	// If set to `true` , images will be scanned after being pushed. If this parameter is not specified, it will default to `false` and images will not be scanned unless a scan is manually started.
	ScanOnPush interface{} `json:"scanOnPush" yaml:"scanOnPush"`
}

// The `LifecyclePolicy` property type specifies a lifecycle policy.
//
// For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html) in the *Amazon ECR User Guide* .
//
// TODO: EXAMPLE
//
type CfnRepository_LifecyclePolicyProperty struct {
	// The JSON repository policy text to apply to the repository.
	LifecyclePolicyText *string `json:"lifecyclePolicyText" yaml:"lifecyclePolicyText"`
	// The AWS account ID associated with the registry that contains the repository.
	//
	// If you do not specify a registry, the default registry is assumed.
	RegistryId *string `json:"registryId" yaml:"registryId"`
}

// Properties for defining a `CfnRepository`.
//
// TODO: EXAMPLE
//
type CfnRepositoryProps struct {
	// The encryption configuration for the repository.
	//
	// This determines how the contents of your repository are encrypted at rest.
	EncryptionConfiguration interface{} `json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The image scanning configuration for the repository.
	//
	// This determines whether images are scanned for known vulnerabilities after being pushed to the repository.
	ImageScanningConfiguration interface{} `json:"imageScanningConfiguration" yaml:"imageScanningConfiguration"`
	// The tag mutability setting for the repository.
	//
	// If this parameter is omitted, the default setting of `MUTABLE` will be used which will allow image tags to be overwritten. If `IMMUTABLE` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.
	ImageTagMutability *string `json:"imageTagMutability" yaml:"imageTagMutability"`
	// Creates or updates a lifecycle policy.
	//
	// For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html) .
	LifecyclePolicy interface{} `json:"lifecyclePolicy" yaml:"lifecyclePolicy"`
	// The name to use for the repository.
	//
	// The repository name may be specified on its own (such as `nginx-web-app` ) or it can be prepended with a namespace to group the repository into a category (such as `project-a/nginx-web-app` ). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
	// The JSON repository policy text to apply to the repository.
	//
	// For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide* .
	RepositoryPolicyText interface{} `json:"repositoryPolicyText" yaml:"repositoryPolicyText"`
	// An array of key-value pairs to apply to this resource.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// Represents an ECR repository.
type IRepository interface {
	awscdk.IResource
	// Add a policy statement to the repository's resource policy.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Grant the given principal identity permissions to perform the actions on this repository.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to pull images in this repository.
	GrantPull(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to pull and push images to this repository.
	GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant
	// Define a CloudWatch event that triggers when something happens to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event rule that can trigger a target when an image is pushed to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers for repository events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event rule that can trigger a target when the image scan is completed.
	OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule
	// Returns the URI of the repository for a certain tag. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[@DIGEST]
	RepositoryUriForDigest(digest *string) *string
	// Returns the URI of the repository for a certain tag. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[:TAG]
	RepositoryUriForTag(tag *string) *string
	// The ARN of the repository.
	RepositoryArn() *string
	// The name of the repository.
	RepositoryName() *string
	// The URI of this repository (represents the latest image):.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY
	RepositoryUri() *string
}

// The jsii proxy for IRepository
type jsiiProxy_IRepository struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IRepository) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) GrantPull(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPull",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPullPush",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCloudTrailEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCloudTrailImagePushed",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onImageScanCompleted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) RepositoryUriForDigest(digest *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"repositoryUriForDigest",
		[]interface{}{digest},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) RepositoryUriForTag(tag *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"repositoryUriForTag",
		[]interface{}{tag},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IRepository) RepositoryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) RepositoryUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUri",
		&returns,
	)
	return returns
}

// An ECR life cycle rule.
//
// TODO: EXAMPLE
//
type LifecycleRule struct {
	// Describes the purpose of the rule.
	Description *string `json:"description" yaml:"description"`
	// The maximum age of images to retain. The value must represent a number of days.
	//
	// Specify exactly one of maxImageCount and maxImageAge.
	MaxImageAge awscdk.Duration `json:"maxImageAge" yaml:"maxImageAge"`
	// The maximum number of images to retain.
	//
	// Specify exactly one of maxImageCount and maxImageAge.
	MaxImageCount *float64 `json:"maxImageCount" yaml:"maxImageCount"`
	// Controls the order in which rules are evaluated (low to high).
	//
	// All rules must have a unique priority, where lower numbers have
	// higher precedence. The first rule that matches is applied to an image.
	//
	// There can only be one rule with a tagStatus of Any, and it must have
	// the highest rulePriority.
	//
	// All rules without a specified priority will have incrementing priorities
	// automatically assigned to them, higher than any rules that DO have priorities.
	RulePriority *float64 `json:"rulePriority" yaml:"rulePriority"`
	// Select images that have ALL the given prefixes in their tag.
	//
	// Only if tagStatus == TagStatus.Tagged
	TagPrefixList *[]*string `json:"tagPrefixList" yaml:"tagPrefixList"`
	// Select images based on tags.
	//
	// Only one rule is allowed to select untagged images, and it must
	// have the highest rulePriority.
	TagStatus TagStatus `json:"tagStatus" yaml:"tagStatus"`
}

// Options for the onCloudTrailImagePushed method.
//
// TODO: EXAMPLE
//
type OnCloudTrailImagePushedOptions struct {
	// A description of the rule's purpose.
	Description *string `json:"description" yaml:"description"`
	// Additional restrictions for the event to route to the specified target.
	//
	// The method that generates the rule probably imposes some type of event
	// filtering. The filtering implied by what you pass here is added
	// on top of that filtering.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	EventPattern *awsevents.EventPattern `json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	RuleName *string `json:"ruleName" yaml:"ruleName"`
	// The target to register for the event.
	Target awsevents.IRuleTarget `json:"target" yaml:"target"`
	// Only watch changes to this image tag.
	ImageTag *string `json:"imageTag" yaml:"imageTag"`
}

// Options for the OnImageScanCompleted method.
//
// TODO: EXAMPLE
//
type OnImageScanCompletedOptions struct {
	// A description of the rule's purpose.
	Description *string `json:"description" yaml:"description"`
	// Additional restrictions for the event to route to the specified target.
	//
	// The method that generates the rule probably imposes some type of event
	// filtering. The filtering implied by what you pass here is added
	// on top of that filtering.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	EventPattern *awsevents.EventPattern `json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	RuleName *string `json:"ruleName" yaml:"ruleName"`
	// The target to register for the event.
	Target awsevents.IRuleTarget `json:"target" yaml:"target"`
	// Only watch changes to the image tags spedified.
	//
	// Leave it undefined to watch the full repository.
	ImageTags *[]*string `json:"imageTags" yaml:"imageTags"`
}

// Authorization token to access the global public ECR Gallery via Docker CLI.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AmazonECR/latest/public/public-registries.html#public-registry-auth
//
type PublicGalleryAuthorizationToken interface {
}

// The jsii proxy struct for PublicGalleryAuthorizationToken
type jsiiProxy_PublicGalleryAuthorizationToken struct {
	_ byte // padding
}

// Grant access to retrieve an authorization token.
func PublicGalleryAuthorizationToken_GrantRead(grantee awsiam.IGrantable) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_ecr.PublicGalleryAuthorizationToken",
		"grantRead",
		[]interface{}{grantee},
	)
}

// Define an ECR repository.
//
// TODO: EXAMPLE
//
type Repository interface {
	RepositoryBase
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	RepositoryArn() *string
	RepositoryName() *string
	RepositoryUri() *string
	Stack() awscdk.Stack
	AddLifecycleRule(rule *LifecycleRule)
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	GrantPull(grantee awsiam.IGrantable) awsiam.Grant
	GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant
	OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule
	RepositoryUriForDigest(digest *string) *string
	RepositoryUriForTag(tag *string) *string
	ToString() *string
}

// The jsii proxy struct for Repository
type jsiiProxy_Repository struct {
	jsiiProxy_RepositoryBase
}

func (j *jsiiProxy_Repository) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) RepositoryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) RepositoryUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewRepository(scope constructs.Construct, id *string, props *RepositoryProps) Repository {
	_init_.Initialize()

	j := jsiiProxy_Repository{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.Repository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRepository_Override(r Repository, scope constructs.Construct, id *string, props *RepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.Repository",
		[]interface{}{scope, id, props},
		r,
	)
}

// Returns an ECR ARN for a repository that resides in the same account/region as the current stack.
func Repository_ArnForLocalRepository(repositoryName *string, scope constructs.IConstruct, account *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.Repository",
		"arnForLocalRepository",
		[]interface{}{repositoryName, scope, account},
		&returns,
	)

	return returns
}

func Repository_FromRepositoryArn(scope constructs.Construct, id *string, repositoryArn *string) IRepository {
	_init_.Initialize()

	var returns IRepository

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.Repository",
		"fromRepositoryArn",
		[]interface{}{scope, id, repositoryArn},
		&returns,
	)

	return returns
}

// Import a repository.
func Repository_FromRepositoryAttributes(scope constructs.Construct, id *string, attrs *RepositoryAttributes) IRepository {
	_init_.Initialize()

	var returns IRepository

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.Repository",
		"fromRepositoryAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

func Repository_FromRepositoryName(scope constructs.Construct, id *string, repositoryName *string) IRepository {
	_init_.Initialize()

	var returns IRepository

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.Repository",
		"fromRepositoryName",
		[]interface{}{scope, id, repositoryName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Repository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.Repository",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Repository_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.Repository",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a life cycle rule to the repository.
//
// Life cycle rules automatically expire images from the repository that match
// certain conditions.
func (r *jsiiProxy_Repository) AddLifecycleRule(rule *LifecycleRule) {
	_jsii_.InvokeVoid(
		r,
		"addLifecycleRule",
		[]interface{}{rule},
	)
}

// Add a policy statement to the repository's resource policy.
func (r *jsiiProxy_Repository) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		r,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (r *jsiiProxy_Repository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_Repository) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (r *jsiiProxy_Repository) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (r *jsiiProxy_Repository) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given principal identity permissions to perform the actions on this repository.
func (r *jsiiProxy_Repository) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grant",
		args,
		&returns,
	)

	return returns
}

// Grant the given identity permissions to use the images in this repository.
func (r *jsiiProxy_Repository) GrantPull(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grantPull",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to pull and push images to this repository.
func (r *jsiiProxy_Repository) GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grantPullPush",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Define a CloudWatch event that triggers when something happens to this repository.
//
// Requires that there exists at least one CloudTrail Trail in your account
// that captures the event. This method will not create the Trail.
func (r *jsiiProxy_Repository) OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onCloudTrailEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an AWS CloudWatch event rule that can trigger a target when an image is pushed to this repository.
//
// Requires that there exists at least one CloudTrail Trail in your account
// that captures the event. This method will not create the Trail.
func (r *jsiiProxy_Repository) OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onCloudTrailImagePushed",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a CloudWatch event rule which triggers for repository events.
//
// Use
// `rule.addEventPattern(pattern)` to specify a filter.
func (r *jsiiProxy_Repository) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an AWS CloudWatch event rule that can trigger a target when an image scan is completed.
func (r *jsiiProxy_Repository) OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onImageScanCompleted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Returns the URL of the repository. Can be used in `docker push/pull`.
//
// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[@DIGEST]
func (r *jsiiProxy_Repository) RepositoryUriForDigest(digest *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"repositoryUriForDigest",
		[]interface{}{digest},
		&returns,
	)

	return returns
}

// Returns the URL of the repository. Can be used in `docker push/pull`.
//
// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[:TAG]
func (r *jsiiProxy_Repository) RepositoryUriForTag(tag *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"repositoryUriForTag",
		[]interface{}{tag},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_Repository) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type RepositoryAttributes struct {
	RepositoryArn *string `json:"repositoryArn" yaml:"repositoryArn"`
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
}

// Base class for ECR repository.
//
// Reused between imported repositories and owned repositories.
type RepositoryBase interface {
	awscdk.Resource
	IRepository
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	RepositoryArn() *string
	RepositoryName() *string
	RepositoryUri() *string
	Stack() awscdk.Stack
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	GrantPull(grantee awsiam.IGrantable) awsiam.Grant
	GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant
	OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule
	RepositoryUriForDigest(digest *string) *string
	RepositoryUriForTag(tag *string) *string
	ToString() *string
}

// The jsii proxy struct for RepositoryBase
type jsiiProxy_RepositoryBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IRepository
}

func (j *jsiiProxy_RepositoryBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryBase) RepositoryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryBase) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryBase) RepositoryUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewRepositoryBase_Override(r RepositoryBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.RepositoryBase",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func RepositoryBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.RepositoryBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func RepositoryBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.RepositoryBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a policy statement to the repository's resource policy.
func (r *jsiiProxy_RepositoryBase) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		r,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (r *jsiiProxy_RepositoryBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_RepositoryBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (r *jsiiProxy_RepositoryBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (r *jsiiProxy_RepositoryBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given principal identity permissions to perform the actions on this repository.
func (r *jsiiProxy_RepositoryBase) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grant",
		args,
		&returns,
	)

	return returns
}

// Grant the given identity permissions to use the images in this repository.
func (r *jsiiProxy_RepositoryBase) GrantPull(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grantPull",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to pull and push images to this repository.
func (r *jsiiProxy_RepositoryBase) GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grantPullPush",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Define a CloudWatch event that triggers when something happens to this repository.
//
// Requires that there exists at least one CloudTrail Trail in your account
// that captures the event. This method will not create the Trail.
func (r *jsiiProxy_RepositoryBase) OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onCloudTrailEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an AWS CloudWatch event rule that can trigger a target when an image is pushed to this repository.
//
// Requires that there exists at least one CloudTrail Trail in your account
// that captures the event. This method will not create the Trail.
func (r *jsiiProxy_RepositoryBase) OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onCloudTrailImagePushed",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a CloudWatch event rule which triggers for repository events.
//
// Use
// `rule.addEventPattern(pattern)` to specify a filter.
func (r *jsiiProxy_RepositoryBase) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an AWS CloudWatch event rule that can trigger a target when an image scan is completed.
func (r *jsiiProxy_RepositoryBase) OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onImageScanCompleted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Returns the URL of the repository. Can be used in `docker push/pull`.
//
// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[@DIGEST]
func (r *jsiiProxy_RepositoryBase) RepositoryUriForDigest(digest *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"repositoryUriForDigest",
		[]interface{}{digest},
		&returns,
	)

	return returns
}

// Returns the URL of the repository. Can be used in `docker push/pull`.
//
// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[:TAG]
func (r *jsiiProxy_RepositoryBase) RepositoryUriForTag(tag *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"repositoryUriForTag",
		[]interface{}{tag},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_RepositoryBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Indicates whether server-side encryption is enabled for the object, and whether that encryption is from the AWS Key Management Service (AWS KMS) or from Amazon S3 managed encryption (SSE-S3).
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
type RepositoryEncryption interface {
	Value() *string
}

// The jsii proxy struct for RepositoryEncryption
type jsiiProxy_RepositoryEncryption struct {
	_ byte // padding
}

func (j *jsiiProxy_RepositoryEncryption) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewRepositoryEncryption(value *string) RepositoryEncryption {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEncryption{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.RepositoryEncryption",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewRepositoryEncryption_Override(r RepositoryEncryption, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.RepositoryEncryption",
		[]interface{}{value},
		r,
	)
}

func RepositoryEncryption_AES_256() RepositoryEncryption {
	_init_.Initialize()
	var returns RepositoryEncryption
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr.RepositoryEncryption",
		"AES_256",
		&returns,
	)
	return returns
}

func RepositoryEncryption_KMS() RepositoryEncryption {
	_init_.Initialize()
	var returns RepositoryEncryption
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr.RepositoryEncryption",
		"KMS",
		&returns,
	)
	return returns
}

// TODO: EXAMPLE
//
type RepositoryProps struct {
	// The kind of server-side encryption to apply to this repository.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryptionKey is not specified, an AWS managed KMS key is used.
	Encryption RepositoryEncryption `json:"encryption" yaml:"encryption"`
	// External KMS key to use for repository encryption.
	//
	// The 'encryption' property must be either not specified or set to "KMS".
	// An error will be emitted if encryption is set to "AES256".
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// Enable the scan on push when creating the repository.
	ImageScanOnPush *bool `json:"imageScanOnPush" yaml:"imageScanOnPush"`
	// The tag mutability setting for the repository.
	//
	// If this parameter is omitted, the default setting of MUTABLE will be used which will allow image tags to be overwritten.
	ImageTagMutability TagMutability `json:"imageTagMutability" yaml:"imageTagMutability"`
	// The AWS account ID associated with the registry that contains the repository.
	// See: https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_PutLifecyclePolicy.html
	//
	LifecycleRegistryId *string `json:"lifecycleRegistryId" yaml:"lifecycleRegistryId"`
	// Life cycle rules to apply to this registry.
	LifecycleRules *[]*LifecycleRule `json:"lifecycleRules" yaml:"lifecycleRules"`
	// Determine what happens to the repository when the resource/stack is deleted.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// Name for this repository.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
}

// The tag mutability setting for your repository.
//
// TODO: EXAMPLE
//
type TagMutability string

const (
	TagMutability_MUTABLE TagMutability = "MUTABLE"
	TagMutability_IMMUTABLE TagMutability = "IMMUTABLE"
)

// Select images based on tags.
type TagStatus string

const (
	TagStatus_ANY TagStatus = "ANY"
	TagStatus_TAGGED TagStatus = "TAGGED"
	TagStatus_UNTAGGED TagStatus = "UNTAGGED"
)

