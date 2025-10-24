package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// An import source from an inline string.
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
//   inlineImportSource := awscdk.Aws_cloudfront.InlineImportSource_FromAsset(jsii.String("path"), &AssetOptions{
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
type InlineImportSource interface {
	ImportSource
	// the contents of the KeyValueStore.
	Data() *string
}

// The jsii proxy struct for InlineImportSource
type jsiiProxy_InlineImportSource struct {
	jsiiProxy_ImportSource
}

func (j *jsiiProxy_InlineImportSource) Data() *string {
	var returns *string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}


func NewInlineImportSource(data *string) InlineImportSource {
	_init_.Initialize()

	if err := validateNewInlineImportSourceParameters(data); err != nil {
		panic(err)
	}
	j := jsiiProxy_InlineImportSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.InlineImportSource",
		[]interface{}{data},
		&j,
	)

	return &j
}

func NewInlineImportSource_Override(i InlineImportSource, data *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.InlineImportSource",
		[]interface{}{data},
		i,
	)
}

// An import source that exists as a local file.
func InlineImportSource_FromAsset(path *string, options *awss3assets.AssetOptions) ImportSource {
	_init_.Initialize()

	if err := validateInlineImportSource_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.InlineImportSource",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// An import source that exists as an object in an S3 bucket.
func InlineImportSource_FromBucket(bucket awss3.IBucket, key *string) ImportSource {
	_init_.Initialize()

	if err := validateInlineImportSource_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.InlineImportSource",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// An import source that uses an inline string.
func InlineImportSource_FromInline(data *string) ImportSource {
	_init_.Initialize()

	if err := validateInlineImportSource_FromInlineParameters(data); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.InlineImportSource",
		"fromInline",
		[]interface{}{data},
		&returns,
	)

	return returns
}

