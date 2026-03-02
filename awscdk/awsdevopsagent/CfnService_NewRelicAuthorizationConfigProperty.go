package awsdevopsagent


// New Relic authorization configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   newRelicAuthorizationConfigProperty := &NewRelicAuthorizationConfigProperty{
//   	ApiKey: &NewRelicApiKeyConfigProperty{
//   		AccountId: jsii.String("accountId"),
//   		ApiKey: jsii.String("apiKey"),
//   		Region: jsii.String("region"),
//
//   		// the properties below are optional
//   		AlertPolicyIds: []*string{
//   			jsii.String("alertPolicyIds"),
//   		},
//   		ApplicationIds: []*string{
//   			jsii.String("applicationIds"),
//   		},
//   		EntityGuids: []*string{
//   			jsii.String("entityGuids"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicauthorizationconfig.html
//
type CfnService_NewRelicAuthorizationConfigProperty struct {
	// New Relic API key configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicauthorizationconfig.html#cfn-devopsagent-service-newrelicauthorizationconfig-apikey
	//
	ApiKey interface{} `field:"required" json:"apiKey" yaml:"apiKey"`
}

