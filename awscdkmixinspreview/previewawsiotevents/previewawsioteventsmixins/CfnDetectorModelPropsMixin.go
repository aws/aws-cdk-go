package previewawsioteventsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiotevents/previewawsioteventsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::IoTEvents::DetectorModel resource creates a detector model.
//
// You create a *detector model* (a model of your equipment or process) using *states* . For each state, you define conditional (Boolean) logic that evaluates the incoming inputs to detect significant events. When an event is detected, it can change the state or trigger custom-built or predefined actions using other AWS services. You can define additional events that trigger actions when entering or exiting a state and, optionally, when a condition is met. For more information, see [How to Use AWS IoT Events](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *AWS IoT Events Developer Guide* .
//
// > When you successfully update a detector model (using the AWS IoT Events console, AWS IoT Events API or CLI commands, or CloudFormation ) all detector instances created by the model are reset to their initial states. (The detector's `state` , and the values of any variables and timers are reset.)
// >
// > When you successfully update a detector model (using the AWS IoT Events console, AWS IoT Events API or CLI commands, or CloudFormation ) the version number of the detector model is incremented. (A detector model with version number 1 before the update has version number 2 after the update succeeds.)
// >
// > If you attempt to update a detector model using CloudFormation and the update does not succeed, the system may, in some cases, restore the original detector model. When this occurs, the detector model's version is incremented twice (for example, from version 1 to version 3) and the detector instances are reset.
// >
// > Also, be aware that if you attempt to update several detector models at once using CloudFormation , some updates may succeed and others fail. In this case, the effects on each detector model's detector instances and version number depend on whether the update succeeded or failed, with the results as stated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDetectorModelPropsMixin := awscdkmixinspreview.Mixins.NewCfnDetectorModelPropsMixin(&CfnDetectorModelMixinProps{
//   	DetectorModelDefinition: &DetectorModelDefinitionProperty{
//   		InitialStateName: jsii.String("initialStateName"),
//   		States: []interface{}{
//   			&StateProperty{
//   				OnEnter: &OnEnterProperty{
//   					Events: []interface{}{
//   						&EventProperty{
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyType: jsii.String("hashKeyType"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   										TableName: jsii.String("tableName"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TableName: jsii.String("tableName"),
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//   											},
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//   										},
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TargetArn: jsii.String("targetArn"),
//   									},
//   									Sqs: &SqsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										QueueUrl: jsii.String("queueUrl"),
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   							EventName: jsii.String("eventName"),
//   						},
//   					},
//   				},
//   				OnExit: &OnExitProperty{
//   					Events: []interface{}{
//   						&EventProperty{
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyType: jsii.String("hashKeyType"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   										TableName: jsii.String("tableName"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TableName: jsii.String("tableName"),
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//   											},
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//   										},
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TargetArn: jsii.String("targetArn"),
//   									},
//   									Sqs: &SqsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										QueueUrl: jsii.String("queueUrl"),
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   							EventName: jsii.String("eventName"),
//   						},
//   					},
//   				},
//   				OnInput: &OnInputProperty{
//   					Events: []interface{}{
//   						&EventProperty{
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyType: jsii.String("hashKeyType"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   										TableName: jsii.String("tableName"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TableName: jsii.String("tableName"),
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//   											},
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//   										},
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TargetArn: jsii.String("targetArn"),
//   									},
//   									Sqs: &SqsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										QueueUrl: jsii.String("queueUrl"),
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   							EventName: jsii.String("eventName"),
//   						},
//   					},
//   					TransitionEvents: []interface{}{
//   						&TransitionEventProperty{
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyType: jsii.String("hashKeyType"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   										TableName: jsii.String("tableName"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TableName: jsii.String("tableName"),
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//   											},
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//   										},
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TargetArn: jsii.String("targetArn"),
//   									},
//   									Sqs: &SqsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										QueueUrl: jsii.String("queueUrl"),
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   							EventName: jsii.String("eventName"),
//   							NextState: jsii.String("nextState"),
//   						},
//   					},
//   				},
//   				StateName: jsii.String("stateName"),
//   			},
//   		},
//   	},
//   	DetectorModelDescription: jsii.String("detectorModelDescription"),
//   	DetectorModelName: jsii.String("detectorModelName"),
//   	EvaluationMethod: jsii.String("evaluationMethod"),
//   	Key: jsii.String("key"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html
//
type CfnDetectorModelPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDetectorModelMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDetectorModelPropsMixin
type jsiiProxy_CfnDetectorModelPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDetectorModelPropsMixin) Props() *CfnDetectorModelMixinProps {
	var returns *CfnDetectorModelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModelPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTEvents::DetectorModel`.
func NewCfnDetectorModelPropsMixin(props *CfnDetectorModelMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDetectorModelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDetectorModelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDetectorModelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotevents.mixins.CfnDetectorModelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTEvents::DetectorModel`.
func NewCfnDetectorModelPropsMixin_Override(c CfnDetectorModelPropsMixin, props *CfnDetectorModelMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotevents.mixins.CfnDetectorModelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDetectorModelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDetectorModelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotevents.mixins.CfnDetectorModelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDetectorModelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotevents.mixins.CfnDetectorModelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDetectorModelPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDetectorModelPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

