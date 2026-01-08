package previewawscodebuildevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents",
		reflect.TypeOf((*ProjectEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "codeBuildBuildPhaseChangePattern", GoMethod: "CodeBuildBuildPhaseChangePattern"},
			_jsii_.MemberMethod{JsiiMethod: "codeBuildBuildStateChangePattern", GoMethod: "CodeBuildBuildStateChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_ProjectEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ProjectEvents_CodeBuildBuildPhaseChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.AdditionalInformation",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_AdditionalInformation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.AdditionalInformationItem",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_AdditionalInformationItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.Artifact",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_Artifact)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.Auth",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_Auth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.Cache",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_Cache)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.CodeBuildBuildPhaseChangeProps",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_CodeBuildBuildPhaseChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.Environment",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_Environment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.EnvironmentItem",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_EnvironmentItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.Logs",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_Logs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.NetworkInterface",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_NetworkInterface)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.Source",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_Source)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.VpcConfig",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_VpcConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildPhaseChange.VpcConfigItem",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildPhaseChange_VpcConfigItem)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ProjectEvents_CodeBuildBuildStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.AdditionalInformation",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_AdditionalInformation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.AdditionalInformationItem",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_AdditionalInformationItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.Artifact",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_Artifact)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.Auth",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_Auth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.Cache",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_Cache)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.CodeBuildBuildStateChangeProps",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_CodeBuildBuildStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.Environment",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_Environment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.EnvironmentItem",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_EnvironmentItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.Logs",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_Logs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.NetworkInterface",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_NetworkInterface)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.Source",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_Source)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.VpcConfig",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_VpcConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codebuild.events.ProjectEvents.CodeBuildBuildStateChange.VpcConfigItem",
		reflect.TypeOf((*ProjectEvents_CodeBuildBuildStateChange_VpcConfigItem)(nil)).Elem(),
	)
}
