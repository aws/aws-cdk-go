package awsimagebuilder


// Container distribution settings for encryption, licensing, and sharing in a specific Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerDistributionConfigurationProperty := &containerDistributionConfigurationProperty{
//   	containerTags: []*string{
//   		jsii.String("containerTags"),
//   	},
//   	description: jsii.String("description"),
//   	targetRepository: &targetContainerRepositoryProperty{
//   		repositoryName: jsii.String("repositoryName"),
//   		service: jsii.String("service"),
//   	},
//   }
//
type CfnDistributionConfiguration_ContainerDistributionConfigurationProperty struct {
	// Tags that are attached to the container distribution configuration.
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// The description of the container distribution configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The destination repository for the container distribution configuration.
	TargetRepository interface{} `field:"optional" json:"targetRepository" yaml:"targetRepository"`
}

