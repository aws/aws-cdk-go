package awsimagebuilder


// The container repository where the output container image is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetContainerRepositoryProperty := &targetContainerRepositoryProperty{
//   	repositoryName: jsii.String("repositoryName"),
//   	service: jsii.String("service"),
//   }
//
type CfnContainerRecipe_TargetContainerRepositoryProperty struct {
	// The name of the container repository where the output container image is stored.
	//
	// This name is prefixed by the repository location.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Specifies the service in which this image was registered.
	Service *string `field:"optional" json:"service" yaml:"service"`
}

