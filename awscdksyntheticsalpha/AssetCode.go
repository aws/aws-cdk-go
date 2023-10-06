package awscdksyntheticsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdksyntheticsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Canary code from an Asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import synthetics_alpha "github.com/aws/aws-cdk-go/awscdksyntheticsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var grantable iGrantable
//   var localBundling iLocalBundling
//
//   assetCode := synthetics_alpha.NewAssetCode(jsii.String("assetPath"), &AssetOptions{
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
	// Called when the canary is initialized to allow this object to bind to the stack, add resources and have fun.
	// Experimental.
	Bind(scope constructs.Construct, handler *string, family RuntimeFamily) *CodeConfig
}

// The jsii proxy struct for AssetCode
type jsiiProxy_AssetCode struct {
	jsiiProxy_Code
}

// Experimental.
func NewAssetCode(assetPath *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateNewAssetCodeParameters(assetPath, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetCode{}

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.AssetCode",
		[]interface{}{assetPath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetCode_Override(a AssetCode, assetPath *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.AssetCode",
		[]interface{}{assetPath, options},
		a,
	)
}

// Specify code from a local path.
//
// Path must include the folder structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `AssetCode` associated with the specified path.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
// Experimental.
func AssetCode_FromAsset(assetPath *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateAssetCode_FromAssetParameters(assetPath, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.AssetCode",
		"fromAsset",
		[]interface{}{assetPath, options},
		&returns,
	)

	return returns
}

// Specify code from an s3 bucket.
//
// The object in the s3 bucket must be a .zip file that contains
// the structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `S3Code` associated with the specified S3 object.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
// Experimental.
func AssetCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	if err := validateAssetCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.AssetCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Specify code inline.
//
// Returns: `InlineCode` with inline code.
// Experimental.
func AssetCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateAssetCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.AssetCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetCode) Bind(scope constructs.Construct, handler *string, family RuntimeFamily) *CodeConfig {
	if err := a.validateBindParameters(scope, handler, family); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, handler, family},
		&returns,
	)

	return returns
}

