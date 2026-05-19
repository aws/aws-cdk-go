package awsbedrockagentcorealpha


// API Key additional configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   apiKeyAdditionalConfiguration := &ApiKeyAdditionalConfiguration{
//   	CredentialParameterName: jsii.String("credentialParameterName"),
//   	CredentialPrefix: jsii.String("credentialPrefix"),
//   }
//
// Experimental.
type ApiKeyAdditionalConfiguration struct {
	// The name of the credential parameter for the API key.
	//
	// This parameter name is used when sending the API key to the target endpoint.
	//
	// Length Constraints: Minimum length of 1. Maximum length of 64.
	// Default: - 'Authorization' for HEADER, 'api_key' for QUERY_PARAMETER.
	//
	// Experimental.
	CredentialParameterName *string `field:"optional" json:"credentialParameterName" yaml:"credentialParameterName"`
	// The prefix for the API key credential.
	//
	// This prefix is added to the API key when sending it to the target endpoint.
	//
	// Length Constraints: Minimum length of 1. Maximum length of 64.
	// Default: - 'Bearer ' for HEADER, no prefix for QUERY_PARAMETER.
	//
	// Experimental.
	CredentialPrefix *string `field:"optional" json:"credentialPrefix" yaml:"credentialPrefix"`
}

