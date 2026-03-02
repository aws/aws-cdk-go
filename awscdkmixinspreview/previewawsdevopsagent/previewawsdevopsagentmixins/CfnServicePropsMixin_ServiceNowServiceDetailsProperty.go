package previewawsdevopsagentmixins


// ServiceNow service configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var exchangeParameters interface{}
//
//   serviceNowServiceDetailsProperty := &ServiceNowServiceDetailsProperty{
//   	AuthorizationConfig: &ServiceNowAuthorizationConfigProperty{
//   		OAuthClientCredentials: &OAuthClientDetailsProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientName: jsii.String("clientName"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ExchangeParameters: exchangeParameters,
//   		},
//   	},
//   	InstanceUrl: jsii.String("instanceUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicenowservicedetails.html
//
type CfnServicePropsMixin_ServiceNowServiceDetailsProperty struct {
	// ServiceNow OAuth authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicenowservicedetails.html#cfn-devopsagent-service-servicenowservicedetails-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// ServiceNow instance URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicenowservicedetails.html#cfn-devopsagent-service-servicenowservicedetails-instanceurl
	//
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
}

