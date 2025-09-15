package awsapplicationsignals


// A reference to a Discovery resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   discoveryReference := &DiscoveryReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type DiscoveryReference struct {
	// The AccountId of the Discovery resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

