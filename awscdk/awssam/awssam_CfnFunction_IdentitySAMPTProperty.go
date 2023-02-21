package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identitySAMPTProperty := &identitySAMPTProperty{
//   	identityName: jsii.String("identityName"),
//   }
//
type CfnFunction_IdentitySAMPTProperty struct {
	// `CfnFunction.IdentitySAMPTProperty.IdentityName`.
	IdentityName *string `field:"required" json:"identityName" yaml:"identityName"`
}

