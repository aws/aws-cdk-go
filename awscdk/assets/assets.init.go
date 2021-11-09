package assets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.assets.CopyOptions",
		reflect.TypeOf((*CopyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.assets.FingerprintOptions",
		reflect.TypeOf((*FingerprintOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.assets.FollowMode",
		reflect.TypeOf((*FollowMode)(nil)).Elem(),
		map[string]interface{}{
			"NEVER": FollowMode_NEVER,
			"ALWAYS": FollowMode_ALWAYS,
			"EXTERNAL": FollowMode_EXTERNAL,
			"BLOCK_EXTERNAL": FollowMode_BLOCK_EXTERNAL,
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.assets.IAsset",
		reflect.TypeOf((*IAsset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "sourceHash", GoGetter: "SourceHash"},
		},
		func() interface{} {
			return &jsiiProxy_IAsset{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.assets.Staging",
		reflect.TypeOf((*Staging)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absoluteStagedPath", GoGetter: "AbsoluteStagedPath"},
			_jsii_.MemberProperty{JsiiProperty: "assetHash", GoGetter: "AssetHash"},
			_jsii_.MemberProperty{JsiiProperty: "isArchive", GoGetter: "IsArchive"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "packaging", GoGetter: "Packaging"},
			_jsii_.MemberMethod{JsiiMethod: "relativeStagedPath", GoMethod: "RelativeStagedPath"},
			_jsii_.MemberProperty{JsiiProperty: "sourceHash", GoGetter: "SourceHash"},
			_jsii_.MemberProperty{JsiiProperty: "sourcePath", GoGetter: "SourcePath"},
			_jsii_.MemberProperty{JsiiProperty: "stagedPath", GoGetter: "StagedPath"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Staging{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkAssetStaging)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.assets.StagingProps",
		reflect.TypeOf((*StagingProps)(nil)).Elem(),
	)
}
