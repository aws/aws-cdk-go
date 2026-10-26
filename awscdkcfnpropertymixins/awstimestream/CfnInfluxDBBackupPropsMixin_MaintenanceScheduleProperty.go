package awstimestream


// The maintenance schedule of the resource at the time of backup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   maintenanceScheduleProperty := &MaintenanceScheduleProperty{
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	Timezone: jsii.String("timezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbbackup-maintenanceschedule.html
//
type CfnInfluxDBBackupPropsMixin_MaintenanceScheduleProperty struct {
	// The preferred maintenance window in the format ddd:HH:MM-ddd:HH:MM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbbackup-maintenanceschedule.html#cfn-timestream-influxdbbackup-maintenanceschedule-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The IANA timezone identifier for the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbbackup-maintenanceschedule.html#cfn-timestream-influxdbbackup-maintenanceschedule-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

