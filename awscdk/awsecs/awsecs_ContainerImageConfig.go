package awsecs


// The configuration for creating a container image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerImageConfig := &containerImageConfig{
//   	imageName: jsii.String("imageName"),
//
//   	// the properties below are optional
//   	repositoryCredentials: &repositoryCredentialsProperty{
//   		credentialsParameter: jsii.String("credentialsParameter"),
//   	},
//   }
//
type ContainerImageConfig struct {
	// Specifies the name of the container image.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// Specifies the credentials used to access the image repository.
	RepositoryCredentials *CfnTaskDefinition_RepositoryCredentialsProperty `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
}

