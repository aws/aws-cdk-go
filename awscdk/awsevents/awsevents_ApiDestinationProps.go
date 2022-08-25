package awsevents


// The event API Destination properties.
//
// Example:
//   connection := events.NewConnection(this, jsii.String("Connection"), &connectionProps{
//   	authorization: events.authorization.apiKey(jsii.String("x-api-key"), awscdk.SecretValue.secretsManager(jsii.String("ApiSecretName"))),
//   	description: jsii.String("Connection with API Key x-api-key"),
//   })
//
//   destination := events.NewApiDestination(this, jsii.String("Destination"), &apiDestinationProps{
//   	connection: connection,
//   	endpoint: jsii.String("https://example.com"),
//   	description: jsii.String("Calling example.com with API key x-api-key"),
//   })
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
//   	targets: []iRuleTarget{
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
	ApiDestinationName *string `field:"optional" json:"apiDestinationName" yaml:"apiDestinationName"`
	// A description for the API destination.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The method to use for the request to the HTTP invocation endpoint.
	HttpMethod HttpMethod `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The maximum number of requests per second to send to the HTTP invocation endpoint.
	RateLimitPerSecond *float64 `field:"optional" json:"rateLimitPerSecond" yaml:"rateLimitPerSecond"`
}

