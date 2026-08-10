package awssagemaker


// The schedule configuration for automatic patching.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   patchScheduleProperty := &PatchScheduleProperty{
//   	NextPatchDate: jsii.String("nextPatchDate"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-patchschedule.html
//
type CfnCluster_PatchScheduleProperty struct {
	// The date and time of the next scheduled patch, set by the system when a patch AMI is detected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-patchschedule.html#cfn-sagemaker-cluster-patchschedule-nextpatchdate
	//
	NextPatchDate *string `field:"optional" json:"nextPatchDate" yaml:"nextPatchDate"`
}

