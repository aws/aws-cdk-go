package awseks


// A key-value pair that describes a required claim in the identity token.
//
// If set, each claim is verified to be present in the token with a matching value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requiredClaimProperty := &requiredClaimProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnIdentityProviderConfig_RequiredClaimProperty struct {
	// The key to match from the token.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the key from the token.
	Value *string `field:"required" json:"value" yaml:"value"`
}

