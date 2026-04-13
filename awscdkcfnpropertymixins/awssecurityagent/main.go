package awssecurityagent

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnAgentSpaceMixinProps",
		reflect.TypeOf((*CfnAgentSpaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnAgentSpacePropsMixin",
		reflect.TypeOf((*CfnAgentSpacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAgentSpacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnAgentSpacePropsMixin.AWSResourcesProperty",
		reflect.TypeOf((*CfnAgentSpacePropsMixin_AWSResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnAgentSpacePropsMixin.CodeReviewSettingsProperty",
		reflect.TypeOf((*CfnAgentSpacePropsMixin_CodeReviewSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnAgentSpacePropsMixin.IntegratedResourceProperty",
		reflect.TypeOf((*CfnAgentSpacePropsMixin_IntegratedResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnAgentSpacePropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnAgentSpacePropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnApplicationPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnApplicationPropsMixin.IdCConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_IdCConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestMixinProps",
		reflect.TypeOf((*CfnPentestMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin",
		reflect.TypeOf((*CfnPentestPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPentestPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin.ActorProperty",
		reflect.TypeOf((*CfnPentestPropsMixin_ActorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin.AssetsProperty",
		reflect.TypeOf((*CfnPentestPropsMixin_AssetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin.AuthenticationProperty",
		reflect.TypeOf((*CfnPentestPropsMixin_AuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin.CloudWatchLogProperty",
		reflect.TypeOf((*CfnPentestPropsMixin_CloudWatchLogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin.CustomHeaderProperty",
		reflect.TypeOf((*CfnPentestPropsMixin_CustomHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin.DocumentInfoProperty",
		reflect.TypeOf((*CfnPentestPropsMixin_DocumentInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin.EndpointProperty",
		reflect.TypeOf((*CfnPentestPropsMixin_EndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin.IntegratedRepositoryProperty",
		reflect.TypeOf((*CfnPentestPropsMixin_IntegratedRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin.NetworkTrafficConfigProperty",
		reflect.TypeOf((*CfnPentestPropsMixin_NetworkTrafficConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin.NetworkTrafficRuleProperty",
		reflect.TypeOf((*CfnPentestPropsMixin_NetworkTrafficRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin.SourceCodeRepositoryProperty",
		reflect.TypeOf((*CfnPentestPropsMixin_SourceCodeRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnPentestPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnPentestPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnTargetDomainMixinProps",
		reflect.TypeOf((*CfnTargetDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnTargetDomainPropsMixin",
		reflect.TypeOf((*CfnTargetDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTargetDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnTargetDomainPropsMixin.DnsVerificationProperty",
		reflect.TypeOf((*CfnTargetDomainPropsMixin_DnsVerificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnTargetDomainPropsMixin.HttpVerificationProperty",
		reflect.TypeOf((*CfnTargetDomainPropsMixin_HttpVerificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_securityagent.CfnTargetDomainPropsMixin.VerificationDetailsProperty",
		reflect.TypeOf((*CfnTargetDomainPropsMixin_VerificationDetailsProperty)(nil)).Elem(),
	)
}
