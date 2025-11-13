package interfacesawsgamelift


// A reference to a Build resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   buildReference := &BuildReference{
//   	BuildArn: jsii.String("buildArn"),
//   	BuildId: jsii.String("buildId"),
//   }
//
type BuildReference struct {
	// The ARN of the Build resource.
	BuildArn *string `field:"required" json:"buildArn" yaml:"buildArn"`
	// The BuildId of the Build resource.
	BuildId *string `field:"required" json:"buildId" yaml:"buildId"`
}

