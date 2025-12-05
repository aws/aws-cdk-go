package previewawsdevopsagentmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceMixinProps",
		reflect.TypeOf((*CfnAgentSpaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpacePropsMixin",
		reflect.TypeOf((*CfnAgentSpacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAgentSpacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationMixinProps",
		reflect.TypeOf((*CfnAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin",
		reflect.TypeOf((*CfnAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.AWSConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_AWSConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.AWSResourceProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_AWSResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.DynatraceConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_DynatraceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.EventChannelConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_EventChannelConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.GitHubConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_GitHubConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.GitLabConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_GitLabConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.KeyValuePairProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_KeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.MCPServerConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.MCPServerDatadogConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerDatadogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.MCPServerNewRelicConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerNewRelicConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.MCPServerSplunkConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_MCPServerSplunkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.ServiceConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_ServiceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.ServiceNowConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_ServiceNowConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.SlackChannelProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_SlackChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.SlackConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_SlackConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.SlackTransmissionTargetProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_SlackTransmissionTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin.SourceAwsConfigurationProperty",
		reflect.TypeOf((*CfnAssociationPropsMixin_SourceAwsConfigurationProperty)(nil)).Elem(),
	)
}
