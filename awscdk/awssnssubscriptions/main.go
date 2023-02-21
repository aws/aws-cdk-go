package awssnssubscriptions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns_subscriptions.EmailSubscription",
		reflect.TypeOf((*EmailSubscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_EmailSubscription{}
			_jsii_.InitJsiiProxy(&j.Type__awssnsITopicSubscription)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns_subscriptions.EmailSubscriptionProps",
		reflect.TypeOf((*EmailSubscriptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns_subscriptions.LambdaSubscription",
		reflect.TypeOf((*LambdaSubscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaSubscription{}
			_jsii_.InitJsiiProxy(&j.Type__awssnsITopicSubscription)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns_subscriptions.LambdaSubscriptionProps",
		reflect.TypeOf((*LambdaSubscriptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns_subscriptions.SmsSubscription",
		reflect.TypeOf((*SmsSubscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SmsSubscription{}
			_jsii_.InitJsiiProxy(&j.Type__awssnsITopicSubscription)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns_subscriptions.SmsSubscriptionProps",
		reflect.TypeOf((*SmsSubscriptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns_subscriptions.SqsSubscription",
		reflect.TypeOf((*SqsSubscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SqsSubscription{}
			_jsii_.InitJsiiProxy(&j.Type__awssnsITopicSubscription)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns_subscriptions.SqsSubscriptionProps",
		reflect.TypeOf((*SqsSubscriptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns_subscriptions.SubscriptionProps",
		reflect.TypeOf((*SubscriptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_sns_subscriptions.UrlSubscription",
		reflect.TypeOf((*UrlSubscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_UrlSubscription{}
			_jsii_.InitJsiiProxy(&j.Type__awssnsITopicSubscription)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_sns_subscriptions.UrlSubscriptionProps",
		reflect.TypeOf((*UrlSubscriptionProps)(nil)).Elem(),
	)
}
