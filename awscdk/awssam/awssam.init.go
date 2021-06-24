package awssam

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_sam.CfnApi",
		reflect.TypeOf((*CfnApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessLogSetting", GoGetter: "AccessLogSetting"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "auth", GoGetter: "Auth"},
			_jsii_.MemberProperty{JsiiProperty: "binaryMediaTypes", GoGetter: "BinaryMediaTypes"},
			_jsii_.MemberProperty{JsiiProperty: "cacheClusterEnabled", GoGetter: "CacheClusterEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "cacheClusterSize", GoGetter: "CacheClusterSize"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cors", GoGetter: "Cors"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "definitionBody", GoGetter: "DefinitionBody"},
			_jsii_.MemberProperty{JsiiProperty: "definitionUri", GoGetter: "DefinitionUri"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfiguration", GoGetter: "EndpointConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "methodSettings", GoGetter: "MethodSettings"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "openApiVersion", GoGetter: "OpenApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "tracingEnabled", GoGetter: "TracingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApi{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnApi.AccessLogSettingProperty",
		reflect.TypeOf((*CfnApi_AccessLogSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnApi.AuthProperty",
		reflect.TypeOf((*CfnApi_AuthProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnApi.CorsConfigurationProperty",
		reflect.TypeOf((*CfnApi_CorsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnApi.S3LocationProperty",
		reflect.TypeOf((*CfnApi_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnApiProps",
		reflect.TypeOf((*CfnApiProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_sam.CfnApplication",
		reflect.TypeOf((*CfnApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationArns", GoGetter: "NotificationArns"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInMinutes", GoGetter: "TimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnApplication.ApplicationLocationProperty",
		reflect.TypeOf((*CfnApplication_ApplicationLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnApplicationProps",
		reflect.TypeOf((*CfnApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_sam.CfnFunction",
		reflect.TypeOf((*CfnFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoPublishAlias", GoGetter: "AutoPublishAlias"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "codeUri", GoGetter: "CodeUri"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterQueue", GoGetter: "DeadLetterQueue"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentPreference", GoGetter: "DeploymentPreference"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "events", GoGetter: "Events"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemConfigs", GoGetter: "FileSystemConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyArn", GoGetter: "KmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "layers", GoGetter: "Layers"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "memorySize", GoGetter: "MemorySize"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "permissionsBoundary", GoGetter: "PermissionsBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "policies", GoGetter: "Policies"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "provisionedConcurrencyConfig", GoGetter: "ProvisionedConcurrencyConfig"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "reservedConcurrentExecutions", GoGetter: "ReservedConcurrentExecutions"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "runtime", GoGetter: "Runtime"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "tracing", GoGetter: "Tracing"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfig", GoGetter: "VpcConfig"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFunction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.AlexaSkillEventProperty",
		reflect.TypeOf((*CfnFunction_AlexaSkillEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.ApiEventProperty",
		reflect.TypeOf((*CfnFunction_ApiEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.BucketSAMPTProperty",
		reflect.TypeOf((*CfnFunction_BucketSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.CloudWatchEventEventProperty",
		reflect.TypeOf((*CfnFunction_CloudWatchEventEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.CloudWatchLogsEventProperty",
		reflect.TypeOf((*CfnFunction_CloudWatchLogsEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.CollectionSAMPTProperty",
		reflect.TypeOf((*CfnFunction_CollectionSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.DeadLetterQueueProperty",
		reflect.TypeOf((*CfnFunction_DeadLetterQueueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.DeploymentPreferenceProperty",
		reflect.TypeOf((*CfnFunction_DeploymentPreferenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.DestinationConfigProperty",
		reflect.TypeOf((*CfnFunction_DestinationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.DomainSAMPTProperty",
		reflect.TypeOf((*CfnFunction_DomainSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.DynamoDBEventProperty",
		reflect.TypeOf((*CfnFunction_DynamoDBEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.EmptySAMPTProperty",
		reflect.TypeOf((*CfnFunction_EmptySAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.EventBridgeRuleEventProperty",
		reflect.TypeOf((*CfnFunction_EventBridgeRuleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.EventSourceProperty",
		reflect.TypeOf((*CfnFunction_EventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.FileSystemConfigProperty",
		reflect.TypeOf((*CfnFunction_FileSystemConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.FunctionEnvironmentProperty",
		reflect.TypeOf((*CfnFunction_FunctionEnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.FunctionSAMPTProperty",
		reflect.TypeOf((*CfnFunction_FunctionSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_sam.CfnFunction.IAMPolicyDocumentProperty",
		reflect.TypeOf((*CfnFunction_IAMPolicyDocumentProperty)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "statement", GoGetter: "Statement"},
		},
		func() interface{} {
			return &jsiiProxy_CfnFunction_IAMPolicyDocumentProperty{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.IdentitySAMPTProperty",
		reflect.TypeOf((*CfnFunction_IdentitySAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.IoTRuleEventProperty",
		reflect.TypeOf((*CfnFunction_IoTRuleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.KeySAMPTProperty",
		reflect.TypeOf((*CfnFunction_KeySAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.KinesisEventProperty",
		reflect.TypeOf((*CfnFunction_KinesisEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.LogGroupSAMPTProperty",
		reflect.TypeOf((*CfnFunction_LogGroupSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.OnFailureProperty",
		reflect.TypeOf((*CfnFunction_OnFailureProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.ProvisionedConcurrencyConfigProperty",
		reflect.TypeOf((*CfnFunction_ProvisionedConcurrencyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.QueueSAMPTProperty",
		reflect.TypeOf((*CfnFunction_QueueSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.S3EventProperty",
		reflect.TypeOf((*CfnFunction_S3EventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.S3KeyFilterProperty",
		reflect.TypeOf((*CfnFunction_S3KeyFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.S3KeyFilterRuleProperty",
		reflect.TypeOf((*CfnFunction_S3KeyFilterRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.S3LocationProperty",
		reflect.TypeOf((*CfnFunction_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.S3NotificationFilterProperty",
		reflect.TypeOf((*CfnFunction_S3NotificationFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.SAMPolicyTemplateProperty",
		reflect.TypeOf((*CfnFunction_SAMPolicyTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.SNSEventProperty",
		reflect.TypeOf((*CfnFunction_SNSEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.SQSEventProperty",
		reflect.TypeOf((*CfnFunction_SQSEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.ScheduleEventProperty",
		reflect.TypeOf((*CfnFunction_ScheduleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.StateMachineSAMPTProperty",
		reflect.TypeOf((*CfnFunction_StateMachineSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.StreamSAMPTProperty",
		reflect.TypeOf((*CfnFunction_StreamSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.TableSAMPTProperty",
		reflect.TypeOf((*CfnFunction_TableSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.TableStreamSAMPTProperty",
		reflect.TypeOf((*CfnFunction_TableStreamSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.TopicSAMPTProperty",
		reflect.TypeOf((*CfnFunction_TopicSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunction.VpcConfigProperty",
		reflect.TypeOf((*CfnFunction_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnFunctionProps",
		reflect.TypeOf((*CfnFunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_sam.CfnLayerVersion",
		reflect.TypeOf((*CfnLayerVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "compatibleRuntimes", GoGetter: "CompatibleRuntimes"},
			_jsii_.MemberProperty{JsiiProperty: "contentUri", GoGetter: "ContentUri"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "layerName", GoGetter: "LayerName"},
			_jsii_.MemberProperty{JsiiProperty: "licenseInfo", GoGetter: "LicenseInfo"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "retentionPolicy", GoGetter: "RetentionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLayerVersion{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnLayerVersion.S3LocationProperty",
		reflect.TypeOf((*CfnLayerVersion_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnLayerVersionProps",
		reflect.TypeOf((*CfnLayerVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_sam.CfnSimpleTable",
		reflect.TypeOf((*CfnSimpleTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberProperty{JsiiProperty: "provisionedThroughput", GoGetter: "ProvisionedThroughput"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sseSpecification", GoGetter: "SseSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSimpleTable{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnSimpleTable.PrimaryKeyProperty",
		reflect.TypeOf((*CfnSimpleTable_PrimaryKeyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnSimpleTable.ProvisionedThroughputProperty",
		reflect.TypeOf((*CfnSimpleTable_ProvisionedThroughputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnSimpleTable.SSESpecificationProperty",
		reflect.TypeOf((*CfnSimpleTable_SSESpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnSimpleTableProps",
		reflect.TypeOf((*CfnSimpleTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_sam.CfnStateMachine",
		reflect.TypeOf((*CfnStateMachine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "definitionSubstitutions", GoGetter: "DefinitionSubstitutions"},
			_jsii_.MemberProperty{JsiiProperty: "definitionUri", GoGetter: "DefinitionUri"},
			_jsii_.MemberProperty{JsiiProperty: "events", GoGetter: "Events"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policies", GoGetter: "Policies"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStateMachine{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachine.ApiEventProperty",
		reflect.TypeOf((*CfnStateMachine_ApiEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachine.CloudWatchEventEventProperty",
		reflect.TypeOf((*CfnStateMachine_CloudWatchEventEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachine.CloudWatchLogsLogGroupProperty",
		reflect.TypeOf((*CfnStateMachine_CloudWatchLogsLogGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachine.EventBridgeRuleEventProperty",
		reflect.TypeOf((*CfnStateMachine_EventBridgeRuleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachine.EventSourceProperty",
		reflect.TypeOf((*CfnStateMachine_EventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachine.FunctionSAMPTProperty",
		reflect.TypeOf((*CfnStateMachine_FunctionSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_sam.CfnStateMachine.IAMPolicyDocumentProperty",
		reflect.TypeOf((*CfnStateMachine_IAMPolicyDocumentProperty)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "statement", GoGetter: "Statement"},
		},
		func() interface{} {
			return &jsiiProxy_CfnStateMachine_IAMPolicyDocumentProperty{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachine.LogDestinationProperty",
		reflect.TypeOf((*CfnStateMachine_LogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachine.LoggingConfigurationProperty",
		reflect.TypeOf((*CfnStateMachine_LoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachine.S3LocationProperty",
		reflect.TypeOf((*CfnStateMachine_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachine.SAMPolicyTemplateProperty",
		reflect.TypeOf((*CfnStateMachine_SAMPolicyTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachine.ScheduleEventProperty",
		reflect.TypeOf((*CfnStateMachine_ScheduleEventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachine.StateMachineSAMPTProperty",
		reflect.TypeOf((*CfnStateMachine_StateMachineSAMPTProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_sam.CfnStateMachineProps",
		reflect.TypeOf((*CfnStateMachineProps)(nil)).Elem(),
	)
}
