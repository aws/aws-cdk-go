package previewawsconnectcampaignsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaigns.mixins.CfnCampaignMixinProps",
		reflect.TypeOf((*CfnCampaignMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_connectcampaigns.mixins.CfnCampaignPropsMixin",
		reflect.TypeOf((*CfnCampaignPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCampaignPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaigns.mixins.CfnCampaignPropsMixin.AgentlessDialerConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_AgentlessDialerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaigns.mixins.CfnCampaignPropsMixin.AnswerMachineDetectionConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_AnswerMachineDetectionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaigns.mixins.CfnCampaignPropsMixin.DialerConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_DialerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaigns.mixins.CfnCampaignPropsMixin.OutboundCallConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_OutboundCallConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaigns.mixins.CfnCampaignPropsMixin.PredictiveDialerConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_PredictiveDialerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaigns.mixins.CfnCampaignPropsMixin.ProgressiveDialerConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_ProgressiveDialerConfigProperty)(nil)).Elem(),
	)
}
