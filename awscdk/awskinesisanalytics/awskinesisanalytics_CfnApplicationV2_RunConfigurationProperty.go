package awskinesisanalytics


// Describes the starting parameters for an Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runConfigurationProperty := &runConfigurationProperty{
//   	applicationRestoreConfiguration: &applicationRestoreConfigurationProperty{
//   		applicationRestoreType: jsii.String("applicationRestoreType"),
//
//   		// the properties below are optional
//   		snapshotName: jsii.String("snapshotName"),
//   	},
//   	flinkRunConfiguration: &flinkRunConfigurationProperty{
//   		allowNonRestoredState: jsii.Boolean(false),
//   	},
//   }
//
type CfnApplicationV2_RunConfigurationProperty struct {
	// Describes the restore behavior of a restarting application.
	ApplicationRestoreConfiguration interface{} `field:"optional" json:"applicationRestoreConfiguration" yaml:"applicationRestoreConfiguration"`
	// Describes the starting parameters for a Flink-based Kinesis Data Analytics application.
	FlinkRunConfiguration interface{} `field:"optional" json:"flinkRunConfiguration" yaml:"flinkRunConfiguration"`
}

