package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// Extract an archive into a directory.
//
// Example:
//   var myBucket bucket
//
//
//   handle := ec2.NewInitServiceRestartHandle()
//
//   ec2.CloudFormationInit_FromElements(ec2.InitFile_FromString(jsii.String("/etc/nginx/nginx.conf"), jsii.String("..."), &InitFileOptions{
//   	ServiceRestartHandles: []initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.InitSource_FromS3Object(jsii.String("/var/www/html"), myBucket, jsii.String("html.zip"), &InitSourceOptions{
//   	ServiceRestartHandles: []*initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.InitService_Enable(jsii.String("nginx"), &InitServiceOptions{
//   	ServiceRestartHandle: handle,
//   }))
//
type InitSource interface {
	InitElement
	// Returns the init element type for this element.
	ElementType() *string
}

// The jsii proxy struct for InitSource
type jsiiProxy_InitSource struct {
	jsiiProxy_InitElement
}

func (j *jsiiProxy_InitSource) ElementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementType",
		&returns,
	)
	return returns
}


func NewInitSource_Override(i InitSource, targetDirectory *string, serviceHandles *[]InitServiceRestartHandle) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitSource",
		[]interface{}{targetDirectory, serviceHandles},
		i,
	)
}

// Create an InitSource from an asset created from the given path.
func InitSource_FromAsset(targetDirectory *string, path *string, options *InitSourceAssetOptions) InitSource {
	_init_.Initialize()

	if err := validateInitSource_FromAssetParameters(targetDirectory, path, options); err != nil {
		panic(err)
	}
	var returns InitSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitSource",
		"fromAsset",
		[]interface{}{targetDirectory, path, options},
		&returns,
	)

	return returns
}

// Extract a directory from an existing directory asset.
func InitSource_FromExistingAsset(targetDirectory *string, asset awss3assets.Asset, options *InitSourceOptions) InitSource {
	_init_.Initialize()

	if err := validateInitSource_FromExistingAssetParameters(targetDirectory, asset, options); err != nil {
		panic(err)
	}
	var returns InitSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitSource",
		"fromExistingAsset",
		[]interface{}{targetDirectory, asset, options},
		&returns,
	)

	return returns
}

// Extract a GitHub branch into a given directory.
func InitSource_FromGitHub(targetDirectory *string, owner *string, repo *string, refSpec *string, options *InitSourceOptions) InitSource {
	_init_.Initialize()

	if err := validateInitSource_FromGitHubParameters(targetDirectory, owner, repo, options); err != nil {
		panic(err)
	}
	var returns InitSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitSource",
		"fromGitHub",
		[]interface{}{targetDirectory, owner, repo, refSpec, options},
		&returns,
	)

	return returns
}

// Extract an archive stored in an S3 bucket into the given directory.
func InitSource_FromS3Object(targetDirectory *string, bucket awss3.IBucket, key *string, options *InitSourceOptions) InitSource {
	_init_.Initialize()

	if err := validateInitSource_FromS3ObjectParameters(targetDirectory, bucket, key, options); err != nil {
		panic(err)
	}
	var returns InitSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitSource",
		"fromS3Object",
		[]interface{}{targetDirectory, bucket, key, options},
		&returns,
	)

	return returns
}

// Retrieve a URL and extract it into the given directory.
func InitSource_FromUrl(targetDirectory *string, url *string, options *InitSourceOptions) InitSource {
	_init_.Initialize()

	if err := validateInitSource_FromUrlParameters(targetDirectory, url, options); err != nil {
		panic(err)
	}
	var returns InitSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitSource",
		"fromUrl",
		[]interface{}{targetDirectory, url, options},
		&returns,
	)

	return returns
}

