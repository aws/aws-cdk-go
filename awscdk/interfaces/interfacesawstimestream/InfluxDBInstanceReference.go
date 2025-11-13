package interfacesawstimestream


// A reference to a InfluxDBInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   influxDBInstanceReference := &InfluxDBInstanceReference{
//   	InfluxDbInstanceArn: jsii.String("influxDbInstanceArn"),
//   	InfluxDbInstanceId: jsii.String("influxDbInstanceId"),
//   }
//
type InfluxDBInstanceReference struct {
	// The ARN of the InfluxDBInstance resource.
	InfluxDbInstanceArn *string `field:"required" json:"influxDbInstanceArn" yaml:"influxDbInstanceArn"`
	// The Id of the InfluxDBInstance resource.
	InfluxDbInstanceId *string `field:"required" json:"influxDbInstanceId" yaml:"influxDbInstanceId"`
}

