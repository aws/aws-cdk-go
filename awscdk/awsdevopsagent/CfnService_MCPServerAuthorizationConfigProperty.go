package awsdevopsagent


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var exchangeParameters interface{}
//
//   mCPServerAuthorizationConfigProperty := &MCPServerAuthorizationConfigProperty{
//   	ApiKey: &ApiKeyDetailsProperty{
//   		ApiKeyHeader: jsii.String("apiKeyHeader"),
//   		ApiKeyName: jsii.String("apiKeyName"),
//   		ApiKeyValue: jsii.String("apiKeyValue"),
//   	},
//   	OAuthClientCredentials: &MCPServerOAuthClientCredentialsConfigProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		ExchangeUrl: jsii.String("exchangeUrl"),
//
//   		// the properties below are optional
//   		ClientName: jsii.String("clientName"),
//   		ExchangeParameters: exchangeParameters,
//   		Scopes: []*string{
//   			jsii.String("scopes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverauthorizationconfig.html
//
type CfnService_MCPServerAuthorizationConfigProperty struct {
	// API key authentication details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverauthorizationconfig.html#cfn-devopsagent-service-mcpserverauthorizationconfig-apikey
	//
	ApiKey interface{} `field:"optional" json:"apiKey" yaml:"apiKey"`
	// MCP server OAuth client credentials configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverauthorizationconfig.html#cfn-devopsagent-service-mcpserverauthorizationconfig-oauthclientcredentials
	//
	OAuthClientCredentials interface{} `field:"optional" json:"oAuthClientCredentials" yaml:"oAuthClientCredentials"`
}

