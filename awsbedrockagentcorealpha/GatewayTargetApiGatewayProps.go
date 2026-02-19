package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
)

// Properties for creating an API Gateway-based Gateway Target.
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
//   // Create a gateway target using the static factory method
//   apiGatewayTarget := agentcore.GatewayTarget_ForApiGateway(this, jsii.String("MyApiGatewayTarget"), &GatewayTargetApiGatewayProps{
//   	GatewayTargetName: jsii.String("my-api-gateway-target"),
//   	Description: jsii.String("Target for API Gateway REST API integration"),
//   	Gateway: gateway,
//   	RestApi: api,
//   	ApiGatewayToolConfiguration: &ApiGatewayToolConfiguration{
//   		ToolFilters: []ApiGatewayToolFilter{
//   			&ApiGatewayToolFilter{
//   				FilterPath: jsii.String("/pets/*"),
//   				Methods: []ApiGatewayHttpMethod{
//   					agentcore.ApiGatewayHttpMethod_GET,
//   					agentcore.ApiGatewayHttpMethod_POST,
//   				},
//   			},
//   		},
//   	},
//   	MetadataConfiguration: &MetadataConfiguration{
//   		AllowedRequestHeaders: []*string{
//   			jsii.String("X-User-Id"),
//   		},
//   		AllowedQueryParameters: []*string{
//   			jsii.String("limit"),
//   		},
//   	},
//   })
//
// Experimental.
type GatewayTargetApiGatewayProps struct {
	// Optional description for the gateway target The description can have up to 200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target The name must be unique within the gateway Pattern: ^([0-9a-zA-Z][-]?){1,100}$.
	// Default: - auto generate.
	//
	// Experimental.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// Tool configuration defining which operations to expose.
	// Experimental.
	ApiGatewayToolConfiguration *ApiGatewayToolConfiguration `field:"required" json:"apiGatewayToolConfiguration" yaml:"apiGatewayToolConfiguration"`
	// The gateway this target belongs to [disable-awslint:prefer-ref-interface].
	// Experimental.
	Gateway IGateway `field:"required" json:"gateway" yaml:"gateway"`
	// The REST API to integrate with Must be in the same account and region as the gateway [disable-awslint:prefer-ref-interface].
	// Experimental.
	RestApi awsapigateway.IRestApi `field:"required" json:"restApi" yaml:"restApi"`
	// Credential providers for authentication API Gateway targets support IAM and API key authentication.
	// Default: - Empty array (service handles IAM automatically).
	//
	// Experimental.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
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

