package awsimagebuilder


// Properties for defining a `CfnImagePipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImagePipelineProps := &cfnImagePipelineProps{
//   	infrastructureConfigurationArn: jsii.String("infrastructureConfigurationArn"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	containerRecipeArn: jsii.String("containerRecipeArn"),
//   	description: jsii.String("description"),
//   	distributionConfigurationArn: jsii.String("distributionConfigurationArn"),
//   	enhancedImageMetadataEnabled: jsii.Boolean(false),
//   	imageRecipeArn: jsii.String("imageRecipeArn"),
//   	imageTestsConfiguration: &imageTestsConfigurationProperty{
//   		imageTestsEnabled: jsii.Boolean(false),
//   		timeoutMinutes: jsii.Number(123),
//   	},
//   	schedule: &scheduleProperty{
//   		pipelineExecutionStartCondition: jsii.String("pipelineExecutionStartCondition"),
//   		scheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	status: jsii.String("status"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnImagePipelineProps struct {
	// The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.
	InfrastructureConfigurationArn *string `field:"required" json:"infrastructureConfigurationArn" yaml:"infrastructureConfigurationArn"`
	// The name of the image pipeline.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the container recipe that is used for this pipeline.
	ContainerRecipeArn *string `field:"optional" json:"containerRecipeArn" yaml:"containerRecipeArn"`
	// The description of this image pipeline.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.
	DistributionConfigurationArn *string `field:"optional" json:"distributionConfigurationArn" yaml:"distributionConfigurationArn"`
	// Collects additional information about the image being created, including the operating system (OS) version and package list.
	//
	// This information is used to enhance the overall experience of using EC2 Image Builder. Enabled by default.
	EnhancedImageMetadataEnabled interface{} `field:"optional" json:"enhancedImageMetadataEnabled" yaml:"enhancedImageMetadataEnabled"`
	// The Amazon Resource Name (ARN) of the image recipe associated with this image pipeline.
	ImageRecipeArn *string `field:"optional" json:"imageRecipeArn" yaml:"imageRecipeArn"`
	// The configuration of the image tests that run after image creation to ensure the quality of the image that was created.
	ImageTestsConfiguration interface{} `field:"optional" json:"imageTestsConfiguration" yaml:"imageTestsConfiguration"`
	// The schedule of the image pipeline.
	//
	// A schedule configures how often and when a pipeline automatically creates a new image.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// The status of the image pipeline.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags of this image pipeline.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

