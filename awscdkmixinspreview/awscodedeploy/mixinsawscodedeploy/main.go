package mixinsawscodedeploy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentConfigMixinProps",
		reflect.TypeOf((*CfnDeploymentConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentConfigPropsMixin",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeploymentConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentConfigPropsMixin.MinimumHealthyHostsPerZoneProperty",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin_MinimumHealthyHostsPerZoneProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentConfigPropsMixin.MinimumHealthyHostsProperty",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin_MinimumHealthyHostsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentConfigPropsMixin.TimeBasedCanaryProperty",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin_TimeBasedCanaryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentConfigPropsMixin.TimeBasedLinearProperty",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin_TimeBasedLinearProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentConfigPropsMixin.TrafficRoutingConfigProperty",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin_TrafficRoutingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentConfigPropsMixin.ZonalConfigProperty",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin_ZonalConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupMixinProps",
		reflect.TypeOf((*CfnDeploymentGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeploymentGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.AlarmConfigurationProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_AlarmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.AlarmProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_AlarmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.AutoRollbackConfigurationProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_AutoRollbackConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.BlueGreenDeploymentConfigurationProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_BlueGreenDeploymentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.BlueInstanceTerminationOptionProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_BlueInstanceTerminationOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.DeploymentProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_DeploymentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.DeploymentReadyOptionProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_DeploymentReadyOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.DeploymentStyleProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_DeploymentStyleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.EC2TagFilterProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_EC2TagFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.EC2TagSetListObjectProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_EC2TagSetListObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.EC2TagSetProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_EC2TagSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.ECSServiceProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_ECSServiceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.ELBInfoProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_ELBInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.GitHubLocationProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_GitHubLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.GreenFleetProvisioningOptionProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_GreenFleetProvisioningOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.LoadBalancerInfoProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_LoadBalancerInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.OnPremisesTagSetListObjectProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_OnPremisesTagSetListObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.OnPremisesTagSetProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_OnPremisesTagSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.RevisionLocationProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_RevisionLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.TagFilterProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_TagFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.TargetGroupInfoProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_TargetGroupInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.TargetGroupPairInfoProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_TargetGroupPairInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.TrafficRouteProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_TrafficRouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin.TriggerConfigProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_TriggerConfigProperty)(nil)).Elem(),
	)
}
