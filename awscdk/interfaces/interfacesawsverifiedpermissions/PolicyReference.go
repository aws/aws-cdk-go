package interfacesawsverifiedpermissions


// A reference to a Policy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyReference := &PolicyReference{
//   	PolicyId: jsii.String("policyId"),
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   }
//
type PolicyReference struct {
	// The PolicyId of the Policy resource.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// The PolicyStoreId of the Policy resource.
	PolicyStoreId *string `field:"required" json:"policyStoreId" yaml:"policyStoreId"`
}

