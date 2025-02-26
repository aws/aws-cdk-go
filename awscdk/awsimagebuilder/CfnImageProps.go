package awsimagebuilder


// Properties for defining a `CfnImage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageProps := &CfnImageProps{
//   	ContainerRecipeArn: jsii.String("containerRecipeArn"),
//   	DistributionConfigurationArn: jsii.String("distributionConfigurationArn"),
//   	EnhancedImageMetadataEnabled: jsii.Boolean(false),
//   	ExecutionRole: jsii.String("executionRole"),
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
//   	InfrastructureConfigurationArn: jsii.String("infrastructureConfigurationArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Workflows: []interface{}{
//   		&WorkflowConfigurationProperty{
//   			OnFailure: jsii.String("onFailure"),
//   			ParallelGroup: jsii.String("parallelGroup"),
//   			Parameters: []interface{}{
//   				&WorkflowParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   			WorkflowArn: jsii.String("workflowArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-image.html
//
type CfnImageProps struct {
	// The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-image.html#cfn-imagebuilder-image-containerrecipearn
	//
	ContainerRecipeArn *string `field:"optional" json:"containerRecipeArn" yaml:"containerRecipeArn"`
	// The Amazon Resource Name (ARN) of the distribution configuration that defines and configures the outputs of your pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-image.html#cfn-imagebuilder-image-distributionconfigurationarn
	//
	DistributionConfigurationArn *string `field:"optional" json:"distributionConfigurationArn" yaml:"distributionConfigurationArn"`
	// Collects additional information about the image being created, including the operating system (OS) version and package list.
	//
	// This information is used to enhance the overall experience of using EC2 Image Builder. Enabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-image.html#cfn-imagebuilder-image-enhancedimagemetadataenabled
	//
	EnhancedImageMetadataEnabled interface{} `field:"optional" json:"enhancedImageMetadataEnabled" yaml:"enhancedImageMetadataEnabled"`
	// The name or Amazon Resource Name (ARN) for the IAM role you create that grants Image Builder access to perform workflow actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-image.html#cfn-imagebuilder-image-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-image.html#cfn-imagebuilder-image-imagerecipearn
	//
	ImageRecipeArn *string `field:"optional" json:"imageRecipeArn" yaml:"imageRecipeArn"`
	// Contains settings for vulnerability scans.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-image.html#cfn-imagebuilder-image-imagescanningconfiguration
	//
	ImageScanningConfiguration interface{} `field:"optional" json:"imageScanningConfiguration" yaml:"imageScanningConfiguration"`
	// The image tests configuration of the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-image.html#cfn-imagebuilder-image-imagetestsconfiguration
	//
	ImageTestsConfiguration interface{} `field:"optional" json:"imageTestsConfiguration" yaml:"imageTestsConfiguration"`
	// The Amazon Resource Name (ARN) of the infrastructure configuration that defines the environment in which your image will be built and tested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-image.html#cfn-imagebuilder-image-infrastructureconfigurationarn
	//
	InfrastructureConfigurationArn *string `field:"optional" json:"infrastructureConfigurationArn" yaml:"infrastructureConfigurationArn"`
	// The tags of the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-image.html#cfn-imagebuilder-image-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Contains an array of workflow configuration objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-image.html#cfn-imagebuilder-image-workflows
	//
	Workflows interface{} `field:"optional" json:"workflows" yaml:"workflows"`
}

