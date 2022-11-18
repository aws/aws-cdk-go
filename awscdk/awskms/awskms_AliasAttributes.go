package awskms


// Properties of a reference to an existing KMS Alias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//
//   aliasAttributes := &aliasAttributes{
//   	aliasName: jsii.String("aliasName"),
//   	aliasTargetKey: key,
//   }
//
type AliasAttributes struct {
	// Specifies the alias name.
	//
	// This value must begin with alias/ followed by a name (i.e. alias/ExampleAlias)
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// The customer master key (CMK) to which the Alias refers.
	AliasTargetKey IKey `field:"required" json:"aliasTargetKey" yaml:"aliasTargetKey"`
}

