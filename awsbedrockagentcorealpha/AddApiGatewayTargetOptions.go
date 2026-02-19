package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
)

// Options for adding an API Gateway target to a gateway.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("MyApi"), &RestApiProps{
//   	RestApiName: jsii.String("my-api"),
//   })
//
//   // Uses IAM authorization for outbound auth by default
//   apiGatewayTarget := gateway.AddApiGatewayTarget(jsii.String("MyApiGatewayTarget"), &AddApiGatewayTargetOptions{
//   	RestApi: api,
//   	ApiGatewayToolConfiguration: &ApiGatewayToolConfiguration{
//   		ToolFilters: []ApiGatewayToolFilter{
//   			&ApiGatewayToolFilter{
//   				FilterPath: jsii.String("/pets/*"),
//   				Methods: []ApiGatewayHttpMethod{
//   					agentcore.ApiGatewayHttpMethod_GET,
//   				},
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type AddApiGatewayTargetOptions struct {
	// Tool configuration defining which operations to expose.
	// Experimental.
	ApiGatewayToolConfiguration *ApiGatewayToolConfiguration `field:"required" json:"apiGatewayToolConfiguration" yaml:"apiGatewayToolConfiguration"`
	// The REST API to integrate with Must be in the same account and region as the gateway [disable-awslint:prefer-ref-interface].
	// Experimental.
	RestApi awsapigateway.IRestApi `field:"required" json:"restApi" yaml:"restApi"`
	// Credential providers for authentication API Gateway targets support IAM and API key authentication.
	// Default: - Empty array (service handles IAM automatically).
	//
	// Experimental.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// Optional description for the gateway target.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target Valid characters are a-z, A-Z, 0-9, _ (underscore) and - (hyphen).
	// Default: - auto generate.
	//
	// Experimental.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// Metadata configuration for passing headers and query parameters Allows you to pass additional context through headers and query parameters.
	// Default: - No metadata configuration.
	//
	// Experimental.
	MetadataConfiguration *MetadataConfiguration `field:"optional" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// The stage name of the REST API.
	// Default: - Uses the deployment stage from the RestApi (restApi.deploymentStage.stageName)
	//
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

