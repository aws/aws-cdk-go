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
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
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
	// For connections to private APIs, the parameters to use for invoking the API.
	//
	// For more information, see [Connecting to private APIs](https://docs.aws.amazon.com/eventbridge/latest/userguide/connection-private.html) in the **Amazon EventBridge User Guide** .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html#cfn-events-connection-invocationconnectivityparameters
	//
	InvocationConnectivityParameters interface{} `field:"optional" json:"invocationConnectivityParameters" yaml:"invocationConnectivityParameters"`
	// The identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this connection.
	//
	// The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.
	//
	// If you do not specify a customer managed key identifier, EventBridge uses an AWS owned key to encrypt the connection.
	//
	// For more information, see [Identify and view keys](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html#cfn-events-connection-kmskeyidentifier
	//
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// The name for the connection to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html#cfn-events-connection-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

