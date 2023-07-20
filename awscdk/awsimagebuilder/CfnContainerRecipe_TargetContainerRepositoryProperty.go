package awsimagebuilder


// The container repository where the output container image is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetContainerRepositoryProperty := &TargetContainerRepositoryProperty{
//   	RepositoryName: jsii.String("repositoryName"),
//   	Service: jsii.String("service"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-targetcontainerrepository.html
//
type CfnContainerRecipe_TargetContainerRepositoryProperty struct {
	// The name of the container repository where the output container image is stored.
	//
	// This name is prefixed by the repository location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-targetcontainerrepository.html#cfn-imagebuilder-containerrecipe-targetcontainerrepository-repositoryname
	//
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Specifies the service in which this image was registered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-targetcontainerrepository.html#cfn-imagebuilder-containerrecipe-targetcontainerrepository-service
	//
	Service *string `field:"optional" json:"service" yaml:"service"`
}

