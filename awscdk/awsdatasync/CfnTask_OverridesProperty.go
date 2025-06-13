package awsdatasync


// Customizes the reporting level for aspects of your task report.
//
// For example, your report might generally only include errors, but you could specify that you want a list of successes and errors just for the files that Datasync attempted to delete in your destination location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   overridesProperty := &OverridesProperty{
//   	Deleted: &DeletedProperty{
//   		ReportLevel: jsii.String("reportLevel"),
//   	},
//   	Skipped: &SkippedProperty{
//   		ReportLevel: jsii.String("reportLevel"),
//   	},
//   	Transferred: &TransferredProperty{
//   		ReportLevel: jsii.String("reportLevel"),
//   	},
//   	Verified: &VerifiedProperty{
//   		ReportLevel: jsii.String("reportLevel"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-overrides.html
//
type CfnTask_OverridesProperty struct {
	// Specifies the level of reporting for the files, objects, and directories that Datasync attempted to delete in your destination location.
	//
	// This only applies if you configure your task to delete data in the destination that isn't in the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-overrides.html#cfn-datasync-task-overrides-deleted
	//
	Deleted interface{} `field:"optional" json:"deleted" yaml:"deleted"`
	// Specifies the level of reporting for the files, objects, and directories that Datasync attempted to skip during your transfer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-overrides.html#cfn-datasync-task-overrides-skipped
	//
	Skipped interface{} `field:"optional" json:"skipped" yaml:"skipped"`
	// Specifies the level of reporting for the files, objects, and directories that Datasync attempted to transfer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-overrides.html#cfn-datasync-task-overrides-transferred
	//
	Transferred interface{} `field:"optional" json:"transferred" yaml:"transferred"`
	// Specifies the level of reporting for the files, objects, and directories that Datasync attempted to verify at the end of your transfer.
	//
	// This only applies if you configure your task to verify data during and after the transfer (which Datasync does by default).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-overrides.html#cfn-datasync-task-overrides-verified
	//
	Verified interface{} `field:"optional" json:"verified" yaml:"verified"`
}

