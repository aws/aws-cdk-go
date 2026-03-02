package awsdevopsagent


// OAuth client credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var exchangeParameters interface{}
//
//   oAuthClientDetailsProperty := &OAuthClientDetailsProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	ClientName: jsii.String("clientName"),
//   	ExchangeParameters: exchangeParameters,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-oauthclientdetails.html
//
type CfnService_OAuthClientDetailsProperty struct {
	// OAuth client ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-oauthclientdetails.html#cfn-devopsagent-service-oauthclientdetails-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// OAuth client secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-oauthclientdetails.html#cfn-devopsagent-service-oauthclientdetails-clientsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// User friendly OAuth client name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-oauthclientdetails.html#cfn-devopsagent-service-oauthclientdetails-clientname
	//
	ClientName *string `field:"optional" json:"clientName" yaml:"clientName"`
	// OAuth token exchange parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-oauthclientdetails.html#cfn-devopsagent-service-oauthclientdetails-exchangeparameters
	//
	ExchangeParameters interface{} `field:"optional" json:"exchangeParameters" yaml:"exchangeParameters"`
}

