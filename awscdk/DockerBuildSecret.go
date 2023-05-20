package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Methods to build Docker CLI arguments for builds using secrets.
//
// Docker BuildKit must be enabled to use build secrets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerBuildSecret := cdk.NewDockerBuildSecret()
//
// See: https://docs.docker.com/build/buildkit/
//
type DockerBuildSecret interface {
}

// The jsii proxy struct for DockerBuildSecret
type jsiiProxy_DockerBuildSecret struct {
	_ byte // padding
}

func NewDockerBuildSecret() DockerBuildSecret {
	_init_.Initialize()

	j := jsiiProxy_DockerBuildSecret{}

	_jsii_.Create(
		"aws-cdk-lib.DockerBuildSecret",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDockerBuildSecret_Override(d DockerBuildSecret) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.DockerBuildSecret",
		nil, // no parameters
		d,
	)
}

// A Docker build secret from a file source.
//
// Returns: The latter half required for `--secret`.
func DockerBuildSecret_FromSrc(src *string) *string {
	_init_.Initialize()

	if err := validateDockerBuildSecret_FromSrcParameters(src); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerBuildSecret",
		"fromSrc",
		[]interface{}{src},
		&returns,
	)

	return returns
}

