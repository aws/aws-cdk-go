package interfacesawsiotsecuretunneling


// A reference to a Tunnel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tunnelReference := &TunnelReference{
//   	TunnelArn: jsii.String("tunnelArn"),
//   }
//
type TunnelReference struct {
	// The TunnelArn of the Tunnel resource.
	TunnelArn *string `field:"required" json:"tunnelArn" yaml:"tunnelArn"`
}

