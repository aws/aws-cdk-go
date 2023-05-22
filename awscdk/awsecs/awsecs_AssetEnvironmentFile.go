package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

// Environment file from a local directory.
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
//   var localBundling iLocalBundling
//
//   assetEnvironmentFile := awscdk.Aws_ecs.NewAssetEnvironmentFile(jsii.String("path"), &AssetOptions{
//   	AssetHash: jsii.String("assetHash"),
//   	AssetHashType: monocdk.AssetHashType_SOURCE,
//   	Bundling: &BundlingOptions{
//   		Image: dockerImage,
//
//   		// the properties below are optional
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
//   		OutputType: monocdk.BundlingOutput_ARCHIVED,
//   		SecurityOpt: jsii.String("securityOpt"),
//   		User: jsii.String("user"),
//   		Volumes: []dockerVolume{
//   			&dockerVolume{
//   				ContainerPath: jsii.String("containerPath"),
//   				HostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				Consistency: monocdk.DockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		WorkingDirectory: jsii.String("workingDirectory"),
//   	},
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	Follow: awscdk.Assets.FollowMode_NEVER,
//   	FollowSymlinks: monocdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: monocdk.IgnoreMode_GLOB,
//   	Readers: []*iGrantable{
//   		grantable,
//   	},
//   	SourceHash: jsii.String("sourceHash"),
//   })
//
// Experimental.
type AssetEnvironmentFile interface {
	EnvironmentFile
	// The path to the asset file or directory.
	// Experimental.
	Path() *string
	// Called when the container is initialized to allow this object to bind to the stack.
	// Experimental.
	Bind(scope awscdk.Construct) *EnvironmentFileConfig
}

// The jsii proxy struct for AssetEnvironmentFile
type jsiiProxy_AssetEnvironmentFile struct {
	jsiiProxy_EnvironmentFile
}

func (j *jsiiProxy_AssetEnvironmentFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


// Experimental.
func NewAssetEnvironmentFile(path *string, options *awss3assets.AssetOptions) AssetEnvironmentFile {
	_init_.Initialize()

	if err := validateNewAssetEnvironmentFileParameters(path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetEnvironmentFile{}

	_jsii_.Create(
		"monocdk.aws_ecs.AssetEnvironmentFile",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetEnvironmentFile_Override(a AssetEnvironmentFile, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.AssetEnvironmentFile",
		[]interface{}{path, options},
		a,
	)
}

// Loads the environment file from a local disk path.
// Experimental.
func AssetEnvironmentFile_FromAsset(path *string, options *awss3assets.AssetOptions) AssetEnvironmentFile {
	_init_.Initialize()

	if err := validateAssetEnvironmentFile_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetEnvironmentFile

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.AssetEnvironmentFile",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Loads the environment file from an S3 bucket.
//
// Returns: `S3EnvironmentFile` associated with the specified S3 object.
// Experimental.
func AssetEnvironmentFile_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3EnvironmentFile {
	_init_.Initialize()

	if err := validateAssetEnvironmentFile_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3EnvironmentFile

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.AssetEnvironmentFile",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetEnvironmentFile) Bind(scope awscdk.Construct) *EnvironmentFileConfig {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *EnvironmentFileConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

