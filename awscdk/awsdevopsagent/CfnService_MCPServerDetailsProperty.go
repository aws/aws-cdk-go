package awsdevopsagent


// MCP server configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var exchangeParameters interface{}
//
//   mCPServerDetailsProperty := &MCPServerDetailsProperty{
//   	AuthorizationConfig: &MCPServerAuthorizationConfigProperty{
//   		ApiKey: &ApiKeyDetailsProperty{
//   			ApiKeyHeader: jsii.String("apiKeyHeader"),
//   			ApiKeyName: jsii.String("apiKeyName"),
//   			ApiKeyValue: jsii.String("apiKeyValue"),
//   		},
//   		OAuthClientCredentials: &MCPServerOAuthClientCredentialsConfigProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ExchangeUrl: jsii.String("exchangeUrl"),
//
//   			// the properties below are optional
//   			ClientName: jsii.String("clientName"),
//   			ExchangeParameters: exchangeParameters,
//   			Scopes: []*string{
//   				jsii.String("scopes"),
//   			},
//   		},
//   	},
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverdetails.html
//
type CfnService_MCPServerDetailsProperty struct {
	// MCP server authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverdetails.html#cfn-devopsagent-service-mcpserverdetails-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"required" json:"authorizationConfig" yaml:"authorizationConfig"`
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverdetails.html#cfn-devopsagent-service-mcpserverdetails-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// MCP server name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverdetails.html#cfn-devopsagent-service-mcpserverdetails-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional description for the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverdetails.html#cfn-devopsagent-service-mcpserverdetails-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

