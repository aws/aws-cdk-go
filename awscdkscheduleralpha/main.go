// The CDK Construct Library for Amazon Scheduler
package awscdkscheduleralpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-scheduler-alpha.CronOptionsWithTimezone",
		reflect.TypeOf((*CronOptionsWithTimezone)(nil)).Elem(),
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
}
