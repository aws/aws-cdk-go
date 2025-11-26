package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// Represents a container instance image that is used to launch the instance used for building the container for an EC2 Image Builder container build.
//
// Example:
//   containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("InstanceConfigContainerRecipe"), &ContainerRecipeProps{
//   	BaseImage: imagebuilder.BaseContainerImage_*FromDockerHub(jsii.String("amazonlinux"), jsii.String("latest")),
//   	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
//   	// Custom ECS-optimized AMI for building
//   	InstanceImage: imagebuilder.ContainerInstanceImage_FromSsmParameterName(jsii.String("/aws/service/ecs/optimized-ami/amazon-linux-2023/recommended/image_id")),
//   	// Additional storage for build process
//   	InstanceBlockDevices: []BlockDevice{
//   		&BlockDevice{
//   			DeviceName: jsii.String("/dev/xvda"),
//   			Volume: ec2.BlockDeviceVolume_Ebs(jsii.Number(50), &EbsDeviceOptions{
//   				Encrypted: jsii.Boolean(true),
//   				VolumeType: ec2.EbsDeviceVolumeType_GENERAL_PURPOSE_SSD_GP3,
//   			}),
//   		},
//   	},
//   })
//
// Experimental.
type ContainerInstanceImage interface {
	// The rendered container instance image to use.
	// Experimental.
	Image() *string
}

// The jsii proxy struct for ContainerInstanceImage
type jsiiProxy_ContainerInstanceImage struct {
	_ byte // padding
}

func (j *jsiiProxy_ContainerInstanceImage) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}


// Experimental.
func NewContainerInstanceImage(image *string) ContainerInstanceImage {
	_init_.Initialize()

	if err := validateNewContainerInstanceImageParameters(image); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerInstanceImage{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerInstanceImage",
		[]interface{}{image},
		&j,
	)

	return &j
}

// Experimental.
func NewContainerInstanceImage_Override(c ContainerInstanceImage, image *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerInstanceImage",
		[]interface{}{image},
		c,
	)
}

// The AMI ID to use to launch the instance for building the container image.
// Experimental.
func ContainerInstanceImage_FromAmiId(amiId *string) ContainerInstanceImage {
	_init_.Initialize()

	if err := validateContainerInstanceImage_FromAmiIdParameters(amiId); err != nil {
		panic(err)
	}
	var returns ContainerInstanceImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerInstanceImage",
		"fromAmiId",
		[]interface{}{amiId},
		&returns,
	)

	return returns
}

// The SSM parameter to use to launch the instance for building the container image.
// Experimental.
func ContainerInstanceImage_FromSsmParameter(parameter awsssm.IStringParameter) ContainerInstanceImage {
	_init_.Initialize()

	if err := validateContainerInstanceImage_FromSsmParameterParameters(parameter); err != nil {
		panic(err)
	}
	var returns ContainerInstanceImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerInstanceImage",
		"fromSsmParameter",
		[]interface{}{parameter},
		&returns,
	)

	return returns
}

// The ARN of the SSM parameter used to launch the instance for building the container image.
// Experimental.
func ContainerInstanceImage_FromSsmParameterName(parameterName *string) ContainerInstanceImage {
	_init_.Initialize()

	if err := validateContainerInstanceImage_FromSsmParameterNameParameters(parameterName); err != nil {
		panic(err)
	}
	var returns ContainerInstanceImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerInstanceImage",
		"fromSsmParameterName",
		[]interface{}{parameterName},
		&returns,
	)

	return returns
}

// The string value of the container instance image to use in a container recipe.
//
// This can either be:
// - an SSM parameter reference, prefixed with `ssm:` and followed by the parameter name or ARN
// - an AMI ID.
// Experimental.
func ContainerInstanceImage_FromString(containerInstanceImageString *string) ContainerInstanceImage {
	_init_.Initialize()

	if err := validateContainerInstanceImage_FromStringParameters(containerInstanceImageString); err != nil {
		panic(err)
	}
	var returns ContainerInstanceImage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerInstanceImage",
		"fromString",
		[]interface{}{containerInstanceImageString},
		&returns,
	)

	return returns
}

