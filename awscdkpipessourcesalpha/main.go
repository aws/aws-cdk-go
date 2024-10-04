// The CDK Construct Library for Amazon EventBridge Pipes Sources
package awscdkpipessourcesalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-sources-alpha.DynamoDBSource",
		reflect.TypeOf((*DynamoDBSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterTarget", GoGetter: "DeadLetterTarget"},
			_jsii_.MemberMethod{JsiiMethod: "getDeadLetterTargetArn", GoMethod: "GetDeadLetterTargetArn"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "sourceArn", GoGetter: "SourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceParameters", GoGetter: "SourceParameters"},
		},
		func() interface{} {
			j := jsiiProxy_DynamoDBSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_StreamSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-sources-alpha.DynamoDBSourceParameters",
		reflect.TypeOf((*DynamoDBSourceParameters)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-pipes-sources-alpha.DynamoDBStartingPosition",
		reflect.TypeOf((*DynamoDBStartingPosition)(nil)).Elem(),
		map[string]interface{}{
			"TRIM_HORIZON": DynamoDBStartingPosition_TRIM_HORIZON,
			"LATEST": DynamoDBStartingPosition_LATEST,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-sources-alpha.KinesisSource",
		reflect.TypeOf((*KinesisSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterTarget", GoGetter: "DeadLetterTarget"},
			_jsii_.MemberMethod{JsiiMethod: "getDeadLetterTargetArn", GoMethod: "GetDeadLetterTargetArn"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "sourceArn", GoGetter: "SourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceParameters", GoGetter: "SourceParameters"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_StreamSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-sources-alpha.KinesisSourceParameters",
		reflect.TypeOf((*KinesisSourceParameters)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-pipes-sources-alpha.KinesisStartingPosition",
		reflect.TypeOf((*KinesisStartingPosition)(nil)).Elem(),
		map[string]interface{}{
			"TRIM_HORIZON": KinesisStartingPosition_TRIM_HORIZON,
			"LATEST": KinesisStartingPosition_LATEST,
			"AT_TIMESTAMP": KinesisStartingPosition_AT_TIMESTAMP,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-pipes-sources-alpha.OnPartialBatchItemFailure",
		reflect.TypeOf((*OnPartialBatchItemFailure)(nil)).Elem(),
		map[string]interface{}{
			"AUTOMATIC_BISECT": OnPartialBatchItemFailure_AUTOMATIC_BISECT,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-sources-alpha.SqsSource",
		reflect.TypeOf((*SqsSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "sourceArn", GoGetter: "SourceArn"},
		},
		func() interface{} {
			j := jsiiProxy_SqsSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-sources-alpha.SqsSourceParameters",
		reflect.TypeOf((*SqsSourceParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-pipes-sources-alpha.StreamSource",
		reflect.TypeOf((*StreamSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterTarget", GoGetter: "DeadLetterTarget"},
			_jsii_.MemberMethod{JsiiMethod: "getDeadLetterTargetArn", GoMethod: "GetDeadLetterTargetArn"},
			_jsii_.MemberMethod{JsiiMethod: "grantPush", GoMethod: "GrantPush"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "sourceArn", GoGetter: "SourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceParameters", GoGetter: "SourceParameters"},
		},
		func() interface{} {
			j := jsiiProxy_StreamSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkpipesalphaSourceWithDeadLetterTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-pipes-sources-alpha.StreamSourceParameters",
		reflect.TypeOf((*StreamSourceParameters)(nil)).Elem(),
	)
}
