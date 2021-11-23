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
// TODO: EXAMPLE
//
type CfnSlackChannelConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConfigurationName() *string
	SetConfigurationName(val *string)
	CreationStack() *[]*string
	GuardrailPolicies() *[]*string
	SetGuardrailPolicies(val *[]*string)
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	LoggingLevel() *string
	SetLoggingLevel(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	SlackChannelId() *string
	SetSlackChannelId(val *string)
	SlackWorkspaceId() *string
	SetSlackWorkspaceId(val *string)
	SnsTopicArns() *[]*string
	SetSnsTopicArns(val *[]*string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	UserRoleRequired() interface{}
	SetUserRoleRequired(val interface{})
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnSlackChannelConfiguration) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnSlackChannelConfiguration) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnSlackChannelConfiguration) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnSlackChannelConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnSlackChannelConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnSlackChannelConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnSlackChannelConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
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

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnSlackChannelConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnSlackChannelConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnSlackChannelConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnSlackChannelConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
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

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnSlackChannelConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnSlackChannelConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::Chatbot::SlackChannelConfiguration`.
//
// TODO: EXAMPLE
//
type CfnSlackChannelConfigurationProps struct {
	// `AWS::Chatbot::SlackChannelConfiguration.ConfigurationName`.
	ConfigurationName *string `json:"configurationName"`
	// `AWS::Chatbot::SlackChannelConfiguration.GuardrailPolicies`.
	GuardrailPolicies *[]*string `json:"guardrailPolicies"`
	// `AWS::Chatbot::SlackChannelConfiguration.IamRoleArn`.
	IamRoleArn *string `json:"iamRoleArn"`
	// `AWS::Chatbot::SlackChannelConfiguration.LoggingLevel`.
	LoggingLevel *string `json:"loggingLevel"`
	// `AWS::Chatbot::SlackChannelConfiguration.SlackChannelId`.
	SlackChannelId *string `json:"slackChannelId"`
	// `AWS::Chatbot::SlackChannelConfiguration.SlackWorkspaceId`.
	SlackWorkspaceId *string `json:"slackWorkspaceId"`
	// `AWS::Chatbot::SlackChannelConfiguration.SnsTopicArns`.
	SnsTopicArns *[]*string `json:"snsTopicArns"`
	// `AWS::Chatbot::SlackChannelConfiguration.UserRoleRequired`.
	UserRoleRequired interface{} `json:"userRoleRequired"`
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
	LoggingLevel_ERROR LoggingLevel = "ERROR"
	LoggingLevel_INFO LoggingLevel = "INFO"
	LoggingLevel_NONE LoggingLevel = "NONE"
)

// A new Slack channel configuration.
//
// TODO: EXAMPLE
//
// Experimental.
type SlackChannelConfiguration interface {
	awscdk.Resource
	ISlackChannelConfiguration
	Env() *awscdk.ResourceEnvironment
	GrantPrincipal() awsiam.IPrincipal
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Role() awsiam.IRole
	SlackChannelConfigurationArn() *string
	SlackChannelConfigurationName() *string
	Stack() awscdk.Stack
	AddNotificationTopic(notificationTopic awssns.ITopic)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
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
// Returns: a reference to the existing Slack channel configuration
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

// Adds a SNS topic that deliver notifications to AWS Chatbot.
// Experimental.
func (s *jsiiProxy_SlackChannelConfiguration) AddNotificationTopic(notificationTopic awssns.ITopic) {
	_jsii_.InvokeVoid(
		s,
		"addNotificationTopic",
		[]interface{}{notificationTopic},
	)
}

// Adds extra permission to iam-role of Slack channel configuration.
// Experimental.
func (s *jsiiProxy_SlackChannelConfiguration) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		s,
		"addToRolePolicy",
		[]interface{}{statement},
	)
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
// Experimental.
func (s *jsiiProxy_SlackChannelConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Returns a target configuration for notification rule.
// Experimental.
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

// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
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

// Return the given named metric for this SlackChannelConfiguration.
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SlackChannelConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SlackChannelConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SlackChannelConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SlackChannelConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type SlackChannelConfigurationProps struct {
	// The name of Slack channel configuration.
	// Experimental.
	SlackChannelConfigurationName *string `json:"slackChannelConfigurationName"`
	// The ID of the Slack channel.
	//
	// To get the ID, open Slack, right click on the channel name in the left pane, then choose Copy Link.
	// The channel ID is the 9-character string at the end of the URL. For example, ABCBBLZZZ.
	// Experimental.
	SlackChannelId *string `json:"slackChannelId"`
	// The ID of the Slack workspace authorized with AWS Chatbot.
	//
	// To get the workspace ID, you must perform the initial authorization flow with Slack in the AWS Chatbot console.
	// Then you can copy and paste the workspace ID from the console.
	// For more details, see steps 1-4 in Setting Up AWS Chatbot with Slack in the AWS Chatbot User Guide.
	// See: https://docs.aws.amazon.com/chatbot/latest/adminguide/setting-up.html#Setup_intro
	//
	// Experimental.
	SlackWorkspaceId *string `json:"slackWorkspaceId"`
	// Specifies the logging level for this configuration.
	//
	// This property affects the log entries pushed to Amazon CloudWatch Logs.
	// Experimental.
	LoggingLevel LoggingLevel `json:"loggingLevel"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `json:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	// Experimental.
	LogRetentionRetryOptions *awslogs.LogRetentionRetryOptions `json:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	LogRetentionRole awsiam.IRole `json:"logRetentionRole"`
	// The SNS topics that deliver notifications to AWS Chatbot.
	// Experimental.
	NotificationTopics *[]awssns.ITopic `json:"notificationTopics"`
	// The permission role of Slack channel configuration.
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

