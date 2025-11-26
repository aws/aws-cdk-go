package previewawsdlmmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyMixinProps",
		reflect.TypeOf((*CfnLifecyclePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLifecyclePolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.ArchiveRetainRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ArchiveRetainRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.ArchiveRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ArchiveRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.CreateRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_CreateRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.CrossRegionCopyActionProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_CrossRegionCopyActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.CrossRegionCopyDeprecateRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_CrossRegionCopyDeprecateRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.CrossRegionCopyRetainRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_CrossRegionCopyRetainRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.CrossRegionCopyRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_CrossRegionCopyRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.DeprecateRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_DeprecateRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.EventParametersProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_EventParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.EventSourceProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_EventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.ExclusionsProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ExclusionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.FastRestoreRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_FastRestoreRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.ParametersProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.PolicyDetailsProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_PolicyDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.RetainRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_RetainRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.RetentionArchiveTierProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_RetentionArchiveTierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.ScheduleProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.ScriptProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ScriptProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin.ShareRuleProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ShareRuleProperty)(nil)).Elem(),
	)
}
