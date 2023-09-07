package awsdatasync


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-overrides.html#cfn-datasync-task-overrides-deleted
	//
	Deleted interface{} `field:"optional" json:"deleted" yaml:"deleted"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-overrides.html#cfn-datasync-task-overrides-skipped
	//
	Skipped interface{} `field:"optional" json:"skipped" yaml:"skipped"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-overrides.html#cfn-datasync-task-overrides-transferred
	//
	Transferred interface{} `field:"optional" json:"transferred" yaml:"transferred"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-overrides.html#cfn-datasync-task-overrides-verified
	//
	Verified interface{} `field:"optional" json:"verified" yaml:"verified"`
}

