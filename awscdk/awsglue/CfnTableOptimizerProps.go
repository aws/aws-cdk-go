package awsglue


// Properties for defining a `CfnTableOptimizer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTableOptimizerProps := &CfnTableOptimizerProps{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	TableName: jsii.String("tableName"),
//   	TableOptimizerConfiguration: &TableOptimizerConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		OrphanFileDeletionConfiguration: &OrphanFileDeletionConfigurationProperty{
//   			IcebergConfiguration: &IcebergConfigurationProperty{
//   				Location: jsii.String("location"),
//   				OrphanFileRetentionPeriodInDays: jsii.Number(123),
//   			},
//   		},
//   		RetentionConfiguration: &RetentionConfigurationProperty{
//   			IcebergConfiguration: &IcebergConfigurationProperty{
//   				Location: jsii.String("location"),
//   				OrphanFileRetentionPeriodInDays: jsii.Number(123),
//   			},
//   		},
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html
//
type CfnTableOptimizerProps struct {
	// The catalog ID of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html#cfn-glue-tableoptimizer-catalogid
	//
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The name of the database.
	//
	// For Hive compatibility, this is folded to lowercase when it is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html#cfn-glue-tableoptimizer-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The table name.
	//
	// For Hive compatibility, this must be entirely lowercase.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html#cfn-glue-tableoptimizer-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Specifies configuration details of a table optimizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html#cfn-glue-tableoptimizer-tableoptimizerconfiguration
	//
	TableOptimizerConfiguration interface{} `field:"required" json:"tableOptimizerConfiguration" yaml:"tableOptimizerConfiguration"`
	// The type of table optimizer.
	//
	// Currently, the only valid value is compaction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html#cfn-glue-tableoptimizer-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

