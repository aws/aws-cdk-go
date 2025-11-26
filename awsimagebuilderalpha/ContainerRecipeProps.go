package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a Container Recipe resource.
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
type ContainerRecipeProps struct {
	// The base image for customizations specified in the container recipe.
	// Experimental.
	BaseImage BaseContainerImage `field:"required" json:"baseImage" yaml:"baseImage"`
	// The container repository where the output container image is stored.
	// Experimental.
	TargetRepository Repository `field:"required" json:"targetRepository" yaml:"targetRepository"`
	// The list of component configurations to apply in the image build.
	// Default: None.
	//
	// Experimental.
	Components *[]*ComponentConfiguration `field:"optional" json:"components" yaml:"components"`
	// The name of the container recipe.
	// Default: a name is generated.
	//
	// Experimental.
	ContainerRecipeName *string `field:"optional" json:"containerRecipeName" yaml:"containerRecipeName"`
	// The version of the container recipe.
	// Default: 1.0.x
	//
	// Experimental.
	ContainerRecipeVersion *string `field:"optional" json:"containerRecipeVersion" yaml:"containerRecipeVersion"`
	// The description of the container recipe.
	// Default: None.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The dockerfile template used to build the container image.
	// Default: - a standard dockerfile template will be generated to pull the base image, perform environment setup, and
	// run all components in the recipe.
	//
	// Experimental.
	Dockerfile DockerfileData `field:"optional" json:"dockerfile" yaml:"dockerfile"`
	// The block devices to attach to the instance used for building, testing, and distributing the container image.
	// Default: the block devices of the instance image will be used.
	//
	// Experimental.
	InstanceBlockDevices *[]*awsec2.BlockDevice `field:"optional" json:"instanceBlockDevices" yaml:"instanceBlockDevices"`
	// The image to use to launch the instance used for building, testing, and distributing the container image.
	// Default: Image Builder will use the appropriate ECS-optimized AMI.
	//
	// Experimental.
	InstanceImage ContainerInstanceImage `field:"optional" json:"instanceImage" yaml:"instanceImage"`
	// The KMS key used to encrypt the dockerfile template.
	// Default: None.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The operating system (OS) version of the base image.
	// Default: - Image Builder will determine the OS version of the base image, if sourced from a third-party container
	// registry. Otherwise, the OS version of the base image is required.
	//
	// Experimental.
	OsVersion OSVersion `field:"optional" json:"osVersion" yaml:"osVersion"`
	// The tags to apply to the container recipe.
	// Default: None.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The working directory for use during build and test workflows.
	// Default: - the Image Builder default working directory is used. For Linux and macOS builds, this would be /tmp. For
	// Windows builds, this would be C:/.
	//
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

