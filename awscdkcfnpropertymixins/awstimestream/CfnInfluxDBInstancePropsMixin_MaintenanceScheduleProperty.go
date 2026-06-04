package awstimestream


// The maintenance schedule for the InfluxDB instance.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbinstance-maintenanceschedule.html
//
type CfnInfluxDBInstancePropsMixin_MaintenanceScheduleProperty struct {
	// The preferred maintenance window in format ddd:HH:MM-ddd:HH:MM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbinstance-maintenanceschedule.html#cfn-timestream-influxdbinstance-maintenanceschedule-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The IANA timezone identifier for the maintenance schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbinstance-maintenanceschedule.html#cfn-timestream-influxdbinstance-maintenanceschedule-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

