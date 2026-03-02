package previewawsdevopsagentmixins


// MCP server configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   			ClientName: jsii.String("clientName"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ExchangeParameters: exchangeParameters,
//   			ExchangeUrl: jsii.String("exchangeUrl"),
//   			Scopes: []*string{
//   				jsii.String("scopes"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverdetails.html
//
type CfnServicePropsMixin_MCPServerDetailsProperty struct {
	// MCP server authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverdetails.html#cfn-devopsagent-service-mcpserverdetails-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// Optional description for the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverdetails.html#cfn-devopsagent-service-mcpserverdetails-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverdetails.html#cfn-devopsagent-service-mcpserverdetails-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// MCP server name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserverdetails.html#cfn-devopsagent-service-mcpserverdetails-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

