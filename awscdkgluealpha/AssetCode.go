package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Job Code from a local file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var grantable iGrantable
//   var localBundling iLocalBundling
//
//   assetCode := glue_alpha.NewAssetCode(jsii.String("path"), &AssetOptions{
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
// Experimental.
type AssetCode interface {
	Code
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig
}

// The jsii proxy struct for AssetCode
type jsiiProxy_AssetCode struct {
	jsiiProxy_Code
}

// Experimental.
func NewAssetCode(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateNewAssetCodeParameters(path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetCode{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.AssetCode",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetCode_Override(a AssetCode, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.AssetCode",
		[]interface{}{path, options},
		a,
	)
}

// Job code from a local disk path.
// Experimental.
func AssetCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateAssetCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.AssetCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Job code as an S3 object.
// Experimental.
func AssetCode_FromBucket(bucket awss3.IBucket, key *string) S3Code {
	_init_.Initialize()

	if err := validateAssetCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.AssetCode",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetCode) Bind(scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig {
	if err := a.validateBindParameters(scope, grantable); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, grantable},
		&returns,
	)

	return returns
}

