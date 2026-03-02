package awsdevopsagent


// MCP server OAuth client credentials configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var exchangeParameters interface{}
//
//   mCPServerOAuthClientCredentialsConfigProperty := &MCPServerOAuthClientCredentialsConfigProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	ExchangeUrl: jsii.String("exchangeUrl"),
//
//   	// the properties below are optional
//   	ClientName: jsii.String("clientName"),
//   	ExchangeParameters: exchangeParameters,
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html
//
type CfnService_MCPServerOAuthClientCredentialsConfigProperty struct {
	// OAuth client ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html#cfn-devopsagent-service-mcpserveroauthclientcredentialsconfig-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// OAuth client secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html#cfn-devopsagent-service-mcpserveroauthclientcredentialsconfig-clientsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// OAuth token exchange URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html#cfn-devopsagent-service-mcpserveroauthclientcredentialsconfig-exchangeurl
	//
	ExchangeUrl *string `field:"required" json:"exchangeUrl" yaml:"exchangeUrl"`
	// User friendly OAuth client name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html#cfn-devopsagent-service-mcpserveroauthclientcredentialsconfig-clientname
	//
	ClientName *string `field:"optional" json:"clientName" yaml:"clientName"`
	// OAuth token exchange parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html#cfn-devopsagent-service-mcpserveroauthclientcredentialsconfig-exchangeparameters
	//
	ExchangeParameters interface{} `field:"optional" json:"exchangeParameters" yaml:"exchangeParameters"`
	// OAuth scopes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html#cfn-devopsagent-service-mcpserveroauthclientcredentialsconfig-scopes
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

