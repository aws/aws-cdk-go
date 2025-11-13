package interfacesawsiam


// A reference to a Policy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyReference := &PolicyReference{
//   	PolicyId: jsii.String("policyId"),
//   }
//
type PolicyReference struct {
	// The Id of the Policy resource.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
}

