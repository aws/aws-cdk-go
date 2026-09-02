package awsdeadline


// The usage details of the allotted budget.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   usageTrackingResourceProperty := &UsageTrackingResourceProperty{
//   	QueueId: jsii.String("queueId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-usagetrackingresource.html
//
type CfnBudget_UsageTrackingResourceProperty struct {
	// The queue ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-usagetrackingresource.html#cfn-deadline-budget-usagetrackingresource-queueid
	//
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

