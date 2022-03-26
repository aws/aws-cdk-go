package awsecr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecr/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/constructs-go/constructs/v3"
)

// Authorization token to access private ECR repositories in the current environment via Docker CLI.
//
// Example:
//   user := iam.NewUser(this, jsii.String("User"))
//   ecr.authorizationToken.grantRead(user)
//
// See: https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry_auth.html
//
// Experimental.
type AuthorizationToken interface {
}

// The jsii proxy struct for AuthorizationToken
type jsiiProxy_AuthorizationToken struct {
	_ byte // padding
}

// Grant access to retrieve an authorization token.
// Experimental.
func AuthorizationToken_GrantRead(grantee awsiam.IGrantable) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_ecr.AuthorizationToken",
		"grantRead",
		[]interface{}{grantee},
	)
}

// A CloudFormation `AWS::ECR::PublicRepository`.
//
// The `AWS::ECR::PublicRepository` resource specifies an Amazon Elastic Container Registry Public (Amazon ECR Public) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see [Amazon ECR public repositories](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repositories.html) in the *Amazon ECR Public User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//
//   var repositoryCatalogData interface{}
//   var repositoryPolicyText interface{}
//   cfnPublicRepository := ecr.NewCfnPublicRepository(this, jsii.String("MyCfnPublicRepository"), &cfnPublicRepositoryProps{
//   	repositoryCatalogData: repositoryCatalogData,
//   	repositoryName: jsii.String("repositoryName"),
//   	repositoryPolicyText: repositoryPolicyText,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPublicRepository interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the Amazon Resource Name (ARN) for the specified `AWS::ECR::PublicRepository` resource.
	//
	// For example, `arn:aws:ecr-public:: *123456789012* :repository/ *test-repository*` .
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
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
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The details about the repository that are publicly visible in the Amazon ECR Public Gallery.
	//
	// For more information, see [Amazon ECR Public repository catalog data](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repository-catalog-data.html) in the *Amazon ECR Public User Guide* .
	RepositoryCatalogData() interface{}
	SetRepositoryCatalogData(val interface{})
	// The name to use for the public repository.
	//
	// The repository name may be specified on its own (such as `nginx-web-app` ) or it can be prepended with a namespace to group the repository into a category (such as `project-a/nginx-web-app` ). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	RepositoryName() *string
	SetRepositoryName(val *string)
	// The JSON repository policy text to apply to the public repository.
	//
	// For more information, see [Amazon ECR Public repository policies](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repository-policies.html) in the *Amazon ECR Public User Guide* .
	RepositoryPolicyText() interface{}
	SetRepositoryPolicyText(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
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

func (j *jsiiProxy_CfnPublicRepository) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnPublicRepository(scope awscdk.Construct, id *string, props *CfnPublicRepositoryProps) CfnPublicRepository {
	_init_.Initialize()

	j := jsiiProxy_CfnPublicRepository{}

	_jsii_.Create(
		"monocdk.aws_ecr.CfnPublicRepository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECR::PublicRepository`.
func NewCfnPublicRepository_Override(c CfnPublicRepository, scope awscdk.Construct, id *string, props *CfnPublicRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecr.CfnPublicRepository",
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
// Experimental.
func CfnPublicRepository_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnPublicRepository",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPublicRepository_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnPublicRepository",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPublicRepository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnPublicRepository",
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
		"monocdk.aws_ecr.CfnPublicRepository",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPublicRepository) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPublicRepository) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPublicRepository) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPublicRepository) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPublicRepository) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPublicRepository) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPublicRepository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnPublicRepository) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPublicRepository) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPublicRepository) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPublicRepository) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicRepository) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPublicRepository) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnPublicRepository) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnPublicRepository) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//
//   var repositoryCatalogData interface{}
//   var repositoryPolicyText interface{}
//   cfnPublicRepositoryProps := &cfnPublicRepositoryProps{
//   	repositoryCatalogData: repositoryCatalogData,
//   	repositoryName: jsii.String("repositoryName"),
//   	repositoryPolicyText: repositoryPolicyText,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//   cfnPullThroughCacheRule := ecr.NewCfnPullThroughCacheRule(this, jsii.String("MyCfnPullThroughCacheRule"), &cfnPullThroughCacheRuleProps{
//   	ecrRepositoryPrefix: jsii.String("ecrRepositoryPrefix"),
//   	upstreamRegistryUrl: jsii.String("upstreamRegistryUrl"),
//   })
//
type CfnPullThroughCacheRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The Amazon ECR repository prefix associated with the pull through cache rule.
	EcrRepositoryPrefix() *string
	SetEcrRepositoryPrefix(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The upstream registry URL associated with the pull through cache rule.
	UpstreamRegistryUrl() *string
	SetUpstreamRegistryUrl(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
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

func (j *jsiiProxy_CfnPullThroughCacheRule) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnPullThroughCacheRule(scope awscdk.Construct, id *string, props *CfnPullThroughCacheRuleProps) CfnPullThroughCacheRule {
	_init_.Initialize()

	j := jsiiProxy_CfnPullThroughCacheRule{}

	_jsii_.Create(
		"monocdk.aws_ecr.CfnPullThroughCacheRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECR::PullThroughCacheRule`.
func NewCfnPullThroughCacheRule_Override(c CfnPullThroughCacheRule, scope awscdk.Construct, id *string, props *CfnPullThroughCacheRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecr.CfnPullThroughCacheRule",
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
// Experimental.
func CfnPullThroughCacheRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnPullThroughCacheRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPullThroughCacheRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnPullThroughCacheRule",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPullThroughCacheRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnPullThroughCacheRule",
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
		"monocdk.aws_ecr.CfnPullThroughCacheRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPullThroughCacheRule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPullThroughCacheRule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPullThroughCacheRule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPullThroughCacheRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPullThroughCacheRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPullThroughCacheRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPullThroughCacheRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnPullThroughCacheRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPullThroughCacheRule) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPullThroughCacheRule) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPullThroughCacheRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPullThroughCacheRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPullThroughCacheRule) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnPullThroughCacheRule) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnPullThroughCacheRule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//   cfnPullThroughCacheRuleProps := &cfnPullThroughCacheRuleProps{
//   	ecrRepositoryPrefix: jsii.String("ecrRepositoryPrefix"),
//   	upstreamRegistryUrl: jsii.String("upstreamRegistryUrl"),
//   }
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//
//   var policyText interface{}
//   cfnRegistryPolicy := ecr.NewCfnRegistryPolicy(this, jsii.String("MyCfnRegistryPolicy"), &cfnRegistryPolicyProps{
//   	policyText: policyText,
//   })
//
type CfnRegistryPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The account ID of the private registry the policy is associated with.
	AttrRegistryId() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
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
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The JSON policy text for your registry.
	PolicyText() interface{}
	SetPolicyText(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
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

func (j *jsiiProxy_CfnRegistryPolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnRegistryPolicy(scope awscdk.Construct, id *string, props *CfnRegistryPolicyProps) CfnRegistryPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnRegistryPolicy{}

	_jsii_.Create(
		"monocdk.aws_ecr.CfnRegistryPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECR::RegistryPolicy`.
func NewCfnRegistryPolicy_Override(c CfnRegistryPolicy, scope awscdk.Construct, id *string, props *CfnRegistryPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecr.CfnRegistryPolicy",
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
// Experimental.
func CfnRegistryPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnRegistryPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnRegistryPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnRegistryPolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnRegistryPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnRegistryPolicy",
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
		"monocdk.aws_ecr.CfnRegistryPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRegistryPolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRegistryPolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRegistryPolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRegistryPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRegistryPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRegistryPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRegistryPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnRegistryPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRegistryPolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnRegistryPolicy) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnRegistryPolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRegistryPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRegistryPolicy) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnRegistryPolicy) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnRegistryPolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//
//   var policyText interface{}
//   cfnRegistryPolicyProps := &cfnRegistryPolicyProps{
//   	policyText: policyText,
//   }
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//   cfnReplicationConfiguration := ecr.NewCfnReplicationConfiguration(this, jsii.String("MyCfnReplicationConfiguration"), &cfnReplicationConfigurationProps{
//   	replicationConfiguration: &replicationConfigurationProperty{
//   		rules: []interface{}{
//   			&replicationRuleProperty{
//   				destinations: []interface{}{
//   					&replicationDestinationProperty{
//   						region: jsii.String("region"),
//   						registryId: jsii.String("registryId"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				repositoryFilters: []interface{}{
//   					&repositoryFilterProperty{
//   						filter: jsii.String("filter"),
//   						filterType: jsii.String("filterType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   })
//
type CfnReplicationConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The account ID of the destination registry.
	AttrRegistryId() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
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
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The replication configuration for a registry.
	ReplicationConfiguration() interface{}
	SetReplicationConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
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

func (j *jsiiProxy_CfnReplicationConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnReplicationConfiguration(scope awscdk.Construct, id *string, props *CfnReplicationConfigurationProps) CfnReplicationConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnReplicationConfiguration{}

	_jsii_.Create(
		"monocdk.aws_ecr.CfnReplicationConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECR::ReplicationConfiguration`.
func NewCfnReplicationConfiguration_Override(c CfnReplicationConfiguration, scope awscdk.Construct, id *string, props *CfnReplicationConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecr.CfnReplicationConfiguration",
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
// Experimental.
func CfnReplicationConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnReplicationConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnReplicationConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnReplicationConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnReplicationConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnReplicationConfiguration",
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
		"monocdk.aws_ecr.CfnReplicationConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReplicationConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnReplicationConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnReplicationConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnReplicationConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnReplicationConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnReplicationConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnReplicationConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnReplicationConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnReplicationConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnReplicationConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnReplicationConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReplicationConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnReplicationConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnReplicationConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//   replicationConfigurationProperty := &replicationConfigurationProperty{
//   	rules: []interface{}{
//   		&replicationRuleProperty{
//   			destinations: []interface{}{
//   				&replicationDestinationProperty{
//   					region: jsii.String("region"),
//   					registryId: jsii.String("registryId"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			repositoryFilters: []interface{}{
//   				&repositoryFilterProperty{
//   					filter: jsii.String("filter"),
//   					filterType: jsii.String("filterType"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnReplicationConfiguration_ReplicationConfigurationProperty struct {
	// An array of objects representing the replication destinations and repository filters for a replication configuration.
	Rules interface{} `json:"rules" yaml:"rules"`
}

// An array of objects representing the destination for a replication rule.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//   replicationDestinationProperty := &replicationDestinationProperty{
//   	region: jsii.String("region"),
//   	registryId: jsii.String("registryId"),
//   }
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//   replicationRuleProperty := &replicationRuleProperty{
//   	destinations: []interface{}{
//   		&replicationDestinationProperty{
//   			region: jsii.String("region"),
//   			registryId: jsii.String("registryId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	repositoryFilters: []interface{}{
//   		&repositoryFilterProperty{
//   			filter: jsii.String("filter"),
//   			filterType: jsii.String("filterType"),
//   		},
//   	},
//   }
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//   repositoryFilterProperty := &repositoryFilterProperty{
//   	filter: jsii.String("filter"),
//   	filterType: jsii.String("filterType"),
//   }
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//   cfnReplicationConfigurationProps := &cfnReplicationConfigurationProps{
//   	replicationConfiguration: &replicationConfigurationProperty{
//   		rules: []interface{}{
//   			&replicationRuleProperty{
//   				destinations: []interface{}{
//   					&replicationDestinationProperty{
//   						region: jsii.String("region"),
//   						registryId: jsii.String("registryId"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				repositoryFilters: []interface{}{
//   					&repositoryFilterProperty{
//   						filter: jsii.String("filter"),
//   						filterType: jsii.String("filterType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnReplicationConfigurationProps struct {
	// The replication configuration for a registry.
	ReplicationConfiguration interface{} `json:"replicationConfiguration" yaml:"replicationConfiguration"`
}

// A CloudFormation `AWS::ECR::Repository`.
//
// The `AWS::ECR::Repository` resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see [Amazon ECR private repositories](https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html) in the *Amazon ECR User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//
//   var repositoryPolicyText interface{}
//   cfnRepository := ecr.NewCfnRepository(this, jsii.String("MyCfnRepository"), &cfnRepositoryProps{
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		encryptionType: jsii.String("encryptionType"),
//
//   		// the properties below are optional
//   		kmsKey: jsii.String("kmsKey"),
//   	},
//   	imageScanningConfiguration: &imageScanningConfigurationProperty{
//   		scanOnPush: jsii.Boolean(false),
//   	},
//   	imageTagMutability: jsii.String("imageTagMutability"),
//   	lifecyclePolicy: &lifecyclePolicyProperty{
//   		lifecyclePolicyText: jsii.String("lifecyclePolicyText"),
//   		registryId: jsii.String("registryId"),
//   	},
//   	repositoryName: jsii.String("repositoryName"),
//   	repositoryPolicyText: repositoryPolicyText,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnRepository interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the Amazon Resource Name (ARN) for the specified `AWS::ECR::Repository` resource.
	//
	// For example, `arn:aws:ecr: *eu-west-1* : *123456789012* :repository/ *test-repository*` .
	AttrArn() *string
	// Returns the URI for the specified `AWS::ECR::Repository` resource.
	//
	// For example, `*123456789012* .dkr.ecr. *us-west-2* .amazonaws.com/repository` .
	AttrRepositoryUri() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The encryption configuration for the repository.
	//
	// This determines how the contents of your repository are encrypted at rest.
	EncryptionConfiguration() interface{}
	SetEncryptionConfiguration(val interface{})
	// The image scanning configuration for the repository.
	//
	// This determines whether images are scanned for known vulnerabilities after being pushed to the repository.
	ImageScanningConfiguration() interface{}
	SetImageScanningConfiguration(val interface{})
	// The tag mutability setting for the repository.
	//
	// If this parameter is omitted, the default setting of `MUTABLE` will be used which will allow image tags to be overwritten. If `IMMUTABLE` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.
	ImageTagMutability() *string
	SetImageTagMutability(val *string)
	// Creates or updates a lifecycle policy.
	//
	// For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html) .
	LifecyclePolicy() interface{}
	SetLifecyclePolicy(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The name to use for the repository.
	//
	// The repository name may be specified on its own (such as `nginx-web-app` ) or it can be prepended with a namespace to group the repository into a category (such as `project-a/nginx-web-app` ). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	RepositoryName() *string
	SetRepositoryName(val *string)
	// The JSON repository policy text to apply to the repository.
	//
	// For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide* .
	RepositoryPolicyText() interface{}
	SetRepositoryPolicyText(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
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

func (j *jsiiProxy_CfnRepository) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnRepository(scope awscdk.Construct, id *string, props *CfnRepositoryProps) CfnRepository {
	_init_.Initialize()

	j := jsiiProxy_CfnRepository{}

	_jsii_.Create(
		"monocdk.aws_ecr.CfnRepository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECR::Repository`.
func NewCfnRepository_Override(c CfnRepository, scope awscdk.Construct, id *string, props *CfnRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecr.CfnRepository",
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
// Experimental.
func CfnRepository_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnRepository",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnRepository_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnRepository",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnRepository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.CfnRepository",
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
		"monocdk.aws_ecr.CfnRepository",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRepository) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRepository) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRepository) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRepository) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRepository) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRepository) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRepository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnRepository) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRepository) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnRepository) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnRepository) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRepository) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRepository) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnRepository) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnRepository) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//   encryptionConfigurationProperty := &encryptionConfigurationProperty{
//   	encryptionType: jsii.String("encryptionType"),
//
//   	// the properties below are optional
//   	kmsKey: jsii.String("kmsKey"),
//   }
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//   imageScanningConfigurationProperty := &imageScanningConfigurationProperty{
//   	scanOnPush: jsii.Boolean(false),
//   }
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//   lifecyclePolicyProperty := &lifecyclePolicyProperty{
//   	lifecyclePolicyText: jsii.String("lifecyclePolicyText"),
//   	registryId: jsii.String("registryId"),
//   }
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//
//   var repositoryPolicyText interface{}
//   cfnRepositoryProps := &cfnRepositoryProps{
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		encryptionType: jsii.String("encryptionType"),
//
//   		// the properties below are optional
//   		kmsKey: jsii.String("kmsKey"),
//   	},
//   	imageScanningConfiguration: &imageScanningConfigurationProperty{
//   		scanOnPush: jsii.Boolean(false),
//   	},
//   	imageTagMutability: jsii.String("imageTagMutability"),
//   	lifecyclePolicy: &lifecyclePolicyProperty{
//   		lifecyclePolicyText: jsii.String("lifecyclePolicyText"),
//   		registryId: jsii.String("registryId"),
//   	},
//   	repositoryName: jsii.String("repositoryName"),
//   	repositoryPolicyText: repositoryPolicyText,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
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
// Experimental.
type IRepository interface {
	awscdk.IResource
	// Add a policy statement to the repository's resource policy.
	// Experimental.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Grant the given principal identity permissions to perform the actions on this repository.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to pull images in this repository.
	// Experimental.
	GrantPull(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to pull and push images to this repository.
	// Experimental.
	GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant
	// Define a CloudWatch event that triggers when something happens to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	// Experimental.
	OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event rule that can trigger a target when an image is pushed to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	// Experimental.
	OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers for repository events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event rule that can trigger a target when the image scan is completed.
	// Experimental.
	OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule
	// Returns the URI of the repository for a certain tag. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[@DIGEST]
	// Experimental.
	RepositoryUriForDigest(digest *string) *string
	// Returns the URI of the repository for a certain tag. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[:TAG]
	// Experimental.
	RepositoryUriForTag(tag *string) *string
	// The ARN of the repository.
	// Experimental.
	RepositoryArn() *string
	// The name of the repository.
	// Experimental.
	RepositoryName() *string
	// The URI of this repository (represents the latest image):.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY
	// Experimental.
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
// Example:
//   var repository repository
//
//   repository.addLifecycleRule(&lifecycleRule{
//   	tagPrefixList: []*string{
//   		jsii.String("prod"),
//   	},
//   	maxImageCount: jsii.Number(9999),
//   })
//   repository.addLifecycleRule(&lifecycleRule{
//   	maxImageAge: duration.days(jsii.Number(30)),
//   })
//
// Experimental.
type LifecycleRule struct {
	// Describes the purpose of the rule.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// The maximum age of images to retain. The value must represent a number of days.
	//
	// Specify exactly one of maxImageCount and maxImageAge.
	// Experimental.
	MaxImageAge awscdk.Duration `json:"maxImageAge" yaml:"maxImageAge"`
	// The maximum number of images to retain.
	//
	// Specify exactly one of maxImageCount and maxImageAge.
	// Experimental.
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
	// Experimental.
	RulePriority *float64 `json:"rulePriority" yaml:"rulePriority"`
	// Select images that have ALL the given prefixes in their tag.
	//
	// Only if tagStatus == TagStatus.Tagged
	// Experimental.
	TagPrefixList *[]*string `json:"tagPrefixList" yaml:"tagPrefixList"`
	// Select images based on tags.
	//
	// Only one rule is allowed to select untagged images, and it must
	// have the highest rulePriority.
	// Experimental.
	TagStatus TagStatus `json:"tagStatus" yaml:"tagStatus"`
}

// Options for the onCloudTrailImagePushed method.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"import awscdk "github.com/aws/aws-cdk-go/awscdk"import events "github.com/aws/aws-cdk-go/awscdk/aws_events"
//
//   var detail interface{}
//   var ruleTarget iRuleTarget
//   onCloudTrailImagePushedOptions := &onCloudTrailImagePushedOptions{
//   	description: jsii.String("description"),
//   	eventPattern: &eventPattern{
//   		account: []*string{
//   			jsii.String("account"),
//   		},
//   		detail: map[string]interface{}{
//   			"detailKey": detail,
//   		},
//   		detailType: []*string{
//   			jsii.String("detailType"),
//   		},
//   		id: []*string{
//   			jsii.String("id"),
//   		},
//   		region: []*string{
//   			jsii.String("region"),
//   		},
//   		resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		source: []*string{
//   			jsii.String("source"),
//   		},
//   		time: []*string{
//   			jsii.String("time"),
//   		},
//   		version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	imageTag: jsii.String("imageTag"),
//   	ruleName: jsii.String("ruleName"),
//   	target: ruleTarget,
//   }
//
// Experimental.
type OnCloudTrailImagePushedOptions struct {
	// A description of the rule's purpose.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Additional restrictions for the event to route to the specified target.
	//
	// The method that generates the rule probably imposes some type of event
	// filtering. The filtering implied by what you pass here is added
	// on top of that filtering.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	// Experimental.
	EventPattern *awsevents.EventPattern `json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	// Experimental.
	RuleName *string `json:"ruleName" yaml:"ruleName"`
	// The target to register for the event.
	// Experimental.
	Target awsevents.IRuleTarget `json:"target" yaml:"target"`
	// Only watch changes to this image tag.
	// Experimental.
	ImageTag *string `json:"imageTag" yaml:"imageTag"`
}

// Options for the OnImageScanCompleted method.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"import awscdk "github.com/aws/aws-cdk-go/awscdk"import events "github.com/aws/aws-cdk-go/awscdk/aws_events"
//
//   var detail interface{}
//   var ruleTarget iRuleTarget
//   onImageScanCompletedOptions := &onImageScanCompletedOptions{
//   	description: jsii.String("description"),
//   	eventPattern: &eventPattern{
//   		account: []*string{
//   			jsii.String("account"),
//   		},
//   		detail: map[string]interface{}{
//   			"detailKey": detail,
//   		},
//   		detailType: []*string{
//   			jsii.String("detailType"),
//   		},
//   		id: []*string{
//   			jsii.String("id"),
//   		},
//   		region: []*string{
//   			jsii.String("region"),
//   		},
//   		resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		source: []*string{
//   			jsii.String("source"),
//   		},
//   		time: []*string{
//   			jsii.String("time"),
//   		},
//   		version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	imageTags: []*string{
//   		jsii.String("imageTags"),
//   	},
//   	ruleName: jsii.String("ruleName"),
//   	target: ruleTarget,
//   }
//
// Experimental.
type OnImageScanCompletedOptions struct {
	// A description of the rule's purpose.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Additional restrictions for the event to route to the specified target.
	//
	// The method that generates the rule probably imposes some type of event
	// filtering. The filtering implied by what you pass here is added
	// on top of that filtering.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	// Experimental.
	EventPattern *awsevents.EventPattern `json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	// Experimental.
	RuleName *string `json:"ruleName" yaml:"ruleName"`
	// The target to register for the event.
	// Experimental.
	Target awsevents.IRuleTarget `json:"target" yaml:"target"`
	// Only watch changes to the image tags spedified.
	//
	// Leave it undefined to watch the full repository.
	// Experimental.
	ImageTags *[]*string `json:"imageTags" yaml:"imageTags"`
}

// Authorization token to access the global public ECR Gallery via Docker CLI.
//
// Example:
//   user := iam.NewUser(this, jsii.String("User"))
//   ecr.publicGalleryAuthorizationToken.grantRead(user)
//
// See: https://docs.aws.amazon.com/AmazonECR/latest/public/public-registries.html#public-registry-auth
//
// Experimental.
type PublicGalleryAuthorizationToken interface {
}

// The jsii proxy struct for PublicGalleryAuthorizationToken
type jsiiProxy_PublicGalleryAuthorizationToken struct {
	_ byte // padding
}

// Grant access to retrieve an authorization token.
// Experimental.
func PublicGalleryAuthorizationToken_GrantRead(grantee awsiam.IGrantable) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_ecr.PublicGalleryAuthorizationToken",
		"grantRead",
		[]interface{}{grantee},
	)
}

// Define an ECR repository.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromEcr(&ecrProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(80),
//   		},
//   		repository: ecr.repository.fromRepositoryName(this, jsii.String("NginxRepository"), jsii.String("nginx")),
//   		tag: jsii.String("latest"),
//   	}),
//   })
//
// Experimental.
type Repository interface {
	RepositoryBase
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The ARN of the repository.
	// Experimental.
	RepositoryArn() *string
	// The name of the repository.
	// Experimental.
	RepositoryName() *string
	// The URI of this repository (represents the latest image):.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY
	// Experimental.
	RepositoryUri() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Add a life cycle rule to the repository.
	//
	// Life cycle rules automatically expire images from the repository that match
	// certain conditions.
	// Experimental.
	AddLifecycleRule(rule *LifecycleRule)
	// Add a policy statement to the repository's resource policy.
	// Experimental.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant the given principal identity permissions to perform the actions on this repository.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to use the images in this repository.
	// Experimental.
	GrantPull(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to pull and push images to this repository.
	// Experimental.
	GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant
	// Define a CloudWatch event that triggers when something happens to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	// Experimental.
	OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event rule that can trigger a target when an image is pushed to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	// Experimental.
	OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers for repository events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event rule that can trigger a target when an image scan is completed.
	// Experimental.
	OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Returns the URL of the repository. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[@DIGEST]
	// Experimental.
	RepositoryUriForDigest(digest *string) *string
	// Returns the URL of the repository. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[:TAG]
	// Experimental.
	RepositoryUriForTag(tag *string) *string
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	// Experimental.
	Validate() *[]*string
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

func (j *jsiiProxy_Repository) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Experimental.
func NewRepository(scope constructs.Construct, id *string, props *RepositoryProps) Repository {
	_init_.Initialize()

	j := jsiiProxy_Repository{}

	_jsii_.Create(
		"monocdk.aws_ecr.Repository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRepository_Override(r Repository, scope constructs.Construct, id *string, props *RepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecr.Repository",
		[]interface{}{scope, id, props},
		r,
	)
}

// Returns an ECR ARN for a repository that resides in the same account/region as the current stack.
// Experimental.
func Repository_ArnForLocalRepository(repositoryName *string, scope constructs.IConstruct, account *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.Repository",
		"arnForLocalRepository",
		[]interface{}{repositoryName, scope, account},
		&returns,
	)

	return returns
}

// Experimental.
func Repository_FromRepositoryArn(scope constructs.Construct, id *string, repositoryArn *string) IRepository {
	_init_.Initialize()

	var returns IRepository

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.Repository",
		"fromRepositoryArn",
		[]interface{}{scope, id, repositoryArn},
		&returns,
	)

	return returns
}

// Import a repository.
// Experimental.
func Repository_FromRepositoryAttributes(scope constructs.Construct, id *string, attrs *RepositoryAttributes) IRepository {
	_init_.Initialize()

	var returns IRepository

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.Repository",
		"fromRepositoryAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Experimental.
func Repository_FromRepositoryName(scope constructs.Construct, id *string, repositoryName *string) IRepository {
	_init_.Initialize()

	var returns IRepository

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.Repository",
		"fromRepositoryName",
		[]interface{}{scope, id, repositoryName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Repository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.Repository",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Repository_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.Repository",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) AddLifecycleRule(rule *LifecycleRule) {
	_jsii_.InvokeVoid(
		r,
		"addLifecycleRule",
		[]interface{}{rule},
	)
}

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

func (r *jsiiProxy_Repository) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_Repository) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) Prepare() {
	_jsii_.InvokeVoid(
		r,
		"prepare",
		nil, // no parameters
	)
}

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

func (r *jsiiProxy_Repository) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (r *jsiiProxy_Repository) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr "github.com/aws/aws-cdk-go/awscdk/aws_ecr"
//   repositoryAttributes := &repositoryAttributes{
//   	repositoryArn: jsii.String("repositoryArn"),
//   	repositoryName: jsii.String("repositoryName"),
//   }
//
// Experimental.
type RepositoryAttributes struct {
	// Experimental.
	RepositoryArn *string `json:"repositoryArn" yaml:"repositoryArn"`
	// Experimental.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
}

// Base class for ECR repository.
//
// Reused between imported repositories and owned repositories.
// Experimental.
type RepositoryBase interface {
	awscdk.Resource
	IRepository
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The ARN of the repository.
	// Experimental.
	RepositoryArn() *string
	// The name of the repository.
	// Experimental.
	RepositoryName() *string
	// The URI of this repository (represents the latest image):.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY
	// Experimental.
	RepositoryUri() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Add a policy statement to the repository's resource policy.
	// Experimental.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant the given principal identity permissions to perform the actions on this repository.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to use the images in this repository.
	// Experimental.
	GrantPull(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to pull and push images to this repository.
	// Experimental.
	GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant
	// Define a CloudWatch event that triggers when something happens to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	// Experimental.
	OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event rule that can trigger a target when an image is pushed to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	// Experimental.
	OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers for repository events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event rule that can trigger a target when an image scan is completed.
	// Experimental.
	OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Returns the URL of the repository. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[@DIGEST]
	// Experimental.
	RepositoryUriForDigest(digest *string) *string
	// Returns the URL of the repository. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[:TAG]
	// Experimental.
	RepositoryUriForTag(tag *string) *string
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
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

func (j *jsiiProxy_RepositoryBase) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Experimental.
func NewRepositoryBase_Override(r RepositoryBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecr.RepositoryBase",
		[]interface{}{scope, id, props},
		r,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func RepositoryBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.RepositoryBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func RepositoryBase_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr.RepositoryBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

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

func (r *jsiiProxy_RepositoryBase) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryBase) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_RepositoryBase) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) Prepare() {
	_jsii_.InvokeVoid(
		r,
		"prepare",
		nil, // no parameters
	)
}

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

func (r *jsiiProxy_RepositoryBase) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (r *jsiiProxy_RepositoryBase) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Indicates whether server-side encryption is enabled for the object, and whether that encryption is from the AWS Key Management Service (AWS KMS) or from Amazon S3 managed encryption (SSE-S3).
//
// Example:
//   ecr.NewRepository(this, jsii.String("Repo"), &repositoryProps{
//   	encryption: ecr.repositoryEncryption_KMS(),
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
// Experimental.
type RepositoryEncryption interface {
	// the string value of the encryption.
	// Experimental.
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


// Experimental.
func NewRepositoryEncryption(value *string) RepositoryEncryption {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEncryption{}

	_jsii_.Create(
		"monocdk.aws_ecr.RepositoryEncryption",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewRepositoryEncryption_Override(r RepositoryEncryption, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecr.RepositoryEncryption",
		[]interface{}{value},
		r,
	)
}

func RepositoryEncryption_AES_256() RepositoryEncryption {
	_init_.Initialize()
	var returns RepositoryEncryption
	_jsii_.StaticGet(
		"monocdk.aws_ecr.RepositoryEncryption",
		"AES_256",
		&returns,
	)
	return returns
}

func RepositoryEncryption_KMS() RepositoryEncryption {
	_init_.Initialize()
	var returns RepositoryEncryption
	_jsii_.StaticGet(
		"monocdk.aws_ecr.RepositoryEncryption",
		"KMS",
		&returns,
	)
	return returns
}

// Example:
//   ecr.NewRepository(this, jsii.String("Repo"), &repositoryProps{
//   	imageTagMutability: ecr.tagMutability_IMMUTABLE,
//   })
//
// Experimental.
type RepositoryProps struct {
	// The kind of server-side encryption to apply to this repository.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryptionKey is not specified, an AWS managed KMS key is used.
	// Experimental.
	Encryption RepositoryEncryption `json:"encryption" yaml:"encryption"`
	// External KMS key to use for repository encryption.
	//
	// The 'encryption' property must be either not specified or set to "KMS".
	// An error will be emitted if encryption is set to "AES256".
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// Enable the scan on push when creating the repository.
	// Experimental.
	ImageScanOnPush *bool `json:"imageScanOnPush" yaml:"imageScanOnPush"`
	// The tag mutability setting for the repository.
	//
	// If this parameter is omitted, the default setting of MUTABLE will be used which will allow image tags to be overwritten.
	// Experimental.
	ImageTagMutability TagMutability `json:"imageTagMutability" yaml:"imageTagMutability"`
	// The AWS account ID associated with the registry that contains the repository.
	// See: https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_PutLifecyclePolicy.html
	//
	// Experimental.
	LifecycleRegistryId *string `json:"lifecycleRegistryId" yaml:"lifecycleRegistryId"`
	// Life cycle rules to apply to this registry.
	// Experimental.
	LifecycleRules *[]*LifecycleRule `json:"lifecycleRules" yaml:"lifecycleRules"`
	// Determine what happens to the repository when the resource/stack is deleted.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// Name for this repository.
	// Experimental.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
}

// The tag mutability setting for your repository.
//
// Example:
//   ecr.NewRepository(this, jsii.String("Repo"), &repositoryProps{
//   	imageTagMutability: ecr.tagMutability_IMMUTABLE,
//   })
//
// Experimental.
type TagMutability string

const (
	// allow image tags to be overwritten.
	// Experimental.
	TagMutability_MUTABLE TagMutability = "MUTABLE"
	// all image tags within the repository will be immutable which will prevent them from being overwritten.
	// Experimental.
	TagMutability_IMMUTABLE TagMutability = "IMMUTABLE"
)

// Select images based on tags.
// Experimental.
type TagStatus string

const (
	// Rule applies to all images.
	// Experimental.
	TagStatus_ANY TagStatus = "ANY"
	// Rule applies to tagged images.
	// Experimental.
	TagStatus_TAGGED TagStatus = "TAGGED"
	// Rule applies to untagged images.
	// Experimental.
	TagStatus_UNTAGGED TagStatus = "UNTAGGED"
)

