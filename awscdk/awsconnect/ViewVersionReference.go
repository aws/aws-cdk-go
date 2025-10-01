package awsconnect


// A reference to a ViewVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   viewVersionReference := &ViewVersionReference{
//   	ViewVersionArn: jsii.String("viewVersionArn"),
//   }
//
type ViewVersionReference struct {
	// The ViewVersionArn of the ViewVersion resource.
	ViewVersionArn *string `field:"required" json:"viewVersionArn" yaml:"viewVersionArn"`
}

