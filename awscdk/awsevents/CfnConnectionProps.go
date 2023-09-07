package awsevents


// Properties for defining a `CfnConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectionProps := &CfnConnectionProps{
//   	AuthorizationType: jsii.String("authorizationType"),
//
//   	// the properties below are optional
//   	AuthParameters: &AuthParametersProperty{
//   		ApiKeyAuthParameters: &ApiKeyAuthParametersProperty{
//   			ApiKeyName: jsii.String("apiKeyName"),
//   			ApiKeyValue: jsii.String("apiKeyValue"),
//   		},
//   		BasicAuthParameters: &BasicAuthParametersProperty{
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//   		},
//   		InvocationHttpParameters: &ConnectionHttpParametersProperty{
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
//   		OAuthParameters: &OAuthParametersProperty{
//   			AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   			ClientParameters: &ClientParametersProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//   			},
//   			HttpMethod: jsii.String("httpMethod"),
//
//   			// the properties below are optional
//   			OAuthHttpParameters: &ConnectionHttpParametersProperty{
//   				BodyParameters: []interface{}{
//   					&ParameterProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//
//   						// the properties below are optional
//   						IsValueSecret: jsii.Boolean(false),
//   					},
//   				},
//   				HeaderParameters: []interface{}{
//   					&ParameterProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//
//   						// the properties below are optional
//   						IsValueSecret: jsii.Boolean(false),
//   					},
//   				},
//   				QueryStringParameters: []interface{}{
//   					&ParameterProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//
//   						// the properties below are optional
//   						IsValueSecret: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html
//
type CfnConnectionProps struct {
	// The type of authorization to use for the connection.
	//
	// > OAUTH tokens are refreshed when a 401 or 407 response is returned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html#cfn-events-connection-authorizationtype
	//
	AuthorizationType *string `field:"required" json:"authorizationType" yaml:"authorizationType"`
	// A `CreateConnectionAuthRequestParameters` object that contains the authorization parameters to use to authorize with the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html#cfn-events-connection-authparameters
	//
	AuthParameters interface{} `field:"optional" json:"authParameters" yaml:"authParameters"`
	// A description for the connection to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html#cfn-events-connection-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name for the connection to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html#cfn-events-connection-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

