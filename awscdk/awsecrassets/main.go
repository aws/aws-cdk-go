package awsecrassets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ecr_assets.DockerImageAsset",
		reflect.TypeOf((*DockerImageAsset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addResourceMetadata", GoMethod: "AddResourceMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "assetHash", GoGetter: "AssetHash"},
			_jsii_.MemberProperty{JsiiProperty: "imageTag", GoGetter: "ImageTag"},
			_jsii_.MemberProperty{JsiiProperty: "imageUri", GoGetter: "ImageUri"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "repository", GoGetter: "Repository"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DockerImageAsset{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ecr_assets.DockerImageAssetInvalidationOptions",
		reflect.TypeOf((*DockerImageAssetInvalidationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ecr_assets.DockerImageAssetOptions",
		reflect.TypeOf((*DockerImageAssetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ecr_assets.DockerImageAssetProps",
		reflect.TypeOf((*DockerImageAssetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ecr_assets.NetworkMode",
		reflect.TypeOf((*NetworkMode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkMode{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ecr_assets.Platform",
		reflect.TypeOf((*Platform)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "platform", GoGetter: "Platform"},
		},
		func() interface{} {
			return &jsiiProxy_Platform{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ecr_assets.TarballImageAsset",
		reflect.TypeOf((*TarballImageAsset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assetHash", GoGetter: "AssetHash"},
			_jsii_.MemberProperty{JsiiProperty: "imageTag", GoGetter: "ImageTag"},
			_jsii_.MemberProperty{JsiiProperty: "imageUri", GoGetter: "ImageUri"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "repository", GoGetter: "Repository"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TarballImageAsset{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ecr_assets.TarballImageAssetProps",
		reflect.TypeOf((*TarballImageAssetProps)(nil)).Elem(),
	)
}
