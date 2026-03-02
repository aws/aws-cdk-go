package interfacesawsec2


// A reference to a IPAMPrefixListResolver resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPAMPrefixListResolverReference := map[string]*string{
//   	"ipamPrefixListResolverArn": jsii.String("ipamPrefixListResolverArn"),
//   	"ipamPrefixListResolverId": jsii.String("ipamPrefixListResolverId"),
//   }
//
type IPAMPrefixListResolverReference struct {
	// The ARN of the IPAMPrefixListResolver resource.
	IpamPrefixListResolverArn *string `field:"required" json:"ipamPrefixListResolverArn" yaml:"ipamPrefixListResolverArn"`
	// The IpamPrefixListResolverId of the IPAMPrefixListResolver resource.
	IpamPrefixListResolverId *string `field:"required" json:"ipamPrefixListResolverId" yaml:"ipamPrefixListResolverId"`
}

