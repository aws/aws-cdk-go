package awsdevopsagent


// New Relic API key configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   newRelicApiKeyConfigProperty := &NewRelicApiKeyConfigProperty{
//   	AccountId: jsii.String("accountId"),
//   	ApiKey: jsii.String("apiKey"),
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	AlertPolicyIds: []*string{
//   		jsii.String("alertPolicyIds"),
//   	},
//   	ApplicationIds: []*string{
//   		jsii.String("applicationIds"),
//   	},
//   	EntityGuids: []*string{
//   		jsii.String("entityGuids"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html
//
type CfnService_NewRelicApiKeyConfigProperty struct {
	// New Relic Account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html#cfn-devopsagent-service-newrelicapikeyconfig-accountid
	//
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// New Relic User API Key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html#cfn-devopsagent-service-newrelicapikeyconfig-apikey
	//
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// New Relic region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html#cfn-devopsagent-service-newrelicapikeyconfig-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	// List of alert policy IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html#cfn-devopsagent-service-newrelicapikeyconfig-alertpolicyids
	//
	AlertPolicyIds *[]*string `field:"optional" json:"alertPolicyIds" yaml:"alertPolicyIds"`
	// List of monitored APM application IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html#cfn-devopsagent-service-newrelicapikeyconfig-applicationids
	//
	ApplicationIds *[]*string `field:"optional" json:"applicationIds" yaml:"applicationIds"`
	// List of globally unique IDs for New Relic resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html#cfn-devopsagent-service-newrelicapikeyconfig-entityguids
	//
	EntityGuids *[]*string `field:"optional" json:"entityGuids" yaml:"entityGuids"`
}

