package awsstepfunctions


// Contains details about all child workflow executions started by a Map Run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mapRunExecutionCountsProperty := &MapRunExecutionCountsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunexecutioncounts.html
//
type CfnMapRunPropsMixin_MapRunExecutionCountsProperty struct {
	// The total number of child workflow executions that were stopped.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunexecutioncounts.html#cfn-stepfunctions-maprun-maprunexecutioncounts-aborted
	//
	Aborted *float64 `field:"optional" json:"aborted" yaml:"aborted"`
	// The total number of child workflow executions that have failed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunexecutioncounts.html#cfn-stepfunctions-maprun-maprunexecutioncounts-failed
	//
	Failed *float64 `field:"optional" json:"failed" yaml:"failed"`
	// The number of child workflow executions that cannot be redriven.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunexecutioncounts.html#cfn-stepfunctions-maprun-maprunexecutioncounts-failuresnotredrivable
	//
	FailuresNotRedrivable *float64 `field:"optional" json:"failuresNotRedrivable" yaml:"failuresNotRedrivable"`
	// The total number of child workflow executions that haven't started executing yet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunexecutioncounts.html#cfn-stepfunctions-maprun-maprunexecutioncounts-pending
	//
	Pending *float64 `field:"optional" json:"pending" yaml:"pending"`
	// The number of unsuccessful child workflow executions waiting to be redriven.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunexecutioncounts.html#cfn-stepfunctions-maprun-maprunexecutioncounts-pendingredrive
	//
	PendingRedrive *float64 `field:"optional" json:"pendingRedrive" yaml:"pendingRedrive"`
	// The count of child workflow executions whose results were written by ResultWriter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunexecutioncounts.html#cfn-stepfunctions-maprun-maprunexecutioncounts-resultswritten
	//
	ResultsWritten *float64 `field:"optional" json:"resultsWritten" yaml:"resultsWritten"`
	// The total number of child workflow executions that are currently in-progress.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunexecutioncounts.html#cfn-stepfunctions-maprun-maprunexecutioncounts-running
	//
	Running *float64 `field:"optional" json:"running" yaml:"running"`
	// The total number of child workflow executions that have completed successfully.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunexecutioncounts.html#cfn-stepfunctions-maprun-maprunexecutioncounts-succeeded
	//
	Succeeded *float64 `field:"optional" json:"succeeded" yaml:"succeeded"`
	// The total number of child workflow executions that have timed out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunexecutioncounts.html#cfn-stepfunctions-maprun-maprunexecutioncounts-timedout
	//
	TimedOut *float64 `field:"optional" json:"timedOut" yaml:"timedOut"`
	// The total number of child workflow executions started by a Map Run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-maprun-maprunexecutioncounts.html#cfn-stepfunctions-maprun-maprunexecutioncounts-total
	//
	Total *float64 `field:"optional" json:"total" yaml:"total"`
}

