package awsiotevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelMixinProps",
		reflect.TypeOf((*CfnAlarmModelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin",
		reflect.TypeOf((*CfnAlarmModelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlarmModelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.AcknowledgeFlowProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_AcknowledgeFlowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.AlarmActionProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_AlarmActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.AlarmCapabilitiesProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_AlarmCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.AlarmEventActionsProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_AlarmEventActionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.AlarmRuleProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_AlarmRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.AssetPropertyTimestampProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_AssetPropertyTimestampProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.AssetPropertyValueProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_AssetPropertyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.AssetPropertyVariantProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_AssetPropertyVariantProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.DynamoDBProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_DynamoDBProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.DynamoDBv2Property",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_DynamoDBv2Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.FirehoseProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_FirehoseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.InitializationConfigurationProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_InitializationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.IotEventsProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_IotEventsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.IotSiteWiseProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_IotSiteWiseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.IotTopicPublishProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_IotTopicPublishProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.LambdaProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_LambdaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.PayloadProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_PayloadProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.SimpleRuleProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_SimpleRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.SnsProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_SnsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnAlarmModelPropsMixin.SqsProperty",
		reflect.TypeOf((*CfnAlarmModelPropsMixin_SqsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelMixinProps",
		reflect.TypeOf((*CfnDetectorModelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin",
		reflect.TypeOf((*CfnDetectorModelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDetectorModelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.AssetPropertyTimestampProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_AssetPropertyTimestampProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.AssetPropertyValueProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_AssetPropertyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.AssetPropertyVariantProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_AssetPropertyVariantProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.ClearTimerProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_ClearTimerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.DetectorModelDefinitionProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_DetectorModelDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.DynamoDBProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_DynamoDBProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.DynamoDBv2Property",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_DynamoDBv2Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.EventProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_EventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.FirehoseProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_FirehoseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.IotEventsProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_IotEventsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.IotSiteWiseProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_IotSiteWiseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.IotTopicPublishProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_IotTopicPublishProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.LambdaProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_LambdaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.OnEnterProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_OnEnterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.OnExitProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_OnExitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.OnInputProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_OnInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.PayloadProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_PayloadProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.ResetTimerProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_ResetTimerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.SetTimerProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_SetTimerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.SetVariableProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_SetVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.SnsProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_SnsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.SqsProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_SqsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.StateProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_StateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnDetectorModelPropsMixin.TransitionEventProperty",
		reflect.TypeOf((*CfnDetectorModelPropsMixin_TransitionEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnInputMixinProps",
		reflect.TypeOf((*CfnInputMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnInputPropsMixin",
		reflect.TypeOf((*CfnInputPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInputPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnInputPropsMixin.AttributeProperty",
		reflect.TypeOf((*CfnInputPropsMixin_AttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotevents.CfnInputPropsMixin.InputDefinitionProperty",
		reflect.TypeOf((*CfnInputPropsMixin_InputDefinitionProperty)(nil)).Elem(),
	)
}
