package previewawsioteventsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiotevents/previewawsioteventsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an alarm model to monitor an AWS IoT Events input attribute.
//
// You can use the alarm to get notified when the value is outside a specified range. For more information, see [Create an alarm model](https://docs.aws.amazon.com/iotevents/latest/developerguide/create-alarms.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAlarmModelPropsMixin := awscdkmixinspreview.Mixins.NewCfnAlarmModelPropsMixin(&CfnAlarmModelMixinProps{
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
//   					HashKeyType: jsii.String("hashKeyType"),
//   					HashKeyValue: jsii.String("hashKeyValue"),
//   					Operation: jsii.String("operation"),
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					PayloadField: jsii.String("payloadField"),
//   					RangeKeyField: jsii.String("rangeKeyField"),
//   					RangeKeyType: jsii.String("rangeKeyType"),
//   					RangeKeyValue: jsii.String("rangeKeyValue"),
//   					TableName: jsii.String("tableName"),
//   				},
//   				DynamoDBv2: &DynamoDBv2Property{
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					TableName: jsii.String("tableName"),
//   				},
//   				Firehose: &FirehoseProperty{
//   					DeliveryStreamName: jsii.String("deliveryStreamName"),
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					Separator: jsii.String("separator"),
//   				},
//   				IotEvents: &IotEventsProperty{
//   					InputName: jsii.String("inputName"),
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
//   						Quality: jsii.String("quality"),
//   						Timestamp: &AssetPropertyTimestampProperty{
//   							OffsetInNanos: jsii.String("offsetInNanos"),
//   							TimeInSeconds: jsii.String("timeInSeconds"),
//   						},
//   						Value: &AssetPropertyVariantProperty{
//   							BooleanValue: jsii.String("booleanValue"),
//   							DoubleValue: jsii.String("doubleValue"),
//   							IntegerValue: jsii.String("integerValue"),
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   				},
//   				IotTopicPublish: &IotTopicPublishProperty{
//   					MqttTopic: jsii.String("mqttTopic"),
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Lambda: &LambdaProperty{
//   					FunctionArn: jsii.String("functionArn"),
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Sns: &SnsProperty{
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					TargetArn: jsii.String("targetArn"),
//   				},
//   				Sqs: &SqsProperty{
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					QueueUrl: jsii.String("queueUrl"),
//   					UseBase64: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	AlarmModelDescription: jsii.String("alarmModelDescription"),
//   	AlarmModelName: jsii.String("alarmModelName"),
//   	AlarmRule: &AlarmRuleProperty{
//   		SimpleRule: &SimpleRuleProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			InputProperty: jsii.String("inputProperty"),
//   			Threshold: jsii.String("threshold"),
//   		},
//   	},
//   	Key: jsii.String("key"),
//   	RoleArn: jsii.String("roleArn"),
//   	Severity: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html
//
type CfnAlarmModelPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAlarmModelMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAlarmModelPropsMixin
type jsiiProxy_CfnAlarmModelPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAlarmModelPropsMixin) Props() *CfnAlarmModelMixinProps {
	var returns *CfnAlarmModelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmModelPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTEvents::AlarmModel`.
func NewCfnAlarmModelPropsMixin(props *CfnAlarmModelMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAlarmModelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAlarmModelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAlarmModelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotevents.mixins.CfnAlarmModelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTEvents::AlarmModel`.
func NewCfnAlarmModelPropsMixin_Override(c CfnAlarmModelPropsMixin, props *CfnAlarmModelMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotevents.mixins.CfnAlarmModelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAlarmModelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAlarmModelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotevents.mixins.CfnAlarmModelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAlarmModelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotevents.mixins.CfnAlarmModelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAlarmModelPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAlarmModelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

