package mixinsawscodebuild

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnFleetMixinProps",
		reflect.TypeOf((*CfnFleetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnFleetPropsMixin",
		reflect.TypeOf((*CfnFleetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFleetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnFleetPropsMixin.ComputeConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ComputeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnFleetPropsMixin.FleetProxyRuleProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_FleetProxyRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnFleetPropsMixin.ProxyConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ProxyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnFleetPropsMixin.ScalingConfigurationInputProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ScalingConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnFleetPropsMixin.TargetTrackingScalingConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_TargetTrackingScalingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnFleetPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.ArtifactsProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ArtifactsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.BatchRestrictionsProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_BatchRestrictionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.BuildStatusConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_BuildStatusConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.CloudWatchLogsConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_CloudWatchLogsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.DockerServerProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_DockerServerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.EnvironmentProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_EnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.EnvironmentVariableProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_EnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.GitSubmodulesConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_GitSubmodulesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.LogsConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_LogsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.ProjectBuildBatchConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProjectBuildBatchConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.ProjectCacheProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProjectCacheProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.ProjectFileSystemLocationProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProjectFileSystemLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.ProjectFleetProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProjectFleetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.ProjectSourceVersionProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProjectSourceVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.ProjectTriggersProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ProjectTriggersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.PullRequestBuildPolicyProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_PullRequestBuildPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.RegistryCredentialProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_RegistryCredentialProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.S3LogsConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_S3LogsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.ScopeConfigurationProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_ScopeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.SourceAuthProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_SourceAuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.SourceProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin.WebhookFilterProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_WebhookFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnReportGroupMixinProps",
		reflect.TypeOf((*CfnReportGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnReportGroupPropsMixin",
		reflect.TypeOf((*CfnReportGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReportGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnReportGroupPropsMixin.ReportExportConfigProperty",
		reflect.TypeOf((*CfnReportGroupPropsMixin_ReportExportConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnReportGroupPropsMixin.S3ReportExportConfigProperty",
		reflect.TypeOf((*CfnReportGroupPropsMixin_S3ReportExportConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnSourceCredentialMixinProps",
		reflect.TypeOf((*CfnSourceCredentialMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnSourceCredentialPropsMixin",
		reflect.TypeOf((*CfnSourceCredentialPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSourceCredentialPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
