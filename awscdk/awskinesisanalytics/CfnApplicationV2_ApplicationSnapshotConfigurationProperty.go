package awskinesisanalytics


// Describes whether snapshots are enabled for a Managed Service for Apache Flink application.
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
type CfnApplicationV2_ApplicationSnapshotConfigurationProperty struct {
	// Describes whether snapshots are enabled for a Managed Service for Apache Flink application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationsnapshotconfiguration.html#cfn-kinesisanalyticsv2-application-applicationsnapshotconfiguration-snapshotsenabled
	//
	SnapshotsEnabled interface{} `field:"required" json:"snapshotsEnabled" yaml:"snapshotsEnabled"`
}

