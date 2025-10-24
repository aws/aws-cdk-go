package awsiotevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an alarm model to monitor an AWS IoT Events input attribute.
//
// You can use the alarm to get notified when the value is outside a specified range. For more information, see [Create an alarm model](https://docs.aws.amazon.com/iotevents/latest/developerguide/create-alarms.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlarmModel := awscdk.Aws_iotevents.NewCfnAlarmModel(this, jsii.String("MyCfnAlarmModel"), &CfnAlarmModelProps{
//   	AlarmRule: &AlarmRuleProperty{
//   		SimpleRule: &SimpleRuleProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			InputProperty: jsii.String("inputProperty"),
//   			Threshold: jsii.String("threshold"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AlarmCapabilities: &AlarmCapabilitiesProperty{
//   		AcknowledgeFlow: &AcknowledgeFlowProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		InitializationConfiguration: &InitializationConfigurationProperty{
//   			DisabledOnInitialization: jsii.Boolean(false),
//   		},
//   	},
//   	AlarmEventActions: &AlarmEventActionsProperty{
//   		AlarmActions: []interface{}{
//   			&AlarmActionProperty{
//   				DynamoDb: &DynamoDBProperty{
//   					HashKeyField: jsii.String("hashKeyField"),
//   					HashKeyValue: jsii.String("hashKeyValue"),
//   					TableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					HashKeyType: jsii.String("hashKeyType"),
//   					Operation: jsii.String("operation"),
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					PayloadField: jsii.String("payloadField"),
//   					RangeKeyField: jsii.String("rangeKeyField"),
//   					RangeKeyType: jsii.String("rangeKeyType"),
//   					RangeKeyValue: jsii.String("rangeKeyValue"),
//   				},
//   				DynamoDBv2: &DynamoDBv2Property{
//   					TableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Firehose: &FirehoseProperty{
//   					DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					Separator: jsii.String("separator"),
//   				},
//   				IotEvents: &IotEventsProperty{
//   					InputName: jsii.String("inputName"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				IotSiteWise: &IotSiteWiseProperty{
//   					AssetId: jsii.String("assetId"),
//   					EntryId: jsii.String("entryId"),
//   					PropertyAlias: jsii.String("propertyAlias"),
//   					PropertyId: jsii.String("propertyId"),
//   					PropertyValue: &AssetPropertyValueProperty{
//   						Value: &AssetPropertyVariantProperty{
//   							BooleanValue: jsii.String("booleanValue"),
//   							DoubleValue: jsii.String("doubleValue"),
//   							IntegerValue: jsii.String("integerValue"),
//   							StringValue: jsii.String("stringValue"),
//   						},
//
//   						// the properties below are optional
//   						Quality: jsii.String("quality"),
//   						Timestamp: &AssetPropertyTimestampProperty{
//   							TimeInSeconds: jsii.String("timeInSeconds"),
//
//   							// the properties below are optional
//   							OffsetInNanos: jsii.String("offsetInNanos"),
//   						},
//   					},
//   				},
//   				IotTopicPublish: &IotTopicPublishProperty{
//   					MqttTopic: jsii.String("mqttTopic"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Lambda: &LambdaProperty{
//   					FunctionArn: jsii.String("functionArn"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Sns: &SnsProperty{
//   					TargetArn: jsii.String("targetArn"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Sqs: &SqsProperty{
//   					QueueUrl: jsii.String("queueUrl"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					UseBase64: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	AlarmModelDescription: jsii.String("alarmModelDescription"),
//   	AlarmModelName: jsii.String("alarmModelName"),
//   	Key: jsii.String("key"),
//   	Severity: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html
//
type CfnAlarmModel interface {
	awscdk.CfnResource
	IAlarmModelRef
	awscdk.IInspectable
	awscdk.ITaggable
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
	// A reference to a AlarmModel resource.
	AlarmModelRef() *AlarmModelReference
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
	RoleArn() *string
	SetRoleArn(val *string)
	// A non-negative integer that reflects the severity level of the alarm.
	Severity() *float64
	SetSeverity(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// A list of key-value pairs that contain metadata for the alarm model.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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
	jsiiProxy_IAlarmModelRef
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
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

func (j *jsiiProxy_CfnAlarmModel) AlarmModelRef() *AlarmModelReference {
	var returns *AlarmModelReference
	_jsii_.Get(
		j,
		"alarmModelRef",
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

func (j *jsiiProxy_CfnAlarmModel) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
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

func (j *jsiiProxy_CfnAlarmModel) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnAlarmModel(scope constructs.Construct, id *string, props *CfnAlarmModelProps) CfnAlarmModel {
	_init_.Initialize()

	if err := validateNewCfnAlarmModelParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAlarmModel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotevents.CfnAlarmModel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnAlarmModel_Override(c CfnAlarmModel, scope constructs.Construct, id *string, props *CfnAlarmModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotevents.CfnAlarmModel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAlarmModel)SetAlarmCapabilities(val interface{}) {
	if err := j.validateSetAlarmCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmCapabilities",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel)SetAlarmEventActions(val interface{}) {
	if err := j.validateSetAlarmEventActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmEventActions",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel)SetAlarmModelDescription(val *string) {
	_jsii_.Set(
		j,
		"alarmModelDescription",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel)SetAlarmModelName(val *string) {
	_jsii_.Set(
		j,
		"alarmModelName",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel)SetAlarmRule(val interface{}) {
	if err := j.validateSetAlarmRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmRule",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel)SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel)SetSeverity(val *float64) {
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_CfnAlarmModel)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

// Creates a new IAlarmModelRef from a alarmModelName.
func CfnAlarmModel_FromAlarmModelName(scope constructs.Construct, id *string, alarmModelName *string) IAlarmModelRef {
	_init_.Initialize()

	if err := validateCfnAlarmModel_FromAlarmModelNameParameters(scope, id, alarmModelName); err != nil {
		panic(err)
	}
	var returns IAlarmModelRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotevents.CfnAlarmModel",
		"fromAlarmModelName",
		[]interface{}{scope, id, alarmModelName},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAlarmModel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAlarmModel_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotevents.CfnAlarmModel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnAlarmModel_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAlarmModel_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotevents.CfnAlarmModel",
		"isCfnResource",
		[]interface{}{x},
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

	if err := validateCfnAlarmModel_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAlarmModel) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAlarmModel) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAlarmModel) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAlarmModel) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAlarmModel) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAlarmModel) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAlarmModel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAlarmModel) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarmModel) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
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
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAlarmModel) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarmModel) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarmModel) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAlarmModel) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAlarmModel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarmModel) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

