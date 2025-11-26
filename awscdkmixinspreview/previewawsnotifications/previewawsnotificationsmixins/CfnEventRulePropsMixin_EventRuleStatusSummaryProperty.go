package previewawsnotificationsmixins


// Provides additional information about the current `EventRule` status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventRuleStatusSummaryProperty := &EventRuleStatusSummaryProperty{
//   	Reason: jsii.String("reason"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notifications-eventrule-eventrulestatussummary.html
//
type CfnEventRulePropsMixin_EventRuleStatusSummaryProperty struct {
	// A human-readable reason for `EventRuleStatus` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notifications-eventrule-eventrulestatussummary.html#cfn-notifications-eventrule-eventrulestatussummary-reason
	//
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// The status of the `EventRule` .
	//
	// - Values:
	//
	// - `ACTIVE`
	//
	// - The `EventRule` can process events.
	// - `INACTIVE`
	//
	// - The `EventRule` may be unable to process events.
	// - `CREATING`
	//
	// - The `EventRule` is being created.
	//
	// Only `GET` and `LIST` calls can be run.
	// - `UPDATING`
	//
	// - The `EventRule` is being updated.
	//
	// Only `GET` and `LIST` calls can be run.
	// - `DELETING`
	//
	// - The `EventRule` is being deleted.
	//
	// Only `GET` and `LIST` calls can be run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notifications-eventrule-eventrulestatussummary.html#cfn-notifications-eventrule-eventrulestatussummary-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

