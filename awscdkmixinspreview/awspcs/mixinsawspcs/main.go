package mixinsawspcs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin",
		reflect.TypeOf((*CfnClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.AccountingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AccountingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.AuthKeyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AuthKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.EndpointProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.ErrorInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ErrorInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.NetworkingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_NetworkingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.SchedulerProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SchedulerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.SlurmConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SlurmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPropsMixin.SlurmCustomSettingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SlurmCustomSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupMixinProps",
		reflect.TypeOf((*CfnComputeNodeGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComputeNodeGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.CustomLaunchTemplateProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_CustomLaunchTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.ErrorInfoProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_ErrorInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.InstanceConfigProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_InstanceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.ScalingConfigurationProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_ScalingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.SlurmConfigurationProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_SlurmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.SlurmCustomSettingProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_SlurmCustomSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnComputeNodeGroupPropsMixin.SpotOptionsProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_SpotOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueueMixinProps",
		reflect.TypeOf((*CfnQueueMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin",
		reflect.TypeOf((*CfnQueuePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueuePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin.ComputeNodeGroupConfigurationProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_ComputeNodeGroupConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin.ErrorInfoProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_ErrorInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin.SlurmConfigurationProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_SlurmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnQueuePropsMixin.SlurmCustomSettingProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_SlurmCustomSettingProperty)(nil)).Elem(),
	)
}
