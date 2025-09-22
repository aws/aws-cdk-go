package awsec2


// A reference to a IPAMResourceDiscovery resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPAMResourceDiscoveryReference := map[string]*string{
//   	"ipamResourceDiscoveryId": jsii.String("ipamResourceDiscoveryId"),
//   }
//
type IPAMResourceDiscoveryReference struct {
	// The IpamResourceDiscoveryId of the IPAMResourceDiscovery resource.
	IpamResourceDiscoveryId *string `field:"required" json:"ipamResourceDiscoveryId" yaml:"ipamResourceDiscoveryId"`
}

