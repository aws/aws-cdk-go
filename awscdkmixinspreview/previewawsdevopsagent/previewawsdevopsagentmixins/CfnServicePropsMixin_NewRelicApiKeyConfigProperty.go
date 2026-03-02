package previewawsdevopsagentmixins


// New Relic API key configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   newRelicApiKeyConfigProperty := &NewRelicApiKeyConfigProperty{
//   	AccountId: jsii.String("accountId"),
//   	AlertPolicyIds: []*string{
//   		jsii.String("alertPolicyIds"),
//   	},
//   	ApiKey: jsii.String("apiKey"),
//   	ApplicationIds: []*string{
//   		jsii.String("applicationIds"),
//   	},
//   	EntityGuids: []*string{
//   		jsii.String("entityGuids"),
//   	},
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html
//
type CfnServicePropsMixin_NewRelicApiKeyConfigProperty struct {
	// New Relic Account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html#cfn-devopsagent-service-newrelicapikeyconfig-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// List of alert policy IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html#cfn-devopsagent-service-newrelicapikeyconfig-alertpolicyids
	//
	AlertPolicyIds *[]*string `field:"optional" json:"alertPolicyIds" yaml:"alertPolicyIds"`
	// New Relic User API Key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html#cfn-devopsagent-service-newrelicapikeyconfig-apikey
	//
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// List of monitored APM application IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html#cfn-devopsagent-service-newrelicapikeyconfig-applicationids
	//
	ApplicationIds *[]*string `field:"optional" json:"applicationIds" yaml:"applicationIds"`
	// List of globally unique IDs for New Relic resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html#cfn-devopsagent-service-newrelicapikeyconfig-entityguids
	//
	EntityGuids *[]*string `field:"optional" json:"entityGuids" yaml:"entityGuids"`
	// New Relic region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicapikeyconfig.html#cfn-devopsagent-service-newrelicapikeyconfig-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

