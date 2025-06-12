package awskinesisanalyticsv2


// Describes the starting parameters for an Managed Service for Apache Flink application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runConfigurationProperty := &RunConfigurationProperty{
//   	ApplicationRestoreConfiguration: &ApplicationRestoreConfigurationProperty{
//   		ApplicationRestoreType: jsii.String("applicationRestoreType"),
//
//   		// the properties below are optional
//   		SnapshotName: jsii.String("snapshotName"),
//   	},
//   	FlinkRunConfiguration: &FlinkRunConfigurationProperty{
//   		AllowNonRestoredState: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-runconfiguration.html
//
type CfnApplication_RunConfigurationProperty struct {
	// Describes the restore behavior of a restarting application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-runconfiguration.html#cfn-kinesisanalyticsv2-application-runconfiguration-applicationrestoreconfiguration
	//
	ApplicationRestoreConfiguration interface{} `field:"optional" json:"applicationRestoreConfiguration" yaml:"applicationRestoreConfiguration"`
	// Describes the starting parameters for a Managed Service for Apache Flink application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-runconfiguration.html#cfn-kinesisanalyticsv2-application-runconfiguration-flinkrunconfiguration
	//
	FlinkRunConfiguration interface{} `field:"optional" json:"flinkRunConfiguration" yaml:"flinkRunConfiguration"`
}

