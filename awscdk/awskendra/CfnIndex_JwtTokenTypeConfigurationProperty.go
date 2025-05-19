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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html
//
type CfnIndex_JwtTokenTypeConfigurationProperty struct {
	// The location of the key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-keylocation
	//
	KeyLocation *string `field:"required" json:"keyLocation" yaml:"keyLocation"`
	// The regular expression that identifies the claim.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-claimregex
	//
	ClaimRegex *string `field:"optional" json:"claimRegex" yaml:"claimRegex"`
	// The group attribute field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-groupattributefield
	//
	GroupAttributeField *string `field:"optional" json:"groupAttributeField" yaml:"groupAttributeField"`
	// The issuer of the token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-issuer
	//
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The Amazon Resource Name (arn) of the secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-secretmanagerarn
	//
	SecretManagerArn *string `field:"optional" json:"secretManagerArn" yaml:"secretManagerArn"`
	// The signing key URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The user name attribute field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-usernameattributefield
	//
	UserNameAttributeField *string `field:"optional" json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

