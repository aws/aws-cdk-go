package awsdevopsagent


// ServiceNow service configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var exchangeParameters interface{}
//
//   serviceNowServiceDetailsProperty := &ServiceNowServiceDetailsProperty{
//   	InstanceUrl: jsii.String("instanceUrl"),
//
//   	// the properties below are optional
//   	AuthorizationConfig: &ServiceNowAuthorizationConfigProperty{
//   		OAuthClientCredentials: &OAuthClientDetailsProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//
//   			// the properties below are optional
//   			ClientName: jsii.String("clientName"),
//   			ExchangeParameters: exchangeParameters,
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicenowservicedetails.html
//
type CfnService_ServiceNowServiceDetailsProperty struct {
	// ServiceNow instance URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicenowservicedetails.html#cfn-devopsagent-service-servicenowservicedetails-instanceurl
	//
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
	// ServiceNow OAuth authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicenowservicedetails.html#cfn-devopsagent-service-servicenowservicedetails-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

