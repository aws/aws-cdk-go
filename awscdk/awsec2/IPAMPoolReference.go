package awsec2


// A reference to a IPAMPool resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPAMPoolReference := map[string]*string{
//   	"ipamPoolArn": jsii.String("ipamPoolArn"),
//   	"ipamPoolId": jsii.String("ipamPoolId"),
//   }
//
type IPAMPoolReference struct {
	// The ARN of the IPAMPool resource.
	IpamPoolArn *string `field:"required" json:"ipamPoolArn" yaml:"ipamPoolArn"`
	// The IpamPoolId of the IPAMPool resource.
	IpamPoolId *string `field:"required" json:"ipamPoolId" yaml:"ipamPoolId"`
}

