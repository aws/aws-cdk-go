package awsimagebuilder


// Container distribution settings for encryption, licensing, and sharing in a specific Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerDistributionConfigurationProperty := &ContainerDistributionConfigurationProperty{
//   	ContainerTags: []*string{
//   		jsii.String("containerTags"),
//   	},
//   	Description: jsii.String("description"),
//   	TargetRepository: &TargetContainerRepositoryProperty{
//   		RepositoryName: jsii.String("repositoryName"),
//   		Service: jsii.String("service"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-containerdistributionconfiguration.html
//
type CfnDistributionConfiguration_ContainerDistributionConfigurationProperty struct {
	// Tags that are attached to the container distribution configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-containerdistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-containerdistributionconfiguration-containertags
	//
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// The description of the container distribution configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-containerdistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-containerdistributionconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The destination repository for the container distribution configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-containerdistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-containerdistributionconfiguration-targetrepository
	//
	TargetRepository interface{} `field:"optional" json:"targetRepository" yaml:"targetRepository"`
}

