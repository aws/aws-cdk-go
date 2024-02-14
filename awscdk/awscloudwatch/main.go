package awscloudwatch

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.Alarm",
		reflect.TypeOf((*Alarm)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlarmAction", GoMethod: "AddAlarmAction"},
			_jsii_.MemberMethod{JsiiMethod: "addInsufficientDataAction", GoMethod: "AddInsufficientDataAction"},
			_jsii_.MemberMethod{JsiiMethod: "addOkAction", GoMethod: "AddOkAction"},
			_jsii_.MemberProperty{JsiiProperty: "alarmActionArns", GoGetter: "AlarmActionArns"},
			_jsii_.MemberProperty{JsiiProperty: "alarmArn", GoGetter: "AlarmArn"},
			_jsii_.MemberProperty{JsiiProperty: "alarmName", GoGetter: "AlarmName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "insufficientDataActionArns", GoGetter: "InsufficientDataActionArns"},
			_jsii_.MemberProperty{JsiiProperty: "metric", GoGetter: "Metric"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "okActionArns", GoGetter: "OkActionArns"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "renderAlarmRule", GoMethod: "RenderAlarmRule"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toAnnotation", GoMethod: "ToAnnotation"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Alarm{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AlarmBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.AlarmActionConfig",
		reflect.TypeOf((*AlarmActionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.AlarmBase",
		reflect.TypeOf((*AlarmBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlarmAction", GoMethod: "AddAlarmAction"},
			_jsii_.MemberMethod{JsiiMethod: "addInsufficientDataAction", GoMethod: "AddInsufficientDataAction"},
			_jsii_.MemberMethod{JsiiMethod: "addOkAction", GoMethod: "AddOkAction"},
			_jsii_.MemberProperty{JsiiProperty: "alarmActionArns", GoGetter: "AlarmActionArns"},
			_jsii_.MemberProperty{JsiiProperty: "alarmArn", GoGetter: "AlarmArn"},
			_jsii_.MemberProperty{JsiiProperty: "alarmName", GoGetter: "AlarmName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "insufficientDataActionArns", GoGetter: "InsufficientDataActionArns"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "okActionArns", GoGetter: "OkActionArns"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "renderAlarmRule", GoMethod: "RenderAlarmRule"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AlarmBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAlarm)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.AlarmProps",
		reflect.TypeOf((*AlarmProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.AlarmRule",
		reflect.TypeOf((*AlarmRule)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AlarmRule{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.AlarmState",
		reflect.TypeOf((*AlarmState)(nil)).Elem(),
		map[string]interface{}{
			"ALARM": AlarmState_ALARM,
			"OK": AlarmState_OK,
			"INSUFFICIENT_DATA": AlarmState_INSUFFICIENT_DATA,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.AlarmStatusWidget",
		reflect.TypeOf((*AlarmStatusWidget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "copyMetricWarnings", GoMethod: "CopyMetricWarnings"},
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
			_jsii_.MemberProperty{JsiiProperty: "x", GoGetter: "X"},
			_jsii_.MemberProperty{JsiiProperty: "y", GoGetter: "Y"},
		},
		func() interface{} {
			j := jsiiProxy_AlarmStatusWidget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConcreteWidget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.AlarmStatusWidgetProps",
		reflect.TypeOf((*AlarmStatusWidgetProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.AlarmStatusWidgetSortBy",
		reflect.TypeOf((*AlarmStatusWidgetSortBy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": AlarmStatusWidgetSortBy_DEFAULT,
			"STATE_UPDATED_TIMESTAMP": AlarmStatusWidgetSortBy_STATE_UPDATED_TIMESTAMP,
			"TIMESTAMP": AlarmStatusWidgetSortBy_TIMESTAMP,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.AlarmWidget",
		reflect.TypeOf((*AlarmWidget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "copyMetricWarnings", GoMethod: "CopyMetricWarnings"},
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
			_jsii_.MemberProperty{JsiiProperty: "x", GoGetter: "X"},
			_jsii_.MemberProperty{JsiiProperty: "y", GoGetter: "Y"},
		},
		func() interface{} {
			j := jsiiProxy_AlarmWidget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConcreteWidget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.AlarmWidgetProps",
		reflect.TypeOf((*AlarmWidgetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.CfnAlarm",
		reflect.TypeOf((*CfnAlarm)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionsEnabled", GoGetter: "ActionsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alarmActions", GoGetter: "AlarmActions"},
			_jsii_.MemberProperty{JsiiProperty: "alarmDescription", GoGetter: "AlarmDescription"},
			_jsii_.MemberProperty{JsiiProperty: "alarmName", GoGetter: "AlarmName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "comparisonOperator", GoGetter: "ComparisonOperator"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "dimensions", GoGetter: "Dimensions"},
			_jsii_.MemberProperty{JsiiProperty: "evaluateLowSampleCountPercentile", GoGetter: "EvaluateLowSampleCountPercentile"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberProperty{JsiiProperty: "extendedStatistic", GoGetter: "ExtendedStatistic"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "insufficientDataActions", GoGetter: "InsufficientDataActions"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "metricName", GoGetter: "MetricName"},
			_jsii_.MemberProperty{JsiiProperty: "metrics", GoGetter: "Metrics"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "okActions", GoGetter: "OkActions"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdMetricId", GoGetter: "ThresholdMetricId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "treatMissingData", GoGetter: "TreatMissingData"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlarm{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAlarm.DimensionProperty",
		reflect.TypeOf((*CfnAlarm_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAlarm.MetricDataQueryProperty",
		reflect.TypeOf((*CfnAlarm_MetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAlarm.MetricProperty",
		reflect.TypeOf((*CfnAlarm_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAlarm.MetricStatProperty",
		reflect.TypeOf((*CfnAlarm_MetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAlarmProps",
		reflect.TypeOf((*CfnAlarmProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector",
		reflect.TypeOf((*CfnAnomalyDetector)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dimensions", GoGetter: "Dimensions"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "metricMathAnomalyDetector", GoGetter: "MetricMathAnomalyDetector"},
			_jsii_.MemberProperty{JsiiProperty: "metricName", GoGetter: "MetricName"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "singleMetricAnomalyDetector", GoGetter: "SingleMetricAnomalyDetector"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stat", GoGetter: "Stat"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnomalyDetector{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector.ConfigurationProperty",
		reflect.TypeOf((*CfnAnomalyDetector_ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector.DimensionProperty",
		reflect.TypeOf((*CfnAnomalyDetector_DimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector.MetricDataQueryProperty",
		reflect.TypeOf((*CfnAnomalyDetector_MetricDataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector.MetricMathAnomalyDetectorProperty",
		reflect.TypeOf((*CfnAnomalyDetector_MetricMathAnomalyDetectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector.MetricProperty",
		reflect.TypeOf((*CfnAnomalyDetector_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector.MetricStatProperty",
		reflect.TypeOf((*CfnAnomalyDetector_MetricStatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector.RangeProperty",
		reflect.TypeOf((*CfnAnomalyDetector_RangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector.SingleMetricAnomalyDetectorProperty",
		reflect.TypeOf((*CfnAnomalyDetector_SingleMetricAnomalyDetectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetectorProps",
		reflect.TypeOf((*CfnAnomalyDetectorProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.CfnCompositeAlarm",
		reflect.TypeOf((*CfnCompositeAlarm)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionsEnabled", GoGetter: "ActionsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "actionsSuppressor", GoGetter: "ActionsSuppressor"},
			_jsii_.MemberProperty{JsiiProperty: "actionsSuppressorExtensionPeriod", GoGetter: "ActionsSuppressorExtensionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "actionsSuppressorWaitPeriod", GoGetter: "ActionsSuppressorWaitPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alarmActions", GoGetter: "AlarmActions"},
			_jsii_.MemberProperty{JsiiProperty: "alarmDescription", GoGetter: "AlarmDescription"},
			_jsii_.MemberProperty{JsiiProperty: "alarmName", GoGetter: "AlarmName"},
			_jsii_.MemberProperty{JsiiProperty: "alarmRule", GoGetter: "AlarmRule"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "insufficientDataActions", GoGetter: "InsufficientDataActions"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "okActions", GoGetter: "OkActions"},
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
			j := jsiiProxy_CfnCompositeAlarm{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnCompositeAlarmProps",
		reflect.TypeOf((*CfnCompositeAlarmProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.CfnDashboard",
		reflect.TypeOf((*CfnDashboard)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dashboardBody", GoGetter: "DashboardBody"},
			_jsii_.MemberProperty{JsiiProperty: "dashboardName", GoGetter: "DashboardName"},
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
			j := jsiiProxy_CfnDashboard{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnDashboardProps",
		reflect.TypeOf((*CfnDashboardProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.CfnInsightRule",
		reflect.TypeOf((*CfnInsightRule)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrRuleName", GoGetter: "AttrRuleName"},
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
			_jsii_.MemberProperty{JsiiProperty: "ruleBody", GoGetter: "RuleBody"},
			_jsii_.MemberProperty{JsiiProperty: "ruleName", GoGetter: "RuleName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleState", GoGetter: "RuleState"},
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
			j := jsiiProxy_CfnInsightRule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnInsightRuleProps",
		reflect.TypeOf((*CfnInsightRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.CfnMetricStream",
		reflect.TypeOf((*CfnMetricStream)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrCreationDate", GoGetter: "AttrCreationDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdateDate", GoGetter: "AttrLastUpdateDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludeFilters", GoGetter: "ExcludeFilters"},
			_jsii_.MemberProperty{JsiiProperty: "firehoseArn", GoGetter: "FirehoseArn"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "includeFilters", GoGetter: "IncludeFilters"},
			_jsii_.MemberProperty{JsiiProperty: "includeLinkedAccountsMetrics", GoGetter: "IncludeLinkedAccountsMetrics"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "outputFormat", GoGetter: "OutputFormat"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "statisticsConfigurations", GoGetter: "StatisticsConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMetricStream{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnMetricStream.MetricStreamFilterProperty",
		reflect.TypeOf((*CfnMetricStream_MetricStreamFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnMetricStream.MetricStreamStatisticsConfigurationProperty",
		reflect.TypeOf((*CfnMetricStream_MetricStreamStatisticsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnMetricStream.MetricStreamStatisticsMetricProperty",
		reflect.TypeOf((*CfnMetricStream_MetricStreamStatisticsMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CfnMetricStreamProps",
		reflect.TypeOf((*CfnMetricStreamProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.Color",
		reflect.TypeOf((*Color)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Color{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.Column",
		reflect.TypeOf((*Column)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addWidget", GoMethod: "AddWidget"},
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "widgets", GoGetter: "Widgets"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
		},
		func() interface{} {
			j := jsiiProxy_Column{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWidget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CommonMetricOptions",
		reflect.TypeOf((*CommonMetricOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.ComparisonOperator",
		reflect.TypeOf((*ComparisonOperator)(nil)).Elem(),
		map[string]interface{}{
			"GREATER_THAN_OR_EQUAL_TO_THRESHOLD": ComparisonOperator_GREATER_THAN_OR_EQUAL_TO_THRESHOLD,
			"GREATER_THAN_THRESHOLD": ComparisonOperator_GREATER_THAN_THRESHOLD,
			"LESS_THAN_THRESHOLD": ComparisonOperator_LESS_THAN_THRESHOLD,
			"LESS_THAN_OR_EQUAL_TO_THRESHOLD": ComparisonOperator_LESS_THAN_OR_EQUAL_TO_THRESHOLD,
			"LESS_THAN_LOWER_OR_GREATER_THAN_UPPER_THRESHOLD": ComparisonOperator_LESS_THAN_LOWER_OR_GREATER_THAN_UPPER_THRESHOLD,
			"GREATER_THAN_UPPER_THRESHOLD": ComparisonOperator_GREATER_THAN_UPPER_THRESHOLD,
			"LESS_THAN_LOWER_THRESHOLD": ComparisonOperator_LESS_THAN_LOWER_THRESHOLD,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.CompositeAlarm",
		reflect.TypeOf((*CompositeAlarm)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlarmAction", GoMethod: "AddAlarmAction"},
			_jsii_.MemberMethod{JsiiMethod: "addInsufficientDataAction", GoMethod: "AddInsufficientDataAction"},
			_jsii_.MemberMethod{JsiiMethod: "addOkAction", GoMethod: "AddOkAction"},
			_jsii_.MemberProperty{JsiiProperty: "alarmActionArns", GoGetter: "AlarmActionArns"},
			_jsii_.MemberProperty{JsiiProperty: "alarmArn", GoGetter: "AlarmArn"},
			_jsii_.MemberProperty{JsiiProperty: "alarmName", GoGetter: "AlarmName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "insufficientDataActionArns", GoGetter: "InsufficientDataActionArns"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "okActionArns", GoGetter: "OkActionArns"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "renderAlarmRule", GoMethod: "RenderAlarmRule"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CompositeAlarm{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AlarmBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CompositeAlarmProps",
		reflect.TypeOf((*CompositeAlarmProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.ConcreteWidget",
		reflect.TypeOf((*ConcreteWidget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "copyMetricWarnings", GoMethod: "CopyMetricWarnings"},
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
			_jsii_.MemberProperty{JsiiProperty: "x", GoGetter: "X"},
			_jsii_.MemberProperty{JsiiProperty: "y", GoGetter: "Y"},
		},
		func() interface{} {
			j := jsiiProxy_ConcreteWidget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWidget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CreateAlarmOptions",
		reflect.TypeOf((*CreateAlarmOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.CustomWidget",
		reflect.TypeOf((*CustomWidget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "copyMetricWarnings", GoMethod: "CopyMetricWarnings"},
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
			_jsii_.MemberProperty{JsiiProperty: "x", GoGetter: "X"},
			_jsii_.MemberProperty{JsiiProperty: "y", GoGetter: "Y"},
		},
		func() interface{} {
			j := jsiiProxy_CustomWidget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConcreteWidget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.CustomWidgetProps",
		reflect.TypeOf((*CustomWidgetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.Dashboard",
		reflect.TypeOf((*Dashboard)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVariable", GoMethod: "AddVariable"},
			_jsii_.MemberMethod{JsiiMethod: "addWidgets", GoMethod: "AddWidgets"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "dashboardArn", GoGetter: "DashboardArn"},
			_jsii_.MemberProperty{JsiiProperty: "dashboardName", GoGetter: "DashboardName"},
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
			j := jsiiProxy_Dashboard{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.DashboardProps",
		reflect.TypeOf((*DashboardProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.DashboardVariable",
		reflect.TypeOf((*DashboardVariable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
		},
		func() interface{} {
			j := jsiiProxy_DashboardVariable{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVariable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.DashboardVariableOptions",
		reflect.TypeOf((*DashboardVariableOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.DefaultValue",
		reflect.TypeOf((*DefaultValue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "val", GoGetter: "Val"},
		},
		func() interface{} {
			return &jsiiProxy_DefaultValue{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.Dimension",
		reflect.TypeOf((*Dimension)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.GaugeWidget",
		reflect.TypeOf((*GaugeWidget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMetric", GoMethod: "AddMetric"},
			_jsii_.MemberMethod{JsiiMethod: "copyMetricWarnings", GoMethod: "CopyMetricWarnings"},
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
			_jsii_.MemberProperty{JsiiProperty: "x", GoGetter: "X"},
			_jsii_.MemberProperty{JsiiProperty: "y", GoGetter: "Y"},
		},
		func() interface{} {
			j := jsiiProxy_GaugeWidget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConcreteWidget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.GaugeWidgetProps",
		reflect.TypeOf((*GaugeWidgetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.GraphWidget",
		reflect.TypeOf((*GraphWidget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addLeftMetric", GoMethod: "AddLeftMetric"},
			_jsii_.MemberMethod{JsiiMethod: "addRightMetric", GoMethod: "AddRightMetric"},
			_jsii_.MemberMethod{JsiiMethod: "copyMetricWarnings", GoMethod: "CopyMetricWarnings"},
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
			_jsii_.MemberProperty{JsiiProperty: "x", GoGetter: "X"},
			_jsii_.MemberProperty{JsiiProperty: "y", GoGetter: "Y"},
		},
		func() interface{} {
			j := jsiiProxy_GraphWidget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConcreteWidget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.GraphWidgetProps",
		reflect.TypeOf((*GraphWidgetProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.GraphWidgetView",
		reflect.TypeOf((*GraphWidgetView)(nil)).Elem(),
		map[string]interface{}{
			"TIME_SERIES": GraphWidgetView_TIME_SERIES,
			"BAR": GraphWidgetView_BAR,
			"PIE": GraphWidgetView_PIE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.HorizontalAnnotation",
		reflect.TypeOf((*HorizontalAnnotation)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudwatch.IAlarm",
		reflect.TypeOf((*IAlarm)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarmArn", GoGetter: "AlarmArn"},
			_jsii_.MemberProperty{JsiiProperty: "alarmName", GoGetter: "AlarmName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "renderAlarmRule", GoMethod: "RenderAlarmRule"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAlarm{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAlarmRule)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudwatch.IAlarmAction",
		reflect.TypeOf((*IAlarmAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IAlarmAction{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudwatch.IAlarmRule",
		reflect.TypeOf((*IAlarmRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "renderAlarmRule", GoMethod: "RenderAlarmRule"},
		},
		func() interface{} {
			return &jsiiProxy_IAlarmRule{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudwatch.IMetric",
		reflect.TypeOf((*IMetric)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toMetricConfig", GoMethod: "ToMetricConfig"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
		},
		func() interface{} {
			return &jsiiProxy_IMetric{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudwatch.IVariable",
		reflect.TypeOf((*IVariable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
		},
		func() interface{} {
			return &jsiiProxy_IVariable{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cloudwatch.IWidget",
		reflect.TypeOf((*IWidget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
		},
		func() interface{} {
			return &jsiiProxy_IWidget{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.LegendPosition",
		reflect.TypeOf((*LegendPosition)(nil)).Elem(),
		map[string]interface{}{
			"BOTTOM": LegendPosition_BOTTOM,
			"RIGHT": LegendPosition_RIGHT,
			"HIDDEN": LegendPosition_HIDDEN,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.LogQueryVisualizationType",
		reflect.TypeOf((*LogQueryVisualizationType)(nil)).Elem(),
		map[string]interface{}{
			"TABLE": LogQueryVisualizationType_TABLE,
			"LINE": LogQueryVisualizationType_LINE,
			"STACKEDAREA": LogQueryVisualizationType_STACKEDAREA,
			"BAR": LogQueryVisualizationType_BAR,
			"PIE": LogQueryVisualizationType_PIE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.LogQueryWidget",
		reflect.TypeOf((*LogQueryWidget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "copyMetricWarnings", GoMethod: "CopyMetricWarnings"},
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
			_jsii_.MemberProperty{JsiiProperty: "x", GoGetter: "X"},
			_jsii_.MemberProperty{JsiiProperty: "y", GoGetter: "Y"},
		},
		func() interface{} {
			j := jsiiProxy_LogQueryWidget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConcreteWidget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.LogQueryWidgetProps",
		reflect.TypeOf((*LogQueryWidgetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.MathExpression",
		reflect.TypeOf((*MathExpression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "color", GoGetter: "Color"},
			_jsii_.MemberMethod{JsiiMethod: "createAlarm", GoMethod: "CreateAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "expression", GoGetter: "Expression"},
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "searchAccount", GoGetter: "SearchAccount"},
			_jsii_.MemberProperty{JsiiProperty: "searchRegion", GoGetter: "SearchRegion"},
			_jsii_.MemberMethod{JsiiMethod: "toMetricConfig", GoMethod: "ToMetricConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "usingMetrics", GoGetter: "UsingMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_MathExpression{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMetric)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.MathExpressionOptions",
		reflect.TypeOf((*MathExpressionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.MathExpressionProps",
		reflect.TypeOf((*MathExpressionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.Metric",
		reflect.TypeOf((*Metric)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberMethod{JsiiMethod: "attachTo", GoMethod: "AttachTo"},
			_jsii_.MemberProperty{JsiiProperty: "color", GoGetter: "Color"},
			_jsii_.MemberMethod{JsiiMethod: "createAlarm", GoMethod: "CreateAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "dimensions", GoGetter: "Dimensions"},
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "metricName", GoGetter: "MetricName"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberMethod{JsiiMethod: "toMetricConfig", GoMethod: "ToMetricConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Metric{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMetric)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.MetricConfig",
		reflect.TypeOf((*MetricConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.MetricExpressionConfig",
		reflect.TypeOf((*MetricExpressionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.MetricOptions",
		reflect.TypeOf((*MetricOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.MetricProps",
		reflect.TypeOf((*MetricProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.MetricStatConfig",
		reflect.TypeOf((*MetricStatConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.MetricWidgetProps",
		reflect.TypeOf((*MetricWidgetProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.PeriodOverride",
		reflect.TypeOf((*PeriodOverride)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": PeriodOverride_AUTO,
			"INHERIT": PeriodOverride_INHERIT,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.Row",
		reflect.TypeOf((*Row)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addWidget", GoMethod: "AddWidget"},
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "widgets", GoGetter: "Widgets"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
		},
		func() interface{} {
			j := jsiiProxy_Row{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWidget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.SearchComponents",
		reflect.TypeOf((*SearchComponents)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.Shading",
		reflect.TypeOf((*Shading)(nil)).Elem(),
		map[string]interface{}{
			"NONE": Shading_NONE,
			"ABOVE": Shading_ABOVE,
			"BELOW": Shading_BELOW,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.SingleValueWidget",
		reflect.TypeOf((*SingleValueWidget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "copyMetricWarnings", GoMethod: "CopyMetricWarnings"},
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
			_jsii_.MemberProperty{JsiiProperty: "x", GoGetter: "X"},
			_jsii_.MemberProperty{JsiiProperty: "y", GoGetter: "Y"},
		},
		func() interface{} {
			j := jsiiProxy_SingleValueWidget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConcreteWidget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.SingleValueWidgetProps",
		reflect.TypeOf((*SingleValueWidgetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.Spacer",
		reflect.TypeOf((*Spacer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
		},
		func() interface{} {
			j := jsiiProxy_Spacer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWidget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.SpacerProps",
		reflect.TypeOf((*SpacerProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.Statistic",
		reflect.TypeOf((*Statistic)(nil)).Elem(),
		map[string]interface{}{
			"SAMPLE_COUNT": Statistic_SAMPLE_COUNT,
			"AVERAGE": Statistic_AVERAGE,
			"SUM": Statistic_SUM,
			"MINIMUM": Statistic_MINIMUM,
			"MAXIMUM": Statistic_MAXIMUM,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.Stats",
		reflect.TypeOf((*Stats)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Stats{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.TableLayout",
		reflect.TypeOf((*TableLayout)(nil)).Elem(),
		map[string]interface{}{
			"HORIZONTAL": TableLayout_HORIZONTAL,
			"VERTICAL": TableLayout_VERTICAL,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.TableSummaryColumn",
		reflect.TypeOf((*TableSummaryColumn)(nil)).Elem(),
		map[string]interface{}{
			"MINIMUM": TableSummaryColumn_MINIMUM,
			"MAXIMUM": TableSummaryColumn_MAXIMUM,
			"SUM": TableSummaryColumn_SUM,
			"AVERAGE": TableSummaryColumn_AVERAGE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.TableSummaryProps",
		reflect.TypeOf((*TableSummaryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.TableThreshold",
		reflect.TypeOf((*TableThreshold)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
		},
		func() interface{} {
			return &jsiiProxy_TableThreshold{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.TableWidget",
		reflect.TypeOf((*TableWidget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMetric", GoMethod: "AddMetric"},
			_jsii_.MemberMethod{JsiiMethod: "copyMetricWarnings", GoMethod: "CopyMetricWarnings"},
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
			_jsii_.MemberProperty{JsiiProperty: "x", GoGetter: "X"},
			_jsii_.MemberProperty{JsiiProperty: "y", GoGetter: "Y"},
		},
		func() interface{} {
			j := jsiiProxy_TableWidget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConcreteWidget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.TableWidgetProps",
		reflect.TypeOf((*TableWidgetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.TextWidget",
		reflect.TypeOf((*TextWidget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "copyMetricWarnings", GoMethod: "CopyMetricWarnings"},
			_jsii_.MemberProperty{JsiiProperty: "height", GoGetter: "Height"},
			_jsii_.MemberMethod{JsiiMethod: "position", GoMethod: "Position"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "warnings", GoGetter: "Warnings"},
			_jsii_.MemberProperty{JsiiProperty: "warningsV2", GoGetter: "WarningsV2"},
			_jsii_.MemberProperty{JsiiProperty: "width", GoGetter: "Width"},
			_jsii_.MemberProperty{JsiiProperty: "x", GoGetter: "X"},
			_jsii_.MemberProperty{JsiiProperty: "y", GoGetter: "Y"},
		},
		func() interface{} {
			j := jsiiProxy_TextWidget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConcreteWidget)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.TextWidgetBackground",
		reflect.TypeOf((*TextWidgetBackground)(nil)).Elem(),
		map[string]interface{}{
			"SOLID": TextWidgetBackground_SOLID,
			"TRANSPARENT": TextWidgetBackground_TRANSPARENT,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.TextWidgetProps",
		reflect.TypeOf((*TextWidgetProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.TreatMissingData",
		reflect.TypeOf((*TreatMissingData)(nil)).Elem(),
		map[string]interface{}{
			"BREACHING": TreatMissingData_BREACHING,
			"NOT_BREACHING": TreatMissingData_NOT_BREACHING,
			"IGNORE": TreatMissingData_IGNORE,
			"MISSING": TreatMissingData_MISSING,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.Unit",
		reflect.TypeOf((*Unit)(nil)).Elem(),
		map[string]interface{}{
			"SECONDS": Unit_SECONDS,
			"MICROSECONDS": Unit_MICROSECONDS,
			"MILLISECONDS": Unit_MILLISECONDS,
			"BYTES": Unit_BYTES,
			"KILOBYTES": Unit_KILOBYTES,
			"MEGABYTES": Unit_MEGABYTES,
			"GIGABYTES": Unit_GIGABYTES,
			"TERABYTES": Unit_TERABYTES,
			"BITS": Unit_BITS,
			"KILOBITS": Unit_KILOBITS,
			"MEGABITS": Unit_MEGABITS,
			"GIGABITS": Unit_GIGABITS,
			"TERABITS": Unit_TERABITS,
			"PERCENT": Unit_PERCENT,
			"COUNT": Unit_COUNT,
			"BYTES_PER_SECOND": Unit_BYTES_PER_SECOND,
			"KILOBYTES_PER_SECOND": Unit_KILOBYTES_PER_SECOND,
			"MEGABYTES_PER_SECOND": Unit_MEGABYTES_PER_SECOND,
			"GIGABYTES_PER_SECOND": Unit_GIGABYTES_PER_SECOND,
			"TERABYTES_PER_SECOND": Unit_TERABYTES_PER_SECOND,
			"BITS_PER_SECOND": Unit_BITS_PER_SECOND,
			"KILOBITS_PER_SECOND": Unit_KILOBITS_PER_SECOND,
			"MEGABITS_PER_SECOND": Unit_MEGABITS_PER_SECOND,
			"GIGABITS_PER_SECOND": Unit_GIGABITS_PER_SECOND,
			"TERABITS_PER_SECOND": Unit_TERABITS_PER_SECOND,
			"COUNT_PER_SECOND": Unit_COUNT_PER_SECOND,
			"NONE": Unit_NONE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudwatch.Values",
		reflect.TypeOf((*Values)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
		},
		func() interface{} {
			return &jsiiProxy_Values{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.VariableInputType",
		reflect.TypeOf((*VariableInputType)(nil)).Elem(),
		map[string]interface{}{
			"INPUT": VariableInputType_INPUT,
			"RADIO": VariableInputType_RADIO,
			"SELECT": VariableInputType_SELECT,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.VariableType",
		reflect.TypeOf((*VariableType)(nil)).Elem(),
		map[string]interface{}{
			"PROPERTY": VariableType_PROPERTY,
			"PATTERN": VariableType_PATTERN,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.VariableValue",
		reflect.TypeOf((*VariableValue)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.VerticalAnnotation",
		reflect.TypeOf((*VerticalAnnotation)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudwatch.VerticalShading",
		reflect.TypeOf((*VerticalShading)(nil)).Elem(),
		map[string]interface{}{
			"NONE": VerticalShading_NONE,
			"BEFORE": VerticalShading_BEFORE,
			"AFTER": VerticalShading_AFTER,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudwatch.YAxisProps",
		reflect.TypeOf((*YAxisProps)(nil)).Elem(),
	)
}
