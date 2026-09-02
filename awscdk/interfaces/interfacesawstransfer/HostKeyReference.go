package interfacesawstransfer


// A reference to a HostKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostKeyReference := &HostKeyReference{
//   	HostKeyArn: jsii.String("hostKeyArn"),
//   }
//
type HostKeyReference struct {
	// The Arn of the HostKey resource.
	HostKeyArn *string `field:"required" json:"hostKeyArn" yaml:"hostKeyArn"`
}

