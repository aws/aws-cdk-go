package awssesactions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ses_actions.AddHeader",
		reflect.TypeOf((*AddHeader)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_AddHeader{}
			_jsii_.InitJsiiProxy(&j.Type__awssesIReceiptRuleAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ses_actions.AddHeaderProps",
		reflect.TypeOf((*AddHeaderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ses_actions.Bounce",
		reflect.TypeOf((*Bounce)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_Bounce{}
			_jsii_.InitJsiiProxy(&j.Type__awssesIReceiptRuleAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ses_actions.BounceProps",
		reflect.TypeOf((*BounceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		reflect.TypeOf((*BounceTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			return &jsiiProxy_BounceTemplate{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ses_actions.BounceTemplateProps",
		reflect.TypeOf((*BounceTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ses_actions.EmailEncoding",
		reflect.TypeOf((*EmailEncoding)(nil)).Elem(),
		map[string]interface{}{
			"BASE64": EmailEncoding_BASE64,
			"UTF8": EmailEncoding_UTF8,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ses_actions.Lambda",
		reflect.TypeOf((*Lambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_Lambda{}
			_jsii_.InitJsiiProxy(&j.Type__awssesIReceiptRuleAction)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ses_actions.LambdaInvocationType",
		reflect.TypeOf((*LambdaInvocationType)(nil)).Elem(),
		map[string]interface{}{
			"EVENT": LambdaInvocationType_EVENT,
			"REQUEST_RESPONSE": LambdaInvocationType_REQUEST_RESPONSE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ses_actions.LambdaProps",
		reflect.TypeOf((*LambdaProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ses_actions.S3",
		reflect.TypeOf((*S3)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3{}
			_jsii_.InitJsiiProxy(&j.Type__awssesIReceiptRuleAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ses_actions.S3Props",
		reflect.TypeOf((*S3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ses_actions.Sns",
		reflect.TypeOf((*Sns)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_Sns{}
			_jsii_.InitJsiiProxy(&j.Type__awssesIReceiptRuleAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ses_actions.SnsProps",
		reflect.TypeOf((*SnsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ses_actions.Stop",
		reflect.TypeOf((*Stop)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_Stop{}
			_jsii_.InitJsiiProxy(&j.Type__awssesIReceiptRuleAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ses_actions.StopProps",
		reflect.TypeOf((*StopProps)(nil)).Elem(),
	)
}
