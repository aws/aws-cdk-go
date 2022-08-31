package awsecrassets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// platform supported by docker.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewDockerImageAsset(this, jsii.String("MyBuildImage"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("my-image")),
//   	platform: awscdk.Platform_LINUX_ARM64(),
//   })
//
// Experimental.
type Platform interface {
	// The platform to use for docker build.
	// Experimental.
	Platform() *string
}

// The jsii proxy struct for Platform
type jsiiProxy_Platform struct {
	_ byte // padding
}

func (j *jsiiProxy_Platform) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}


// Used to specify a custom platform Use this if the platform name is not yet supported by the CDK.
// Experimental.
func Platform_Custom(platform *string) Platform {
	_init_.Initialize()

	var returns Platform

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr_assets.Platform",
		"custom",
		[]interface{}{platform},
		&returns,
	)

	return returns
}

func Platform_LINUX_AMD64() Platform {
	_init_.Initialize()
	var returns Platform
	_jsii_.StaticGet(
		"monocdk.aws_ecr_assets.Platform",
		"LINUX_AMD64",
		&returns,
	)
	return returns
}

func Platform_LINUX_ARM64() Platform {
	_init_.Initialize()
	var returns Platform
	_jsii_.StaticGet(
		"monocdk.aws_ecr_assets.Platform",
		"LINUX_ARM64",
		&returns,
	)
	return returns
}

