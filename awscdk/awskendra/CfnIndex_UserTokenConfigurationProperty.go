package awskendra


// Provides the configuration information for a token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userTokenConfigurationProperty := &UserTokenConfigurationProperty{
//   	JsonTokenTypeConfiguration: &JsonTokenTypeConfigurationProperty{
//   		GroupAttributeField: jsii.String("groupAttributeField"),
//   		UserNameAttributeField: jsii.String("userNameAttributeField"),
//   	},
//   	JwtTokenTypeConfiguration: &JwtTokenTypeConfigurationProperty{
//   		KeyLocation: jsii.String("keyLocation"),
//
//   		// the properties below are optional
//   		ClaimRegex: jsii.String("claimRegex"),
//   		GroupAttributeField: jsii.String("groupAttributeField"),
//   		Issuer: jsii.String("issuer"),
//   		SecretManagerArn: jsii.String("secretManagerArn"),
//   		Url: jsii.String("url"),
//   		UserNameAttributeField: jsii.String("userNameAttributeField"),
//   	},
//   }
//
type CfnIndex_UserTokenConfigurationProperty struct {
	// Information about the JSON token type configuration.
	JsonTokenTypeConfiguration interface{} `field:"optional" json:"jsonTokenTypeConfiguration" yaml:"jsonTokenTypeConfiguration"`
	// Information about the JWT token type configuration.
	JwtTokenTypeConfiguration interface{} `field:"optional" json:"jwtTokenTypeConfiguration" yaml:"jwtTokenTypeConfiguration"`
}

