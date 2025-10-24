package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// An import source from an S3 object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage DockerImage
//   var grantable IGrantable
//   var keyRef IKeyRef
//   var localBundling ILocalBundling
//
//   s3ImportSource := awscdk.Aws_cloudfront.S3ImportSource_FromAsset(jsii.String("path"), &AssetOptions{
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
//   		Volumes: []DockerVolume{
//   			&DockerVolume{
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
//   	DisplayName: jsii.String("displayName"),
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	FollowSymlinks: cdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   	Readers: []IGrantable{
//   		grantable,
//   	},
//   	SourceKMSKey: keyRef,
//   })
//
type S3ImportSource interface {
	ImportSource
	// the S3 bucket that contains the data.
	Bucket() awss3.IBucket
	// the key within the S3 bucket that contains the data.
	Key() *string
}

// The jsii proxy struct for S3ImportSource
type jsiiProxy_S3ImportSource struct {
	jsiiProxy_ImportSource
}

func (j *jsiiProxy_S3ImportSource) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ImportSource) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}


func NewS3ImportSource(bucket awss3.IBucket, key *string) S3ImportSource {
	_init_.Initialize()

	if err := validateNewS3ImportSourceParameters(bucket, key); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ImportSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.S3ImportSource",
		[]interface{}{bucket, key},
		&j,
	)

	return &j
}

func NewS3ImportSource_Override(s S3ImportSource, bucket awss3.IBucket, key *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.S3ImportSource",
		[]interface{}{bucket, key},
		s,
	)
}

// An import source that exists as a local file.
func S3ImportSource_FromAsset(path *string, options *awss3assets.AssetOptions) ImportSource {
	_init_.Initialize()

	if err := validateS3ImportSource_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.S3ImportSource",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// An import source that exists as an object in an S3 bucket.
func S3ImportSource_FromBucket(bucket awss3.IBucket, key *string) ImportSource {
	_init_.Initialize()

	if err := validateS3ImportSource_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.S3ImportSource",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// An import source that uses an inline string.
func S3ImportSource_FromInline(data *string) ImportSource {
	_init_.Initialize()

	if err := validateS3ImportSource_FromInlineParameters(data); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.S3ImportSource",
		"fromInline",
		[]interface{}{data},
		&returns,
	)

	return returns
}

