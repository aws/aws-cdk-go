package awsdatazone


// The Amaon Redshift lineage sync configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftLineageSyncConfigurationInputProperty := &RedshiftLineageSyncConfigurationInputProperty{
//   	Enabled: jsii.Boolean(false),
//   	Schedule: &LineageSyncScheduleProperty{
//   		Schedule: jsii.String("schedule"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftlineagesyncconfigurationinput.html
//
type CfnConnection_RedshiftLineageSyncConfigurationInputProperty struct {
	// Specifies whether the Amaon Redshift lineage sync configuration is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftlineagesyncconfigurationinput.html#cfn-datazone-connection-redshiftlineagesyncconfigurationinput-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The schedule of the Amaon Redshift lineage sync configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftlineagesyncconfigurationinput.html#cfn-datazone-connection-redshiftlineagesyncconfigurationinput-schedule
	//
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
}

