package interfacesawslogs


// A reference to a StorageTierPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageTierPolicyReference := &StorageTierPolicyReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type StorageTierPolicyReference struct {
	// The AccountId of the StorageTierPolicy resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

