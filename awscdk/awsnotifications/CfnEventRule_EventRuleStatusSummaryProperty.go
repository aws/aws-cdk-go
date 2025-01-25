package awsnotifications


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventRuleStatusSummaryProperty := &EventRuleStatusSummaryProperty{
//   	Reason: jsii.String("reason"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notifications-eventrule-eventrulestatussummary.html
//
type CfnEventRule_EventRuleStatusSummaryProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notifications-eventrule-eventrulestatussummary.html#cfn-notifications-eventrule-eventrulestatussummary-reason
	//
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notifications-eventrule-eventrulestatussummary.html#cfn-notifications-eventrule-eventrulestatussummary-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
}

