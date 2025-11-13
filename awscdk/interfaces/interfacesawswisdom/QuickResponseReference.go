package interfacesawswisdom


// A reference to a QuickResponse resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quickResponseReference := &QuickResponseReference{
//   	QuickResponseArn: jsii.String("quickResponseArn"),
//   }
//
type QuickResponseReference struct {
	// The QuickResponseArn of the QuickResponse resource.
	QuickResponseArn *string `field:"required" json:"quickResponseArn" yaml:"quickResponseArn"`
}

