package awsbedrock

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.FoundationModel",
		reflect.TypeOf((*FoundationModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "modelArn", GoGetter: "ModelArn"},
			_jsii_.MemberProperty{JsiiProperty: "modelId", GoGetter: "ModelId"},
		},
		func() interface{} {
			j := jsiiProxy_FoundationModel{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IModel)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.FoundationModelIdentifier",
		reflect.TypeOf((*FoundationModelIdentifier)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "modelId", GoGetter: "ModelId"},
		},
		func() interface{} {
			return &jsiiProxy_FoundationModelIdentifier{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_bedrock.IModel",
		reflect.TypeOf((*IModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "modelArn", GoGetter: "ModelArn"},
		},
		func() interface{} {
			return &jsiiProxy_IModel{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_bedrock.ProvisionedModel",
		reflect.TypeOf((*ProvisionedModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "modelArn", GoGetter: "ModelArn"},
		},
		func() interface{} {
			j := jsiiProxy_ProvisionedModel{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IModel)
			return &j
		},
	)
}
