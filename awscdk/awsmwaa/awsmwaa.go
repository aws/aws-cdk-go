package awsmwaa

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment",
		reflect.TypeOf((*CfnEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "airflowConfigurationOptions", GoGetter: "AirflowConfigurationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "airflowVersion", GoGetter: "AirflowVersion"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrLoggingConfigurationDagProcessingLogsCloudWatchLogGroupArn", GoGetter: "AttrLoggingConfigurationDagProcessingLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrLoggingConfigurationSchedulerLogsCloudWatchLogGroupArn", GoGetter: "AttrLoggingConfigurationSchedulerLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrLoggingConfigurationTaskLogsCloudWatchLogGroupArn", GoGetter: "AttrLoggingConfigurationTaskLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrLoggingConfigurationWebserverLogsCloudWatchLogGroupArn", GoGetter: "AttrLoggingConfigurationWebserverLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrLoggingConfigurationWorkerLogsCloudWatchLogGroupArn", GoGetter: "AttrLoggingConfigurationWorkerLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrWebserverUrl", GoGetter: "AttrWebserverUrl"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dagS3Path", GoGetter: "DagS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "environmentClass", GoGetter: "EnvironmentClass"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfiguration", GoGetter: "LoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maxWorkers", GoGetter: "MaxWorkers"},
			_jsii_.MemberProperty{JsiiProperty: "minWorkers", GoGetter: "MinWorkers"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkConfiguration", GoGetter: "NetworkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsS3ObjectVersion", GoGetter: "PluginsS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsS3Path", GoGetter: "PluginsS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsS3ObjectVersion", GoGetter: "RequirementsS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsS3Path", GoGetter: "RequirementsS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "schedulers", GoGetter: "Schedulers"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceBucketArn", GoGetter: "SourceBucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "webserverAccessMode", GoGetter: "WebserverAccessMode"},
			_jsii_.MemberProperty{JsiiProperty: "weeklyMaintenanceWindowStart", GoGetter: "WeeklyMaintenanceWindowStart"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment.LoggingConfigurationProperty",
		reflect.TypeOf((*CfnEnvironment_LoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment.ModuleLoggingConfigurationProperty",
		reflect.TypeOf((*CfnEnvironment_ModuleLoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnEnvironment_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_mwaa.CfnEnvironmentProps",
		reflect.TypeOf((*CfnEnvironmentProps)(nil)).Elem(),
	)
}
