package awspcs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin",
		reflect.TypeOf((*CfnClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.AccountingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AccountingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.AuthKeyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AuthKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.CgroupCustomSettingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_CgroupCustomSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.EndpointProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.ErrorInfoProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ErrorInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.JwtAuthProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_JwtAuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.JwtKeyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_JwtKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.NetworkingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_NetworkingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.SchedulerProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SchedulerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.SlurmConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SlurmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.SlurmCustomSettingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SlurmCustomSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.SlurmRestProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SlurmRestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnClusterPropsMixin.SlurmdbdCustomSettingProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SlurmdbdCustomSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnComputeNodeGroupMixinProps",
		reflect.TypeOf((*CfnComputeNodeGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnComputeNodeGroupPropsMixin",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComputeNodeGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnComputeNodeGroupPropsMixin.CustomLaunchTemplateProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_CustomLaunchTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnComputeNodeGroupPropsMixin.ErrorInfoProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_ErrorInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnComputeNodeGroupPropsMixin.InstanceConfigProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_InstanceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnComputeNodeGroupPropsMixin.ScalingConfigurationProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_ScalingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnComputeNodeGroupPropsMixin.SlurmConfigurationProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_SlurmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnComputeNodeGroupPropsMixin.SlurmCustomSettingProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_SlurmCustomSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnComputeNodeGroupPropsMixin.SpotOptionsProperty",
		reflect.TypeOf((*CfnComputeNodeGroupPropsMixin_SpotOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnQueueMixinProps",
		reflect.TypeOf((*CfnQueueMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnQueuePropsMixin",
		reflect.TypeOf((*CfnQueuePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueuePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnQueuePropsMixin.ComputeNodeGroupConfigurationProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_ComputeNodeGroupConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnQueuePropsMixin.ErrorInfoProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_ErrorInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnQueuePropsMixin.SlurmConfigurationProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_SlurmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_pcs.CfnQueuePropsMixin.SlurmCustomSettingProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_SlurmCustomSettingProperty)(nil)).Elem(),
	)
}
