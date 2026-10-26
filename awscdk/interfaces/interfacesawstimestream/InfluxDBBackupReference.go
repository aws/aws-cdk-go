package interfacesawstimestream


// A reference to a InfluxDBBackup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   influxDBBackupReference := &InfluxDBBackupReference{
//   	InfluxDbBackupArn: jsii.String("influxDbBackupArn"),
//   }
//
type InfluxDBBackupReference struct {
	// The Arn of the InfluxDBBackup resource.
	InfluxDbBackupArn *string `field:"required" json:"influxDbBackupArn" yaml:"influxDbBackupArn"`
}

