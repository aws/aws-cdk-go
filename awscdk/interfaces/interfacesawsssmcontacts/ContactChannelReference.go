package interfacesawsssmcontacts


// A reference to a ContactChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactChannelReference := &ContactChannelReference{
//   	ContactChannelArn: jsii.String("contactChannelArn"),
//   }
//
type ContactChannelReference struct {
	// The Arn of the ContactChannel resource.
	ContactChannelArn *string `field:"required" json:"contactChannelArn" yaml:"contactChannelArn"`
}

