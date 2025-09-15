package awsec2


// A reference to a NetworkAcl resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkAclReference := &NetworkAclReference{
//   	NetworkAclId: jsii.String("networkAclId"),
//   }
//
type NetworkAclReference struct {
	// The Id of the NetworkAcl resource.
	NetworkAclId *string `field:"required" json:"networkAclId" yaml:"networkAclId"`
}

