package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

// Extract an archive into a directory.
//
// Example:
//   var myBucket bucket
//
//
//   handle := ec2.NewInitServiceRestartHandle()
//
//   ec2.cloudFormationInit.fromElements(ec2.initFile.fromString(jsii.String("/etc/nginx/nginx.conf"), jsii.String("..."), &initFileOptions{
//   	serviceRestartHandles: []initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.initSource.fromS3Object(jsii.String("/var/www/html"), myBucket, jsii.String("html.zip"), &initSourceOptions{
//   	serviceRestartHandles: []*initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.initService.enable(jsii.String("nginx"), &initServiceOptions{
//   	serviceRestartHandle: handle,
//   }))
//
// Experimental.
type InitSource interface {
	InitElement
	// Returns the init element type for this element.
	// Experimental.
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


// Experimental.
func NewInitSource_Override(i InitSource, targetDirectory *string, serviceHandles *[]InitServiceRestartHandle) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.InitSource",
		[]interface{}{targetDirectory, serviceHandles},
		i,
	)
}

// Create an InitSource from an asset created from the given path.
// Experimental.
func InitSource_FromAsset(targetDirectory *string, path *string, options *InitSourceAssetOptions) InitSource {
	_init_.Initialize()

	var returns InitSource

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.InitSource",
		"fromAsset",
		[]interface{}{targetDirectory, path, options},
		&returns,
	)

	return returns
}

// Extract a directory from an existing directory asset.
// Experimental.
func InitSource_FromExistingAsset(targetDirectory *string, asset awss3assets.Asset, options *InitSourceOptions) InitSource {
	_init_.Initialize()

	var returns InitSource

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.InitSource",
		"fromExistingAsset",
		[]interface{}{targetDirectory, asset, options},
		&returns,
	)

	return returns
}

// Extract a GitHub branch into a given directory.
// Experimental.
func InitSource_FromGitHub(targetDirectory *string, owner *string, repo *string, refSpec *string, options *InitSourceOptions) InitSource {
	_init_.Initialize()

	var returns InitSource

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.InitSource",
		"fromGitHub",
		[]interface{}{targetDirectory, owner, repo, refSpec, options},
		&returns,
	)

	return returns
}

// Extract an archive stored in an S3 bucket into the given directory.
// Experimental.
func InitSource_FromS3Object(targetDirectory *string, bucket awss3.IBucket, key *string, options *InitSourceOptions) InitSource {
	_init_.Initialize()

	var returns InitSource

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.InitSource",
		"fromS3Object",
		[]interface{}{targetDirectory, bucket, key, options},
		&returns,
	)

	return returns
}

// Retrieve a URL and extract it into the given directory.
// Experimental.
func InitSource_FromUrl(targetDirectory *string, url *string, options *InitSourceOptions) InitSource {
	_init_.Initialize()

	var returns InitSource

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.InitSource",
		"fromUrl",
		[]interface{}{targetDirectory, url, options},
		&returns,
	)

	return returns
}

