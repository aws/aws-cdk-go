// This module is deprecated. All constructs are now available under aws-cdk-lib/aws-appconfig
package awscdkappconfigalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.Action",
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
		"@aws-cdk/aws-appconfig-alpha.ActionPoint",
		reflect.TypeOf((*ActionPoint)(nil)).Elem(),
		map[string]interface{}{
			"PRE_CREATE_HOSTED_CONFIGURATION_VERSION": ActionPoint_PRE_CREATE_HOSTED_CONFIGURATION_VERSION,
			"PRE_START_DEPLOYMENT": ActionPoint_PRE_START_DEPLOYMENT,
			"ON_DEPLOYMENT_START": ActionPoint_ON_DEPLOYMENT_START,
			"ON_DEPLOYMENT_STEP": ActionPoint_ON_DEPLOYMENT_STEP,
			"ON_DEPLOYMENT_BAKING": ActionPoint_ON_DEPLOYMENT_BAKING,
			"ON_DEPLOYMENT_COMPLETE": ActionPoint_ON_DEPLOYMENT_COMPLETE,
			"ON_DEPLOYMENT_ROLLED_BACK": ActionPoint_ON_DEPLOYMENT_ROLLED_BACK,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appconfig-alpha.ActionProps",
		reflect.TypeOf((*ActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.Application",
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
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environments", GoGetter: "Environments"},
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
		"@aws-cdk/aws-appconfig-alpha.ApplicationProps",
		reflect.TypeOf((*ApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationContent",
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
		"@aws-cdk/aws-appconfig-alpha.ConfigurationOptions",
		reflect.TypeOf((*ConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationProps",
		reflect.TypeOf((*ConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.ConfigurationSource",
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
		"@aws-cdk/aws-appconfig-alpha.ConfigurationSourceType",
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
		"@aws-cdk/aws-appconfig-alpha.ConfigurationType",
		reflect.TypeOf((*ConfigurationType)(nil)).Elem(),
		map[string]interface{}{
			"FREEFORM": ConfigurationType_FREEFORM,
			"FEATURE_FLAGS": ConfigurationType_FEATURE_FLAGS,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.DeploymentStrategy",
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
		"@aws-cdk/aws-appconfig-alpha.DeploymentStrategyId",
		reflect.TypeOf((*DeploymentStrategyId)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentStrategyId{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appconfig-alpha.DeploymentStrategyProps",
		reflect.TypeOf((*DeploymentStrategyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.Environment",
		reflect.TypeOf((*Environment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environmentArn", GoGetter: "EnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "environmentId", GoGetter: "EnvironmentId"},
			_jsii_.MemberProperty{JsiiProperty: "extensible", GoGetter: "Extensible"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
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
		"@aws-cdk/aws-appconfig-alpha.EnvironmentAttributes",
		reflect.TypeOf((*EnvironmentAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appconfig-alpha.EnvironmentOptions",
		reflect.TypeOf((*EnvironmentOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appconfig-alpha.EnvironmentProps",
		reflect.TypeOf((*EnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.EventBridgeDestination",
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
		"@aws-cdk/aws-appconfig-alpha.ExtensibleBase",
		reflect.TypeOf((*ExtensibleBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
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
		"@aws-cdk/aws-appconfig-alpha.Extension",
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
		"@aws-cdk/aws-appconfig-alpha.ExtensionAttributes",
		reflect.TypeOf((*ExtensionAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appconfig-alpha.ExtensionOptions",
		reflect.TypeOf((*ExtensionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appconfig-alpha.ExtensionProps",
		reflect.TypeOf((*ExtensionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-appconfig-alpha.GrowthType",
		reflect.TypeOf((*GrowthType)(nil)).Elem(),
		map[string]interface{}{
			"LINEAR": GrowthType_LINEAR,
			"EXPONENTIAL": GrowthType_EXPONENTIAL,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.HostedConfiguration",
		reflect.TypeOf((*HostedConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExistingEnvironmentsToApplication", GoMethod: "AddExistingEnvironmentsToApplication"},
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileArn", GoGetter: "ConfigurationProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileId", GoGetter: "ConfigurationProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "contentType", GoGetter: "ContentType"},
			_jsii_.MemberMethod{JsiiMethod: "deploy", GoMethod: "Deploy"},
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
		"@aws-cdk/aws-appconfig-alpha.HostedConfigurationOptions",
		reflect.TypeOf((*HostedConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appconfig-alpha.HostedConfigurationProps",
		reflect.TypeOf((*HostedConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-appconfig-alpha.IApplication",
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
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environments", GoGetter: "Environments"},
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
		"@aws-cdk/aws-appconfig-alpha.IConfiguration",
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
		"@aws-cdk/aws-appconfig-alpha.IDeploymentStrategy",
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
		"@aws-cdk/aws-appconfig-alpha.IEnvironment",
		reflect.TypeOf((*IEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environmentArn", GoGetter: "EnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "environmentId", GoGetter: "EnvironmentId"},
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
		"@aws-cdk/aws-appconfig-alpha.IEventDestination",
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
		"@aws-cdk/aws-appconfig-alpha.IExtensible",
		reflect.TypeOf((*IExtensible)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
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
		"@aws-cdk/aws-appconfig-alpha.IExtension",
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
		"@aws-cdk/aws-appconfig-alpha.IValidator",
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
		"@aws-cdk/aws-appconfig-alpha.JsonSchemaValidator",
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
		"@aws-cdk/aws-appconfig-alpha.LambdaDestination",
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
		"@aws-cdk/aws-appconfig-alpha.LambdaValidator",
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
		"@aws-cdk/aws-appconfig-alpha.Monitor",
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
		"@aws-cdk/aws-appconfig-alpha.MonitorType",
		reflect.TypeOf((*MonitorType)(nil)).Elem(),
		map[string]interface{}{
			"CLOUDWATCH": MonitorType_CLOUDWATCH,
			"CFN_MONITORS_PROPERTY": MonitorType_CFN_MONITORS_PROPERTY,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.Parameter",
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
		"@aws-cdk/aws-appconfig-alpha.Platform",
		reflect.TypeOf((*Platform)(nil)).Elem(),
		map[string]interface{}{
			"X86_64": Platform_X86_64,
			"ARM_64": Platform_ARM_64,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.RolloutStrategy",
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
		"@aws-cdk/aws-appconfig-alpha.RolloutStrategyProps",
		reflect.TypeOf((*RolloutStrategyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.SnsDestination",
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
		"@aws-cdk/aws-appconfig-alpha.SourceType",
		reflect.TypeOf((*SourceType)(nil)).Elem(),
		map[string]interface{}{
			"LAMBDA": SourceType_LAMBDA,
			"SQS": SourceType_SQS,
			"SNS": SourceType_SNS,
			"EVENTS": SourceType_EVENTS,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.SourcedConfiguration",
		reflect.TypeOf((*SourcedConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExistingEnvironmentsToApplication", GoMethod: "AddExistingEnvironmentsToApplication"},
			_jsii_.MemberMethod{JsiiMethod: "addExtension", GoMethod: "AddExtension"},
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileArn", GoGetter: "ConfigurationProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "configurationProfileId", GoGetter: "ConfigurationProfileId"},
			_jsii_.MemberMethod{JsiiMethod: "deploy", GoMethod: "Deploy"},
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
		"@aws-cdk/aws-appconfig-alpha.SourcedConfigurationOptions",
		reflect.TypeOf((*SourcedConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-appconfig-alpha.SourcedConfigurationProps",
		reflect.TypeOf((*SourcedConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-appconfig-alpha.SqsDestination",
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
		"@aws-cdk/aws-appconfig-alpha.ValidatorType",
		reflect.TypeOf((*ValidatorType)(nil)).Elem(),
		map[string]interface{}{
			"JSON_SCHEMA": ValidatorType_JSON_SCHEMA,
			"LAMBDA": ValidatorType_LAMBDA,
		},
	)
}
