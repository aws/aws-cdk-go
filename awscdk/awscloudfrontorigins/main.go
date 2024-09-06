package awscloudfrontorigins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront_origins.FunctionUrlOrigin",
		reflect.TypeOf((*FunctionUrlOrigin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "renderCustomOriginConfig", GoMethod: "RenderCustomOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "renderS3OriginConfig", GoMethod: "RenderS3OriginConfig"},
		},
		func() interface{} {
			j := jsiiProxy_FunctionUrlOrigin{}
			_jsii_.InitJsiiProxy(&j.Type__awscloudfrontOriginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront_origins.FunctionUrlOriginProps",
		reflect.TypeOf((*FunctionUrlOriginProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront_origins.HttpOrigin",
		reflect.TypeOf((*HttpOrigin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "renderCustomOriginConfig", GoMethod: "RenderCustomOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "renderS3OriginConfig", GoMethod: "RenderS3OriginConfig"},
		},
		func() interface{} {
			j := jsiiProxy_HttpOrigin{}
			_jsii_.InitJsiiProxy(&j.Type__awscloudfrontOriginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront_origins.HttpOriginProps",
		reflect.TypeOf((*HttpOriginProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront_origins.LoadBalancerV2Origin",
		reflect.TypeOf((*LoadBalancerV2Origin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "renderCustomOriginConfig", GoMethod: "RenderCustomOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "renderS3OriginConfig", GoMethod: "RenderS3OriginConfig"},
		},
		func() interface{} {
			j := jsiiProxy_LoadBalancerV2Origin{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_HttpOrigin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront_origins.LoadBalancerV2OriginProps",
		reflect.TypeOf((*LoadBalancerV2OriginProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront_origins.OriginGroup",
		reflect.TypeOf((*OriginGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_OriginGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscloudfrontIOrigin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront_origins.OriginGroupProps",
		reflect.TypeOf((*OriginGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront_origins.RestApiOrigin",
		reflect.TypeOf((*RestApiOrigin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "renderCustomOriginConfig", GoMethod: "RenderCustomOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "renderS3OriginConfig", GoMethod: "RenderS3OriginConfig"},
		},
		func() interface{} {
			j := jsiiProxy_RestApiOrigin{}
			_jsii_.InitJsiiProxy(&j.Type__awscloudfrontOriginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront_origins.RestApiOriginProps",
		reflect.TypeOf((*RestApiOriginProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront_origins.S3BucketOrigin",
		reflect.TypeOf((*S3BucketOrigin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "renderCustomOriginConfig", GoMethod: "RenderCustomOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "renderS3OriginConfig", GoMethod: "RenderS3OriginConfig"},
		},
		func() interface{} {
			j := jsiiProxy_S3BucketOrigin{}
			_jsii_.InitJsiiProxy(&j.Type__awscloudfrontOriginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront_origins.S3BucketOriginBaseProps",
		reflect.TypeOf((*S3BucketOriginBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront_origins.S3BucketOriginWithOACProps",
		reflect.TypeOf((*S3BucketOriginWithOACProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront_origins.S3BucketOriginWithOAIProps",
		reflect.TypeOf((*S3BucketOriginWithOAIProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront_origins.S3Origin",
		reflect.TypeOf((*S3Origin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3Origin{}
			_jsii_.InitJsiiProxy(&j.Type__awscloudfrontIOrigin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront_origins.S3OriginProps",
		reflect.TypeOf((*S3OriginProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudfront_origins.S3StaticWebsiteOrigin",
		reflect.TypeOf((*S3StaticWebsiteOrigin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "renderCustomOriginConfig", GoMethod: "RenderCustomOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "renderS3OriginConfig", GoMethod: "RenderS3OriginConfig"},
		},
		func() interface{} {
			j := jsiiProxy_S3StaticWebsiteOrigin{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_HttpOrigin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudfront_origins.S3StaticWebsiteOriginProps",
		reflect.TypeOf((*S3StaticWebsiteOriginProps)(nil)).Elem(),
	)
}
