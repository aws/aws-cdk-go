package awsglue


// An object representing an AWS Glue table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataQualityTargetTableProperty := &DataQualityTargetTableProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	TableName: jsii.String("tableName"),
//   }
//
type CfnDataQualityRuleset_DataQualityTargetTableProperty struct {
	// The name of the database where the AWS Glue table exists.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of the AWS Glue table.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

