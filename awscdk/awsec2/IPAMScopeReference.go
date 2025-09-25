package awsec2


// A reference to a IPAMScope resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPAMScopeReference := map[string]*string{
//   	"ipamScopeArn": jsii.String("ipamScopeArn"),
//   	"ipamScopeId": jsii.String("ipamScopeId"),
//   }
//
type IPAMScopeReference struct {
	// The ARN of the IPAMScope resource.
	IpamScopeArn *string `field:"required" json:"ipamScopeArn" yaml:"ipamScopeArn"`
	// The IpamScopeId of the IPAMScope resource.
	IpamScopeId *string `field:"required" json:"ipamScopeId" yaml:"ipamScopeId"`
}

