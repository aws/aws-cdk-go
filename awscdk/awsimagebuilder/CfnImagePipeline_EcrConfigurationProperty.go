package awsimagebuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecrConfigurationProperty := &EcrConfigurationProperty{
//   	ContainerTags: []*string{
//   		jsii.String("containerTags"),
//   	},
//   	RepositoryName: jsii.String("repositoryName"),
//   }
//
type CfnImagePipeline_EcrConfigurationProperty struct {
	// `CfnImagePipeline.EcrConfigurationProperty.ContainerTags`.
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// `CfnImagePipeline.EcrConfigurationProperty.RepositoryName`.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}

