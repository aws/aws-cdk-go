package previewawslookoutequipmentmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutequipment.mixins.CfnInferenceSchedulerMixinProps",
		reflect.TypeOf((*CfnInferenceSchedulerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lookoutequipment.mixins.CfnInferenceSchedulerPropsMixin",
		reflect.TypeOf((*CfnInferenceSchedulerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInferenceSchedulerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutequipment.mixins.CfnInferenceSchedulerPropsMixin.DataInputConfigurationProperty",
		reflect.TypeOf((*CfnInferenceSchedulerPropsMixin_DataInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutequipment.mixins.CfnInferenceSchedulerPropsMixin.DataOutputConfigurationProperty",
		reflect.TypeOf((*CfnInferenceSchedulerPropsMixin_DataOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutequipment.mixins.CfnInferenceSchedulerPropsMixin.InputNameConfigurationProperty",
		reflect.TypeOf((*CfnInferenceSchedulerPropsMixin_InputNameConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutequipment.mixins.CfnInferenceSchedulerPropsMixin.S3InputConfigurationProperty",
		reflect.TypeOf((*CfnInferenceSchedulerPropsMixin_S3InputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutequipment.mixins.CfnInferenceSchedulerPropsMixin.S3OutputConfigurationProperty",
		reflect.TypeOf((*CfnInferenceSchedulerPropsMixin_S3OutputConfigurationProperty)(nil)).Elem(),
	)
}
