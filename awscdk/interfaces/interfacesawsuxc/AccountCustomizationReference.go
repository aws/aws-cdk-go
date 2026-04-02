package interfacesawsuxc


// A reference to a AccountCustomization resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountCustomizationReference := &AccountCustomizationReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type AccountCustomizationReference struct {
	// The AccountId of the AccountCustomization resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

