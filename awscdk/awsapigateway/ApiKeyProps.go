package awsapigateway


// ApiKey Properties.
//
// Example:
//   stack := awscdk.NewStack(app, jsii.String("my-stack"), &StackProps{
//   	PropertyInjectors: []iPropertyInjector{
//   		NewApiKeyPropsInjector(),
//   	},
//   })
//   api.NewApiKey(stack, jsii.String("my-api-key"), &apiKeyProps{
//   })
//
type ApiKeyProps struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	// Default: - CORS is disabled.
	//
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	// Default: - Inherited from parent.
	//
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	// Default: - Inherited from parent.
	//
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// A name for the API key.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.
	// Default: automically generated name.
	//
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// A description of the purpose of the API key.
	// Default: none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The value of the API key.
	//
	// Must be at least 20 characters long.
	// Default: none.
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
	// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
	// Default: none.
	//
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
	// Indicates whether the API key can be used by clients.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies whether the key identifier is distinct from the created API key value.
	// Default: false.
	//
	GenerateDistinctId *bool `field:"optional" json:"generateDistinctId" yaml:"generateDistinctId"`
	// A list of resources this api key is associated with.
	// Default: none.
	//
	// Deprecated: - use `stages` instead.
	Resources *[]IRestApi `field:"optional" json:"resources" yaml:"resources"`
	// A list of Stages this api key is associated with.
	// Default: - the api key is not associated with any stages.
	//
	Stages *[]IStage `field:"optional" json:"stages" yaml:"stages"`
}

