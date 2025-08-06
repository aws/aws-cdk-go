package awss3tables


// Contains details about the snapshot management settings for an Iceberg table.
//
// A snapshot is expired when it exceeds MinSnapshotsToKeep and MaxSnapshotAgeHours.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapshotManagementProperty := &SnapshotManagementProperty{
//   	MaxSnapshotAgeHours: jsii.Number(123),
//   	MinSnapshotsToKeep: jsii.Number(123),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-snapshotmanagement.html
//
type CfnTable_SnapshotManagementProperty struct {
	// The maximum age of a snapshot before it can be expired.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-snapshotmanagement.html#cfn-s3tables-table-snapshotmanagement-maxsnapshotagehours
	//
	MaxSnapshotAgeHours *float64 `field:"optional" json:"maxSnapshotAgeHours" yaml:"maxSnapshotAgeHours"`
	// The minimum number of snapshots to keep.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-snapshotmanagement.html#cfn-s3tables-table-snapshotmanagement-minsnapshotstokeep
	//
	MinSnapshotsToKeep *float64 `field:"optional" json:"minSnapshotsToKeep" yaml:"minSnapshotsToKeep"`
	// Indicates whether the SnapshotManagement maintenance action is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-table-snapshotmanagement.html#cfn-s3tables-table-snapshotmanagement-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

