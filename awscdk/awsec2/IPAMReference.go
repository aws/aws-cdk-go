package awsec2


// A reference to a IPAM resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPAMReference := map[string]*string{
//   	"ipamArn": jsii.String("ipamArn"),
//   	"ipamId": jsii.String("ipamId"),
//   }
//
type IPAMReference struct {
	// The ARN of the IPAM resource.
	IpamArn *string `field:"required" json:"ipamArn" yaml:"ipamArn"`
	// The IpamId of the IPAM resource.
	IpamId *string `field:"required" json:"ipamId" yaml:"ipamId"`
}

