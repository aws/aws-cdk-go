package awsdatazone


// Redshift Lineage Sync Configuration Input.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftlineagesyncconfigurationinput.html#cfn-datazone-connection-redshiftlineagesyncconfigurationinput-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Lineage Sync Schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftlineagesyncconfigurationinput.html#cfn-datazone-connection-redshiftlineagesyncconfigurationinput-schedule
	//
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
}

