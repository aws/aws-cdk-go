package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Environment file from a local directory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
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
//   })
//
type AssetEnvironmentFile interface {
	EnvironmentFile
	// The path to the asset file or directory.
	Path() *string
	// Called when the container is initialized to allow this object to bind to the stack.
	Bind(scope constructs.Construct) *EnvironmentFileConfig
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


func NewAssetEnvironmentFile(path *string, options *awss3assets.AssetOptions) AssetEnvironmentFile {
	_init_.Initialize()

	if err := validateNewAssetEnvironmentFileParameters(path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetEnvironmentFile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AssetEnvironmentFile",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

func NewAssetEnvironmentFile_Override(a AssetEnvironmentFile, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AssetEnvironmentFile",
		[]interface{}{path, options},
		a,
	)
}

// Loads the environment file from a local disk path.
func AssetEnvironmentFile_FromAsset(path *string, options *awss3assets.AssetOptions) AssetEnvironmentFile {
	_init_.Initialize()

	if err := validateAssetEnvironmentFile_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetEnvironmentFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetEnvironmentFile",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Loads the environment file from an S3 bucket.
//
// Returns: `S3EnvironmentFile` associated with the specified S3 object.
func AssetEnvironmentFile_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3EnvironmentFile {
	_init_.Initialize()

	if err := validateAssetEnvironmentFile_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3EnvironmentFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetEnvironmentFile",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetEnvironmentFile) Bind(scope constructs.Construct) *EnvironmentFileConfig {
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

