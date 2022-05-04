package awscodestarnotifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::CodeStarNotifications::NotificationRule`.
//
// Creates a notification rule for a resource. The rule specifies the events you want notifications about and the targets (such as AWS Chatbot topics or AWS Chatbot clients configured for Slack) where you want to receive them.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codestarnotifications "github.com/aws/aws-cdk-go/awscdk/aws_codestarnotifications"
//
//   var tags interface{}
//   cfnNotificationRule := codestarnotifications.NewCfnNotificationRule(this, jsii.String("MyCfnNotificationRule"), &cfnNotificationRuleProps{
//   	detailType: jsii.String("detailType"),
//   	eventTypeIds: []*string{
//   		jsii.String("eventTypeIds"),
//   	},
//   	name: jsii.String("name"),
//   	resource: jsii.String("resource"),
//   	targets: []interface{}{
//   		&targetProperty{
//   			targetAddress: jsii.String("targetAddress"),
//   			targetType: jsii.String("targetType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	createdBy: jsii.String("createdBy"),
//   	eventTypeId: jsii.String("eventTypeId"),
//   	status: jsii.String("status"),
//   	tags: tags,
//   	targetAddress: jsii.String("targetAddress"),
//   })
//
type CfnNotificationRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// `AWS::CodeStarNotifications::NotificationRule.CreatedBy`.
	CreatedBy() *string
	SetCreatedBy(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The level of detail to include in the notifications for this resource.
	//
	// `BASIC` will include only the contents of the event as it would appear in Amazon CloudWatch. `FULL` will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	DetailType() *string
	SetDetailType(val *string)
	// `AWS::CodeStarNotifications::NotificationRule.EventTypeId`.
	EventTypeId() *string
	SetEventTypeId(val *string)
	// A list of event types associated with this notification rule.
	//
	// For a complete list of event types and IDs, see [Notification concepts](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api) in the *Developer Tools Console User Guide* .
	EventTypeIds() *[]*string
	SetEventTypeIds(val *[]*string)
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
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account .
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the resource to associate with the notification rule.
	//
	// Supported resources include pipelines in AWS CodePipeline , repositories in AWS CodeCommit , and build projects in AWS CodeBuild .
	Resource() *string
	SetResource(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The status of the notification rule.
	//
	// The default value is `ENABLED` . If the status is set to `DISABLED` , notifications aren't sent for the notification rule.
	Status() *string
	SetStatus(val *string)
	// A list of tags to apply to this notification rule.
	//
	// Key names cannot start with " `aws` ".
	Tags() awscdk.TagManager
	// `AWS::CodeStarNotifications::NotificationRule.TargetAddress`.
	TargetAddress() *string
	SetTargetAddress(val *string)
	// A list of Amazon Resource Names (ARNs) of AWS Chatbot topics and AWS Chatbot clients to associate with the notification rule.
	Targets() interface{}
	SetTargets(val interface{})
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

// The jsii proxy struct for CfnNotificationRule
type jsiiProxy_CfnNotificationRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnNotificationRule) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) DetailType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detailType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) EventTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) EventTypeIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventTypeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) TargetAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CodeStarNotifications::NotificationRule`.
func NewCfnNotificationRule(scope constructs.Construct, id *string, props *CfnNotificationRuleProps) CfnNotificationRule {
	_init_.Initialize()

	j := jsiiProxy_CfnNotificationRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codestarnotifications.CfnNotificationRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeStarNotifications::NotificationRule`.
func NewCfnNotificationRule_Override(c CfnNotificationRule, scope constructs.Construct, id *string, props *CfnNotificationRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codestarnotifications.CfnNotificationRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnNotificationRule) SetCreatedBy(val *string) {
	_jsii_.Set(
		j,
		"createdBy",
		val,
	)
}

func (j *jsiiProxy_CfnNotificationRule) SetDetailType(val *string) {
	_jsii_.Set(
		j,
		"detailType",
		val,
	)
}

func (j *jsiiProxy_CfnNotificationRule) SetEventTypeId(val *string) {
	_jsii_.Set(
		j,
		"eventTypeId",
		val,
	)
}

func (j *jsiiProxy_CfnNotificationRule) SetEventTypeIds(val *[]*string) {
	_jsii_.Set(
		j,
		"eventTypeIds",
		val,
	)
}

func (j *jsiiProxy_CfnNotificationRule) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnNotificationRule) SetResource(val *string) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_CfnNotificationRule) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CfnNotificationRule) SetTargetAddress(val *string) {
	_jsii_.Set(
		j,
		"targetAddress",
		val,
	)
}

func (j *jsiiProxy_CfnNotificationRule) SetTargets(val interface{}) {
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnNotificationRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codestarnotifications.CfnNotificationRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnNotificationRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codestarnotifications.CfnNotificationRule",
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
func CfnNotificationRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codestarnotifications.CfnNotificationRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNotificationRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codestarnotifications.CfnNotificationRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNotificationRule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnNotificationRule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnNotificationRule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnNotificationRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnNotificationRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnNotificationRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnNotificationRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnNotificationRule) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotificationRule) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotificationRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnNotificationRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnNotificationRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotificationRule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotificationRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNotificationRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information about the AWS Chatbot topics or AWS Chatbot clients associated with a notification rule.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codestarnotifications "github.com/aws/aws-cdk-go/awscdk/aws_codestarnotifications"
//   targetProperty := &targetProperty{
//   	targetAddress: jsii.String("targetAddress"),
//   	targetType: jsii.String("targetType"),
//   }
//
type CfnNotificationRule_TargetProperty struct {
	// The Amazon Resource Name (ARN) of the AWS Chatbot topic or AWS Chatbot client.
	TargetAddress *string `json:"targetAddress" yaml:"targetAddress"`
	// The target type. Can be an Amazon Simple Notification Service topic or AWS Chatbot client.
	//
	// - Amazon Simple Notification Service topics are specified as `SNS` .
	// - AWS Chatbot clients are specified as `AWSChatbotSlack` .
	TargetType *string `json:"targetType" yaml:"targetType"`
}

// Properties for defining a `CfnNotificationRule`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codestarnotifications "github.com/aws/aws-cdk-go/awscdk/aws_codestarnotifications"
//
//   var tags interface{}
//   cfnNotificationRuleProps := &cfnNotificationRuleProps{
//   	detailType: jsii.String("detailType"),
//   	eventTypeIds: []*string{
//   		jsii.String("eventTypeIds"),
//   	},
//   	name: jsii.String("name"),
//   	resource: jsii.String("resource"),
//   	targets: []interface{}{
//   		&targetProperty{
//   			targetAddress: jsii.String("targetAddress"),
//   			targetType: jsii.String("targetType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	createdBy: jsii.String("createdBy"),
//   	eventTypeId: jsii.String("eventTypeId"),
//   	status: jsii.String("status"),
//   	tags: tags,
//   	targetAddress: jsii.String("targetAddress"),
//   }
//
type CfnNotificationRuleProps struct {
	// The level of detail to include in the notifications for this resource.
	//
	// `BASIC` will include only the contents of the event as it would appear in Amazon CloudWatch. `FULL` will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	DetailType *string `json:"detailType" yaml:"detailType"`
	// A list of event types associated with this notification rule.
	//
	// For a complete list of event types and IDs, see [Notification concepts](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api) in the *Developer Tools Console User Guide* .
	EventTypeIds *[]*string `json:"eventTypeIds" yaml:"eventTypeIds"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account .
	Name *string `json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the resource to associate with the notification rule.
	//
	// Supported resources include pipelines in AWS CodePipeline , repositories in AWS CodeCommit , and build projects in AWS CodeBuild .
	Resource *string `json:"resource" yaml:"resource"`
	// A list of Amazon Resource Names (ARNs) of AWS Chatbot topics and AWS Chatbot clients to associate with the notification rule.
	Targets interface{} `json:"targets" yaml:"targets"`
	// `AWS::CodeStarNotifications::NotificationRule.CreatedBy`.
	CreatedBy *string `json:"createdBy" yaml:"createdBy"`
	// `AWS::CodeStarNotifications::NotificationRule.EventTypeId`.
	EventTypeId *string `json:"eventTypeId" yaml:"eventTypeId"`
	// The status of the notification rule.
	//
	// The default value is `ENABLED` . If the status is set to `DISABLED` , notifications aren't sent for the notification rule.
	Status *string `json:"status" yaml:"status"`
	// A list of tags to apply to this notification rule.
	//
	// Key names cannot start with " `aws` ".
	Tags interface{} `json:"tags" yaml:"tags"`
	// `AWS::CodeStarNotifications::NotificationRule.TargetAddress`.
	TargetAddress *string `json:"targetAddress" yaml:"targetAddress"`
}

// The level of detail to include in the notifications for this resource.
type DetailType string

const (
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	DetailType_BASIC DetailType = "BASIC"
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	DetailType_FULL DetailType = "FULL"
)

// Represents a notification rule.
type INotificationRule interface {
	awscdk.IResource
	// Adds target to notification rule.
	//
	// Returns: boolean - return true if it had any effect.
	AddTarget(target INotificationRuleTarget) *bool
	// The ARN of the notification rule (i.e. arn:aws:codestar-notifications:::notificationrule/01234abcde).
	NotificationRuleArn() *string
}

// The jsii proxy for INotificationRule
type jsiiProxy_INotificationRule struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_INotificationRule) AddTarget(target INotificationRuleTarget) *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"addTarget",
		[]interface{}{target},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_INotificationRule) NotificationRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationRuleArn",
		&returns,
	)
	return returns
}

// Represents a notification source The source that allows CodeBuild and CodePipeline to associate with this rule.
type INotificationRuleSource interface {
	// Returns a source configuration for notification rule.
	BindAsNotificationRuleSource(scope constructs.Construct) *NotificationRuleSourceConfig
}

// The jsii proxy for INotificationRuleSource
type jsiiProxy_INotificationRuleSource struct {
	_ byte // padding
}

func (i *jsiiProxy_INotificationRuleSource) BindAsNotificationRuleSource(scope constructs.Construct) *NotificationRuleSourceConfig {
	var returns *NotificationRuleSourceConfig

	_jsii_.Invoke(
		i,
		"bindAsNotificationRuleSource",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Represents a notification target That allows AWS Chatbot and SNS topic to associate with this rule target.
type INotificationRuleTarget interface {
	// Returns a target configuration for notification rule.
	BindAsNotificationRuleTarget(scope constructs.Construct) *NotificationRuleTargetConfig
}

// The jsii proxy for INotificationRuleTarget
type jsiiProxy_INotificationRuleTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_INotificationRuleTarget) BindAsNotificationRuleTarget(scope constructs.Construct) *NotificationRuleTargetConfig {
	var returns *NotificationRuleTargetConfig

	_jsii_.Invoke(
		i,
		"bindAsNotificationRuleTarget",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// A new notification rule.
//
// Example:
//   import notifications "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk"import chatbot "github.com/aws/aws-cdk-go/awscdk"
//
//   project := codebuild.NewPipelineProject(this, jsii.String("MyProject"))
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic1"))
//
//   slack := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &slackChannelConfigurationProps{
//   	slackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
//   	slackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
//   	slackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
//   })
//
//   rule := notifications.NewNotificationRule(this, jsii.String("NotificationRule"), &notificationRuleProps{
//   	source: project,
//   	events: []*string{
//   		jsii.String("codebuild-project-build-state-succeeded"),
//   		jsii.String("codebuild-project-build-state-failed"),
//   	},
//   	targets: []iNotificationRuleTarget{
//   		topic,
//   	},
//   })
//   rule.addTarget(slack)
//
type NotificationRule interface {
	awscdk.Resource
	INotificationRule
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// The ARN of the notification rule (i.e. arn:aws:codestar-notifications:::notificationrule/01234abcde).
	NotificationRuleArn() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds target to notification rule.
	AddTarget(target INotificationRuleTarget) *bool
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NotificationRule
type jsiiProxy_NotificationRule struct {
	internal.Type__awscdkResource
	jsiiProxy_INotificationRule
}

func (j *jsiiProxy_NotificationRule) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationRule) NotificationRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationRule) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewNotificationRule(scope constructs.Construct, id *string, props *NotificationRuleProps) NotificationRule {
	_init_.Initialize()

	j := jsiiProxy_NotificationRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codestarnotifications.NotificationRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNotificationRule_Override(n NotificationRule, scope constructs.Construct, id *string, props *NotificationRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codestarnotifications.NotificationRule",
		[]interface{}{scope, id, props},
		n,
	)
}

// Import an existing notification rule provided an ARN.
func NotificationRule_FromNotificationRuleArn(scope constructs.Construct, id *string, notificationRuleArn *string) INotificationRule {
	_init_.Initialize()

	var returns INotificationRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codestarnotifications.NotificationRule",
		"fromNotificationRuleArn",
		[]interface{}{scope, id, notificationRuleArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func NotificationRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codestarnotifications.NotificationRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func NotificationRule_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codestarnotifications.NotificationRule",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationRule) AddTarget(target INotificationRuleTarget) *bool {
	var returns *bool

	_jsii_.Invoke(
		n,
		"addTarget",
		[]interface{}{target},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_NotificationRule) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationRule) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationRule) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Standard set of options for `notifyOnXxx` codestar notification handler on construct.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codestarnotifications "github.com/aws/aws-cdk-go/awscdk/aws_codestarnotifications"
//   notificationRuleOptions := &notificationRuleOptions{
//   	detailType: codestarnotifications.detailType_BASIC,
//   	enabled: jsii.Boolean(false),
//   	notificationRuleName: jsii.String("notificationRuleName"),
//   }
//
type NotificationRuleOptions struct {
	// The level of detail to include in the notifications for this resource.
	//
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	DetailType DetailType `json:"detailType" yaml:"detailType"`
	// The status of the notification rule.
	//
	// If the enabled is set to DISABLED, notifications aren't sent for the notification rule.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account.
	NotificationRuleName *string `json:"notificationRuleName" yaml:"notificationRuleName"`
}

// Properties for a new notification rule.
//
// Example:
//   import notifications "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk"import chatbot "github.com/aws/aws-cdk-go/awscdk"
//
//   project := codebuild.NewPipelineProject(this, jsii.String("MyProject"))
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic1"))
//
//   slack := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &slackChannelConfigurationProps{
//   	slackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
//   	slackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
//   	slackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
//   })
//
//   rule := notifications.NewNotificationRule(this, jsii.String("NotificationRule"), &notificationRuleProps{
//   	source: project,
//   	events: []*string{
//   		jsii.String("codebuild-project-build-state-succeeded"),
//   		jsii.String("codebuild-project-build-state-failed"),
//   	},
//   	targets: []iNotificationRuleTarget{
//   		topic,
//   	},
//   })
//   rule.addTarget(slack)
//
type NotificationRuleProps struct {
	// The level of detail to include in the notifications for this resource.
	//
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	DetailType DetailType `json:"detailType" yaml:"detailType"`
	// The status of the notification rule.
	//
	// If the enabled is set to DISABLED, notifications aren't sent for the notification rule.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account.
	NotificationRuleName *string `json:"notificationRuleName" yaml:"notificationRuleName"`
	// A list of event types associated with this notification rule.
	//
	// For a complete list of event types and IDs, see Notification concepts in the Developer Tools Console User Guide.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#concepts-api
	//
	Events *[]*string `json:"events" yaml:"events"`
	// The Amazon Resource Name (ARN) of the resource to associate with the notification rule.
	//
	// Currently, Supported sources include pipelines in AWS CodePipeline, build projects in AWS CodeBuild, and repositories in AWS CodeCommit in this L2 constructor.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html#cfn-codestarnotifications-notificationrule-resource
	//
	Source INotificationRuleSource `json:"source" yaml:"source"`
	// The targets to register for the notification destination.
	Targets *[]INotificationRuleTarget `json:"targets" yaml:"targets"`
}

// Information about the Codebuild or CodePipeline associated with a notification source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codestarnotifications "github.com/aws/aws-cdk-go/awscdk/aws_codestarnotifications"
//   notificationRuleSourceConfig := &notificationRuleSourceConfig{
//   	sourceArn: jsii.String("sourceArn"),
//   }
//
type NotificationRuleSourceConfig struct {
	// The Amazon Resource Name (ARN) of the notification source.
	SourceArn *string `json:"sourceArn" yaml:"sourceArn"`
}

// Information about the SNS topic or AWS Chatbot client associated with a notification target.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codestarnotifications "github.com/aws/aws-cdk-go/awscdk/aws_codestarnotifications"
//   notificationRuleTargetConfig := &notificationRuleTargetConfig{
//   	targetAddress: jsii.String("targetAddress"),
//   	targetType: jsii.String("targetType"),
//   }
//
type NotificationRuleTargetConfig struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic or AWS Chatbot client.
	TargetAddress *string `json:"targetAddress" yaml:"targetAddress"`
	// The target type.
	//
	// Can be an Amazon SNS topic or AWS Chatbot client.
	TargetType *string `json:"targetType" yaml:"targetType"`
}

