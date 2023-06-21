package awssecurityhub


// Used to update information about the investigation into the finding.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowUpdateProperty := &WorkflowUpdateProperty{
//   	Status: jsii.String("status"),
//   }
//
type CfnAutomationRule_WorkflowUpdateProperty struct {
	// The status of the investigation into the finding.
	//
	// The workflow status is specific to an individual finding. It does not affect the generation of new findings. For example, setting the workflow status to `SUPPRESSED` or `RESOLVED` does not prevent a new finding for the same issue.
	//
	// The allowed values are the following.
	//
	// - `NEW` - The initial state of a finding, before it is reviewed.
	//
	// Security Hub also resets `WorkFlowStatus` from `NOTIFIED` or `RESOLVED` to `NEW` in the following cases:
	//
	// - The record state changes from `ARCHIVED` to `ACTIVE` .
	// - The compliance status changes from `PASSED` to either `WARNING` , `FAILED` , or `NOT_AVAILABLE` .
	// - `NOTIFIED` - Indicates that you notified the resource owner about the security issue. Used when the initial reviewer is not the resource owner, and needs intervention from the resource owner.
	// - `RESOLVED` - The finding was reviewed and remediated and is now considered resolved.
	// - `SUPPRESSED` - Indicates that you reviewed the finding and do not believe that any action is needed. The finding is no longer updated.
	Status *string `field:"required" json:"status" yaml:"status"`
}

