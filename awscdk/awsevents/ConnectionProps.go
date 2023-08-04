package awsevents


// An API Destination Connection.
//
// A connection defines the authorization type and credentials to use for authorization with an API destination HTTP endpoint.
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
type ConnectionProps struct {
	// The authorization type for the connection.
	Authorization Authorization `field:"required" json:"authorization" yaml:"authorization"`
	// Additional string parameters to add to the invocation bodies.
	// Default: - No additional parameters.
	//
	BodyParameters *map[string]HttpParameter `field:"optional" json:"bodyParameters" yaml:"bodyParameters"`
	// The name of the connection.
	// Default: - A name is automatically generated.
	//
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The name of the connection.
	// Default: - none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Additional string parameters to add to the invocation headers.
	// Default: - No additional parameters.
	//
	HeaderParameters *map[string]HttpParameter `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// Additional string parameters to add to the invocation query strings.
	// Default: - No additional parameters.
	//
	QueryStringParameters *map[string]HttpParameter `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

