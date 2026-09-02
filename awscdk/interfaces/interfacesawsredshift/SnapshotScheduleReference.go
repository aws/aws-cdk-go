package interfacesawsredshift


// A reference to a SnapshotSchedule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapshotScheduleReference := &SnapshotScheduleReference{
//   	SnapshotScheduleArn: jsii.String("snapshotScheduleArn"),
//   }
//
type SnapshotScheduleReference struct {
	// The Arn of the SnapshotSchedule resource.
	SnapshotScheduleArn *string `field:"required" json:"snapshotScheduleArn" yaml:"snapshotScheduleArn"`
}

