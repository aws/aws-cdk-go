package awskinesisanalyticsv2


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
type CfnApplication_RunConfigurationProperty struct {
	// `CfnApplication.RunConfigurationProperty.ApplicationRestoreConfiguration`.
	ApplicationRestoreConfiguration interface{} `field:"optional" json:"applicationRestoreConfiguration" yaml:"applicationRestoreConfiguration"`
	// `CfnApplication.RunConfigurationProperty.FlinkRunConfiguration`.
	FlinkRunConfiguration interface{} `field:"optional" json:"flinkRunConfiguration" yaml:"flinkRunConfiguration"`
}

