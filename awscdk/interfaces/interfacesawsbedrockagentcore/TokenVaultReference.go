package interfacesawsbedrockagentcore


// A reference to a TokenVault resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tokenVaultReference := &TokenVaultReference{
//   	TokenVaultArn: jsii.String("tokenVaultArn"),
//   }
//
type TokenVaultReference struct {
	// The Arn of the TokenVault resource.
	TokenVaultArn *string `field:"required" json:"tokenVaultArn" yaml:"tokenVaultArn"`
}

