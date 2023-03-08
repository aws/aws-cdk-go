package awsimagebuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecrConfigurationProperty := &ecrConfigurationProperty{
//   	containerTags: []*string{
//   		jsii.String("containerTags"),
//   	},
//   	repositoryName: jsii.String("repositoryName"),
//   }
//
type CfnImage_EcrConfigurationProperty struct {
	// `CfnImage.EcrConfigurationProperty.ContainerTags`.
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// `CfnImage.EcrConfigurationProperty.RepositoryName`.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}

