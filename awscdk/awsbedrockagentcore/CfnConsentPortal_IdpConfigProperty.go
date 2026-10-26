package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idpConfigProperty := &IdpConfigProperty{
//   	CredentialProviderArn: jsii.String("credentialProviderArn"),
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//
//   	// the properties below are optional
//   	Audience: jsii.String("audience"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-consentportal-idpconfig.html
//
type CfnConsentPortal_IdpConfigProperty struct {
	// The ARN of an AWS::BedrockAgentCore::OAuth2CredentialProvider resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-consentportal-idpconfig.html#cfn-bedrockagentcore-consentportal-idpconfig-credentialproviderarn
	//
	CredentialProviderArn *string `field:"required" json:"credentialProviderArn" yaml:"credentialProviderArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-consentportal-idpconfig.html#cfn-bedrockagentcore-consentportal-idpconfig-scopes
	//
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-consentportal-idpconfig.html#cfn-bedrockagentcore-consentportal-idpconfig-audience
	//
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
}

