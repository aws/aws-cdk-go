package awsbedrockagentcore


// The API key credential provider for the gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyCredentialProviderProperty := &ApiKeyCredentialProviderProperty{
//   	ProviderArn: jsii.String("providerArn"),
//
//   	// the properties below are optional
//   	CredentialLocation: jsii.String("credentialLocation"),
//   	CredentialParameterName: jsii.String("credentialParameterName"),
//   	CredentialPrefix: jsii.String("credentialPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apikeycredentialprovider.html
//
type CfnGatewayTarget_ApiKeyCredentialProviderProperty struct {
	// The provider ARN for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apikeycredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-apikeycredentialprovider-providerarn
	//
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// The credential location for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apikeycredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-apikeycredentialprovider-credentiallocation
	//
	CredentialLocation *string `field:"optional" json:"credentialLocation" yaml:"credentialLocation"`
	// The credential parameter name for the provider for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apikeycredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-apikeycredentialprovider-credentialparametername
	//
	CredentialParameterName *string `field:"optional" json:"credentialParameterName" yaml:"credentialParameterName"`
	// The API key credential provider for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apikeycredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-apikeycredentialprovider-credentialprefix
	//
	CredentialPrefix *string `field:"optional" json:"credentialPrefix" yaml:"credentialPrefix"`
}

