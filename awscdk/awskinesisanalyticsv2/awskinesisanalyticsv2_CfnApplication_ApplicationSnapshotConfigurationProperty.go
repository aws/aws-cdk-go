package awskinesisanalyticsv2


// Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationSnapshotConfigurationProperty := &applicationSnapshotConfigurationProperty{
//   	snapshotsEnabled: jsii.Boolean(false),
//   }
//
type CfnApplication_ApplicationSnapshotConfigurationProperty struct {
	// Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.
	SnapshotsEnabled interface{} `field:"required" json:"snapshotsEnabled" yaml:"snapshotsEnabled"`
}

