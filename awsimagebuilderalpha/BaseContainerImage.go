package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
)

// Represents a base image that is used to start from in EC2 Image Builder image builds.
//
// Example:
//   parameterizedComponent := imagebuilder.Component_FromComponentName(this, jsii.String("ParameterizedComponent"), jsii.String("my-parameterized-component"))
//
//   imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("ParameterizedImageRecipe"), &ImageRecipeProps{
//   	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
//   	Components: []ComponentConfiguration{
//   		&ComponentConfiguration{
//   			Component: parameterizedComponent,
//   			Parameters: map[string]ComponentParameterValue{
//   				"environment": imagebuilder.ComponentParameterValue_fromString(jsii.String("production")),
//   				"version": imagebuilder.ComponentParameterValue_fromString(jsii.String("1.0.0")),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type BaseContainerImage interface {
	// The rendered base image to use.
	// Experimental.
	Image() *string
}

// The jsii proxy struct for BaseContainerImage
type jsiiProxy_BaseContainerImage struct {
	_ byte // padding
}

func (j *jsiiProxy_BaseContainerImage) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}


// Experimental.
func NewBaseContainerImage(image *string) BaseContainerImage {
	_init_.Initialize()

	if err := validateNewBaseContainerImageParameters(image); err != nil {
		panic(err)
	}
	j := jsiiProxy_BaseContainerImage{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.BaseContainerImage",
		[]interface{}{image},
		&j,
	)

	return &j
}

// Experimental.
func NewBaseContainerImage_Override(b BaseContainerImage, image *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.BaseContainerImage",
		[]interface{}{image},
		b,
	)
}

// The DockerHub image to use as the base image in a container recipe.
// Experimental.
func BaseContainerImage_FromDockerHub(repository *string, tag *string) BaseContainerImage {
	_init_.Initialize()

	if err := validateBaseContainerImage_FromDockerHubParameters(repository, tag); err != nil {
		panic(err)
	}
	var returns BaseContainerImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.BaseContainerImage",
		"fromDockerHub",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// The ECR container image to use as the base image in a container recipe.
// Experimental.
func BaseContainerImage_FromEcr(repository awsecr.IRepository, tag *string) BaseContainerImage {
	_init_.Initialize()

	if err := validateBaseContainerImage_FromEcrParameters(repository, tag); err != nil {
		panic(err)
	}
	var returns BaseContainerImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.BaseContainerImage",
		"fromEcr",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// The ECR public container image to use as the base image in a container recipe.
// Experimental.
func BaseContainerImage_FromEcrPublic(registryAlias *string, repositoryName *string, tag *string) BaseContainerImage {
	_init_.Initialize()

	if err := validateBaseContainerImage_FromEcrPublicParameters(registryAlias, repositoryName, tag); err != nil {
		panic(err)
	}
	var returns BaseContainerImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.BaseContainerImage",
		"fromEcrPublic",
		[]interface{}{registryAlias, repositoryName, tag},
		&returns,
	)

	return returns
}

// The string value of the base image to use in a container recipe.
//
// This can be an EC2 Image Builder image ARN,
// an ECR or ECR public image, or a container URI sourced from a third-party container registry such as DockerHub.
// Experimental.
func BaseContainerImage_FromString(baseContainerImageString *string) BaseContainerImage {
	_init_.Initialize()

	if err := validateBaseContainerImage_FromStringParameters(baseContainerImageString); err != nil {
		panic(err)
	}
	var returns BaseContainerImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.BaseContainerImage",
		"fromString",
		[]interface{}{baseContainerImageString},
		&returns,
	)

	return returns
}

