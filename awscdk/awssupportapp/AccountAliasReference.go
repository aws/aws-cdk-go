package awssupportapp


// A reference to a AccountAlias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountAliasReference := &AccountAliasReference{
//   	AccountAliasResourceId: jsii.String("accountAliasResourceId"),
//   }
//
type AccountAliasReference struct {
	// The AccountAliasResourceId of the AccountAlias resource.
	AccountAliasResourceId *string `field:"required" json:"accountAliasResourceId" yaml:"accountAliasResourceId"`
}

