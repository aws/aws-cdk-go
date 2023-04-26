package awskendra


// Provides the configuration information for the JWT token type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jwtTokenTypeConfigurationProperty := &JwtTokenTypeConfigurationProperty{
//   	KeyLocation: jsii.String("keyLocation"),
//
//   	// the properties below are optional
//   	ClaimRegex: jsii.String("claimRegex"),
//   	GroupAttributeField: jsii.String("groupAttributeField"),
//   	Issuer: jsii.String("issuer"),
//   	SecretManagerArn: jsii.String("secretManagerArn"),
//   	Url: jsii.String("url"),
//   	UserNameAttributeField: jsii.String("userNameAttributeField"),
//   }
//
type CfnIndex_JwtTokenTypeConfigurationProperty struct {
	// The location of the key.
	KeyLocation *string `field:"required" json:"keyLocation" yaml:"keyLocation"`
	// The regular expression that identifies the claim.
	ClaimRegex *string `field:"optional" json:"claimRegex" yaml:"claimRegex"`
	// The group attribute field.
	GroupAttributeField *string `field:"optional" json:"groupAttributeField" yaml:"groupAttributeField"`
	// The issuer of the token.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The Amazon Resource Name (arn) of the secret.
	SecretManagerArn *string `field:"optional" json:"secretManagerArn" yaml:"secretManagerArn"`
	// The signing key URL.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The user name attribute field.
	UserNameAttributeField *string `field:"optional" json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

