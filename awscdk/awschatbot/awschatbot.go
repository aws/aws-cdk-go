package awschatbot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awschatbot/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::Chatbot::SlackChannelConfiguration`.
//
// The `AWS::Chatbot::SlackChannelConfiguration` resource configures a Slack channel to allow users to use AWS Chatbot with AWS CloudFormation templates.
//
// This resource requires some setup to be done in the AWS Chatbot console. To provide the required Slack workspace ID, you must perform the initial authorization flow with Slack in the AWS Chatbot console, then copy and paste the workspace ID from the console. For more details, see steps 1-4 in [Setting Up AWS Chatbot with Slack](https://docs.aws.amazon.com/chatbot/latest/adminguide/setting-up.html#Setup_intro) in the *AWS Chatbot User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSlackChannelConfiguration := awscdk.Aws_chatbot.NewCfnSlackChannelConfiguration(this, jsii.String("MyCfnSlackChannelConfiguration"), &cfnSlackChannelConfigurationProps{
//   	configurationName: jsii.String("configurationName"),
//   	iamRoleArn: jsii.String("iamRoleArn"),
//   	slackChannelId: jsii.String("slackChannelId"),
//   	slackWorkspaceId: jsii.String("slackWorkspaceId"),
//
//   	// the properties below are optional
//   	guardrailPolicies: []*string{
//   		jsii.String("guardrailPolicies"),
//   	},
//   	loggingLevel: jsii.String("loggingLevel"),
//   	snsTopicArns: []*string{
//   		jsii.String("snsTopicArns"),
//   	},
//   	userRoleRequired: jsii.Boolean(false),
//   })
//
type CfnSlackChannelConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The name of the configuration.
	ConfigurationName() *string
	SetConfigurationName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The list of IAM policy ARNs that are applied as channel guardrails.
	//
	// The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	GuardrailPolicies() *[]*string
	SetGuardrailPolicies(val *[]*string)
	// The ARN of the IAM role that defines the permissions for AWS Chatbot .
	//
	// This is a user-definworked role that AWS Chatbot will assume. This is not the service-linked role. For more information, see [IAM Policies for AWS Chatbot](https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html) .
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	// Specifies the logging level for this configuration. This property affects the log entries pushed to Amazon CloudWatch Logs.
	//
	// Logging levels include `ERROR` , `INFO` , or `NONE` .
	LoggingLevel() *string
	SetLoggingLevel(val *string)
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
	// The ID of the Slack channel.
	//
	// To get the ID, open Slack, right click on the channel name in the left pane, then choose Copy Link. The channel ID is the 9-character string at the end of the URL. For example, `ABCBBLZZZ` .
	SlackChannelId() *string
	SetSlackChannelId(val *string)
	// The ID of the Slack workspace authorized with AWS Chatbot .
	//
	// To get the workspace ID, you must perform the initial authorization flow with Slack in the AWS Chatbot console. Then you can copy and paste the workspace ID from the console. For more details, see steps 1-4 in [Setting Up AWS Chatbot with Slack](https://docs.aws.amazon.com/chatbot/latest/adminguide/setting-up.html#Setup_intro) in the *AWS Chatbot User Guide* .
	SlackWorkspaceId() *string
	SetSlackWorkspaceId(val *string)
	// The ARNs of the SNS topics that deliver notifications to AWS Chatbot .
	SnsTopicArns() *[]*string
	SetSnsTopicArns(val *[]*string)
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
	// Enables use of a user role requirement in your chat configuration.
	UserRoleRequired() interface{}
	SetUserRoleRequired(val interface{})
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

// The jsii proxy struct for CfnSlackChannelConfiguration
type jsiiProxy_CfnSlackChannelConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) ConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) GuardrailPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrailPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) LoggingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) SlackChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) SlackWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) SnsTopicArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snsTopicArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) UserRoleRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userRoleRequired",
		&returns,
	)
	return returns
}


// Create a new `AWS::Chatbot::SlackChannelConfiguration`.
func NewCfnSlackChannelConfiguration(scope awscdk.Construct, id *string, props *CfnSlackChannelConfigurationProps) CfnSlackChannelConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnSlackChannelConfiguration{}

	_jsii_.Create(
		"monocdk.aws_chatbot.CfnSlackChannelConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Chatbot::SlackChannelConfiguration`.
func NewCfnSlackChannelConfiguration_Override(c CfnSlackChannelConfiguration, scope awscdk.Construct, id *string, props *CfnSlackChannelConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_chatbot.CfnSlackChannelConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) SetConfigurationName(val *string) {
	_jsii_.Set(
		j,
		"configurationName",
		val,
	)
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) SetGuardrailPolicies(val *[]*string) {
	_jsii_.Set(
		j,
		"guardrailPolicies",
		val,
	)
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) SetIamRoleArn(val *string) {
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) SetLoggingLevel(val *string) {
	_jsii_.Set(
		j,
		"loggingLevel",
		val,
	)
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) SetSlackChannelId(val *string) {
	_jsii_.Set(
		j,
		"slackChannelId",
		val,
	)
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) SetSlackWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"slackWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) SetSnsTopicArns(val *[]*string) {
	_jsii_.Set(
		j,
		"snsTopicArns",
		val,
	)
}

func (j *jsiiProxy_CfnSlackChannelConfiguration) SetUserRoleRequired(val interface{}) {
	_jsii_.Set(
		j,
		"userRoleRequired",
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
func CfnSlackChannelConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_chatbot.CfnSlackChannelConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnSlackChannelConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_chatbot.CfnSlackChannelConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnSlackChannelConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_chatbot.CfnSlackChannelConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSlackChannelConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_chatbot.CfnSlackChannelConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSlackChannelConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnSlackChannelConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSlackChannelConfigurationProps := &cfnSlackChannelConfigurationProps{
//   	configurationName: jsii.String("configurationName"),
//   	iamRoleArn: jsii.String("iamRoleArn"),
//   	slackChannelId: jsii.String("slackChannelId"),
//   	slackWorkspaceId: jsii.String("slackWorkspaceId"),
//
//   	// the properties below are optional
//   	guardrailPolicies: []*string{
//   		jsii.String("guardrailPolicies"),
//   	},
//   	loggingLevel: jsii.String("loggingLevel"),
//   	snsTopicArns: []*string{
//   		jsii.String("snsTopicArns"),
//   	},
//   	userRoleRequired: jsii.Boolean(false),
//   }
//
type CfnSlackChannelConfigurationProps struct {
	// The name of the configuration.
	ConfigurationName *string `field:"required" json:"configurationName" yaml:"configurationName"`
	// The ARN of the IAM role that defines the permissions for AWS Chatbot .
	//
	// This is a user-definworked role that AWS Chatbot will assume. This is not the service-linked role. For more information, see [IAM Policies for AWS Chatbot](https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html) .
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The ID of the Slack channel.
	//
	// To get the ID, open Slack, right click on the channel name in the left pane, then choose Copy Link. The channel ID is the 9-character string at the end of the URL. For example, `ABCBBLZZZ` .
	SlackChannelId *string `field:"required" json:"slackChannelId" yaml:"slackChannelId"`
	// The ID of the Slack workspace authorized with AWS Chatbot .
	//
	// To get the workspace ID, you must perform the initial authorization flow with Slack in the AWS Chatbot console. Then you can copy and paste the workspace ID from the console. For more details, see steps 1-4 in [Setting Up AWS Chatbot with Slack](https://docs.aws.amazon.com/chatbot/latest/adminguide/setting-up.html#Setup_intro) in the *AWS Chatbot User Guide* .
	SlackWorkspaceId *string `field:"required" json:"slackWorkspaceId" yaml:"slackWorkspaceId"`
	// The list of IAM policy ARNs that are applied as channel guardrails.
	//
	// The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	GuardrailPolicies *[]*string `field:"optional" json:"guardrailPolicies" yaml:"guardrailPolicies"`
	// Specifies the logging level for this configuration. This property affects the log entries pushed to Amazon CloudWatch Logs.
	//
	// Logging levels include `ERROR` , `INFO` , or `NONE` .
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// The ARNs of the SNS topics that deliver notifications to AWS Chatbot .
	SnsTopicArns *[]*string `field:"optional" json:"snsTopicArns" yaml:"snsTopicArns"`
	// Enables use of a user role requirement in your chat configuration.
	UserRoleRequired interface{} `field:"optional" json:"userRoleRequired" yaml:"userRoleRequired"`
}

// Represents a Slack channel configuration.
// Experimental.
type ISlackChannelConfiguration interface {
	awsiam.IGrantable
	awscodestarnotifications.INotificationRuleTarget
	awscdk.IResource
	// Adds a statement to the IAM role.
	// Experimental.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Return the given named metric for this SlackChannelConfiguration.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The permission role of Slack channel configuration.
	// Experimental.
	Role() awsiam.IRole
	// The ARN of the Slack channel configuration In the form of arn:aws:chatbot:{region}:{account}:chat-configuration/slack-channel/{slackChannelName}.
	// Experimental.
	SlackChannelConfigurationArn() *string
	// The name of Slack channel configuration.
	// Experimental.
	SlackChannelConfigurationName() *string
}

// The jsii proxy for ISlackChannelConfiguration
type jsiiProxy_ISlackChannelConfiguration struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscodestarnotificationsINotificationRuleTarget
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ISlackChannelConfiguration) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		i,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (i *jsiiProxy_ISlackChannelConfiguration) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISlackChannelConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ISlackChannelConfiguration) BindAsNotificationRuleTarget(scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig {
	var returns *awscodestarnotifications.NotificationRuleTargetConfig

	_jsii_.Invoke(
		i,
		"bindAsNotificationRuleTarget",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) SlackChannelConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) SlackChannelConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Logging levels include ERROR, INFO, or NONE.
// Experimental.
type LoggingLevel string

const (
	// ERROR.
	// Experimental.
	LoggingLevel_ERROR LoggingLevel = "ERROR"
	// INFO.
	// Experimental.
	LoggingLevel_INFO LoggingLevel = "INFO"
	// NONE.
	// Experimental.
	LoggingLevel_NONE LoggingLevel = "NONE"
)

// A new Slack channel configuration.
//
// Example:
//   import chatbot "github.com/aws/aws-cdk-go/awscdk"
//
//   var project project
//
//
//   target := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &slackChannelConfigurationProps{
//   	slackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
//   	slackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
//   	slackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
//   })
//
//   rule := project.notifyOnBuildSucceeded(jsii.String("NotifyOnBuildSucceeded"), target)
//
// Experimental.
type SlackChannelConfiguration interface {
	awscdk.Resource
	ISlackChannelConfiguration
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
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
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
	// The permission role of Slack channel configuration.
	// Experimental.
	Role() awsiam.IRole
	// The ARN of the Slack channel configuration In the form of arn:aws:chatbot:{region}:{account}:chat-configuration/slack-channel/{slackChannelName}.
	// Experimental.
	SlackChannelConfigurationArn() *string
	// The name of Slack channel configuration.
	// Experimental.
	SlackChannelConfigurationName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Adds a SNS topic that deliver notifications to AWS Chatbot.
	// Experimental.
	AddNotificationTopic(notificationTopic awssns.ITopic)
	// Adds extra permission to iam-role of Slack channel configuration.
	// Experimental.
	AddToRolePolicy(statement awsiam.PolicyStatement)
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
	// Returns a target configuration for notification rule.
	// Experimental.
	BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig
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
	// Return the given named metric for this SlackChannelConfiguration.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
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

// The jsii proxy struct for SlackChannelConfiguration
type jsiiProxy_SlackChannelConfiguration struct {
	internal.Type__awscdkResource
	jsiiProxy_ISlackChannelConfiguration
}

func (j *jsiiProxy_SlackChannelConfiguration) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) SlackChannelConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) SlackChannelConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewSlackChannelConfiguration(scope constructs.Construct, id *string, props *SlackChannelConfigurationProps) SlackChannelConfiguration {
	_init_.Initialize()

	j := jsiiProxy_SlackChannelConfiguration{}

	_jsii_.Create(
		"monocdk.aws_chatbot.SlackChannelConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSlackChannelConfiguration_Override(s SlackChannelConfiguration, scope constructs.Construct, id *string, props *SlackChannelConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_chatbot.SlackChannelConfiguration",
		[]interface{}{scope, id, props},
		s,
	)
}

// Import an existing Slack channel configuration provided an ARN.
//
// Returns: a reference to the existing Slack channel configuration.
// Experimental.
func SlackChannelConfiguration_FromSlackChannelConfigurationArn(scope constructs.Construct, id *string, slackChannelConfigurationArn *string) ISlackChannelConfiguration {
	_init_.Initialize()

	var returns ISlackChannelConfiguration

	_jsii_.StaticInvoke(
		"monocdk.aws_chatbot.SlackChannelConfiguration",
		"fromSlackChannelConfigurationArn",
		[]interface{}{scope, id, slackChannelConfigurationArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func SlackChannelConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_chatbot.SlackChannelConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func SlackChannelConfiguration_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_chatbot.SlackChannelConfiguration",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for All SlackChannelConfigurations.
// Experimental.
func SlackChannelConfiguration_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_chatbot.SlackChannelConfiguration",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) AddNotificationTopic(notificationTopic awssns.ITopic) {
	_jsii_.InvokeVoid(
		s,
		"addNotificationTopic",
		[]interface{}{notificationTopic},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		s,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig {
	var returns *awscodestarnotifications.NotificationRuleTargetConfig

	_jsii_.Invoke(
		s,
		"bindAsNotificationRuleTarget",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a new Slack channel configuration.
//
// Example:
//   import chatbot "github.com/aws/aws-cdk-go/awscdk"
//
//   var project project
//
//
//   target := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &slackChannelConfigurationProps{
//   	slackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
//   	slackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
//   	slackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
//   })
//
//   rule := project.notifyOnBuildSucceeded(jsii.String("NotifyOnBuildSucceeded"), target)
//
// Experimental.
type SlackChannelConfigurationProps struct {
	// The name of Slack channel configuration.
	// Experimental.
	SlackChannelConfigurationName *string `field:"required" json:"slackChannelConfigurationName" yaml:"slackChannelConfigurationName"`
	// The ID of the Slack channel.
	//
	// To get the ID, open Slack, right click on the channel name in the left pane, then choose Copy Link.
	// The channel ID is the 9-character string at the end of the URL. For example, ABCBBLZZZ.
	// Experimental.
	SlackChannelId *string `field:"required" json:"slackChannelId" yaml:"slackChannelId"`
	// The ID of the Slack workspace authorized with AWS Chatbot.
	//
	// To get the workspace ID, you must perform the initial authorization flow with Slack in the AWS Chatbot console.
	// Then you can copy and paste the workspace ID from the console.
	// For more details, see steps 1-4 in Setting Up AWS Chatbot with Slack in the AWS Chatbot User Guide.
	// See: https://docs.aws.amazon.com/chatbot/latest/adminguide/setting-up.html#Setup_intro
	//
	// Experimental.
	SlackWorkspaceId *string `field:"required" json:"slackWorkspaceId" yaml:"slackWorkspaceId"`
	// Specifies the logging level for this configuration.
	//
	// This property affects the log entries pushed to Amazon CloudWatch Logs.
	// Experimental.
	LoggingLevel LoggingLevel `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	// Experimental.
	LogRetentionRetryOptions *awslogs.LogRetentionRetryOptions `field:"optional" json:"logRetentionRetryOptions" yaml:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	LogRetentionRole awsiam.IRole `field:"optional" json:"logRetentionRole" yaml:"logRetentionRole"`
	// The SNS topics that deliver notifications to AWS Chatbot.
	// Experimental.
	NotificationTopics *[]awssns.ITopic `field:"optional" json:"notificationTopics" yaml:"notificationTopics"`
	// The permission role of Slack channel configuration.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

