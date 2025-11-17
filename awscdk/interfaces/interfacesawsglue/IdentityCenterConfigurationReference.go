package interfacesawsglue


// A reference to a IdentityCenterConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityCenterConfigurationReference := &IdentityCenterConfigurationReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type IdentityCenterConfigurationReference struct {
	// The AccountId of the IdentityCenterConfiguration resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

