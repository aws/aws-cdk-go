package awsevents


// Properties for defining a `CfnConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectionProps := &cfnConnectionProps{
//   	authorizationType: jsii.String("authorizationType"),
//   	authParameters: &authParametersProperty{
//   		apiKeyAuthParameters: &apiKeyAuthParametersProperty{
//   			apiKeyName: jsii.String("apiKeyName"),
//   			apiKeyValue: jsii.String("apiKeyValue"),
//   		},
//   		basicAuthParameters: &basicAuthParametersProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//   		},
//   		invocationHttpParameters: &connectionHttpParametersProperty{
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
//   		oAuthParameters: &oAuthParametersProperty{
//   			authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   			clientParameters: &clientParametersProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//   			},
//   			httpMethod: jsii.String("httpMethod"),
//
//   			// the properties below are optional
//   			oAuthHttpParameters: &connectionHttpParametersProperty{
//   				bodyParameters: []interface{}{
//   					&parameterProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//
//   						// the properties below are optional
//   						isValueSecret: jsii.Boolean(false),
//   					},
//   				},
//   				headerParameters: []interface{}{
//   					&parameterProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//
//   						// the properties below are optional
//   						isValueSecret: jsii.Boolean(false),
//   					},
//   				},
//   				queryStringParameters: []interface{}{
//   					&parameterProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//
//   						// the properties below are optional
//   						isValueSecret: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   }
//
type CfnConnectionProps struct {
	// The type of authorization to use for the connection.
	//
	// > OAUTH tokens are refreshed when a 401 or 407 response is returned.
	AuthorizationType *string `field:"required" json:"authorizationType" yaml:"authorizationType"`
	// A `CreateConnectionAuthRequestParameters` object that contains the authorization parameters to use to authorize with the endpoint.
	AuthParameters interface{} `field:"required" json:"authParameters" yaml:"authParameters"`
	// A description for the connection to create.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name for the connection to create.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

