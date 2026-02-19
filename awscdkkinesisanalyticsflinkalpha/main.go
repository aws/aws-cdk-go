// A CDK Construct Library for Kinesis Analytics Flink applications
package awscdkkinesisanalyticsflinkalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Application",
		reflect.TypeOf((*Application)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackPressuredTimeMsPerSecond", GoMethod: "MetricBackPressuredTimeMsPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "metricBusyTimePerMsPerSecond", GoMethod: "MetricBusyTimePerMsPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "metricCpuUtilization", GoMethod: "MetricCpuUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricCurrentInputWatermark", GoMethod: "MetricCurrentInputWatermark"},
			_jsii_.MemberMethod{JsiiMethod: "metricCurrentOutputWatermark", GoMethod: "MetricCurrentOutputWatermark"},
			_jsii_.MemberMethod{JsiiMethod: "metricDowntime", GoMethod: "MetricDowntime"},
			_jsii_.MemberMethod{JsiiMethod: "metricFullRestarts", GoMethod: "MetricFullRestarts"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeapMemoryUtilization", GoMethod: "MetricHeapMemoryUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricIdleTimeMsPerSecond", GoMethod: "MetricIdleTimeMsPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "metricKpus", GoMethod: "MetricKpus"},
			_jsii_.MemberMethod{JsiiMethod: "metricLastCheckpointDuration", GoMethod: "MetricLastCheckpointDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricLastCheckpointSize", GoMethod: "MetricLastCheckpointSize"},
			_jsii_.MemberMethod{JsiiMethod: "metricManagedMemoryTotal", GoMethod: "MetricManagedMemoryTotal"},
			_jsii_.MemberMethod{JsiiMethod: "metricManagedMemoryUsed", GoMethod: "MetricManagedMemoryUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricManagedMemoryUtilization", GoMethod: "MetricManagedMemoryUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfFailedCheckpoints", GoMethod: "MetricNumberOfFailedCheckpoints"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumLateRecordsDropped", GoMethod: "MetricNumLateRecordsDropped"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumRecordsIn", GoMethod: "MetricNumRecordsIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumRecordsInPerSecond", GoMethod: "MetricNumRecordsInPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumRecordsOut", GoMethod: "MetricNumRecordsOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumRecordsOutPerSecond", GoMethod: "MetricNumRecordsOutPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "metricOldGenerationGCCount", GoMethod: "MetricOldGenerationGCCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricOldGenerationGCTime", GoMethod: "MetricOldGenerationGCTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricThreadsCount", GoMethod: "MetricThreadsCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricUptime", GoMethod: "MetricUptime"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Application{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApplication)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.ApplicationAttributes",
		reflect.TypeOf((*ApplicationAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.ApplicationCode",
		reflect.TypeOf((*ApplicationCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ApplicationCode{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.ApplicationCodeConfig",
		reflect.TypeOf((*ApplicationCodeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.ApplicationProps",
		reflect.TypeOf((*ApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.IApplication",
		reflect.TypeOf((*IApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackPressuredTimeMsPerSecond", GoMethod: "MetricBackPressuredTimeMsPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "metricBusyTimePerMsPerSecond", GoMethod: "MetricBusyTimePerMsPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "metricCpuUtilization", GoMethod: "MetricCpuUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricCurrentInputWatermark", GoMethod: "MetricCurrentInputWatermark"},
			_jsii_.MemberMethod{JsiiMethod: "metricCurrentOutputWatermark", GoMethod: "MetricCurrentOutputWatermark"},
			_jsii_.MemberMethod{JsiiMethod: "metricDowntime", GoMethod: "MetricDowntime"},
			_jsii_.MemberMethod{JsiiMethod: "metricFullRestarts", GoMethod: "MetricFullRestarts"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeapMemoryUtilization", GoMethod: "MetricHeapMemoryUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricIdleTimeMsPerSecond", GoMethod: "MetricIdleTimeMsPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "metricKpus", GoMethod: "MetricKpus"},
			_jsii_.MemberMethod{JsiiMethod: "metricLastCheckpointDuration", GoMethod: "MetricLastCheckpointDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricLastCheckpointSize", GoMethod: "MetricLastCheckpointSize"},
			_jsii_.MemberMethod{JsiiMethod: "metricManagedMemoryTotal", GoMethod: "MetricManagedMemoryTotal"},
			_jsii_.MemberMethod{JsiiMethod: "metricManagedMemoryUsed", GoMethod: "MetricManagedMemoryUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricManagedMemoryUtilization", GoMethod: "MetricManagedMemoryUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumberOfFailedCheckpoints", GoMethod: "MetricNumberOfFailedCheckpoints"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumLateRecordsDropped", GoMethod: "MetricNumLateRecordsDropped"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumRecordsIn", GoMethod: "MetricNumRecordsIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumRecordsInPerSecond", GoMethod: "MetricNumRecordsInPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumRecordsOut", GoMethod: "MetricNumRecordsOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumRecordsOutPerSecond", GoMethod: "MetricNumRecordsOutPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "metricOldGenerationGCCount", GoMethod: "MetricOldGenerationGCCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricOldGenerationGCTime", GoMethod: "MetricOldGenerationGCTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricThreadsCount", GoMethod: "MetricThreadsCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricUptime", GoMethod: "MetricUptime"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.LogLevel",
		reflect.TypeOf((*LogLevel)(nil)).Elem(),
		map[string]interface{}{
			"DEBUG": LogLevel_DEBUG,
			"INFO": LogLevel_INFO,
			"WARN": LogLevel_WARN,
			"ERROR": LogLevel_ERROR,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.MetricsLevel",
		reflect.TypeOf((*MetricsLevel)(nil)).Elem(),
		map[string]interface{}{
			"APPLICATION": MetricsLevel_APPLICATION,
			"TASK": MetricsLevel_TASK,
			"OPERATOR": MetricsLevel_OPERATOR,
			"PARALLELISM": MetricsLevel_PARALLELISM,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.PropertyGroups",
		reflect.TypeOf((*PropertyGroups)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		reflect.TypeOf((*Runtime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_Runtime{}
		},
	)
}
