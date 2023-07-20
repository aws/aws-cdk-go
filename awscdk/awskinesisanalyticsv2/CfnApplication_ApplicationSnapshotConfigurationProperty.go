package awskinesisanalyticsv2


// Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationSnapshotConfigurationProperty := &ApplicationSnapshotConfigurationProperty{
//   	SnapshotsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationsnapshotconfiguration.html
//
type CfnApplication_ApplicationSnapshotConfigurationProperty struct {
	// Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationsnapshotconfiguration.html#cfn-kinesisanalyticsv2-application-applicationsnapshotconfiguration-snapshotsenabled
	//
	SnapshotsEnabled interface{} `field:"required" json:"snapshotsEnabled" yaml:"snapshotsEnabled"`
}

