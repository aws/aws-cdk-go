package previewawsdevopsagentmixins


// New Relic authorization configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   newRelicAuthorizationConfigProperty := &NewRelicAuthorizationConfigProperty{
//   	ApiKey: &NewRelicApiKeyConfigProperty{
//   		AccountId: jsii.String("accountId"),
//   		AlertPolicyIds: []*string{
//   			jsii.String("alertPolicyIds"),
//   		},
//   		ApiKey: jsii.String("apiKey"),
//   		ApplicationIds: []*string{
//   			jsii.String("applicationIds"),
//   		},
//   		EntityGuids: []*string{
//   			jsii.String("entityGuids"),
//   		},
//   		Region: jsii.String("region"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicauthorizationconfig.html
//
type CfnServicePropsMixin_NewRelicAuthorizationConfigProperty struct {
	// New Relic API key configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicauthorizationconfig.html#cfn-devopsagent-service-newrelicauthorizationconfig-apikey
	//
	ApiKey interface{} `field:"optional" json:"apiKey" yaml:"apiKey"`
}

