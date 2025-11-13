package interfacesawsec2


// A reference to a IPAMPoolCidr resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPAMPoolCidrReference := map[string]*string{
//   	"ipamPoolCidrId": jsii.String("ipamPoolCidrId"),
//   	"ipamPoolId": jsii.String("ipamPoolId"),
//   }
//
type IPAMPoolCidrReference struct {
	// The IpamPoolCidrId of the IPAMPoolCidr resource.
	IpamPoolCidrId *string `field:"required" json:"ipamPoolCidrId" yaml:"ipamPoolCidrId"`
	// The IpamPoolId of the IPAMPoolCidr resource.
	IpamPoolId *string `field:"required" json:"ipamPoolId" yaml:"ipamPoolId"`
}

