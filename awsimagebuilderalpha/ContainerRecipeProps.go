package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a Container Recipe resource.
//
// Example:
//   sourceRepo := ecr.Repository_FromRepositoryName(this, jsii.String("SourceRepo"), jsii.String("my-base-image"))
//   targetRepo := ecr.Repository_FromRepositoryName(this, jsii.String("TargetRepo"), jsii.String("my-container-repo"))
//
//   containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("EcrContainerRecipe"), &ContainerRecipeProps{
//   	BaseImage: imagebuilder.BaseContainerImage_FromEcr(sourceRepo, jsii.String("1.0.0")),
//   	TargetRepository: imagebuilder.Repository_FromEcr(targetRepo),
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

