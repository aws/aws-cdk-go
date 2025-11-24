package mixinsawskendra


// Provides the configuration information for a token.
//
// > If you're using an Amazon Kendra Gen AI Enterprise Edition index and you try to use `UserTokenConfigurations` to configure user context policy, Amazon Kendra returns a `ValidationException` error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   userTokenConfigurationProperty := &UserTokenConfigurationProperty{
//   	JsonTokenTypeConfiguration: &JsonTokenTypeConfigurationProperty{
//   		GroupAttributeField: jsii.String("groupAttributeField"),
//   		UserNameAttributeField: jsii.String("userNameAttributeField"),
//   	},
//   	JwtTokenTypeConfiguration: &JwtTokenTypeConfigurationProperty{
//   		ClaimRegex: jsii.String("claimRegex"),
//   		GroupAttributeField: jsii.String("groupAttributeField"),
//   		Issuer: jsii.String("issuer"),
//   		KeyLocation: jsii.String("keyLocation"),
//   		SecretManagerArn: jsii.String("secretManagerArn"),
//   		Url: jsii.String("url"),
//   		UserNameAttributeField: jsii.String("userNameAttributeField"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-usertokenconfiguration.html
//
type CfnIndexPropsMixin_UserTokenConfigurationProperty struct {
	// Information about the JSON token type configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-usertokenconfiguration.html#cfn-kendra-index-usertokenconfiguration-jsontokentypeconfiguration
	//
	JsonTokenTypeConfiguration interface{} `field:"optional" json:"jsonTokenTypeConfiguration" yaml:"jsonTokenTypeConfiguration"`
	// Information about the JWT token type configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-usertokenconfiguration.html#cfn-kendra-index-usertokenconfiguration-jwttokentypeconfiguration
	//
	JwtTokenTypeConfiguration interface{} `field:"optional" json:"jwtTokenTypeConfiguration" yaml:"jwtTokenTypeConfiguration"`
}

