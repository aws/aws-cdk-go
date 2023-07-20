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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-authparameters.html
//
type CfnConnection_AuthParametersProperty struct {
	// The API Key parameters to use for authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-authparameters.html#cfn-events-connection-authparameters-apikeyauthparameters
	//
	ApiKeyAuthParameters interface{} `field:"optional" json:"apiKeyAuthParameters" yaml:"apiKeyAuthParameters"`
	// The authorization parameters for Basic authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-authparameters.html#cfn-events-connection-authparameters-basicauthparameters
	//
	BasicAuthParameters interface{} `field:"optional" json:"basicAuthParameters" yaml:"basicAuthParameters"`
	// Additional parameters for the connection that are passed through with every invocation to the HTTP endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-authparameters.html#cfn-events-connection-authparameters-invocationhttpparameters
	//
	InvocationHttpParameters interface{} `field:"optional" json:"invocationHttpParameters" yaml:"invocationHttpParameters"`
	// The OAuth parameters to use for authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-authparameters.html#cfn-events-connection-authparameters-oauthparameters
	//
	OAuthParameters interface{} `field:"optional" json:"oAuthParameters" yaml:"oAuthParameters"`
}

