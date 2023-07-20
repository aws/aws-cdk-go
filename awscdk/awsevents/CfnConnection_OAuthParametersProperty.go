package awsevents


// Contains the OAuth authorization parameters to use for the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuthParametersProperty := &OAuthParametersProperty{
//   	AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	ClientParameters: &ClientParametersProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   	},
//   	HttpMethod: jsii.String("httpMethod"),
//
//   	// the properties below are optional
//   	OAuthHttpParameters: &ConnectionHttpParametersProperty{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-oauthparameters.html
//
type CfnConnection_OAuthParametersProperty struct {
	// The URL to the authorization endpoint when OAuth is specified as the authorization type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-oauthparameters.html#cfn-events-connection-oauthparameters-authorizationendpoint
	//
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// A `CreateConnectionOAuthClientRequestParameters` object that contains the client parameters for OAuth authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-oauthparameters.html#cfn-events-connection-oauthparameters-clientparameters
	//
	ClientParameters interface{} `field:"required" json:"clientParameters" yaml:"clientParameters"`
	// The method to use for the authorization request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-oauthparameters.html#cfn-events-connection-oauthparameters-httpmethod
	//
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// A `ConnectionHttpParameters` object that contains details about the additional parameters to use for the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-oauthparameters.html#cfn-events-connection-oauthparameters-oauthhttpparameters
	//
	OAuthHttpParameters interface{} `field:"optional" json:"oAuthHttpParameters" yaml:"oAuthHttpParameters"`
}

