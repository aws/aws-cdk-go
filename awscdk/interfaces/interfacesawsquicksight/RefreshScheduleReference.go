package interfacesawsquicksight


// A reference to a RefreshSchedule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   refreshScheduleReference := &RefreshScheduleReference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	DataSetId: jsii.String("dataSetId"),
//   	RefreshScheduleArn: jsii.String("refreshScheduleArn"),
//   	ScheduleId: jsii.String("scheduleId"),
//   }
//
type RefreshScheduleReference struct {
	// The AwsAccountId of the RefreshSchedule resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The DataSetId of the RefreshSchedule resource.
	DataSetId *string `field:"required" json:"dataSetId" yaml:"dataSetId"`
	// The ARN of the RefreshSchedule resource.
	RefreshScheduleArn *string `field:"required" json:"refreshScheduleArn" yaml:"refreshScheduleArn"`
	// The Schedule/ScheduleId of the RefreshSchedule resource.
	ScheduleId *string `field:"required" json:"scheduleId" yaml:"scheduleId"`
}

