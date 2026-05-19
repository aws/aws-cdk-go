package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// API Key location within the request.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   // Create an API key credential provider in Token Vault
//   apiKeyProvider := agentcore.NewApiKeyCredentialProvider(this, jsii.String("MyApiKeyProvider"), &ApiKeyCredentialProviderProps{
//   	ApiKeyCredentialProviderName: jsii.String("my-apikey"),
//   })
//
//   bucket := s3.Bucket_FromBucketName(this, jsii.String("ExistingBucket"), jsii.String("my-schema-bucket"))
//   s3mySchema := agentcore.ApiSchema_FromS3File(bucket, jsii.String("schemas/myschema.yaml"))
//
//   // Add an OpenAPI target using the L2 construct directly
//   target := gateway.AddOpenApiTarget(jsii.String("MyTarget"), &AddOpenApiTargetOptions{
//   	GatewayTargetName: jsii.String("my-api-target"),
//   	Description: jsii.String("Target for external API integration"),
//   	ApiSchema: s3mySchema,
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		agentcore.GatewayCredentialProvider_FromApiKeyIdentity(apiKeyProvider, &FromApiKeyIdentityOptions{
//   			CredentialLocation: agentcore.ApiKeyCredentialLocation_Header(&ApiKeyAdditionalConfiguration{
//   				CredentialParameterName: jsii.String("X-API-Key"),
//   			}),
//   		}),
//   	},
//   })
//
//   // This makes sure your s3 bucket is available before target
//   target.Node.AddDependency(bucket)
//
type ApiKeyCredentialLocation interface {
	// The type of credential location (HEADER or QUERY_PARAMETER).
	CredentialLocationType() *string
	// The name of the credential parameter.
	CredentialParameterName() *string
	// The prefix for the credential value.
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
func ApiKeyCredentialLocation_Header(config *ApiKeyAdditionalConfiguration) ApiKeyCredentialLocation {
	_init_.Initialize()

	if err := validateApiKeyCredentialLocation_HeaderParameters(config); err != nil {
		panic(err)
	}
	var returns ApiKeyCredentialLocation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialLocation",
		"header",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Create a query parameter-based API key credential location.
//
// Returns: ApiKeyCredentialLocation configured for query parameter placement.
func ApiKeyCredentialLocation_QueryParameter(config *ApiKeyAdditionalConfiguration) ApiKeyCredentialLocation {
	_init_.Initialize()

	if err := validateApiKeyCredentialLocation_QueryParameterParameters(config); err != nil {
		panic(err)
	}
	var returns ApiKeyCredentialLocation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.ApiKeyCredentialLocation",
		"queryParameter",
		[]interface{}{config},
		&returns,
	)

	return returns
}

