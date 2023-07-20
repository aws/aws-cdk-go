package awsiotanalytics


// A structure that contains the name and configuration information of a late data rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lateDataRuleProperty := &LateDataRuleProperty{
//   	RuleConfiguration: &LateDataRuleConfigurationProperty{
//   		DeltaTimeSessionWindowConfiguration: &DeltaTimeSessionWindowConfigurationProperty{
//   			TimeoutInMinutes: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	RuleName: jsii.String("ruleName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-latedatarule.html
//
type CfnDataset_LateDataRuleProperty struct {
	// The information needed to configure the late data rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-latedatarule.html#cfn-iotanalytics-dataset-latedatarule-ruleconfiguration
	//
	RuleConfiguration interface{} `field:"required" json:"ruleConfiguration" yaml:"ruleConfiguration"`
	// The name of the late data rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-latedatarule.html#cfn-iotanalytics-dataset-latedatarule-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
}

