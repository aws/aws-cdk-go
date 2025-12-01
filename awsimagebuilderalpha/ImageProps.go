package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for creating an Image resource.
//
// Example:
//   containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("MyContainerRecipe"), &ContainerRecipeProps{
//   	BaseImage: imagebuilder.BaseContainerImage_*FromDockerHub(jsii.String("amazonlinux"), jsii.String("latest")),
//   	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
//   })
//
//   containerImage := imagebuilder.NewImage(this, jsii.String("MyContainerImage"), &ImageProps{
//   	Recipe: containerRecipe,
//   })
//
// Experimental.
type ImageProps struct {
	// The recipe that defines the base image, components, and customizations used to build the image.
	//
	// This can either be
	// an image recipe, or a container recipe.
	// Experimental.
	Recipe IRecipeBase `field:"required" json:"recipe" yaml:"recipe"`
	// The execution role to use for deleting the image as well as the underlying resources, such as the AMIs, snapshots, and containers.
	//
	// This role should contain resource lifecycle permissions required to delete the underlying
	// AMIs/containers.
	// Default: - no execution role. Only the Image Builder image will be deleted.
	//
	// Experimental.
	DeletionExecutionRole awsiam.IRole `field:"optional" json:"deletionExecutionRole" yaml:"deletionExecutionRole"`
	// The distribution configuration used for distributing the image.
	// Default: None.
	//
	// Experimental.
	DistributionConfiguration IDistributionConfiguration `field:"optional" json:"distributionConfiguration" yaml:"distributionConfiguration"`
	// If enabled, collects additional information about the image being created, including the operating system (OS) version and package list for the AMI.
	// Default: true.
	//
	// Experimental.
	EnhancedImageMetadataEnabled *bool `field:"optional" json:"enhancedImageMetadataEnabled" yaml:"enhancedImageMetadataEnabled"`
	// The execution role used to perform workflow actions to build the image.
	//
	// By default, the Image Builder Service Linked Role (SLR) will be created automatically and used as the execution
	// role. However, when providing a custom set of image workflows for the image, an execution role will be
	// generated with the minimal permissions needed to execute the workflows.
	// Default: - Image Builder will use the SLR if possible. Otherwise, an execution role will be generated
	//
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The container repository that Amazon Inspector scans to identify findings for your container images.
	//
	// If a
	// repository is not provided, Image Builder creates a repository named `image-builder-image-scanning-repository`
	// for vulnerability scanning.
	// Default: - if scanning is enabled, a repository will be created by Image Builder if one is not provided.
	//
	// Experimental.
	ImageScanningEcrRepository awsecr.IRepository `field:"optional" json:"imageScanningEcrRepository" yaml:"imageScanningEcrRepository"`
	// The tags for Image Builder to apply to the output container image that Amazon Inspector scans.
	// Default: None.
	//
	// Experimental.
	ImageScanningEcrTags *[]*string `field:"optional" json:"imageScanningEcrTags" yaml:"imageScanningEcrTags"`
	// Indicates whether Image Builder keeps a snapshot of the vulnerability scans that Amazon Inspector runs against the build instance when you create a new image.
	// Default: false.
	//
	// Experimental.
	ImageScanningEnabled *bool `field:"optional" json:"imageScanningEnabled" yaml:"imageScanningEnabled"`
	// Whether to run tests after building an image.
	// Default: true.
	//
	// Experimental.
	ImageTestsEnabled *bool `field:"optional" json:"imageTestsEnabled" yaml:"imageTestsEnabled"`
	// The infrastructure configuration used for building the image.
	//
	// A default infrastructure configuration will be used if one is not provided.
	//
	// The default configuration will create an instance profile and role with minimal permissions needed to build the
	// image, attached to the EC2 instance.
	//
	// IMDSv2 will be required by default on the instances used to build and test the image.
	// Default: - an infrastructure configuration will be created with the default settings.
	//
	// Experimental.
	InfrastructureConfiguration IInfrastructureConfiguration `field:"optional" json:"infrastructureConfiguration" yaml:"infrastructureConfiguration"`
	// The log group to use for the image.
	//
	// By default, a log group will be created with the format
	// `/aws/imagebuilder/<image-name>`.
	// Default: - a log group will be created.
	//
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The tags to apply to the image.
	// Default: None.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The list of workflow configurations used to build the image.
	// Default: - Image Builder will use a default set of workflows for the build to build, test, and distribute the
	// image.
	//
	// Experimental.
	Workflows *[]*WorkflowConfiguration `field:"optional" json:"workflows" yaml:"workflows"`
}

