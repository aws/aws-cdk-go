package previewawsdevopsagentmixins


// MCP server OAuth client credentials configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var exchangeParameters interface{}
//
//   mCPServerOAuthClientCredentialsConfigProperty := &MCPServerOAuthClientCredentialsConfigProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientName: jsii.String("clientName"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	ExchangeParameters: exchangeParameters,
//   	ExchangeUrl: jsii.String("exchangeUrl"),
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html
//
type CfnServicePropsMixin_MCPServerOAuthClientCredentialsConfigProperty struct {
	// OAuth client ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html#cfn-devopsagent-service-mcpserveroauthclientcredentialsconfig-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// User friendly OAuth client name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html#cfn-devopsagent-service-mcpserveroauthclientcredentialsconfig-clientname
	//
	ClientName *string `field:"optional" json:"clientName" yaml:"clientName"`
	// OAuth client secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html#cfn-devopsagent-service-mcpserveroauthclientcredentialsconfig-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// OAuth token exchange parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html#cfn-devopsagent-service-mcpserveroauthclientcredentialsconfig-exchangeparameters
	//
	ExchangeParameters interface{} `field:"optional" json:"exchangeParameters" yaml:"exchangeParameters"`
	// OAuth token exchange URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html#cfn-devopsagent-service-mcpserveroauthclientcredentialsconfig-exchangeurl
	//
	ExchangeUrl *string `field:"optional" json:"exchangeUrl" yaml:"exchangeUrl"`
	// OAuth scopes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.html#cfn-devopsagent-service-mcpserveroauthclientcredentialsconfig-scopes
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

