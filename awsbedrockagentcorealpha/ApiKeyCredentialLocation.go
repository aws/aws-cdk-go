package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// API Key location within the request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   apiKeyCredentialLocation := bedrock_agentcore_alpha.ApiKeyCredentialLocation_Header(&ApiKeyAdditionalConfiguration{
//   	CredentialParameterName: jsii.String("credentialParameterName"),
//   	CredentialPrefix: jsii.String("credentialPrefix"),
//   })
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type ApiKeyCredentialLocation interface {
	// The type of credential location (HEADER or QUERY_PARAMETER).
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialLocationType() *string
	// The name of the credential parameter.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialParameterName() *string
	// The prefix for the credential value.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialPrefix() *string
}

// The jsii proxy struct for ApiKeyCredentialLocation
type jsiiProxy_ApiKeyCredentialLocation struct {
	_ byte // padding
}

func (j *jsiiProxy_ApiKeyCredentialLocation) CredentialLocationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialLocationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiKeyCredentialLocation) CredentialParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialParameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiKeyCredentialLocation) CredentialPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialPrefix",
		&returns,
	)
	return returns
}


// Create a header-based API key credential location.
//
// Returns: ApiKeyCredentialLocation configured for header placement.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func ApiKeyCredentialLocation_Header(config *ApiKeyAdditionalConfiguration) ApiKeyCredentialLocation {
	_init_.Initialize()

	if err := validateApiKeyCredentialLocation_HeaderParameters(config); err != nil {
		panic(err)
	}
	var returns ApiKeyCredentialLocation

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiKeyCredentialLocation",
		"header",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Create a query parameter-based API key credential location.
//
// Returns: ApiKeyCredentialLocation configured for query parameter placement.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func ApiKeyCredentialLocation_QueryParameter(config *ApiKeyAdditionalConfiguration) ApiKeyCredentialLocation {
	_init_.Initialize()

	if err := validateApiKeyCredentialLocation_QueryParameterParameters(config); err != nil {
		panic(err)
	}
	var returns ApiKeyCredentialLocation

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiKeyCredentialLocation",
		"queryParameter",
		[]interface{}{config},
		&returns,
	)

	return returns
}

