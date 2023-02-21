package awsevents


// Contains the authorization parameters to use for the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authParametersProperty := &authParametersProperty{
//   	apiKeyAuthParameters: &apiKeyAuthParametersProperty{
//   		apiKeyName: jsii.String("apiKeyName"),
//   		apiKeyValue: jsii.String("apiKeyValue"),
//   	},
//   	basicAuthParameters: &basicAuthParametersProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	invocationHttpParameters: &connectionHttpParametersProperty{
//   		bodyParameters: []interface{}{
//   			&parameterProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//
//   				// the properties below are optional
//   				isValueSecret: jsii.Boolean(false),
//   			},
//   		},
//   		headerParameters: []interface{}{
//   			&parameterProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//
//   				// the properties below are optional
//   				isValueSecret: jsii.Boolean(false),
//   			},
//   		},
//   		queryStringParameters: []interface{}{
//   			&parameterProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//
//   				// the properties below are optional
//   				isValueSecret: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	oAuthParameters: &oAuthParametersProperty{
//   		authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		clientParameters: &clientParametersProperty{
//   			clientId: jsii.String("clientId"),
//   			clientSecret: jsii.String("clientSecret"),
//   		},
//   		httpMethod: jsii.String("httpMethod"),
//
//   		// the properties below are optional
//   		oAuthHttpParameters: &connectionHttpParametersProperty{
//   			bodyParameters: []interface{}{
//   				&parameterProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//
//   					// the properties below are optional
//   					isValueSecret: jsii.Boolean(false),
//   				},
//   			},
//   			headerParameters: []interface{}{
//   				&parameterProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//
//   					// the properties below are optional
//   					isValueSecret: jsii.Boolean(false),
//   				},
//   			},
//   			queryStringParameters: []interface{}{
//   				&parameterProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//
//   					// the properties below are optional
//   					isValueSecret: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnConnection_AuthParametersProperty struct {
	// The API Key parameters to use for authorization.
	ApiKeyAuthParameters interface{} `field:"optional" json:"apiKeyAuthParameters" yaml:"apiKeyAuthParameters"`
	// The authorization parameters for Basic authorization.
	BasicAuthParameters interface{} `field:"optional" json:"basicAuthParameters" yaml:"basicAuthParameters"`
	// Additional parameters for the connection that are passed through with every invocation to the HTTP endpoint.
	InvocationHttpParameters interface{} `field:"optional" json:"invocationHttpParameters" yaml:"invocationHttpParameters"`
	// The OAuth parameters to use for authorization.
	OAuthParameters interface{} `field:"optional" json:"oAuthParameters" yaml:"oAuthParameters"`
}

