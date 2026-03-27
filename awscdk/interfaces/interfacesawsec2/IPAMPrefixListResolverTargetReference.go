package interfacesawsec2


// A reference to a IPAMPrefixListResolverTarget resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPAMPrefixListResolverTargetReference := map[string]*string{
//   	"ipamPrefixListResolverTargetArn": jsii.String("ipamPrefixListResolverTargetArn"),
//   	"ipamPrefixListResolverTargetId": jsii.String("ipamPrefixListResolverTargetId"),
//   }
//
type IPAMPrefixListResolverTargetReference struct {
	// The ARN of the IPAMPrefixListResolverTarget resource.
	IpamPrefixListResolverTargetArn *string `field:"required" json:"ipamPrefixListResolverTargetArn" yaml:"ipamPrefixListResolverTargetArn"`
	// The IpamPrefixListResolverTargetId of the IPAMPrefixListResolverTarget resource.
	IpamPrefixListResolverTargetId *string `field:"required" json:"ipamPrefixListResolverTargetId" yaml:"ipamPrefixListResolverTargetId"`
}

