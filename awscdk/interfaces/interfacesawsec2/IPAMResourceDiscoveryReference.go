package interfacesawsec2


// A reference to a IPAMResourceDiscovery resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPAMResourceDiscoveryReference := map[string]*string{
//   	"ipamResourceDiscoveryArn": jsii.String("ipamResourceDiscoveryArn"),
//   	"ipamResourceDiscoveryId": jsii.String("ipamResourceDiscoveryId"),
//   }
//
type IPAMResourceDiscoveryReference struct {
	// The ARN of the IPAMResourceDiscovery resource.
	IpamResourceDiscoveryArn *string `field:"required" json:"ipamResourceDiscoveryArn" yaml:"ipamResourceDiscoveryArn"`
	// The IpamResourceDiscoveryId of the IPAMResourceDiscovery resource.
	IpamResourceDiscoveryId *string `field:"required" json:"ipamResourceDiscoveryId" yaml:"ipamResourceDiscoveryId"`
}

