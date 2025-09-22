package awsworkspaces


// A reference to a ConnectionAlias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionAliasReference := &ConnectionAliasReference{
//   	AliasId: jsii.String("aliasId"),
//   }
//
type ConnectionAliasReference struct {
	// The AliasId of the ConnectionAlias resource.
	AliasId *string `field:"required" json:"aliasId" yaml:"aliasId"`
}

