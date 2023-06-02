package awsevents


// Contains the authorization parameters to use for the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authParametersProperty := &AuthParametersProperty{
//   	ApiKeyAuthParameters: &ApiKeyAuthParametersProperty{
//   		ApiKeyName: jsii.String("apiKeyName"),
//   		ApiKeyValue: jsii.String("apiKeyValue"),
//   	},
//   	BasicAuthParameters: &BasicAuthParametersProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	InvocationHttpParameters: &ConnectionHttpParametersProperty{
//   		BodyParameters: []interface{}{
//   			&ParameterProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//
//   				// the properties below are optional
//   				IsValueSecret: jsii.Boolean(false),
//   			},
//   		},
//   		HeaderParameters: []interface{}{
//   			&ParameterProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//
//   				// the properties below are optional
//   				IsValueSecret: jsii.Boolean(false),
//   			},
//   		},
//   		QueryStringParameters: []interface{}{
//   			&ParameterProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//
//   				// the properties below are optional
//   				IsValueSecret: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	OAuthParameters: &OAuthParametersProperty{
//   		AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		ClientParameters: &ClientParametersProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   		HttpMethod: jsii.String("httpMethod"),
//
//   		// the properties below are optional
//   		OAuthHttpParameters: &ConnectionHttpParametersProperty{
//   			BodyParameters: []interface{}{
//   				&ParameterProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//
//   					// the properties below are optional
//   					IsValueSecret: jsii.Boolean(false),
//   				},
//   			},
//   			HeaderParameters: []interface{}{
//   				&ParameterProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//
//   					// the properties below are optional
//   					IsValueSecret: jsii.Boolean(false),
//   				},
//   			},
//   			QueryStringParameters: []interface{}{
//   				&ParameterProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//
//   					// the properties below are optional
//   					IsValueSecret: jsii.Boolean(false),
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

