package awsbedrock

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnAgent",
		reflect.TypeOf((*CfnAgent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionGroups", GoGetter: "ActionGroups"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "agentCollaboration", GoGetter: "AgentCollaboration"},
			_jsii_.MemberProperty{JsiiProperty: "agentCollaborators", GoGetter: "AgentCollaborators"},
			_jsii_.MemberProperty{JsiiProperty: "agentName", GoGetter: "AgentName"},
			_jsii_.MemberProperty{JsiiProperty: "agentResourceRoleArn", GoGetter: "AgentResourceRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentArn", GoGetter: "AttrAgentArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentId", GoGetter: "AttrAgentId"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentStatus", GoGetter: "AttrAgentStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentVersion", GoGetter: "AttrAgentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrFailureReasons", GoGetter: "AttrFailureReasons"},
			_jsii_.MemberProperty{JsiiProperty: "attrPreparedAt", GoGetter: "AttrPreparedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrRecommendedActions", GoGetter: "AttrRecommendedActions"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "autoPrepare", GoGetter: "AutoPrepare"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customerEncryptionKeyArn", GoGetter: "CustomerEncryptionKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "customOrchestration", GoGetter: "CustomOrchestration"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "foundationModel", GoGetter: "FoundationModel"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailConfiguration", GoGetter: "GuardrailConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "idleSessionTtlInSeconds", GoGetter: "IdleSessionTtlInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBases", GoGetter: "KnowledgeBases"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "memoryConfiguration", GoGetter: "MemoryConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "orchestrationType", GoGetter: "OrchestrationType"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "promptOverrideConfiguration", GoGetter: "PromptOverrideConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "skipResourceInUseCheckOnDelete", GoGetter: "SkipResourceInUseCheckOnDelete"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "testAliasTags", GoGetter: "TestAliasTags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAgent{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.APISchemaProperty",
		reflect.TypeOf((*CfnAgent_APISchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.ActionGroupExecutorProperty",
		reflect.TypeOf((*CfnAgent_ActionGroupExecutorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.AgentActionGroupProperty",
		reflect.TypeOf((*CfnAgent_AgentActionGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.AgentCollaboratorProperty",
		reflect.TypeOf((*CfnAgent_AgentCollaboratorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.AgentDescriptorProperty",
		reflect.TypeOf((*CfnAgent_AgentDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.AgentKnowledgeBaseProperty",
		reflect.TypeOf((*CfnAgent_AgentKnowledgeBaseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.CustomOrchestrationProperty",
		reflect.TypeOf((*CfnAgent_CustomOrchestrationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.FunctionProperty",
		reflect.TypeOf((*CfnAgent_FunctionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.FunctionSchemaProperty",
		reflect.TypeOf((*CfnAgent_FunctionSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.GuardrailConfigurationProperty",
		reflect.TypeOf((*CfnAgent_GuardrailConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.InferenceConfigurationProperty",
		reflect.TypeOf((*CfnAgent_InferenceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.MemoryConfigurationProperty",
		reflect.TypeOf((*CfnAgent_MemoryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.OrchestrationExecutorProperty",
		reflect.TypeOf((*CfnAgent_OrchestrationExecutorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.ParameterDetailProperty",
		reflect.TypeOf((*CfnAgent_ParameterDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.PromptConfigurationProperty",
		reflect.TypeOf((*CfnAgent_PromptConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.PromptOverrideConfigurationProperty",
		reflect.TypeOf((*CfnAgent_PromptOverrideConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.S3IdentifierProperty",
		reflect.TypeOf((*CfnAgent_S3IdentifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgent.SessionSummaryConfigurationProperty",
		reflect.TypeOf((*CfnAgent_SessionSummaryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnAgentAlias",
		reflect.TypeOf((*CfnAgentAlias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "agentAliasName", GoGetter: "AgentAliasName"},
			_jsii_.MemberProperty{JsiiProperty: "agentId", GoGetter: "AgentId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentAliasArn", GoGetter: "AttrAgentAliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentAliasHistoryEvents", GoGetter: "AttrAgentAliasHistoryEvents"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentAliasId", GoGetter: "AttrAgentAliasId"},
			_jsii_.MemberProperty{JsiiProperty: "attrAgentAliasStatus", GoGetter: "AttrAgentAliasStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "routingConfiguration", GoGetter: "RoutingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAgentAlias{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgentAlias.AgentAliasHistoryEventProperty",
		reflect.TypeOf((*CfnAgentAlias_AgentAliasHistoryEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgentAlias.AgentAliasRoutingConfigurationListItemProperty",
		reflect.TypeOf((*CfnAgentAlias_AgentAliasRoutingConfigurationListItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgentAliasProps",
		reflect.TypeOf((*CfnAgentAliasProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnAgentProps",
		reflect.TypeOf((*CfnAgentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnApplicationInferenceProfile",
		reflect.TypeOf((*CfnApplicationInferenceProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrInferenceProfileArn", GoGetter: "AttrInferenceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrInferenceProfileId", GoGetter: "AttrInferenceProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "attrInferenceProfileIdentifier", GoGetter: "AttrInferenceProfileIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "attrModels", GoGetter: "AttrModels"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrType", GoGetter: "AttrType"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileName", GoGetter: "InferenceProfileName"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "modelSource", GoGetter: "ModelSource"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationInferenceProfile{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnApplicationInferenceProfile.InferenceProfileModelProperty",
		reflect.TypeOf((*CfnApplicationInferenceProfile_InferenceProfileModelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnApplicationInferenceProfile.InferenceProfileModelSourceProperty",
		reflect.TypeOf((*CfnApplicationInferenceProfile_InferenceProfileModelSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnApplicationInferenceProfileProps",
		reflect.TypeOf((*CfnApplicationInferenceProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnBlueprint",
		reflect.TypeOf((*CfnBlueprint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrBlueprintArn", GoGetter: "AttrBlueprintArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrBlueprintStage", GoGetter: "AttrBlueprintStage"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTime", GoGetter: "AttrCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastModifiedTime", GoGetter: "AttrLastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "blueprintName", GoGetter: "BlueprintName"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionContext", GoGetter: "KmsEncryptionContext"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBlueprint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnBlueprintProps",
		reflect.TypeOf((*CfnBlueprintProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject",
		reflect.TypeOf((*CfnDataAutomationProject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTime", GoGetter: "AttrCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastModifiedTime", GoGetter: "AttrLastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrProjectArn", GoGetter: "AttrProjectArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrProjectStage", GoGetter: "AttrProjectStage"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customOutputConfiguration", GoGetter: "CustomOutputConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionContext", GoGetter: "KmsEncryptionContext"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "overrideConfiguration", GoGetter: "OverrideConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "projectDescription", GoGetter: "ProjectDescription"},
			_jsii_.MemberProperty{JsiiProperty: "projectName", GoGetter: "ProjectName"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "standardOutputConfiguration", GoGetter: "StandardOutputConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataAutomationProject{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.AudioExtractionCategoryProperty",
		reflect.TypeOf((*CfnDataAutomationProject_AudioExtractionCategoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.AudioOverrideConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_AudioOverrideConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.AudioStandardExtractionProperty",
		reflect.TypeOf((*CfnDataAutomationProject_AudioStandardExtractionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.AudioStandardGenerativeFieldProperty",
		reflect.TypeOf((*CfnDataAutomationProject_AudioStandardGenerativeFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.AudioStandardOutputConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_AudioStandardOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.BlueprintItemProperty",
		reflect.TypeOf((*CfnDataAutomationProject_BlueprintItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.CustomOutputConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_CustomOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.DocumentBoundingBoxProperty",
		reflect.TypeOf((*CfnDataAutomationProject_DocumentBoundingBoxProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.DocumentExtractionGranularityProperty",
		reflect.TypeOf((*CfnDataAutomationProject_DocumentExtractionGranularityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.DocumentOutputAdditionalFileFormatProperty",
		reflect.TypeOf((*CfnDataAutomationProject_DocumentOutputAdditionalFileFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.DocumentOutputFormatProperty",
		reflect.TypeOf((*CfnDataAutomationProject_DocumentOutputFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.DocumentOutputTextFormatProperty",
		reflect.TypeOf((*CfnDataAutomationProject_DocumentOutputTextFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.DocumentOverrideConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_DocumentOverrideConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.DocumentStandardExtractionProperty",
		reflect.TypeOf((*CfnDataAutomationProject_DocumentStandardExtractionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.DocumentStandardGenerativeFieldProperty",
		reflect.TypeOf((*CfnDataAutomationProject_DocumentStandardGenerativeFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.DocumentStandardOutputConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_DocumentStandardOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.ImageBoundingBoxProperty",
		reflect.TypeOf((*CfnDataAutomationProject_ImageBoundingBoxProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.ImageExtractionCategoryProperty",
		reflect.TypeOf((*CfnDataAutomationProject_ImageExtractionCategoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.ImageOverrideConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_ImageOverrideConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.ImageStandardExtractionProperty",
		reflect.TypeOf((*CfnDataAutomationProject_ImageStandardExtractionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.ImageStandardGenerativeFieldProperty",
		reflect.TypeOf((*CfnDataAutomationProject_ImageStandardGenerativeFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.ImageStandardOutputConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_ImageStandardOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.ModalityProcessingConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_ModalityProcessingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.ModalityRoutingConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_ModalityRoutingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.OverrideConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_OverrideConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.SplitterConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_SplitterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.StandardOutputConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_StandardOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.VideoBoundingBoxProperty",
		reflect.TypeOf((*CfnDataAutomationProject_VideoBoundingBoxProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.VideoExtractionCategoryProperty",
		reflect.TypeOf((*CfnDataAutomationProject_VideoExtractionCategoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.VideoOverrideConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_VideoOverrideConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.VideoStandardExtractionProperty",
		reflect.TypeOf((*CfnDataAutomationProject_VideoStandardExtractionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.VideoStandardGenerativeFieldProperty",
		reflect.TypeOf((*CfnDataAutomationProject_VideoStandardGenerativeFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProject.VideoStandardOutputConfigurationProperty",
		reflect.TypeOf((*CfnDataAutomationProject_VideoStandardOutputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataAutomationProjectProps",
		reflect.TypeOf((*CfnDataAutomationProjectProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnDataSource",
		reflect.TypeOf((*CfnDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrDataSourceConfigurationWebConfigurationCrawlerConfigurationUserAgentHeader", GoGetter: "AttrDataSourceConfigurationWebConfigurationCrawlerConfigurationUserAgentHeader"},
			_jsii_.MemberProperty{JsiiProperty: "attrDataSourceId", GoGetter: "AttrDataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "attrDataSourceStatus", GoGetter: "AttrDataSourceStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrFailureReasons", GoGetter: "AttrFailureReasons"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataDeletionPolicy", GoGetter: "DataDeletionPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceConfiguration", GoGetter: "DataSourceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideEncryptionConfiguration", GoGetter: "ServerSideEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vectorIngestionConfiguration", GoGetter: "VectorIngestionConfiguration"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.BedrockDataAutomationConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_BedrockDataAutomationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.BedrockFoundationModelConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_BedrockFoundationModelConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.BedrockFoundationModelContextEnrichmentConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_BedrockFoundationModelContextEnrichmentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.ChunkingConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ChunkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.ConfluenceCrawlerConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ConfluenceCrawlerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.ConfluenceDataSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ConfluenceDataSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.ConfluenceSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ConfluenceSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.ContextEnrichmentConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ContextEnrichmentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.CrawlFilterConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_CrawlFilterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.CustomTransformationConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_CustomTransformationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.DataSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_DataSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.EnrichmentStrategyConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_EnrichmentStrategyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.FixedSizeChunkingConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_FixedSizeChunkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.HierarchicalChunkingConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_HierarchicalChunkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.HierarchicalChunkingLevelConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_HierarchicalChunkingLevelConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.IntermediateStorageProperty",
		reflect.TypeOf((*CfnDataSource_IntermediateStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.ParsingConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ParsingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.ParsingPromptProperty",
		reflect.TypeOf((*CfnDataSource_ParsingPromptProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.PatternObjectFilterConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_PatternObjectFilterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.PatternObjectFilterProperty",
		reflect.TypeOf((*CfnDataSource_PatternObjectFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.S3DataSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_S3DataSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.S3LocationProperty",
		reflect.TypeOf((*CfnDataSource_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.SalesforceCrawlerConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SalesforceCrawlerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.SalesforceDataSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SalesforceDataSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.SalesforceSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SalesforceSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.SeedUrlProperty",
		reflect.TypeOf((*CfnDataSource_SeedUrlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.SemanticChunkingConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SemanticChunkingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.ServerSideEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ServerSideEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.SharePointCrawlerConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SharePointCrawlerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.SharePointDataSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SharePointDataSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.SharePointSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SharePointSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.TransformationFunctionProperty",
		reflect.TypeOf((*CfnDataSource_TransformationFunctionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.TransformationLambdaConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_TransformationLambdaConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.TransformationProperty",
		reflect.TypeOf((*CfnDataSource_TransformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.UrlConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_UrlConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.VectorIngestionConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_VectorIngestionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.WebCrawlerConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_WebCrawlerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.WebCrawlerLimitsProperty",
		reflect.TypeOf((*CfnDataSource_WebCrawlerLimitsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.WebDataSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_WebDataSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSource.WebSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_WebSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnDataSourceProps",
		reflect.TypeOf((*CfnDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnFlow",
		reflect.TypeOf((*CfnFlow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrValidations", GoGetter: "AttrValidations"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersion", GoGetter: "AttrVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customerEncryptionKeyArn", GoGetter: "CustomerEncryptionKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "definitionS3Location", GoGetter: "DefinitionS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "definitionString", GoGetter: "DefinitionString"},
			_jsii_.MemberProperty{JsiiProperty: "definitionSubstitutions", GoGetter: "DefinitionSubstitutions"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
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
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "testAliasTags", GoGetter: "TestAliasTags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlow{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.AgentFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlow_AgentFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.ConditionFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlow_ConditionFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.FlowConditionProperty",
		reflect.TypeOf((*CfnFlow_FlowConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.FlowConditionalConnectionConfigurationProperty",
		reflect.TypeOf((*CfnFlow_FlowConditionalConnectionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.FlowConnectionConfigurationProperty",
		reflect.TypeOf((*CfnFlow_FlowConnectionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.FlowConnectionProperty",
		reflect.TypeOf((*CfnFlow_FlowConnectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.FlowDataConnectionConfigurationProperty",
		reflect.TypeOf((*CfnFlow_FlowDataConnectionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.FlowDefinitionProperty",
		reflect.TypeOf((*CfnFlow_FlowDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.FlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlow_FlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.FlowNodeInputProperty",
		reflect.TypeOf((*CfnFlow_FlowNodeInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.FlowNodeOutputProperty",
		reflect.TypeOf((*CfnFlow_FlowNodeOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.FlowNodeProperty",
		reflect.TypeOf((*CfnFlow_FlowNodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.FlowValidationProperty",
		reflect.TypeOf((*CfnFlow_FlowValidationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.GuardrailConfigurationProperty",
		reflect.TypeOf((*CfnFlow_GuardrailConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.KnowledgeBaseFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlow_KnowledgeBaseFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.LambdaFunctionFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlow_LambdaFunctionFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.LexFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlow_LexFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.PromptFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlow_PromptFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.PromptFlowNodeInlineConfigurationProperty",
		reflect.TypeOf((*CfnFlow_PromptFlowNodeInlineConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.PromptFlowNodeResourceConfigurationProperty",
		reflect.TypeOf((*CfnFlow_PromptFlowNodeResourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.PromptFlowNodeSourceConfigurationProperty",
		reflect.TypeOf((*CfnFlow_PromptFlowNodeSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.PromptInferenceConfigurationProperty",
		reflect.TypeOf((*CfnFlow_PromptInferenceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.PromptInputVariableProperty",
		reflect.TypeOf((*CfnFlow_PromptInputVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.PromptModelInferenceConfigurationProperty",
		reflect.TypeOf((*CfnFlow_PromptModelInferenceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.PromptTemplateConfigurationProperty",
		reflect.TypeOf((*CfnFlow_PromptTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.RetrievalFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlow_RetrievalFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.RetrievalFlowNodeS3ConfigurationProperty",
		reflect.TypeOf((*CfnFlow_RetrievalFlowNodeS3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.RetrievalFlowNodeServiceConfigurationProperty",
		reflect.TypeOf((*CfnFlow_RetrievalFlowNodeServiceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.S3LocationProperty",
		reflect.TypeOf((*CfnFlow_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.StorageFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlow_StorageFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.StorageFlowNodeS3ConfigurationProperty",
		reflect.TypeOf((*CfnFlow_StorageFlowNodeS3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.StorageFlowNodeServiceConfigurationProperty",
		reflect.TypeOf((*CfnFlow_StorageFlowNodeServiceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlow.TextPromptTemplateConfigurationProperty",
		reflect.TypeOf((*CfnFlow_TextPromptTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnFlowAlias",
		reflect.TypeOf((*CfnFlowAlias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrFlowId", GoGetter: "AttrFlowId"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "flowArn", GoGetter: "FlowArn"},
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
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "routingConfiguration", GoGetter: "RoutingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlowAlias{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowAlias.FlowAliasRoutingConfigurationListItemProperty",
		reflect.TypeOf((*CfnFlowAlias_FlowAliasRoutingConfigurationListItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowAliasProps",
		reflect.TypeOf((*CfnFlowAliasProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowProps",
		reflect.TypeOf((*CfnFlowProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion",
		reflect.TypeOf((*CfnFlowVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrCustomerEncryptionKeyArn", GoGetter: "AttrCustomerEncryptionKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrDefinition", GoGetter: "AttrDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "attrExecutionRoleArn", GoGetter: "AttrExecutionRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrFlowId", GoGetter: "AttrFlowId"},
			_jsii_.MemberProperty{JsiiProperty: "attrName", GoGetter: "AttrName"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersion", GoGetter: "AttrVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "flowArn", GoGetter: "FlowArn"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlowVersion{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.AgentFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_AgentFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.ConditionFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_ConditionFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.FlowConditionProperty",
		reflect.TypeOf((*CfnFlowVersion_FlowConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.FlowConditionalConnectionConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_FlowConditionalConnectionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.FlowConnectionConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_FlowConnectionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.FlowConnectionProperty",
		reflect.TypeOf((*CfnFlowVersion_FlowConnectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.FlowDataConnectionConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_FlowDataConnectionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.FlowDefinitionProperty",
		reflect.TypeOf((*CfnFlowVersion_FlowDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.FlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_FlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.FlowNodeInputProperty",
		reflect.TypeOf((*CfnFlowVersion_FlowNodeInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.FlowNodeOutputProperty",
		reflect.TypeOf((*CfnFlowVersion_FlowNodeOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.FlowNodeProperty",
		reflect.TypeOf((*CfnFlowVersion_FlowNodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.GuardrailConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_GuardrailConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.KnowledgeBaseFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_KnowledgeBaseFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.LambdaFunctionFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_LambdaFunctionFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.LexFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_LexFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.PromptFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_PromptFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.PromptFlowNodeInlineConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_PromptFlowNodeInlineConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.PromptFlowNodeResourceConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_PromptFlowNodeResourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.PromptFlowNodeSourceConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_PromptFlowNodeSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.PromptInferenceConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_PromptInferenceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.PromptInputVariableProperty",
		reflect.TypeOf((*CfnFlowVersion_PromptInputVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.PromptModelInferenceConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_PromptModelInferenceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.PromptTemplateConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_PromptTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.RetrievalFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_RetrievalFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.RetrievalFlowNodeS3ConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_RetrievalFlowNodeS3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.RetrievalFlowNodeServiceConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_RetrievalFlowNodeServiceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.StorageFlowNodeConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_StorageFlowNodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.StorageFlowNodeS3ConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_StorageFlowNodeS3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.StorageFlowNodeServiceConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_StorageFlowNodeServiceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersion.TextPromptTemplateConfigurationProperty",
		reflect.TypeOf((*CfnFlowVersion_TextPromptTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnFlowVersionProps",
		reflect.TypeOf((*CfnFlowVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail",
		reflect.TypeOf((*CfnGuardrail)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrFailureRecommendations", GoGetter: "AttrFailureRecommendations"},
			_jsii_.MemberProperty{JsiiProperty: "attrGuardrailArn", GoGetter: "AttrGuardrailArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrGuardrailId", GoGetter: "AttrGuardrailId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatusReasons", GoGetter: "AttrStatusReasons"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersion", GoGetter: "AttrVersion"},
			_jsii_.MemberProperty{JsiiProperty: "blockedInputMessaging", GoGetter: "BlockedInputMessaging"},
			_jsii_.MemberProperty{JsiiProperty: "blockedOutputsMessaging", GoGetter: "BlockedOutputsMessaging"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "contentPolicyConfig", GoGetter: "ContentPolicyConfig"},
			_jsii_.MemberProperty{JsiiProperty: "contextualGroundingPolicyConfig", GoGetter: "ContextualGroundingPolicyConfig"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyArn", GoGetter: "KmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "sensitiveInformationPolicyConfig", GoGetter: "SensitiveInformationPolicyConfig"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "topicPolicyConfig", GoGetter: "TopicPolicyConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "wordPolicyConfig", GoGetter: "WordPolicyConfig"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGuardrail{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail.ContentFilterConfigProperty",
		reflect.TypeOf((*CfnGuardrail_ContentFilterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail.ContentPolicyConfigProperty",
		reflect.TypeOf((*CfnGuardrail_ContentPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail.ContextualGroundingFilterConfigProperty",
		reflect.TypeOf((*CfnGuardrail_ContextualGroundingFilterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail.ContextualGroundingPolicyConfigProperty",
		reflect.TypeOf((*CfnGuardrail_ContextualGroundingPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail.ManagedWordsConfigProperty",
		reflect.TypeOf((*CfnGuardrail_ManagedWordsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail.PiiEntityConfigProperty",
		reflect.TypeOf((*CfnGuardrail_PiiEntityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail.RegexConfigProperty",
		reflect.TypeOf((*CfnGuardrail_RegexConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail.SensitiveInformationPolicyConfigProperty",
		reflect.TypeOf((*CfnGuardrail_SensitiveInformationPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail.TopicConfigProperty",
		reflect.TypeOf((*CfnGuardrail_TopicConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail.TopicPolicyConfigProperty",
		reflect.TypeOf((*CfnGuardrail_TopicPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail.WordConfigProperty",
		reflect.TypeOf((*CfnGuardrail_WordConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail.WordPolicyConfigProperty",
		reflect.TypeOf((*CfnGuardrail_WordPolicyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrailProps",
		reflect.TypeOf((*CfnGuardrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnGuardrailVersion",
		reflect.TypeOf((*CfnGuardrailVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrGuardrailArn", GoGetter: "AttrGuardrailArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrGuardrailId", GoGetter: "AttrGuardrailId"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersion", GoGetter: "AttrVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailIdentifier", GoGetter: "GuardrailIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGuardrailVersion{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnGuardrailVersionProps",
		reflect.TypeOf((*CfnGuardrailVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase",
		reflect.TypeOf((*CfnKnowledgeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrFailureReasons", GoGetter: "AttrFailureReasons"},
			_jsii_.MemberProperty{JsiiProperty: "attrKnowledgeBaseArn", GoGetter: "AttrKnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrKnowledgeBaseId", GoGetter: "AttrKnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseConfiguration", GoGetter: "KnowledgeBaseConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "storageConfiguration", GoGetter: "StorageConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKnowledgeBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.BedrockEmbeddingModelConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_BedrockEmbeddingModelConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.CuratedQueryProperty",
		reflect.TypeOf((*CfnKnowledgeBase_CuratedQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.EmbeddingModelConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_EmbeddingModelConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.KendraKnowledgeBaseConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_KendraKnowledgeBaseConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.KnowledgeBaseConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_KnowledgeBaseConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.MongoDbAtlasConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_MongoDbAtlasConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.MongoDbAtlasFieldMappingProperty",
		reflect.TypeOf((*CfnKnowledgeBase_MongoDbAtlasFieldMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.NeptuneAnalyticsConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_NeptuneAnalyticsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.NeptuneAnalyticsFieldMappingProperty",
		reflect.TypeOf((*CfnKnowledgeBase_NeptuneAnalyticsFieldMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.OpenSearchManagedClusterConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_OpenSearchManagedClusterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.OpenSearchManagedClusterFieldMappingProperty",
		reflect.TypeOf((*CfnKnowledgeBase_OpenSearchManagedClusterFieldMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.OpenSearchServerlessConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_OpenSearchServerlessConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.OpenSearchServerlessFieldMappingProperty",
		reflect.TypeOf((*CfnKnowledgeBase_OpenSearchServerlessFieldMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.PineconeConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_PineconeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.PineconeFieldMappingProperty",
		reflect.TypeOf((*CfnKnowledgeBase_PineconeFieldMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.QueryGenerationColumnProperty",
		reflect.TypeOf((*CfnKnowledgeBase_QueryGenerationColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.QueryGenerationConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_QueryGenerationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.QueryGenerationContextProperty",
		reflect.TypeOf((*CfnKnowledgeBase_QueryGenerationContextProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.QueryGenerationTableProperty",
		reflect.TypeOf((*CfnKnowledgeBase_QueryGenerationTableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.RdsConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_RdsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.RdsFieldMappingProperty",
		reflect.TypeOf((*CfnKnowledgeBase_RdsFieldMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.RedshiftConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_RedshiftConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.RedshiftProvisionedAuthConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_RedshiftProvisionedAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.RedshiftProvisionedConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_RedshiftProvisionedConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.RedshiftQueryEngineAwsDataCatalogStorageConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_RedshiftQueryEngineAwsDataCatalogStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.RedshiftQueryEngineConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_RedshiftQueryEngineConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.RedshiftQueryEngineRedshiftStorageConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_RedshiftQueryEngineRedshiftStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.RedshiftQueryEngineStorageConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_RedshiftQueryEngineStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.RedshiftServerlessAuthConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_RedshiftServerlessAuthConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.RedshiftServerlessConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_RedshiftServerlessConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.S3LocationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.SqlKnowledgeBaseConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_SqlKnowledgeBaseConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.StorageConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_StorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.SupplementalDataStorageConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_SupplementalDataStorageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.SupplementalDataStorageLocationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_SupplementalDataStorageLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBase.VectorKnowledgeBaseConfigurationProperty",
		reflect.TypeOf((*CfnKnowledgeBase_VectorKnowledgeBaseConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnKnowledgeBaseProps",
		reflect.TypeOf((*CfnKnowledgeBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnPrompt",
		reflect.TypeOf((*CfnPrompt)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersion", GoGetter: "AttrVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customerEncryptionKeyArn", GoGetter: "CustomerEncryptionKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVariant", GoGetter: "DefaultVariant"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
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
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "variants", GoGetter: "Variants"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrompt{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.CachePointBlockProperty",
		reflect.TypeOf((*CfnPrompt_CachePointBlockProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.ChatPromptTemplateConfigurationProperty",
		reflect.TypeOf((*CfnPrompt_ChatPromptTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.ContentBlockProperty",
		reflect.TypeOf((*CfnPrompt_ContentBlockProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.MessageProperty",
		reflect.TypeOf((*CfnPrompt_MessageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.PromptAgentResourceProperty",
		reflect.TypeOf((*CfnPrompt_PromptAgentResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.PromptGenAiResourceProperty",
		reflect.TypeOf((*CfnPrompt_PromptGenAiResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.PromptInferenceConfigurationProperty",
		reflect.TypeOf((*CfnPrompt_PromptInferenceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.PromptInputVariableProperty",
		reflect.TypeOf((*CfnPrompt_PromptInputVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.PromptMetadataEntryProperty",
		reflect.TypeOf((*CfnPrompt_PromptMetadataEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.PromptModelInferenceConfigurationProperty",
		reflect.TypeOf((*CfnPrompt_PromptModelInferenceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.PromptTemplateConfigurationProperty",
		reflect.TypeOf((*CfnPrompt_PromptTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.PromptVariantProperty",
		reflect.TypeOf((*CfnPrompt_PromptVariantProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.SpecificToolChoiceProperty",
		reflect.TypeOf((*CfnPrompt_SpecificToolChoiceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.SystemContentBlockProperty",
		reflect.TypeOf((*CfnPrompt_SystemContentBlockProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.TextPromptTemplateConfigurationProperty",
		reflect.TypeOf((*CfnPrompt_TextPromptTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.TextS3LocationProperty",
		reflect.TypeOf((*CfnPrompt_TextS3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.ToolChoiceProperty",
		reflect.TypeOf((*CfnPrompt_ToolChoiceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.ToolConfigurationProperty",
		reflect.TypeOf((*CfnPrompt_ToolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.ToolInputSchemaProperty",
		reflect.TypeOf((*CfnPrompt_ToolInputSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.ToolProperty",
		reflect.TypeOf((*CfnPrompt_ToolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPrompt.ToolSpecificationProperty",
		reflect.TypeOf((*CfnPrompt_ToolSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptProps",
		reflect.TypeOf((*CfnPromptProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion",
		reflect.TypeOf((*CfnPromptVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrCustomerEncryptionKeyArn", GoGetter: "AttrCustomerEncryptionKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrDefaultVariant", GoGetter: "AttrDefaultVariant"},
			_jsii_.MemberProperty{JsiiProperty: "attrName", GoGetter: "AttrName"},
			_jsii_.MemberProperty{JsiiProperty: "attrPromptId", GoGetter: "AttrPromptId"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedAt", GoGetter: "AttrUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrVariants", GoGetter: "AttrVariants"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersion", GoGetter: "AttrVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "promptArn", GoGetter: "PromptArn"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPromptVersion{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.CachePointBlockProperty",
		reflect.TypeOf((*CfnPromptVersion_CachePointBlockProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.ChatPromptTemplateConfigurationProperty",
		reflect.TypeOf((*CfnPromptVersion_ChatPromptTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.ContentBlockProperty",
		reflect.TypeOf((*CfnPromptVersion_ContentBlockProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.MessageProperty",
		reflect.TypeOf((*CfnPromptVersion_MessageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.PromptAgentResourceProperty",
		reflect.TypeOf((*CfnPromptVersion_PromptAgentResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.PromptGenAiResourceProperty",
		reflect.TypeOf((*CfnPromptVersion_PromptGenAiResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.PromptInferenceConfigurationProperty",
		reflect.TypeOf((*CfnPromptVersion_PromptInferenceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.PromptInputVariableProperty",
		reflect.TypeOf((*CfnPromptVersion_PromptInputVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.PromptMetadataEntryProperty",
		reflect.TypeOf((*CfnPromptVersion_PromptMetadataEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.PromptModelInferenceConfigurationProperty",
		reflect.TypeOf((*CfnPromptVersion_PromptModelInferenceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.PromptTemplateConfigurationProperty",
		reflect.TypeOf((*CfnPromptVersion_PromptTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.PromptVariantProperty",
		reflect.TypeOf((*CfnPromptVersion_PromptVariantProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.SpecificToolChoiceProperty",
		reflect.TypeOf((*CfnPromptVersion_SpecificToolChoiceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.SystemContentBlockProperty",
		reflect.TypeOf((*CfnPromptVersion_SystemContentBlockProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.TextPromptTemplateConfigurationProperty",
		reflect.TypeOf((*CfnPromptVersion_TextPromptTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.ToolChoiceProperty",
		reflect.TypeOf((*CfnPromptVersion_ToolChoiceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.ToolConfigurationProperty",
		reflect.TypeOf((*CfnPromptVersion_ToolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.ToolInputSchemaProperty",
		reflect.TypeOf((*CfnPromptVersion_ToolInputSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.ToolProperty",
		reflect.TypeOf((*CfnPromptVersion_ToolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersion.ToolSpecificationProperty",
		reflect.TypeOf((*CfnPromptVersion_ToolSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_bedrock.CfnPromptVersionProps",
		reflect.TypeOf((*CfnPromptVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.FoundationModel",
		reflect.TypeOf((*FoundationModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "modelArn", GoGetter: "ModelArn"},
			_jsii_.MemberProperty{JsiiProperty: "modelId", GoGetter: "ModelId"},
		},
		func() interface{} {
			j := jsiiProxy_FoundationModel{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IModel)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		reflect.TypeOf((*FoundationModelIdentifier)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "modelId", GoGetter: "ModelId"},
		},
		func() interface{} {
			return &jsiiProxy_FoundationModelIdentifier{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_bedrock.IModel",
		reflect.TypeOf((*IModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "modelArn", GoGetter: "ModelArn"},
		},
		func() interface{} {
			return &jsiiProxy_IModel{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.ProvisionedModel",
		reflect.TypeOf((*ProvisionedModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "modelArn", GoGetter: "ModelArn"},
		},
		func() interface{} {
			j := jsiiProxy_ProvisionedModel{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IModel)
			return &j
		},
	)
}
