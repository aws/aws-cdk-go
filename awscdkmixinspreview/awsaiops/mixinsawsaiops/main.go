package mixinsawsaiops

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aiops.mixins.CfnInvestigationGroupMixinProps",
		reflect.TypeOf((*CfnInvestigationGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_aiops.mixins.CfnInvestigationGroupPropsMixin",
		reflect.TypeOf((*CfnInvestigationGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInvestigationGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aiops.mixins.CfnInvestigationGroupPropsMixin.ChatbotNotificationChannelProperty",
		reflect.TypeOf((*CfnInvestigationGroupPropsMixin_ChatbotNotificationChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aiops.mixins.CfnInvestigationGroupPropsMixin.CrossAccountConfigurationProperty",
		reflect.TypeOf((*CfnInvestigationGroupPropsMixin_CrossAccountConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_aiops.mixins.CfnInvestigationGroupPropsMixin.EncryptionConfigMapProperty",
		reflect.TypeOf((*CfnInvestigationGroupPropsMixin_EncryptionConfigMapProperty)(nil)).Elem(),
	)
}
