package previewawsdevopsagentmixins


// ServiceNow OAuth authorization configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var exchangeParameters interface{}
//
//   serviceNowAuthorizationConfigProperty := &ServiceNowAuthorizationConfigProperty{
//   	OAuthClientCredentials: &OAuthClientDetailsProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientName: jsii.String("clientName"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		ExchangeParameters: exchangeParameters,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicenowauthorizationconfig.html
//
type CfnServicePropsMixin_ServiceNowAuthorizationConfigProperty struct {
	// OAuth client credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicenowauthorizationconfig.html#cfn-devopsagent-service-servicenowauthorizationconfig-oauthclientcredentials
	//
	OAuthClientCredentials interface{} `field:"optional" json:"oAuthClientCredentials" yaml:"oAuthClientCredentials"`
}

