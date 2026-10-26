package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   idpConfigProperty := &IdpConfigProperty{
//   	Audience: jsii.String("audience"),
//   	CredentialProviderArn: jsii.String("credentialProviderArn"),
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-consentportal-idpconfig.html
//
type CfnConsentPortalPropsMixin_IdpConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-consentportal-idpconfig.html#cfn-bedrockagentcore-consentportal-idpconfig-audience
	//
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// The ARN of an AWS::BedrockAgentCore::OAuth2CredentialProvider resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-consentportal-idpconfig.html#cfn-bedrockagentcore-consentportal-idpconfig-credentialproviderarn
	//
	CredentialProviderArn *string `field:"optional" json:"credentialProviderArn" yaml:"credentialProviderArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-consentportal-idpconfig.html#cfn-bedrockagentcore-consentportal-idpconfig-scopes
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

