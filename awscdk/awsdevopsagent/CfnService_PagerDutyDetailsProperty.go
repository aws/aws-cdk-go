package awsdevopsagent


// PagerDuty service configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var exchangeParameters interface{}
//
//   pagerDutyDetailsProperty := &PagerDutyDetailsProperty{
//   	AuthorizationConfig: &PagerDutyAuthorizationConfigProperty{
//   		OAuthClientCredentials: &OAuthClientDetailsProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//
//   			// the properties below are optional
//   			ClientName: jsii.String("clientName"),
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
type CfnService_PagerDutyDetailsProperty struct {
	// PagerDuty OAuth authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-pagerdutydetails.html#cfn-devopsagent-service-pagerdutydetails-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"required" json:"authorizationConfig" yaml:"authorizationConfig"`
	// PagerDuty scopes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-pagerdutydetails.html#cfn-devopsagent-service-pagerdutydetails-scopes
	//
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
}

