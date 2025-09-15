package awsapigatewayv2


// ApiKey Properties.
//
// Example:
//   // Auto-generated name and value
//   autoKey := apigwv2.NewApiKey(this, jsii.String("AutoKey"))
//
//   // Explicit name and value
//   explicitKey := apigwv2.NewApiKey(this, jsii.String("ExplicitKey"), &ApiKeyProps{
//   	ApiKeyName: jsii.String("MyWebSocketApiKey"),
//   	Value: jsii.String("MyApiKeyThatIsAtLeast20Characters"),
//   })
//
type ApiKeyProps struct {
	// A name for the API key.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.
	// Default: automatically generated name.
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
}

