package awsevents


// Contains the OAuth authorization parameters to use for the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuthParametersProperty := &oAuthParametersProperty{
//   	authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	clientParameters: &clientParametersProperty{
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//   	},
//   	httpMethod: jsii.String("httpMethod"),
//
//   	// the properties below are optional
//   	oAuthHttpParameters: &connectionHttpParametersProperty{
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
//   }
//
type CfnConnection_OAuthParametersProperty struct {
	// The URL to the authorization endpoint when OAuth is specified as the authorization type.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// A `CreateConnectionOAuthClientRequestParameters` object that contains the client parameters for OAuth authorization.
	ClientParameters interface{} `field:"required" json:"clientParameters" yaml:"clientParameters"`
	// The method to use for the authorization request.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// A `ConnectionHttpParameters` object that contains details about the additional parameters to use for the connection.
	OAuthHttpParameters interface{} `field:"optional" json:"oAuthHttpParameters" yaml:"oAuthHttpParameters"`
}

