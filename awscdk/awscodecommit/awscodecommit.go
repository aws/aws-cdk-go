package awscodecommit

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodecommit/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::CodeCommit::Repository`.
//
// Creates a new, empty repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codecommit "github.com/aws/aws-cdk-go/awscdk/aws_codecommit"
//   cfnRepository := codecommit.NewCfnRepository(this, jsii.String("MyCfnRepository"), &cfnRepositoryProps{
//   	repositoryName: jsii.String("repositoryName"),
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
//
//   		// the properties below are optional
//   		branchName: jsii.String("branchName"),
//   	},
//   	repositoryDescription: jsii.String("repositoryDescription"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	triggers: []interface{}{
//   		&repositoryTriggerProperty{
//   			destinationArn: jsii.String("destinationArn"),
//   			events: []*string{
//   				jsii.String("events"),
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			branches: []*string{
//   				jsii.String("branches"),
//   			},
//   			customData: jsii.String("customData"),
//   		},
//   	},
//   })
//
type CfnRepository interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// When you pass the logical ID of this resource, the function returns the Amazon Resource Name (ARN) of the repository.
	AttrArn() *string
	// When you pass the logical ID of this resource, the function returns the URL to use for cloning the repository over HTTPS.
	AttrCloneUrlHttp() *string
	// When you pass the logical ID of this resource, the function returns the URL to use for cloning the repository over SSH.
	AttrCloneUrlSsh() *string
	// When you pass the logical ID of this resource, the function returns the repository's name.
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Information about code to be committed to a repository after it is created in an AWS CloudFormation stack.
	//
	// Information about code is only used in resource creation. Updates to a stack will not reflect changes made to code properties after initial resource creation.
	//
	// > You can only use this property to add code when creating a repository with a AWS CloudFormation template at creation time. This property cannot be used for updating code to an existing repository.
	Code() interface{}
	SetCode(val interface{})
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
	// A comment or description about the new repository.
	//
	// > The description field for a repository accepts all HTML characters and all valid Unicode characters. Applications that do not HTML-encode the description and display it in a webpage can expose users to potentially malicious code. Make sure that you HTML-encode the description field in any application that uses this API to display the repository description on a webpage.
	RepositoryDescription() *string
	SetRepositoryDescription(val *string)
	// The name of the new repository to be created.
	//
	// > The repository name must be unique across the calling AWS account . Repository names are limited to 100 alphanumeric, dash, and underscore characters, and cannot include certain characters. For more information about the limits on repository names, see [Quotas](https://docs.aws.amazon.com/codecommit/latest/userguide/limits.html) in the *AWS CodeCommit User Guide* . The suffix .git is prohibited.
	RepositoryName() *string
	SetRepositoryName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// One or more tag key-value pairs to use when tagging this repository.
	Tags() awscdk.TagManager
	// The JSON block of configuration information for each trigger.
	Triggers() interface{}
	SetTriggers(val interface{})
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

func (j *jsiiProxy_CfnRepository) AttrCloneUrlHttp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCloneUrlHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) AttrCloneUrlSsh() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCloneUrlSsh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRepository) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
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

func (j *jsiiProxy_CfnRepository) Code() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"code",
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

func (j *jsiiProxy_CfnRepository) RepositoryDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryDescription",
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

func (j *jsiiProxy_CfnRepository) Triggers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggers",
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


// Create a new `AWS::CodeCommit::Repository`.
func NewCfnRepository(scope awscdk.Construct, id *string, props *CfnRepositoryProps) CfnRepository {
	_init_.Initialize()

	j := jsiiProxy_CfnRepository{}

	_jsii_.Create(
		"monocdk.aws_codecommit.CfnRepository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeCommit::Repository`.
func NewCfnRepository_Override(c CfnRepository, scope awscdk.Construct, id *string, props *CfnRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codecommit.CfnRepository",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRepository) SetCode(val interface{}) {
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_CfnRepository) SetRepositoryDescription(val *string) {
	_jsii_.Set(
		j,
		"repositoryDescription",
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

func (j *jsiiProxy_CfnRepository) SetTriggers(val interface{}) {
	_jsii_.Set(
		j,
		"triggers",
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
		"monocdk.aws_codecommit.CfnRepository",
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
		"monocdk.aws_codecommit.CfnRepository",
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
		"monocdk.aws_codecommit.CfnRepository",
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
		"monocdk.aws_codecommit.CfnRepository",
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

// Information about code to be committed.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codecommit "github.com/aws/aws-cdk-go/awscdk/aws_codecommit"
//   codeProperty := &codeProperty{
//   	s3: &s3Property{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//
//   	// the properties below are optional
//   	branchName: jsii.String("branchName"),
//   }
//
type CfnRepository_CodeProperty struct {
	// Information about the Amazon S3 bucket that contains a ZIP file of code to be committed to the repository.
	//
	// Changes to this property are ignored after initial resource creation.
	S3 interface{} `json:"s3" yaml:"s3"`
	// Optional.
	//
	// Specifies a branch name to be used as the default branch when importing code into a repository on initial creation. If this property is not set, the name *main* will be used for the default branch for the repository. Changes to this property are ignored after initial resource creation. We recommend using this parameter to set the name to *main* to align with the default behavior of CodeCommit unless another name is needed.
	BranchName *string `json:"branchName" yaml:"branchName"`
}

// Information about a trigger for a repository.
//
// > If you want to receive notifications about repository events, consider using notifications instead of triggers. For more information, see [Configuring notifications for repository events](https://docs.aws.amazon.com/codecommit/latest/userguide/how-to-repository-email.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codecommit "github.com/aws/aws-cdk-go/awscdk/aws_codecommit"
//   repositoryTriggerProperty := &repositoryTriggerProperty{
//   	destinationArn: jsii.String("destinationArn"),
//   	events: []*string{
//   		jsii.String("events"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	branches: []*string{
//   		jsii.String("branches"),
//   	},
//   	customData: jsii.String("customData"),
//   }
//
type CfnRepository_RepositoryTriggerProperty struct {
	// The ARN of the resource that is the target for a trigger (for example, the ARN of a topic in Amazon SNS).
	DestinationArn *string `json:"destinationArn" yaml:"destinationArn"`
	// The repository events that cause the trigger to run actions in another service, such as sending a notification through Amazon SNS.
	//
	// > The valid value "all" cannot be used with any other values.
	Events *[]*string `json:"events" yaml:"events"`
	// The name of the trigger.
	Name *string `json:"name" yaml:"name"`
	// The branches to be included in the trigger configuration.
	//
	// If you specify an empty array, the trigger applies to all branches.
	//
	// > Although no content is required in the array, you must include the array itself.
	Branches *[]*string `json:"branches" yaml:"branches"`
	// Any custom data associated with the trigger to be included in the information sent to the target of the trigger.
	CustomData *string `json:"customData" yaml:"customData"`
}

// Information about the Amazon S3 bucket that contains the code that will be committed to the new repository.
//
// Changes to this property are ignored after initial resource creation.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codecommit "github.com/aws/aws-cdk-go/awscdk/aws_codecommit"
//   s3Property := &s3Property{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	objectVersion: jsii.String("objectVersion"),
//   }
//
type CfnRepository_S3Property struct {
	// The name of the Amazon S3 bucket that contains the ZIP file with the content that will be committed to the new repository.
	//
	// This can be specified using the name of the bucket in the AWS account . Changes to this property are ignored after initial resource creation.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// The key to use for accessing the Amazon S3 bucket.
	//
	// Changes to this property are ignored after initial resource creation. For more information, see [Creating object key names](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html) and [Uploading objects](https://docs.aws.amazon.com/AmazonS3/latest/userguide/upload-objects.html) in the Amazon S3 User Guide.
	Key *string `json:"key" yaml:"key"`
	// The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket.
	//
	// Changes to this property are ignored after initial resource creation.
	ObjectVersion *string `json:"objectVersion" yaml:"objectVersion"`
}

// Properties for defining a `CfnRepository`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codecommit "github.com/aws/aws-cdk-go/awscdk/aws_codecommit"
//   cfnRepositoryProps := &cfnRepositoryProps{
//   	repositoryName: jsii.String("repositoryName"),
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
//
//   		// the properties below are optional
//   		branchName: jsii.String("branchName"),
//   	},
//   	repositoryDescription: jsii.String("repositoryDescription"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	triggers: []interface{}{
//   		&repositoryTriggerProperty{
//   			destinationArn: jsii.String("destinationArn"),
//   			events: []*string{
//   				jsii.String("events"),
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			branches: []*string{
//   				jsii.String("branches"),
//   			},
//   			customData: jsii.String("customData"),
//   		},
//   	},
//   }
//
type CfnRepositoryProps struct {
	// The name of the new repository to be created.
	//
	// > The repository name must be unique across the calling AWS account . Repository names are limited to 100 alphanumeric, dash, and underscore characters, and cannot include certain characters. For more information about the limits on repository names, see [Quotas](https://docs.aws.amazon.com/codecommit/latest/userguide/limits.html) in the *AWS CodeCommit User Guide* . The suffix .git is prohibited.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
	// Information about code to be committed to a repository after it is created in an AWS CloudFormation stack.
	//
	// Information about code is only used in resource creation. Updates to a stack will not reflect changes made to code properties after initial resource creation.
	//
	// > You can only use this property to add code when creating a repository with a AWS CloudFormation template at creation time. This property cannot be used for updating code to an existing repository.
	Code interface{} `json:"code" yaml:"code"`
	// A comment or description about the new repository.
	//
	// > The description field for a repository accepts all HTML characters and all valid Unicode characters. Applications that do not HTML-encode the description and display it in a webpage can expose users to potentially malicious code. Make sure that you HTML-encode the description field in any application that uses this API to display the repository description on a webpage.
	RepositoryDescription *string `json:"repositoryDescription" yaml:"repositoryDescription"`
	// One or more tag key-value pairs to use when tagging this repository.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The JSON block of configuration information for each trigger.
	Triggers interface{} `json:"triggers" yaml:"triggers"`
}

// Represents the contents to initialize the repository with.
//
// Example:
//   repo := codecommit.NewRepository(this, jsii.String("Repository"), &repositoryProps{
//   	repositoryName: jsii.String("MyRepositoryName"),
//   	code: codecommit.code.fromDirectory(path.join(__dirname, jsii.String("directory/")), jsii.String("develop")),
//   })
//
// Experimental.
type Code interface {
	// This method is called after a repository is passed this instance of Code in its 'code' property.
	// Experimental.
	Bind(scope constructs.Construct) *CodeConfig
}

// The jsii proxy struct for Code
type jsiiProxy_Code struct {
	_ byte // padding
}

// Experimental.
func NewCode_Override(c Code) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codecommit.Code",
		nil, // no parameters
		c,
	)
}

// Code from user-supplied asset.
// Experimental.
func Code_FromAsset(asset awss3assets.Asset, branch *string) Code {
	_init_.Initialize()

	var returns Code

	_jsii_.StaticInvoke(
		"monocdk.aws_codecommit.Code",
		"fromAsset",
		[]interface{}{asset, branch},
		&returns,
	)

	return returns
}

// Code from directory.
// Experimental.
func Code_FromDirectory(directoryPath *string, branch *string) Code {
	_init_.Initialize()

	var returns Code

	_jsii_.StaticInvoke(
		"monocdk.aws_codecommit.Code",
		"fromDirectory",
		[]interface{}{directoryPath, branch},
		&returns,
	)

	return returns
}

// Code from preexisting ZIP file.
// Experimental.
func Code_FromZipFile(filePath *string, branch *string) Code {
	_init_.Initialize()

	var returns Code

	_jsii_.StaticInvoke(
		"monocdk.aws_codecommit.Code",
		"fromZipFile",
		[]interface{}{filePath, branch},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Code) Bind(scope constructs.Construct) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Represents the structure to pass into the underlying CfnRepository class.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codecommit "github.com/aws/aws-cdk-go/awscdk/aws_codecommit"
//   codeConfig := &codeConfig{
//   	code: &codeProperty{
//   		s3: &s3Property{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			objectVersion: jsii.String("objectVersion"),
//   		},
//
//   		// the properties below are optional
//   		branchName: jsii.String("branchName"),
//   	},
//   }
//
// Experimental.
type CodeConfig struct {
	// represents the underlying code structure.
	// Experimental.
	Code *CfnRepository_CodeProperty `json:"code" yaml:"code"`
}

// Experimental.
type IRepository interface {
	awscodestarnotifications.INotificationRuleSource
	awscdk.IResource
	// Grant the given principal identity permissions to perform the actions on this repository.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to pull this repository.
	// Experimental.
	GrantPull(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to pull and push this repository.
	// Experimental.
	GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to read this repository.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Defines a CodeStar Notification rule which triggers when a pull request is merged.
	// Deprecated: this method has a typo in its name, use notifyOnPullRequestMerged instead.
	NotifiyOnPullRequestMerged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule triggered when the project events specified by you are emitted. Similar to `onEvent` API.
	//
	// You can also use the methods to define rules for the specific event emitted.
	// eg: `notifyOnPullRequstCreated`.
	//
	// Returns: CodeStar Notifications rule associated with this repository.
	// Experimental.
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *RepositoryNotifyOnOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when an approval rule is overridden.
	// Experimental.
	NotifyOnApprovalRuleOverridden(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when an approval status is changed.
	// Experimental.
	NotifyOnApprovalStatusChanged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a new branch or tag is created.
	// Experimental.
	NotifyOnBranchOrTagCreated(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a branch or tag is deleted.
	// Experimental.
	NotifyOnBranchOrTagDeleted(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a comment is made on a pull request.
	// Experimental.
	NotifyOnPullRequestComment(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a pull request is created.
	// Experimental.
	NotifyOnPullRequestCreated(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a pull request is merged.
	// Experimental.
	NotifyOnPullRequestMerged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CloudWatch event rule which triggers when a comment is made on a commit.
	// Experimental.
	OnCommentOnCommit(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a comment is made on a pull request.
	// Experimental.
	OnCommentOnPullRequest(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a commit is pushed to a branch.
	// Experimental.
	OnCommit(id *string, options *OnCommitOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers for repository events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a pull request state is changed.
	// Experimental.
	OnPullRequestStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a reference is created (i.e. a new branch/tag is created) to the repository.
	// Experimental.
	OnReferenceCreated(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a reference is delete (i.e. a branch/tag is deleted) from the repository.
	// Experimental.
	OnReferenceDeleted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a reference is updated (i.e. a commit is pushed to an existing or new branch) from the repository.
	// Experimental.
	OnReferenceUpdated(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a "CodeCommit Repository State Change" event occurs.
	// Experimental.
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The ARN of this Repository.
	// Experimental.
	RepositoryArn() *string
	// The HTTPS (GRC) clone URL.
	//
	// HTTPS (GRC) is the protocol to use with git-remote-codecommit (GRC).
	//
	// It is the recommended method for supporting connections made with federated
	// access, identity providers, and temporary credentials.
	// See: https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-git-remote-codecommit.html
	//
	// Experimental.
	RepositoryCloneUrlGrc() *string
	// The HTTP clone URL.
	// Experimental.
	RepositoryCloneUrlHttp() *string
	// The SSH clone URL.
	// Experimental.
	RepositoryCloneUrlSsh() *string
	// The human-visible name of this Repository.
	// Experimental.
	RepositoryName() *string
}

// The jsii proxy for IRepository
type jsiiProxy_IRepository struct {
	internal.Type__awscodestarnotificationsINotificationRuleSource
	internal.Type__awscdkIResource
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

func (i *jsiiProxy_IRepository) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifiyOnPullRequestMerged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifiyOnPullRequestMerged",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *RepositoryNotifyOnOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOn",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnApprovalRuleOverridden(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnApprovalRuleOverridden",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnApprovalStatusChanged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnApprovalStatusChanged",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnBranchOrTagCreated(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnBranchOrTagCreated",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnBranchOrTagDeleted(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnBranchOrTagDeleted",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnPullRequestComment(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnPullRequestComment",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnPullRequestCreated(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnPullRequestCreated",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnPullRequestMerged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnPullRequestMerged",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnCommentOnCommit(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCommentOnCommit",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnCommentOnPullRequest(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCommentOnPullRequest",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnCommit(id *string, options *OnCommitOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCommit",
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

func (i *jsiiProxy_IRepository) OnPullRequestStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onPullRequestStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnReferenceCreated(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onReferenceCreated",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnReferenceDeleted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onReferenceDeleted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnReferenceUpdated(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onReferenceUpdated",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IRepository) BindAsNotificationRuleSource(scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig {
	var returns *awscodestarnotifications.NotificationRuleSourceConfig

	_jsii_.Invoke(
		i,
		"bindAsNotificationRuleSource",
		[]interface{}{scope},
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

func (j *jsiiProxy_IRepository) RepositoryCloneUrlGrc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCloneUrlGrc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) RepositoryCloneUrlHttp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCloneUrlHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) RepositoryCloneUrlSsh() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCloneUrlSsh",
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

func (j *jsiiProxy_IRepository) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Options for the onCommit() method.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var repo repository
//   myTopic := sns.NewTopic(this, jsii.String("Topic"))
//
//   repo.onCommit(jsii.String("OnCommit"), &onCommitOptions{
//   	target: targets.NewSnsTopic(myTopic),
//   })
//
// Experimental.
type OnCommitOptions struct {
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
	// The branch to monitor.
	// Experimental.
	Branches *[]*string `json:"branches" yaml:"branches"`
}

// Fields of CloudWatch Events that change references.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#codebuild_event_type
//
// Experimental.
type ReferenceEvent interface {
}

// The jsii proxy struct for ReferenceEvent
type jsiiProxy_ReferenceEvent struct {
	_ byte // padding
}

func ReferenceEvent_CommitId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"commitId",
		&returns,
	)
	return returns
}

func ReferenceEvent_EventType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"eventType",
		&returns,
	)
	return returns
}

func ReferenceEvent_ReferenceFullName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"referenceFullName",
		&returns,
	)
	return returns
}

func ReferenceEvent_ReferenceName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"referenceName",
		&returns,
	)
	return returns
}

func ReferenceEvent_ReferenceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"referenceType",
		&returns,
	)
	return returns
}

func ReferenceEvent_RepositoryId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"repositoryId",
		&returns,
	)
	return returns
}

func ReferenceEvent_RepositoryName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_codecommit.ReferenceEvent",
		"repositoryName",
		&returns,
	)
	return returns
}

// Provides a CodeCommit Repository.
//
// Example:
//   var project pipelineProject
//   repository := codecommit.NewRepository(this, jsii.String("MyRepository"), &repositoryProps{
//   	repositoryName: jsii.String("MyRepository"),
//   })
//   project := codebuild.NewPipelineProject(this, jsii.String("MyProject"))
//
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("CodeCommit"),
//   	repository: repository,
//   	output: sourceOutput,
//   })
//   buildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   	actionName: jsii.String("CodeBuild"),
//   	project: project,
//   	input: sourceOutput,
//   	outputs: []artifact{
//   		codepipeline.NewArtifact(),
//   	},
//   	 // optional
//   	executeBatchBuild: jsii.Boolean(true),
//   	 // optional, defaults to false
//   	combineBatchBuildArtifacts: jsii.Boolean(true),
//   })
//
//   codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		&stageProps{
//   			stageName: jsii.String("Source"),
//   			actions: []iAction{
//   				sourceAction,
//   			},
//   		},
//   		&stageProps{
//   			stageName: jsii.String("Build"),
//   			actions: []*iAction{
//   				buildAction,
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type Repository interface {
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
	// The ARN of this Repository.
	// Experimental.
	RepositoryArn() *string
	// The HTTPS (GRC) clone URL.
	//
	// HTTPS (GRC) is the protocol to use with git-remote-codecommit (GRC).
	//
	// It is the recommended method for supporting connections made with federated
	// access, identity providers, and temporary credentials.
	// Experimental.
	RepositoryCloneUrlGrc() *string
	// The HTTP clone URL.
	// Experimental.
	RepositoryCloneUrlHttp() *string
	// The SSH clone URL.
	// Experimental.
	RepositoryCloneUrlSsh() *string
	// The human-visible name of this Repository.
	// Experimental.
	RepositoryName() *string
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
	// Returns a source configuration for notification rule.
	// Experimental.
	BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig
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
	// Grant the given identity permissions to pull this repository.
	// Experimental.
	GrantPull(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to pull and push this repository.
	// Experimental.
	GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to read this repository.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Defines a CodeStar Notification rule which triggers when a pull request is merged.
	// Experimental.
	NotifiyOnPullRequestMerged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Create a trigger to notify another service to run actions on repository events.
	// Experimental.
	Notify(arn *string, options *RepositoryTriggerOptions) Repository
	// Defines a CodeStar Notification rule triggered when the project events specified by you are emitted. Similar to `onEvent` API.
	//
	// You can also use the methods to define rules for the specific event emitted.
	// eg: `notifyOnPullRequstCreated`.
	// Experimental.
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *RepositoryNotifyOnOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when an approval rule is overridden.
	// Experimental.
	NotifyOnApprovalRuleOverridden(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when an approval status is changed.
	// Experimental.
	NotifyOnApprovalStatusChanged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a new branch or tag is created.
	// Experimental.
	NotifyOnBranchOrTagCreated(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a branch or tag is deleted.
	// Experimental.
	NotifyOnBranchOrTagDeleted(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a comment is made on a pull request.
	// Experimental.
	NotifyOnPullRequestComment(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a pull request is created.
	// Experimental.
	NotifyOnPullRequestCreated(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a pull request is merged.
	// Experimental.
	NotifyOnPullRequestMerged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CloudWatch event rule which triggers when a comment is made on a commit.
	// Experimental.
	OnCommentOnCommit(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a comment is made on a pull request.
	// Experimental.
	OnCommentOnPullRequest(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a commit is pushed to a branch.
	// Experimental.
	OnCommit(id *string, options *OnCommitOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers for repository events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
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
	// Defines a CloudWatch event rule which triggers when a pull request state is changed.
	// Experimental.
	OnPullRequestStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a reference is created (i.e. a new branch/tag is created) to the repository.
	// Experimental.
	OnReferenceCreated(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a reference is delete (i.e. a branch/tag is deleted) from the repository.
	// Experimental.
	OnReferenceDeleted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a reference is updated (i.e. a commit is pushed to an existing or new branch) from the repository.
	// Experimental.
	OnReferenceUpdated(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a "CodeCommit Repository State Change" event occurs.
	// Experimental.
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
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

// The jsii proxy struct for Repository
type jsiiProxy_Repository struct {
	internal.Type__awscdkResource
	jsiiProxy_IRepository
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

func (j *jsiiProxy_Repository) RepositoryCloneUrlGrc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCloneUrlGrc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) RepositoryCloneUrlHttp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCloneUrlHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) RepositoryCloneUrlSsh() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCloneUrlSsh",
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
		"monocdk.aws_codecommit.Repository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRepository_Override(r Repository, scope constructs.Construct, id *string, props *RepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codecommit.Repository",
		[]interface{}{scope, id, props},
		r,
	)
}

// Imports a codecommit repository.
// Experimental.
func Repository_FromRepositoryArn(scope constructs.Construct, id *string, repositoryArn *string) IRepository {
	_init_.Initialize()

	var returns IRepository

	_jsii_.StaticInvoke(
		"monocdk.aws_codecommit.Repository",
		"fromRepositoryArn",
		[]interface{}{scope, id, repositoryArn},
		&returns,
	)

	return returns
}

// Experimental.
func Repository_FromRepositoryName(scope constructs.Construct, id *string, repositoryName *string) IRepository {
	_init_.Initialize()

	var returns IRepository

	_jsii_.StaticInvoke(
		"monocdk.aws_codecommit.Repository",
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
		"monocdk.aws_codecommit.Repository",
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
		"monocdk.aws_codecommit.Repository",
		"isResource",
		[]interface{}{construct},
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

func (r *jsiiProxy_Repository) BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig {
	var returns *awscodestarnotifications.NotificationRuleSourceConfig

	_jsii_.Invoke(
		r,
		"bindAsNotificationRuleSource",
		[]interface{}{_scope},
		&returns,
	)

	return returns
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

func (r *jsiiProxy_Repository) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) NotifiyOnPullRequestMerged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		r,
		"notifiyOnPullRequestMerged",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) Notify(arn *string, options *RepositoryTriggerOptions) Repository {
	var returns Repository

	_jsii_.Invoke(
		r,
		"notify",
		[]interface{}{arn, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *RepositoryNotifyOnOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		r,
		"notifyOn",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) NotifyOnApprovalRuleOverridden(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		r,
		"notifyOnApprovalRuleOverridden",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) NotifyOnApprovalStatusChanged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		r,
		"notifyOnApprovalStatusChanged",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) NotifyOnBranchOrTagCreated(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		r,
		"notifyOnBranchOrTagCreated",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) NotifyOnBranchOrTagDeleted(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		r,
		"notifyOnBranchOrTagDeleted",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) NotifyOnPullRequestComment(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		r,
		"notifyOnPullRequestComment",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) NotifyOnPullRequestCreated(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		r,
		"notifyOnPullRequestCreated",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) NotifyOnPullRequestMerged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		r,
		"notifyOnPullRequestMerged",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) OnCommentOnCommit(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onCommentOnCommit",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) OnCommentOnPullRequest(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onCommentOnPullRequest",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) OnCommit(id *string, options *OnCommitOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onCommit",
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

func (r *jsiiProxy_Repository) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) OnPullRequestStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onPullRequestStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) OnReferenceCreated(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onReferenceCreated",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) OnReferenceDeleted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onReferenceDeleted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) OnReferenceUpdated(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onReferenceUpdated",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
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

// Repository events that will cause the trigger to run actions in another service.
// Experimental.
type RepositoryEventTrigger string

const (
	// Experimental.
	RepositoryEventTrigger_ALL RepositoryEventTrigger = "ALL"
	// Experimental.
	RepositoryEventTrigger_UPDATE_REF RepositoryEventTrigger = "UPDATE_REF"
	// Experimental.
	RepositoryEventTrigger_CREATE_REF RepositoryEventTrigger = "CREATE_REF"
	// Experimental.
	RepositoryEventTrigger_DELETE_REF RepositoryEventTrigger = "DELETE_REF"
)

// List of event types for AWS CodeCommit.
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-repositories
//
// Experimental.
type RepositoryNotificationEvents string

const (
	// Trigger notication when comment made on commit.
	// Experimental.
	RepositoryNotificationEvents_COMMIT_COMMENT RepositoryNotificationEvents = "COMMIT_COMMENT"
	// Trigger notification when comment made on pull request.
	// Experimental.
	RepositoryNotificationEvents_PULL_REQUEST_COMMENT RepositoryNotificationEvents = "PULL_REQUEST_COMMENT"
	// Trigger notification when approval status changed.
	// Experimental.
	RepositoryNotificationEvents_APPROVAL_STATUS_CHANGED RepositoryNotificationEvents = "APPROVAL_STATUS_CHANGED"
	// Trigger notifications when approval rule is overridden.
	// Experimental.
	RepositoryNotificationEvents_APPROVAL_RULE_OVERRIDDEN RepositoryNotificationEvents = "APPROVAL_RULE_OVERRIDDEN"
	// Trigger notification when pull request created.
	// Experimental.
	RepositoryNotificationEvents_PULL_REQUEST_CREATED RepositoryNotificationEvents = "PULL_REQUEST_CREATED"
	// Trigger notification when pull request source updated.
	// Experimental.
	RepositoryNotificationEvents_PULL_REQUEST_SOURCE_UPDATED RepositoryNotificationEvents = "PULL_REQUEST_SOURCE_UPDATED"
	// Trigger notification when pull request status is changed.
	// Experimental.
	RepositoryNotificationEvents_PULL_REQUEST_STATUS_CHANGED RepositoryNotificationEvents = "PULL_REQUEST_STATUS_CHANGED"
	// Trigger notification when pull requset is merged.
	// Experimental.
	RepositoryNotificationEvents_PULL_REQUEST_MERGED RepositoryNotificationEvents = "PULL_REQUEST_MERGED"
	// Trigger notification when a branch or tag is created.
	// Experimental.
	RepositoryNotificationEvents_BRANCH_OR_TAG_CREATED RepositoryNotificationEvents = "BRANCH_OR_TAG_CREATED"
	// Trigger notification when a branch or tag is deleted.
	// Experimental.
	RepositoryNotificationEvents_BRANCH_OR_TAG_DELETED RepositoryNotificationEvents = "BRANCH_OR_TAG_DELETED"
	// Trigger notification when a branch or tag is updated.
	// Experimental.
	RepositoryNotificationEvents_BRANCH_OR_TAG_UPDATED RepositoryNotificationEvents = "BRANCH_OR_TAG_UPDATED"
)

// Additional options to pass to the notification rule.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codecommit "github.com/aws/aws-cdk-go/awscdk/aws_codecommit"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codestarnotifications "github.com/aws/aws-cdk-go/awscdk/aws_codestarnotifications"
//   repositoryNotifyOnOptions := &repositoryNotifyOnOptions{
//   	events: []repositoryNotificationEvents{
//   		codecommit.*repositoryNotificationEvents_COMMIT_COMMENT,
//   	},
//
//   	// the properties below are optional
//   	detailType: codestarnotifications.detailType_BASIC,
//   	enabled: jsii.Boolean(false),
//   	notificationRuleName: jsii.String("notificationRuleName"),
//   }
//
// Experimental.
type RepositoryNotifyOnOptions struct {
	// The level of detail to include in the notifications for this resource.
	//
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	// Experimental.
	DetailType awscodestarnotifications.DetailType `json:"detailType" yaml:"detailType"`
	// The status of the notification rule.
	//
	// If the enabled is set to DISABLED, notifications aren't sent for the notification rule.
	// Experimental.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account.
	// Experimental.
	NotificationRuleName *string `json:"notificationRuleName" yaml:"notificationRuleName"`
	// A list of event types associated with this notification rule for CodeCommit repositories.
	//
	// For a complete list of event types and IDs, see Notification concepts in the Developer Tools Console User Guide.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#concepts-api
	//
	// Experimental.
	Events *[]RepositoryNotificationEvents `json:"events" yaml:"events"`
}

// Example:
//   // Source stage: read from repository
//   repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &repositoryProps{
//   	repositoryName: jsii.String("template-repo"),
//   })
//   sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
//   source := cpactions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("Source"),
//   	repository: repo,
//   	output: sourceOutput,
//   	trigger: cpactions.codeCommitTrigger_POLL,
//   })
//   sourceStage := map[string]interface{}{
//   	"stageName": jsii.String("Source"),
//   	"actions": []CodeCommitSourceAction{
//   		source,
//   	},
//   }
//
//   // Deployment stage: create and deploy changeset with manual approval
//   stackName := "OurStack"
//   changeSetName := "StagedChangeSet"
//
//   prodStage := map[string]interface{}{
//   	"stageName": jsii.String("Deploy"),
//   	"actions": []interface{}{
//   		cpactions.NewCloudFormationCreateReplaceChangeSetAction(&CloudFormationCreateReplaceChangeSetActionProps{
//   			"actionName": jsii.String("PrepareChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"adminPermissions": jsii.Boolean(true),
//   			"templatePath": sourceOutput.atPath(jsii.String("template.yaml")),
//   			"runOrder": jsii.Number(1),
//   		}),
//   		cpactions.NewManualApprovalAction(&ManualApprovalActionProps{
//   			"actionName": jsii.String("ApproveChanges"),
//   			"runOrder": jsii.Number(2),
//   		}),
//   		cpactions.NewCloudFormationExecuteChangeSetAction(&CloudFormationExecuteChangeSetActionProps{
//   			"actionName": jsii.String("ExecuteChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"runOrder": jsii.Number(3),
//   		}),
//   	},
//   }
//
//   codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
// Experimental.
type RepositoryProps struct {
	// Name of the repository.
	//
	// This property is required for all CodeCommit repositories.
	// Experimental.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
	// The contents with which to initialize the repository after it has been created.
	// Experimental.
	Code Code `json:"code" yaml:"code"`
	// A description of the repository.
	//
	// Use the description to identify the
	// purpose of the repository.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
}

// Creates for a repository trigger to an SNS topic or Lambda function.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codecommit "github.com/aws/aws-cdk-go/awscdk/aws_codecommit"
//   repositoryTriggerOptions := &repositoryTriggerOptions{
//   	branches: []*string{
//   		jsii.String("branches"),
//   	},
//   	customData: jsii.String("customData"),
//   	events: []repositoryEventTrigger{
//   		codecommit.*repositoryEventTrigger_ALL,
//   	},
//   	name: jsii.String("name"),
//   }
//
// Experimental.
type RepositoryTriggerOptions struct {
	// The names of the branches in the AWS CodeCommit repository that contain events that you want to include in the trigger.
	//
	// If you don't specify at
	// least one branch, the trigger applies to all branches.
	// Experimental.
	Branches *[]*string `json:"branches" yaml:"branches"`
	// When an event is triggered, additional information that AWS CodeCommit includes when it sends information to the target.
	// Experimental.
	CustomData *string `json:"customData" yaml:"customData"`
	// The repository events for which AWS CodeCommit sends information to the target, which you specified in the DestinationArn property.If you don't specify events, the trigger runs for all repository events.
	// Experimental.
	Events *[]RepositoryEventTrigger `json:"events" yaml:"events"`
	// A name for the trigger.Triggers on a repository must have unique names.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
}

