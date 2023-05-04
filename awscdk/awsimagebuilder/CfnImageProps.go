package awsimagebuilder


// Properties for defining a `CfnImage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageProps := &CfnImageProps{
//   	InfrastructureConfigurationArn: jsii.String("infrastructureConfigurationArn"),
//
//   	// the properties below are optional
//   	ContainerRecipeArn: jsii.String("containerRecipeArn"),
//   	DistributionConfigurationArn: jsii.String("distributionConfigurationArn"),
//   	EnhancedImageMetadataEnabled: jsii.Boolean(false),
//   	ImageRecipeArn: jsii.String("imageRecipeArn"),
//   	ImageScanningConfiguration: &ImageScanningConfigurationProperty{
//   		EcrConfiguration: &EcrConfigurationProperty{
//   			ContainerTags: []*string{
//   				jsii.String("containerTags"),
//   			},
//   			RepositoryName: jsii.String("repositoryName"),
//   		},
//   		ImageScanningEnabled: jsii.Boolean(false),
//   	},
//   	ImageTestsConfiguration: &ImageTestsConfigurationProperty{
//   		ImageTestsEnabled: jsii.Boolean(false),
//   		TimeoutMinutes: jsii.Number(123),
//   	},
//   	Tags: map[string]*string{
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

