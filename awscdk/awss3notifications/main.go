package awss3notifications

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_s3_notifications.LambdaDestination",
		reflect.TypeOf((*LambdaDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaDestination{}
			_jsii_.InitJsiiProxy(&j.Type__awss3IBucketNotificationDestination)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_s3_notifications.SnsDestination",
		reflect.TypeOf((*SnsDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SnsDestination{}
			_jsii_.InitJsiiProxy(&j.Type__awss3IBucketNotificationDestination)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_s3_notifications.SqsDestination",
		reflect.TypeOf((*SqsDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SqsDestination{}
			_jsii_.InitJsiiProxy(&j.Type__awss3IBucketNotificationDestination)
			return &j
		},
	)
}
