package previewawseksmixins


// A key-value pair that describes a required claim in the identity token.
//
// If set, each claim is verified to be present in the token with a matching value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requiredClaimProperty := &RequiredClaimProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-identityproviderconfig-requiredclaim.html
//
type CfnIdentityProviderConfigPropsMixin_RequiredClaimProperty struct {
	// The key to match from the token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-identityproviderconfig-requiredclaim.html#cfn-eks-identityproviderconfig-requiredclaim-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the key from the token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-identityproviderconfig-requiredclaim.html#cfn-eks-identityproviderconfig-requiredclaim-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

