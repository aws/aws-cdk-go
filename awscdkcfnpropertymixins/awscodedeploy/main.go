package awscodedeploy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentConfigMixinProps",
		reflect.TypeOf((*CfnDeploymentConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentConfigPropsMixin",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeploymentConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentConfigPropsMixin.MinimumHealthyHostsPerZoneProperty",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin_MinimumHealthyHostsPerZoneProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentConfigPropsMixin.MinimumHealthyHostsProperty",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin_MinimumHealthyHostsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentConfigPropsMixin.TimeBasedCanaryProperty",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin_TimeBasedCanaryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentConfigPropsMixin.TimeBasedLinearProperty",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin_TimeBasedLinearProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentConfigPropsMixin.TrafficRoutingConfigProperty",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin_TrafficRoutingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentConfigPropsMixin.ZonalConfigProperty",
		reflect.TypeOf((*CfnDeploymentConfigPropsMixin_ZonalConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupMixinProps",
		reflect.TypeOf((*CfnDeploymentGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeploymentGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.AlarmConfigurationProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_AlarmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.AlarmProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_AlarmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.AutoRollbackConfigurationProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_AutoRollbackConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.BlueGreenDeploymentConfigurationProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_BlueGreenDeploymentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.BlueInstanceTerminationOptionProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_BlueInstanceTerminationOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.DeploymentProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_DeploymentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.DeploymentReadyOptionProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_DeploymentReadyOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.DeploymentStyleProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_DeploymentStyleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.EC2TagFilterProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_EC2TagFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.EC2TagSetListObjectProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_EC2TagSetListObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.EC2TagSetProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_EC2TagSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.ECSServiceProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_ECSServiceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.ELBInfoProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_ELBInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.GitHubLocationProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_GitHubLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.GreenFleetProvisioningOptionProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_GreenFleetProvisioningOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.LoadBalancerInfoProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_LoadBalancerInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.OnPremisesTagSetListObjectProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_OnPremisesTagSetListObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.OnPremisesTagSetProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_OnPremisesTagSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.RevisionLocationProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_RevisionLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.TagFilterProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_TagFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.TargetGroupInfoProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_TargetGroupInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.TargetGroupPairInfoProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_TargetGroupPairInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.TrafficRouteProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_TrafficRouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codedeploy.CfnDeploymentGroupPropsMixin.TriggerConfigProperty",
		reflect.TypeOf((*CfnDeploymentGroupPropsMixin_TriggerConfigProperty)(nil)).Elem(),
	)
}
