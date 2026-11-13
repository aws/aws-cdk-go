package awscloudwatch


// The alert states that trigger the notification rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerProperty := &TriggerProperty{
//   	StateValues: []*string{
//   		jsii.String("stateValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-trigger.html
//
type CfnAlert_TriggerProperty struct {
	// The alert states that trigger this rule.
	//
	// A notification is sent when the alert enters any listed state. If omitted or empty, every state triggers the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-trigger.html#cfn-cloudwatch-alert-trigger-statevalues
	//
	StateValues *[]*string `field:"optional" json:"stateValues" yaml:"stateValues"`
}

