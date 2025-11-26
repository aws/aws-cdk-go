package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
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
type BaseImage interface {
	// The rendered base image to use.
	// Experimental.
	Image() *string
}

// The jsii proxy struct for BaseImage
type jsiiProxy_BaseImage struct {
	_ byte // padding
}

func (j *jsiiProxy_BaseImage) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}


// Experimental.
func NewBaseImage(image *string) BaseImage {
	_init_.Initialize()

	if err := validateNewBaseImageParameters(image); err != nil {
		panic(err)
	}
	j := jsiiProxy_BaseImage{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.BaseImage",
		[]interface{}{image},
		&j,
	)

	return &j
}

// Experimental.
func NewBaseImage_Override(b BaseImage, image *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.BaseImage",
		[]interface{}{image},
		b,
	)
}

// The AMI ID to use as a base image in an image recipe.
// Experimental.
func BaseImage_FromAmiId(amiId *string) BaseImage {
	_init_.Initialize()

	if err := validateBaseImage_FromAmiIdParameters(amiId); err != nil {
		panic(err)
	}
	var returns BaseImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.BaseImage",
		"fromAmiId",
		[]interface{}{amiId},
		&returns,
	)

	return returns
}

// The marketplace product ID for an AMI product to use as the base image in an image recipe.
// Experimental.
func BaseImage_FromMarketplaceProductId(productId *string) BaseImage {
	_init_.Initialize()

	if err := validateBaseImage_FromMarketplaceProductIdParameters(productId); err != nil {
		panic(err)
	}
	var returns BaseImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.BaseImage",
		"fromMarketplaceProductId",
		[]interface{}{productId},
		&returns,
	)

	return returns
}

// The SSM parameter to use as the base image in an image recipe.
// Experimental.
func BaseImage_FromSsmParameter(parameter awsssm.IParameter) BaseImage {
	_init_.Initialize()

	if err := validateBaseImage_FromSsmParameterParameters(parameter); err != nil {
		panic(err)
	}
	var returns BaseImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.BaseImage",
		"fromSsmParameter",
		[]interface{}{parameter},
		&returns,
	)

	return returns
}

// The parameter name for the SSM parameter to use as the base image in an image recipe.
// Experimental.
func BaseImage_FromSsmParameterName(parameterName *string) BaseImage {
	_init_.Initialize()

	if err := validateBaseImage_FromSsmParameterNameParameters(parameterName); err != nil {
		panic(err)
	}
	var returns BaseImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.BaseImage",
		"fromSsmParameterName",
		[]interface{}{parameterName},
		&returns,
	)

	return returns
}

// The direct string value of the base image to use in an image recipe.
//
// This can be an EC2 Image Builder image ARN,
// an SSM parameter, an AWS Marketplace product ID, or an AMI ID.
// Experimental.
func BaseImage_FromString(baseImageString *string) BaseImage {
	_init_.Initialize()

	if err := validateBaseImage_FromStringParameters(baseImageString); err != nil {
		panic(err)
	}
	var returns BaseImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.BaseImage",
		"fromString",
		[]interface{}{baseImageString},
		&returns,
	)

	return returns
}

