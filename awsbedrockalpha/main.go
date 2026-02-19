// The CDK Construct Library for Amazon Bedrock
package awsbedrockalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.ActionGroupExecutor",
		reflect.TypeOf((*ActionGroupExecutor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "customControl", GoGetter: "CustomControl"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
		},
		func() interface{} {
			return &jsiiProxy_ActionGroupExecutor{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.Agent",
		reflect.TypeOf((*Agent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionGroups", GoGetter: "ActionGroups"},
			_jsii_.MemberMethod{JsiiMethod: "addActionGroup", GoMethod: "AddActionGroup"},
			_jsii_.MemberMethod{JsiiMethod: "addActionGroups", GoMethod: "AddActionGroups"},
			_jsii_.MemberMethod{JsiiMethod: "addGuardrail", GoMethod: "AddGuardrail"},
			_jsii_.MemberProperty{JsiiProperty: "agentArn", GoGetter: "AgentArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentId", GoGetter: "AgentId"},
			_jsii_.MemberProperty{JsiiProperty: "agentVersion", GoGetter: "AgentVersion"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "guardrail", GoGetter: "Guardrail"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdated", GoGetter: "LastUpdated"},
			_jsii_.MemberMethod{JsiiMethod: "metricCount", GoMethod: "MetricCount"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "testAlias", GoGetter: "TestAlias"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Agent{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AgentBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAgent)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.AgentActionGroup",
		reflect.TypeOf((*AgentActionGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiSchema", GoGetter: "ApiSchema"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "executor", GoGetter: "Executor"},
			_jsii_.MemberProperty{JsiiProperty: "forceDelete", GoGetter: "ForceDelete"},
			_jsii_.MemberProperty{JsiiProperty: "functionSchema", GoGetter: "FunctionSchema"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentActionGroupSignature", GoGetter: "ParentActionGroupSignature"},
		},
		func() interface{} {
			return &jsiiProxy_AgentActionGroup{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.AgentActionGroupProps",
		reflect.TypeOf((*AgentActionGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.AgentAlias",
		reflect.TypeOf((*AgentAlias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agent", GoGetter: "Agent"},
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasId", GoGetter: "AliasId"},
			_jsii_.MemberProperty{JsiiProperty: "aliasName", GoGetter: "AliasName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantGet", GoMethod: "GrantGet"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_AgentAlias{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AgentAliasBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.AgentAliasAttributes",
		reflect.TypeOf((*AgentAliasAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.AgentAliasBase",
		reflect.TypeOf((*AgentAliasBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agent", GoGetter: "Agent"},
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasId", GoGetter: "AliasId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantGet", GoMethod: "GrantGet"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_AgentAliasBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAgentAlias)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.AgentAliasProps",
		reflect.TypeOf((*AgentAliasProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.AgentAttributes",
		reflect.TypeOf((*AgentAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.AgentBase",
		reflect.TypeOf((*AgentBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentArn", GoGetter: "AgentArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentId", GoGetter: "AgentId"},
			_jsii_.MemberProperty{JsiiProperty: "agentVersion", GoGetter: "AgentVersion"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdated", GoGetter: "LastUpdated"},
			_jsii_.MemberMethod{JsiiMethod: "metricCount", GoMethod: "MetricCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_AgentBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAgent)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.AgentCollaboration",
		reflect.TypeOf((*AgentCollaboration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "collaborators", GoGetter: "Collaborators"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_AgentCollaboration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.AgentCollaborationConfig",
		reflect.TypeOf((*AgentCollaborationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.AgentCollaborator",
		reflect.TypeOf((*AgentCollaborator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentAlias", GoGetter: "AgentAlias"},
			_jsii_.MemberProperty{JsiiProperty: "collaborationInstruction", GoGetter: "CollaborationInstruction"},
			_jsii_.MemberProperty{JsiiProperty: "collaboratorName", GoGetter: "CollaboratorName"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "relayConversationHistory", GoGetter: "RelayConversationHistory"},
		},
		func() interface{} {
			return &jsiiProxy_AgentCollaborator{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.AgentCollaboratorProps",
		reflect.TypeOf((*AgentCollaboratorProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.AgentCollaboratorType",
		reflect.TypeOf((*AgentCollaboratorType)(nil)).Elem(),
		map[string]interface{}{
			"SUPERVISOR": AgentCollaboratorType_SUPERVISOR,
			"DISABLED": AgentCollaboratorType_DISABLED,
			"SUPERVISOR_ROUTER": AgentCollaboratorType_SUPERVISOR_ROUTER,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.AgentGenAiResourceProps",
		reflect.TypeOf((*AgentGenAiResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.AgentPromptVariantProps",
		reflect.TypeOf((*AgentPromptVariantProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.AgentProps",
		reflect.TypeOf((*AgentProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.AgentStepType",
		reflect.TypeOf((*AgentStepType)(nil)).Elem(),
		map[string]interface{}{
			"PRE_PROCESSING": AgentStepType_PRE_PROCESSING,
			"ORCHESTRATION": AgentStepType_ORCHESTRATION,
			"POST_PROCESSING": AgentStepType_POST_PROCESSING,
			"ROUTING_CLASSIFIER": AgentStepType_ROUTING_CLASSIFIER,
			"MEMORY_SUMMARIZATION": AgentStepType_MEMORY_SUMMARIZATION,
			"KNOWLEDGE_BASE_RESPONSE_GENERATION": AgentStepType_KNOWLEDGE_BASE_RESPONSE_GENERATION,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.ApiSchema",
		reflect.TypeOf((*ApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			return &jsiiProxy_ApiSchema{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.ApplicationInferenceProfile",
		reflect.TypeOf((*ApplicationInferenceProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantProfileUsage", GoMethod: "GrantProfileUsage"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileArn", GoGetter: "InferenceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileId", GoGetter: "InferenceProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileModel", GoGetter: "InferenceProfileModel"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileName", GoGetter: "InferenceProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "invokableArn", GoGetter: "InvokableArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationInferenceProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InferenceProfileBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBedrockInvokable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.ApplicationInferenceProfileAttributes",
		reflect.TypeOf((*ApplicationInferenceProfileAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.ApplicationInferenceProfileProps",
		reflect.TypeOf((*ApplicationInferenceProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.AssetApiSchema",
		reflect.TypeOf((*AssetApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			j := jsiiProxy_AssetApiSchema{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ApiSchema)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModel",
		reflect.TypeOf((*BedrockFoundationModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "asArn", GoMethod: "AsArn"},
			_jsii_.MemberMethod{JsiiMethod: "asIModel", GoMethod: "AsIModel"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeAllRegions", GoMethod: "GrantInvokeAllRegions"},
			_jsii_.MemberProperty{JsiiProperty: "invokableArn", GoGetter: "InvokableArn"},
			_jsii_.MemberProperty{JsiiProperty: "modelArn", GoGetter: "ModelArn"},
			_jsii_.MemberProperty{JsiiProperty: "modelId", GoGetter: "ModelId"},
			_jsii_.MemberProperty{JsiiProperty: "supportedVectorType", GoGetter: "SupportedVectorType"},
			_jsii_.MemberProperty{JsiiProperty: "supportsAgents", GoGetter: "SupportsAgents"},
			_jsii_.MemberProperty{JsiiProperty: "supportsCrossRegion", GoGetter: "SupportsCrossRegion"},
			_jsii_.MemberProperty{JsiiProperty: "supportsKnowledgeBase", GoGetter: "SupportsKnowledgeBase"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorDimensions", GoGetter: "VectorDimensions"},
		},
		func() interface{} {
			j := jsiiProxy_BedrockFoundationModel{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBedrockInvokable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.BedrockFoundationModelProps",
		reflect.TypeOf((*BedrockFoundationModelProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.CanadaSpecificPIIType",
		reflect.TypeOf((*CanadaSpecificPIIType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_CanadaSpecificPIIType{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PIIType)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.ChatMessage",
		reflect.TypeOf((*ChatMessage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
		},
		func() interface{} {
			return &jsiiProxy_ChatMessage{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.ChatMessageRole",
		reflect.TypeOf((*ChatMessageRole)(nil)).Elem(),
		map[string]interface{}{
			"USER": ChatMessageRole_USER,
			"ASSISTANT": ChatMessageRole_ASSISTANT,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.ChatPromptVariantProps",
		reflect.TypeOf((*ChatPromptVariantProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.ChatTemplateConfigurationProps",
		reflect.TypeOf((*ChatTemplateConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.CommonPromptVariantProps",
		reflect.TypeOf((*CommonPromptVariantProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.ContentFilter",
		reflect.TypeOf((*ContentFilter)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.ContentFilterStrength",
		reflect.TypeOf((*ContentFilterStrength)(nil)).Elem(),
		map[string]interface{}{
			"NONE": ContentFilterStrength_NONE,
			"LOW": ContentFilterStrength_LOW,
			"MEDIUM": ContentFilterStrength_MEDIUM,
			"HIGH": ContentFilterStrength_HIGH,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.ContentFilterType",
		reflect.TypeOf((*ContentFilterType)(nil)).Elem(),
		map[string]interface{}{
			"SEXUAL": ContentFilterType_SEXUAL,
			"VIOLENCE": ContentFilterType_VIOLENCE,
			"HATE": ContentFilterType_HATE,
			"INSULTS": ContentFilterType_INSULTS,
			"MISCONDUCT": ContentFilterType_MISCONDUCT,
			"PROMPT_ATTACK": ContentFilterType_PROMPT_ATTACK,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.ContextualGroundingFilter",
		reflect.TypeOf((*ContextualGroundingFilter)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.ContextualGroundingFilterType",
		reflect.TypeOf((*ContextualGroundingFilterType)(nil)).Elem(),
		map[string]interface{}{
			"GROUNDING": ContextualGroundingFilterType_GROUNDING,
			"RELEVANCE": ContextualGroundingFilterType_RELEVANCE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.CrossRegionInferenceProfile",
		reflect.TypeOf((*CrossRegionInferenceProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantProfileUsage", GoMethod: "GrantProfileUsage"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileArn", GoGetter: "InferenceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileId", GoGetter: "InferenceProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileModel", GoGetter: "InferenceProfileModel"},
			_jsii_.MemberProperty{JsiiProperty: "invokableArn", GoGetter: "InvokableArn"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_CrossRegionInferenceProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBedrockInvokable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInferenceProfile)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.CrossRegionInferenceProfileProps",
		reflect.TypeOf((*CrossRegionInferenceProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.CrossRegionInferenceProfileRegion",
		reflect.TypeOf((*CrossRegionInferenceProfileRegion)(nil)).Elem(),
		map[string]interface{}{
			"GLOBAL": CrossRegionInferenceProfileRegion_GLOBAL,
			"EU": CrossRegionInferenceProfileRegion_EU,
			"US": CrossRegionInferenceProfileRegion_US,
			"US_GOV": CrossRegionInferenceProfileRegion_US_GOV,
			"APAC": CrossRegionInferenceProfileRegion_APAC,
			"JP": CrossRegionInferenceProfileRegion_JP,
			"AU": CrossRegionInferenceProfileRegion_AU,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.CustomControl",
		reflect.TypeOf((*CustomControl)(nil)).Elem(),
		map[string]interface{}{
			"RETURN_CONTROL": CustomControl_RETURN_CONTROL,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.CustomOrchestrationExecutor",
		reflect.TypeOf((*CustomOrchestrationExecutor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_CustomOrchestrationExecutor{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.CustomParserProps",
		reflect.TypeOf((*CustomParserProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.CustomTopicProps",
		reflect.TypeOf((*CustomTopicProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.DefaultPromptRouterIdentifier",
		reflect.TypeOf((*DefaultPromptRouterIdentifier)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "promptRouterId", GoGetter: "PromptRouterId"},
			_jsii_.MemberProperty{JsiiProperty: "routingModels", GoGetter: "RoutingModels"},
		},
		func() interface{} {
			return &jsiiProxy_DefaultPromptRouterIdentifier{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.FinancePIIType",
		reflect.TypeOf((*FinancePIIType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_FinancePIIType{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PIIType)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.Function",
		reflect.TypeOf((*Function)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "requireConfirmation", GoGetter: "RequireConfirmation"},
		},
		func() interface{} {
			return &jsiiProxy_Function{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.FunctionParameter",
		reflect.TypeOf((*FunctionParameter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "required", GoGetter: "Required"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_FunctionParameter{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.FunctionParameterProps",
		reflect.TypeOf((*FunctionParameterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.FunctionProps",
		reflect.TypeOf((*FunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.FunctionSchema",
		reflect.TypeOf((*FunctionSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "functions", GoGetter: "Functions"},
		},
		func() interface{} {
			return &jsiiProxy_FunctionSchema{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.FunctionSchemaProps",
		reflect.TypeOf((*FunctionSchemaProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.FunctionToolProps",
		reflect.TypeOf((*FunctionToolProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.GeneralPIIType",
		reflect.TypeOf((*GeneralPIIType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_GeneralPIIType{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PIIType)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.Guardrail",
		reflect.TypeOf((*Guardrail)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContentFilter", GoMethod: "AddContentFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addContextualGroundingFilter", GoMethod: "AddContextualGroundingFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addDeniedTopicFilter", GoMethod: "AddDeniedTopicFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addManagedWordListFilter", GoMethod: "AddManagedWordListFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addPIIFilter", GoMethod: "AddPIIFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addRegexFilter", GoMethod: "AddRegexFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addWordFilter", GoMethod: "AddWordFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addWordFilterFromFile", GoMethod: "AddWordFilterFromFile"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "blockedInputMessaging", GoGetter: "BlockedInputMessaging"},
			_jsii_.MemberProperty{JsiiProperty: "blockedOutputsMessaging", GoGetter: "BlockedOutputsMessaging"},
			_jsii_.MemberProperty{JsiiProperty: "contentFilters", GoGetter: "ContentFilters"},
			_jsii_.MemberProperty{JsiiProperty: "contentFiltersTierConfig", GoGetter: "ContentFiltersTierConfig"},
			_jsii_.MemberProperty{JsiiProperty: "contextualGroundingFilters", GoGetter: "ContextualGroundingFilters"},
			_jsii_.MemberMethod{JsiiMethod: "createVersion", GoMethod: "CreateVersion"},
			_jsii_.MemberProperty{JsiiProperty: "crossRegionConfig", GoGetter: "CrossRegionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "deniedTopics", GoGetter: "DeniedTopics"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantApply", GoMethod: "GrantApply"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailArn", GoGetter: "GuardrailArn"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailId", GoGetter: "GuardrailId"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailVersion", GoGetter: "GuardrailVersion"},
			_jsii_.MemberProperty{JsiiProperty: "hash", GoGetter: "Hash"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdated", GoGetter: "LastUpdated"},
			_jsii_.MemberProperty{JsiiProperty: "managedWordListFilters", GoGetter: "ManagedWordListFilters"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationClientErrors", GoMethod: "MetricInvocationClientErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationLatency", GoMethod: "MetricInvocationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationServerErrors", GoMethod: "MetricInvocationServerErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsIntervened", GoMethod: "MetricInvocationsIntervened"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationThrottles", GoMethod: "MetricInvocationThrottles"},
			_jsii_.MemberMethod{JsiiMethod: "metricTextUnitCount", GoMethod: "MetricTextUnitCount"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "piiFilters", GoGetter: "PiiFilters"},
			_jsii_.MemberProperty{JsiiProperty: "regexFilters", GoGetter: "RegexFilters"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "topicsTierConfig", GoGetter: "TopicsTierConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateVersion", GoMethod: "UpdateVersion"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
			_jsii_.MemberProperty{JsiiProperty: "wordFilters", GoGetter: "WordFilters"},
		},
		func() interface{} {
			j := jsiiProxy_Guardrail{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_GuardrailBase)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.GuardrailAction",
		reflect.TypeOf((*GuardrailAction)(nil)).Elem(),
		map[string]interface{}{
			"BLOCK": GuardrailAction_BLOCK,
			"ANONYMIZE": GuardrailAction_ANONYMIZE,
			"NONE": GuardrailAction_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.GuardrailAttributes",
		reflect.TypeOf((*GuardrailAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.GuardrailBase",
		reflect.TypeOf((*GuardrailBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantApply", GoMethod: "GrantApply"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailArn", GoGetter: "GuardrailArn"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailId", GoGetter: "GuardrailId"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailVersion", GoGetter: "GuardrailVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdated", GoGetter: "LastUpdated"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationClientErrors", GoMethod: "MetricInvocationClientErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationLatency", GoMethod: "MetricInvocationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationServerErrors", GoMethod: "MetricInvocationServerErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsIntervened", GoMethod: "MetricInvocationsIntervened"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationThrottles", GoMethod: "MetricInvocationThrottles"},
			_jsii_.MemberMethod{JsiiMethod: "metricTextUnitCount", GoMethod: "MetricTextUnitCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateVersion", GoMethod: "UpdateVersion"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_GuardrailBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGuardrail)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.GuardrailCrossRegionConfigProperty",
		reflect.TypeOf((*GuardrailCrossRegionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.GuardrailProps",
		reflect.TypeOf((*GuardrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-alpha.IAgent",
		reflect.TypeOf((*IAgent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentArn", GoGetter: "AgentArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentId", GoGetter: "AgentId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdated", GoGetter: "LastUpdated"},
			_jsii_.MemberMethod{JsiiMethod: "metricCount", GoMethod: "MetricCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IAgent{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-alpha.IAgentAlias",
		reflect.TypeOf((*IAgentAlias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agent", GoGetter: "Agent"},
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasId", GoGetter: "AliasId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantGet", GoMethod: "GrantGet"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IAgentAlias{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-alpha.IBedrockInvokable",
		reflect.TypeOf((*IBedrockInvokable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "invokableArn", GoGetter: "InvokableArn"},
		},
		func() interface{} {
			return &jsiiProxy_IBedrockInvokable{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-alpha.IGuardrail",
		reflect.TypeOf((*IGuardrail)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantApply", GoMethod: "GrantApply"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailArn", GoGetter: "GuardrailArn"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailId", GoGetter: "GuardrailId"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailVersion", GoGetter: "GuardrailVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdated", GoGetter: "LastUpdated"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationClientErrors", GoMethod: "MetricInvocationClientErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationLatency", GoMethod: "MetricInvocationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationServerErrors", GoMethod: "MetricInvocationServerErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsIntervened", GoMethod: "MetricInvocationsIntervened"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationThrottles", GoMethod: "MetricInvocationThrottles"},
			_jsii_.MemberMethod{JsiiMethod: "metricTextUnitCount", GoMethod: "MetricTextUnitCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IGuardrail{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-alpha.IInferenceProfile",
		reflect.TypeOf((*IInferenceProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantProfileUsage", GoMethod: "GrantProfileUsage"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileArn", GoGetter: "InferenceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileId", GoGetter: "InferenceProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IInferenceProfile{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-alpha.IPrompt",
		reflect.TypeOf((*IPrompt)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantGet", GoMethod: "GrantGet"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "promptArn", GoGetter: "PromptArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptId", GoGetter: "PromptId"},
			_jsii_.MemberProperty{JsiiProperty: "promptVersion", GoGetter: "PromptVersion"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IPrompt{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-alpha.IPromptRouter",
		reflect.TypeOf((*IPromptRouter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "promptRouterArn", GoGetter: "PromptRouterArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptRouterId", GoGetter: "PromptRouterId"},
			_jsii_.MemberProperty{JsiiProperty: "routingEndpoints", GoGetter: "RoutingEndpoints"},
		},
		func() interface{} {
			return &jsiiProxy_IPromptRouter{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-bedrock-alpha.IPromptVariant",
		reflect.TypeOf((*IPromptVariant)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "genAiResource", GoGetter: "GenAiResource"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceConfiguration", GoGetter: "InferenceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "modelId", GoGetter: "ModelId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "templateConfiguration", GoGetter: "TemplateConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "templateType", GoGetter: "TemplateType"},
		},
		func() interface{} {
			return &jsiiProxy_IPromptVariant{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.InferenceConfiguration",
		reflect.TypeOf((*InferenceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.InferenceProfileBase",
		reflect.TypeOf((*InferenceProfileBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantProfileUsage", GoMethod: "GrantProfileUsage"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileArn", GoGetter: "InferenceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileId", GoGetter: "InferenceProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_InferenceProfileBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInferenceProfile)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.InferenceProfileType",
		reflect.TypeOf((*InferenceProfileType)(nil)).Elem(),
		map[string]interface{}{
			"SYSTEM_DEFINED": InferenceProfileType_SYSTEM_DEFINED,
			"APPLICATION": InferenceProfileType_APPLICATION,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.InformationTechnologyPIIType",
		reflect.TypeOf((*InformationTechnologyPIIType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_InformationTechnologyPIIType{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PIIType)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.InlineApiSchema",
		reflect.TypeOf((*InlineApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			j := jsiiProxy_InlineApiSchema{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ApiSchema)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.ManagedWordFilter",
		reflect.TypeOf((*ManagedWordFilter)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.ManagedWordFilterType",
		reflect.TypeOf((*ManagedWordFilterType)(nil)).Elem(),
		map[string]interface{}{
			"PROFANITY": ManagedWordFilterType_PROFANITY,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.Memory",
		reflect.TypeOf((*Memory)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Memory{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.ModalityType",
		reflect.TypeOf((*ModalityType)(nil)).Elem(),
		map[string]interface{}{
			"TEXT": ModalityType_TEXT,
			"IMAGE": ModalityType_IMAGE,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.OrchestrationType",
		reflect.TypeOf((*OrchestrationType)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": OrchestrationType_DEFAULT,
			"CUSTOM_ORCHESTRATION": OrchestrationType_CUSTOM_ORCHESTRATION,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PIIFilter",
		reflect.TypeOf((*PIIFilter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.PIIType",
		reflect.TypeOf((*PIIType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_PIIType{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.ParameterType",
		reflect.TypeOf((*ParameterType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": ParameterType_STRING,
			"NUMBER": ParameterType_NUMBER,
			"INTEGER": ParameterType_INTEGER,
			"BOOLEAN": ParameterType_BOOLEAN,
			"ARRAY": ParameterType_ARRAY,
			"OBJECT": ParameterType_OBJECT,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.ParentActionGroupSignature",
		reflect.TypeOf((*ParentActionGroupSignature)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ParentActionGroupSignature{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.Prompt",
		reflect.TypeOf((*Prompt)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVariant", GoMethod: "AddVariant"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "createVersion", GoMethod: "CreateVersion"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantGet", GoMethod: "GrantGet"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "promptArn", GoGetter: "PromptArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptId", GoGetter: "PromptId"},
			_jsii_.MemberProperty{JsiiProperty: "promptName", GoGetter: "PromptName"},
			_jsii_.MemberProperty{JsiiProperty: "promptVersion", GoGetter: "PromptVersion"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "variants", GoGetter: "Variants"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Prompt{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PromptBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPrompt)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PromptAttributes",
		reflect.TypeOf((*PromptAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.PromptBase",
		reflect.TypeOf((*PromptBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantGet", GoMethod: "GrantGet"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "promptArn", GoGetter: "PromptArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptId", GoGetter: "PromptId"},
			_jsii_.MemberProperty{JsiiProperty: "promptVersion", GoGetter: "PromptVersion"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_PromptBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPrompt)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.PromptGenAiResource",
		reflect.TypeOf((*PromptGenAiResource)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_PromptGenAiResource{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.PromptInferenceConfiguration",
		reflect.TypeOf((*PromptInferenceConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_PromptInferenceConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PromptInferenceConfigurationProps",
		reflect.TypeOf((*PromptInferenceConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PromptKnowledgeBaseResponseGenerationConfigCustomParser",
		reflect.TypeOf((*PromptKnowledgeBaseResponseGenerationConfigCustomParser)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PromptMemorySummarizationConfigCustomParser",
		reflect.TypeOf((*PromptMemorySummarizationConfigCustomParser)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PromptOrchestrationConfigCustomParser",
		reflect.TypeOf((*PromptOrchestrationConfigCustomParser)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.PromptOverrideConfiguration",
		reflect.TypeOf((*PromptOverrideConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseResponseGenerationStep", GoGetter: "KnowledgeBaseResponseGenerationStep"},
			_jsii_.MemberProperty{JsiiProperty: "memorySummarizationStep", GoGetter: "MemorySummarizationStep"},
			_jsii_.MemberProperty{JsiiProperty: "orchestrationStep", GoGetter: "OrchestrationStep"},
			_jsii_.MemberProperty{JsiiProperty: "parser", GoGetter: "Parser"},
			_jsii_.MemberProperty{JsiiProperty: "postProcessingStep", GoGetter: "PostProcessingStep"},
			_jsii_.MemberProperty{JsiiProperty: "preProcessingStep", GoGetter: "PreProcessingStep"},
			_jsii_.MemberProperty{JsiiProperty: "routingClassifierStep", GoGetter: "RoutingClassifierStep"},
		},
		func() interface{} {
			return &jsiiProxy_PromptOverrideConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PromptPostProcessingConfigCustomParser",
		reflect.TypeOf((*PromptPostProcessingConfigCustomParser)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PromptPreProcessingConfigCustomParser",
		reflect.TypeOf((*PromptPreProcessingConfigCustomParser)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PromptProps",
		reflect.TypeOf((*PromptProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.PromptRouter",
		reflect.TypeOf((*PromptRouter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "invokableArn", GoGetter: "InvokableArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptRouterArn", GoGetter: "PromptRouterArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptRouterId", GoGetter: "PromptRouterId"},
			_jsii_.MemberProperty{JsiiProperty: "routingEndpoints", GoGetter: "RoutingEndpoints"},
		},
		func() interface{} {
			j := jsiiProxy_PromptRouter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBedrockInvokable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPromptRouter)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PromptRouterProps",
		reflect.TypeOf((*PromptRouterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PromptRoutingClassifierConfigCustomParser",
		reflect.TypeOf((*PromptRoutingClassifierConfigCustomParser)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PromptStepConfigBase",
		reflect.TypeOf((*PromptStepConfigBase)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.PromptTemplateConfiguration",
		reflect.TypeOf((*PromptTemplateConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_PromptTemplateConfiguration{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.PromptTemplateType",
		reflect.TypeOf((*PromptTemplateType)(nil)).Elem(),
		map[string]interface{}{
			"TEXT": PromptTemplateType_TEXT,
			"CHAT": PromptTemplateType_CHAT,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.PromptVariant",
		reflect.TypeOf((*PromptVariant)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_PromptVariant{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.PromptVersion",
		reflect.TypeOf((*PromptVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "prompt", GoGetter: "Prompt"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "versionArn", GoGetter: "VersionArn"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_PromptVersion{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.PromptVersionProps",
		reflect.TypeOf((*PromptVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.RegexFilter",
		reflect.TypeOf((*RegexFilter)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.RequireConfirmation",
		reflect.TypeOf((*RequireConfirmation)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": RequireConfirmation_ENABLED,
			"DISABLED": RequireConfirmation_DISABLED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.S3ApiSchema",
		reflect.TypeOf((*S3ApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			j := jsiiProxy_S3ApiSchema{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ApiSchema)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.SessionSummaryMemoryProps",
		reflect.TypeOf((*SessionSummaryMemoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.TextPromptVariantProps",
		reflect.TypeOf((*TextPromptVariantProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.TextTemplateConfigurationProps",
		reflect.TypeOf((*TextTemplateConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.TierConfig",
		reflect.TypeOf((*TierConfig)(nil)).Elem(),
		map[string]interface{}{
			"CLASSIC": TierConfig_CLASSIC,
			"STANDARD": TierConfig_STANDARD,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.Tool",
		reflect.TypeOf((*Tool)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Tool{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.ToolChoice",
		reflect.TypeOf((*ToolChoice)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "any", GoGetter: "Any"},
			_jsii_.MemberProperty{JsiiProperty: "auto", GoGetter: "Auto"},
			_jsii_.MemberProperty{JsiiProperty: "tool", GoGetter: "Tool"},
		},
		func() interface{} {
			return &jsiiProxy_ToolChoice{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.ToolConfiguration",
		reflect.TypeOf((*ToolConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.Topic",
		reflect.TypeOf((*Topic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "examples", GoGetter: "Examples"},
			_jsii_.MemberProperty{JsiiProperty: "inputAction", GoGetter: "InputAction"},
			_jsii_.MemberProperty{JsiiProperty: "inputEnabled", GoGetter: "InputEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "outputAction", GoGetter: "OutputAction"},
			_jsii_.MemberProperty{JsiiProperty: "outputEnabled", GoGetter: "OutputEnabled"},
		},
		func() interface{} {
			return &jsiiProxy_Topic{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.UKSpecificPIIType",
		reflect.TypeOf((*UKSpecificPIIType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_UKSpecificPIIType{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PIIType)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.USASpecificPIIType",
		reflect.TypeOf((*USASpecificPIIType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_USASpecificPIIType{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PIIType)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-bedrock-alpha.VectorType",
		reflect.TypeOf((*VectorType)(nil)).Elem(),
		map[string]interface{}{
			"FLOATING_POINT": VectorType_FLOATING_POINT,
			"BINARY": VectorType_BINARY,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.WordFilter",
		reflect.TypeOf((*WordFilter)(nil)).Elem(),
	)
}
