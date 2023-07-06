// The CDK Construct Library for Amazon Scheduler
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
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-scheduler-alpha.ISchedule",
		reflect.TypeOf((*ISchedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ISchedule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
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
		"@aws-cdk/aws-scheduler-alpha.ScheduleTargetInput",
		reflect.TypeOf((*ScheduleTargetInput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ScheduleTargetInput{}
		},
	)
}
