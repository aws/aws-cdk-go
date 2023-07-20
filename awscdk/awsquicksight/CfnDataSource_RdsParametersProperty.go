package awsquicksight


// The parameters for Amazon RDS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rdsParametersProperty := &RdsParametersProperty{
//   	Database: jsii.String("database"),
//   	InstanceId: jsii.String("instanceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-rdsparameters.html
//
type CfnDataSource_RdsParametersProperty struct {
	// Database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-rdsparameters.html#cfn-quicksight-datasource-rdsparameters-database
	//
	Database *string `field:"required" json:"database" yaml:"database"`
	// Instance ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-rdsparameters.html#cfn-quicksight-datasource-rdsparameters-instanceid
	//
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
}

