package awscdksagemakeralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdksagemakeralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Constructs for types of container images.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   var image containerImage
//   var modelData modelData
//
//
//   model := sagemaker.NewModel(this, jsii.String("ContainerModel"), &ModelProps{
//   	Containers: []containerDefinition{
//   		&containerDefinition{
//   			Image: *Image,
//   			ModelData: *ModelData,
//   		},
//   	},
//   	NetworkIsolation: jsii.Boolean(true),
//   })
//
// Experimental.
type ContainerImage interface {
	// Called when the image is used by a Model.
	// Experimental.
	Bind(scope constructs.Construct, model Model) *ContainerImageConfig
}

// The jsii proxy struct for ContainerImage
type jsiiProxy_ContainerImage struct {
	_ byte // padding
}

// Experimental.
func NewContainerImage_Override(c ContainerImage) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-sagemaker-alpha.ContainerImage",
		nil, // no parameters
		c,
	)
}

// Reference an image that's constructed directly from sources on disk.
// Experimental.
func ContainerImage_FromAsset(directory *string, options *awsecrassets.DockerImageAssetOptions) ContainerImage {
	_init_.Initialize()

	if err := validateContainerImage_FromAssetParameters(directory, options); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.ContainerImage",
		"fromAsset",
		[]interface{}{directory, options},
		&returns,
	)

	return returns
}

// Reference an AWS Deep Learning Container image.
// Experimental.
func ContainerImage_FromDlc(repositoryName *string, tag *string, accountId *string) ContainerImage {
	_init_.Initialize()

	if err := validateContainerImage_FromDlcParameters(repositoryName, tag); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.ContainerImage",
		"fromDlc",
		[]interface{}{repositoryName, tag, accountId},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
// Experimental.
func ContainerImage_FromEcrRepository(repository awsecr.IRepository, tag *string) ContainerImage {
	_init_.Initialize()

	if err := validateContainerImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.ContainerImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerImage) Bind(scope constructs.Construct, model Model) *ContainerImageConfig {
	if err := c.validateBindParameters(scope, model); err != nil {
		panic(err)
	}
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, model},
		&returns,
	)

	return returns
}

