package awsecrassets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// networking mode on build time supported by docker.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewDockerImageAsset(this, jsii.String("MyBuildImage"), &DockerImageAssetProps{
//   	Directory: path.join(__dirname, jsii.String("my-image")),
//   	NetworkMode: awscdk.NetworkMode_HOST(),
//   })
//
type NetworkMode interface {
	// The networking mode to use for docker build.
	Mode() *string
}

// The jsii proxy struct for NetworkMode
type jsiiProxy_NetworkMode struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkMode) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}


// Used to specify a custom networking mode Use this if the networking mode name is not yet supported by the CDK.
func NetworkMode_Custom(mode *string) NetworkMode {
	_init_.Initialize()

	if err := validateNetworkMode_CustomParameters(mode); err != nil {
		panic(err)
	}
	var returns NetworkMode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr_assets.NetworkMode",
		"custom",
		[]interface{}{mode},
		&returns,
	)

	return returns
}

// Reuse another container's network stack.
func NetworkMode_FromContainer(containerId *string) NetworkMode {
	_init_.Initialize()

	if err := validateNetworkMode_FromContainerParameters(containerId); err != nil {
		panic(err)
	}
	var returns NetworkMode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr_assets.NetworkMode",
		"fromContainer",
		[]interface{}{containerId},
		&returns,
	)

	return returns
}

func NetworkMode_DEFAULT() NetworkMode {
	_init_.Initialize()
	var returns NetworkMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr_assets.NetworkMode",
		"DEFAULT",
		&returns,
	)
	return returns
}

func NetworkMode_HOST() NetworkMode {
	_init_.Initialize()
	var returns NetworkMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr_assets.NetworkMode",
		"HOST",
		&returns,
	)
	return returns
}

func NetworkMode_NONE() NetworkMode {
	_init_.Initialize()
	var returns NetworkMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr_assets.NetworkMode",
		"NONE",
		&returns,
	)
	return returns
}

