package awsresiliencehub


// A reference to a App resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appReference := &AppReference{
//   	AppArn: jsii.String("appArn"),
//   }
//
type AppReference struct {
	// The AppArn of the App resource.
	AppArn *string `field:"required" json:"appArn" yaml:"appArn"`
}

