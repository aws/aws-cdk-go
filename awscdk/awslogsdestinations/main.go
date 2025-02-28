package awslogsdestinations

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_logs_destinations.KinesisDestination",
		reflect.TypeOf((*KinesisDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisDestination{}
			_jsii_.InitJsiiProxy(&j.Type__awslogsILogSubscriptionDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_logs_destinations.KinesisDestinationProps",
		reflect.TypeOf((*KinesisDestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_logs_destinations.LambdaDestination",
		reflect.TypeOf((*LambdaDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaDestination{}
			_jsii_.InitJsiiProxy(&j.Type__awslogsILogSubscriptionDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_logs_destinations.LambdaDestinationOptions",
		reflect.TypeOf((*LambdaDestinationOptions)(nil)).Elem(),
	)
}
