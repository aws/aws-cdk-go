package awsiotevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::IoTEvents::AlarmModel`.
//
// Represents an alarm model to monitor an AWS IoT Events input attribute. You can use the alarm to get notified when the value is outside a specified range. For more information, see [Create an alarm model](https://docs.aws.amazon.com/iotevents/latest/developerguide/create-alarms.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlarmModel := awscdk.Aws_iotevents.NewCfnAlarmModel(this, jsii.String("MyCfnAlarmModel"), &cfnAlarmModelProps{
//   	alarmRule: &alarmRuleProperty{
//   		simpleRule: &simpleRuleProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			inputProperty: jsii.String("inputProperty"),
//   			threshold: jsii.String("threshold"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	alarmCapabilities: &alarmCapabilitiesProperty{
//   		acknowledgeFlow: &acknowledgeFlowProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		initializationConfiguration: &initializationConfigurationProperty{
//   			disabledOnInitialization: jsii.Boolean(false),
//   		},
//   	},
//   	alarmEventActions: &alarmEventActionsProperty{
//   		alarmActions: []interface{}{
//   			&alarmActionProperty{
//   				dynamoDb: &dynamoDBProperty{
//   					hashKeyField: jsii.String("hashKeyField"),
//   					hashKeyValue: jsii.String("hashKeyValue"),
//   					tableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					hashKeyType: jsii.String("hashKeyType"),
//   					operation: jsii.String("operation"),
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   					payloadField: jsii.String("payloadField"),
//   					rangeKeyField: jsii.String("rangeKeyField"),
//   					rangeKeyType: jsii.String("rangeKeyType"),
//   					rangeKeyValue: jsii.String("rangeKeyValue"),
//   				},
//   				dynamoDBv2: &dynamoDBv2Property{
//   					tableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				firehose: &firehoseProperty{
//   					deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   					separator: jsii.String("separator"),
//   				},
//   				iotEvents: &iotEventsProperty{
//   					inputName: jsii.String("inputName"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				iotSiteWise: &iotSiteWiseProperty{
//   					propertyValue: &assetPropertyValueProperty{
//   						value: &assetPropertyVariantProperty{
//   							booleanValue: jsii.String("booleanValue"),
//   							doubleValue: jsii.String("doubleValue"),
//   							integerValue: jsii.String("integerValue"),
//   							stringValue: jsii.String("stringValue"),
//   						},
//
//   						// the properties below are optional
//   						quality: jsii.String("quality"),
//   						timestamp: &assetPropertyTimestampProperty{
//   							timeInSeconds: jsii.String("timeInSeconds"),
//
//   							// the properties below are optional
//   							offsetInNanos: jsii.String("offsetInNanos"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					assetId: jsii.String("assetId"),
//   					entryId: jsii.String("entryId"),
//   					propertyAlias: jsii.String("propertyAlias"),
//   					propertyId: jsii.String("propertyId"),
//   				},
//   				iotTopicPublish: &iotTopicPublishProperty{
//   					mqttTopic: jsii.String("mqttTopic"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				lambda: &lambdaProperty{
//   					functionArn: jsii.String("functionArn"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				sns: &snsProperty{
//   					targetArn: jsii.String("targetArn"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				sqs: &sqsProperty{
//   					queueUrl: jsii.String("queueUrl"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   					useBase64: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	alarmModelDescription: jsii.String("alarmModelDescription"),
//   	alarmModelName: jsii.String("alarmModelName"),
//   	key: jsii.String("key"),
//   	severity: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnAlarmModel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Contains the configuration information of alarm state changes.
	AlarmCapabilities() interface{}
	SetAlarmCapabilities(val interface{})
	// Contains information about one or more alarm actions.
	AlarmEventActions() interface{}
	SetAlarmEventActions(val interface{})
	// The description of the alarm model.
	AlarmModelDescription() *string
	SetAlarmModelDescription(val *string)
	// The name of the alarm model.
	AlarmModelName() *string
	SetAlarmModelName(val *string)
	// Defines when your alarm is invoked.
	AlarmRule() interface{}
	SetAlarmRule(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// An input attribute used as a key to create an alarm.
	//
	// AWS IoT Events routes [inputs](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Input.html) associated with this key to the alarm.
	Key() *string
	SetKey(val *string)
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
	// The ARN of the IAM role that allows the alarm to perform actions and access AWS resources.
	//
	// For more information, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* .
	RoleArn() *string
	SetRoleArn(val *string)
	// A non-negative integer that reflects the severity level of the alarm.
	Severity() *float64
	SetSeverity(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A list of key-value pairs that contain metadata for the alarm model.
	//
	// The tags help you manage the alarm model. For more information, see [Tagging your AWS IoT Events resources](https://docs.aws.amazon.com/iotevents/latest/developerguide/tagging-iotevents.html) in the *AWS IoT Events Developer Guide* .
	//
	// You can create up to 50 tags for one alarm model.
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

// The jsii proxy struct for CfnAlarmModel
type jsiiProxy_CfnAlarmModel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAlarmModel) AlarmCapabilities() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) AlarmEventActions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmEventActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) AlarmModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmModelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) AlarmModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) AlarmRule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) Severity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTEvents::AlarmModel`.
func NewCfnAlarmModel(scope constructs.Construct, id *string, props *CfnAlarmModelProps) CfnAlarmModel {
	_init_.Initialize()

	j := jsiiProxy_CfnAlarmModel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotevents.CfnAlarmModel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTEvents::AlarmModel`.
func NewCfnAlarmModel_Override(c CfnAlarmModel, scope constructs.Construct, id *string, props *CfnAlarmModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotevents.CfnAlarmModel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAlarmModel) SetAlarmCapabilities(val interface{}) {
	_jsii_.Set(
		j,
		"alarmCapabilities",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel) SetAlarmEventActions(val interface{}) {
	_jsii_.Set(
		j,
		"alarmEventActions",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel) SetAlarmModelDescription(val *string) {
	_jsii_.Set(
		j,
		"alarmModelDescription",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel) SetAlarmModelName(val *string) {
	_jsii_.Set(
		j,
		"alarmModelName",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel) SetAlarmRule(val interface{}) {
	_jsii_.Set(
		j,
		"alarmRule",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel) SetSeverity(val *float64) {
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAlarmModel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotevents.CfnAlarmModel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAlarmModel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotevents.CfnAlarmModel",
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
func CfnAlarmModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotevents.CfnAlarmModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAlarmModel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotevents.CfnAlarmModel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAlarmModel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAlarmModel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAlarmModel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAlarmModel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAlarmModel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAlarmModel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAlarmModel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAlarmModel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarmModel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarmModel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAlarmModel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAlarmModel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarmModel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarmModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarmModel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies whether to get notified for alarm state changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acknowledgeFlowProperty := &acknowledgeFlowProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnAlarmModel_AcknowledgeFlowProperty struct {
	// The value must be `TRUE` or `FALSE` .
	//
	// If `TRUE` , you receive a notification when the alarm state changes. You must choose to acknowledge the notification before the alarm state can return to `NORMAL` . If `FALSE` , you won't receive notifications. The alarm automatically changes to the `NORMAL` state when the input property value returns to the specified range.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

// Specifies one of the following actions to receive notifications when the alarm state changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmActionProperty := &alarmActionProperty{
//   	dynamoDb: &dynamoDBProperty{
//   		hashKeyField: jsii.String("hashKeyField"),
//   		hashKeyValue: jsii.String("hashKeyValue"),
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		hashKeyType: jsii.String("hashKeyType"),
//   		operation: jsii.String("operation"),
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   		payloadField: jsii.String("payloadField"),
//   		rangeKeyField: jsii.String("rangeKeyField"),
//   		rangeKeyType: jsii.String("rangeKeyType"),
//   		rangeKeyValue: jsii.String("rangeKeyValue"),
//   	},
//   	dynamoDBv2: &dynamoDBv2Property{
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	firehose: &firehoseProperty{
//   		deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   		separator: jsii.String("separator"),
//   	},
//   	iotEvents: &iotEventsProperty{
//   		inputName: jsii.String("inputName"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	iotSiteWise: &iotSiteWiseProperty{
//   		propertyValue: &assetPropertyValueProperty{
//   			value: &assetPropertyVariantProperty{
//   				booleanValue: jsii.String("booleanValue"),
//   				doubleValue: jsii.String("doubleValue"),
//   				integerValue: jsii.String("integerValue"),
//   				stringValue: jsii.String("stringValue"),
//   			},
//
//   			// the properties below are optional
//   			quality: jsii.String("quality"),
//   			timestamp: &assetPropertyTimestampProperty{
//   				timeInSeconds: jsii.String("timeInSeconds"),
//
//   				// the properties below are optional
//   				offsetInNanos: jsii.String("offsetInNanos"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		assetId: jsii.String("assetId"),
//   		entryId: jsii.String("entryId"),
//   		propertyAlias: jsii.String("propertyAlias"),
//   		propertyId: jsii.String("propertyId"),
//   	},
//   	iotTopicPublish: &iotTopicPublishProperty{
//   		mqttTopic: jsii.String("mqttTopic"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	lambda: &lambdaProperty{
//   		functionArn: jsii.String("functionArn"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	sns: &snsProperty{
//   		targetArn: jsii.String("targetArn"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	sqs: &sqsProperty{
//   		queueUrl: jsii.String("queueUrl"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   		useBase64: jsii.Boolean(false),
//   	},
//   }
//
type CfnAlarmModel_AlarmActionProperty struct {
	// Defines an action to write to the Amazon DynamoDB table that you created.
	//
	// The standard action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify.
	//
	// You must use expressions for all parameters in `DynamoDBAction` . The expressions accept literals, operators, functions, references, and substitution templates.
	//
	// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `hashKeyType` parameter can be `'STRING'` .
	// - For references, you must specify either variables or input values. For example, the value for the `hashKeyField` parameter can be `$input.GreenhouseInput.name` .
	// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
	//
	// In the following example, the value for the `hashKeyValue` parameter uses a substitution template.
	//
	// `'${$input.GreenhouseInput.temperature * 6 / 5 + 32} in Fahrenheit'`
	// - For a string concatenation, you must use `+` . A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
	//
	// In the following example, the value for the `tableName` parameter uses a string concatenation.
	//
	// `'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`
	//
	// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
	//
	// If the defined payload type is a string, `DynamoDBAction` writes non-JSON data to the DynamoDB table as binary data. The DynamoDB console displays the data as Base64-encoded text. The value for the `payloadField` parameter is `<payload-field>_raw` .
	DynamoDb interface{} `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Defines an action to write to the Amazon DynamoDB table that you created.
	//
	// The default action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.
	//
	// You must use expressions for all parameters in `DynamoDBv2Action` . The expressions accept literals, operators, functions, references, and substitution templates.
	//
	// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `tableName` parameter can be `'GreenhouseTemperatureTable'` .
	// - For references, you must specify either variables or input values. For example, the value for the `tableName` parameter can be `$variable.ddbtableName` .
	// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
	//
	// In the following example, the value for the `contentExpression` parameter in `Payload` uses a substitution template.
	//
	// `'{\"sensorID\": \"${$input.GreenhouseInput.sensor_id}\", \"temperature\": \"${$input.GreenhouseInput.temperature * 9 / 5 + 32}\"}'`
	// - For a string concatenation, you must use `+` . A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
	//
	// In the following example, the value for the `tableName` parameter uses a string concatenation.
	//
	// `'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`
	//
	// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
	//
	// The value for the `type` parameter in `Payload` must be `JSON` .
	DynamoDBv2 interface{} `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.
	IotEvents interface{} `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// Sends information about the detector model instance and the event that triggered the action to a specified asset property in AWS IoT SiteWise .
	//
	// You must use expressions for all parameters in `IotSiteWiseAction` . The expressions accept literals, operators, functions, references, and substitutions templates.
	//
	// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `propertyAlias` parameter can be `'/company/windfarm/3/turbine/7/temperature'` .
	// - For references, you must specify either variables or input values. For example, the value for the `assetId` parameter can be `$input.TurbineInput.assetId1` .
	// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
	//
	// In the following example, the value for the `propertyAlias` parameter uses a substitution template.
	//
	// `'company/windfarm/${$input.TemperatureInput.sensorData.windfarmID}/turbine/ ${$input.TemperatureInput.sensorData.turbineID}/temperature'`
	//
	// You must specify either `propertyAlias` or both `assetId` and `propertyId` to identify the target asset property in AWS IoT SiteWise .
	//
	// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
	IotSiteWise interface{} `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Information required to publish the MQTT message through the AWS IoT message broker.
	IotTopicPublish interface{} `field:"optional" json:"iotTopicPublish" yaml:"iotTopicPublish"`
	// Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// `CfnAlarmModel.AlarmActionProperty.Sns`.
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
	// Sends information about the detector model instance and the event that triggered the action to an Amazon SQS queue.
	Sqs interface{} `field:"optional" json:"sqs" yaml:"sqs"`
}

// Contains the configuration information of alarm state changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmCapabilitiesProperty := &alarmCapabilitiesProperty{
//   	acknowledgeFlow: &acknowledgeFlowProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	initializationConfiguration: &initializationConfigurationProperty{
//   		disabledOnInitialization: jsii.Boolean(false),
//   	},
//   }
//
type CfnAlarmModel_AlarmCapabilitiesProperty struct {
	// Specifies whether to get notified for alarm state changes.
	AcknowledgeFlow interface{} `field:"optional" json:"acknowledgeFlow" yaml:"acknowledgeFlow"`
	// Specifies the default alarm state.
	//
	// The configuration applies to all alarms that were created based on this alarm model.
	InitializationConfiguration interface{} `field:"optional" json:"initializationConfiguration" yaml:"initializationConfiguration"`
}

// Contains information about one or more alarm actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmEventActionsProperty := &alarmEventActionsProperty{
//   	alarmActions: []interface{}{
//   		&alarmActionProperty{
//   			dynamoDb: &dynamoDBProperty{
//   				hashKeyField: jsii.String("hashKeyField"),
//   				hashKeyValue: jsii.String("hashKeyValue"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				hashKeyType: jsii.String("hashKeyType"),
//   				operation: jsii.String("operation"),
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   				payloadField: jsii.String("payloadField"),
//   				rangeKeyField: jsii.String("rangeKeyField"),
//   				rangeKeyType: jsii.String("rangeKeyType"),
//   				rangeKeyValue: jsii.String("rangeKeyValue"),
//   			},
//   			dynamoDBv2: &dynamoDBv2Property{
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			firehose: &firehoseProperty{
//   				deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   				separator: jsii.String("separator"),
//   			},
//   			iotEvents: &iotEventsProperty{
//   				inputName: jsii.String("inputName"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			iotSiteWise: &iotSiteWiseProperty{
//   				propertyValue: &assetPropertyValueProperty{
//   					value: &assetPropertyVariantProperty{
//   						booleanValue: jsii.String("booleanValue"),
//   						doubleValue: jsii.String("doubleValue"),
//   						integerValue: jsii.String("integerValue"),
//   						stringValue: jsii.String("stringValue"),
//   					},
//
//   					// the properties below are optional
//   					quality: jsii.String("quality"),
//   					timestamp: &assetPropertyTimestampProperty{
//   						timeInSeconds: jsii.String("timeInSeconds"),
//
//   						// the properties below are optional
//   						offsetInNanos: jsii.String("offsetInNanos"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				assetId: jsii.String("assetId"),
//   				entryId: jsii.String("entryId"),
//   				propertyAlias: jsii.String("propertyAlias"),
//   				propertyId: jsii.String("propertyId"),
//   			},
//   			iotTopicPublish: &iotTopicPublishProperty{
//   				mqttTopic: jsii.String("mqttTopic"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			lambda: &lambdaProperty{
//   				functionArn: jsii.String("functionArn"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			sns: &snsProperty{
//   				targetArn: jsii.String("targetArn"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			sqs: &sqsProperty{
//   				queueUrl: jsii.String("queueUrl"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   				useBase64: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
//
type CfnAlarmModel_AlarmEventActionsProperty struct {
	// Specifies one or more supported actions to receive notifications when the alarm state changes.
	AlarmActions interface{} `field:"optional" json:"alarmActions" yaml:"alarmActions"`
}

// Defines when your alarm is invoked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmRuleProperty := &alarmRuleProperty{
//   	simpleRule: &simpleRuleProperty{
//   		comparisonOperator: jsii.String("comparisonOperator"),
//   		inputProperty: jsii.String("inputProperty"),
//   		threshold: jsii.String("threshold"),
//   	},
//   }
//
type CfnAlarmModel_AlarmRuleProperty struct {
	// A rule that compares an input property value to a threshold value with a comparison operator.
	SimpleRule interface{} `field:"optional" json:"simpleRule" yaml:"simpleRule"`
}

// A structure that contains timestamp information. For more information, see [TimeInNanos](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_TimeInNanos.html) in the *AWS IoT SiteWise API Reference* .
//
// You must use expressions for all parameters in `AssetPropertyTimestamp` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `timeInSeconds` parameter can be `'1586400675'` .
// - For references, you must specify either variables or input values. For example, the value for the `offsetInNanos` parameter can be `$variable.time` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `timeInSeconds` parameter uses a substitution template.
//
// `'${$input.TemperatureInput.sensorData.timestamp / 1000}'`
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyTimestampProperty := &assetPropertyTimestampProperty{
//   	timeInSeconds: jsii.String("timeInSeconds"),
//
//   	// the properties below are optional
//   	offsetInNanos: jsii.String("offsetInNanos"),
//   }
//
type CfnAlarmModel_AssetPropertyTimestampProperty struct {
	// The timestamp, in seconds, in the Unix epoch format.
	//
	// The valid range is between 1-31556889864403199.
	TimeInSeconds *string `field:"required" json:"timeInSeconds" yaml:"timeInSeconds"`
	// The nanosecond offset converted from `timeInSeconds` .
	//
	// The valid range is between 0-999999999.
	OffsetInNanos *string `field:"optional" json:"offsetInNanos" yaml:"offsetInNanos"`
}

// A structure that contains value information. For more information, see [AssetPropertyValue](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_AssetPropertyValue.html) in the *AWS IoT SiteWise API Reference* .
//
// You must use expressions for all parameters in `AssetPropertyValue` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `quality` parameter can be `'GOOD'` .
// - For references, you must specify either variables or input values. For example, the value for the `quality` parameter can be `$input.TemperatureInput.sensorData.quality` .
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyValueProperty := &assetPropertyValueProperty{
//   	value: &assetPropertyVariantProperty{
//   		booleanValue: jsii.String("booleanValue"),
//   		doubleValue: jsii.String("doubleValue"),
//   		integerValue: jsii.String("integerValue"),
//   		stringValue: jsii.String("stringValue"),
//   	},
//
//   	// the properties below are optional
//   	quality: jsii.String("quality"),
//   	timestamp: &assetPropertyTimestampProperty{
//   		timeInSeconds: jsii.String("timeInSeconds"),
//
//   		// the properties below are optional
//   		offsetInNanos: jsii.String("offsetInNanos"),
//   	},
//   }
//
type CfnAlarmModel_AssetPropertyValueProperty struct {
	// The value to send to an asset property.
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// The quality of the asset property value.
	//
	// The value must be `'GOOD'` , `'BAD'` , or `'UNCERTAIN'` .
	Quality *string `field:"optional" json:"quality" yaml:"quality"`
	// The timestamp associated with the asset property value.
	//
	// The default is the current event time.
	Timestamp interface{} `field:"optional" json:"timestamp" yaml:"timestamp"`
}

// A structure that contains an asset property value.
//
// For more information, see [Variant](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_Variant.html) in the *AWS IoT SiteWise API Reference* .
//
// You must use expressions for all parameters in `AssetPropertyVariant` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `integerValue` parameter can be `'100'` .
// - For references, you must specify either variables or parameters. For example, the value for the `booleanValue` parameter can be `$variable.offline` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `doubleValue` parameter uses a substitution template.
//
// `'${$input.TemperatureInput.sensorData.temperature * 6 / 5 + 32}'`
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// You must specify one of the following value types, depending on the `dataType` of the specified asset property. For more information, see [AssetProperty](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_AssetProperty.html) in the *AWS IoT SiteWise API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyVariantProperty := &assetPropertyVariantProperty{
//   	booleanValue: jsii.String("booleanValue"),
//   	doubleValue: jsii.String("doubleValue"),
//   	integerValue: jsii.String("integerValue"),
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnAlarmModel_AssetPropertyVariantProperty struct {
	// The asset property value is a Boolean value that must be `'TRUE'` or `'FALSE'` .
	//
	// You must use an expression, and the evaluated result should be a Boolean value.
	BooleanValue *string `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// The asset property value is a double.
	//
	// You must use an expression, and the evaluated result should be a double.
	DoubleValue *string `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// The asset property value is an integer.
	//
	// You must use an expression, and the evaluated result should be an integer.
	IntegerValue *string `field:"optional" json:"integerValue" yaml:"integerValue"`
	// The asset property value is a string.
	//
	// You must use an expression, and the evaluated result should be a string.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

// Defines an action to write to the Amazon DynamoDB table that you created.
//
// The standard action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify.
//
// You must use expressions for all parameters in `DynamoDBAction` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `hashKeyType` parameter can be `'STRING'` .
// - For references, you must specify either variables or input values. For example, the value for the `hashKeyField` parameter can be `$input.GreenhouseInput.name` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `hashKeyValue` parameter uses a substitution template.
//
// `'${$input.GreenhouseInput.temperature * 6 / 5 + 32} in Fahrenheit'`
// - For a string concatenation, you must use `+` . A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `tableName` parameter uses a string concatenation.
//
// `'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// If the defined payload type is a string, `DynamoDBAction` writes non-JSON data to the DynamoDB table as binary data. The DynamoDB console displays the data as Base64-encoded text. The value for the `payloadField` parameter is `<payload-field>_raw` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDBProperty := &dynamoDBProperty{
//   	hashKeyField: jsii.String("hashKeyField"),
//   	hashKeyValue: jsii.String("hashKeyValue"),
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	hashKeyType: jsii.String("hashKeyType"),
//   	operation: jsii.String("operation"),
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   	payloadField: jsii.String("payloadField"),
//   	rangeKeyField: jsii.String("rangeKeyField"),
//   	rangeKeyType: jsii.String("rangeKeyType"),
//   	rangeKeyValue: jsii.String("rangeKeyValue"),
//   }
//
type CfnAlarmModel_DynamoDBProperty struct {
	// The name of the hash key (also called the partition key).
	//
	// The `hashKeyField` value must match the partition key of the target DynamoDB table.
	HashKeyField *string `field:"required" json:"hashKeyField" yaml:"hashKeyField"`
	// The value of the hash key (also called the partition key).
	HashKeyValue *string `field:"required" json:"hashKeyValue" yaml:"hashKeyValue"`
	// The name of the DynamoDB table.
	//
	// The `tableName` value must match the table name of the target DynamoDB table.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The data type for the hash key (also called the partition key). You can specify the following values:.
	//
	// - `'STRING'` - The hash key is a string.
	// - `'NUMBER'` - The hash key is a number.
	//
	// If you don't specify `hashKeyType` , the default value is `'STRING'` .
	HashKeyType *string `field:"optional" json:"hashKeyType" yaml:"hashKeyType"`
	// The type of operation to perform. You can specify the following values:.
	//
	// - `'INSERT'` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.
	// - `'UPDATE'` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.
	// - `'DELETE'` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.
	//
	// If you don't specify this parameter, AWS IoT Events triggers the `'INSERT'` operation.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Information needed to configure the payload.
	//
	// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use `contentExpression` .
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
	// The name of the DynamoDB column that receives the action payload.
	//
	// If you don't specify this parameter, the name of the DynamoDB column is `payload` .
	PayloadField *string `field:"optional" json:"payloadField" yaml:"payloadField"`
	// The name of the range key (also called the sort key).
	//
	// The `rangeKeyField` value must match the sort key of the target DynamoDB table.
	RangeKeyField *string `field:"optional" json:"rangeKeyField" yaml:"rangeKeyField"`
	// The data type for the range key (also called the sort key), You can specify the following values:.
	//
	// - `'STRING'` - The range key is a string.
	// - `'NUMBER'` - The range key is number.
	//
	// If you don't specify `rangeKeyField` , the default value is `'STRING'` .
	RangeKeyType *string `field:"optional" json:"rangeKeyType" yaml:"rangeKeyType"`
	// The value of the range key (also called the sort key).
	RangeKeyValue *string `field:"optional" json:"rangeKeyValue" yaml:"rangeKeyValue"`
}

// Defines an action to write to the Amazon DynamoDB table that you created.
//
// The default action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.
//
// You must use expressions for all parameters in `DynamoDBv2Action` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `tableName` parameter can be `'GreenhouseTemperatureTable'` .
// - For references, you must specify either variables or input values. For example, the value for the `tableName` parameter can be `$variable.ddbtableName` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `contentExpression` parameter in `Payload` uses a substitution template.
//
// `'{\"sensorID\": \"${$input.GreenhouseInput.sensor_id}\", \"temperature\": \"${$input.GreenhouseInput.temperature * 9 / 5 + 32}\"}'`
// - For a string concatenation, you must use `+` . A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `tableName` parameter uses a string concatenation.
//
// `'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// The value for the `type` parameter in `Payload` must be `JSON` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDBv2Property := &dynamoDBv2Property{
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnAlarmModel_DynamoDBv2Property struct {
	// The name of the DynamoDB table.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// `CfnAlarmModel.DynamoDBv2Property.Payload`.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseProperty := &firehoseProperty{
//   	deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   	separator: jsii.String("separator"),
//   }
//
type CfnAlarmModel_FirehoseProperty struct {
	// The name of the Kinesis Data Firehose delivery stream where the data is written.
	DeliveryStreamName *string `field:"required" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// You can configure the action payload when you send a message to an Amazon Kinesis Data Firehose delivery stream.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
	// A character separator that is used to separate records written to the Kinesis Data Firehose delivery stream.
	//
	// Valid values are: '\n' (newline), '\t' (tab), '\r\n' (Windows newline), ',' (comma).
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
}

// Specifies the default alarm state.
//
// The configuration applies to all alarms that were created based on this alarm model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   initializationConfigurationProperty := &initializationConfigurationProperty{
//   	disabledOnInitialization: jsii.Boolean(false),
//   }
//
type CfnAlarmModel_InitializationConfigurationProperty struct {
	// The value must be `TRUE` or `FALSE` .
	//
	// If `FALSE` , all alarm instances created based on the alarm model are activated. The default value is `TRUE` .
	DisabledOnInitialization interface{} `field:"required" json:"disabledOnInitialization" yaml:"disabledOnInitialization"`
}

// Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotEventsProperty := &iotEventsProperty{
//   	inputName: jsii.String("inputName"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnAlarmModel_IotEventsProperty struct {
	// The name of the AWS IoT Events input where the data is sent.
	InputName *string `field:"required" json:"inputName" yaml:"inputName"`
	// You can configure the action payload when you send a message to an AWS IoT Events input.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

// Sends information about the detector model instance and the event that triggered the action to a specified asset property in AWS IoT SiteWise .
//
// You must use expressions for all parameters in `IotSiteWiseAction` . The expressions accept literals, operators, functions, references, and substitutions templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `propertyAlias` parameter can be `'/company/windfarm/3/turbine/7/temperature'` .
// - For references, you must specify either variables or input values. For example, the value for the `assetId` parameter can be `$input.TurbineInput.assetId1` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `propertyAlias` parameter uses a substitution template.
//
// `'company/windfarm/${$input.TemperatureInput.sensorData.windfarmID}/turbine/ ${$input.TemperatureInput.sensorData.turbineID}/temperature'`
//
// You must specify either `propertyAlias` or both `assetId` and `propertyId` to identify the target asset property in AWS IoT SiteWise .
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotSiteWiseProperty := &iotSiteWiseProperty{
//   	propertyValue: &assetPropertyValueProperty{
//   		value: &assetPropertyVariantProperty{
//   			booleanValue: jsii.String("booleanValue"),
//   			doubleValue: jsii.String("doubleValue"),
//   			integerValue: jsii.String("integerValue"),
//   			stringValue: jsii.String("stringValue"),
//   		},
//
//   		// the properties below are optional
//   		quality: jsii.String("quality"),
//   		timestamp: &assetPropertyTimestampProperty{
//   			timeInSeconds: jsii.String("timeInSeconds"),
//
//   			// the properties below are optional
//   			offsetInNanos: jsii.String("offsetInNanos"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	assetId: jsii.String("assetId"),
//   	entryId: jsii.String("entryId"),
//   	propertyAlias: jsii.String("propertyAlias"),
//   	propertyId: jsii.String("propertyId"),
//   }
//
type CfnAlarmModel_IotSiteWiseProperty struct {
	// The value to send to the asset property.
	//
	// This value contains timestamp, quality, and value (TQV) information.
	PropertyValue interface{} `field:"required" json:"propertyValue" yaml:"propertyValue"`
	// The ID of the asset that has the specified property.
	AssetId *string `field:"optional" json:"assetId" yaml:"assetId"`
	// A unique identifier for this entry.
	//
	// You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier.
	EntryId *string `field:"optional" json:"entryId" yaml:"entryId"`
	// The alias of the asset property.
	PropertyAlias *string `field:"optional" json:"propertyAlias" yaml:"propertyAlias"`
	// The ID of the asset property.
	PropertyId *string `field:"optional" json:"propertyId" yaml:"propertyId"`
}

// Information required to publish the MQTT message through the AWS IoT message broker.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotTopicPublishProperty := &iotTopicPublishProperty{
//   	mqttTopic: jsii.String("mqttTopic"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnAlarmModel_IotTopicPublishProperty struct {
	// The MQTT topic of the message.
	//
	// You can use a string expression that includes variables ( `$variable.<variable-name>` ) and input values ( `$input.<input-name>.<path-to-datum>` ) as the topic string.
	MqttTopic *string `field:"required" json:"mqttTopic" yaml:"mqttTopic"`
	// You can configure the action payload when you publish a message to an AWS IoT Core topic.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

// Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaProperty := &lambdaProperty{
//   	functionArn: jsii.String("functionArn"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnAlarmModel_LambdaProperty struct {
	// The ARN of the Lambda function that is executed.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// You can configure the action payload when you send a message to a Lambda function.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

// Information needed to configure the payload.
//
// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use `contentExpression` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   payloadProperty := &payloadProperty{
//   	contentExpression: jsii.String("contentExpression"),
//   	type: jsii.String("type"),
//   }
//
type CfnAlarmModel_PayloadProperty struct {
	// The content of the payload.
	//
	// You can use a string expression that includes quoted strings ( `'<string>'` ), variables ( `$variable.<variable-name>` ), input values ( `$input.<input-name>.<path-to-datum>` ), string concatenations, and quoted strings that contain `${}` as the content. The recommended maximum size of a content expression is 1 KB.
	ContentExpression *string `field:"required" json:"contentExpression" yaml:"contentExpression"`
	// The value of the payload type can be either `STRING` or `JSON` .
	Type *string `field:"required" json:"type" yaml:"type"`
}

// A rule that compares an input property value to a threshold value with a comparison operator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simpleRuleProperty := &simpleRuleProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	inputProperty: jsii.String("inputProperty"),
//   	threshold: jsii.String("threshold"),
//   }
//
type CfnAlarmModel_SimpleRuleProperty struct {
	// The comparison operator.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The value on the left side of the comparison operator.
	//
	// You can specify an AWS IoT Events input attribute as an input property.
	InputProperty *string `field:"required" json:"inputProperty" yaml:"inputProperty"`
	// The value on the right side of the comparison operator.
	//
	// You can enter a number or specify an AWS IoT Events input attribute.
	Threshold *string `field:"required" json:"threshold" yaml:"threshold"`
}

// Information required to publish the Amazon SNS message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snsProperty := &snsProperty{
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnAlarmModel_SnsProperty struct {
	// The ARN of the Amazon SNS target where the message is sent.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// You can configure the action payload when you send a message as an Amazon SNS push notification.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

// Sends information about the detector model instance and the event that triggered the action to an Amazon SQS queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqsProperty := &sqsProperty{
//   	queueUrl: jsii.String("queueUrl"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   	useBase64: jsii.Boolean(false),
//   }
//
type CfnAlarmModel_SqsProperty struct {
	// The URL of the SQS queue where the data is written.
	QueueUrl *string `field:"required" json:"queueUrl" yaml:"queueUrl"`
	// You can configure the action payload when you send a message to an Amazon SQS queue.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
	// Set this to TRUE if you want the data to be base-64 encoded before it is written to the queue.
	//
	// Otherwise, set this to FALSE.
	UseBase64 interface{} `field:"optional" json:"useBase64" yaml:"useBase64"`
}

// Properties for defining a `CfnAlarmModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlarmModelProps := &cfnAlarmModelProps{
//   	alarmRule: &alarmRuleProperty{
//   		simpleRule: &simpleRuleProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			inputProperty: jsii.String("inputProperty"),
//   			threshold: jsii.String("threshold"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	alarmCapabilities: &alarmCapabilitiesProperty{
//   		acknowledgeFlow: &acknowledgeFlowProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		initializationConfiguration: &initializationConfigurationProperty{
//   			disabledOnInitialization: jsii.Boolean(false),
//   		},
//   	},
//   	alarmEventActions: &alarmEventActionsProperty{
//   		alarmActions: []interface{}{
//   			&alarmActionProperty{
//   				dynamoDb: &dynamoDBProperty{
//   					hashKeyField: jsii.String("hashKeyField"),
//   					hashKeyValue: jsii.String("hashKeyValue"),
//   					tableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					hashKeyType: jsii.String("hashKeyType"),
//   					operation: jsii.String("operation"),
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   					payloadField: jsii.String("payloadField"),
//   					rangeKeyField: jsii.String("rangeKeyField"),
//   					rangeKeyType: jsii.String("rangeKeyType"),
//   					rangeKeyValue: jsii.String("rangeKeyValue"),
//   				},
//   				dynamoDBv2: &dynamoDBv2Property{
//   					tableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				firehose: &firehoseProperty{
//   					deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   					separator: jsii.String("separator"),
//   				},
//   				iotEvents: &iotEventsProperty{
//   					inputName: jsii.String("inputName"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				iotSiteWise: &iotSiteWiseProperty{
//   					propertyValue: &assetPropertyValueProperty{
//   						value: &assetPropertyVariantProperty{
//   							booleanValue: jsii.String("booleanValue"),
//   							doubleValue: jsii.String("doubleValue"),
//   							integerValue: jsii.String("integerValue"),
//   							stringValue: jsii.String("stringValue"),
//   						},
//
//   						// the properties below are optional
//   						quality: jsii.String("quality"),
//   						timestamp: &assetPropertyTimestampProperty{
//   							timeInSeconds: jsii.String("timeInSeconds"),
//
//   							// the properties below are optional
//   							offsetInNanos: jsii.String("offsetInNanos"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					assetId: jsii.String("assetId"),
//   					entryId: jsii.String("entryId"),
//   					propertyAlias: jsii.String("propertyAlias"),
//   					propertyId: jsii.String("propertyId"),
//   				},
//   				iotTopicPublish: &iotTopicPublishProperty{
//   					mqttTopic: jsii.String("mqttTopic"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				lambda: &lambdaProperty{
//   					functionArn: jsii.String("functionArn"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				sns: &snsProperty{
//   					targetArn: jsii.String("targetArn"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				sqs: &sqsProperty{
//   					queueUrl: jsii.String("queueUrl"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   					useBase64: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	alarmModelDescription: jsii.String("alarmModelDescription"),
//   	alarmModelName: jsii.String("alarmModelName"),
//   	key: jsii.String("key"),
//   	severity: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAlarmModelProps struct {
	// Defines when your alarm is invoked.
	AlarmRule interface{} `field:"required" json:"alarmRule" yaml:"alarmRule"`
	// The ARN of the IAM role that allows the alarm to perform actions and access AWS resources.
	//
	// For more information, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Contains the configuration information of alarm state changes.
	AlarmCapabilities interface{} `field:"optional" json:"alarmCapabilities" yaml:"alarmCapabilities"`
	// Contains information about one or more alarm actions.
	AlarmEventActions interface{} `field:"optional" json:"alarmEventActions" yaml:"alarmEventActions"`
	// The description of the alarm model.
	AlarmModelDescription *string `field:"optional" json:"alarmModelDescription" yaml:"alarmModelDescription"`
	// The name of the alarm model.
	AlarmModelName *string `field:"optional" json:"alarmModelName" yaml:"alarmModelName"`
	// An input attribute used as a key to create an alarm.
	//
	// AWS IoT Events routes [inputs](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Input.html) associated with this key to the alarm.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A non-negative integer that reflects the severity level of the alarm.
	Severity *float64 `field:"optional" json:"severity" yaml:"severity"`
	// A list of key-value pairs that contain metadata for the alarm model.
	//
	// The tags help you manage the alarm model. For more information, see [Tagging your AWS IoT Events resources](https://docs.aws.amazon.com/iotevents/latest/developerguide/tagging-iotevents.html) in the *AWS IoT Events Developer Guide* .
	//
	// You can create up to 50 tags for one alarm model.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTEvents::DetectorModel`.
//
// The AWS::IoTEvents::DetectorModel resource creates a detector model. You create a *detector model* (a model of your equipment or process) using *states* . For each state, you define conditional (Boolean) logic that evaluates the incoming inputs to detect significant events. When an event is detected, it can change the state or trigger custom-built or predefined actions using other AWS services. You can define additional events that trigger actions when entering or exiting a state and, optionally, when a condition is met. For more information, see [How to Use AWS IoT Events](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *AWS IoT Events Developer Guide* .
//
// > When you successfully update a detector model (using the AWS IoT Events console, AWS IoT Events API or CLI commands, or AWS CloudFormation ) all detector instances created by the model are reset to their initial states. (The detector's `state` , and the values of any variables and timers are reset.)
// >
// > When you successfully update a detector model (using the AWS IoT Events console, AWS IoT Events API or CLI commands, or AWS CloudFormation ) the version number of the detector model is incremented. (A detector model with version number 1 before the update has version number 2 after the update succeeds.)
// >
// > If you attempt to update a detector model using AWS CloudFormation and the update does not succeed, the system may, in some cases, restore the original detector model. When this occurs, the detector model's version is incremented twice (for example, from version 1 to version 3) and the detector instances are reset.
// >
// > Also, be aware that if you attempt to update several detector models at once using AWS CloudFormation , some updates may succeed and others fail. In this case, the effects on each detector model's detector instances and version number depend on whether the update succeeded or failed, with the results as stated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDetectorModel := awscdk.Aws_iotevents.NewCfnDetectorModel(this, jsii.String("MyCfnDetectorModel"), &cfnDetectorModelProps{
//   	detectorModelDefinition: &detectorModelDefinitionProperty{
//   		initialStateName: jsii.String("initialStateName"),
//   		states: []interface{}{
//   			&stateProperty{
//   				stateName: jsii.String("stateName"),
//
//   				// the properties below are optional
//   				onEnter: &onEnterProperty{
//   					events: []interface{}{
//   						&eventProperty{
//   							eventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							actions: []interface{}{
//   								&actionProperty{
//   									clearTimer: &clearTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									dynamoDb: &dynamoDBProperty{
//   										hashKeyField: jsii.String("hashKeyField"),
//   										hashKeyValue: jsii.String("hashKeyValue"),
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										hashKeyType: jsii.String("hashKeyType"),
//   										operation: jsii.String("operation"),
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										payloadField: jsii.String("payloadField"),
//   										rangeKeyField: jsii.String("rangeKeyField"),
//   										rangeKeyType: jsii.String("rangeKeyType"),
//   										rangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									dynamoDBv2: &dynamoDBv2Property{
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									firehose: &firehoseProperty{
//   										deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										separator: jsii.String("separator"),
//   									},
//   									iotEvents: &iotEventsProperty{
//   										inputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									iotSiteWise: &iotSiteWiseProperty{
//   										propertyValue: &assetPropertyValueProperty{
//   											value: &assetPropertyVariantProperty{
//   												booleanValue: jsii.String("booleanValue"),
//   												doubleValue: jsii.String("doubleValue"),
//   												integerValue: jsii.String("integerValue"),
//   												stringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											quality: jsii.String("quality"),
//   											timestamp: &assetPropertyTimestampProperty{
//   												timeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												offsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										assetId: jsii.String("assetId"),
//   										entryId: jsii.String("entryId"),
//   										propertyAlias: jsii.String("propertyAlias"),
//   										propertyId: jsii.String("propertyId"),
//   									},
//   									iotTopicPublish: &iotTopicPublishProperty{
//   										mqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									lambda: &lambdaProperty{
//   										functionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									resetTimer: &resetTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									setTimer: &setTimerProperty{
//   										timerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										durationExpression: jsii.String("durationExpression"),
//   										seconds: jsii.Number(123),
//   									},
//   									setVariable: &setVariableProperty{
//   										value: jsii.String("value"),
//   										variableName: jsii.String("variableName"),
//   									},
//   									sns: &snsProperty{
//   										targetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									sqs: &sqsProperty{
//   										queueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										useBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							condition: jsii.String("condition"),
//   						},
//   					},
//   				},
//   				onExit: &onExitProperty{
//   					events: []interface{}{
//   						&eventProperty{
//   							eventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							actions: []interface{}{
//   								&actionProperty{
//   									clearTimer: &clearTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									dynamoDb: &dynamoDBProperty{
//   										hashKeyField: jsii.String("hashKeyField"),
//   										hashKeyValue: jsii.String("hashKeyValue"),
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										hashKeyType: jsii.String("hashKeyType"),
//   										operation: jsii.String("operation"),
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										payloadField: jsii.String("payloadField"),
//   										rangeKeyField: jsii.String("rangeKeyField"),
//   										rangeKeyType: jsii.String("rangeKeyType"),
//   										rangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									dynamoDBv2: &dynamoDBv2Property{
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									firehose: &firehoseProperty{
//   										deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										separator: jsii.String("separator"),
//   									},
//   									iotEvents: &iotEventsProperty{
//   										inputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									iotSiteWise: &iotSiteWiseProperty{
//   										propertyValue: &assetPropertyValueProperty{
//   											value: &assetPropertyVariantProperty{
//   												booleanValue: jsii.String("booleanValue"),
//   												doubleValue: jsii.String("doubleValue"),
//   												integerValue: jsii.String("integerValue"),
//   												stringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											quality: jsii.String("quality"),
//   											timestamp: &assetPropertyTimestampProperty{
//   												timeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												offsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										assetId: jsii.String("assetId"),
//   										entryId: jsii.String("entryId"),
//   										propertyAlias: jsii.String("propertyAlias"),
//   										propertyId: jsii.String("propertyId"),
//   									},
//   									iotTopicPublish: &iotTopicPublishProperty{
//   										mqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									lambda: &lambdaProperty{
//   										functionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									resetTimer: &resetTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									setTimer: &setTimerProperty{
//   										timerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										durationExpression: jsii.String("durationExpression"),
//   										seconds: jsii.Number(123),
//   									},
//   									setVariable: &setVariableProperty{
//   										value: jsii.String("value"),
//   										variableName: jsii.String("variableName"),
//   									},
//   									sns: &snsProperty{
//   										targetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									sqs: &sqsProperty{
//   										queueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										useBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							condition: jsii.String("condition"),
//   						},
//   					},
//   				},
//   				onInput: &onInputProperty{
//   					events: []interface{}{
//   						&eventProperty{
//   							eventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							actions: []interface{}{
//   								&actionProperty{
//   									clearTimer: &clearTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									dynamoDb: &dynamoDBProperty{
//   										hashKeyField: jsii.String("hashKeyField"),
//   										hashKeyValue: jsii.String("hashKeyValue"),
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										hashKeyType: jsii.String("hashKeyType"),
//   										operation: jsii.String("operation"),
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										payloadField: jsii.String("payloadField"),
//   										rangeKeyField: jsii.String("rangeKeyField"),
//   										rangeKeyType: jsii.String("rangeKeyType"),
//   										rangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									dynamoDBv2: &dynamoDBv2Property{
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									firehose: &firehoseProperty{
//   										deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										separator: jsii.String("separator"),
//   									},
//   									iotEvents: &iotEventsProperty{
//   										inputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									iotSiteWise: &iotSiteWiseProperty{
//   										propertyValue: &assetPropertyValueProperty{
//   											value: &assetPropertyVariantProperty{
//   												booleanValue: jsii.String("booleanValue"),
//   												doubleValue: jsii.String("doubleValue"),
//   												integerValue: jsii.String("integerValue"),
//   												stringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											quality: jsii.String("quality"),
//   											timestamp: &assetPropertyTimestampProperty{
//   												timeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												offsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										assetId: jsii.String("assetId"),
//   										entryId: jsii.String("entryId"),
//   										propertyAlias: jsii.String("propertyAlias"),
//   										propertyId: jsii.String("propertyId"),
//   									},
//   									iotTopicPublish: &iotTopicPublishProperty{
//   										mqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									lambda: &lambdaProperty{
//   										functionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									resetTimer: &resetTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									setTimer: &setTimerProperty{
//   										timerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										durationExpression: jsii.String("durationExpression"),
//   										seconds: jsii.Number(123),
//   									},
//   									setVariable: &setVariableProperty{
//   										value: jsii.String("value"),
//   										variableName: jsii.String("variableName"),
//   									},
//   									sns: &snsProperty{
//   										targetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									sqs: &sqsProperty{
//   										queueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										useBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							condition: jsii.String("condition"),
//   						},
//   					},
//   					transitionEvents: []interface{}{
//   						&transitionEventProperty{
//   							condition: jsii.String("condition"),
//   							eventName: jsii.String("eventName"),
//   							nextState: jsii.String("nextState"),
//
//   							// the properties below are optional
//   							actions: []interface{}{
//   								&actionProperty{
//   									clearTimer: &clearTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									dynamoDb: &dynamoDBProperty{
//   										hashKeyField: jsii.String("hashKeyField"),
//   										hashKeyValue: jsii.String("hashKeyValue"),
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										hashKeyType: jsii.String("hashKeyType"),
//   										operation: jsii.String("operation"),
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										payloadField: jsii.String("payloadField"),
//   										rangeKeyField: jsii.String("rangeKeyField"),
//   										rangeKeyType: jsii.String("rangeKeyType"),
//   										rangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									dynamoDBv2: &dynamoDBv2Property{
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									firehose: &firehoseProperty{
//   										deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										separator: jsii.String("separator"),
//   									},
//   									iotEvents: &iotEventsProperty{
//   										inputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									iotSiteWise: &iotSiteWiseProperty{
//   										propertyValue: &assetPropertyValueProperty{
//   											value: &assetPropertyVariantProperty{
//   												booleanValue: jsii.String("booleanValue"),
//   												doubleValue: jsii.String("doubleValue"),
//   												integerValue: jsii.String("integerValue"),
//   												stringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											quality: jsii.String("quality"),
//   											timestamp: &assetPropertyTimestampProperty{
//   												timeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												offsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										assetId: jsii.String("assetId"),
//   										entryId: jsii.String("entryId"),
//   										propertyAlias: jsii.String("propertyAlias"),
//   										propertyId: jsii.String("propertyId"),
//   									},
//   									iotTopicPublish: &iotTopicPublishProperty{
//   										mqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									lambda: &lambdaProperty{
//   										functionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									resetTimer: &resetTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									setTimer: &setTimerProperty{
//   										timerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										durationExpression: jsii.String("durationExpression"),
//   										seconds: jsii.Number(123),
//   									},
//   									setVariable: &setVariableProperty{
//   										value: jsii.String("value"),
//   										variableName: jsii.String("variableName"),
//   									},
//   									sns: &snsProperty{
//   										targetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									sqs: &sqsProperty{
//   										queueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										useBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	detectorModelDescription: jsii.String("detectorModelDescription"),
//   	detectorModelName: jsii.String("detectorModelName"),
//   	evaluationMethod: jsii.String("evaluationMethod"),
//   	key: jsii.String("key"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDetectorModel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Information that defines how a detector operates.
	DetectorModelDefinition() interface{}
	SetDetectorModelDefinition(val interface{})
	// A brief description of the detector model.
	DetectorModelDescription() *string
	SetDetectorModelDescription(val *string)
	// The name of the detector model.
	DetectorModelName() *string
	SetDetectorModelName(val *string)
	// Information about the order in which events are evaluated and how actions are executed.
	EvaluationMethod() *string
	SetEvaluationMethod(val *string)
	// The value used to identify a detector instance.
	//
	// When a device or system sends input, a new detector instance with a unique key value is created. AWS IoT Events can continue to route input to its corresponding detector instance based on this identifying information.
	//
	// This parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct detector instance, the device must send a message payload that contains the same attribute-value.
	Key() *string
	SetKey(val *string)
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
	// The ARN of the role that grants permission to AWS IoT Events to perform its operations.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
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

// The jsii proxy struct for CfnDetectorModel
type jsiiProxy_CfnDetectorModel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDetectorModel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) DetectorModelDefinition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectorModelDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) DetectorModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorModelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) DetectorModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) EvaluationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTEvents::DetectorModel`.
func NewCfnDetectorModel(scope constructs.Construct, id *string, props *CfnDetectorModelProps) CfnDetectorModel {
	_init_.Initialize()

	j := jsiiProxy_CfnDetectorModel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotevents.CfnDetectorModel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTEvents::DetectorModel`.
func NewCfnDetectorModel_Override(c CfnDetectorModel, scope constructs.Construct, id *string, props *CfnDetectorModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotevents.CfnDetectorModel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDetectorModel) SetDetectorModelDefinition(val interface{}) {
	_jsii_.Set(
		j,
		"detectorModelDefinition",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel) SetDetectorModelDescription(val *string) {
	_jsii_.Set(
		j,
		"detectorModelDescription",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel) SetDetectorModelName(val *string) {
	_jsii_.Set(
		j,
		"detectorModelName",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel) SetEvaluationMethod(val *string) {
	_jsii_.Set(
		j,
		"evaluationMethod",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDetectorModel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotevents.CfnDetectorModel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDetectorModel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotevents.CfnDetectorModel",
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
func CfnDetectorModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotevents.CfnDetectorModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDetectorModel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotevents.CfnDetectorModel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDetectorModel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDetectorModel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDetectorModel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDetectorModel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDetectorModel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDetectorModel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDetectorModel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDetectorModel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDetectorModel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDetectorModel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An action to be performed when the `condition` is TRUE.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	clearTimer: &clearTimerProperty{
//   		timerName: jsii.String("timerName"),
//   	},
//   	dynamoDb: &dynamoDBProperty{
//   		hashKeyField: jsii.String("hashKeyField"),
//   		hashKeyValue: jsii.String("hashKeyValue"),
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		hashKeyType: jsii.String("hashKeyType"),
//   		operation: jsii.String("operation"),
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   		payloadField: jsii.String("payloadField"),
//   		rangeKeyField: jsii.String("rangeKeyField"),
//   		rangeKeyType: jsii.String("rangeKeyType"),
//   		rangeKeyValue: jsii.String("rangeKeyValue"),
//   	},
//   	dynamoDBv2: &dynamoDBv2Property{
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	firehose: &firehoseProperty{
//   		deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   		separator: jsii.String("separator"),
//   	},
//   	iotEvents: &iotEventsProperty{
//   		inputName: jsii.String("inputName"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	iotSiteWise: &iotSiteWiseProperty{
//   		propertyValue: &assetPropertyValueProperty{
//   			value: &assetPropertyVariantProperty{
//   				booleanValue: jsii.String("booleanValue"),
//   				doubleValue: jsii.String("doubleValue"),
//   				integerValue: jsii.String("integerValue"),
//   				stringValue: jsii.String("stringValue"),
//   			},
//
//   			// the properties below are optional
//   			quality: jsii.String("quality"),
//   			timestamp: &assetPropertyTimestampProperty{
//   				timeInSeconds: jsii.String("timeInSeconds"),
//
//   				// the properties below are optional
//   				offsetInNanos: jsii.String("offsetInNanos"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		assetId: jsii.String("assetId"),
//   		entryId: jsii.String("entryId"),
//   		propertyAlias: jsii.String("propertyAlias"),
//   		propertyId: jsii.String("propertyId"),
//   	},
//   	iotTopicPublish: &iotTopicPublishProperty{
//   		mqttTopic: jsii.String("mqttTopic"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	lambda: &lambdaProperty{
//   		functionArn: jsii.String("functionArn"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	resetTimer: &resetTimerProperty{
//   		timerName: jsii.String("timerName"),
//   	},
//   	setTimer: &setTimerProperty{
//   		timerName: jsii.String("timerName"),
//
//   		// the properties below are optional
//   		durationExpression: jsii.String("durationExpression"),
//   		seconds: jsii.Number(123),
//   	},
//   	setVariable: &setVariableProperty{
//   		value: jsii.String("value"),
//   		variableName: jsii.String("variableName"),
//   	},
//   	sns: &snsProperty{
//   		targetArn: jsii.String("targetArn"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	sqs: &sqsProperty{
//   		queueUrl: jsii.String("queueUrl"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   		useBase64: jsii.Boolean(false),
//   	},
//   }
//
type CfnDetectorModel_ActionProperty struct {
	// Information needed to clear the timer.
	ClearTimer interface{} `field:"optional" json:"clearTimer" yaml:"clearTimer"`
	// Writes to the DynamoDB table that you created.
	//
	// The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *AWS IoT Events Developer Guide* .
	DynamoDb interface{} `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Writes to the DynamoDB table that you created.
	//
	// The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *AWS IoT Events Developer Guide* .
	DynamoDBv2 interface{} `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// Sends AWS IoT Events input, which passes information about the detector model instance and the event that triggered the action.
	IotEvents interface{} `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// Sends information about the detector model instance and the event that triggered the action to an asset property in AWS IoT SiteWise .
	IotSiteWise interface{} `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Publishes an MQTT message with the given topic to the AWS IoT message broker.
	IotTopicPublish interface{} `field:"optional" json:"iotTopicPublish" yaml:"iotTopicPublish"`
	// Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Information needed to reset the timer.
	ResetTimer interface{} `field:"optional" json:"resetTimer" yaml:"resetTimer"`
	// Information needed to set the timer.
	SetTimer interface{} `field:"optional" json:"setTimer" yaml:"setTimer"`
	// Sets a variable to a specified value.
	SetVariable interface{} `field:"optional" json:"setVariable" yaml:"setVariable"`
	// Sends an Amazon SNS message.
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
	// Sends an Amazon SNS message.
	Sqs interface{} `field:"optional" json:"sqs" yaml:"sqs"`
}

// A structure that contains timestamp information. For more information, see [TimeInNanos](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_TimeInNanos.html) in the *AWS IoT SiteWise API Reference* .
//
// You must use expressions for all parameters in `AssetPropertyTimestamp` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `timeInSeconds` parameter can be `'1586400675'` .
// - For references, you must specify either variables or input values. For example, the value for the `offsetInNanos` parameter can be `$variable.time` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `timeInSeconds` parameter uses a substitution template.
//
// `'${$input.TemperatureInput.sensorData.timestamp / 1000}'`
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyTimestampProperty := &assetPropertyTimestampProperty{
//   	timeInSeconds: jsii.String("timeInSeconds"),
//
//   	// the properties below are optional
//   	offsetInNanos: jsii.String("offsetInNanos"),
//   }
//
type CfnDetectorModel_AssetPropertyTimestampProperty struct {
	// The timestamp, in seconds, in the Unix epoch format.
	//
	// The valid range is between 1-31556889864403199.
	TimeInSeconds *string `field:"required" json:"timeInSeconds" yaml:"timeInSeconds"`
	// The nanosecond offset converted from `timeInSeconds` .
	//
	// The valid range is between 0-999999999.
	OffsetInNanos *string `field:"optional" json:"offsetInNanos" yaml:"offsetInNanos"`
}

// A structure that contains value information. For more information, see [AssetPropertyValue](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_AssetPropertyValue.html) in the *AWS IoT SiteWise API Reference* .
//
// You must use expressions for all parameters in `AssetPropertyValue` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `quality` parameter can be `'GOOD'` .
// - For references, you must specify either variables or input values. For example, the value for the `quality` parameter can be `$input.TemperatureInput.sensorData.quality` .
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyValueProperty := &assetPropertyValueProperty{
//   	value: &assetPropertyVariantProperty{
//   		booleanValue: jsii.String("booleanValue"),
//   		doubleValue: jsii.String("doubleValue"),
//   		integerValue: jsii.String("integerValue"),
//   		stringValue: jsii.String("stringValue"),
//   	},
//
//   	// the properties below are optional
//   	quality: jsii.String("quality"),
//   	timestamp: &assetPropertyTimestampProperty{
//   		timeInSeconds: jsii.String("timeInSeconds"),
//
//   		// the properties below are optional
//   		offsetInNanos: jsii.String("offsetInNanos"),
//   	},
//   }
//
type CfnDetectorModel_AssetPropertyValueProperty struct {
	// The value to send to an asset property.
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// The quality of the asset property value.
	//
	// The value must be `'GOOD'` , `'BAD'` , or `'UNCERTAIN'` .
	Quality *string `field:"optional" json:"quality" yaml:"quality"`
	// The timestamp associated with the asset property value.
	//
	// The default is the current event time.
	Timestamp interface{} `field:"optional" json:"timestamp" yaml:"timestamp"`
}

// A structure that contains an asset property value.
//
// For more information, see [Variant](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_Variant.html) in the *AWS IoT SiteWise API Reference* .
//
// You must use expressions for all parameters in `AssetPropertyVariant` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `integerValue` parameter can be `'100'` .
// - For references, you must specify either variables or parameters. For example, the value for the `booleanValue` parameter can be `$variable.offline` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `doubleValue` parameter uses a substitution template.
//
// `'${$input.TemperatureInput.sensorData.temperature * 6 / 5 + 32}'`
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// You must specify one of the following value types, depending on the `dataType` of the specified asset property. For more information, see [AssetProperty](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_AssetProperty.html) in the *AWS IoT SiteWise API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyVariantProperty := &assetPropertyVariantProperty{
//   	booleanValue: jsii.String("booleanValue"),
//   	doubleValue: jsii.String("doubleValue"),
//   	integerValue: jsii.String("integerValue"),
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnDetectorModel_AssetPropertyVariantProperty struct {
	// The asset property value is a Boolean value that must be `'TRUE'` or `'FALSE'` .
	//
	// You must use an expression, and the evaluated result should be a Boolean value.
	BooleanValue *string `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// The asset property value is a double.
	//
	// You must use an expression, and the evaluated result should be a double.
	DoubleValue *string `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// The asset property value is an integer.
	//
	// You must use an expression, and the evaluated result should be an integer.
	IntegerValue *string `field:"optional" json:"integerValue" yaml:"integerValue"`
	// The asset property value is a string.
	//
	// You must use an expression, and the evaluated result should be a string.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

// Information needed to clear the timer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clearTimerProperty := &clearTimerProperty{
//   	timerName: jsii.String("timerName"),
//   }
//
type CfnDetectorModel_ClearTimerProperty struct {
	// The name of the timer to clear.
	TimerName *string `field:"required" json:"timerName" yaml:"timerName"`
}

// Information that defines how a detector operates.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   detectorModelDefinitionProperty := &detectorModelDefinitionProperty{
//   	initialStateName: jsii.String("initialStateName"),
//   	states: []interface{}{
//   		&stateProperty{
//   			stateName: jsii.String("stateName"),
//
//   			// the properties below are optional
//   			onEnter: &onEnterProperty{
//   				events: []interface{}{
//   					&eventProperty{
//   						eventName: jsii.String("eventName"),
//
//   						// the properties below are optional
//   						actions: []interface{}{
//   							&actionProperty{
//   								clearTimer: &clearTimerProperty{
//   									timerName: jsii.String("timerName"),
//   								},
//   								dynamoDb: &dynamoDBProperty{
//   									hashKeyField: jsii.String("hashKeyField"),
//   									hashKeyValue: jsii.String("hashKeyValue"),
//   									tableName: jsii.String("tableName"),
//
//   									// the properties below are optional
//   									hashKeyType: jsii.String("hashKeyType"),
//   									operation: jsii.String("operation"),
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   									payloadField: jsii.String("payloadField"),
//   									rangeKeyField: jsii.String("rangeKeyField"),
//   									rangeKeyType: jsii.String("rangeKeyType"),
//   									rangeKeyValue: jsii.String("rangeKeyValue"),
//   								},
//   								dynamoDBv2: &dynamoDBv2Property{
//   									tableName: jsii.String("tableName"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								firehose: &firehoseProperty{
//   									deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   									separator: jsii.String("separator"),
//   								},
//   								iotEvents: &iotEventsProperty{
//   									inputName: jsii.String("inputName"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								iotSiteWise: &iotSiteWiseProperty{
//   									propertyValue: &assetPropertyValueProperty{
//   										value: &assetPropertyVariantProperty{
//   											booleanValue: jsii.String("booleanValue"),
//   											doubleValue: jsii.String("doubleValue"),
//   											integerValue: jsii.String("integerValue"),
//   											stringValue: jsii.String("stringValue"),
//   										},
//
//   										// the properties below are optional
//   										quality: jsii.String("quality"),
//   										timestamp: &assetPropertyTimestampProperty{
//   											timeInSeconds: jsii.String("timeInSeconds"),
//
//   											// the properties below are optional
//   											offsetInNanos: jsii.String("offsetInNanos"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									assetId: jsii.String("assetId"),
//   									entryId: jsii.String("entryId"),
//   									propertyAlias: jsii.String("propertyAlias"),
//   									propertyId: jsii.String("propertyId"),
//   								},
//   								iotTopicPublish: &iotTopicPublishProperty{
//   									mqttTopic: jsii.String("mqttTopic"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								lambda: &lambdaProperty{
//   									functionArn: jsii.String("functionArn"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								resetTimer: &resetTimerProperty{
//   									timerName: jsii.String("timerName"),
//   								},
//   								setTimer: &setTimerProperty{
//   									timerName: jsii.String("timerName"),
//
//   									// the properties below are optional
//   									durationExpression: jsii.String("durationExpression"),
//   									seconds: jsii.Number(123),
//   								},
//   								setVariable: &setVariableProperty{
//   									value: jsii.String("value"),
//   									variableName: jsii.String("variableName"),
//   								},
//   								sns: &snsProperty{
//   									targetArn: jsii.String("targetArn"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								sqs: &sqsProperty{
//   									queueUrl: jsii.String("queueUrl"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   									useBase64: jsii.Boolean(false),
//   								},
//   							},
//   						},
//   						condition: jsii.String("condition"),
//   					},
//   				},
//   			},
//   			onExit: &onExitProperty{
//   				events: []interface{}{
//   					&eventProperty{
//   						eventName: jsii.String("eventName"),
//
//   						// the properties below are optional
//   						actions: []interface{}{
//   							&actionProperty{
//   								clearTimer: &clearTimerProperty{
//   									timerName: jsii.String("timerName"),
//   								},
//   								dynamoDb: &dynamoDBProperty{
//   									hashKeyField: jsii.String("hashKeyField"),
//   									hashKeyValue: jsii.String("hashKeyValue"),
//   									tableName: jsii.String("tableName"),
//
//   									// the properties below are optional
//   									hashKeyType: jsii.String("hashKeyType"),
//   									operation: jsii.String("operation"),
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   									payloadField: jsii.String("payloadField"),
//   									rangeKeyField: jsii.String("rangeKeyField"),
//   									rangeKeyType: jsii.String("rangeKeyType"),
//   									rangeKeyValue: jsii.String("rangeKeyValue"),
//   								},
//   								dynamoDBv2: &dynamoDBv2Property{
//   									tableName: jsii.String("tableName"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								firehose: &firehoseProperty{
//   									deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   									separator: jsii.String("separator"),
//   								},
//   								iotEvents: &iotEventsProperty{
//   									inputName: jsii.String("inputName"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								iotSiteWise: &iotSiteWiseProperty{
//   									propertyValue: &assetPropertyValueProperty{
//   										value: &assetPropertyVariantProperty{
//   											booleanValue: jsii.String("booleanValue"),
//   											doubleValue: jsii.String("doubleValue"),
//   											integerValue: jsii.String("integerValue"),
//   											stringValue: jsii.String("stringValue"),
//   										},
//
//   										// the properties below are optional
//   										quality: jsii.String("quality"),
//   										timestamp: &assetPropertyTimestampProperty{
//   											timeInSeconds: jsii.String("timeInSeconds"),
//
//   											// the properties below are optional
//   											offsetInNanos: jsii.String("offsetInNanos"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									assetId: jsii.String("assetId"),
//   									entryId: jsii.String("entryId"),
//   									propertyAlias: jsii.String("propertyAlias"),
//   									propertyId: jsii.String("propertyId"),
//   								},
//   								iotTopicPublish: &iotTopicPublishProperty{
//   									mqttTopic: jsii.String("mqttTopic"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								lambda: &lambdaProperty{
//   									functionArn: jsii.String("functionArn"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								resetTimer: &resetTimerProperty{
//   									timerName: jsii.String("timerName"),
//   								},
//   								setTimer: &setTimerProperty{
//   									timerName: jsii.String("timerName"),
//
//   									// the properties below are optional
//   									durationExpression: jsii.String("durationExpression"),
//   									seconds: jsii.Number(123),
//   								},
//   								setVariable: &setVariableProperty{
//   									value: jsii.String("value"),
//   									variableName: jsii.String("variableName"),
//   								},
//   								sns: &snsProperty{
//   									targetArn: jsii.String("targetArn"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								sqs: &sqsProperty{
//   									queueUrl: jsii.String("queueUrl"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   									useBase64: jsii.Boolean(false),
//   								},
//   							},
//   						},
//   						condition: jsii.String("condition"),
//   					},
//   				},
//   			},
//   			onInput: &onInputProperty{
//   				events: []interface{}{
//   					&eventProperty{
//   						eventName: jsii.String("eventName"),
//
//   						// the properties below are optional
//   						actions: []interface{}{
//   							&actionProperty{
//   								clearTimer: &clearTimerProperty{
//   									timerName: jsii.String("timerName"),
//   								},
//   								dynamoDb: &dynamoDBProperty{
//   									hashKeyField: jsii.String("hashKeyField"),
//   									hashKeyValue: jsii.String("hashKeyValue"),
//   									tableName: jsii.String("tableName"),
//
//   									// the properties below are optional
//   									hashKeyType: jsii.String("hashKeyType"),
//   									operation: jsii.String("operation"),
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   									payloadField: jsii.String("payloadField"),
//   									rangeKeyField: jsii.String("rangeKeyField"),
//   									rangeKeyType: jsii.String("rangeKeyType"),
//   									rangeKeyValue: jsii.String("rangeKeyValue"),
//   								},
//   								dynamoDBv2: &dynamoDBv2Property{
//   									tableName: jsii.String("tableName"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								firehose: &firehoseProperty{
//   									deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   									separator: jsii.String("separator"),
//   								},
//   								iotEvents: &iotEventsProperty{
//   									inputName: jsii.String("inputName"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								iotSiteWise: &iotSiteWiseProperty{
//   									propertyValue: &assetPropertyValueProperty{
//   										value: &assetPropertyVariantProperty{
//   											booleanValue: jsii.String("booleanValue"),
//   											doubleValue: jsii.String("doubleValue"),
//   											integerValue: jsii.String("integerValue"),
//   											stringValue: jsii.String("stringValue"),
//   										},
//
//   										// the properties below are optional
//   										quality: jsii.String("quality"),
//   										timestamp: &assetPropertyTimestampProperty{
//   											timeInSeconds: jsii.String("timeInSeconds"),
//
//   											// the properties below are optional
//   											offsetInNanos: jsii.String("offsetInNanos"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									assetId: jsii.String("assetId"),
//   									entryId: jsii.String("entryId"),
//   									propertyAlias: jsii.String("propertyAlias"),
//   									propertyId: jsii.String("propertyId"),
//   								},
//   								iotTopicPublish: &iotTopicPublishProperty{
//   									mqttTopic: jsii.String("mqttTopic"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								lambda: &lambdaProperty{
//   									functionArn: jsii.String("functionArn"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								resetTimer: &resetTimerProperty{
//   									timerName: jsii.String("timerName"),
//   								},
//   								setTimer: &setTimerProperty{
//   									timerName: jsii.String("timerName"),
//
//   									// the properties below are optional
//   									durationExpression: jsii.String("durationExpression"),
//   									seconds: jsii.Number(123),
//   								},
//   								setVariable: &setVariableProperty{
//   									value: jsii.String("value"),
//   									variableName: jsii.String("variableName"),
//   								},
//   								sns: &snsProperty{
//   									targetArn: jsii.String("targetArn"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								sqs: &sqsProperty{
//   									queueUrl: jsii.String("queueUrl"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   									useBase64: jsii.Boolean(false),
//   								},
//   							},
//   						},
//   						condition: jsii.String("condition"),
//   					},
//   				},
//   				transitionEvents: []interface{}{
//   					&transitionEventProperty{
//   						condition: jsii.String("condition"),
//   						eventName: jsii.String("eventName"),
//   						nextState: jsii.String("nextState"),
//
//   						// the properties below are optional
//   						actions: []interface{}{
//   							&actionProperty{
//   								clearTimer: &clearTimerProperty{
//   									timerName: jsii.String("timerName"),
//   								},
//   								dynamoDb: &dynamoDBProperty{
//   									hashKeyField: jsii.String("hashKeyField"),
//   									hashKeyValue: jsii.String("hashKeyValue"),
//   									tableName: jsii.String("tableName"),
//
//   									// the properties below are optional
//   									hashKeyType: jsii.String("hashKeyType"),
//   									operation: jsii.String("operation"),
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   									payloadField: jsii.String("payloadField"),
//   									rangeKeyField: jsii.String("rangeKeyField"),
//   									rangeKeyType: jsii.String("rangeKeyType"),
//   									rangeKeyValue: jsii.String("rangeKeyValue"),
//   								},
//   								dynamoDBv2: &dynamoDBv2Property{
//   									tableName: jsii.String("tableName"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								firehose: &firehoseProperty{
//   									deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   									separator: jsii.String("separator"),
//   								},
//   								iotEvents: &iotEventsProperty{
//   									inputName: jsii.String("inputName"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								iotSiteWise: &iotSiteWiseProperty{
//   									propertyValue: &assetPropertyValueProperty{
//   										value: &assetPropertyVariantProperty{
//   											booleanValue: jsii.String("booleanValue"),
//   											doubleValue: jsii.String("doubleValue"),
//   											integerValue: jsii.String("integerValue"),
//   											stringValue: jsii.String("stringValue"),
//   										},
//
//   										// the properties below are optional
//   										quality: jsii.String("quality"),
//   										timestamp: &assetPropertyTimestampProperty{
//   											timeInSeconds: jsii.String("timeInSeconds"),
//
//   											// the properties below are optional
//   											offsetInNanos: jsii.String("offsetInNanos"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									assetId: jsii.String("assetId"),
//   									entryId: jsii.String("entryId"),
//   									propertyAlias: jsii.String("propertyAlias"),
//   									propertyId: jsii.String("propertyId"),
//   								},
//   								iotTopicPublish: &iotTopicPublishProperty{
//   									mqttTopic: jsii.String("mqttTopic"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								lambda: &lambdaProperty{
//   									functionArn: jsii.String("functionArn"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								resetTimer: &resetTimerProperty{
//   									timerName: jsii.String("timerName"),
//   								},
//   								setTimer: &setTimerProperty{
//   									timerName: jsii.String("timerName"),
//
//   									// the properties below are optional
//   									durationExpression: jsii.String("durationExpression"),
//   									seconds: jsii.Number(123),
//   								},
//   								setVariable: &setVariableProperty{
//   									value: jsii.String("value"),
//   									variableName: jsii.String("variableName"),
//   								},
//   								sns: &snsProperty{
//   									targetArn: jsii.String("targetArn"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   								},
//   								sqs: &sqsProperty{
//   									queueUrl: jsii.String("queueUrl"),
//
//   									// the properties below are optional
//   									payload: &payloadProperty{
//   										contentExpression: jsii.String("contentExpression"),
//   										type: jsii.String("type"),
//   									},
//   									useBase64: jsii.Boolean(false),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDetectorModel_DetectorModelDefinitionProperty struct {
	// The state that is entered at the creation of each detector (instance).
	InitialStateName *string `field:"required" json:"initialStateName" yaml:"initialStateName"`
	// Information about the states of the detector.
	States interface{} `field:"required" json:"states" yaml:"states"`
}

// Defines an action to write to the Amazon DynamoDB table that you created.
//
// The standard action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify.
//
// You must use expressions for all parameters in `DynamoDBAction` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `hashKeyType` parameter can be `'STRING'` .
// - For references, you must specify either variables or input values. For example, the value for the `hashKeyField` parameter can be `$input.GreenhouseInput.name` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `hashKeyValue` parameter uses a substitution template.
//
// `'${$input.GreenhouseInput.temperature * 6 / 5 + 32} in Fahrenheit'`
// - For a string concatenation, you must use `+` . A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `tableName` parameter uses a string concatenation.
//
// `'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// If the defined payload type is a string, `DynamoDBAction` writes non-JSON data to the DynamoDB table as binary data. The DynamoDB console displays the data as Base64-encoded text. The value for the `payloadField` parameter is `<payload-field>_raw` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDBProperty := &dynamoDBProperty{
//   	hashKeyField: jsii.String("hashKeyField"),
//   	hashKeyValue: jsii.String("hashKeyValue"),
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	hashKeyType: jsii.String("hashKeyType"),
//   	operation: jsii.String("operation"),
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   	payloadField: jsii.String("payloadField"),
//   	rangeKeyField: jsii.String("rangeKeyField"),
//   	rangeKeyType: jsii.String("rangeKeyType"),
//   	rangeKeyValue: jsii.String("rangeKeyValue"),
//   }
//
type CfnDetectorModel_DynamoDBProperty struct {
	// The name of the hash key (also called the partition key).
	//
	// The `hashKeyField` value must match the partition key of the target DynamoDB table.
	HashKeyField *string `field:"required" json:"hashKeyField" yaml:"hashKeyField"`
	// The value of the hash key (also called the partition key).
	HashKeyValue *string `field:"required" json:"hashKeyValue" yaml:"hashKeyValue"`
	// The name of the DynamoDB table.
	//
	// The `tableName` value must match the table name of the target DynamoDB table.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The data type for the hash key (also called the partition key). You can specify the following values:.
	//
	// - `'STRING'` - The hash key is a string.
	// - `'NUMBER'` - The hash key is a number.
	//
	// If you don't specify `hashKeyType` , the default value is `'STRING'` .
	HashKeyType *string `field:"optional" json:"hashKeyType" yaml:"hashKeyType"`
	// The type of operation to perform. You can specify the following values:.
	//
	// - `'INSERT'` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.
	// - `'UPDATE'` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.
	// - `'DELETE'` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.
	//
	// If you don't specify this parameter, AWS IoT Events triggers the `'INSERT'` operation.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Information needed to configure the payload.
	//
	// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use `contentExpression` .
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
	// The name of the DynamoDB column that receives the action payload.
	//
	// If you don't specify this parameter, the name of the DynamoDB column is `payload` .
	PayloadField *string `field:"optional" json:"payloadField" yaml:"payloadField"`
	// The name of the range key (also called the sort key).
	//
	// The `rangeKeyField` value must match the sort key of the target DynamoDB table.
	RangeKeyField *string `field:"optional" json:"rangeKeyField" yaml:"rangeKeyField"`
	// The data type for the range key (also called the sort key), You can specify the following values:.
	//
	// - `'STRING'` - The range key is a string.
	// - `'NUMBER'` - The range key is number.
	//
	// If you don't specify `rangeKeyField` , the default value is `'STRING'` .
	RangeKeyType *string `field:"optional" json:"rangeKeyType" yaml:"rangeKeyType"`
	// The value of the range key (also called the sort key).
	RangeKeyValue *string `field:"optional" json:"rangeKeyValue" yaml:"rangeKeyValue"`
}

// Defines an action to write to the Amazon DynamoDB table that you created.
//
// The default action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.
//
// You must use expressions for all parameters in `DynamoDBv2Action` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `tableName` parameter can be `'GreenhouseTemperatureTable'` .
// - For references, you must specify either variables or input values. For example, the value for the `tableName` parameter can be `$variable.ddbtableName` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `contentExpression` parameter in `Payload` uses a substitution template.
//
// `'{\"sensorID\": \"${$input.GreenhouseInput.sensor_id}\", \"temperature\": \"${$input.GreenhouseInput.temperature * 9 / 5 + 32}\"}'`
// - For a string concatenation, you must use `+` . A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `tableName` parameter uses a string concatenation.
//
// `'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// The value for the `type` parameter in `Payload` must be `JSON` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDBv2Property := &dynamoDBv2Property{
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnDetectorModel_DynamoDBv2Property struct {
	// The name of the DynamoDB table.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Information needed to configure the payload.
	//
	// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use `contentExpression` .
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

// Specifies the `actions` to be performed when the `condition` evaluates to TRUE.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventProperty := &eventProperty{
//   	eventName: jsii.String("eventName"),
//
//   	// the properties below are optional
//   	actions: []interface{}{
//   		&actionProperty{
//   			clearTimer: &clearTimerProperty{
//   				timerName: jsii.String("timerName"),
//   			},
//   			dynamoDb: &dynamoDBProperty{
//   				hashKeyField: jsii.String("hashKeyField"),
//   				hashKeyValue: jsii.String("hashKeyValue"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				hashKeyType: jsii.String("hashKeyType"),
//   				operation: jsii.String("operation"),
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   				payloadField: jsii.String("payloadField"),
//   				rangeKeyField: jsii.String("rangeKeyField"),
//   				rangeKeyType: jsii.String("rangeKeyType"),
//   				rangeKeyValue: jsii.String("rangeKeyValue"),
//   			},
//   			dynamoDBv2: &dynamoDBv2Property{
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			firehose: &firehoseProperty{
//   				deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   				separator: jsii.String("separator"),
//   			},
//   			iotEvents: &iotEventsProperty{
//   				inputName: jsii.String("inputName"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			iotSiteWise: &iotSiteWiseProperty{
//   				propertyValue: &assetPropertyValueProperty{
//   					value: &assetPropertyVariantProperty{
//   						booleanValue: jsii.String("booleanValue"),
//   						doubleValue: jsii.String("doubleValue"),
//   						integerValue: jsii.String("integerValue"),
//   						stringValue: jsii.String("stringValue"),
//   					},
//
//   					// the properties below are optional
//   					quality: jsii.String("quality"),
//   					timestamp: &assetPropertyTimestampProperty{
//   						timeInSeconds: jsii.String("timeInSeconds"),
//
//   						// the properties below are optional
//   						offsetInNanos: jsii.String("offsetInNanos"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				assetId: jsii.String("assetId"),
//   				entryId: jsii.String("entryId"),
//   				propertyAlias: jsii.String("propertyAlias"),
//   				propertyId: jsii.String("propertyId"),
//   			},
//   			iotTopicPublish: &iotTopicPublishProperty{
//   				mqttTopic: jsii.String("mqttTopic"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			lambda: &lambdaProperty{
//   				functionArn: jsii.String("functionArn"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			resetTimer: &resetTimerProperty{
//   				timerName: jsii.String("timerName"),
//   			},
//   			setTimer: &setTimerProperty{
//   				timerName: jsii.String("timerName"),
//
//   				// the properties below are optional
//   				durationExpression: jsii.String("durationExpression"),
//   				seconds: jsii.Number(123),
//   			},
//   			setVariable: &setVariableProperty{
//   				value: jsii.String("value"),
//   				variableName: jsii.String("variableName"),
//   			},
//   			sns: &snsProperty{
//   				targetArn: jsii.String("targetArn"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			sqs: &sqsProperty{
//   				queueUrl: jsii.String("queueUrl"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   				useBase64: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	condition: jsii.String("condition"),
//   }
//
type CfnDetectorModel_EventProperty struct {
	// The name of the event.
	EventName *string `field:"required" json:"eventName" yaml:"eventName"`
	// The actions to be performed.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// Optional.
	//
	// The Boolean expression that, when TRUE, causes the `actions` to be performed. If not present, the actions are performed (=TRUE). If the expression result is not a Boolean value, the actions are not performed (=FALSE).
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
}

// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseProperty := &firehoseProperty{
//   	deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   	separator: jsii.String("separator"),
//   }
//
type CfnDetectorModel_FirehoseProperty struct {
	// The name of the Kinesis Data Firehose delivery stream where the data is written.
	DeliveryStreamName *string `field:"required" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// You can configure the action payload when you send a message to an Amazon Kinesis Data Firehose delivery stream.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
	// A character separator that is used to separate records written to the Kinesis Data Firehose delivery stream.
	//
	// Valid values are: '\n' (newline), '\t' (tab), '\r\n' (Windows newline), ',' (comma).
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
}

// Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotEventsProperty := &iotEventsProperty{
//   	inputName: jsii.String("inputName"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnDetectorModel_IotEventsProperty struct {
	// The name of the AWS IoT Events input where the data is sent.
	InputName *string `field:"required" json:"inputName" yaml:"inputName"`
	// You can configure the action payload when you send a message to an AWS IoT Events input.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

// Sends information about the detector model instance and the event that triggered the action to a specified asset property in AWS IoT SiteWise .
//
// You must use expressions for all parameters in `IotSiteWiseAction` . The expressions accept literals, operators, functions, references, and substitutions templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `propertyAlias` parameter can be `'/company/windfarm/3/turbine/7/temperature'` .
// - For references, you must specify either variables or input values. For example, the value for the `assetId` parameter can be `$input.TurbineInput.assetId1` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `propertyAlias` parameter uses a substitution template.
//
// `'company/windfarm/${$input.TemperatureInput.sensorData.windfarmID}/turbine/ ${$input.TemperatureInput.sensorData.turbineID}/temperature'`
//
// You must specify either `propertyAlias` or both `assetId` and `propertyId` to identify the target asset property in AWS IoT SiteWise .
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotSiteWiseProperty := &iotSiteWiseProperty{
//   	propertyValue: &assetPropertyValueProperty{
//   		value: &assetPropertyVariantProperty{
//   			booleanValue: jsii.String("booleanValue"),
//   			doubleValue: jsii.String("doubleValue"),
//   			integerValue: jsii.String("integerValue"),
//   			stringValue: jsii.String("stringValue"),
//   		},
//
//   		// the properties below are optional
//   		quality: jsii.String("quality"),
//   		timestamp: &assetPropertyTimestampProperty{
//   			timeInSeconds: jsii.String("timeInSeconds"),
//
//   			// the properties below are optional
//   			offsetInNanos: jsii.String("offsetInNanos"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	assetId: jsii.String("assetId"),
//   	entryId: jsii.String("entryId"),
//   	propertyAlias: jsii.String("propertyAlias"),
//   	propertyId: jsii.String("propertyId"),
//   }
//
type CfnDetectorModel_IotSiteWiseProperty struct {
	// The value to send to the asset property.
	//
	// This value contains timestamp, quality, and value (TQV) information.
	PropertyValue interface{} `field:"required" json:"propertyValue" yaml:"propertyValue"`
	// The ID of the asset that has the specified property.
	AssetId *string `field:"optional" json:"assetId" yaml:"assetId"`
	// A unique identifier for this entry.
	//
	// You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier.
	EntryId *string `field:"optional" json:"entryId" yaml:"entryId"`
	// The alias of the asset property.
	PropertyAlias *string `field:"optional" json:"propertyAlias" yaml:"propertyAlias"`
	// The ID of the asset property.
	PropertyId *string `field:"optional" json:"propertyId" yaml:"propertyId"`
}

// Information required to publish the MQTT message through the AWS IoT message broker.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotTopicPublishProperty := &iotTopicPublishProperty{
//   	mqttTopic: jsii.String("mqttTopic"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnDetectorModel_IotTopicPublishProperty struct {
	// The MQTT topic of the message.
	//
	// You can use a string expression that includes variables ( `$variable.<variable-name>` ) and input values ( `$input.<input-name>.<path-to-datum>` ) as the topic string.
	MqttTopic *string `field:"required" json:"mqttTopic" yaml:"mqttTopic"`
	// You can configure the action payload when you publish a message to an AWS IoT Core topic.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

// Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaProperty := &lambdaProperty{
//   	functionArn: jsii.String("functionArn"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnDetectorModel_LambdaProperty struct {
	// The ARN of the Lambda function that is executed.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// You can configure the action payload when you send a message to a Lambda function.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

// When entering this state, perform these `actions` if the `condition` is TRUE.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onEnterProperty := &onEnterProperty{
//   	events: []interface{}{
//   		&eventProperty{
//   			eventName: jsii.String("eventName"),
//
//   			// the properties below are optional
//   			actions: []interface{}{
//   				&actionProperty{
//   					clearTimer: &clearTimerProperty{
//   						timerName: jsii.String("timerName"),
//   					},
//   					dynamoDb: &dynamoDBProperty{
//   						hashKeyField: jsii.String("hashKeyField"),
//   						hashKeyValue: jsii.String("hashKeyValue"),
//   						tableName: jsii.String("tableName"),
//
//   						// the properties below are optional
//   						hashKeyType: jsii.String("hashKeyType"),
//   						operation: jsii.String("operation"),
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						payloadField: jsii.String("payloadField"),
//   						rangeKeyField: jsii.String("rangeKeyField"),
//   						rangeKeyType: jsii.String("rangeKeyType"),
//   						rangeKeyValue: jsii.String("rangeKeyValue"),
//   					},
//   					dynamoDBv2: &dynamoDBv2Property{
//   						tableName: jsii.String("tableName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					firehose: &firehoseProperty{
//   						deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						separator: jsii.String("separator"),
//   					},
//   					iotEvents: &iotEventsProperty{
//   						inputName: jsii.String("inputName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					iotSiteWise: &iotSiteWiseProperty{
//   						propertyValue: &assetPropertyValueProperty{
//   							value: &assetPropertyVariantProperty{
//   								booleanValue: jsii.String("booleanValue"),
//   								doubleValue: jsii.String("doubleValue"),
//   								integerValue: jsii.String("integerValue"),
//   								stringValue: jsii.String("stringValue"),
//   							},
//
//   							// the properties below are optional
//   							quality: jsii.String("quality"),
//   							timestamp: &assetPropertyTimestampProperty{
//   								timeInSeconds: jsii.String("timeInSeconds"),
//
//   								// the properties below are optional
//   								offsetInNanos: jsii.String("offsetInNanos"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						assetId: jsii.String("assetId"),
//   						entryId: jsii.String("entryId"),
//   						propertyAlias: jsii.String("propertyAlias"),
//   						propertyId: jsii.String("propertyId"),
//   					},
//   					iotTopicPublish: &iotTopicPublishProperty{
//   						mqttTopic: jsii.String("mqttTopic"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					lambda: &lambdaProperty{
//   						functionArn: jsii.String("functionArn"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					resetTimer: &resetTimerProperty{
//   						timerName: jsii.String("timerName"),
//   					},
//   					setTimer: &setTimerProperty{
//   						timerName: jsii.String("timerName"),
//
//   						// the properties below are optional
//   						durationExpression: jsii.String("durationExpression"),
//   						seconds: jsii.Number(123),
//   					},
//   					setVariable: &setVariableProperty{
//   						value: jsii.String("value"),
//   						variableName: jsii.String("variableName"),
//   					},
//   					sns: &snsProperty{
//   						targetArn: jsii.String("targetArn"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					sqs: &sqsProperty{
//   						queueUrl: jsii.String("queueUrl"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						useBase64: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			condition: jsii.String("condition"),
//   		},
//   	},
//   }
//
type CfnDetectorModel_OnEnterProperty struct {
	// Specifies the actions that are performed when the state is entered and the `condition` is `TRUE` .
	Events interface{} `field:"optional" json:"events" yaml:"events"`
}

// When exiting this state, perform these `actions` if the specified `condition` is `TRUE` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onExitProperty := &onExitProperty{
//   	events: []interface{}{
//   		&eventProperty{
//   			eventName: jsii.String("eventName"),
//
//   			// the properties below are optional
//   			actions: []interface{}{
//   				&actionProperty{
//   					clearTimer: &clearTimerProperty{
//   						timerName: jsii.String("timerName"),
//   					},
//   					dynamoDb: &dynamoDBProperty{
//   						hashKeyField: jsii.String("hashKeyField"),
//   						hashKeyValue: jsii.String("hashKeyValue"),
//   						tableName: jsii.String("tableName"),
//
//   						// the properties below are optional
//   						hashKeyType: jsii.String("hashKeyType"),
//   						operation: jsii.String("operation"),
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						payloadField: jsii.String("payloadField"),
//   						rangeKeyField: jsii.String("rangeKeyField"),
//   						rangeKeyType: jsii.String("rangeKeyType"),
//   						rangeKeyValue: jsii.String("rangeKeyValue"),
//   					},
//   					dynamoDBv2: &dynamoDBv2Property{
//   						tableName: jsii.String("tableName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					firehose: &firehoseProperty{
//   						deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						separator: jsii.String("separator"),
//   					},
//   					iotEvents: &iotEventsProperty{
//   						inputName: jsii.String("inputName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					iotSiteWise: &iotSiteWiseProperty{
//   						propertyValue: &assetPropertyValueProperty{
//   							value: &assetPropertyVariantProperty{
//   								booleanValue: jsii.String("booleanValue"),
//   								doubleValue: jsii.String("doubleValue"),
//   								integerValue: jsii.String("integerValue"),
//   								stringValue: jsii.String("stringValue"),
//   							},
//
//   							// the properties below are optional
//   							quality: jsii.String("quality"),
//   							timestamp: &assetPropertyTimestampProperty{
//   								timeInSeconds: jsii.String("timeInSeconds"),
//
//   								// the properties below are optional
//   								offsetInNanos: jsii.String("offsetInNanos"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						assetId: jsii.String("assetId"),
//   						entryId: jsii.String("entryId"),
//   						propertyAlias: jsii.String("propertyAlias"),
//   						propertyId: jsii.String("propertyId"),
//   					},
//   					iotTopicPublish: &iotTopicPublishProperty{
//   						mqttTopic: jsii.String("mqttTopic"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					lambda: &lambdaProperty{
//   						functionArn: jsii.String("functionArn"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					resetTimer: &resetTimerProperty{
//   						timerName: jsii.String("timerName"),
//   					},
//   					setTimer: &setTimerProperty{
//   						timerName: jsii.String("timerName"),
//
//   						// the properties below are optional
//   						durationExpression: jsii.String("durationExpression"),
//   						seconds: jsii.Number(123),
//   					},
//   					setVariable: &setVariableProperty{
//   						value: jsii.String("value"),
//   						variableName: jsii.String("variableName"),
//   					},
//   					sns: &snsProperty{
//   						targetArn: jsii.String("targetArn"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					sqs: &sqsProperty{
//   						queueUrl: jsii.String("queueUrl"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						useBase64: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			condition: jsii.String("condition"),
//   		},
//   	},
//   }
//
type CfnDetectorModel_OnExitProperty struct {
	// Specifies the `actions` that are performed when the state is exited and the `condition` is `TRUE` .
	Events interface{} `field:"optional" json:"events" yaml:"events"`
}

// Specifies the actions performed when the `condition` evaluates to TRUE.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onInputProperty := &onInputProperty{
//   	events: []interface{}{
//   		&eventProperty{
//   			eventName: jsii.String("eventName"),
//
//   			// the properties below are optional
//   			actions: []interface{}{
//   				&actionProperty{
//   					clearTimer: &clearTimerProperty{
//   						timerName: jsii.String("timerName"),
//   					},
//   					dynamoDb: &dynamoDBProperty{
//   						hashKeyField: jsii.String("hashKeyField"),
//   						hashKeyValue: jsii.String("hashKeyValue"),
//   						tableName: jsii.String("tableName"),
//
//   						// the properties below are optional
//   						hashKeyType: jsii.String("hashKeyType"),
//   						operation: jsii.String("operation"),
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						payloadField: jsii.String("payloadField"),
//   						rangeKeyField: jsii.String("rangeKeyField"),
//   						rangeKeyType: jsii.String("rangeKeyType"),
//   						rangeKeyValue: jsii.String("rangeKeyValue"),
//   					},
//   					dynamoDBv2: &dynamoDBv2Property{
//   						tableName: jsii.String("tableName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					firehose: &firehoseProperty{
//   						deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						separator: jsii.String("separator"),
//   					},
//   					iotEvents: &iotEventsProperty{
//   						inputName: jsii.String("inputName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					iotSiteWise: &iotSiteWiseProperty{
//   						propertyValue: &assetPropertyValueProperty{
//   							value: &assetPropertyVariantProperty{
//   								booleanValue: jsii.String("booleanValue"),
//   								doubleValue: jsii.String("doubleValue"),
//   								integerValue: jsii.String("integerValue"),
//   								stringValue: jsii.String("stringValue"),
//   							},
//
//   							// the properties below are optional
//   							quality: jsii.String("quality"),
//   							timestamp: &assetPropertyTimestampProperty{
//   								timeInSeconds: jsii.String("timeInSeconds"),
//
//   								// the properties below are optional
//   								offsetInNanos: jsii.String("offsetInNanos"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						assetId: jsii.String("assetId"),
//   						entryId: jsii.String("entryId"),
//   						propertyAlias: jsii.String("propertyAlias"),
//   						propertyId: jsii.String("propertyId"),
//   					},
//   					iotTopicPublish: &iotTopicPublishProperty{
//   						mqttTopic: jsii.String("mqttTopic"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					lambda: &lambdaProperty{
//   						functionArn: jsii.String("functionArn"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					resetTimer: &resetTimerProperty{
//   						timerName: jsii.String("timerName"),
//   					},
//   					setTimer: &setTimerProperty{
//   						timerName: jsii.String("timerName"),
//
//   						// the properties below are optional
//   						durationExpression: jsii.String("durationExpression"),
//   						seconds: jsii.Number(123),
//   					},
//   					setVariable: &setVariableProperty{
//   						value: jsii.String("value"),
//   						variableName: jsii.String("variableName"),
//   					},
//   					sns: &snsProperty{
//   						targetArn: jsii.String("targetArn"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					sqs: &sqsProperty{
//   						queueUrl: jsii.String("queueUrl"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						useBase64: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			condition: jsii.String("condition"),
//   		},
//   	},
//   	transitionEvents: []interface{}{
//   		&transitionEventProperty{
//   			condition: jsii.String("condition"),
//   			eventName: jsii.String("eventName"),
//   			nextState: jsii.String("nextState"),
//
//   			// the properties below are optional
//   			actions: []interface{}{
//   				&actionProperty{
//   					clearTimer: &clearTimerProperty{
//   						timerName: jsii.String("timerName"),
//   					},
//   					dynamoDb: &dynamoDBProperty{
//   						hashKeyField: jsii.String("hashKeyField"),
//   						hashKeyValue: jsii.String("hashKeyValue"),
//   						tableName: jsii.String("tableName"),
//
//   						// the properties below are optional
//   						hashKeyType: jsii.String("hashKeyType"),
//   						operation: jsii.String("operation"),
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						payloadField: jsii.String("payloadField"),
//   						rangeKeyField: jsii.String("rangeKeyField"),
//   						rangeKeyType: jsii.String("rangeKeyType"),
//   						rangeKeyValue: jsii.String("rangeKeyValue"),
//   					},
//   					dynamoDBv2: &dynamoDBv2Property{
//   						tableName: jsii.String("tableName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					firehose: &firehoseProperty{
//   						deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						separator: jsii.String("separator"),
//   					},
//   					iotEvents: &iotEventsProperty{
//   						inputName: jsii.String("inputName"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					iotSiteWise: &iotSiteWiseProperty{
//   						propertyValue: &assetPropertyValueProperty{
//   							value: &assetPropertyVariantProperty{
//   								booleanValue: jsii.String("booleanValue"),
//   								doubleValue: jsii.String("doubleValue"),
//   								integerValue: jsii.String("integerValue"),
//   								stringValue: jsii.String("stringValue"),
//   							},
//
//   							// the properties below are optional
//   							quality: jsii.String("quality"),
//   							timestamp: &assetPropertyTimestampProperty{
//   								timeInSeconds: jsii.String("timeInSeconds"),
//
//   								// the properties below are optional
//   								offsetInNanos: jsii.String("offsetInNanos"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						assetId: jsii.String("assetId"),
//   						entryId: jsii.String("entryId"),
//   						propertyAlias: jsii.String("propertyAlias"),
//   						propertyId: jsii.String("propertyId"),
//   					},
//   					iotTopicPublish: &iotTopicPublishProperty{
//   						mqttTopic: jsii.String("mqttTopic"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					lambda: &lambdaProperty{
//   						functionArn: jsii.String("functionArn"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					resetTimer: &resetTimerProperty{
//   						timerName: jsii.String("timerName"),
//   					},
//   					setTimer: &setTimerProperty{
//   						timerName: jsii.String("timerName"),
//
//   						// the properties below are optional
//   						durationExpression: jsii.String("durationExpression"),
//   						seconds: jsii.Number(123),
//   					},
//   					setVariable: &setVariableProperty{
//   						value: jsii.String("value"),
//   						variableName: jsii.String("variableName"),
//   					},
//   					sns: &snsProperty{
//   						targetArn: jsii.String("targetArn"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   					},
//   					sqs: &sqsProperty{
//   						queueUrl: jsii.String("queueUrl"),
//
//   						// the properties below are optional
//   						payload: &payloadProperty{
//   							contentExpression: jsii.String("contentExpression"),
//   							type: jsii.String("type"),
//   						},
//   						useBase64: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDetectorModel_OnInputProperty struct {
	// Specifies the actions performed when the `condition` evaluates to TRUE.
	Events interface{} `field:"optional" json:"events" yaml:"events"`
	// Specifies the actions performed, and the next state entered, when a `condition` evaluates to TRUE.
	TransitionEvents interface{} `field:"optional" json:"transitionEvents" yaml:"transitionEvents"`
}

// Information needed to configure the payload.
//
// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use `contentExpression` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   payloadProperty := &payloadProperty{
//   	contentExpression: jsii.String("contentExpression"),
//   	type: jsii.String("type"),
//   }
//
type CfnDetectorModel_PayloadProperty struct {
	// The content of the payload.
	//
	// You can use a string expression that includes quoted strings ( `'<string>'` ), variables ( `$variable.<variable-name>` ), input values ( `$input.<input-name>.<path-to-datum>` ), string concatenations, and quoted strings that contain `${}` as the content. The recommended maximum size of a content expression is 1 KB.
	ContentExpression *string `field:"required" json:"contentExpression" yaml:"contentExpression"`
	// The value of the payload type can be either `STRING` or `JSON` .
	Type *string `field:"required" json:"type" yaml:"type"`
}

// Information required to reset the timer.
//
// The timer is reset to the previously evaluated result of the duration. The duration expression isn't reevaluated when you reset the timer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resetTimerProperty := &resetTimerProperty{
//   	timerName: jsii.String("timerName"),
//   }
//
type CfnDetectorModel_ResetTimerProperty struct {
	// The name of the timer to reset.
	TimerName *string `field:"required" json:"timerName" yaml:"timerName"`
}

// Information needed to set the timer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   setTimerProperty := &setTimerProperty{
//   	timerName: jsii.String("timerName"),
//
//   	// the properties below are optional
//   	durationExpression: jsii.String("durationExpression"),
//   	seconds: jsii.Number(123),
//   }
//
type CfnDetectorModel_SetTimerProperty struct {
	// The name of the timer.
	TimerName *string `field:"required" json:"timerName" yaml:"timerName"`
	// The duration of the timer, in seconds.
	//
	// You can use a string expression that includes numbers, variables ( `$variable.<variable-name>` ), and input values ( `$input.<input-name>.<path-to-datum>` ) as the duration. The range of the duration is 1-31622400 seconds. To ensure accuracy, the minimum duration is 60 seconds. The evaluated result of the duration is rounded down to the nearest whole number.
	DurationExpression *string `field:"optional" json:"durationExpression" yaml:"durationExpression"`
	// The number of seconds until the timer expires.
	//
	// The minimum value is 60 seconds to ensure accuracy. The maximum value is 31622400 seconds.
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
}

// Information about the variable and its new value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   setVariableProperty := &setVariableProperty{
//   	value: jsii.String("value"),
//   	variableName: jsii.String("variableName"),
//   }
//
type CfnDetectorModel_SetVariableProperty struct {
	// The new value of the variable.
	Value *string `field:"required" json:"value" yaml:"value"`
	// The name of the variable.
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
}

// Information required to publish the Amazon SNS message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snsProperty := &snsProperty{
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnDetectorModel_SnsProperty struct {
	// The ARN of the Amazon SNS target where the message is sent.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// You can configure the action payload when you send a message as an Amazon SNS push notification.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
}

// Sends information about the detector model instance and the event that triggered the action to an Amazon SQS queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqsProperty := &sqsProperty{
//   	queueUrl: jsii.String("queueUrl"),
//
//   	// the properties below are optional
//   	payload: &payloadProperty{
//   		contentExpression: jsii.String("contentExpression"),
//   		type: jsii.String("type"),
//   	},
//   	useBase64: jsii.Boolean(false),
//   }
//
type CfnDetectorModel_SqsProperty struct {
	// The URL of the SQS queue where the data is written.
	QueueUrl *string `field:"required" json:"queueUrl" yaml:"queueUrl"`
	// You can configure the action payload when you send a message to an Amazon SQS queue.
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
	// Set this to TRUE if you want the data to be base-64 encoded before it is written to the queue.
	//
	// Otherwise, set this to FALSE.
	UseBase64 interface{} `field:"optional" json:"useBase64" yaml:"useBase64"`
}

// Information that defines a state of a detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stateProperty := &stateProperty{
//   	stateName: jsii.String("stateName"),
//
//   	// the properties below are optional
//   	onEnter: &onEnterProperty{
//   		events: []interface{}{
//   			&eventProperty{
//   				eventName: jsii.String("eventName"),
//
//   				// the properties below are optional
//   				actions: []interface{}{
//   					&actionProperty{
//   						clearTimer: &clearTimerProperty{
//   							timerName: jsii.String("timerName"),
//   						},
//   						dynamoDb: &dynamoDBProperty{
//   							hashKeyField: jsii.String("hashKeyField"),
//   							hashKeyValue: jsii.String("hashKeyValue"),
//   							tableName: jsii.String("tableName"),
//
//   							// the properties below are optional
//   							hashKeyType: jsii.String("hashKeyType"),
//   							operation: jsii.String("operation"),
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   							payloadField: jsii.String("payloadField"),
//   							rangeKeyField: jsii.String("rangeKeyField"),
//   							rangeKeyType: jsii.String("rangeKeyType"),
//   							rangeKeyValue: jsii.String("rangeKeyValue"),
//   						},
//   						dynamoDBv2: &dynamoDBv2Property{
//   							tableName: jsii.String("tableName"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						firehose: &firehoseProperty{
//   							deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   							separator: jsii.String("separator"),
//   						},
//   						iotEvents: &iotEventsProperty{
//   							inputName: jsii.String("inputName"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						iotSiteWise: &iotSiteWiseProperty{
//   							propertyValue: &assetPropertyValueProperty{
//   								value: &assetPropertyVariantProperty{
//   									booleanValue: jsii.String("booleanValue"),
//   									doubleValue: jsii.String("doubleValue"),
//   									integerValue: jsii.String("integerValue"),
//   									stringValue: jsii.String("stringValue"),
//   								},
//
//   								// the properties below are optional
//   								quality: jsii.String("quality"),
//   								timestamp: &assetPropertyTimestampProperty{
//   									timeInSeconds: jsii.String("timeInSeconds"),
//
//   									// the properties below are optional
//   									offsetInNanos: jsii.String("offsetInNanos"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							assetId: jsii.String("assetId"),
//   							entryId: jsii.String("entryId"),
//   							propertyAlias: jsii.String("propertyAlias"),
//   							propertyId: jsii.String("propertyId"),
//   						},
//   						iotTopicPublish: &iotTopicPublishProperty{
//   							mqttTopic: jsii.String("mqttTopic"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						lambda: &lambdaProperty{
//   							functionArn: jsii.String("functionArn"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						resetTimer: &resetTimerProperty{
//   							timerName: jsii.String("timerName"),
//   						},
//   						setTimer: &setTimerProperty{
//   							timerName: jsii.String("timerName"),
//
//   							// the properties below are optional
//   							durationExpression: jsii.String("durationExpression"),
//   							seconds: jsii.Number(123),
//   						},
//   						setVariable: &setVariableProperty{
//   							value: jsii.String("value"),
//   							variableName: jsii.String("variableName"),
//   						},
//   						sns: &snsProperty{
//   							targetArn: jsii.String("targetArn"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						sqs: &sqsProperty{
//   							queueUrl: jsii.String("queueUrl"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   							useBase64: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				condition: jsii.String("condition"),
//   			},
//   		},
//   	},
//   	onExit: &onExitProperty{
//   		events: []interface{}{
//   			&eventProperty{
//   				eventName: jsii.String("eventName"),
//
//   				// the properties below are optional
//   				actions: []interface{}{
//   					&actionProperty{
//   						clearTimer: &clearTimerProperty{
//   							timerName: jsii.String("timerName"),
//   						},
//   						dynamoDb: &dynamoDBProperty{
//   							hashKeyField: jsii.String("hashKeyField"),
//   							hashKeyValue: jsii.String("hashKeyValue"),
//   							tableName: jsii.String("tableName"),
//
//   							// the properties below are optional
//   							hashKeyType: jsii.String("hashKeyType"),
//   							operation: jsii.String("operation"),
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   							payloadField: jsii.String("payloadField"),
//   							rangeKeyField: jsii.String("rangeKeyField"),
//   							rangeKeyType: jsii.String("rangeKeyType"),
//   							rangeKeyValue: jsii.String("rangeKeyValue"),
//   						},
//   						dynamoDBv2: &dynamoDBv2Property{
//   							tableName: jsii.String("tableName"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						firehose: &firehoseProperty{
//   							deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   							separator: jsii.String("separator"),
//   						},
//   						iotEvents: &iotEventsProperty{
//   							inputName: jsii.String("inputName"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						iotSiteWise: &iotSiteWiseProperty{
//   							propertyValue: &assetPropertyValueProperty{
//   								value: &assetPropertyVariantProperty{
//   									booleanValue: jsii.String("booleanValue"),
//   									doubleValue: jsii.String("doubleValue"),
//   									integerValue: jsii.String("integerValue"),
//   									stringValue: jsii.String("stringValue"),
//   								},
//
//   								// the properties below are optional
//   								quality: jsii.String("quality"),
//   								timestamp: &assetPropertyTimestampProperty{
//   									timeInSeconds: jsii.String("timeInSeconds"),
//
//   									// the properties below are optional
//   									offsetInNanos: jsii.String("offsetInNanos"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							assetId: jsii.String("assetId"),
//   							entryId: jsii.String("entryId"),
//   							propertyAlias: jsii.String("propertyAlias"),
//   							propertyId: jsii.String("propertyId"),
//   						},
//   						iotTopicPublish: &iotTopicPublishProperty{
//   							mqttTopic: jsii.String("mqttTopic"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						lambda: &lambdaProperty{
//   							functionArn: jsii.String("functionArn"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						resetTimer: &resetTimerProperty{
//   							timerName: jsii.String("timerName"),
//   						},
//   						setTimer: &setTimerProperty{
//   							timerName: jsii.String("timerName"),
//
//   							// the properties below are optional
//   							durationExpression: jsii.String("durationExpression"),
//   							seconds: jsii.Number(123),
//   						},
//   						setVariable: &setVariableProperty{
//   							value: jsii.String("value"),
//   							variableName: jsii.String("variableName"),
//   						},
//   						sns: &snsProperty{
//   							targetArn: jsii.String("targetArn"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						sqs: &sqsProperty{
//   							queueUrl: jsii.String("queueUrl"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   							useBase64: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				condition: jsii.String("condition"),
//   			},
//   		},
//   	},
//   	onInput: &onInputProperty{
//   		events: []interface{}{
//   			&eventProperty{
//   				eventName: jsii.String("eventName"),
//
//   				// the properties below are optional
//   				actions: []interface{}{
//   					&actionProperty{
//   						clearTimer: &clearTimerProperty{
//   							timerName: jsii.String("timerName"),
//   						},
//   						dynamoDb: &dynamoDBProperty{
//   							hashKeyField: jsii.String("hashKeyField"),
//   							hashKeyValue: jsii.String("hashKeyValue"),
//   							tableName: jsii.String("tableName"),
//
//   							// the properties below are optional
//   							hashKeyType: jsii.String("hashKeyType"),
//   							operation: jsii.String("operation"),
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   							payloadField: jsii.String("payloadField"),
//   							rangeKeyField: jsii.String("rangeKeyField"),
//   							rangeKeyType: jsii.String("rangeKeyType"),
//   							rangeKeyValue: jsii.String("rangeKeyValue"),
//   						},
//   						dynamoDBv2: &dynamoDBv2Property{
//   							tableName: jsii.String("tableName"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						firehose: &firehoseProperty{
//   							deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   							separator: jsii.String("separator"),
//   						},
//   						iotEvents: &iotEventsProperty{
//   							inputName: jsii.String("inputName"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						iotSiteWise: &iotSiteWiseProperty{
//   							propertyValue: &assetPropertyValueProperty{
//   								value: &assetPropertyVariantProperty{
//   									booleanValue: jsii.String("booleanValue"),
//   									doubleValue: jsii.String("doubleValue"),
//   									integerValue: jsii.String("integerValue"),
//   									stringValue: jsii.String("stringValue"),
//   								},
//
//   								// the properties below are optional
//   								quality: jsii.String("quality"),
//   								timestamp: &assetPropertyTimestampProperty{
//   									timeInSeconds: jsii.String("timeInSeconds"),
//
//   									// the properties below are optional
//   									offsetInNanos: jsii.String("offsetInNanos"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							assetId: jsii.String("assetId"),
//   							entryId: jsii.String("entryId"),
//   							propertyAlias: jsii.String("propertyAlias"),
//   							propertyId: jsii.String("propertyId"),
//   						},
//   						iotTopicPublish: &iotTopicPublishProperty{
//   							mqttTopic: jsii.String("mqttTopic"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						lambda: &lambdaProperty{
//   							functionArn: jsii.String("functionArn"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						resetTimer: &resetTimerProperty{
//   							timerName: jsii.String("timerName"),
//   						},
//   						setTimer: &setTimerProperty{
//   							timerName: jsii.String("timerName"),
//
//   							// the properties below are optional
//   							durationExpression: jsii.String("durationExpression"),
//   							seconds: jsii.Number(123),
//   						},
//   						setVariable: &setVariableProperty{
//   							value: jsii.String("value"),
//   							variableName: jsii.String("variableName"),
//   						},
//   						sns: &snsProperty{
//   							targetArn: jsii.String("targetArn"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						sqs: &sqsProperty{
//   							queueUrl: jsii.String("queueUrl"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   							useBase64: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				condition: jsii.String("condition"),
//   			},
//   		},
//   		transitionEvents: []interface{}{
//   			&transitionEventProperty{
//   				condition: jsii.String("condition"),
//   				eventName: jsii.String("eventName"),
//   				nextState: jsii.String("nextState"),
//
//   				// the properties below are optional
//   				actions: []interface{}{
//   					&actionProperty{
//   						clearTimer: &clearTimerProperty{
//   							timerName: jsii.String("timerName"),
//   						},
//   						dynamoDb: &dynamoDBProperty{
//   							hashKeyField: jsii.String("hashKeyField"),
//   							hashKeyValue: jsii.String("hashKeyValue"),
//   							tableName: jsii.String("tableName"),
//
//   							// the properties below are optional
//   							hashKeyType: jsii.String("hashKeyType"),
//   							operation: jsii.String("operation"),
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   							payloadField: jsii.String("payloadField"),
//   							rangeKeyField: jsii.String("rangeKeyField"),
//   							rangeKeyType: jsii.String("rangeKeyType"),
//   							rangeKeyValue: jsii.String("rangeKeyValue"),
//   						},
//   						dynamoDBv2: &dynamoDBv2Property{
//   							tableName: jsii.String("tableName"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						firehose: &firehoseProperty{
//   							deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   							separator: jsii.String("separator"),
//   						},
//   						iotEvents: &iotEventsProperty{
//   							inputName: jsii.String("inputName"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						iotSiteWise: &iotSiteWiseProperty{
//   							propertyValue: &assetPropertyValueProperty{
//   								value: &assetPropertyVariantProperty{
//   									booleanValue: jsii.String("booleanValue"),
//   									doubleValue: jsii.String("doubleValue"),
//   									integerValue: jsii.String("integerValue"),
//   									stringValue: jsii.String("stringValue"),
//   								},
//
//   								// the properties below are optional
//   								quality: jsii.String("quality"),
//   								timestamp: &assetPropertyTimestampProperty{
//   									timeInSeconds: jsii.String("timeInSeconds"),
//
//   									// the properties below are optional
//   									offsetInNanos: jsii.String("offsetInNanos"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							assetId: jsii.String("assetId"),
//   							entryId: jsii.String("entryId"),
//   							propertyAlias: jsii.String("propertyAlias"),
//   							propertyId: jsii.String("propertyId"),
//   						},
//   						iotTopicPublish: &iotTopicPublishProperty{
//   							mqttTopic: jsii.String("mqttTopic"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						lambda: &lambdaProperty{
//   							functionArn: jsii.String("functionArn"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						resetTimer: &resetTimerProperty{
//   							timerName: jsii.String("timerName"),
//   						},
//   						setTimer: &setTimerProperty{
//   							timerName: jsii.String("timerName"),
//
//   							// the properties below are optional
//   							durationExpression: jsii.String("durationExpression"),
//   							seconds: jsii.Number(123),
//   						},
//   						setVariable: &setVariableProperty{
//   							value: jsii.String("value"),
//   							variableName: jsii.String("variableName"),
//   						},
//   						sns: &snsProperty{
//   							targetArn: jsii.String("targetArn"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   						},
//   						sqs: &sqsProperty{
//   							queueUrl: jsii.String("queueUrl"),
//
//   							// the properties below are optional
//   							payload: &payloadProperty{
//   								contentExpression: jsii.String("contentExpression"),
//   								type: jsii.String("type"),
//   							},
//   							useBase64: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDetectorModel_StateProperty struct {
	// The name of the state.
	StateName *string `field:"required" json:"stateName" yaml:"stateName"`
	// When entering this state, perform these `actions` if the `condition` is TRUE.
	OnEnter interface{} `field:"optional" json:"onEnter" yaml:"onEnter"`
	// When exiting this state, perform these `actions` if the specified `condition` is `TRUE` .
	OnExit interface{} `field:"optional" json:"onExit" yaml:"onExit"`
	// When an input is received and the `condition` is TRUE, perform the specified `actions` .
	OnInput interface{} `field:"optional" json:"onInput" yaml:"onInput"`
}

// Specifies the actions performed and the next state entered when a `condition` evaluates to TRUE.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitionEventProperty := &transitionEventProperty{
//   	condition: jsii.String("condition"),
//   	eventName: jsii.String("eventName"),
//   	nextState: jsii.String("nextState"),
//
//   	// the properties below are optional
//   	actions: []interface{}{
//   		&actionProperty{
//   			clearTimer: &clearTimerProperty{
//   				timerName: jsii.String("timerName"),
//   			},
//   			dynamoDb: &dynamoDBProperty{
//   				hashKeyField: jsii.String("hashKeyField"),
//   				hashKeyValue: jsii.String("hashKeyValue"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				hashKeyType: jsii.String("hashKeyType"),
//   				operation: jsii.String("operation"),
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   				payloadField: jsii.String("payloadField"),
//   				rangeKeyField: jsii.String("rangeKeyField"),
//   				rangeKeyType: jsii.String("rangeKeyType"),
//   				rangeKeyValue: jsii.String("rangeKeyValue"),
//   			},
//   			dynamoDBv2: &dynamoDBv2Property{
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			firehose: &firehoseProperty{
//   				deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   				separator: jsii.String("separator"),
//   			},
//   			iotEvents: &iotEventsProperty{
//   				inputName: jsii.String("inputName"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			iotSiteWise: &iotSiteWiseProperty{
//   				propertyValue: &assetPropertyValueProperty{
//   					value: &assetPropertyVariantProperty{
//   						booleanValue: jsii.String("booleanValue"),
//   						doubleValue: jsii.String("doubleValue"),
//   						integerValue: jsii.String("integerValue"),
//   						stringValue: jsii.String("stringValue"),
//   					},
//
//   					// the properties below are optional
//   					quality: jsii.String("quality"),
//   					timestamp: &assetPropertyTimestampProperty{
//   						timeInSeconds: jsii.String("timeInSeconds"),
//
//   						// the properties below are optional
//   						offsetInNanos: jsii.String("offsetInNanos"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				assetId: jsii.String("assetId"),
//   				entryId: jsii.String("entryId"),
//   				propertyAlias: jsii.String("propertyAlias"),
//   				propertyId: jsii.String("propertyId"),
//   			},
//   			iotTopicPublish: &iotTopicPublishProperty{
//   				mqttTopic: jsii.String("mqttTopic"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			lambda: &lambdaProperty{
//   				functionArn: jsii.String("functionArn"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			resetTimer: &resetTimerProperty{
//   				timerName: jsii.String("timerName"),
//   			},
//   			setTimer: &setTimerProperty{
//   				timerName: jsii.String("timerName"),
//
//   				// the properties below are optional
//   				durationExpression: jsii.String("durationExpression"),
//   				seconds: jsii.Number(123),
//   			},
//   			setVariable: &setVariableProperty{
//   				value: jsii.String("value"),
//   				variableName: jsii.String("variableName"),
//   			},
//   			sns: &snsProperty{
//   				targetArn: jsii.String("targetArn"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			sqs: &sqsProperty{
//   				queueUrl: jsii.String("queueUrl"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   				useBase64: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
//
type CfnDetectorModel_TransitionEventProperty struct {
	// Required.
	//
	// A Boolean expression that when TRUE causes the actions to be performed and the `nextState` to be entered.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// The name of the transition event.
	EventName *string `field:"required" json:"eventName" yaml:"eventName"`
	// The next state to enter.
	NextState *string `field:"required" json:"nextState" yaml:"nextState"`
	// The actions to be performed.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
}

// Properties for defining a `CfnDetectorModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDetectorModelProps := &cfnDetectorModelProps{
//   	detectorModelDefinition: &detectorModelDefinitionProperty{
//   		initialStateName: jsii.String("initialStateName"),
//   		states: []interface{}{
//   			&stateProperty{
//   				stateName: jsii.String("stateName"),
//
//   				// the properties below are optional
//   				onEnter: &onEnterProperty{
//   					events: []interface{}{
//   						&eventProperty{
//   							eventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							actions: []interface{}{
//   								&actionProperty{
//   									clearTimer: &clearTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									dynamoDb: &dynamoDBProperty{
//   										hashKeyField: jsii.String("hashKeyField"),
//   										hashKeyValue: jsii.String("hashKeyValue"),
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										hashKeyType: jsii.String("hashKeyType"),
//   										operation: jsii.String("operation"),
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										payloadField: jsii.String("payloadField"),
//   										rangeKeyField: jsii.String("rangeKeyField"),
//   										rangeKeyType: jsii.String("rangeKeyType"),
//   										rangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									dynamoDBv2: &dynamoDBv2Property{
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									firehose: &firehoseProperty{
//   										deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										separator: jsii.String("separator"),
//   									},
//   									iotEvents: &iotEventsProperty{
//   										inputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									iotSiteWise: &iotSiteWiseProperty{
//   										propertyValue: &assetPropertyValueProperty{
//   											value: &assetPropertyVariantProperty{
//   												booleanValue: jsii.String("booleanValue"),
//   												doubleValue: jsii.String("doubleValue"),
//   												integerValue: jsii.String("integerValue"),
//   												stringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											quality: jsii.String("quality"),
//   											timestamp: &assetPropertyTimestampProperty{
//   												timeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												offsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										assetId: jsii.String("assetId"),
//   										entryId: jsii.String("entryId"),
//   										propertyAlias: jsii.String("propertyAlias"),
//   										propertyId: jsii.String("propertyId"),
//   									},
//   									iotTopicPublish: &iotTopicPublishProperty{
//   										mqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									lambda: &lambdaProperty{
//   										functionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									resetTimer: &resetTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									setTimer: &setTimerProperty{
//   										timerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										durationExpression: jsii.String("durationExpression"),
//   										seconds: jsii.Number(123),
//   									},
//   									setVariable: &setVariableProperty{
//   										value: jsii.String("value"),
//   										variableName: jsii.String("variableName"),
//   									},
//   									sns: &snsProperty{
//   										targetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									sqs: &sqsProperty{
//   										queueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										useBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							condition: jsii.String("condition"),
//   						},
//   					},
//   				},
//   				onExit: &onExitProperty{
//   					events: []interface{}{
//   						&eventProperty{
//   							eventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							actions: []interface{}{
//   								&actionProperty{
//   									clearTimer: &clearTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									dynamoDb: &dynamoDBProperty{
//   										hashKeyField: jsii.String("hashKeyField"),
//   										hashKeyValue: jsii.String("hashKeyValue"),
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										hashKeyType: jsii.String("hashKeyType"),
//   										operation: jsii.String("operation"),
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										payloadField: jsii.String("payloadField"),
//   										rangeKeyField: jsii.String("rangeKeyField"),
//   										rangeKeyType: jsii.String("rangeKeyType"),
//   										rangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									dynamoDBv2: &dynamoDBv2Property{
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									firehose: &firehoseProperty{
//   										deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										separator: jsii.String("separator"),
//   									},
//   									iotEvents: &iotEventsProperty{
//   										inputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									iotSiteWise: &iotSiteWiseProperty{
//   										propertyValue: &assetPropertyValueProperty{
//   											value: &assetPropertyVariantProperty{
//   												booleanValue: jsii.String("booleanValue"),
//   												doubleValue: jsii.String("doubleValue"),
//   												integerValue: jsii.String("integerValue"),
//   												stringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											quality: jsii.String("quality"),
//   											timestamp: &assetPropertyTimestampProperty{
//   												timeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												offsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										assetId: jsii.String("assetId"),
//   										entryId: jsii.String("entryId"),
//   										propertyAlias: jsii.String("propertyAlias"),
//   										propertyId: jsii.String("propertyId"),
//   									},
//   									iotTopicPublish: &iotTopicPublishProperty{
//   										mqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									lambda: &lambdaProperty{
//   										functionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									resetTimer: &resetTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									setTimer: &setTimerProperty{
//   										timerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										durationExpression: jsii.String("durationExpression"),
//   										seconds: jsii.Number(123),
//   									},
//   									setVariable: &setVariableProperty{
//   										value: jsii.String("value"),
//   										variableName: jsii.String("variableName"),
//   									},
//   									sns: &snsProperty{
//   										targetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									sqs: &sqsProperty{
//   										queueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										useBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							condition: jsii.String("condition"),
//   						},
//   					},
//   				},
//   				onInput: &onInputProperty{
//   					events: []interface{}{
//   						&eventProperty{
//   							eventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							actions: []interface{}{
//   								&actionProperty{
//   									clearTimer: &clearTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									dynamoDb: &dynamoDBProperty{
//   										hashKeyField: jsii.String("hashKeyField"),
//   										hashKeyValue: jsii.String("hashKeyValue"),
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										hashKeyType: jsii.String("hashKeyType"),
//   										operation: jsii.String("operation"),
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										payloadField: jsii.String("payloadField"),
//   										rangeKeyField: jsii.String("rangeKeyField"),
//   										rangeKeyType: jsii.String("rangeKeyType"),
//   										rangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									dynamoDBv2: &dynamoDBv2Property{
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									firehose: &firehoseProperty{
//   										deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										separator: jsii.String("separator"),
//   									},
//   									iotEvents: &iotEventsProperty{
//   										inputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									iotSiteWise: &iotSiteWiseProperty{
//   										propertyValue: &assetPropertyValueProperty{
//   											value: &assetPropertyVariantProperty{
//   												booleanValue: jsii.String("booleanValue"),
//   												doubleValue: jsii.String("doubleValue"),
//   												integerValue: jsii.String("integerValue"),
//   												stringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											quality: jsii.String("quality"),
//   											timestamp: &assetPropertyTimestampProperty{
//   												timeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												offsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										assetId: jsii.String("assetId"),
//   										entryId: jsii.String("entryId"),
//   										propertyAlias: jsii.String("propertyAlias"),
//   										propertyId: jsii.String("propertyId"),
//   									},
//   									iotTopicPublish: &iotTopicPublishProperty{
//   										mqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									lambda: &lambdaProperty{
//   										functionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									resetTimer: &resetTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									setTimer: &setTimerProperty{
//   										timerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										durationExpression: jsii.String("durationExpression"),
//   										seconds: jsii.Number(123),
//   									},
//   									setVariable: &setVariableProperty{
//   										value: jsii.String("value"),
//   										variableName: jsii.String("variableName"),
//   									},
//   									sns: &snsProperty{
//   										targetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									sqs: &sqsProperty{
//   										queueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										useBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							condition: jsii.String("condition"),
//   						},
//   					},
//   					transitionEvents: []interface{}{
//   						&transitionEventProperty{
//   							condition: jsii.String("condition"),
//   							eventName: jsii.String("eventName"),
//   							nextState: jsii.String("nextState"),
//
//   							// the properties below are optional
//   							actions: []interface{}{
//   								&actionProperty{
//   									clearTimer: &clearTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									dynamoDb: &dynamoDBProperty{
//   										hashKeyField: jsii.String("hashKeyField"),
//   										hashKeyValue: jsii.String("hashKeyValue"),
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										hashKeyType: jsii.String("hashKeyType"),
//   										operation: jsii.String("operation"),
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										payloadField: jsii.String("payloadField"),
//   										rangeKeyField: jsii.String("rangeKeyField"),
//   										rangeKeyType: jsii.String("rangeKeyType"),
//   										rangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									dynamoDBv2: &dynamoDBv2Property{
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									firehose: &firehoseProperty{
//   										deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										separator: jsii.String("separator"),
//   									},
//   									iotEvents: &iotEventsProperty{
//   										inputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									iotSiteWise: &iotSiteWiseProperty{
//   										propertyValue: &assetPropertyValueProperty{
//   											value: &assetPropertyVariantProperty{
//   												booleanValue: jsii.String("booleanValue"),
//   												doubleValue: jsii.String("doubleValue"),
//   												integerValue: jsii.String("integerValue"),
//   												stringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											quality: jsii.String("quality"),
//   											timestamp: &assetPropertyTimestampProperty{
//   												timeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												offsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										assetId: jsii.String("assetId"),
//   										entryId: jsii.String("entryId"),
//   										propertyAlias: jsii.String("propertyAlias"),
//   										propertyId: jsii.String("propertyId"),
//   									},
//   									iotTopicPublish: &iotTopicPublishProperty{
//   										mqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									lambda: &lambdaProperty{
//   										functionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									resetTimer: &resetTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									setTimer: &setTimerProperty{
//   										timerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										durationExpression: jsii.String("durationExpression"),
//   										seconds: jsii.Number(123),
//   									},
//   									setVariable: &setVariableProperty{
//   										value: jsii.String("value"),
//   										variableName: jsii.String("variableName"),
//   									},
//   									sns: &snsProperty{
//   										targetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									sqs: &sqsProperty{
//   										queueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										useBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	detectorModelDescription: jsii.String("detectorModelDescription"),
//   	detectorModelName: jsii.String("detectorModelName"),
//   	evaluationMethod: jsii.String("evaluationMethod"),
//   	key: jsii.String("key"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDetectorModelProps struct {
	// Information that defines how a detector operates.
	DetectorModelDefinition interface{} `field:"required" json:"detectorModelDefinition" yaml:"detectorModelDefinition"`
	// The ARN of the role that grants permission to AWS IoT Events to perform its operations.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A brief description of the detector model.
	DetectorModelDescription *string `field:"optional" json:"detectorModelDescription" yaml:"detectorModelDescription"`
	// The name of the detector model.
	DetectorModelName *string `field:"optional" json:"detectorModelName" yaml:"detectorModelName"`
	// Information about the order in which events are evaluated and how actions are executed.
	EvaluationMethod *string `field:"optional" json:"evaluationMethod" yaml:"evaluationMethod"`
	// The value used to identify a detector instance.
	//
	// When a device or system sends input, a new detector instance with a unique key value is created. AWS IoT Events can continue to route input to its corresponding detector instance based on this identifying information.
	//
	// This parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct detector instance, the device must send a message payload that contains the same attribute-value.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::IoTEvents::Input`.
//
// The AWS::IoTEvents::Input resource creates an input. To monitor your devices and processes, they must have a way to get telemetry data into AWS IoT Events . This is done by sending messages as *inputs* to AWS IoT Events . For more information, see [How to Use AWS IoT Events](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInput := awscdk.Aws_iotevents.NewCfnInput(this, jsii.String("MyCfnInput"), &cfnInputProps{
//   	inputDefinition: &inputDefinitionProperty{
//   		attributes: []interface{}{
//   			&attributeProperty{
//   				jsonPath: jsii.String("jsonPath"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	inputDescription: jsii.String("inputDescription"),
//   	inputName: jsii.String("inputName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnInput interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The definition of the input.
	InputDefinition() interface{}
	SetInputDefinition(val interface{})
	// A brief description of the input.
	InputDescription() *string
	SetInputDescription(val *string)
	// The name of the input.
	InputName() *string
	SetInputName(val *string)
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
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
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

// The jsii proxy struct for CfnInput
type jsiiProxy_CfnInput struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInput) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) InputDefinition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) InputDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) InputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTEvents::Input`.
func NewCfnInput(scope constructs.Construct, id *string, props *CfnInputProps) CfnInput {
	_init_.Initialize()

	j := jsiiProxy_CfnInput{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotevents.CfnInput",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTEvents::Input`.
func NewCfnInput_Override(c CfnInput, scope constructs.Construct, id *string, props *CfnInputProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotevents.CfnInput",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInput) SetInputDefinition(val interface{}) {
	_jsii_.Set(
		j,
		"inputDefinition",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetInputDescription(val *string) {
	_jsii_.Set(
		j,
		"inputDescription",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetInputName(val *string) {
	_jsii_.Set(
		j,
		"inputName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnInput_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotevents.CfnInput",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnInput_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotevents.CfnInput",
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
func CfnInput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotevents.CfnInput",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInput_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotevents.CfnInput",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInput) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnInput) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInput) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnInput) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnInput) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnInput) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnInput) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnInput) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInput) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInput) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnInput) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInput) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInput) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInput) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The attributes from the JSON payload that are made available by the input.
//
// Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage` . Each such message contains a JSON payload. Those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeProperty := &attributeProperty{
//   	jsonPath: jsii.String("jsonPath"),
//   }
//
type CfnInput_AttributeProperty struct {
	// An expression that specifies an attribute-value pair in a JSON structure.
	//
	// Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events ( `BatchPutMessage` ). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the `condition` expressions used by detectors.
	//
	// Syntax: `<field-name>.<field-name>...`
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
}

// The definition of the input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputDefinitionProperty := &inputDefinitionProperty{
//   	attributes: []interface{}{
//   		&attributeProperty{
//   			jsonPath: jsii.String("jsonPath"),
//   		},
//   	},
//   }
//
type CfnInput_InputDefinitionProperty struct {
	// The attributes from the JSON payload that are made available by the input.
	//
	// Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage` . Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
}

// Properties for defining a `CfnInput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInputProps := &cfnInputProps{
//   	inputDefinition: &inputDefinitionProperty{
//   		attributes: []interface{}{
//   			&attributeProperty{
//   				jsonPath: jsii.String("jsonPath"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	inputDescription: jsii.String("inputDescription"),
//   	inputName: jsii.String("inputName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnInputProps struct {
	// The definition of the input.
	InputDefinition interface{} `field:"required" json:"inputDefinition" yaml:"inputDefinition"`
	// A brief description of the input.
	InputDescription *string `field:"optional" json:"inputDescription" yaml:"inputDescription"`
	// The name of the input.
	InputName *string `field:"optional" json:"inputName" yaml:"inputName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

