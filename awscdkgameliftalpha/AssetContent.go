// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Game content from a local directory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var grantable iGrantable
//   var localBundling iLocalBundling
//
//   assetContent := gamelift_alpha.NewAssetContent(jsii.String("path"), &AssetOptions{
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
// Experimental.
type AssetContent interface {
	Content
	// The path to the asset file or directory.
	// Experimental.
	Path() *string
	// Called when the Build is initialized to allow this object to bind.
	// Experimental.
	Bind(scope constructs.Construct, role awsiam.IRole) *ContentConfig
}

// The jsii proxy struct for AssetContent
type jsiiProxy_AssetContent struct {
	jsiiProxy_Content
}

func (j *jsiiProxy_AssetContent) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


// Experimental.
func NewAssetContent(path *string, options *awss3assets.AssetOptions) AssetContent {
	_init_.Initialize()

	if err := validateNewAssetContentParameters(path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetContent{}

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.AssetContent",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetContent_Override(a AssetContent, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.AssetContent",
		[]interface{}{path, options},
		a,
	)
}

// Loads the game content from a local disk path.
// Experimental.
func AssetContent_FromAsset(path *string, options *awss3assets.AssetOptions) AssetContent {
	_init_.Initialize()

	if err := validateAssetContent_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetContent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.AssetContent",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Game content as an S3 object.
// Experimental.
func AssetContent_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Content {
	_init_.Initialize()

	if err := validateAssetContent_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Content

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.AssetContent",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetContent) Bind(scope constructs.Construct, role awsiam.IRole) *ContentConfig {
	if err := a.validateBindParameters(scope, role); err != nil {
		panic(err)
	}
	var returns *ContentConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, role},
		&returns,
	)

	return returns
}

