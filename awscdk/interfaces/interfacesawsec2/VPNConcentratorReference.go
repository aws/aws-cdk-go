package interfacesawsec2


// A reference to a VPNConcentrator resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPNConcentratorReference := &VPNConcentratorReference{
//   	VpnConcentratorId: jsii.String("vpnConcentratorId"),
//   }
//
type VPNConcentratorReference struct {
	// The VpnConcentratorId of the VPNConcentrator resource.
	VpnConcentratorId *string `field:"required" json:"vpnConcentratorId" yaml:"vpnConcentratorId"`
}

