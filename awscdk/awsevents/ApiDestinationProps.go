package awsevents


// The event API Destination properties.
//
// Example:
//   connection := events.NewConnection(this, jsii.String("Connection"), &ConnectionProps{
//   	Authorization: events.Authorization_ApiKey(jsii.String("x-api-key"), awscdk.SecretValue_SecretsManager(jsii.String("ApiSecretName"))),
//   	Description: jsii.String("Connection with API Key x-api-key"),
//   })
//
//   destination := events.NewApiDestination(this, jsii.String("Destination"), &ApiDestinationProps{
//   	Connection: Connection,
//   	Endpoint: jsii.String("https://example.com"),
//   	Description: jsii.String("Calling example.com with API key x-api-key"),
//   })
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(1))),
//   	Targets: []iRuleTarget{
//   		targets.NewApiDestination(destination),
//   	},
//   })
//
type ApiDestinationProps struct {
	// The ARN of the connection to use for the API destination.
	Connection IConnection `field:"required" json:"connection" yaml:"connection"`
	// The URL to the HTTP invocation endpoint for the API destination..
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The name for the API destination.
	// Default: - A unique name will be generated.
	//
	ApiDestinationName *string `field:"optional" json:"apiDestinationName" yaml:"apiDestinationName"`
	// A description for the API destination.
	// Default: - none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The method to use for the request to the HTTP invocation endpoint.
	// Default: HttpMethod.POST
	//
	HttpMethod HttpMethod `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The maximum number of requests per second to send to the HTTP invocation endpoint.
	// Default: - Not rate limited.
	//
	RateLimitPerSecond *float64 `field:"optional" json:"rateLimitPerSecond" yaml:"rateLimitPerSecond"`
}

