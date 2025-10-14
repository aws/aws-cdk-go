package awsappconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.Action",
		reflect.TypeOf((*Action)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionPoints", GoGetter: "ActionPoints"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "eventDestination", GoGetter: "EventDestination"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "invokeWithoutExecutionRole", GoGetter: "InvokeWithoutExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_Action{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appconfig.ActionPoint",
		reflect.TypeOf((*ActionPoint)(nil)).Elem(),
		map[string]interface{}{
			"PRE_CREATE_HOSTED_CONFIGURATION_VERSION": ActionPoint_PRE_CREATE_HOSTED_CONFIGURATION_VERSION,
			"PRE_START_DEPLOYMENT": ActionPoint_PRE_START_DEPLOYMENT,
			"ON_DEPLOYMENT_START": ActionPoint_ON_DEPLOYMENT_START,
			"ON_DEPLOYMENT_STEP": ActionPoint_ON_DEPLOYMENT_STEP,
			"ON_DEPLOYMENT_BAKING": ActionPoint_ON_DEPLOYMENT_BAKING,
			"ON_DEPLOYMENT_COMPLETE": ActionPoint_ON_DEPLOYMENT_COMPLETE,
			"ON_DEPLOYMENT_ROLLED_BACK": ActionPoint_ON_DEPLOYMENT_ROLLED_BACK,
			"AT_DEPLOYMENT_TICK": ActionPoint_AT_DEPLOYMENT_TICK,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.ActionProps",
		reflect.TypeOf((*ActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.Application",
		reflect.TypeOf((*Application)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEnvironment", GoMethod: "AddEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "addExistingEnvironment", GoMethod: "AddExistingEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
			_jsii_.MemberMethod{JsiiMethod: "addHostedConfiguration", GoMethod: "AddHostedConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "addSourcedConfiguration", GoMethod: "AddSourcedConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "atDeploymentTick", GoMethod: "AtDeploymentTick"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "environments", GoMethod: "Environments"},
			_jsii_.MemberProperty{JsiiProperty: "extensible", GoGetter: "Extensible"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "on", GoMethod: "On"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentBaking", GoMethod: "OnDeploymentBaking"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentComplete", GoMethod: "OnDeploymentComplete"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentRolledBack", GoMethod: "OnDeploymentRolledBack"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStart", GoMethod: "OnDeploymentStart"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStep", GoMethod: "OnDeploymentStep"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "preCreateHostedConfigurationVersion", GoMethod: "PreCreateHostedConfigurationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "preStartDeployment", GoMethod: "PreStartDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Application{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApplication)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExtensible)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.ApplicationProps",
		reflect.TypeOf((*ApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.ApplicationReference",
		reflect.TypeOf((*ApplicationReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.CfnApplication",
		reflect.TypeOf((*CfnApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationRef", GoGetter: "ApplicationRef"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrApplicationId", GoGetter: "AttrApplicationId"},
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
		},
		func() interface{} {
			j := jsiiProxy_CfnApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApplicationRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnApplicationProps",
		reflect.TypeOf((*CfnApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.CfnConfigurationProfile",
		reflect.TypeOf((*CfnConfigurationProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrConfigurationProfileId", GoGetter: "AttrConfigurationProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "attrKmsKeyArn", GoGetter: "AttrKmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileRef", GoGetter: "ConfigurationProfileRef"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deletionProtectionCheck", GoGetter: "DeletionProtectionCheck"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdentifier", GoGetter: "KmsKeyIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "locationUri", GoGetter: "LocationUri"},
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
			_jsii_.MemberProperty{JsiiProperty: "retrievalRoleArn", GoGetter: "RetrievalRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "validators", GoGetter: "Validators"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationProfile{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConfigurationProfileRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnConfigurationProfile.ValidatorsProperty",
		reflect.TypeOf((*CfnConfigurationProfile_ValidatorsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnConfigurationProfileProps",
		reflect.TypeOf((*CfnConfigurationProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.CfnDeployment",
		reflect.TypeOf((*CfnDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDeploymentNumber", GoGetter: "AttrDeploymentNumber"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileId", GoGetter: "ConfigurationProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "configurationVersion", GoGetter: "ConfigurationVersion"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentRef", GoGetter: "DeploymentRef"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentStrategyId", GoGetter: "DeploymentStrategyId"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicExtensionParameters", GoGetter: "DynamicExtensionParameters"},
			_jsii_.MemberProperty{JsiiProperty: "environmentId", GoGetter: "EnvironmentId"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdentifier", GoGetter: "KmsKeyIdentifier"},
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
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeployment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDeploymentRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnDeployment.DynamicExtensionParametersProperty",
		reflect.TypeOf((*CfnDeployment_DynamicExtensionParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnDeploymentProps",
		reflect.TypeOf((*CfnDeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.CfnDeploymentStrategy",
		reflect.TypeOf((*CfnDeploymentStrategy)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentDurationInMinutes", GoGetter: "DeploymentDurationInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentStrategyRef", GoGetter: "DeploymentStrategyRef"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "finalBakeTimeInMinutes", GoGetter: "FinalBakeTimeInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "growthFactor", GoGetter: "GrowthFactor"},
			_jsii_.MemberProperty{JsiiProperty: "growthType", GoGetter: "GrowthType"},
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
			_jsii_.MemberProperty{JsiiProperty: "replicateTo", GoGetter: "ReplicateTo"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeploymentStrategy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDeploymentStrategyRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnDeploymentStrategyProps",
		reflect.TypeOf((*CfnDeploymentStrategyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.CfnEnvironment",
		reflect.TypeOf((*CfnEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrEnvironmentId", GoGetter: "AttrEnvironmentId"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deletionProtectionCheck", GoGetter: "DeletionProtectionCheck"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "environmentRef", GoGetter: "EnvironmentRef"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "monitors", GoGetter: "Monitors"},
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
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEnvironmentRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnEnvironment.MonitorProperty",
		reflect.TypeOf((*CfnEnvironment_MonitorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnEnvironment.MonitorsProperty",
		reflect.TypeOf((*CfnEnvironment_MonitorsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnEnvironmentProps",
		reflect.TypeOf((*CfnEnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.CfnExtension",
		reflect.TypeOf((*CfnExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actions", GoGetter: "Actions"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionNumber", GoGetter: "AttrVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "extensionRef", GoGetter: "ExtensionRef"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersionNumber", GoGetter: "LatestVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
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
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnExtension{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExtensionRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnExtension.ActionProperty",
		reflect.TypeOf((*CfnExtension_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnExtension.ParameterProperty",
		reflect.TypeOf((*CfnExtension_ParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.CfnExtensionAssociation",
		reflect.TypeOf((*CfnExtensionAssociation)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrExtensionArn", GoGetter: "AttrExtensionArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceArn", GoGetter: "AttrResourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "extensionAssociationRef", GoGetter: "ExtensionAssociationRef"},
			_jsii_.MemberProperty{JsiiProperty: "extensionIdentifier", GoGetter: "ExtensionIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "extensionVersionNumber", GoGetter: "ExtensionVersionNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "resourceIdentifier", GoGetter: "ResourceIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnExtensionAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExtensionAssociationRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnExtensionAssociationProps",
		reflect.TypeOf((*CfnExtensionAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnExtensionProps",
		reflect.TypeOf((*CfnExtensionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.CfnHostedConfigurationVersion",
		reflect.TypeOf((*CfnHostedConfigurationVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionNumber", GoGetter: "AttrVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileId", GoGetter: "ConfigurationProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "contentType", GoGetter: "ContentType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "hostedConfigurationVersionRef", GoGetter: "HostedConfigurationVersionRef"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersionNumber", GoGetter: "LatestVersionNumber"},
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
			_jsii_.MemberProperty{JsiiProperty: "versionLabel", GoGetter: "VersionLabel"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHostedConfigurationVersion{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IHostedConfigurationVersionRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.CfnHostedConfigurationVersionProps",
		reflect.TypeOf((*CfnHostedConfigurationVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.ConfigurationContent",
		reflect.TypeOf((*ConfigurationContent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "contentType", GoGetter: "ContentType"},
		},
		func() interface{} {
			return &jsiiProxy_ConfigurationContent{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.ConfigurationOptions",
		reflect.TypeOf((*ConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.ConfigurationProfileReference",
		reflect.TypeOf((*ConfigurationProfileReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.ConfigurationProps",
		reflect.TypeOf((*ConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.ConfigurationSource",
		reflect.TypeOf((*ConfigurationSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "locationUri", GoGetter: "LocationUri"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_ConfigurationSource{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appconfig.ConfigurationSourceType",
		reflect.TypeOf((*ConfigurationSourceType)(nil)).Elem(),
		map[string]interface{}{
			"S3": ConfigurationSourceType_S3,
			"SECRETS_MANAGER": ConfigurationSourceType_SECRETS_MANAGER,
			"SSM_PARAMETER": ConfigurationSourceType_SSM_PARAMETER,
			"SSM_DOCUMENT": ConfigurationSourceType_SSM_DOCUMENT,
			"CODE_PIPELINE": ConfigurationSourceType_CODE_PIPELINE,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appconfig.ConfigurationType",
		reflect.TypeOf((*ConfigurationType)(nil)).Elem(),
		map[string]interface{}{
			"FREEFORM": ConfigurationType_FREEFORM,
			"FEATURE_FLAGS": ConfigurationType_FEATURE_FLAGS,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appconfig.DeletionProtectionCheck",
		reflect.TypeOf((*DeletionProtectionCheck)(nil)).Elem(),
		map[string]interface{}{
			"ACCOUNT_DEFAULT": DeletionProtectionCheck_ACCOUNT_DEFAULT,
			"APPLY": DeletionProtectionCheck_APPLY,
			"BYPASS": DeletionProtectionCheck_BYPASS,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.DeploymentReference",
		reflect.TypeOf((*DeploymentReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.DeploymentStrategy",
		reflect.TypeOf((*DeploymentStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentDurationInMinutes", GoGetter: "DeploymentDurationInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentStrategyArn", GoGetter: "DeploymentStrategyArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentStrategyId", GoGetter: "DeploymentStrategyId"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "finalBakeTimeInMinutes", GoGetter: "FinalBakeTimeInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "growthFactor", GoGetter: "GrowthFactor"},
			_jsii_.MemberProperty{JsiiProperty: "growthType", GoGetter: "GrowthType"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DeploymentStrategy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDeploymentStrategy)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.DeploymentStrategyId",
		reflect.TypeOf((*DeploymentStrategyId)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentStrategyId{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.DeploymentStrategyProps",
		reflect.TypeOf((*DeploymentStrategyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.DeploymentStrategyReference",
		reflect.TypeOf((*DeploymentStrategyReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.Environment",
		reflect.TypeOf((*Environment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeployment", GoMethod: "AddDeployment"},
			_jsii_.MemberMethod{JsiiMethod: "addDeployments", GoMethod: "AddDeployments"},
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "atDeploymentTick", GoMethod: "AtDeploymentTick"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentQueue", GoGetter: "DeploymentQueue"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environmentArn", GoGetter: "EnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "environmentId", GoGetter: "EnvironmentId"},
			_jsii_.MemberProperty{JsiiProperty: "extensible", GoGetter: "Extensible"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadConfig", GoMethod: "GrantReadConfig"},
			_jsii_.MemberProperty{JsiiProperty: "monitors", GoGetter: "Monitors"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "on", GoMethod: "On"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentBaking", GoMethod: "OnDeploymentBaking"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentComplete", GoMethod: "OnDeploymentComplete"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentRolledBack", GoMethod: "OnDeploymentRolledBack"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStart", GoMethod: "OnDeploymentStart"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStep", GoMethod: "OnDeploymentStep"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "preCreateHostedConfigurationVersion", GoMethod: "PreCreateHostedConfigurationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "preStartDeployment", GoMethod: "PreStartDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Environment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEnvironment)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExtensible)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.EnvironmentAttributes",
		reflect.TypeOf((*EnvironmentAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.EnvironmentOptions",
		reflect.TypeOf((*EnvironmentOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.EnvironmentProps",
		reflect.TypeOf((*EnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.EnvironmentReference",
		reflect.TypeOf((*EnvironmentReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.EventBridgeDestination",
		reflect.TypeOf((*EventBridgeDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "extensionUri", GoGetter: "ExtensionUri"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_EventBridgeDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEventDestination)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.ExtensibleBase",
		reflect.TypeOf((*ExtensibleBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
			_jsii_.MemberMethod{JsiiMethod: "atDeploymentTick", GoMethod: "AtDeploymentTick"},
			_jsii_.MemberMethod{JsiiMethod: "on", GoMethod: "On"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentBaking", GoMethod: "OnDeploymentBaking"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentComplete", GoMethod: "OnDeploymentComplete"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentRolledBack", GoMethod: "OnDeploymentRolledBack"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStart", GoMethod: "OnDeploymentStart"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStep", GoMethod: "OnDeploymentStep"},
			_jsii_.MemberMethod{JsiiMethod: "preCreateHostedConfigurationVersion", GoMethod: "PreCreateHostedConfigurationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "preStartDeployment", GoMethod: "PreStartDeployment"},
		},
		func() interface{} {
			j := jsiiProxy_ExtensibleBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExtensible)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.Extension",
		reflect.TypeOf((*Extension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actions", GoGetter: "Actions"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "extensionArn", GoGetter: "ExtensionArn"},
			_jsii_.MemberProperty{JsiiProperty: "extensionId", GoGetter: "ExtensionId"},
			_jsii_.MemberProperty{JsiiProperty: "extensionVersionNumber", GoGetter: "ExtensionVersionNumber"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersionNumber", GoGetter: "LatestVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Extension{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExtension)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.ExtensionAssociationReference",
		reflect.TypeOf((*ExtensionAssociationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.ExtensionAttributes",
		reflect.TypeOf((*ExtensionAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.ExtensionOptions",
		reflect.TypeOf((*ExtensionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.ExtensionProps",
		reflect.TypeOf((*ExtensionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.ExtensionReference",
		reflect.TypeOf((*ExtensionReference)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appconfig.GrowthType",
		reflect.TypeOf((*GrowthType)(nil)).Elem(),
		map[string]interface{}{
			"LINEAR": GrowthType_LINEAR,
			"EXPONENTIAL": GrowthType_EXPONENTIAL,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.HostedConfiguration",
		reflect.TypeOf((*HostedConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExistingEnvironmentsToApplication", GoMethod: "AddExistingEnvironmentsToApplication"},
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "atDeploymentTick", GoMethod: "AtDeploymentTick"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileArn", GoGetter: "ConfigurationProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileId", GoGetter: "ConfigurationProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "contentType", GoGetter: "ContentType"},
			_jsii_.MemberProperty{JsiiProperty: "deletionProtectionCheck", GoGetter: "DeletionProtectionCheck"},
			_jsii_.MemberMethod{JsiiMethod: "deployConfigToEnvironments", GoMethod: "DeployConfigToEnvironments"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentKey", GoGetter: "DeploymentKey"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentStrategy", GoGetter: "DeploymentStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "deployTo", GoGetter: "DeployTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "extensible", GoGetter: "Extensible"},
			_jsii_.MemberProperty{JsiiProperty: "hostedConfigurationVersionArn", GoGetter: "HostedConfigurationVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersionNumber", GoGetter: "LatestVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "on", GoMethod: "On"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentBaking", GoMethod: "OnDeploymentBaking"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentComplete", GoMethod: "OnDeploymentComplete"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentRolledBack", GoMethod: "OnDeploymentRolledBack"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStart", GoMethod: "OnDeploymentStart"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStep", GoMethod: "OnDeploymentStep"},
			_jsii_.MemberMethod{JsiiMethod: "preCreateHostedConfigurationVersion", GoMethod: "PreCreateHostedConfigurationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "preStartDeployment", GoMethod: "PreStartDeployment"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "validators", GoGetter: "Validators"},
			_jsii_.MemberProperty{JsiiProperty: "versionLabel", GoGetter: "VersionLabel"},
			_jsii_.MemberProperty{JsiiProperty: "versionNumber", GoGetter: "VersionNumber"},
		},
		func() interface{} {
			j := jsiiProxy_HostedConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConfiguration)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExtensible)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.HostedConfigurationOptions",
		reflect.TypeOf((*HostedConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.HostedConfigurationProps",
		reflect.TypeOf((*HostedConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.HostedConfigurationVersionReference",
		reflect.TypeOf((*HostedConfigurationVersionReference)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IApplication",
		reflect.TypeOf((*IApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEnvironment", GoMethod: "AddEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "addExistingEnvironment", GoMethod: "AddExistingEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
			_jsii_.MemberMethod{JsiiMethod: "addHostedConfiguration", GoMethod: "AddHostedConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "addSourcedConfiguration", GoMethod: "AddSourcedConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "atDeploymentTick", GoMethod: "AtDeploymentTick"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "environments", GoMethod: "Environments"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "on", GoMethod: "On"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentBaking", GoMethod: "OnDeploymentBaking"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentComplete", GoMethod: "OnDeploymentComplete"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentRolledBack", GoMethod: "OnDeploymentRolledBack"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStart", GoMethod: "OnDeploymentStart"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStep", GoMethod: "OnDeploymentStep"},
			_jsii_.MemberMethod{JsiiMethod: "preCreateHostedConfigurationVersion", GoMethod: "PreCreateHostedConfigurationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "preStartDeployment", GoMethod: "PreStartDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IApplicationRef",
		reflect.TypeOf((*IApplicationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationRef", GoGetter: "ApplicationRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IApplicationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IConfiguration",
		reflect.TypeOf((*IConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileId", GoGetter: "ConfigurationProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentKey", GoGetter: "DeploymentKey"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentStrategy", GoGetter: "DeploymentStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "deployTo", GoGetter: "DeployTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "validators", GoGetter: "Validators"},
			_jsii_.MemberProperty{JsiiProperty: "versionNumber", GoGetter: "VersionNumber"},
		},
		func() interface{} {
			j := jsiiProxy_IConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IConfigurationProfileRef",
		reflect.TypeOf((*IConfigurationProfileRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileRef", GoGetter: "ConfigurationProfileRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IConfigurationProfileRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IDeploymentRef",
		reflect.TypeOf((*IDeploymentRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploymentRef", GoGetter: "DeploymentRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IDeploymentRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IDeploymentStrategy",
		reflect.TypeOf((*IDeploymentStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentDurationInMinutes", GoGetter: "DeploymentDurationInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentStrategyArn", GoGetter: "DeploymentStrategyArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentStrategyId", GoGetter: "DeploymentStrategyId"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "finalBakeTimeInMinutes", GoGetter: "FinalBakeTimeInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "growthFactor", GoGetter: "GrowthFactor"},
			_jsii_.MemberProperty{JsiiProperty: "growthType", GoGetter: "GrowthType"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDeploymentStrategy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IDeploymentStrategyRef",
		reflect.TypeOf((*IDeploymentStrategyRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploymentStrategyRef", GoGetter: "DeploymentStrategyRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IDeploymentStrategyRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IEnvironment",
		reflect.TypeOf((*IEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeployment", GoMethod: "AddDeployment"},
			_jsii_.MemberMethod{JsiiMethod: "addDeployments", GoMethod: "AddDeployments"},
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "atDeploymentTick", GoMethod: "AtDeploymentTick"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environmentArn", GoGetter: "EnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "environmentId", GoGetter: "EnvironmentId"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadConfig", GoMethod: "GrantReadConfig"},
			_jsii_.MemberProperty{JsiiProperty: "monitors", GoGetter: "Monitors"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "on", GoMethod: "On"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentBaking", GoMethod: "OnDeploymentBaking"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentComplete", GoMethod: "OnDeploymentComplete"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentRolledBack", GoMethod: "OnDeploymentRolledBack"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStart", GoMethod: "OnDeploymentStart"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStep", GoMethod: "OnDeploymentStep"},
			_jsii_.MemberMethod{JsiiMethod: "preCreateHostedConfigurationVersion", GoMethod: "PreCreateHostedConfigurationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "preStartDeployment", GoMethod: "PreStartDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IEnvironmentRef",
		reflect.TypeOf((*IEnvironmentRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "environmentRef", GoGetter: "EnvironmentRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IEnvironmentRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IEventDestination",
		reflect.TypeOf((*IEventDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "extensionUri", GoGetter: "ExtensionUri"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IEventDestination{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IExtensible",
		reflect.TypeOf((*IExtensible)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
			_jsii_.MemberMethod{JsiiMethod: "atDeploymentTick", GoMethod: "AtDeploymentTick"},
			_jsii_.MemberMethod{JsiiMethod: "on", GoMethod: "On"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentBaking", GoMethod: "OnDeploymentBaking"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentComplete", GoMethod: "OnDeploymentComplete"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentRolledBack", GoMethod: "OnDeploymentRolledBack"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStart", GoMethod: "OnDeploymentStart"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStep", GoMethod: "OnDeploymentStep"},
			_jsii_.MemberMethod{JsiiMethod: "preCreateHostedConfigurationVersion", GoMethod: "PreCreateHostedConfigurationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "preStartDeployment", GoMethod: "PreStartDeployment"},
		},
		func() interface{} {
			return &jsiiProxy_IExtensible{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IExtension",
		reflect.TypeOf((*IExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actions", GoGetter: "Actions"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "extensionArn", GoGetter: "ExtensionArn"},
			_jsii_.MemberProperty{JsiiProperty: "extensionId", GoGetter: "ExtensionId"},
			_jsii_.MemberProperty{JsiiProperty: "extensionVersionNumber", GoGetter: "ExtensionVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersionNumber", GoGetter: "LatestVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IExtension{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IExtensionAssociationRef",
		reflect.TypeOf((*IExtensionAssociationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "extensionAssociationRef", GoGetter: "ExtensionAssociationRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IExtensionAssociationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IExtensionRef",
		reflect.TypeOf((*IExtensionRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "extensionRef", GoGetter: "ExtensionRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IExtensionRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IHostedConfigurationVersionRef",
		reflect.TypeOf((*IHostedConfigurationVersionRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "hostedConfigurationVersionRef", GoGetter: "HostedConfigurationVersionRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IHostedConfigurationVersionRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_appconfig.IValidator",
		reflect.TypeOf((*IValidator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IValidator{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.JsonSchemaValidator",
		reflect.TypeOf((*JsonSchemaValidator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_JsonSchemaValidator{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IValidator)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.LambdaDestination",
		reflect.TypeOf((*LambdaDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "extensionUri", GoGetter: "ExtensionUri"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEventDestination)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.LambdaValidator",
		reflect.TypeOf((*LambdaValidator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaValidator{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IValidator)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.Monitor",
		reflect.TypeOf((*Monitor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmArn", GoGetter: "AlarmArn"},
			_jsii_.MemberProperty{JsiiProperty: "alarmRoleArn", GoGetter: "AlarmRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "isCompositeAlarm", GoGetter: "IsCompositeAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "monitorType", GoGetter: "MonitorType"},
		},
		func() interface{} {
			return &jsiiProxy_Monitor{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appconfig.MonitorType",
		reflect.TypeOf((*MonitorType)(nil)).Elem(),
		map[string]interface{}{
			"CLOUDWATCH": MonitorType_CLOUDWATCH,
			"CFN_MONITORS_PROPERTY": MonitorType_CFN_MONITORS_PROPERTY,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.Parameter",
		reflect.TypeOf((*Parameter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "isRequired", GoGetter: "IsRequired"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_Parameter{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appconfig.Platform",
		reflect.TypeOf((*Platform)(nil)).Elem(),
		map[string]interface{}{
			"X86_64": Platform_X86_64,
			"ARM_64": Platform_ARM_64,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.RolloutStrategy",
		reflect.TypeOf((*RolloutStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploymentDuration", GoGetter: "DeploymentDuration"},
			_jsii_.MemberProperty{JsiiProperty: "finalBakeTime", GoGetter: "FinalBakeTime"},
			_jsii_.MemberProperty{JsiiProperty: "growthFactor", GoGetter: "GrowthFactor"},
			_jsii_.MemberProperty{JsiiProperty: "growthType", GoGetter: "GrowthType"},
		},
		func() interface{} {
			return &jsiiProxy_RolloutStrategy{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.RolloutStrategyProps",
		reflect.TypeOf((*RolloutStrategyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.SnsDestination",
		reflect.TypeOf((*SnsDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "extensionUri", GoGetter: "ExtensionUri"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_SnsDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEventDestination)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appconfig.SourceType",
		reflect.TypeOf((*SourceType)(nil)).Elem(),
		map[string]interface{}{
			"LAMBDA": SourceType_LAMBDA,
			"SQS": SourceType_SQS,
			"SNS": SourceType_SNS,
			"EVENTS": SourceType_EVENTS,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.SourcedConfiguration",
		reflect.TypeOf((*SourcedConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExistingEnvironmentsToApplication", GoMethod: "AddExistingEnvironmentsToApplication"},
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "atDeploymentTick", GoMethod: "AtDeploymentTick"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileArn", GoGetter: "ConfigurationProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileId", GoGetter: "ConfigurationProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "deletionProtectionCheck", GoGetter: "DeletionProtectionCheck"},
			_jsii_.MemberMethod{JsiiMethod: "deployConfigToEnvironments", GoMethod: "DeployConfigToEnvironments"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentKey", GoGetter: "DeploymentKey"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentStrategy", GoGetter: "DeploymentStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "deployTo", GoGetter: "DeployTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "extensible", GoGetter: "Extensible"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "on", GoMethod: "On"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentBaking", GoMethod: "OnDeploymentBaking"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentComplete", GoMethod: "OnDeploymentComplete"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentRolledBack", GoMethod: "OnDeploymentRolledBack"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStart", GoMethod: "OnDeploymentStart"},
			_jsii_.MemberMethod{JsiiMethod: "onDeploymentStep", GoMethod: "OnDeploymentStep"},
			_jsii_.MemberMethod{JsiiMethod: "preCreateHostedConfigurationVersion", GoMethod: "PreCreateHostedConfigurationVersion"},
			_jsii_.MemberMethod{JsiiMethod: "preStartDeployment", GoMethod: "PreStartDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "retrievalRole", GoGetter: "RetrievalRole"},
			_jsii_.MemberProperty{JsiiProperty: "sourceKey", GoGetter: "SourceKey"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "validators", GoGetter: "Validators"},
			_jsii_.MemberProperty{JsiiProperty: "versionNumber", GoGetter: "VersionNumber"},
		},
		func() interface{} {
			j := jsiiProxy_SourcedConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConfiguration)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExtensible)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.SourcedConfigurationOptions",
		reflect.TypeOf((*SourcedConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_appconfig.SourcedConfigurationProps",
		reflect.TypeOf((*SourcedConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_appconfig.SqsDestination",
		reflect.TypeOf((*SqsDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "extensionUri", GoGetter: "ExtensionUri"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_SqsDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEventDestination)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_appconfig.ValidatorType",
		reflect.TypeOf((*ValidatorType)(nil)).Elem(),
		map[string]interface{}{
			"JSON_SCHEMA": ValidatorType_JSON_SCHEMA,
			"LAMBDA": ValidatorType_LAMBDA,
		},
	)
}
