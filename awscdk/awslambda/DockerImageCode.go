package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
)

// Code property for the DockerImageFunction construct.
//
// Example:
//   lambda.NewDockerImageFunction(this, jsii.String("AssetFunction"), &DockerImageFunctionProps{
//   	Code: lambda.DockerImageCode_FromImageAsset(path.join(__dirname, jsii.String("docker-handler"))),
//   })
//
type DockerImageCode interface {
}

// The jsii proxy struct for DockerImageCode
type jsiiProxy_DockerImageCode struct {
	_ byte // padding
}

func NewDockerImageCode_Override(d DockerImageCode) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.DockerImageCode",
		nil, // no parameters
		d,
	)
}

// Use an existing ECR image as the Lambda code.
func DockerImageCode_FromEcr(repository awsecr.IRepository, props *EcrImageCodeProps) DockerImageCode {
	_init_.Initialize()

	if err := validateDockerImageCode_FromEcrParameters(repository, props); err != nil {
		panic(err)
	}
	var returns DockerImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageCode",
		"fromEcr",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func DockerImageCode_FromImageAsset(directory *string, props *AssetImageCodeProps) DockerImageCode {
	_init_.Initialize()

	if err := validateDockerImageCode_FromImageAssetParameters(directory, props); err != nil {
		panic(err)
	}
	var returns DockerImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageCode",
		"fromImageAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

