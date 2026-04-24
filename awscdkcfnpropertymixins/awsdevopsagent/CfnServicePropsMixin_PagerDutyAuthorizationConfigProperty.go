package awsdevopsagent


// PagerDuty OAuth authorization configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var exchangeParameters interface{}
//
//   pagerDutyAuthorizationConfigProperty := &PagerDutyAuthorizationConfigProperty{
//   	OAuthClientCredentials: &OAuthClientDetailsProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientName: jsii.String("clientName"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		ExchangeParameters: exchangeParameters,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-pagerdutyauthorizationconfig.html
//
type CfnServicePropsMixin_PagerDutyAuthorizationConfigProperty struct {
	// OAuth client credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-pagerdutyauthorizationconfig.html#cfn-devopsagent-service-pagerdutyauthorizationconfig-oauthclientcredentials
	//
	OAuthClientCredentials interface{} `field:"optional" json:"oAuthClientCredentials" yaml:"oAuthClientCredentials"`
}

