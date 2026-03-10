package awss3mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_s3.mixins.BucketAutoDeleteObjects",
		reflect.TypeOf((*BucketAutoDeleteObjects)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_BucketAutoDeleteObjects{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_s3.mixins.BucketBlockPublicAccess",
		reflect.TypeOf((*BucketBlockPublicAccess)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_BucketBlockPublicAccess{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_s3.mixins.BucketPolicyStatements",
		reflect.TypeOf((*BucketPolicyStatements)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_BucketPolicyStatements{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_s3.mixins.BucketVersioning",
		reflect.TypeOf((*BucketVersioning)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_BucketVersioning{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			return &j
		},
	)
}
