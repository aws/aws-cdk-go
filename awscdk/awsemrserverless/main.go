package awsemrserverless

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_emrserverless.CfnApplication",
		reflect.TypeOf((*CfnApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "architecture", GoGetter: "Architecture"},
			_jsii_.MemberProperty{JsiiProperty: "attrApplicationId", GoGetter: "AttrApplicationId"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "autoStartConfiguration", GoGetter: "AutoStartConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "autoStopConfiguration", GoGetter: "AutoStopConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "imageConfiguration", GoGetter: "ImageConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "initialCapacity", GoGetter: "InitialCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "interactiveConfiguration", GoGetter: "InteractiveConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maximumCapacity", GoGetter: "MaximumCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringConfiguration", GoGetter: "MonitoringConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkConfiguration", GoGetter: "NetworkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "releaseLabel", GoGetter: "ReleaseLabel"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "runtimeConfiguration", GoGetter: "RuntimeConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "schedulerConfiguration", GoGetter: "SchedulerConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "workerTypeSpecifications", GoGetter: "WorkerTypeSpecifications"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.AutoStartConfigurationProperty",
		reflect.TypeOf((*CfnApplication_AutoStartConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.AutoStopConfigurationProperty",
		reflect.TypeOf((*CfnApplication_AutoStopConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.CloudWatchLoggingConfigurationProperty",
		reflect.TypeOf((*CfnApplication_CloudWatchLoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.ConfigurationObjectProperty",
		reflect.TypeOf((*CfnApplication_ConfigurationObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.ImageConfigurationInputProperty",
		reflect.TypeOf((*CfnApplication_ImageConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.InitialCapacityConfigKeyValuePairProperty",
		reflect.TypeOf((*CfnApplication_InitialCapacityConfigKeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.InitialCapacityConfigProperty",
		reflect.TypeOf((*CfnApplication_InitialCapacityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.InteractiveConfigurationProperty",
		reflect.TypeOf((*CfnApplication_InteractiveConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.LogTypeMapKeyValuePairProperty",
		reflect.TypeOf((*CfnApplication_LogTypeMapKeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.ManagedPersistenceMonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplication_ManagedPersistenceMonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.MaximumAllowedResourcesProperty",
		reflect.TypeOf((*CfnApplication_MaximumAllowedResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.MonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplication_MonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnApplication_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.S3MonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplication_S3MonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.SchedulerConfigurationProperty",
		reflect.TypeOf((*CfnApplication_SchedulerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.WorkerConfigurationProperty",
		reflect.TypeOf((*CfnApplication_WorkerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplication.WorkerTypeSpecificationInputProperty",
		reflect.TypeOf((*CfnApplication_WorkerTypeSpecificationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_emrserverless.CfnApplicationProps",
		reflect.TypeOf((*CfnApplicationProps)(nil)).Elem(),
	)
}
