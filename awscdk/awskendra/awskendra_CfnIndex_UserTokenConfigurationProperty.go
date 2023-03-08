package awskendra


// Provides the configuration information for a token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userTokenConfigurationProperty := &userTokenConfigurationProperty{
//   	jsonTokenTypeConfiguration: &jsonTokenTypeConfigurationProperty{
//   		groupAttributeField: jsii.String("groupAttributeField"),
//   		userNameAttributeField: jsii.String("userNameAttributeField"),
//   	},
//   	jwtTokenTypeConfiguration: &jwtTokenTypeConfigurationProperty{
//   		keyLocation: jsii.String("keyLocation"),
//
//   		// the properties below are optional
//   		claimRegex: jsii.String("claimRegex"),
//   		groupAttributeField: jsii.String("groupAttributeField"),
//   		issuer: jsii.String("issuer"),
//   		secretManagerArn: jsii.String("secretManagerArn"),
//   		url: jsii.String("url"),
//   		userNameAttributeField: jsii.String("userNameAttributeField"),
//   	},
//   }
//
type CfnIndex_UserTokenConfigurationProperty struct {
	// Information about the JSON token type configuration.
	JsonTokenTypeConfiguration interface{} `field:"optional" json:"jsonTokenTypeConfiguration" yaml:"jsonTokenTypeConfiguration"`
	// Information about the JWT token type configuration.
	JwtTokenTypeConfiguration interface{} `field:"optional" json:"jwtTokenTypeConfiguration" yaml:"jwtTokenTypeConfiguration"`
}

