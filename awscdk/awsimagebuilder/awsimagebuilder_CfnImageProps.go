package awsimagebuilder


// Properties for defining a `CfnImage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageProps := &cfnImageProps{
//   	infrastructureConfigurationArn: jsii.String("infrastructureConfigurationArn"),
//
//   	// the properties below are optional
//   	containerRecipeArn: jsii.String("containerRecipeArn"),
//   	distributionConfigurationArn: jsii.String("distributionConfigurationArn"),
//   	enhancedImageMetadataEnabled: jsii.Boolean(false),
//   	imageRecipeArn: jsii.String("imageRecipeArn"),
//   	imageScanningConfiguration: &imageScanningConfigurationProperty{
//   		ecrConfiguration: &ecrConfigurationProperty{
//   			containerTags: []*string{
//   				jsii.String("containerTags"),
//   			},
//   			repositoryName: jsii.String("repositoryName"),
//   		},
//   		imageScanningEnabled: jsii.Boolean(false),
//   	},
//   	imageTestsConfiguration: &imageTestsConfigurationProperty{
//   		imageTestsEnabled: jsii.Boolean(false),
//   		timeoutMinutes: jsii.Number(123),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnImageProps struct {
	// The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.
	InfrastructureConfigurationArn *string `field:"required" json:"infrastructureConfigurationArn" yaml:"infrastructureConfigurationArn"`
	// The Amazon Resource Name (ARN) of the container recipe that is used for this pipeline.
	ContainerRecipeArn *string `field:"optional" json:"containerRecipeArn" yaml:"containerRecipeArn"`
	// The Amazon Resource Name (ARN) of the distribution configuration.
	DistributionConfigurationArn *string `field:"optional" json:"distributionConfigurationArn" yaml:"distributionConfigurationArn"`
	// Indicates whether Image Builder collects additional information about the image, such as the operating system (OS) version and package list.
	EnhancedImageMetadataEnabled interface{} `field:"optional" json:"enhancedImageMetadataEnabled" yaml:"enhancedImageMetadataEnabled"`
	// The Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn *string `field:"optional" json:"imageRecipeArn" yaml:"imageRecipeArn"`
	// `AWS::ImageBuilder::Image.ImageScanningConfiguration`.
	ImageScanningConfiguration interface{} `field:"optional" json:"imageScanningConfiguration" yaml:"imageScanningConfiguration"`
	// The configuration settings for your image test components, which includes a toggle that allows you to turn off tests, and a timeout setting.
	ImageTestsConfiguration interface{} `field:"optional" json:"imageTestsConfiguration" yaml:"imageTestsConfiguration"`
	// The tags of the image.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

