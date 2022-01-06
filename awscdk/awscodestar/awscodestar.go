package awscodestar

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestar/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::CodeStar::GitHubRepository`.
//
// The `AWS::CodeStar::GitHubRepository` resource creates a GitHub repository where users can store source code for use with AWS workflows. You must provide a location for the source code ZIP file in the AWS CloudFormation template, so the code can be uploaded to the created repository. You must have created a personal access token in GitHub to provide in the AWS CloudFormation template. AWS uses this token to connect to GitHub on your behalf. For more information about using a GitHub source repository with AWS CodeStar projects, see [AWS CodeStar Project Files and Resources](https://docs.aws.amazon.com/codestar/latest/userguide/templates.html#templates-whatis) .
//
// TODO: EXAMPLE
//
type CfnGitHubRepository interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Code() interface{}
	SetCode(val interface{})
	ConnectionArn() *string
	SetConnectionArn(val *string)
	CreationStack() *[]*string
	EnableIssues() interface{}
	SetEnableIssues(val interface{})
	IsPrivate() interface{}
	SetIsPrivate(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RepositoryAccessToken() *string
	SetRepositoryAccessToken(val *string)
	RepositoryDescription() *string
	SetRepositoryDescription(val *string)
	RepositoryName() *string
	SetRepositoryName(val *string)
	RepositoryOwner() *string
	SetRepositoryOwner(val *string)
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

// The jsii proxy struct for CfnGitHubRepository
type jsiiProxy_CfnGitHubRepository struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGitHubRepository) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) Code() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) EnableIssues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIssues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) IsPrivate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrivate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) RepositoryAccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryAccessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) RepositoryDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) RepositoryOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGitHubRepository) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CodeStar::GitHubRepository`.
func NewCfnGitHubRepository(scope constructs.Construct, id *string, props *CfnGitHubRepositoryProps) CfnGitHubRepository {
	_init_.Initialize()

	j := jsiiProxy_CfnGitHubRepository{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codestar.CfnGitHubRepository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeStar::GitHubRepository`.
func NewCfnGitHubRepository_Override(c CfnGitHubRepository, scope constructs.Construct, id *string, props *CfnGitHubRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codestar.CfnGitHubRepository",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGitHubRepository) SetCode(val interface{}) {
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_CfnGitHubRepository) SetConnectionArn(val *string) {
	_jsii_.Set(
		j,
		"connectionArn",
		val,
	)
}

func (j *jsiiProxy_CfnGitHubRepository) SetEnableIssues(val interface{}) {
	_jsii_.Set(
		j,
		"enableIssues",
		val,
	)
}

func (j *jsiiProxy_CfnGitHubRepository) SetIsPrivate(val interface{}) {
	_jsii_.Set(
		j,
		"isPrivate",
		val,
	)
}

func (j *jsiiProxy_CfnGitHubRepository) SetRepositoryAccessToken(val *string) {
	_jsii_.Set(
		j,
		"repositoryAccessToken",
		val,
	)
}

func (j *jsiiProxy_CfnGitHubRepository) SetRepositoryDescription(val *string) {
	_jsii_.Set(
		j,
		"repositoryDescription",
		val,
	)
}

func (j *jsiiProxy_CfnGitHubRepository) SetRepositoryName(val *string) {
	_jsii_.Set(
		j,
		"repositoryName",
		val,
	)
}

func (j *jsiiProxy_CfnGitHubRepository) SetRepositoryOwner(val *string) {
	_jsii_.Set(
		j,
		"repositoryOwner",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnGitHubRepository_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codestar.CfnGitHubRepository",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnGitHubRepository_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codestar.CfnGitHubRepository",
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
func CfnGitHubRepository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codestar.CfnGitHubRepository",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGitHubRepository_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codestar.CfnGitHubRepository",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnGitHubRepository) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnGitHubRepository) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnGitHubRepository) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnGitHubRepository) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnGitHubRepository) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnGitHubRepository) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnGitHubRepository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnGitHubRepository) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnGitHubRepository) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnGitHubRepository) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnGitHubRepository) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGitHubRepository) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnGitHubRepository) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnGitHubRepository) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGitHubRepository) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `Code` property type specifies information about code to be committed.
//
// `Code` is a property of the `AWS::CodeStar::GitHubRepository` resource.
//
// TODO: EXAMPLE
//
type CfnGitHubRepository_CodeProperty struct {
	// Information about the Amazon S3 bucket that contains a ZIP file of code to be committed to the repository.
	S3 interface{} `json:"s3"`
}

// The `S3` property type specifies information about the Amazon S3 bucket that contains the code to be committed to the new repository.
//
// `S3` is a property of the `AWS::CodeStar::GitHubRepository` resource.
//
// TODO: EXAMPLE
//
type CfnGitHubRepository_S3Property struct {
	// The name of the Amazon S3 bucket that contains the ZIP file with the content to be committed to the new repository.
	Bucket *string `json:"bucket"`
	// The S3 object key or file name for the ZIP file.
	Key *string `json:"key"`
	// The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket.
	ObjectVersion *string `json:"objectVersion"`
}

// Properties for defining a `CfnGitHubRepository`.
//
// TODO: EXAMPLE
//
type CfnGitHubRepositoryProps struct {
	// The name of the repository you want to create in GitHub with AWS CloudFormation stack creation.
	RepositoryName *string `json:"repositoryName"`
	// The GitHub user name for the owner of the GitHub repository to be created.
	//
	// If this repository should be owned by a GitHub organization, provide its name.
	RepositoryOwner *string `json:"repositoryOwner"`
	// Information about code to be committed to a repository after it is created in an AWS CloudFormation stack.
	Code interface{} `json:"code"`
	// `AWS::CodeStar::GitHubRepository.ConnectionArn`.
	ConnectionArn *string `json:"connectionArn"`
	// Indicates whether to enable issues for the GitHub repository.
	//
	// You can use GitHub issues to track information and bugs for your repository.
	EnableIssues interface{} `json:"enableIssues"`
	// Indicates whether the GitHub repository is a private repository.
	//
	// If so, you choose who can see and commit to this repository.
	IsPrivate interface{} `json:"isPrivate"`
	// The GitHub user's personal access token for the GitHub repository.
	RepositoryAccessToken *string `json:"repositoryAccessToken"`
	// A comment or description about the new repository.
	//
	// This description is displayed in GitHub after the repository is created.
	RepositoryDescription *string `json:"repositoryDescription"`
}

