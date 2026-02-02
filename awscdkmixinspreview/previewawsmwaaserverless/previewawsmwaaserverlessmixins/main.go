package previewawsmwaaserverlessmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mwaaserverless.mixins.CfnWorkflowMixinProps",
		reflect.TypeOf((*CfnWorkflowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_mwaaserverless.mixins.CfnWorkflowPropsMixin",
		reflect.TypeOf((*CfnWorkflowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkflowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mwaaserverless.mixins.CfnWorkflowPropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnWorkflowPropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mwaaserverless.mixins.CfnWorkflowPropsMixin.LoggingConfigurationProperty",
		reflect.TypeOf((*CfnWorkflowPropsMixin_LoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mwaaserverless.mixins.CfnWorkflowPropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnWorkflowPropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mwaaserverless.mixins.CfnWorkflowPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnWorkflowPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_mwaaserverless.mixins.CfnWorkflowPropsMixin.ScheduleConfigurationProperty",
		reflect.TypeOf((*CfnWorkflowPropsMixin_ScheduleConfigurationProperty)(nil)).Elem(),
	)
}
