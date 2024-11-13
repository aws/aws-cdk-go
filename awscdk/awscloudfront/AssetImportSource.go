package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// An import source from a local file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var grantable iGrantable
//   var key key
//   var localBundling iLocalBundling
//
//   assetImportSource := awscdk.Aws_cloudfront.AssetImportSource_FromAsset(jsii.String("path"), &AssetOptions{
//   	AssetHash: jsii.String("assetHash"),
//   	AssetHashType: cdk.AssetHashType_SOURCE,
//   	Bundling: &BundlingOptions{
//   		Image: dockerImage,
//
//   		// the properties below are optional
//   		BundlingFileAccess: cdk.BundlingFileAccess_VOLUME_COPY,
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		Entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		Local: localBundling,
//   		Network: jsii.String("network"),
//   		OutputType: cdk.BundlingOutput_ARCHIVED,
//   		Platform: jsii.String("platform"),
//   		SecurityOpt: jsii.String("securityOpt"),
//   		User: jsii.String("user"),
//   		Volumes: []dockerVolume{
//   			&dockerVolume{
//   				ContainerPath: jsii.String("containerPath"),
//   				HostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				Consistency: cdk.DockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		VolumesFrom: []*string{
//   			jsii.String("volumesFrom"),
//   		},
//   		WorkingDirectory: jsii.String("workingDirectory"),
//   	},
//   	DeployTime: jsii.Boolean(false),
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	FollowSymlinks: cdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   	Readers: []*iGrantable{
//   		grantable,
//   	},
//   	SourceKMSKey: key,
//   })
//
type AssetImportSource interface {
	ImportSource
	// the path to the local file.
	Path() *string
}

// The jsii proxy struct for AssetImportSource
type jsiiProxy_AssetImportSource struct {
	jsiiProxy_ImportSource
}

func (j *jsiiProxy_AssetImportSource) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


func NewAssetImportSource(path *string, options *awss3assets.AssetOptions) AssetImportSource {
	_init_.Initialize()

	if err := validateNewAssetImportSourceParameters(path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetImportSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.AssetImportSource",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

func NewAssetImportSource_Override(a AssetImportSource, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.AssetImportSource",
		[]interface{}{path, options},
		a,
	)
}

// An import source that exists as a local file.
func AssetImportSource_FromAsset(path *string, options *awss3assets.AssetOptions) ImportSource {
	_init_.Initialize()

	if err := validateAssetImportSource_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.AssetImportSource",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// An import source that exists as an object in an S3 bucket.
func AssetImportSource_FromBucket(bucket awss3.IBucket, key *string) ImportSource {
	_init_.Initialize()

	if err := validateAssetImportSource_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.AssetImportSource",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// An import source that uses an inline string.
func AssetImportSource_FromInline(data *string) ImportSource {
	_init_.Initialize()

	if err := validateAssetImportSource_FromInlineParameters(data); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.AssetImportSource",
		"fromInline",
		[]interface{}{data},
		&returns,
	)

	return returns
}

