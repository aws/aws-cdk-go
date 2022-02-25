package pipelines

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"monocdk.pipelines.AddManualApprovalOptions",
		reflect.TypeOf((*AddManualApprovalOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.AddStackOptions",
		reflect.TypeOf((*AddStackOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.AddStageOptions",
		reflect.TypeOf((*AddStageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.AddStageOpts",
		reflect.TypeOf((*AddStageOpts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.AdditionalArtifact",
		reflect.TypeOf((*AdditionalArtifact)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.ArtifactMap",
		reflect.TypeOf((*ArtifactMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toCodePipeline", GoMethod: "ToCodePipeline"},
		},
		func() interface{} {
			return &jsiiProxy_ArtifactMap{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.AssetPublishingCommand",
		reflect.TypeOf((*AssetPublishingCommand)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.pipelines.AssetType",
		reflect.TypeOf((*AssetType)(nil)).Elem(),
		map[string]interface{}{
			"FILE": AssetType_FILE,
			"DOCKER_IMAGE": AssetType_DOCKER_IMAGE,
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.BaseStageOptions",
		reflect.TypeOf((*BaseStageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.CdkPipeline",
		reflect.TypeOf((*CdkPipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addApplicationStage", GoMethod: "AddApplicationStage"},
			_jsii_.MemberMethod{JsiiMethod: "addStage", GoMethod: "AddStage"},
			_jsii_.MemberProperty{JsiiProperty: "codePipeline", GoGetter: "CodePipeline"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "stackOutput", GoMethod: "StackOutput"},
			_jsii_.MemberMethod{JsiiMethod: "stage", GoMethod: "Stage"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_CdkPipeline{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.CdkPipelineProps",
		reflect.TypeOf((*CdkPipelineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.CdkStackActionFromArtifactOptions",
		reflect.TypeOf((*CdkStackActionFromArtifactOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.CdkStage",
		reflect.TypeOf((*CdkStage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addActions", GoMethod: "AddActions"},
			_jsii_.MemberMethod{JsiiMethod: "addApplication", GoMethod: "AddApplication"},
			_jsii_.MemberMethod{JsiiMethod: "addManualApprovalAction", GoMethod: "AddManualApprovalAction"},
			_jsii_.MemberMethod{JsiiMethod: "addStackArtifactDeployment", GoMethod: "AddStackArtifactDeployment"},
			_jsii_.MemberMethod{JsiiMethod: "deploysStack", GoMethod: "DeploysStack"},
			_jsii_.MemberMethod{JsiiMethod: "nextSequentialRunOrder", GoMethod: "NextSequentialRunOrder"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_CdkStage{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.CdkStageProps",
		reflect.TypeOf((*CdkStageProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.CodeBuildOptions",
		reflect.TypeOf((*CodeBuildOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.CodeBuildStep",
		reflect.TypeOf((*CodeBuildStep)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependencyFileSet", GoMethod: "AddDependencyFileSet"},
			_jsii_.MemberMethod{JsiiMethod: "addOutputDirectory", GoMethod: "AddOutputDirectory"},
			_jsii_.MemberMethod{JsiiMethod: "addStepDependency", GoMethod: "AddStepDependency"},
			_jsii_.MemberProperty{JsiiProperty: "buildEnvironment", GoGetter: "BuildEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "commands", GoGetter: "Commands"},
			_jsii_.MemberMethod{JsiiMethod: "configurePrimaryOutput", GoMethod: "ConfigurePrimaryOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyFileSets", GoGetter: "DependencyFileSets"},
			_jsii_.MemberMethod{JsiiMethod: "discoverReferencedOutputs", GoMethod: "DiscoverReferencedOutputs"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "envFromCfnOutputs", GoGetter: "EnvFromCfnOutputs"},
			_jsii_.MemberMethod{JsiiMethod: "exportedVariable", GoMethod: "ExportedVariable"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputs", GoGetter: "Inputs"},
			_jsii_.MemberProperty{JsiiProperty: "installCommands", GoGetter: "InstallCommands"},
			_jsii_.MemberProperty{JsiiProperty: "isSource", GoGetter: "IsSource"},
			_jsii_.MemberProperty{JsiiProperty: "outputs", GoGetter: "Outputs"},
			_jsii_.MemberProperty{JsiiProperty: "partialBuildSpec", GoGetter: "PartialBuildSpec"},
			_jsii_.MemberProperty{JsiiProperty: "primaryOutput", GoGetter: "PrimaryOutput"},
			_jsii_.MemberMethod{JsiiMethod: "primaryOutputDirectory", GoMethod: "PrimaryOutputDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectName", GoGetter: "ProjectName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "rolePolicyStatements", GoGetter: "RolePolicyStatements"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "subnetSelection", GoGetter: "SubnetSelection"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_CodeBuildStep{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ShellStep)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.CodeBuildStepProps",
		reflect.TypeOf((*CodeBuildStepProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.CodeCommitSourceOptions",
		reflect.TypeOf((*CodeCommitSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.CodePipeline",
		reflect.TypeOf((*CodePipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStage", GoMethod: "AddStage"},
			_jsii_.MemberMethod{JsiiMethod: "addWave", GoMethod: "AddWave"},
			_jsii_.MemberMethod{JsiiMethod: "buildPipeline", GoMethod: "BuildPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "cloudAssemblyFileSet", GoGetter: "CloudAssemblyFileSet"},
			_jsii_.MemberMethod{JsiiMethod: "doBuildPipeline", GoMethod: "DoBuildPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "pipeline", GoGetter: "Pipeline"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "synth", GoGetter: "Synth"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "synthProject", GoGetter: "SynthProject"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberProperty{JsiiProperty: "waves", GoGetter: "Waves"},
		},
		func() interface{} {
			j := jsiiProxy_CodePipeline{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PipelineBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.CodePipelineActionFactoryResult",
		reflect.TypeOf((*CodePipelineActionFactoryResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.CodePipelineFileSet",
		reflect.TypeOf((*CodePipelineFileSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "primaryOutput", GoGetter: "PrimaryOutput"},
			_jsii_.MemberMethod{JsiiMethod: "producedBy", GoMethod: "ProducedBy"},
			_jsii_.MemberProperty{JsiiProperty: "producer", GoGetter: "Producer"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CodePipelineFileSet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FileSet)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.CodePipelineProps",
		reflect.TypeOf((*CodePipelineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.CodePipelineSource",
		reflect.TypeOf((*CodePipelineSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependencyFileSet", GoMethod: "AddDependencyFileSet"},
			_jsii_.MemberMethod{JsiiMethod: "addStepDependency", GoMethod: "AddStepDependency"},
			_jsii_.MemberMethod{JsiiMethod: "configurePrimaryOutput", GoMethod: "ConfigurePrimaryOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyFileSets", GoGetter: "DependencyFileSets"},
			_jsii_.MemberMethod{JsiiMethod: "discoverReferencedOutputs", GoMethod: "DiscoverReferencedOutputs"},
			_jsii_.MemberMethod{JsiiMethod: "getAction", GoMethod: "GetAction"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "isSource", GoGetter: "IsSource"},
			_jsii_.MemberProperty{JsiiProperty: "primaryOutput", GoGetter: "PrimaryOutput"},
			_jsii_.MemberMethod{JsiiMethod: "produceAction", GoMethod: "ProduceAction"},
			_jsii_.MemberMethod{JsiiMethod: "sourceAttribute", GoMethod: "SourceAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CodePipelineSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Step)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICodePipelineActionFactory)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.ConfirmPermissionsBroadening",
		reflect.TypeOf((*ConfirmPermissionsBroadening)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependencyFileSet", GoMethod: "AddDependencyFileSet"},
			_jsii_.MemberMethod{JsiiMethod: "addStepDependency", GoMethod: "AddStepDependency"},
			_jsii_.MemberMethod{JsiiMethod: "configurePrimaryOutput", GoMethod: "ConfigurePrimaryOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyFileSets", GoGetter: "DependencyFileSets"},
			_jsii_.MemberMethod{JsiiMethod: "discoverReferencedOutputs", GoMethod: "DiscoverReferencedOutputs"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "isSource", GoGetter: "IsSource"},
			_jsii_.MemberProperty{JsiiProperty: "primaryOutput", GoGetter: "PrimaryOutput"},
			_jsii_.MemberMethod{JsiiMethod: "produceAction", GoMethod: "ProduceAction"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConfirmPermissionsBroadening{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Step)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICodePipelineActionFactory)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.ConnectionSourceOptions",
		reflect.TypeOf((*ConnectionSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.DeployCdkStackAction",
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
		"monocdk.pipelines.DeployCdkStackActionOptions",
		reflect.TypeOf((*DeployCdkStackActionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.DeployCdkStackActionProps",
		reflect.TypeOf((*DeployCdkStackActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.DockerCredential",
		reflect.TypeOf((*DockerCredential)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "usages", GoGetter: "Usages"},
		},
		func() interface{} {
			return &jsiiProxy_DockerCredential{}
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.pipelines.DockerCredentialUsage",
		reflect.TypeOf((*DockerCredentialUsage)(nil)).Elem(),
		map[string]interface{}{
			"SYNTH": DockerCredentialUsage_SYNTH,
			"SELF_UPDATE": DockerCredentialUsage_SELF_UPDATE,
			"ASSET_PUBLISHING": DockerCredentialUsage_ASSET_PUBLISHING,
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.ECRSourceOptions",
		reflect.TypeOf((*ECRSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.EcrDockerCredentialOptions",
		reflect.TypeOf((*EcrDockerCredentialOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.ExternalDockerCredentialOptions",
		reflect.TypeOf((*ExternalDockerCredentialOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.FileSet",
		reflect.TypeOf((*FileSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "primaryOutput", GoGetter: "PrimaryOutput"},
			_jsii_.MemberMethod{JsiiMethod: "producedBy", GoMethod: "ProducedBy"},
			_jsii_.MemberProperty{JsiiProperty: "producer", GoGetter: "Producer"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FileSet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFileSetProducer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.FileSetLocation",
		reflect.TypeOf((*FileSetLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.FromStackArtifactOptions",
		reflect.TypeOf((*FromStackArtifactOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.GitHubSourceOptions",
		reflect.TypeOf((*GitHubSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"monocdk.pipelines.ICodePipelineActionFactory",
		reflect.TypeOf((*ICodePipelineActionFactory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produceAction", GoMethod: "ProduceAction"},
		},
		func() interface{} {
			return &jsiiProxy_ICodePipelineActionFactory{}
		},
	)
	_jsii_.RegisterInterface(
		"monocdk.pipelines.IFileSetProducer",
		reflect.TypeOf((*IFileSetProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "primaryOutput", GoGetter: "PrimaryOutput"},
		},
		func() interface{} {
			return &jsiiProxy_IFileSetProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"monocdk.pipelines.IStageHost",
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
		"monocdk.pipelines.ManualApprovalStep",
		reflect.TypeOf((*ManualApprovalStep)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependencyFileSet", GoMethod: "AddDependencyFileSet"},
			_jsii_.MemberMethod{JsiiMethod: "addStepDependency", GoMethod: "AddStepDependency"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberMethod{JsiiMethod: "configurePrimaryOutput", GoMethod: "ConfigurePrimaryOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyFileSets", GoGetter: "DependencyFileSets"},
			_jsii_.MemberMethod{JsiiMethod: "discoverReferencedOutputs", GoMethod: "DiscoverReferencedOutputs"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "isSource", GoGetter: "IsSource"},
			_jsii_.MemberProperty{JsiiProperty: "primaryOutput", GoGetter: "PrimaryOutput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ManualApprovalStep{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Step)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.ManualApprovalStepProps",
		reflect.TypeOf((*ManualApprovalStepProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.PermissionsBroadeningCheckProps",
		reflect.TypeOf((*PermissionsBroadeningCheckProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.PipelineBase",
		reflect.TypeOf((*PipelineBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStage", GoMethod: "AddStage"},
			_jsii_.MemberMethod{JsiiMethod: "addWave", GoMethod: "AddWave"},
			_jsii_.MemberMethod{JsiiMethod: "buildPipeline", GoMethod: "BuildPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "cloudAssemblyFileSet", GoGetter: "CloudAssemblyFileSet"},
			_jsii_.MemberMethod{JsiiMethod: "doBuildPipeline", GoMethod: "DoBuildPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "synth", GoGetter: "Synth"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberProperty{JsiiProperty: "waves", GoGetter: "Waves"},
		},
		func() interface{} {
			j := jsiiProxy_PipelineBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.PipelineBaseProps",
		reflect.TypeOf((*PipelineBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.ProduceActionOptions",
		reflect.TypeOf((*ProduceActionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.PublishAssetsAction",
		reflect.TypeOf((*PublishAssetsAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "addPublishCommand", GoMethod: "AddPublishCommand"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_PublishAssetsAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awscodepipelineIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.PublishAssetsActionProps",
		reflect.TypeOf((*PublishAssetsActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.S3SourceOptions",
		reflect.TypeOf((*S3SourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.ShellScriptAction",
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
		"monocdk.pipelines.ShellScriptActionProps",
		reflect.TypeOf((*ShellScriptActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.ShellStep",
		reflect.TypeOf((*ShellStep)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependencyFileSet", GoMethod: "AddDependencyFileSet"},
			_jsii_.MemberMethod{JsiiMethod: "addOutputDirectory", GoMethod: "AddOutputDirectory"},
			_jsii_.MemberMethod{JsiiMethod: "addStepDependency", GoMethod: "AddStepDependency"},
			_jsii_.MemberProperty{JsiiProperty: "commands", GoGetter: "Commands"},
			_jsii_.MemberMethod{JsiiMethod: "configurePrimaryOutput", GoMethod: "ConfigurePrimaryOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyFileSets", GoGetter: "DependencyFileSets"},
			_jsii_.MemberMethod{JsiiMethod: "discoverReferencedOutputs", GoMethod: "DiscoverReferencedOutputs"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "envFromCfnOutputs", GoGetter: "EnvFromCfnOutputs"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputs", GoGetter: "Inputs"},
			_jsii_.MemberProperty{JsiiProperty: "installCommands", GoGetter: "InstallCommands"},
			_jsii_.MemberProperty{JsiiProperty: "isSource", GoGetter: "IsSource"},
			_jsii_.MemberProperty{JsiiProperty: "outputs", GoGetter: "Outputs"},
			_jsii_.MemberProperty{JsiiProperty: "primaryOutput", GoGetter: "PrimaryOutput"},
			_jsii_.MemberMethod{JsiiMethod: "primaryOutputDirectory", GoMethod: "PrimaryOutputDirectory"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ShellStep{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Step)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.ShellStepProps",
		reflect.TypeOf((*ShellStepProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.SimpleSynthAction",
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
		"monocdk.pipelines.SimpleSynthActionProps",
		reflect.TypeOf((*SimpleSynthActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.SimpleSynthOptions",
		reflect.TypeOf((*SimpleSynthOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.StackAsset",
		reflect.TypeOf((*StackAsset)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.StackDeployment",
		reflect.TypeOf((*StackDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absoluteTemplatePath", GoGetter: "AbsoluteTemplatePath"},
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberMethod{JsiiMethod: "addStackDependency", GoMethod: "AddStackDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addStackSteps", GoMethod: "AddStackSteps"},
			_jsii_.MemberProperty{JsiiProperty: "assets", GoGetter: "Assets"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleArn", GoGetter: "AssumeRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "changeSet", GoGetter: "ChangeSet"},
			_jsii_.MemberProperty{JsiiProperty: "constructPath", GoGetter: "ConstructPath"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "post", GoGetter: "Post"},
			_jsii_.MemberProperty{JsiiProperty: "pre", GoGetter: "Pre"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "stackArtifactId", GoGetter: "StackArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "stackDependencies", GoGetter: "StackDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateAsset", GoGetter: "TemplateAsset"},
			_jsii_.MemberProperty{JsiiProperty: "templateUrl", GoGetter: "TemplateUrl"},
		},
		func() interface{} {
			return &jsiiProxy_StackDeployment{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.StackDeploymentProps",
		reflect.TypeOf((*StackDeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.StackOutput",
		reflect.TypeOf((*StackOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "artifactFile", GoGetter: "ArtifactFile"},
			_jsii_.MemberProperty{JsiiProperty: "outputName", GoGetter: "OutputName"},
		},
		func() interface{} {
			return &jsiiProxy_StackOutput{}
		},
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.StackOutputReference",
		reflect.TypeOf((*StackOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "isProducedBy", GoMethod: "IsProducedBy"},
			_jsii_.MemberProperty{JsiiProperty: "outputName", GoGetter: "OutputName"},
			_jsii_.MemberProperty{JsiiProperty: "stackDescription", GoGetter: "StackDescription"},
		},
		func() interface{} {
			return &jsiiProxy_StackOutputReference{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.StackSteps",
		reflect.TypeOf((*StackSteps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.StageDeployment",
		reflect.TypeOf((*StageDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPost", GoMethod: "AddPost"},
			_jsii_.MemberMethod{JsiiMethod: "addPre", GoMethod: "AddPre"},
			_jsii_.MemberProperty{JsiiProperty: "post", GoGetter: "Post"},
			_jsii_.MemberProperty{JsiiProperty: "pre", GoGetter: "Pre"},
			_jsii_.MemberProperty{JsiiProperty: "stacks", GoGetter: "Stacks"},
			_jsii_.MemberProperty{JsiiProperty: "stackSteps", GoGetter: "StackSteps"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
		},
		func() interface{} {
			return &jsiiProxy_StageDeployment{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.StageDeploymentProps",
		reflect.TypeOf((*StageDeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.StandardNpmSynthOptions",
		reflect.TypeOf((*StandardNpmSynthOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.StandardYarnSynthOptions",
		reflect.TypeOf((*StandardYarnSynthOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.Step",
		reflect.TypeOf((*Step)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependencyFileSet", GoMethod: "AddDependencyFileSet"},
			_jsii_.MemberMethod{JsiiMethod: "addStepDependency", GoMethod: "AddStepDependency"},
			_jsii_.MemberMethod{JsiiMethod: "configurePrimaryOutput", GoMethod: "ConfigurePrimaryOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyFileSets", GoGetter: "DependencyFileSets"},
			_jsii_.MemberMethod{JsiiMethod: "discoverReferencedOutputs", GoMethod: "DiscoverReferencedOutputs"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "isSource", GoGetter: "IsSource"},
			_jsii_.MemberProperty{JsiiProperty: "primaryOutput", GoGetter: "PrimaryOutput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Step{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFileSetProducer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.UpdatePipelineAction",
		reflect.TypeOf((*UpdatePipelineAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_UpdatePipelineAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awscodepipelineIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.UpdatePipelineActionProps",
		reflect.TypeOf((*UpdatePipelineActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.pipelines.Wave",
		reflect.TypeOf((*Wave)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPost", GoMethod: "AddPost"},
			_jsii_.MemberMethod{JsiiMethod: "addPre", GoMethod: "AddPre"},
			_jsii_.MemberMethod{JsiiMethod: "addStage", GoMethod: "AddStage"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "post", GoGetter: "Post"},
			_jsii_.MemberProperty{JsiiProperty: "pre", GoGetter: "Pre"},
			_jsii_.MemberProperty{JsiiProperty: "stages", GoGetter: "Stages"},
		},
		func() interface{} {
			return &jsiiProxy_Wave{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.WaveOptions",
		reflect.TypeOf((*WaveOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.pipelines.WaveProps",
		reflect.TypeOf((*WaveProps)(nil)).Elem(),
	)
}
