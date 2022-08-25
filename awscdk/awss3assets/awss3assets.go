package awss3assets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_s3_assets.Asset",
		reflect.TypeOf((*Asset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addResourceMetadata", GoMethod: "AddResourceMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "assetHash", GoGetter: "AssetHash"},
			_jsii_.MemberProperty{JsiiProperty: "assetPath", GoGetter: "AssetPath"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "httpUrl", GoGetter: "HttpUrl"},
			_jsii_.MemberProperty{JsiiProperty: "isFile", GoGetter: "IsFile"},
			_jsii_.MemberProperty{JsiiProperty: "isZipArchive", GoGetter: "IsZipArchive"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketName", GoGetter: "S3BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "s3ObjectKey", GoGetter: "S3ObjectKey"},
			_jsii_.MemberProperty{JsiiProperty: "s3ObjectUrl", GoGetter: "S3ObjectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Asset{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAsset)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_s3_assets.AssetOptions",
		reflect.TypeOf((*AssetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_s3_assets.AssetProps",
		reflect.TypeOf((*AssetProps)(nil)).Elem(),
	)
}
