package awscodestar

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodestar/internal"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::CodeStar::GitHubRepository`.
//
// The `AWS::CodeStar::GitHubRepository` resource creates a GitHub repository where users can store source code for use with AWS workflows. You must provide a location for the source code ZIP file in the AWS CloudFormation template, so the code can be uploaded to the created repository. You must have created a personal access token in GitHub to provide in the AWS CloudFormation template. AWS uses this token to connect to GitHub on your behalf. For more information about using a GitHub source repository with AWS CodeStar projects, see [AWS CodeStar Project Files and Resources](https://docs.aws.amazon.com/codestar/latest/userguide/templates.html#templates-whatis) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGitHubRepository := awscdk.Aws_codestar.NewCfnGitHubRepository(this, jsii.String("MyCfnGitHubRepository"), &cfnGitHubRepositoryProps{
//   	repositoryName: jsii.String("repositoryName"),
//   	repositoryOwner: jsii.String("repositoryOwner"),
//
//   	// the properties below are optional
//   	code: &codeProperty{
//   		s3: &s3Property{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			objectVersion: jsii.String("objectVersion"),
//   		},
//   	},
//   	connectionArn: jsii.String("connectionArn"),
//   	enableIssues: jsii.Boolean(false),
//   	isPrivate: jsii.Boolean(false),
//   	repositoryAccessToken: jsii.String("repositoryAccessToken"),
//   	repositoryDescription: jsii.String("repositoryDescription"),
//   })
//
type CfnGitHubRepository interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Information about code to be committed to a repository after it is created in an AWS CloudFormation stack.
	Code() interface{}
	SetCode(val interface{})
	// `AWS::CodeStar::GitHubRepository.ConnectionArn`.
	ConnectionArn() *string
	SetConnectionArn(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Indicates whether to enable issues for the GitHub repository.
	//
	// You can use GitHub issues to track information and bugs for your repository.
	EnableIssues() interface{}
	SetEnableIssues(val interface{})
	// Indicates whether the GitHub repository is a private repository.
	//
	// If so, you choose who can see and commit to this repository.
	IsPrivate() interface{}
	SetIsPrivate(val interface{})
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
	// The GitHub user's personal access token for the GitHub repository.
	RepositoryAccessToken() *string
	SetRepositoryAccessToken(val *string)
	// A comment or description about the new repository.
	//
	// This description is displayed in GitHub after the repository is created.
	RepositoryDescription() *string
	SetRepositoryDescription(val *string)
	// The name of the repository you want to create in GitHub with AWS CloudFormation stack creation.
	RepositoryName() *string
	SetRepositoryName(val *string)
	// The GitHub user name for the owner of the GitHub repository to be created.
	//
	// If this repository should be owned by a GitHub organization, provide its name.
	RepositoryOwner() *string
	SetRepositoryOwner(val *string)
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

func (j *jsiiProxy_CfnGitHubRepository) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnGitHubRepository(scope awscdk.Construct, id *string, props *CfnGitHubRepositoryProps) CfnGitHubRepository {
	_init_.Initialize()

	j := jsiiProxy_CfnGitHubRepository{}

	_jsii_.Create(
		"monocdk.aws_codestar.CfnGitHubRepository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeStar::GitHubRepository`.
func NewCfnGitHubRepository_Override(c CfnGitHubRepository, scope awscdk.Construct, id *string, props *CfnGitHubRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codestar.CfnGitHubRepository",
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
// Experimental.
func CfnGitHubRepository_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codestar.CfnGitHubRepository",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnGitHubRepository_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codestar.CfnGitHubRepository",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnGitHubRepository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codestar.CfnGitHubRepository",
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
		"monocdk.aws_codestar.CfnGitHubRepository",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGitHubRepository) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGitHubRepository) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGitHubRepository) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGitHubRepository) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGitHubRepository) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGitHubRepository) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGitHubRepository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnGitHubRepository) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnGitHubRepository) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnGitHubRepository) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnGitHubRepository) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGitHubRepository) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGitHubRepository) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnGitHubRepository) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnGitHubRepository) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeProperty := &codeProperty{
//   	s3: &s3Property{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
type CfnGitHubRepository_CodeProperty struct {
	// Information about the Amazon S3 bucket that contains a ZIP file of code to be committed to the repository.
	S3 interface{} `field:"required" json:"s3" yaml:"s3"`
}

// The `S3` property type specifies information about the Amazon S3 bucket that contains the code to be committed to the new repository.
//
// `S3` is a property of the `AWS::CodeStar::GitHubRepository` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3Property := &s3Property{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	objectVersion: jsii.String("objectVersion"),
//   }
//
type CfnGitHubRepository_S3Property struct {
	// The name of the Amazon S3 bucket that contains the ZIP file with the content to be committed to the new repository.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 object key or file name for the ZIP file.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

// Properties for defining a `CfnGitHubRepository`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGitHubRepositoryProps := &cfnGitHubRepositoryProps{
//   	repositoryName: jsii.String("repositoryName"),
//   	repositoryOwner: jsii.String("repositoryOwner"),
//
//   	// the properties below are optional
//   	code: &codeProperty{
//   		s3: &s3Property{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			objectVersion: jsii.String("objectVersion"),
//   		},
//   	},
//   	connectionArn: jsii.String("connectionArn"),
//   	enableIssues: jsii.Boolean(false),
//   	isPrivate: jsii.Boolean(false),
//   	repositoryAccessToken: jsii.String("repositoryAccessToken"),
//   	repositoryDescription: jsii.String("repositoryDescription"),
//   }
//
type CfnGitHubRepositoryProps struct {
	// The name of the repository you want to create in GitHub with AWS CloudFormation stack creation.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// The GitHub user name for the owner of the GitHub repository to be created.
	//
	// If this repository should be owned by a GitHub organization, provide its name.
	RepositoryOwner *string `field:"required" json:"repositoryOwner" yaml:"repositoryOwner"`
	// Information about code to be committed to a repository after it is created in an AWS CloudFormation stack.
	Code interface{} `field:"optional" json:"code" yaml:"code"`
	// `AWS::CodeStar::GitHubRepository.ConnectionArn`.
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
	// Indicates whether to enable issues for the GitHub repository.
	//
	// You can use GitHub issues to track information and bugs for your repository.
	EnableIssues interface{} `field:"optional" json:"enableIssues" yaml:"enableIssues"`
	// Indicates whether the GitHub repository is a private repository.
	//
	// If so, you choose who can see and commit to this repository.
	IsPrivate interface{} `field:"optional" json:"isPrivate" yaml:"isPrivate"`
	// The GitHub user's personal access token for the GitHub repository.
	RepositoryAccessToken *string `field:"optional" json:"repositoryAccessToken" yaml:"repositoryAccessToken"`
	// A comment or description about the new repository.
	//
	// This description is displayed in GitHub after the repository is created.
	RepositoryDescription *string `field:"optional" json:"repositoryDescription" yaml:"repositoryDescription"`
}

// The GitHubRepository resource.
//
// Example:
//   import codestar "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   codestar.NewGitHubRepository(this, jsii.String("GitHubRepo"), &gitHubRepositoryProps{
//   	owner: jsii.String("aws"),
//   	repositoryName: jsii.String("aws-cdk"),
//   	accessToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token"), &secretsManagerSecretOptions{
//   		jsonField: jsii.String("token"),
//   	}),
//   	contentsBucket: s3.bucket.fromBucketName(this, jsii.String("Bucket"), jsii.String("bucket-name")),
//   	contentsKey: jsii.String("import.zip"),
//   })
//
// Experimental.
type GitHubRepository interface {
	awscdk.Resource
	IGitHubRepository
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
	// the repository owner.
	// Experimental.
	Owner() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// the repository name.
	// Experimental.
	Repo() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
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

// The jsii proxy struct for GitHubRepository
type jsiiProxy_GitHubRepository struct {
	internal.Type__awscdkResource
	jsiiProxy_IGitHubRepository
}

func (j *jsiiProxy_GitHubRepository) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) Repo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitHubRepository(scope constructs.Construct, id *string, props *GitHubRepositoryProps) GitHubRepository {
	_init_.Initialize()

	j := jsiiProxy_GitHubRepository{}

	_jsii_.Create(
		"monocdk.aws_codestar.GitHubRepository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubRepository_Override(g GitHubRepository, scope constructs.Construct, id *string, props *GitHubRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codestar.GitHubRepository",
		[]interface{}{scope, id, props},
		g,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func GitHubRepository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codestar.GitHubRepository",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func GitHubRepository_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codestar.GitHubRepository",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRepository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (g *jsiiProxy_GitHubRepository) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRepository) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRepository) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRepository) OnPrepare() {
	_jsii_.InvokeVoid(
		g,
		"onPrepare",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitHubRepository) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (g *jsiiProxy_GitHubRepository) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRepository) Prepare() {
	_jsii_.InvokeVoid(
		g,
		"prepare",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitHubRepository) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		[]interface{}{session},
	)
}

func (g *jsiiProxy_GitHubRepository) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRepository) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties of {@link GitHubRepository}.
//
// Example:
//   import codestar "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   codestar.NewGitHubRepository(this, jsii.String("GitHubRepo"), &gitHubRepositoryProps{
//   	owner: jsii.String("aws"),
//   	repositoryName: jsii.String("aws-cdk"),
//   	accessToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token"), &secretsManagerSecretOptions{
//   		jsonField: jsii.String("token"),
//   	}),
//   	contentsBucket: s3.bucket.fromBucketName(this, jsii.String("Bucket"), jsii.String("bucket-name")),
//   	contentsKey: jsii.String("import.zip"),
//   })
//
// Experimental.
type GitHubRepositoryProps struct {
	// The GitHub user's personal access token for the GitHub repository.
	// Experimental.
	AccessToken awscdk.SecretValue `field:"required" json:"accessToken" yaml:"accessToken"`
	// The name of the Amazon S3 bucket that contains the ZIP file with the content to be committed to the new repository.
	// Experimental.
	ContentsBucket awss3.IBucket `field:"required" json:"contentsBucket" yaml:"contentsBucket"`
	// The S3 object key or file name for the ZIP file.
	// Experimental.
	ContentsKey *string `field:"required" json:"contentsKey" yaml:"contentsKey"`
	// The GitHub user name for the owner of the GitHub repository to be created.
	//
	// If this
	// repository should be owned by a GitHub organization, provide its name.
	// Experimental.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the repository you want to create in GitHub with AWS CloudFormation stack creation.
	// Experimental.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket.
	// Experimental.
	ContentsS3Version *string `field:"optional" json:"contentsS3Version" yaml:"contentsS3Version"`
	// A comment or description about the new repository.
	//
	// This description is displayed in GitHub after the repository
	// is created.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether to enable issues for the GitHub repository.
	//
	// You can use GitHub issues to track information
	// and bugs for your repository.
	// Experimental.
	EnableIssues *bool `field:"optional" json:"enableIssues" yaml:"enableIssues"`
	// Indicates whether the GitHub repository is a private repository.
	//
	// If so, you choose who can see and commit to
	// this repository.
	// Experimental.
	Visibility RepositoryVisibility `field:"optional" json:"visibility" yaml:"visibility"`
}

// GitHubRepository resource interface.
// Experimental.
type IGitHubRepository interface {
	awscdk.IResource
	// the repository owner.
	// Experimental.
	Owner() *string
	// the repository name.
	// Experimental.
	Repo() *string
}

// The jsii proxy for IGitHubRepository
type jsiiProxy_IGitHubRepository struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IGitHubRepository) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGitHubRepository) Repo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repo",
		&returns,
	)
	return returns
}

// Visibility of the GitHubRepository.
// Experimental.
type RepositoryVisibility string

const (
	// private repository.
	// Experimental.
	RepositoryVisibility_PRIVATE RepositoryVisibility = "PRIVATE"
	// public repository.
	// Experimental.
	RepositoryVisibility_PUBLIC RepositoryVisibility = "PUBLIC"
)

