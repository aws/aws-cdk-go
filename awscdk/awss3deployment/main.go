package awss3deployment

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_s3_deployment.BucketDeployment",
		reflect.TypeOf((*BucketDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSource", GoMethod: "AddSource"},
			_jsii_.MemberProperty{JsiiProperty: "deployedBucket", GoGetter: "DeployedBucket"},
			_jsii_.MemberProperty{JsiiProperty: "handlerRole", GoGetter: "HandlerRole"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "objectKeys", GoGetter: "ObjectKeys"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BucketDeployment{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_s3_deployment.BucketDeploymentProps",
		reflect.TypeOf((*BucketDeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		reflect.TypeOf((*CacheControl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CacheControl{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_s3_deployment.DeployTimeSubstitutedFile",
		reflect.TypeOf((*DeployTimeSubstitutedFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSource", GoMethod: "AddSource"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "deployedBucket", GoGetter: "DeployedBucket"},
			_jsii_.MemberProperty{JsiiProperty: "handlerRole", GoGetter: "HandlerRole"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "objectKey", GoGetter: "ObjectKey"},
			_jsii_.MemberProperty{JsiiProperty: "objectKeys", GoGetter: "ObjectKeys"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DeployTimeSubstitutedFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BucketDeployment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_s3_deployment.DeployTimeSubstitutedFileProps",
		reflect.TypeOf((*DeployTimeSubstitutedFileProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_s3_deployment.DeploymentSourceContext",
		reflect.TypeOf((*DeploymentSourceContext)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_s3_deployment.ISource",
		reflect.TypeOf((*ISource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ISource{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_s3_deployment.JsonProcessingOptions",
		reflect.TypeOf((*JsonProcessingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_s3_deployment.MarkersConfig",
		reflect.TypeOf((*MarkersConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_s3_deployment.ServerSideEncryption",
		reflect.TypeOf((*ServerSideEncryption)(nil)).Elem(),
		map[string]interface{}{
			"AES_256": ServerSideEncryption_AES_256,
			"AWS_KMS": ServerSideEncryption_AWS_KMS,
			"AWS_KMS_DSSE": ServerSideEncryption_AWS_KMS_DSSE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_s3_deployment.Source",
		reflect.TypeOf((*Source)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Source{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_s3_deployment.SourceConfig",
		reflect.TypeOf((*SourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_s3_deployment.StorageClass",
		reflect.TypeOf((*StorageClass)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": StorageClass_STANDARD,
			"REDUCED_REDUNDANCY": StorageClass_REDUCED_REDUNDANCY,
			"STANDARD_IA": StorageClass_STANDARD_IA,
			"ONEZONE_IA": StorageClass_ONEZONE_IA,
			"INTELLIGENT_TIERING": StorageClass_INTELLIGENT_TIERING,
			"GLACIER": StorageClass_GLACIER,
			"DEEP_ARCHIVE": StorageClass_DEEP_ARCHIVE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_s3_deployment.UserDefinedObjectMetadata",
		reflect.TypeOf((*UserDefinedObjectMetadata)(nil)).Elem(),
	)
}
