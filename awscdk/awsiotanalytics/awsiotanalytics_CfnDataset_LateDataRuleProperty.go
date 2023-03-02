package awsiotanalytics


// A structure that contains the name and configuration information of a late data rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lateDataRuleProperty := &lateDataRuleProperty{
//   	ruleConfiguration: &lateDataRuleConfigurationProperty{
//   		deltaTimeSessionWindowConfiguration: &deltaTimeSessionWindowConfigurationProperty{
//   			timeoutInMinutes: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ruleName: jsii.String("ruleName"),
//   }
//
type CfnDataset_LateDataRuleProperty struct {
	// The information needed to configure the late data rule.
	RuleConfiguration interface{} `field:"required" json:"ruleConfiguration" yaml:"ruleConfiguration"`
	// The name of the late data rule.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
}

