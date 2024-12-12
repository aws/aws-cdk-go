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
//   	AuthParameters: &AuthParametersProperty{
//   		ApiKeyAuthParameters: &ApiKeyAuthParametersProperty{
//   			ApiKeyName: jsii.String("apiKeyName"),
//   			ApiKeyValue: jsii.String("apiKeyValue"),
//   		},
//   		BasicAuthParameters: &BasicAuthParametersProperty{
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//   		},
//   		ConnectivityParameters: &ConnectivityParametersProperty{
//   			ResourceParameters: &ResourceParametersProperty{
//   				ResourceConfigurationArn: jsii.String("resourceConfigurationArn"),
//
//   				// the properties below are optional
//   				ResourceAssociationArn: jsii.String("resourceAssociationArn"),
//   			},
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
//   	InvocationConnectivityParameters: &InvocationConnectivityParametersProperty{
//   		ResourceParameters: &ResourceParametersProperty{
//   			ResourceConfigurationArn: jsii.String("resourceConfigurationArn"),
//
//   			// the properties below are optional
//   			ResourceAssociationArn: jsii.String("resourceAssociationArn"),
//   		},
//   	},
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
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// The authorization parameters to use to authorize with the endpoint.
	//
	// You must include only authorization parameters for the `AuthorizationType` you specify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html#cfn-events-connection-authparameters
	//
	AuthParameters interface{} `field:"optional" json:"authParameters" yaml:"authParameters"`
	// A description for the connection to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html#cfn-events-connection-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The private resource the HTTP request will be sent to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html#cfn-events-connection-invocationconnectivityparameters
	//
	InvocationConnectivityParameters interface{} `field:"optional" json:"invocationConnectivityParameters" yaml:"invocationConnectivityParameters"`
	// The name for the connection to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html#cfn-events-connection-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

