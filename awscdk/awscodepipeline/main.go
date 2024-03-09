package awscodepipeline

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codepipeline.Action",
		reflect.TypeOf((*Action)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberProperty{JsiiProperty: "providedActionProperties", GoGetter: "ProvidedActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_Action{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.ActionArtifactBounds",
		reflect.TypeOf((*ActionArtifactBounds)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.ActionBindOptions",
		reflect.TypeOf((*ActionBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_codepipeline.ActionCategory",
		reflect.TypeOf((*ActionCategory)(nil)).Elem(),
		map[string]interface{}{
			"SOURCE": ActionCategory_SOURCE,
			"BUILD": ActionCategory_BUILD,
			"TEST": ActionCategory_TEST,
			"APPROVAL": ActionCategory_APPROVAL,
			"DEPLOY": ActionCategory_DEPLOY,
			"INVOKE": ActionCategory_INVOKE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.ActionConfig",
		reflect.TypeOf((*ActionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.ActionProperties",
		reflect.TypeOf((*ActionProperties)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codepipeline.Artifact",
		reflect.TypeOf((*Artifact)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "artifactName", GoGetter: "ArtifactName"},
			_jsii_.MemberMethod{JsiiMethod: "atPath", GoMethod: "AtPath"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "getParam", GoMethod: "GetParam"},
			_jsii_.MemberProperty{JsiiProperty: "objectKey", GoGetter: "ObjectKey"},
			_jsii_.MemberProperty{JsiiProperty: "s3Location", GoGetter: "S3Location"},
			_jsii_.MemberMethod{JsiiMethod: "setMetadata", GoMethod: "SetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			return &jsiiProxy_Artifact{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codepipeline.ArtifactPath",
		reflect.TypeOf((*ArtifactPath)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "artifact", GoGetter: "Artifact"},
			_jsii_.MemberProperty{JsiiProperty: "fileName", GoGetter: "FileName"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
		},
		func() interface{} {
			return &jsiiProxy_ArtifactPath{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codepipeline.CfnCustomActionType",
		reflect.TypeOf((*CfnCustomActionType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "category", GoGetter: "Category"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProperties", GoGetter: "ConfigurationProperties"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "inputArtifactDetails", GoGetter: "InputArtifactDetails"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "outputArtifactDetails", GoGetter: "OutputArtifactDetails"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomActionType{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnCustomActionType.ArtifactDetailsProperty",
		reflect.TypeOf((*CfnCustomActionType_ArtifactDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnCustomActionType.ConfigurationPropertiesProperty",
		reflect.TypeOf((*CfnCustomActionType_ConfigurationPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnCustomActionType.SettingsProperty",
		reflect.TypeOf((*CfnCustomActionType_SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnCustomActionTypeProps",
		reflect.TypeOf((*CfnCustomActionTypeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline",
		reflect.TypeOf((*CfnPipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "artifactStore", GoGetter: "ArtifactStore"},
			_jsii_.MemberProperty{JsiiProperty: "artifactStores", GoGetter: "ArtifactStores"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersion", GoGetter: "AttrVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disableInboundStageTransitions", GoGetter: "DisableInboundStageTransitions"},
			_jsii_.MemberProperty{JsiiProperty: "executionMode", GoGetter: "ExecutionMode"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineType", GoGetter: "PipelineType"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "restartExecutionOnUpdate", GoGetter: "RestartExecutionOnUpdate"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stages", GoGetter: "Stages"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggers", GoGetter: "Triggers"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPipeline{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.ActionDeclarationProperty",
		reflect.TypeOf((*CfnPipeline_ActionDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.ActionTypeIdProperty",
		reflect.TypeOf((*CfnPipeline_ActionTypeIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.ArtifactStoreMapProperty",
		reflect.TypeOf((*CfnPipeline_ArtifactStoreMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.ArtifactStoreProperty",
		reflect.TypeOf((*CfnPipeline_ArtifactStoreProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.BlockerDeclarationProperty",
		reflect.TypeOf((*CfnPipeline_BlockerDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.EncryptionKeyProperty",
		reflect.TypeOf((*CfnPipeline_EncryptionKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.GitBranchFilterCriteriaProperty",
		reflect.TypeOf((*CfnPipeline_GitBranchFilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.GitConfigurationProperty",
		reflect.TypeOf((*CfnPipeline_GitConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.GitFilePathFilterCriteriaProperty",
		reflect.TypeOf((*CfnPipeline_GitFilePathFilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.GitPullRequestFilterProperty",
		reflect.TypeOf((*CfnPipeline_GitPullRequestFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.GitPushFilterProperty",
		reflect.TypeOf((*CfnPipeline_GitPushFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.GitTagFilterCriteriaProperty",
		reflect.TypeOf((*CfnPipeline_GitTagFilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.InputArtifactProperty",
		reflect.TypeOf((*CfnPipeline_InputArtifactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.OutputArtifactProperty",
		reflect.TypeOf((*CfnPipeline_OutputArtifactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.PipelineTriggerDeclarationProperty",
		reflect.TypeOf((*CfnPipeline_PipelineTriggerDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.StageDeclarationProperty",
		reflect.TypeOf((*CfnPipeline_StageDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.StageTransitionProperty",
		reflect.TypeOf((*CfnPipeline_StageTransitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipeline.VariableDeclarationProperty",
		reflect.TypeOf((*CfnPipeline_VariableDeclarationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnPipelineProps",
		reflect.TypeOf((*CfnPipelineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codepipeline.CfnWebhook",
		reflect.TypeOf((*CfnWebhook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrUrl", GoGetter: "AttrUrl"},
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationConfiguration", GoGetter: "AuthenticationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "filters", GoGetter: "Filters"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "registerWithThirdParty", GoGetter: "RegisterWithThirdParty"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "targetAction", GoGetter: "TargetAction"},
			_jsii_.MemberProperty{JsiiProperty: "targetPipeline", GoGetter: "TargetPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "targetPipelineVersion", GoGetter: "TargetPipelineVersion"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWebhook{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnWebhook.WebhookAuthConfigurationProperty",
		reflect.TypeOf((*CfnWebhook_WebhookAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnWebhook.WebhookFilterRuleProperty",
		reflect.TypeOf((*CfnWebhook_WebhookFilterRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CfnWebhookProps",
		reflect.TypeOf((*CfnWebhookProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CommonActionProps",
		reflect.TypeOf((*CommonActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CommonAwsActionProps",
		reflect.TypeOf((*CommonAwsActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CrossRegionSupport",
		reflect.TypeOf((*CrossRegionSupport)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CustomActionProperty",
		reflect.TypeOf((*CustomActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codepipeline.CustomActionRegistration",
		reflect.TypeOf((*CustomActionRegistration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CustomActionRegistration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.CustomActionRegistrationProps",
		reflect.TypeOf((*CustomActionRegistrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_codepipeline.ExecutionMode",
		reflect.TypeOf((*ExecutionMode)(nil)).Elem(),
		map[string]interface{}{
			"QUEUED": ExecutionMode_QUEUED,
			"SUPERSEDED": ExecutionMode_SUPERSEDED,
			"PARALLEL": ExecutionMode_PARALLEL,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.GitConfiguration",
		reflect.TypeOf((*GitConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.GitPushFilter",
		reflect.TypeOf((*GitPushFilter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codepipeline.GlobalVariables",
		reflect.TypeOf((*GlobalVariables)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GlobalVariables{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codepipeline.IAction",
		reflect.TypeOf((*IAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
		},
		func() interface{} {
			return &jsiiProxy_IAction{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codepipeline.IPipeline",
		reflect.TypeOf((*IPipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bindAsNotificationRuleSource", GoMethod: "BindAsNotificationRuleSource"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOn", GoMethod: "NotifyOn"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnAnyActionStateChange", GoMethod: "NotifyOnAnyActionStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnAnyManualApprovalStateChange", GoMethod: "NotifyOnAnyManualApprovalStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnAnyStageStateChange", GoMethod: "NotifyOnAnyStageStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnExecutionStateChange", GoMethod: "NotifyOnExecutionStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineArn", GoGetter: "PipelineArn"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineName", GoGetter: "PipelineName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IPipeline{}
			_jsii_.InitJsiiProxy(&j.Type__awscodestarnotificationsINotificationRuleSource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codepipeline.IStage",
		reflect.TypeOf((*IStage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actions", GoGetter: "Actions"},
			_jsii_.MemberMethod{JsiiMethod: "addAction", GoMethod: "AddAction"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberProperty{JsiiProperty: "pipeline", GoGetter: "Pipeline"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
		},
		func() interface{} {
			return &jsiiProxy_IStage{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codepipeline.Pipeline",
		reflect.TypeOf((*Pipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addStage", GoMethod: "AddStage"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addTrigger", GoMethod: "AddTrigger"},
			_jsii_.MemberMethod{JsiiMethod: "addVariable", GoMethod: "AddVariable"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "artifactBucket", GoGetter: "ArtifactBucket"},
			_jsii_.MemberMethod{JsiiMethod: "bindAsNotificationRuleSource", GoMethod: "BindAsNotificationRuleSource"},
			_jsii_.MemberProperty{JsiiProperty: "crossRegionSupport", GoGetter: "CrossRegionSupport"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOn", GoMethod: "NotifyOn"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnAnyActionStateChange", GoMethod: "NotifyOnAnyActionStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnAnyManualApprovalStateChange", GoMethod: "NotifyOnAnyManualApprovalStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnAnyStageStateChange", GoMethod: "NotifyOnAnyStageStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnExecutionStateChange", GoMethod: "NotifyOnExecutionStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineArn", GoGetter: "PipelineArn"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineName", GoGetter: "PipelineName"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineVersion", GoGetter: "PipelineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "stage", GoMethod: "Stage"},
			_jsii_.MemberProperty{JsiiProperty: "stageCount", GoGetter: "StageCount"},
			_jsii_.MemberProperty{JsiiProperty: "stages", GoGetter: "Stages"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Pipeline{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPipeline)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_codepipeline.PipelineNotificationEvents",
		reflect.TypeOf((*PipelineNotificationEvents)(nil)).Elem(),
		map[string]interface{}{
			"PIPELINE_EXECUTION_FAILED": PipelineNotificationEvents_PIPELINE_EXECUTION_FAILED,
			"PIPELINE_EXECUTION_CANCELED": PipelineNotificationEvents_PIPELINE_EXECUTION_CANCELED,
			"PIPELINE_EXECUTION_STARTED": PipelineNotificationEvents_PIPELINE_EXECUTION_STARTED,
			"PIPELINE_EXECUTION_RESUMED": PipelineNotificationEvents_PIPELINE_EXECUTION_RESUMED,
			"PIPELINE_EXECUTION_SUCCEEDED": PipelineNotificationEvents_PIPELINE_EXECUTION_SUCCEEDED,
			"PIPELINE_EXECUTION_SUPERSEDED": PipelineNotificationEvents_PIPELINE_EXECUTION_SUPERSEDED,
			"STAGE_EXECUTION_STARTED": PipelineNotificationEvents_STAGE_EXECUTION_STARTED,
			"STAGE_EXECUTION_SUCCEEDED": PipelineNotificationEvents_STAGE_EXECUTION_SUCCEEDED,
			"STAGE_EXECUTION_RESUMED": PipelineNotificationEvents_STAGE_EXECUTION_RESUMED,
			"STAGE_EXECUTION_CANCELED": PipelineNotificationEvents_STAGE_EXECUTION_CANCELED,
			"STAGE_EXECUTION_FAILED": PipelineNotificationEvents_STAGE_EXECUTION_FAILED,
			"ACTION_EXECUTION_SUCCEEDED": PipelineNotificationEvents_ACTION_EXECUTION_SUCCEEDED,
			"ACTION_EXECUTION_FAILED": PipelineNotificationEvents_ACTION_EXECUTION_FAILED,
			"ACTION_EXECUTION_CANCELED": PipelineNotificationEvents_ACTION_EXECUTION_CANCELED,
			"ACTION_EXECUTION_STARTED": PipelineNotificationEvents_ACTION_EXECUTION_STARTED,
			"MANUAL_APPROVAL_FAILED": PipelineNotificationEvents_MANUAL_APPROVAL_FAILED,
			"MANUAL_APPROVAL_NEEDED": PipelineNotificationEvents_MANUAL_APPROVAL_NEEDED,
			"MANUAL_APPROVAL_SUCCEEDED": PipelineNotificationEvents_MANUAL_APPROVAL_SUCCEEDED,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.PipelineNotifyOnOptions",
		reflect.TypeOf((*PipelineNotifyOnOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.PipelineProps",
		reflect.TypeOf((*PipelineProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_codepipeline.PipelineType",
		reflect.TypeOf((*PipelineType)(nil)).Elem(),
		map[string]interface{}{
			"V1": PipelineType_V1,
			"V2": PipelineType_V2,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_codepipeline.ProviderType",
		reflect.TypeOf((*ProviderType)(nil)).Elem(),
		map[string]interface{}{
			"CODE_STAR_SOURCE_CONNECTION": ProviderType_CODE_STAR_SOURCE_CONNECTION,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.StageOptions",
		reflect.TypeOf((*StageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.StagePlacement",
		reflect.TypeOf((*StagePlacement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.StageProps",
		reflect.TypeOf((*StageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codepipeline.Trigger",
		reflect.TypeOf((*Trigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "sourceAction", GoGetter: "SourceAction"},
		},
		func() interface{} {
			return &jsiiProxy_Trigger{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.TriggerProps",
		reflect.TypeOf((*TriggerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codepipeline.Variable",
		reflect.TypeOf((*Variable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "reference", GoMethod: "Reference"},
			_jsii_.MemberProperty{JsiiProperty: "variableName", GoGetter: "VariableName"},
		},
		func() interface{} {
			return &jsiiProxy_Variable{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codepipeline.VariableProps",
		reflect.TypeOf((*VariableProps)(nil)).Elem(),
	)
}
