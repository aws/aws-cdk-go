package interfacesawsverifiedpermissions


// A reference to a PolicyTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyTemplateReference := &PolicyTemplateReference{
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   	PolicyTemplateId: jsii.String("policyTemplateId"),
//   }
//
type PolicyTemplateReference struct {
	// The PolicyStoreId of the PolicyTemplate resource.
	PolicyStoreId *string `field:"required" json:"policyStoreId" yaml:"policyStoreId"`
	// The PolicyTemplateId of the PolicyTemplate resource.
	PolicyTemplateId *string `field:"required" json:"policyTemplateId" yaml:"policyTemplateId"`
}

