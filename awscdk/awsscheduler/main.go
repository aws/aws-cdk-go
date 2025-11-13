package awsscheduler

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler.CfnSchedule",
		reflect.TypeOf((*CfnSchedule)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "endDate", GoGetter: "EndDate"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flexibleTimeWindow", GoGetter: "FlexibleTimeWindow"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "groupName", GoGetter: "GroupName"},
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
			_jsii_.MemberProperty{JsiiProperty: "scheduleExpression", GoGetter: "ScheduleExpression"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleExpressionTimezone", GoGetter: "ScheduleExpressionTimezone"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleRef", GoGetter: "ScheduleRef"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "startDate", GoGetter: "StartDate"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchedule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsschedulerIScheduleRef)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.AwsVpcConfigurationProperty",
		reflect.TypeOf((*CfnSchedule_AwsVpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.CapacityProviderStrategyItemProperty",
		reflect.TypeOf((*CfnSchedule_CapacityProviderStrategyItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.DeadLetterConfigProperty",
		reflect.TypeOf((*CfnSchedule_DeadLetterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.EcsParametersProperty",
		reflect.TypeOf((*CfnSchedule_EcsParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.EventBridgeParametersProperty",
		reflect.TypeOf((*CfnSchedule_EventBridgeParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.FlexibleTimeWindowProperty",
		reflect.TypeOf((*CfnSchedule_FlexibleTimeWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.KinesisParametersProperty",
		reflect.TypeOf((*CfnSchedule_KinesisParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnSchedule_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.PlacementConstraintProperty",
		reflect.TypeOf((*CfnSchedule_PlacementConstraintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.PlacementStrategyProperty",
		reflect.TypeOf((*CfnSchedule_PlacementStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.RetryPolicyProperty",
		reflect.TypeOf((*CfnSchedule_RetryPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.SageMakerPipelineParameterProperty",
		reflect.TypeOf((*CfnSchedule_SageMakerPipelineParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.SageMakerPipelineParametersProperty",
		reflect.TypeOf((*CfnSchedule_SageMakerPipelineParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.SqsParametersProperty",
		reflect.TypeOf((*CfnSchedule_SqsParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnSchedule.TargetProperty",
		reflect.TypeOf((*CfnSchedule_TargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler.CfnScheduleGroup",
		reflect.TypeOf((*CfnScheduleGroup)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrLastModificationDate", GoGetter: "AttrLastModificationDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
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
			_jsii_.MemberProperty{JsiiProperty: "scheduleGroupRef", GoGetter: "ScheduleGroupRef"},
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
			j := jsiiProxy_CfnScheduleGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsschedulerIScheduleGroupRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnScheduleGroupProps",
		reflect.TypeOf((*CfnScheduleGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CfnScheduleProps",
		reflect.TypeOf((*CfnScheduleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler.ContextAttribute",
		reflect.TypeOf((*ContextAttribute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_ContextAttribute{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.CronOptionsWithTimezone",
		reflect.TypeOf((*CronOptionsWithTimezone)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_scheduler.ISchedule",
		reflect.TypeOf((*ISchedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleArn", GoGetter: "ScheduleArn"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleGroup", GoGetter: "ScheduleGroup"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleName", GoGetter: "ScheduleName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ISchedule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_scheduler.IScheduleGroup",
		reflect.TypeOf((*IScheduleGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantDeleteSchedules", GoMethod: "GrantDeleteSchedules"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadSchedules", GoMethod: "GrantReadSchedules"},
			_jsii_.MemberMethod{JsiiMethod: "grantWriteSchedules", GoMethod: "GrantWriteSchedules"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAttempts", GoMethod: "MetricAttempts"},
			_jsii_.MemberMethod{JsiiMethod: "metricDropped", GoMethod: "MetricDropped"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailedToBeSentToDLQ", GoMethod: "MetricFailedToBeSentToDLQ"},
			_jsii_.MemberMethod{JsiiMethod: "metricSentToDLQ", GoMethod: "MetricSentToDLQ"},
			_jsii_.MemberMethod{JsiiMethod: "metricSentToDLQTruncated", GoMethod: "MetricSentToDLQTruncated"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetErrors", GoMethod: "MetricTargetErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetThrottled", GoMethod: "MetricTargetThrottled"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottled", GoMethod: "MetricThrottled"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleGroupArn", GoGetter: "ScheduleGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleGroupName", GoGetter: "ScheduleGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IScheduleGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_scheduler.IScheduleTarget",
		reflect.TypeOf((*IScheduleTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IScheduleTarget{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler.Schedule",
		reflect.TypeOf((*Schedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleArn", GoGetter: "ScheduleArn"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleGroup", GoGetter: "ScheduleGroup"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleName", GoGetter: "ScheduleName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Schedule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISchedule)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler.ScheduleExpression",
		reflect.TypeOf((*ScheduleExpression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "expressionString", GoGetter: "ExpressionString"},
			_jsii_.MemberProperty{JsiiProperty: "timeZone", GoGetter: "TimeZone"},
		},
		func() interface{} {
			return &jsiiProxy_ScheduleExpression{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler.ScheduleGroup",
		reflect.TypeOf((*ScheduleGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantDeleteSchedules", GoMethod: "GrantDeleteSchedules"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadSchedules", GoMethod: "GrantReadSchedules"},
			_jsii_.MemberMethod{JsiiMethod: "grantWriteSchedules", GoMethod: "GrantWriteSchedules"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAttempts", GoMethod: "MetricAttempts"},
			_jsii_.MemberMethod{JsiiMethod: "metricDropped", GoMethod: "MetricDropped"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailedToBeSentToDLQ", GoMethod: "MetricFailedToBeSentToDLQ"},
			_jsii_.MemberMethod{JsiiMethod: "metricSentToDLQ", GoMethod: "MetricSentToDLQ"},
			_jsii_.MemberMethod{JsiiMethod: "metricSentToDLQTruncated", GoMethod: "MetricSentToDLQTruncated"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetErrors", GoMethod: "MetricTargetErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetThrottled", GoMethod: "MetricTargetThrottled"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottled", GoMethod: "MetricThrottled"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleGroupArn", GoGetter: "ScheduleGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleGroupName", GoGetter: "ScheduleGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ScheduleGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IScheduleGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.ScheduleGroupProps",
		reflect.TypeOf((*ScheduleGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.ScheduleProps",
		reflect.TypeOf((*ScheduleProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_scheduler.ScheduleTargetConfig",
		reflect.TypeOf((*ScheduleTargetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler.ScheduleTargetInput",
		reflect.TypeOf((*ScheduleTargetInput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ScheduleTargetInput{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_scheduler.TimeWindow",
		reflect.TypeOf((*TimeWindow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "maxWindow", GoGetter: "MaxWindow"},
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
		},
		func() interface{} {
			return &jsiiProxy_TimeWindow{}
		},
	)
}
