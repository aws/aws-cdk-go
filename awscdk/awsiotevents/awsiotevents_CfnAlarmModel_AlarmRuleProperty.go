package awsiotevents


// Defines when your alarm is invoked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmRuleProperty := &alarmRuleProperty{
//   	simpleRule: &simpleRuleProperty{
//   		comparisonOperator: jsii.String("comparisonOperator"),
//   		inputProperty: jsii.String("inputProperty"),
//   		threshold: jsii.String("threshold"),
//   	},
//   }
//
type CfnAlarmModel_AlarmRuleProperty struct {
	// A rule that compares an input property value to a threshold value with a comparison operator.
	SimpleRule interface{} `field:"optional" json:"simpleRule" yaml:"simpleRule"`
}

