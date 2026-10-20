package interfacesawsec2


// A reference to a SecondaryNetwork resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secondaryNetworkReference := &SecondaryNetworkReference{
//   	SecondaryNetworkArn: jsii.String("secondaryNetworkArn"),
//   }
//
type SecondaryNetworkReference struct {
	// The SecondaryNetworkArn of the SecondaryNetwork resource.
	SecondaryNetworkArn *string `field:"required" json:"secondaryNetworkArn" yaml:"secondaryNetworkArn"`
}

