package awsdevopsagent


// PagerDuty service configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var exchangeParameters interface{}
//
//   pagerDutyDetailsProperty := &PagerDutyDetailsProperty{
//   	AuthorizationConfig: &PagerDutyAuthorizationConfigProperty{
//   		OAuthClientCredentials: &OAuthClientDetailsProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientName: jsii.String("clientName"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ExchangeParameters: exchangeParameters,
//   		},
//   	},
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-pagerdutydetails.html
//
type CfnServicePropsMixin_PagerDutyDetailsProperty struct {
	// PagerDuty OAuth authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-pagerdutydetails.html#cfn-devopsagent-service-pagerdutydetails-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// PagerDuty scopes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-pagerdutydetails.html#cfn-devopsagent-service-pagerdutydetails-scopes
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

