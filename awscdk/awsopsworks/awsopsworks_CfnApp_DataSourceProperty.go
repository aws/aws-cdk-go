package awsopsworks


// Describes an app's data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceProperty := &dataSourceProperty{
//   	arn: jsii.String("arn"),
//   	databaseName: jsii.String("databaseName"),
//   	type: jsii.String("type"),
//   }
//
type CfnApp_DataSourceProperty struct {
	// The data source's ARN.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The database name.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The data source's type, `AutoSelectOpsworksMysqlInstance` , `OpsworksMysqlInstance` , `RdsDbInstance` , or `None` .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

