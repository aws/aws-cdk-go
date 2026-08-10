package awsstepfunctions


// Contains details about items processed in all child workflow executions started by a Map Run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mapRunItemCountsProperty := &MapRunItemCountsProperty{
//   	Aborted: jsii.Number(123),
//   	Failed: jsii.Number(123),
//   	FailuresNotRedrivable: jsii.Number(123),
//   	Pending: jsii.Number(123),
//   	PendingRedrive: jsii.Number(123),
//   	ResultsWritten: jsii.Number(123),
//   	Running: jsii.Number(123),
//   	Succeeded: jsii.Number(123),
//   	TimedOut: jsii.Number(123),
//   	Total: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunitemcounts.html
//
type CfnMapRun_MapRunItemCountsProperty struct {
	// The total number of items processed in child workflow executions that were stopped.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunitemcounts.html#cfn-stepfunctions-maprun-maprunitemcounts-aborted
	//
	Aborted *float64 `field:"optional" json:"aborted" yaml:"aborted"`
	// The total number of items processed in child workflow executions that have failed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunitemcounts.html#cfn-stepfunctions-maprun-maprunitemcounts-failed
	//
	Failed *float64 `field:"optional" json:"failed" yaml:"failed"`
	// The number of items in child workflow executions that cannot be redriven.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunitemcounts.html#cfn-stepfunctions-maprun-maprunitemcounts-failuresnotredrivable
	//
	FailuresNotRedrivable *float64 `field:"optional" json:"failuresNotRedrivable" yaml:"failuresNotRedrivable"`
	// The total number of items to process in child workflow executions that haven't started running yet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunitemcounts.html#cfn-stepfunctions-maprun-maprunitemcounts-pending
	//
	Pending *float64 `field:"optional" json:"pending" yaml:"pending"`
	// The number of unsuccessful items currently waiting to be redriven.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunitemcounts.html#cfn-stepfunctions-maprun-maprunitemcounts-pendingredrive
	//
	PendingRedrive *float64 `field:"optional" json:"pendingRedrive" yaml:"pendingRedrive"`
	// The count of items whose results were written by ResultWriter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunitemcounts.html#cfn-stepfunctions-maprun-maprunitemcounts-resultswritten
	//
	ResultsWritten *float64 `field:"optional" json:"resultsWritten" yaml:"resultsWritten"`
	// The total number of items being processed in child workflow executions that are currently in-progress.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunitemcounts.html#cfn-stepfunctions-maprun-maprunitemcounts-running
	//
	Running *float64 `field:"optional" json:"running" yaml:"running"`
	// The total number of items processed in child workflow executions that have completed successfully.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunitemcounts.html#cfn-stepfunctions-maprun-maprunitemcounts-succeeded
	//
	Succeeded *float64 `field:"optional" json:"succeeded" yaml:"succeeded"`
	// The total number of items processed in child workflow executions that have timed out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunitemcounts.html#cfn-stepfunctions-maprun-maprunitemcounts-timedout
	//
	TimedOut *float64 `field:"optional" json:"timedOut" yaml:"timedOut"`
	// The total number of items processed in all the child workflow executions started by a Map Run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunitemcounts.html#cfn-stepfunctions-maprun-maprunitemcounts-total
	//
	Total *float64 `field:"optional" json:"total" yaml:"total"`
}

