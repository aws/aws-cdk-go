package previewawsdevopsagentmixins


// Dynatrace OAuth authorization configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var exchangeParameters interface{}
//
//   dynatraceAuthorizationConfigProperty := &DynatraceAuthorizationConfigProperty{
//   	OAuthClientCredentials: &OAuthClientDetailsProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientName: jsii.String("clientName"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		ExchangeParameters: exchangeParameters,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-dynatraceauthorizationconfig.html
//
type CfnServicePropsMixin_DynatraceAuthorizationConfigProperty struct {
	// OAuth client credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-dynatraceauthorizationconfig.html#cfn-devopsagent-service-dynatraceauthorizationconfig-oauthclientcredentials
	//
	OAuthClientCredentials interface{} `field:"optional" json:"oAuthClientCredentials" yaml:"oAuthClientCredentials"`
}

