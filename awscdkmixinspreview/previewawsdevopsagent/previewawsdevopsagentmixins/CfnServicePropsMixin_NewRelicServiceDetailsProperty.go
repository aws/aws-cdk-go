package previewawsdevopsagentmixins


// New Relic service configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   newRelicServiceDetailsProperty := &NewRelicServiceDetailsProperty{
//   	AuthorizationConfig: &NewRelicAuthorizationConfigProperty{
//   		ApiKey: &NewRelicApiKeyConfigProperty{
//   			AccountId: jsii.String("accountId"),
//   			AlertPolicyIds: []*string{
//   				jsii.String("alertPolicyIds"),
//   			},
//   			ApiKey: jsii.String("apiKey"),
//   			ApplicationIds: []*string{
//   				jsii.String("applicationIds"),
//   			},
//   			EntityGuids: []*string{
//   				jsii.String("entityGuids"),
//   			},
//   			Region: jsii.String("region"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicservicedetails.html
//
type CfnServicePropsMixin_NewRelicServiceDetailsProperty struct {
	// New Relic authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicservicedetails.html#cfn-devopsagent-service-newrelicservicedetails-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

