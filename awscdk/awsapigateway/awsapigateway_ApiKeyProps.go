package awsapigateway


// ApiKey Properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizer authorizer
//   var integration integration
//   var model model
//   var requestValidator requestValidator
//   var restApi restApi
//   var stage stage
//
//   apiKeyProps := &apiKeyProps{
//   	apiKeyName: jsii.String("apiKeyName"),
//   	customerId: jsii.String("customerId"),
//   	defaultCorsPreflightOptions: &corsOptions{
//   		allowOrigins: []*string{
//   			jsii.String("allowOrigins"),
//   		},
//
//   		// the properties below are optional
//   		allowCredentials: jsii.Boolean(false),
//   		allowHeaders: []*string{
//   			jsii.String("allowHeaders"),
//   		},
//   		allowMethods: []*string{
//   			jsii.String("allowMethods"),
//   		},
//   		disableCache: jsii.Boolean(false),
//   		exposeHeaders: []*string{
//   			jsii.String("exposeHeaders"),
//   		},
//   		maxAge: cdk.duration.minutes(jsii.Number(30)),
//   		statusCode: jsii.Number(123),
//   	},
//   	defaultIntegration: integration,
//   	defaultMethodOptions: &methodOptions{
//   		apiKeyRequired: jsii.Boolean(false),
//   		authorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		authorizationType: awscdk.Aws_apigateway.authorizationType_NONE,
//   		authorizer: authorizer,
//   		methodResponses: []methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				responseModels: map[string]iModel{
//   					"responseModelsKey": model,
//   				},
//   				responseParameters: map[string]*bool{
//   					"responseParametersKey": jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		operationName: jsii.String("operationName"),
//   		requestModels: map[string]*iModel{
//   			"requestModelsKey": model,
//   		},
//   		requestParameters: map[string]*bool{
//   			"requestParametersKey": jsii.Boolean(false),
//   		},
//   		requestValidator: requestValidator,
//   		requestValidatorOptions: &requestValidatorOptions{
//   			requestValidatorName: jsii.String("requestValidatorName"),
//   			validateRequestBody: jsii.Boolean(false),
//   			validateRequestParameters: jsii.Boolean(false),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	enabled: jsii.Boolean(false),
//   	generateDistinctId: jsii.Boolean(false),
//   	resources: []iRestApi{
//   		restApi,
//   	},
//   	stages: []iStage{
//   		stage,
//   	},
//   	value: jsii.String("value"),
//   }
//
type ApiKeyProps struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// A name for the API key.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// A description of the purpose of the API key.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The value of the API key.
	//
	// Must be at least 20 characters long.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
	// Indicates whether the API key can be used by clients.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies whether the key identifier is distinct from the created API key value.
	GenerateDistinctId *bool `field:"optional" json:"generateDistinctId" yaml:"generateDistinctId"`
	// A list of resources this api key is associated with.
	// Deprecated: - use `stages` instead.
	Resources *[]IRestApi `field:"optional" json:"resources" yaml:"resources"`
	// A list of Stages this api key is associated with.
	Stages *[]IStage `field:"optional" json:"stages" yaml:"stages"`
}

