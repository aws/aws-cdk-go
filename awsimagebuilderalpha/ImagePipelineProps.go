package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for creating an Image Pipeline resource.
//
// Example:
//   workflowPipeline := imagebuilder.NewImagePipeline(this, jsii.String("WorkflowPipeline"), &ImagePipelineProps{
//   	Recipe: exampleImageRecipe,
//   	Workflows: []WorkflowConfiguration{
//   		&WorkflowConfiguration{
//   			Workflow: imagebuilder.AmazonManagedWorkflow_BuildImage(this, jsii.String("BuildWorkflow")),
//   		},
//   		&WorkflowConfiguration{
//   			Workflow: imagebuilder.AmazonManagedWorkflow_TestImage(this, jsii.String("TestWorkflow")),
//   		},
//   	},
//   })
//
// Experimental.
type ImagePipelineProps struct {
	// The recipe that defines the base image, components, and customizations used to build the image.
	//
	// This can either be
	// an image recipe, or a container recipe.
	// Experimental.
	Recipe IRecipeBase `field:"required" json:"recipe" yaml:"recipe"`
	// The description of the image pipeline.
	// Default: None.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
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
	// The execution role used to perform workflow actions to build this image.
	//
	// By default, the Image Builder Service Linked Role (SLR) will be created automatically and used as the execution
	// role. However, when providing a custom set of image workflows for the pipeline, an execution role will be
	// generated with the minimal permissions needed to execute the workflows.
	// Default: - Image Builder will use the SLR if possible. Otherwise, an execution role will be generated
	//
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The log group to use for images created from the image pipeline.
	//
	// By default, a log group will be created with the
	// format `/aws/imagebuilder/<image-name>`.
	// Default: - a log group will be created.
	//
	// Experimental.
	ImageLogGroup awslogs.ILogGroup `field:"optional" json:"imageLogGroup" yaml:"imageLogGroup"`
	// The log group to use for the image pipeline.
	//
	// By default, a log group will be created with the format
	// `/aws/imagebuilder/pipeline/<pipeline-name>`.
	// Default: - a log group will be created.
	//
	// Experimental.
	ImagePipelineLogGroup awslogs.ILogGroup `field:"optional" json:"imagePipelineLogGroup" yaml:"imagePipelineLogGroup"`
	// The name of the image pipeline.
	// Default: - a name is generated.
	//
	// Experimental.
	ImagePipelineName *string `field:"optional" json:"imagePipelineName" yaml:"imagePipelineName"`
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
	// Default: - an infrastructure configuration will be created with the default settings.
	//
	// Experimental.
	InfrastructureConfiguration IInfrastructureConfiguration `field:"optional" json:"infrastructureConfiguration" yaml:"infrastructureConfiguration"`
	// The schedule of the image pipeline.
	//
	// This configures how often and when a pipeline automatically creates a new
	// image.
	// Default: - none, a manual image pipeline will be created.
	//
	// Experimental.
	Schedule *ImagePipelineSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Indicates whether the pipeline is enabled to be triggered by the provided schedule.
	// Default: ImagePipelineStatus.ENABLED
	//
	// Experimental.
	Status ImagePipelineStatus `field:"optional" json:"status" yaml:"status"`
	// The tags to apply to the image pipeline.
	// Default: None.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The list of workflow configurations used to build the image.
	// Default: - Image Builder will use a default set of workflows for the build to build, test, and distribute the image.
	//
	// Experimental.
	Workflows *[]*WorkflowConfiguration `field:"optional" json:"workflows" yaml:"workflows"`
}

