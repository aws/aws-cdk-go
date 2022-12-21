package awscdk

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.Annotations",
		reflect.TypeOf((*Annotations)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeprecation", GoMethod: "AddDeprecation"},
			_jsii_.MemberMethod{JsiiMethod: "addError", GoMethod: "AddError"},
			_jsii_.MemberMethod{JsiiMethod: "addInfo", GoMethod: "AddInfo"},
			_jsii_.MemberMethod{JsiiMethod: "addWarning", GoMethod: "AddWarning"},
		},
		func() interface{} {
			return &jsiiProxy_Annotations{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.App",
		reflect.TypeOf((*App)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "assetOutdir", GoGetter: "AssetOutdir"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "parentStage", GoGetter: "ParentStage"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_App{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Stage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.AppProps",
		reflect.TypeOf((*AppProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Arn",
		reflect.TypeOf((*Arn)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Arn{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.ArnComponents",
		reflect.TypeOf((*ArnComponents)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.ArnFormat",
		reflect.TypeOf((*ArnFormat)(nil)).Elem(),
		map[string]interface{}{
			"NO_RESOURCE_NAME": ArnFormat_NO_RESOURCE_NAME,
			"COLON_RESOURCE_NAME": ArnFormat_COLON_RESOURCE_NAME,
			"SLASH_RESOURCE_NAME": ArnFormat_SLASH_RESOURCE_NAME,
			"SLASH_RESOURCE_SLASH_RESOURCE_NAME": ArnFormat_SLASH_RESOURCE_SLASH_RESOURCE_NAME,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Aspects",
		reflect.TypeOf((*Aspects)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "add", GoMethod: "Add"},
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
		},
		func() interface{} {
			return &jsiiProxy_Aspects{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.AssetHashType",
		reflect.TypeOf((*AssetHashType)(nil)).Elem(),
		map[string]interface{}{
			"SOURCE": AssetHashType_SOURCE,
			"OUTPUT": AssetHashType_OUTPUT,
			"CUSTOM": AssetHashType_CUSTOM,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.AssetManifestBuilder",
		reflect.TypeOf((*AssetManifestBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDockerImageAsset", GoMethod: "AddDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "defaultAddDockerImageAsset", GoMethod: "DefaultAddDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "defaultAddFileAsset", GoMethod: "DefaultAddFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "emitManifest", GoMethod: "EmitManifest"},
			_jsii_.MemberProperty{JsiiProperty: "hasAssets", GoGetter: "HasAssets"},
		},
		func() interface{} {
			return &jsiiProxy_AssetManifestBuilder{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.AssetManifestDockerImageDestination",
		reflect.TypeOf((*AssetManifestDockerImageDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.AssetManifestFileDestination",
		reflect.TypeOf((*AssetManifestFileDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.AssetOptions",
		reflect.TypeOf((*AssetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.AssetStaging",
		reflect.TypeOf((*AssetStaging)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absoluteStagedPath", GoGetter: "AbsoluteStagedPath"},
			_jsii_.MemberProperty{JsiiProperty: "assetHash", GoGetter: "AssetHash"},
			_jsii_.MemberProperty{JsiiProperty: "isArchive", GoGetter: "IsArchive"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "packaging", GoGetter: "Packaging"},
			_jsii_.MemberMethod{JsiiMethod: "relativeStagedPath", GoMethod: "RelativeStagedPath"},
			_jsii_.MemberProperty{JsiiProperty: "sourcePath", GoGetter: "SourcePath"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AssetStaging{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.AssetStagingProps",
		reflect.TypeOf((*AssetStagingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Aws",
		reflect.TypeOf((*Aws)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Aws{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.BootstraplessSynthesizer",
		reflect.TypeOf((*BootstraplessSynthesizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBootstrapVersionRule", GoMethod: "AddBootstrapVersionRule"},
			_jsii_.MemberMethod{JsiiMethod: "addDockerImageAsset", GoMethod: "AddDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapQualifier", GoGetter: "BootstrapQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "boundStack", GoGetter: "BoundStack"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFormationExecutionRoleArn", GoGetter: "CloudFormationExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromDockerImageAsset", GoMethod: "CloudFormationLocationFromDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromFileAsset", GoMethod: "CloudFormationLocationFromFileAsset"},
			_jsii_.MemberProperty{JsiiProperty: "deployRoleArn", GoGetter: "DeployRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "emitArtifact", GoMethod: "EmitArtifact"},
			_jsii_.MemberMethod{JsiiMethod: "emitStackArtifact", GoMethod: "EmitStackArtifact"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeStackTemplate", GoMethod: "SynthesizeStackTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeTemplate", GoMethod: "SynthesizeTemplate"},
		},
		func() interface{} {
			j := jsiiProxy_BootstraplessSynthesizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DefaultStackSynthesizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.BootstraplessSynthesizerProps",
		reflect.TypeOf((*BootstraplessSynthesizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.BundlingOptions",
		reflect.TypeOf((*BundlingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.BundlingOutput",
		reflect.TypeOf((*BundlingOutput)(nil)).Elem(),
		map[string]interface{}{
			"ARCHIVED": BundlingOutput_ARCHIVED,
			"NOT_ARCHIVED": BundlingOutput_NOT_ARCHIVED,
			"AUTO_DISCOVER": BundlingOutput_AUTO_DISCOVER,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnAutoScalingReplacingUpdate",
		reflect.TypeOf((*CfnAutoScalingReplacingUpdate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnAutoScalingRollingUpdate",
		reflect.TypeOf((*CfnAutoScalingRollingUpdate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnAutoScalingScheduledAction",
		reflect.TypeOf((*CfnAutoScalingScheduledAction)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.CfnCapabilities",
		reflect.TypeOf((*CfnCapabilities)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CfnCapabilities_NONE,
			"ANONYMOUS_IAM": CfnCapabilities_ANONYMOUS_IAM,
			"NAMED_IAM": CfnCapabilities_NAMED_IAM,
			"AUTO_EXPAND": CfnCapabilities_AUTO_EXPAND,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnCodeDeployBlueGreenAdditionalOptions",
		reflect.TypeOf((*CfnCodeDeployBlueGreenAdditionalOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnCodeDeployBlueGreenApplication",
		reflect.TypeOf((*CfnCodeDeployBlueGreenApplication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnCodeDeployBlueGreenApplicationTarget",
		reflect.TypeOf((*CfnCodeDeployBlueGreenApplicationTarget)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnCodeDeployBlueGreenEcsAttributes",
		reflect.TypeOf((*CfnCodeDeployBlueGreenEcsAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnCodeDeployBlueGreenHook",
		reflect.TypeOf((*CfnCodeDeployBlueGreenHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalOptions", GoGetter: "AdditionalOptions"},
			_jsii_.MemberProperty{JsiiProperty: "applications", GoGetter: "Applications"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleEventHooks", GoGetter: "LifecycleEventHooks"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trafficRoutingConfig", GoGetter: "TrafficRoutingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeDeployBlueGreenHook{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnHook)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnCodeDeployBlueGreenHookProps",
		reflect.TypeOf((*CfnCodeDeployBlueGreenHookProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnCodeDeployBlueGreenLifecycleEventHooks",
		reflect.TypeOf((*CfnCodeDeployBlueGreenLifecycleEventHooks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnCodeDeployLambdaAliasUpdate",
		reflect.TypeOf((*CfnCodeDeployLambdaAliasUpdate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnCondition",
		reflect.TypeOf((*CfnCondition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "expression", GoGetter: "Expression"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCondition{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnElement)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICfnConditionExpression)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResolvable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnConditionProps",
		reflect.TypeOf((*CfnConditionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnCreationPolicy",
		reflect.TypeOf((*CfnCreationPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnCustomResource",
		reflect.TypeOf((*CfnCustomResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
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
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "serviceToken", GoGetter: "ServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomResource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnCustomResourceProps",
		reflect.TypeOf((*CfnCustomResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.CfnDeletionPolicy",
		reflect.TypeOf((*CfnDeletionPolicy)(nil)).Elem(),
		map[string]interface{}{
			"DELETE": CfnDeletionPolicy_DELETE,
			"RETAIN": CfnDeletionPolicy_RETAIN,
			"SNAPSHOT": CfnDeletionPolicy_SNAPSHOT,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnDynamicReference",
		reflect.TypeOf((*CfnDynamicReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "newError", GoMethod: "NewError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toStringList", GoMethod: "ToStringList"},
			_jsii_.MemberProperty{JsiiProperty: "typeHint", GoGetter: "TypeHint"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDynamicReference{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Intrinsic)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnDynamicReferenceProps",
		reflect.TypeOf((*CfnDynamicReferenceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.CfnDynamicReferenceService",
		reflect.TypeOf((*CfnDynamicReferenceService)(nil)).Elem(),
		map[string]interface{}{
			"SSM": CfnDynamicReferenceService_SSM,
			"SSM_SECURE": CfnDynamicReferenceService_SSM_SECURE,
			"SECRETS_MANAGER": CfnDynamicReferenceService_SECRETS_MANAGER,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnElement",
		reflect.TypeOf((*CfnElement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CfnElement{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnHook",
		reflect.TypeOf((*CfnHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHook{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnElement)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnHookDefaultVersion",
		reflect.TypeOf((*CfnHookDefaultVersion)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "typeName", GoGetter: "TypeName"},
			_jsii_.MemberProperty{JsiiProperty: "typeVersionArn", GoGetter: "TypeVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "versionId", GoGetter: "VersionId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHookDefaultVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnHookDefaultVersionProps",
		reflect.TypeOf((*CfnHookDefaultVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnHookProps",
		reflect.TypeOf((*CfnHookProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnHookTypeConfig",
		reflect.TypeOf((*CfnHookTypeConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrConfigurationArn", GoGetter: "AttrConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "configurationAlias", GoGetter: "ConfigurationAlias"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "typeArn", GoGetter: "TypeArn"},
			_jsii_.MemberProperty{JsiiProperty: "typeName", GoGetter: "TypeName"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHookTypeConfig{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnHookTypeConfigProps",
		reflect.TypeOf((*CfnHookTypeConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnHookVersion",
		reflect.TypeOf((*CfnHookVersion)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrIsDefaultVersion", GoGetter: "AttrIsDefaultVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrTypeArn", GoGetter: "AttrTypeArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionId", GoGetter: "AttrVersionId"},
			_jsii_.MemberProperty{JsiiProperty: "attrVisibility", GoGetter: "AttrVisibility"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "schemaHandlerPackage", GoGetter: "SchemaHandlerPackage"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "typeName", GoGetter: "TypeName"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHookVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnHookVersion.LoggingConfigProperty",
		reflect.TypeOf((*CfnHookVersion_LoggingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnHookVersionProps",
		reflect.TypeOf((*CfnHookVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnJson",
		reflect.TypeOf((*CfnJson)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_CfnJson{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResolvable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnJsonProps",
		reflect.TypeOf((*CfnJsonProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnMacro",
		reflect.TypeOf((*CfnMacro)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
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
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "logRoleArn", GoGetter: "LogRoleArn"},
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
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMacro{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnMacroProps",
		reflect.TypeOf((*CfnMacroProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnMapping",
		reflect.TypeOf((*CfnMapping)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "findInMap", GoMethod: "FindInMap"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "setValue", GoMethod: "SetValue"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMapping{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnRefElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnMappingProps",
		reflect.TypeOf((*CfnMappingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnModuleDefaultVersion",
		reflect.TypeOf((*CfnModuleDefaultVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "moduleName", GoGetter: "ModuleName"},
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
			_jsii_.MemberProperty{JsiiProperty: "versionId", GoGetter: "VersionId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnModuleDefaultVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnModuleDefaultVersionProps",
		reflect.TypeOf((*CfnModuleDefaultVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnModuleVersion",
		reflect.TypeOf((*CfnModuleVersion)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrDescription", GoGetter: "AttrDescription"},
			_jsii_.MemberProperty{JsiiProperty: "attrDocumentationUrl", GoGetter: "AttrDocumentationUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attrIsDefaultVersion", GoGetter: "AttrIsDefaultVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrSchema", GoGetter: "AttrSchema"},
			_jsii_.MemberProperty{JsiiProperty: "attrTimeCreated", GoGetter: "AttrTimeCreated"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionId", GoGetter: "AttrVersionId"},
			_jsii_.MemberProperty{JsiiProperty: "attrVisibility", GoGetter: "AttrVisibility"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "moduleName", GoGetter: "ModuleName"},
			_jsii_.MemberProperty{JsiiProperty: "modulePackage", GoGetter: "ModulePackage"},
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
			j := jsiiProxy_CfnModuleVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnModuleVersionProps",
		reflect.TypeOf((*CfnModuleVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnOutput",
		reflect.TypeOf((*CfnOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "condition", GoGetter: "Condition"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "exportName", GoGetter: "ExportName"},
			_jsii_.MemberProperty{JsiiProperty: "importValue", GoGetter: "ImportValue"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOutput{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnOutputProps",
		reflect.TypeOf((*CfnOutputProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnParameter",
		reflect.TypeOf((*CfnParameter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedPattern", GoGetter: "AllowedPattern"},
			_jsii_.MemberProperty{JsiiProperty: "allowedValues", GoGetter: "AllowedValues"},
			_jsii_.MemberProperty{JsiiProperty: "constraintDescription", GoGetter: "ConstraintDescription"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "default", GoGetter: "Default"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maxLength", GoGetter: "MaxLength"},
			_jsii_.MemberProperty{JsiiProperty: "maxValue", GoGetter: "MaxValue"},
			_jsii_.MemberProperty{JsiiProperty: "minLength", GoGetter: "MinLength"},
			_jsii_.MemberProperty{JsiiProperty: "minValue", GoGetter: "MinValue"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "noEcho", GoGetter: "NoEcho"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueAsList", GoGetter: "ValueAsList"},
			_jsii_.MemberProperty{JsiiProperty: "valueAsNumber", GoGetter: "ValueAsNumber"},
			_jsii_.MemberProperty{JsiiProperty: "valueAsString", GoGetter: "ValueAsString"},
		},
		func() interface{} {
			j := jsiiProxy_CfnParameter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnParameterProps",
		reflect.TypeOf((*CfnParameterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnPublicTypeVersion",
		reflect.TypeOf((*CfnPublicTypeVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "attrPublicTypeArn", GoGetter: "AttrPublicTypeArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrPublisherId", GoGetter: "AttrPublisherId"},
			_jsii_.MemberProperty{JsiiProperty: "attrTypeVersionArn", GoGetter: "AttrTypeVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logDeliveryBucket", GoGetter: "LogDeliveryBucket"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "publicVersionNumber", GoGetter: "PublicVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeName", GoGetter: "TypeName"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPublicTypeVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnPublicTypeVersionProps",
		reflect.TypeOf((*CfnPublicTypeVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnPublisher",
		reflect.TypeOf((*CfnPublisher)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acceptTermsAndConditions", GoGetter: "AcceptTermsAndConditions"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrIdentityProvider", GoGetter: "AttrIdentityProvider"},
			_jsii_.MemberProperty{JsiiProperty: "attrPublisherId", GoGetter: "AttrPublisherId"},
			_jsii_.MemberProperty{JsiiProperty: "attrPublisherProfile", GoGetter: "AttrPublisherProfile"},
			_jsii_.MemberProperty{JsiiProperty: "attrPublisherStatus", GoGetter: "AttrPublisherStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "connectionArn", GoGetter: "ConnectionArn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			j := jsiiProxy_CfnPublisher{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnPublisherProps",
		reflect.TypeOf((*CfnPublisherProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnRefElement",
		reflect.TypeOf((*CfnRefElement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRefElement{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnElement)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnResource",
		reflect.TypeOf((*CfnResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
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
			j := jsiiProxy_CfnResource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnRefElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnResourceAutoScalingCreationPolicy",
		reflect.TypeOf((*CfnResourceAutoScalingCreationPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnResourceDefaultVersion",
		reflect.TypeOf((*CfnResourceDefaultVersion)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "typeName", GoGetter: "TypeName"},
			_jsii_.MemberProperty{JsiiProperty: "typeVersionArn", GoGetter: "TypeVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "versionId", GoGetter: "VersionId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourceDefaultVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnResourceDefaultVersionProps",
		reflect.TypeOf((*CfnResourceDefaultVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnResourceProps",
		reflect.TypeOf((*CfnResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnResourceSignal",
		reflect.TypeOf((*CfnResourceSignal)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnResourceVersion",
		reflect.TypeOf((*CfnResourceVersion)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrIsDefaultVersion", GoGetter: "AttrIsDefaultVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrProvisioningType", GoGetter: "AttrProvisioningType"},
			_jsii_.MemberProperty{JsiiProperty: "attrTypeArn", GoGetter: "AttrTypeArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionId", GoGetter: "AttrVersionId"},
			_jsii_.MemberProperty{JsiiProperty: "attrVisibility", GoGetter: "AttrVisibility"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "schemaHandlerPackage", GoGetter: "SchemaHandlerPackage"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "typeName", GoGetter: "TypeName"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourceVersion{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnResourceVersion.LoggingConfigProperty",
		reflect.TypeOf((*CfnResourceVersion_LoggingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnResourceVersionProps",
		reflect.TypeOf((*CfnResourceVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnRule",
		reflect.TypeOf((*CfnRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAssertion", GoMethod: "AddAssertion"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRule{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnRefElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnRuleAssertion",
		reflect.TypeOf((*CfnRuleAssertion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnRuleProps",
		reflect.TypeOf((*CfnRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnStack",
		reflect.TypeOf((*CfnStack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
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
			_jsii_.MemberProperty{JsiiProperty: "notificationArns", GoGetter: "NotificationArns"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateUrl", GoGetter: "TemplateUrl"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInMinutes", GoGetter: "TimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStack{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnStackProps",
		reflect.TypeOf((*CfnStackProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnStackSet",
		reflect.TypeOf((*CfnStackSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "administrationRoleArn", GoGetter: "AdministrationRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrStackSetId", GoGetter: "AttrStackSetId"},
			_jsii_.MemberProperty{JsiiProperty: "autoDeployment", GoGetter: "AutoDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "callAs", GoGetter: "CallAs"},
			_jsii_.MemberProperty{JsiiProperty: "capabilities", GoGetter: "Capabilities"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleName", GoGetter: "ExecutionRoleName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "managedExecution", GoGetter: "ManagedExecution"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "operationPreferences", GoGetter: "OperationPreferences"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "permissionModel", GoGetter: "PermissionModel"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stackInstancesGroup", GoGetter: "StackInstancesGroup"},
			_jsii_.MemberProperty{JsiiProperty: "stackSetName", GoGetter: "StackSetName"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateBody", GoGetter: "TemplateBody"},
			_jsii_.MemberProperty{JsiiProperty: "templateUrl", GoGetter: "TemplateUrl"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStackSet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnStackSet.AutoDeploymentProperty",
		reflect.TypeOf((*CfnStackSet_AutoDeploymentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnStackSet.DeploymentTargetsProperty",
		reflect.TypeOf((*CfnStackSet_DeploymentTargetsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnStackSet.ManagedExecutionProperty",
		reflect.TypeOf((*CfnStackSet_ManagedExecutionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnStackSet.OperationPreferencesProperty",
		reflect.TypeOf((*CfnStackSet_OperationPreferencesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnStackSet.ParameterProperty",
		reflect.TypeOf((*CfnStackSet_ParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnStackSet.StackInstancesProperty",
		reflect.TypeOf((*CfnStackSet_StackInstancesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnStackSetProps",
		reflect.TypeOf((*CfnStackSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnTag",
		reflect.TypeOf((*CfnTag)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnTrafficRoute",
		reflect.TypeOf((*CfnTrafficRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnTrafficRouting",
		reflect.TypeOf((*CfnTrafficRouting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnTrafficRoutingConfig",
		reflect.TypeOf((*CfnTrafficRoutingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnTrafficRoutingTimeBasedCanary",
		reflect.TypeOf((*CfnTrafficRoutingTimeBasedCanary)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnTrafficRoutingTimeBasedLinear",
		reflect.TypeOf((*CfnTrafficRoutingTimeBasedLinear)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.CfnTrafficRoutingType",
		reflect.TypeOf((*CfnTrafficRoutingType)(nil)).Elem(),
		map[string]interface{}{
			"ALL_AT_ONCE": CfnTrafficRoutingType_ALL_AT_ONCE,
			"TIME_BASED_CANARY": CfnTrafficRoutingType_TIME_BASED_CANARY,
			"TIME_BASED_LINEAR": CfnTrafficRoutingType_TIME_BASED_LINEAR,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnTypeActivation",
		reflect.TypeOf((*CfnTypeActivation)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "autoUpdate", GoGetter: "AutoUpdate"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "majorVersion", GoGetter: "MajorVersion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "publicTypeArn", GoGetter: "PublicTypeArn"},
			_jsii_.MemberProperty{JsiiProperty: "publisherId", GoGetter: "PublisherId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeName", GoGetter: "TypeName"},
			_jsii_.MemberProperty{JsiiProperty: "typeNameAlias", GoGetter: "TypeNameAlias"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "versionBump", GoGetter: "VersionBump"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTypeActivation{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnTypeActivation.LoggingConfigProperty",
		reflect.TypeOf((*CfnTypeActivation_LoggingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnTypeActivationProps",
		reflect.TypeOf((*CfnTypeActivationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnUpdatePolicy",
		reflect.TypeOf((*CfnUpdatePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnWaitCondition",
		reflect.TypeOf((*CfnWaitCondition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrData", GoGetter: "AttrData"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "handle", GoGetter: "Handle"},
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
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWaitCondition{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CfnWaitConditionHandle",
		reflect.TypeOf((*CfnWaitConditionHandle)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
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
			j := jsiiProxy_CfnWaitConditionHandle{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CfnWaitConditionProps",
		reflect.TypeOf((*CfnWaitConditionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CliCredentialsStackSynthesizer",
		reflect.TypeOf((*CliCredentialsStackSynthesizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBootstrapVersionRule", GoMethod: "AddBootstrapVersionRule"},
			_jsii_.MemberMethod{JsiiMethod: "addDockerImageAsset", GoMethod: "AddDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapQualifier", GoGetter: "BootstrapQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "boundStack", GoGetter: "BoundStack"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromDockerImageAsset", GoMethod: "CloudFormationLocationFromDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromFileAsset", GoMethod: "CloudFormationLocationFromFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "emitArtifact", GoMethod: "EmitArtifact"},
			_jsii_.MemberMethod{JsiiMethod: "emitStackArtifact", GoMethod: "EmitStackArtifact"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeStackTemplate", GoMethod: "SynthesizeStackTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeTemplate", GoMethod: "SynthesizeTemplate"},
		},
		func() interface{} {
			j := jsiiProxy_CliCredentialsStackSynthesizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_StackSynthesizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CliCredentialsStackSynthesizerProps",
		reflect.TypeOf((*CliCredentialsStackSynthesizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.ContextProvider",
		reflect.TypeOf((*ContextProvider)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ContextProvider{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CopyOptions",
		reflect.TypeOf((*CopyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CustomResource",
		reflect.TypeOf((*CustomResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getAttString", GoMethod: "GetAttString"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CustomResource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CustomResourceProps",
		reflect.TypeOf((*CustomResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.CustomResourceProvider",
		reflect.TypeOf((*CustomResourceProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "codeHash", GoGetter: "CodeHash"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "serviceToken", GoGetter: "ServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CustomResourceProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.CustomResourceProviderProps",
		reflect.TypeOf((*CustomResourceProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.CustomResourceProviderRuntime",
		reflect.TypeOf((*CustomResourceProviderRuntime)(nil)).Elem(),
		map[string]interface{}{
			"NODEJS_12_X": CustomResourceProviderRuntime_NODEJS_12_X,
			"NODEJS_14_X": CustomResourceProviderRuntime_NODEJS_14_X,
			"NODEJS_16_X": CustomResourceProviderRuntime_NODEJS_16_X,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.DefaultStackSynthesizer",
		reflect.TypeOf((*DefaultStackSynthesizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBootstrapVersionRule", GoMethod: "AddBootstrapVersionRule"},
			_jsii_.MemberMethod{JsiiMethod: "addDockerImageAsset", GoMethod: "AddDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapQualifier", GoGetter: "BootstrapQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "boundStack", GoGetter: "BoundStack"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFormationExecutionRoleArn", GoGetter: "CloudFormationExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromDockerImageAsset", GoMethod: "CloudFormationLocationFromDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromFileAsset", GoMethod: "CloudFormationLocationFromFileAsset"},
			_jsii_.MemberProperty{JsiiProperty: "deployRoleArn", GoGetter: "DeployRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "emitArtifact", GoMethod: "EmitArtifact"},
			_jsii_.MemberMethod{JsiiMethod: "emitStackArtifact", GoMethod: "EmitStackArtifact"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeStackTemplate", GoMethod: "SynthesizeStackTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeTemplate", GoMethod: "SynthesizeTemplate"},
		},
		func() interface{} {
			j := jsiiProxy_DefaultStackSynthesizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_StackSynthesizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.DefaultStackSynthesizerProps",
		reflect.TypeOf((*DefaultStackSynthesizerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.DefaultTokenResolver",
		reflect.TypeOf((*DefaultTokenResolver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "resolveList", GoMethod: "ResolveList"},
			_jsii_.MemberMethod{JsiiMethod: "resolveString", GoMethod: "ResolveString"},
			_jsii_.MemberMethod{JsiiMethod: "resolveToken", GoMethod: "ResolveToken"},
		},
		func() interface{} {
			j := jsiiProxy_DefaultTokenResolver{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITokenResolver)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.DockerBuildOptions",
		reflect.TypeOf((*DockerBuildOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.DockerIgnoreStrategy",
		reflect.TypeOf((*DockerIgnoreStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "add", GoMethod: "Add"},
			_jsii_.MemberMethod{JsiiMethod: "ignores", GoMethod: "Ignores"},
		},
		func() interface{} {
			j := jsiiProxy_DockerIgnoreStrategy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IgnoreStrategy)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.DockerImage",
		reflect.TypeOf((*DockerImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "cp", GoMethod: "Cp"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberMethod{JsiiMethod: "run", GoMethod: "Run"},
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
		},
		func() interface{} {
			return &jsiiProxy_DockerImage{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.DockerImageAssetLocation",
		reflect.TypeOf((*DockerImageAssetLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.DockerImageAssetSource",
		reflect.TypeOf((*DockerImageAssetSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.DockerRunOptions",
		reflect.TypeOf((*DockerRunOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.DockerVolume",
		reflect.TypeOf((*DockerVolume)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.DockerVolumeConsistency",
		reflect.TypeOf((*DockerVolumeConsistency)(nil)).Elem(),
		map[string]interface{}{
			"CONSISTENT": DockerVolumeConsistency_CONSISTENT,
			"DELEGATED": DockerVolumeConsistency_DELEGATED,
			"CACHED": DockerVolumeConsistency_CACHED,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Duration",
		reflect.TypeOf((*Duration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "formatTokenToNumber", GoMethod: "FormatTokenToNumber"},
			_jsii_.MemberMethod{JsiiMethod: "isUnresolved", GoMethod: "IsUnresolved"},
			_jsii_.MemberMethod{JsiiMethod: "minus", GoMethod: "Minus"},
			_jsii_.MemberMethod{JsiiMethod: "plus", GoMethod: "Plus"},
			_jsii_.MemberMethod{JsiiMethod: "toDays", GoMethod: "ToDays"},
			_jsii_.MemberMethod{JsiiMethod: "toHours", GoMethod: "ToHours"},
			_jsii_.MemberMethod{JsiiMethod: "toHumanString", GoMethod: "ToHumanString"},
			_jsii_.MemberMethod{JsiiMethod: "toIsoString", GoMethod: "ToIsoString"},
			_jsii_.MemberMethod{JsiiMethod: "toMilliseconds", GoMethod: "ToMilliseconds"},
			_jsii_.MemberMethod{JsiiMethod: "toMinutes", GoMethod: "ToMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "toSeconds", GoMethod: "ToSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "unitLabel", GoMethod: "UnitLabel"},
		},
		func() interface{} {
			return &jsiiProxy_Duration{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.EncodingOptions",
		reflect.TypeOf((*EncodingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.Environment",
		reflect.TypeOf((*Environment)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Expiration",
		reflect.TypeOf((*Expiration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "date", GoGetter: "Date"},
			_jsii_.MemberMethod{JsiiMethod: "isAfter", GoMethod: "IsAfter"},
			_jsii_.MemberMethod{JsiiMethod: "isBefore", GoMethod: "IsBefore"},
			_jsii_.MemberMethod{JsiiMethod: "toEpoch", GoMethod: "ToEpoch"},
		},
		func() interface{} {
			return &jsiiProxy_Expiration{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.ExportValueOptions",
		reflect.TypeOf((*ExportValueOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.FeatureFlags",
		reflect.TypeOf((*FeatureFlags)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "isEnabled", GoMethod: "IsEnabled"},
		},
		func() interface{} {
			return &jsiiProxy_FeatureFlags{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.FileAssetLocation",
		reflect.TypeOf((*FileAssetLocation)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.FileAssetPackaging",
		reflect.TypeOf((*FileAssetPackaging)(nil)).Elem(),
		map[string]interface{}{
			"ZIP_DIRECTORY": FileAssetPackaging_ZIP_DIRECTORY,
			"FILE": FileAssetPackaging_FILE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.FileAssetSource",
		reflect.TypeOf((*FileAssetSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.FileCopyOptions",
		reflect.TypeOf((*FileCopyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.FileFingerprintOptions",
		reflect.TypeOf((*FileFingerprintOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.FileSystem",
		reflect.TypeOf((*FileSystem)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_FileSystem{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.FingerprintOptions",
		reflect.TypeOf((*FingerprintOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Fn",
		reflect.TypeOf((*Fn)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Fn{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.GetContextKeyOptions",
		reflect.TypeOf((*GetContextKeyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.GetContextKeyResult",
		reflect.TypeOf((*GetContextKeyResult)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.GetContextValueOptions",
		reflect.TypeOf((*GetContextValueOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.GetContextValueResult",
		reflect.TypeOf((*GetContextValueResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.GitIgnoreStrategy",
		reflect.TypeOf((*GitIgnoreStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "add", GoMethod: "Add"},
			_jsii_.MemberMethod{JsiiMethod: "ignores", GoMethod: "Ignores"},
		},
		func() interface{} {
			j := jsiiProxy_GitIgnoreStrategy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IgnoreStrategy)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.GlobIgnoreStrategy",
		reflect.TypeOf((*GlobIgnoreStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "add", GoMethod: "Add"},
			_jsii_.MemberMethod{JsiiMethod: "ignores", GoMethod: "Ignores"},
		},
		func() interface{} {
			j := jsiiProxy_GlobIgnoreStrategy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IgnoreStrategy)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IAnyProducer",
		reflect.TypeOf((*IAnyProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
		},
		func() interface{} {
			return &jsiiProxy_IAnyProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IAspect",
		reflect.TypeOf((*IAspect)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			return &jsiiProxy_IAspect{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IAsset",
		reflect.TypeOf((*IAsset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assetHash", GoGetter: "AssetHash"},
		},
		func() interface{} {
			return &jsiiProxy_IAsset{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.ICfnConditionExpression",
		reflect.TypeOf((*ICfnConditionExpression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "typeHint", GoGetter: "TypeHint"},
		},
		func() interface{} {
			j := jsiiProxy_ICfnConditionExpression{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResolvable)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.ICfnResourceOptions",
		reflect.TypeOf((*ICfnResourceOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "condition", GoGetter: "Condition"},
			_jsii_.MemberProperty{JsiiProperty: "creationPolicy", GoGetter: "CreationPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deletionPolicy", GoGetter: "DeletionPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "updatePolicy", GoGetter: "UpdatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "updateReplacePolicy", GoGetter: "UpdateReplacePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_ICfnResourceOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.ICfnRuleConditionExpression",
		reflect.TypeOf((*ICfnRuleConditionExpression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disambiguator", GoGetter: "Disambiguator"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "typeHint", GoGetter: "TypeHint"},
		},
		func() interface{} {
			j := jsiiProxy_ICfnRuleConditionExpression{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICfnConditionExpression)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IFragmentConcatenator",
		reflect.TypeOf((*IFragmentConcatenator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "join", GoMethod: "Join"},
		},
		func() interface{} {
			return &jsiiProxy_IFragmentConcatenator{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IInspectable",
		reflect.TypeOf((*IInspectable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
		},
		func() interface{} {
			return &jsiiProxy_IInspectable{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IListProducer",
		reflect.TypeOf((*IListProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
		},
		func() interface{} {
			return &jsiiProxy_IListProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.ILocalBundling",
		reflect.TypeOf((*ILocalBundling)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "tryBundle", GoMethod: "TryBundle"},
		},
		func() interface{} {
			return &jsiiProxy_ILocalBundling{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.INumberProducer",
		reflect.TypeOf((*INumberProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
		},
		func() interface{} {
			return &jsiiProxy_INumberProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IPostProcessor",
		reflect.TypeOf((*IPostProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postProcess", GoMethod: "PostProcess"},
		},
		func() interface{} {
			return &jsiiProxy_IPostProcessor{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IResolvable",
		reflect.TypeOf((*IResolvable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "typeHint", GoGetter: "TypeHint"},
		},
		func() interface{} {
			return &jsiiProxy_IResolvable{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IResolveContext",
		reflect.TypeOf((*IResolveContext)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "documentPath", GoGetter: "DocumentPath"},
			_jsii_.MemberProperty{JsiiProperty: "preparing", GoGetter: "Preparing"},
			_jsii_.MemberMethod{JsiiMethod: "registerPostProcessor", GoMethod: "RegisterPostProcessor"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
		},
		func() interface{} {
			return &jsiiProxy_IResolveContext{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IResource",
		reflect.TypeOf((*IResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IResource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IStableAnyProducer",
		reflect.TypeOf((*IStableAnyProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
		},
		func() interface{} {
			return &jsiiProxy_IStableAnyProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IStableListProducer",
		reflect.TypeOf((*IStableListProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
		},
		func() interface{} {
			return &jsiiProxy_IStableListProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IStableNumberProducer",
		reflect.TypeOf((*IStableNumberProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
		},
		func() interface{} {
			return &jsiiProxy_IStableNumberProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IStableStringProducer",
		reflect.TypeOf((*IStableStringProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
		},
		func() interface{} {
			return &jsiiProxy_IStableStringProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IStackSynthesizer",
		reflect.TypeOf((*IStackSynthesizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDockerImageAsset", GoMethod: "AddDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapQualifier", GoGetter: "BootstrapQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			return &jsiiProxy_IStackSynthesizer{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.IStringProducer",
		reflect.TypeOf((*IStringProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
		},
		func() interface{} {
			return &jsiiProxy_IStringProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.ISynthesisSession",
		reflect.TypeOf((*ISynthesisSession)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assembly", GoGetter: "Assembly"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "validateOnSynth", GoGetter: "ValidateOnSynth"},
		},
		func() interface{} {
			return &jsiiProxy_ISynthesisSession{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.ITaggable",
		reflect.TypeOf((*ITaggable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
		},
		func() interface{} {
			return &jsiiProxy_ITaggable{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.ITemplateOptions",
		reflect.TypeOf((*ITemplateOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "templateFormatVersion", GoGetter: "TemplateFormatVersion"},
			_jsii_.MemberProperty{JsiiProperty: "transforms", GoGetter: "Transforms"},
		},
		func() interface{} {
			return &jsiiProxy_ITemplateOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.ITokenMapper",
		reflect.TypeOf((*ITokenMapper)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "mapToken", GoMethod: "MapToken"},
		},
		func() interface{} {
			return &jsiiProxy_ITokenMapper{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.ITokenResolver",
		reflect.TypeOf((*ITokenResolver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "resolveList", GoMethod: "ResolveList"},
			_jsii_.MemberMethod{JsiiMethod: "resolveString", GoMethod: "ResolveString"},
			_jsii_.MemberMethod{JsiiMethod: "resolveToken", GoMethod: "ResolveToken"},
		},
		func() interface{} {
			return &jsiiProxy_ITokenResolver{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.IgnoreMode",
		reflect.TypeOf((*IgnoreMode)(nil)).Elem(),
		map[string]interface{}{
			"GLOB": IgnoreMode_GLOB,
			"GIT": IgnoreMode_GIT,
			"DOCKER": IgnoreMode_DOCKER,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.IgnoreStrategy",
		reflect.TypeOf((*IgnoreStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "add", GoMethod: "Add"},
			_jsii_.MemberMethod{JsiiMethod: "ignores", GoMethod: "Ignores"},
		},
		func() interface{} {
			return &jsiiProxy_IgnoreStrategy{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Intrinsic",
		reflect.TypeOf((*Intrinsic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "newError", GoMethod: "NewError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toStringList", GoMethod: "ToStringList"},
			_jsii_.MemberProperty{JsiiProperty: "typeHint", GoGetter: "TypeHint"},
		},
		func() interface{} {
			j := jsiiProxy_Intrinsic{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResolvable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.IntrinsicProps",
		reflect.TypeOf((*IntrinsicProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Lazy",
		reflect.TypeOf((*Lazy)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Lazy{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.LazyAnyValueOptions",
		reflect.TypeOf((*LazyAnyValueOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.LazyListValueOptions",
		reflect.TypeOf((*LazyListValueOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.LazyStringValueOptions",
		reflect.TypeOf((*LazyStringValueOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.LegacyStackSynthesizer",
		reflect.TypeOf((*LegacyStackSynthesizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBootstrapVersionRule", GoMethod: "AddBootstrapVersionRule"},
			_jsii_.MemberMethod{JsiiMethod: "addDockerImageAsset", GoMethod: "AddDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapQualifier", GoGetter: "BootstrapQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "boundStack", GoGetter: "BoundStack"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromDockerImageAsset", GoMethod: "CloudFormationLocationFromDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromFileAsset", GoMethod: "CloudFormationLocationFromFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "emitArtifact", GoMethod: "EmitArtifact"},
			_jsii_.MemberMethod{JsiiMethod: "emitStackArtifact", GoMethod: "EmitStackArtifact"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeStackTemplate", GoMethod: "SynthesizeStackTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeTemplate", GoMethod: "SynthesizeTemplate"},
		},
		func() interface{} {
			j := jsiiProxy_LegacyStackSynthesizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_StackSynthesizer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Names",
		reflect.TypeOf((*Names)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Names{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.NestedStack",
		reflect.TypeOf((*NestedStack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addTransform", GoMethod: "AddTransform"},
			_jsii_.MemberMethod{JsiiMethod: "allocateLogicalId", GoMethod: "AllocateLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "bundlingRequired", GoGetter: "BundlingRequired"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberMethod{JsiiMethod: "exportStringListValue", GoMethod: "ExportStringListValue"},
			_jsii_.MemberMethod{JsiiMethod: "exportValue", GoMethod: "ExportValue"},
			_jsii_.MemberMethod{JsiiMethod: "formatArn", GoMethod: "FormatArn"},
			_jsii_.MemberMethod{JsiiMethod: "getLogicalId", GoMethod: "GetLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "nested", GoGetter: "Nested"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackParent", GoGetter: "NestedStackParent"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackResource", GoGetter: "NestedStackResource"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationArns", GoGetter: "NotificationArns"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "regionalFact", GoMethod: "RegionalFact"},
			_jsii_.MemberMethod{JsiiMethod: "renameLogicalId", GoMethod: "RenameLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "reportMissingContextKey", GoMethod: "ReportMissingContextKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "setParameter", GoMethod: "SetParameter"},
			_jsii_.MemberMethod{JsiiMethod: "splitArn", GoMethod: "SplitArn"},
			_jsii_.MemberProperty{JsiiProperty: "stackId", GoGetter: "StackId"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberProperty{JsiiProperty: "synthesizer", GoGetter: "Synthesizer"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateFile", GoGetter: "TemplateFile"},
			_jsii_.MemberProperty{JsiiProperty: "templateOptions", GoGetter: "TemplateOptions"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtection", GoGetter: "TerminationProtection"},
			_jsii_.MemberMethod{JsiiMethod: "toJsonString", GoMethod: "ToJsonString"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlSuffix", GoGetter: "UrlSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_NestedStack{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Stack)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.NestedStackProps",
		reflect.TypeOf((*NestedStackProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.NestedStackSynthesizer",
		reflect.TypeOf((*NestedStackSynthesizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBootstrapVersionRule", GoMethod: "AddBootstrapVersionRule"},
			_jsii_.MemberMethod{JsiiMethod: "addDockerImageAsset", GoMethod: "AddDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapQualifier", GoGetter: "BootstrapQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "boundStack", GoGetter: "BoundStack"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromDockerImageAsset", GoMethod: "CloudFormationLocationFromDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromFileAsset", GoMethod: "CloudFormationLocationFromFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "emitArtifact", GoMethod: "EmitArtifact"},
			_jsii_.MemberMethod{JsiiMethod: "emitStackArtifact", GoMethod: "EmitStackArtifact"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeStackTemplate", GoMethod: "SynthesizeStackTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeTemplate", GoMethod: "SynthesizeTemplate"},
		},
		func() interface{} {
			j := jsiiProxy_NestedStackSynthesizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_StackSynthesizer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.PermissionsBoundary",
		reflect.TypeOf((*PermissionsBoundary)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_PermissionsBoundary{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.PermissionsBoundaryBindOptions",
		reflect.TypeOf((*PermissionsBoundaryBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.PhysicalName",
		reflect.TypeOf((*PhysicalName)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_PhysicalName{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Reference",
		reflect.TypeOf((*Reference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "newError", GoMethod: "NewError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toStringList", GoMethod: "ToStringList"},
			_jsii_.MemberProperty{JsiiProperty: "typeHint", GoGetter: "TypeHint"},
		},
		func() interface{} {
			j := jsiiProxy_Reference{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Intrinsic)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.RemovalPolicy",
		reflect.TypeOf((*RemovalPolicy)(nil)).Elem(),
		map[string]interface{}{
			"DESTROY": RemovalPolicy_DESTROY,
			"RETAIN": RemovalPolicy_RETAIN,
			"SNAPSHOT": RemovalPolicy_SNAPSHOT,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.RemovalPolicyOptions",
		reflect.TypeOf((*RemovalPolicyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.RemoveTag",
		reflect.TypeOf((*RemoveTag)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTag", GoMethod: "ApplyTag"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_RemoveTag{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAspect)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.ResolutionTypeHint",
		reflect.TypeOf((*ResolutionTypeHint)(nil)).Elem(),
		map[string]interface{}{
			"STRING": ResolutionTypeHint_STRING,
			"NUMBER": ResolutionTypeHint_NUMBER,
			"STRING_LIST": ResolutionTypeHint_STRING_LIST,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.ResolveChangeContextOptions",
		reflect.TypeOf((*ResolveChangeContextOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.ResolveOptions",
		reflect.TypeOf((*ResolveOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Resource",
		reflect.TypeOf((*Resource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Resource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.ResourceEnvironment",
		reflect.TypeOf((*ResourceEnvironment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.ResourceProps",
		reflect.TypeOf((*ResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.ReverseOptions",
		reflect.TypeOf((*ReverseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.RoleOptions",
		reflect.TypeOf((*RoleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.ScopedAws",
		reflect.TypeOf((*ScopedAws)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "notificationArns", GoGetter: "NotificationArns"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "stackId", GoGetter: "StackId"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberProperty{JsiiProperty: "urlSuffix", GoGetter: "UrlSuffix"},
		},
		func() interface{} {
			return &jsiiProxy_ScopedAws{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.SecretValue",
		reflect.TypeOf((*SecretValue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "newError", GoMethod: "NewError"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toStringList", GoMethod: "ToStringList"},
			_jsii_.MemberProperty{JsiiProperty: "typeHint", GoGetter: "TypeHint"},
			_jsii_.MemberMethod{JsiiMethod: "unsafeUnwrap", GoMethod: "UnsafeUnwrap"},
		},
		func() interface{} {
			j := jsiiProxy_SecretValue{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Intrinsic)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.SecretsManagerSecretOptions",
		reflect.TypeOf((*SecretsManagerSecretOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Size",
		reflect.TypeOf((*Size)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "isUnresolved", GoMethod: "IsUnresolved"},
			_jsii_.MemberMethod{JsiiMethod: "toGibibytes", GoMethod: "ToGibibytes"},
			_jsii_.MemberMethod{JsiiMethod: "toKibibytes", GoMethod: "ToKibibytes"},
			_jsii_.MemberMethod{JsiiMethod: "toMebibytes", GoMethod: "ToMebibytes"},
			_jsii_.MemberMethod{JsiiMethod: "toPebibytes", GoMethod: "ToPebibytes"},
			_jsii_.MemberMethod{JsiiMethod: "toTebibytes", GoMethod: "ToTebibytes"},
		},
		func() interface{} {
			return &jsiiProxy_Size{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.SizeConversionOptions",
		reflect.TypeOf((*SizeConversionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.SizeRoundingBehavior",
		reflect.TypeOf((*SizeRoundingBehavior)(nil)).Elem(),
		map[string]interface{}{
			"FAIL": SizeRoundingBehavior_FAIL,
			"FLOOR": SizeRoundingBehavior_FLOOR,
			"NONE": SizeRoundingBehavior_NONE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Stack",
		reflect.TypeOf((*Stack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addTransform", GoMethod: "AddTransform"},
			_jsii_.MemberMethod{JsiiMethod: "allocateLogicalId", GoMethod: "AllocateLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "bundlingRequired", GoGetter: "BundlingRequired"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberMethod{JsiiMethod: "exportStringListValue", GoMethod: "ExportStringListValue"},
			_jsii_.MemberMethod{JsiiMethod: "exportValue", GoMethod: "ExportValue"},
			_jsii_.MemberMethod{JsiiMethod: "formatArn", GoMethod: "FormatArn"},
			_jsii_.MemberMethod{JsiiMethod: "getLogicalId", GoMethod: "GetLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "nested", GoGetter: "Nested"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackParent", GoGetter: "NestedStackParent"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackResource", GoGetter: "NestedStackResource"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationArns", GoGetter: "NotificationArns"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "regionalFact", GoMethod: "RegionalFact"},
			_jsii_.MemberMethod{JsiiMethod: "renameLogicalId", GoMethod: "RenameLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "reportMissingContextKey", GoMethod: "ReportMissingContextKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "splitArn", GoMethod: "SplitArn"},
			_jsii_.MemberProperty{JsiiProperty: "stackId", GoGetter: "StackId"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberProperty{JsiiProperty: "synthesizer", GoGetter: "Synthesizer"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateFile", GoGetter: "TemplateFile"},
			_jsii_.MemberProperty{JsiiProperty: "templateOptions", GoGetter: "TemplateOptions"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtection", GoGetter: "TerminationProtection"},
			_jsii_.MemberMethod{JsiiMethod: "toJsonString", GoMethod: "ToJsonString"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlSuffix", GoGetter: "UrlSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_Stack{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.StackProps",
		reflect.TypeOf((*StackProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.StackSynthesizer",
		reflect.TypeOf((*StackSynthesizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBootstrapVersionRule", GoMethod: "AddBootstrapVersionRule"},
			_jsii_.MemberMethod{JsiiMethod: "addDockerImageAsset", GoMethod: "AddDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapQualifier", GoGetter: "BootstrapQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "boundStack", GoGetter: "BoundStack"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromDockerImageAsset", GoMethod: "CloudFormationLocationFromDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromFileAsset", GoMethod: "CloudFormationLocationFromFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "emitArtifact", GoMethod: "EmitArtifact"},
			_jsii_.MemberMethod{JsiiMethod: "emitStackArtifact", GoMethod: "EmitStackArtifact"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeStackTemplate", GoMethod: "SynthesizeStackTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeTemplate", GoMethod: "SynthesizeTemplate"},
		},
		func() interface{} {
			j := jsiiProxy_StackSynthesizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStackSynthesizer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Stage",
		reflect.TypeOf((*Stage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "assetOutdir", GoGetter: "AssetOutdir"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "parentStage", GoGetter: "ParentStage"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Stage{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.StageProps",
		reflect.TypeOf((*StageProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.StageSynthesisOptions",
		reflect.TypeOf((*StageSynthesisOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.StringConcat",
		reflect.TypeOf((*StringConcat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "join", GoMethod: "Join"},
		},
		func() interface{} {
			j := jsiiProxy_StringConcat{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFragmentConcatenator)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.SymlinkFollowMode",
		reflect.TypeOf((*SymlinkFollowMode)(nil)).Elem(),
		map[string]interface{}{
			"NEVER": SymlinkFollowMode_NEVER,
			"ALWAYS": SymlinkFollowMode_ALWAYS,
			"EXTERNAL": SymlinkFollowMode_EXTERNAL,
			"BLOCK_EXTERNAL": SymlinkFollowMode_BLOCK_EXTERNAL,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.SynthesizeStackArtifactOptions",
		reflect.TypeOf((*SynthesizeStackArtifactOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Tag",
		reflect.TypeOf((*Tag)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTag", GoMethod: "ApplyTag"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_Tag{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAspect)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.TagManager",
		reflect.TypeOf((*TagManager)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTagAspectHere", GoMethod: "ApplyTagAspectHere"},
			_jsii_.MemberMethod{JsiiMethod: "hasTags", GoMethod: "HasTags"},
			_jsii_.MemberMethod{JsiiMethod: "removeTag", GoMethod: "RemoveTag"},
			_jsii_.MemberProperty{JsiiProperty: "renderedTags", GoGetter: "RenderedTags"},
			_jsii_.MemberMethod{JsiiMethod: "renderTags", GoMethod: "RenderTags"},
			_jsii_.MemberMethod{JsiiMethod: "setTag", GoMethod: "SetTag"},
			_jsii_.MemberProperty{JsiiProperty: "tagPropertyName", GoGetter: "TagPropertyName"},
			_jsii_.MemberMethod{JsiiMethod: "tagValues", GoMethod: "TagValues"},
		},
		func() interface{} {
			return &jsiiProxy_TagManager{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.TagManagerOptions",
		reflect.TypeOf((*TagManagerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.TagProps",
		reflect.TypeOf((*TagProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.TagType",
		reflect.TypeOf((*TagType)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": TagType_STANDARD,
			"AUTOSCALING_GROUP": TagType_AUTOSCALING_GROUP,
			"MAP": TagType_MAP,
			"KEY_VALUE": TagType_KEY_VALUE,
			"NOT_TAGGABLE": TagType_NOT_TAGGABLE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Tags",
		reflect.TypeOf((*Tags)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "add", GoMethod: "Add"},
			_jsii_.MemberMethod{JsiiMethod: "remove", GoMethod: "Remove"},
		},
		func() interface{} {
			return &jsiiProxy_Tags{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.TimeConversionOptions",
		reflect.TypeOf((*TimeConversionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Token",
		reflect.TypeOf((*Token)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Token{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.TokenComparison",
		reflect.TypeOf((*TokenComparison)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_TokenComparison{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.Tokenization",
		reflect.TypeOf((*Tokenization)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Tokenization{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.TokenizedStringFragments",
		reflect.TypeOf((*TokenizedStringFragments)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIntrinsic", GoMethod: "AddIntrinsic"},
			_jsii_.MemberMethod{JsiiMethod: "addLiteral", GoMethod: "AddLiteral"},
			_jsii_.MemberMethod{JsiiMethod: "addToken", GoMethod: "AddToken"},
			_jsii_.MemberProperty{JsiiProperty: "firstToken", GoGetter: "FirstToken"},
			_jsii_.MemberProperty{JsiiProperty: "firstValue", GoGetter: "FirstValue"},
			_jsii_.MemberMethod{JsiiMethod: "join", GoMethod: "Join"},
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberMethod{JsiiMethod: "mapTokens", GoMethod: "MapTokens"},
			_jsii_.MemberProperty{JsiiProperty: "tokens", GoGetter: "Tokens"},
		},
		func() interface{} {
			return &jsiiProxy_TokenizedStringFragments{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.TreeInspector",
		reflect.TypeOf((*TreeInspector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAttribute", GoMethod: "AddAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "attributes", GoGetter: "Attributes"},
		},
		func() interface{} {
			return &jsiiProxy_TreeInspector{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.UniqueResourceNameOptions",
		reflect.TypeOf((*UniqueResourceNameOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.ValidationResult",
		reflect.TypeOf((*ValidationResult)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assertSuccess", GoMethod: "AssertSuccess"},
			_jsii_.MemberProperty{JsiiProperty: "errorMessage", GoGetter: "ErrorMessage"},
			_jsii_.MemberMethod{JsiiMethod: "errorTree", GoMethod: "ErrorTree"},
			_jsii_.MemberProperty{JsiiProperty: "isSuccess", GoGetter: "IsSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "prefix", GoMethod: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "results", GoGetter: "Results"},
		},
		func() interface{} {
			return &jsiiProxy_ValidationResult{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.ValidationResults",
		reflect.TypeOf((*ValidationResults)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "collect", GoMethod: "Collect"},
			_jsii_.MemberMethod{JsiiMethod: "errorTreeList", GoMethod: "ErrorTreeList"},
			_jsii_.MemberProperty{JsiiProperty: "isSuccess", GoGetter: "IsSuccess"},
			_jsii_.MemberProperty{JsiiProperty: "results", GoGetter: "Results"},
			_jsii_.MemberMethod{JsiiMethod: "wrap", GoMethod: "Wrap"},
		},
		func() interface{} {
			return &jsiiProxy_ValidationResults{}
		},
	)
}
