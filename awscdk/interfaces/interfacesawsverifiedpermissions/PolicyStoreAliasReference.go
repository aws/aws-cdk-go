package interfacesawsverifiedpermissions


// A reference to a PolicyStoreAlias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyStoreAliasReference := &PolicyStoreAliasReference{
//   	AliasName: jsii.String("aliasName"),
//   }
//
type PolicyStoreAliasReference struct {
	// The AliasName of the PolicyStoreAlias resource.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
}

