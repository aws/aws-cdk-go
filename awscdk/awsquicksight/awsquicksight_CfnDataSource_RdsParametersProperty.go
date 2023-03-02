package awsquicksight


// The parameters for Amazon RDS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rdsParametersProperty := &rdsParametersProperty{
//   	database: jsii.String("database"),
//   	instanceId: jsii.String("instanceId"),
//   }
//
type CfnDataSource_RdsParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Instance ID.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
}

