package awsdevopsagent

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAgentSpaceMixinProps",
		reflect.TypeOf((*CfnAgentSpaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAgentSpacePropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAgentSpacePropsMixin.IamAuthConfigurationProperty",
		reflect.TypeOf((*CfnAgentSpacePropsMixin_IamAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAgentSpacePropsMixin.IdcAuthConfigurationProperty",
		reflect.TypeOf((*CfnAgentSpacePropsMixin_IdcAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAgentSpacePropsMixin.OperatorAppProperty",
		reflect.TypeOf((*CfnAgentSpacePropsMixin_OperatorAppProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationMixinProps",
		reflect.TypeOf((*CfnAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin",
		reflect.TypeOf((*CfnAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.AWSConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_AWSConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.AWSResourceProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_AWSResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.AzureConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_AzureConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.DynatraceConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_DynatraceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.EventChannelConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_EventChannelConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.GitHubConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_GitHubConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.GitLabConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_GitLabConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.KeyValuePairProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_KeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.MCPServerConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.MCPServerDatadogConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerDatadogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.MCPServerNewRelicConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerNewRelicConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.MCPServerSigV4ConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerSigV4ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.MCPServerSplunkConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerSplunkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.PagerDutyConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_PagerDutyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.ServiceConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_ServiceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.ServiceNowConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_ServiceNowConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.SlackChannelProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_SlackChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.SlackConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_SlackConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.SlackTransmissionTargetProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_SlackTransmissionTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssociationPropsMixin.SourceAwsConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_SourceAwsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnPrivateConnectionMixinProps",
		reflect.TypeOf((*CfnPrivateConnectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnPrivateConnectionPropsMixin",
		reflect.TypeOf((*CfnPrivateConnectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrivateConnectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnPrivateConnectionPropsMixin.ConnectionConfigurationProperty",
		reflect.TypeOf((*CfnPrivateConnectionPropsMixin_ConnectionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnPrivateConnectionPropsMixin.SelfManagedModeProperty",
		reflect.TypeOf((*CfnPrivateConnectionPropsMixin_SelfManagedModeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnPrivateConnectionPropsMixin.ServiceManagedModeProperty",
		reflect.TypeOf((*CfnPrivateConnectionPropsMixin_ServiceManagedModeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServiceMixinProps",
		reflect.TypeOf((*CfnServiceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin",
		reflect.TypeOf((*CfnServicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.AdditionalServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_AdditionalServiceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.ApiKeyDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ApiKeyDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.AzureIdentityServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_AzureIdentityServiceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.BearerTokenDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_BearerTokenDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.DynatraceAuthorizationConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DynatraceAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.DynatraceServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_DynatraceServiceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.GitLabDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_GitLabDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.MCPServerAuthorizationConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_MCPServerAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.MCPServerDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_MCPServerDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.MCPServerOAuthClientCredentialsConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_MCPServerOAuthClientCredentialsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.MCPServerSigV4AuthorizationConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_MCPServerSigV4AuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.MCPServerSigV4DetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_MCPServerSigV4DetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.MCPServerSplunkAuthorizationConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_MCPServerSplunkAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.MCPServerSplunkDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_MCPServerSplunkDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.NewRelicApiKeyConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_NewRelicApiKeyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.NewRelicAuthorizationConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_NewRelicAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.NewRelicServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_NewRelicServiceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.OAuthClientDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_OAuthClientDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.PagerDutyAuthorizationConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_PagerDutyAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.PagerDutyDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_PagerDutyDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.RegisteredAzureIdentityDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredAzureIdentityDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.RegisteredDynatraceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredDynatraceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.RegisteredGitLabServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredGitLabServiceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.RegisteredMCPServerDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredMCPServerDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.RegisteredMCPServerSigV4DetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredMCPServerSigV4DetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.RegisteredNewRelicDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredNewRelicDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.RegisteredPagerDutyDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredPagerDutyDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.RegisteredServiceNowDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_RegisteredServiceNowDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.ServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.ServiceNowAuthorizationConfigProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceNowAuthorizationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnServicePropsMixin.ServiceNowServiceDetailsProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceNowServiceDetailsProperty)(nil)).Elem(),
	)
}
