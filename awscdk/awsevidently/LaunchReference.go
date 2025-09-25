package awsevidently


// A reference to a Launch resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchReference := &LaunchReference{
//   	LaunchArn: jsii.String("launchArn"),
//   }
//
type LaunchReference struct {
	// The Arn of the Launch resource.
	LaunchArn *string `field:"required" json:"launchArn" yaml:"launchArn"`
}

