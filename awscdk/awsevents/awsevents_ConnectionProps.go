package awsevents


// An API Destination Connection.
//
// A connection defines the authorization type and credentials to use for authorization with an API destination HTTP endpoint.
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
// Experimental.
type ConnectionProps struct {
	// The authorization type for the connection.
	// Experimental.
	Authorization Authorization `field:"required" json:"authorization" yaml:"authorization"`
	// Additional string parameters to add to the invocation bodies.
	// Experimental.
	BodyParameters *map[string]HttpParameter `field:"optional" json:"bodyParameters" yaml:"bodyParameters"`
	// The name of the connection.
	// Experimental.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The name of the connection.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Additional string parameters to add to the invocation headers.
	// Experimental.
	HeaderParameters *map[string]HttpParameter `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// Additional string parameters to add to the invocation query strings.
	// Experimental.
	QueryStringParameters *map[string]HttpParameter `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

