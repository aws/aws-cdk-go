package interfacesawswisdom


// A reference to a Content resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentReference := &ContentReference{
//   	ContentArn: jsii.String("contentArn"),
//   }
//
type ContentReference struct {
	// The ContentArn of the Content resource.
	ContentArn *string `field:"required" json:"contentArn" yaml:"contentArn"`
}

