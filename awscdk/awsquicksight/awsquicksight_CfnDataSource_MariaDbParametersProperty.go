package awsquicksight


// The parameters for MariaDB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mariaDbParametersProperty := &mariaDbParametersProperty{
//   	database: jsii.String("database"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_MariaDbParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

