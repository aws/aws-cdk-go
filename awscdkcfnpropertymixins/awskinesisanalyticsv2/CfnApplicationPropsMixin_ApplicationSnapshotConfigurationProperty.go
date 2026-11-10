package awskinesisanalyticsv2


// Describes whether snapshots are enabled for a Managed Service for Apache Flink application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   applicationSnapshotConfigurationProperty := &ApplicationSnapshotConfigurationProperty{
//   	SnapshotMode: jsii.String("snapshotMode"),
//   	SnapshotsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationsnapshotconfiguration.html
//
type CfnApplicationPropsMixin_ApplicationSnapshotConfigurationProperty struct {
	// The snapshot mode of a Flink-based Kinesis Data Analytics application.
	//
	// Changing from APPLICATION_SNAPSHOTS_LEGACY to ENHANCED_SNAPSHOTS is one-way; an application in ENHANCED_SNAPSHOTS mode cannot return to APPLICATION_SNAPSHOTS_LEGACY.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationsnapshotconfiguration.html#cfn-kinesisanalyticsv2-application-applicationsnapshotconfiguration-snapshotmode
	//
	SnapshotMode *string `field:"optional" json:"snapshotMode" yaml:"snapshotMode"`
	// Describes whether snapshots are enabled for a Managed Service for Apache Flink application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationsnapshotconfiguration.html#cfn-kinesisanalyticsv2-application-applicationsnapshotconfiguration-snapshotsenabled
	//
	SnapshotsEnabled interface{} `field:"optional" json:"snapshotsEnabled" yaml:"snapshotsEnabled"`
}

