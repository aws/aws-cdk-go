package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsecr"
)

// Code property for the DockerImageFunction construct.
//
// Example:
//   lambda.NewDockerImageFunction(this, jsii.String("AssetFunction"), &DockerImageFunctionProps{
//   	Code: lambda.DockerImageCode_FromImageAsset(path.join(__dirname, jsii.String("docker-handler"))),
//   })
//
// Experimental.
type DockerImageCode interface {
}

// The jsii proxy struct for DockerImageCode
type jsiiProxy_DockerImageCode struct {
	_ byte // padding
}

// Experimental.
func NewDockerImageCode_Override(d DockerImageCode) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda.DockerImageCode",
		nil, // no parameters
		d,
	)
}

// Use an existing ECR image as the Lambda code.
// Experimental.
func DockerImageCode_FromEcr(repository awsecr.IRepository, props *EcrImageCodeProps) DockerImageCode {
	_init_.Initialize()

	if err := validateDockerImageCode_FromEcrParameters(repository, props); err != nil {
		panic(err)
	}
	var returns DockerImageCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.DockerImageCode",
		"fromEcr",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
// Experimental.
func DockerImageCode_FromImageAsset(directory *string, props *AssetImageCodeProps) DockerImageCode {
	_init_.Initialize()

	if err := validateDockerImageCode_FromImageAssetParameters(directory, props); err != nil {
		panic(err)
	}
	var returns DockerImageCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.DockerImageCode",
		"fromImageAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

