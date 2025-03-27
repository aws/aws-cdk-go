// This module is deprecated. All constructs are now available under aws-cdk-lib/aws-scheduler.
package awscdkscheduleralpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-alpha.ContextAttribute",
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
		"@aws-cdk/aws-scheduler-alpha.CronOptionsWithTimezone",
		reflect.TypeOf((*CronOptionsWithTimezone)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-alpha.Group",
		reflect.TypeOf((*Group)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "groupArn", GoGetter: "GroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "groupName", GoGetter: "GroupName"},
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
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Group{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-scheduler-alpha.GroupProps",
		reflect.TypeOf((*GroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-scheduler-alpha.IGroup",
		reflect.TypeOf((*IGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantDeleteSchedules", GoMethod: "GrantDeleteSchedules"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadSchedules", GoMethod: "GrantReadSchedules"},
			_jsii_.MemberMethod{JsiiMethod: "grantWriteSchedules", GoMethod: "GrantWriteSchedules"},
			_jsii_.MemberProperty{JsiiProperty: "groupArn", GoGetter: "GroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "groupName", GoGetter: "GroupName"},
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
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-scheduler-alpha.ISchedule",
		reflect.TypeOf((*ISchedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "group", GoGetter: "Group"},
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
		"@aws-cdk/aws-scheduler-alpha.IScheduleGroup",
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
		"@aws-cdk/aws-scheduler-alpha.IScheduleTarget",
		reflect.TypeOf((*IScheduleTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IScheduleTarget{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-alpha.Schedule",
		reflect.TypeOf((*Schedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "group", GoGetter: "Group"},
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
		"@aws-cdk/aws-scheduler-alpha.ScheduleExpression",
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
		"@aws-cdk/aws-scheduler-alpha.ScheduleGroup",
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
		"@aws-cdk/aws-scheduler-alpha.ScheduleGroupProps",
		reflect.TypeOf((*ScheduleGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-scheduler-alpha.ScheduleProps",
		reflect.TypeOf((*ScheduleProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-scheduler-alpha.ScheduleTargetConfig",
		reflect.TypeOf((*ScheduleTargetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-alpha.ScheduleTargetInput",
		reflect.TypeOf((*ScheduleTargetInput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ScheduleTargetInput{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-scheduler-alpha.TimeWindow",
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
