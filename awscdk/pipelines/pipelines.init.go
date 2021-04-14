package pipelines

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.AddManualApprovalOptions",
		reflect.TypeOf((*AddManualApprovalOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.AddStackOptions",
		reflect.TypeOf((*AddStackOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.AddStageOptions",
		reflect.TypeOf((*AddStageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.AdditionalArtifact",
		reflect.TypeOf((*AdditionalArtifact)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.AssetPublishingCommand",
		reflect.TypeOf((*AssetPublishingCommand)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.pipelines.AssetType",
		reflect.TypeOf((*AssetType)(nil)).Elem(),
		map[string]interface{}{
			"FILE": AssetType_FILE,
			"DOCKER_IMAGE": AssetType_DOCKER_IMAGE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.pipelines.CdkPipeline",
		reflect.TypeOf((*CdkPipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addApplicationStage", GoMethod: "AddApplicationStage"},
			_jsii_.MemberMethod{JsiiMethod: "addStage", GoMethod: "AddStage"},
			_jsii_.MemberProperty{JsiiProperty: "codePipeline", GoGetter: "CodePipeline"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "stackOutput", GoMethod: "StackOutput"},
			_jsii_.MemberMethod{JsiiMethod: "stage", GoMethod: "Stage"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CdkPipeline{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.CdkPipelineProps",
		reflect.TypeOf((*CdkPipelineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.CdkStackActionFromArtifactOptions",
		reflect.TypeOf((*CdkStackActionFromArtifactOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.pipelines.CdkStage",
		reflect.TypeOf((*CdkStage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addActions", GoMethod: "AddActions"},
			_jsii_.MemberMethod{JsiiMethod: "addApplication", GoMethod: "AddApplication"},
			_jsii_.MemberMethod{JsiiMethod: "addManualApprovalAction", GoMethod: "AddManualApprovalAction"},
			_jsii_.MemberMethod{JsiiMethod: "addStackArtifactDeployment", GoMethod: "AddStackArtifactDeployment"},
			_jsii_.MemberMethod{JsiiMethod: "deploysStack", GoMethod: "DeploysStack"},
			_jsii_.MemberMethod{JsiiMethod: "nextSequentialRunOrder", GoMethod: "NextSequentialRunOrder"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CdkStage{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.CdkStageProps",
		reflect.TypeOf((*CdkStageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.pipelines.DeployCdkStackAction",
		reflect.TypeOf((*DeployCdkStackAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyStackArtifactIds", GoGetter: "DependencyStackArtifactIds"},
			_jsii_.MemberProperty{JsiiProperty: "executeRunOrder", GoGetter: "ExecuteRunOrder"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberProperty{JsiiProperty: "prepareRunOrder", GoGetter: "PrepareRunOrder"},
			_jsii_.MemberProperty{JsiiProperty: "stackArtifactId", GoGetter: "StackArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
		},
		func() interface{} {
			j := jsiiProxy_DeployCdkStackAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscodepipelineIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.DeployCdkStackActionOptions",
		reflect.TypeOf((*DeployCdkStackActionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.DeployCdkStackActionProps",
		reflect.TypeOf((*DeployCdkStackActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.FromStackArtifactOptions",
		reflect.TypeOf((*FromStackArtifactOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.pipelines.IStageHost",
		reflect.TypeOf((*IStageHost)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "publishAsset", GoMethod: "PublishAsset"},
			_jsii_.MemberMethod{JsiiMethod: "stackOutputArtifact", GoMethod: "StackOutputArtifact"},
		},
		func() interface{} {
			return &jsiiProxy_IStageHost{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.pipelines.PublishAssetsAction",
		reflect.TypeOf((*PublishAssetsAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "addPublishCommand", GoMethod: "AddPublishCommand"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PublishAssetsAction{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awscodepipelineIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.PublishAssetsActionProps",
		reflect.TypeOf((*PublishAssetsActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.pipelines.ShellScriptAction",
		reflect.TypeOf((*ShellScriptAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
		},
		func() interface{} {
			j := jsiiProxy_ShellScriptAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscodepipelineIAction)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.ShellScriptActionProps",
		reflect.TypeOf((*ShellScriptActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.pipelines.SimpleSynthAction",
		reflect.TypeOf((*SimpleSynthAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
		},
		func() interface{} {
			j := jsiiProxy_SimpleSynthAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscodepipelineIAction)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.SimpleSynthActionProps",
		reflect.TypeOf((*SimpleSynthActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.SimpleSynthOptions",
		reflect.TypeOf((*SimpleSynthOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.pipelines.StackOutput",
		reflect.TypeOf((*StackOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "artifactFile", GoGetter: "ArtifactFile"},
			_jsii_.MemberProperty{JsiiProperty: "outputName", GoGetter: "OutputName"},
		},
		func() interface{} {
			return &jsiiProxy_StackOutput{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.StandardNpmSynthOptions",
		reflect.TypeOf((*StandardNpmSynthOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.StandardYarnSynthOptions",
		reflect.TypeOf((*StandardYarnSynthOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.pipelines.UpdatePipelineAction",
		reflect.TypeOf((*UpdatePipelineAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_UpdatePipelineAction{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awscodepipelineIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.pipelines.UpdatePipelineActionProps",
		reflect.TypeOf((*UpdatePipelineActionProps)(nil)).Elem(),
	)
}
