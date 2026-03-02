package previewawsdevopsagentmixins


// Dynatrace service configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var exchangeParameters interface{}
//
//   dynatraceServiceDetailsProperty := &DynatraceServiceDetailsProperty{
//   	AccountUrn: jsii.String("accountUrn"),
//   	AuthorizationConfig: &DynatraceAuthorizationConfigProperty{
//   		OAuthClientCredentials: &OAuthClientDetailsProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientName: jsii.String("clientName"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ExchangeParameters: exchangeParameters,
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-dynatraceservicedetails.html
//
type CfnServicePropsMixin_DynatraceServiceDetailsProperty struct {
	// Dynatrace resource account URN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-dynatraceservicedetails.html#cfn-devopsagent-service-dynatraceservicedetails-accounturn
	//
	AccountUrn *string `field:"optional" json:"accountUrn" yaml:"accountUrn"`
	// Dynatrace OAuth authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-dynatraceservicedetails.html#cfn-devopsagent-service-dynatraceservicedetails-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

