package previewawsec2events


// Type definition for Snapshot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snapshot := &Snapshot{
//   	SnapshotId: []*string{
//   		jsii.String("snapshotId"),
//   	},
//   	Source: []*string{
//   		jsii.String("source"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   }
//
// Experimental.
type EBSMultiVolumeSnapshotsCompletionStatus_Snapshot struct {
	// snapshot_id property.
	//
	// Specify an array of string values to match this event if the actual value of snapshot_id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SnapshotId *[]*string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// source property.
	//
	// Specify an array of string values to match this event if the actual value of source is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Source *[]*string `field:"optional" json:"source" yaml:"source"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
}

