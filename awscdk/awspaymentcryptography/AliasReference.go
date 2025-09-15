package awspaymentcryptography


// A reference to a Alias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aliasReference := &AliasReference{
//   	AliasName: jsii.String("aliasName"),
//   }
//
type AliasReference struct {
	// The AliasName of the Alias resource.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
}

