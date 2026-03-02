package previewawsdevopsagentmixins


// OAuth client credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var exchangeParameters interface{}
//
//   oAuthClientDetailsProperty := &OAuthClientDetailsProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientName: jsii.String("clientName"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	ExchangeParameters: exchangeParameters,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-oauthclientdetails.html
//
type CfnServicePropsMixin_OAuthClientDetailsProperty struct {
	// OAuth client ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-oauthclientdetails.html#cfn-devopsagent-service-oauthclientdetails-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// User friendly OAuth client name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-oauthclientdetails.html#cfn-devopsagent-service-oauthclientdetails-clientname
	//
	ClientName *string `field:"optional" json:"clientName" yaml:"clientName"`
	// OAuth client secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-oauthclientdetails.html#cfn-devopsagent-service-oauthclientdetails-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// OAuth token exchange parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-oauthclientdetails.html#cfn-devopsagent-service-oauthclientdetails-exchangeparameters
	//
	ExchangeParameters interface{} `field:"optional" json:"exchangeParameters" yaml:"exchangeParameters"`
}

