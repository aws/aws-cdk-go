package awsec2


// A reference to a VPCBlockPublicAccessOptions resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCBlockPublicAccessOptionsReference := &VPCBlockPublicAccessOptionsReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type VPCBlockPublicAccessOptionsReference struct {
	// The AccountId of the VPCBlockPublicAccessOptions resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

