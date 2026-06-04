package awstimestream


// The maintenance schedule for the InfluxDB cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceScheduleProperty := &MaintenanceScheduleProperty{
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	Timezone: jsii.String("timezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbcluster-maintenanceschedule.html
//
type CfnInfluxDBCluster_MaintenanceScheduleProperty struct {
	// The preferred maintenance window in format ddd:HH:MM-ddd:HH:MM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbcluster-maintenanceschedule.html#cfn-timestream-influxdbcluster-maintenanceschedule-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"required" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The IANA timezone identifier for the maintenance schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbcluster-maintenanceschedule.html#cfn-timestream-influxdbcluster-maintenanceschedule-timezone
	//
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

