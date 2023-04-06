package awsquicksight


// The parameters for Snowflake.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snowflakeParametersProperty := &SnowflakeParametersProperty{
//   	Database: jsii.String("database"),
//   	Host: jsii.String("host"),
//   	Warehouse: jsii.String("warehouse"),
//   }
//
type CfnDataSource_SnowflakeParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Warehouse.
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
}

