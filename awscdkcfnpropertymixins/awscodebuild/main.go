package awscodebuild

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnFleetMixinProps",
		reflect.TypeOf((*CfnFleetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnFleetPropsMixin",
		reflect.TypeOf((*CfnFleetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFleetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnFleetPropsMixin.ComputeConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ComputeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnFleetPropsMixin.FleetProxyRuleProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_FleetProxyRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnFleetPropsMixin.ProxyConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ProxyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnFleetPropsMixin.ScalingConfigurationInputProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ScalingConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnFleetPropsMixin.TargetTrackingScalingConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_TargetTrackingScalingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnFleetPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.ArtifactsProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ArtifactsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.BatchRestrictionsProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_BatchRestrictionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.BuildStatusConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_BuildStatusConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.CloudWatchLogsConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_CloudWatchLogsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.DockerServerProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_DockerServerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.EnvironmentProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_EnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.EnvironmentVariableProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_EnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.GitSubmodulesConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_GitSubmodulesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.LogsConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_LogsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.ProjectBuildBatchConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProjectBuildBatchConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.ProjectCacheProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProjectCacheProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.ProjectFileSystemLocationProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProjectFileSystemLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.ProjectFleetProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProjectFleetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.ProjectSourceVersionProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProjectSourceVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.ProjectTriggersProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProjectTriggersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.PullRequestBuildPolicyProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_PullRequestBuildPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.RegistryCredentialProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_RegistryCredentialProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.S3LogsConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_S3LogsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.ScopeConfigurationProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ScopeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.SourceAuthProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_SourceAuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.SourceProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnProjectPropsMixin.WebhookFilterProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_WebhookFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnReportGroupMixinProps",
		reflect.TypeOf((*CfnReportGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnReportGroupPropsMixin",
		reflect.TypeOf((*CfnReportGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReportGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnReportGroupPropsMixin.ReportExportConfigProperty",
		reflect.TypeOf((*CfnReportGroupPropsMixin_ReportExportConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnReportGroupPropsMixin.S3ReportExportConfigProperty",
		reflect.TypeOf((*CfnReportGroupPropsMixin_S3ReportExportConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnSourceCredentialMixinProps",
		reflect.TypeOf((*CfnSourceCredentialMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_codebuild.CfnSourceCredentialPropsMixin",
		reflect.TypeOf((*CfnSourceCredentialPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSourceCredentialPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
