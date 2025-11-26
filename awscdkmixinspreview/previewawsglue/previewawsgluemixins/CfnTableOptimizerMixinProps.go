package previewawsgluemixins


// Properties for CfnTableOptimizerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTableOptimizerMixinProps := &CfnTableOptimizerMixinProps{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	TableName: jsii.String("tableName"),
//   	TableOptimizerConfiguration: &TableOptimizerConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		OrphanFileDeletionConfiguration: &OrphanFileDeletionConfigurationProperty{
//   			IcebergConfiguration: &IcebergConfigurationProperty{
//   				Location: jsii.String("location"),
//   				OrphanFileRetentionPeriodInDays: jsii.Number(123),
//   			},
//   		},
//   		RetentionConfiguration: &RetentionConfigurationProperty{
//   			IcebergConfiguration: &IcebergRetentionConfigurationProperty{
//   				CleanExpiredFiles: jsii.Boolean(false),
//   				NumberOfSnapshotsToRetain: jsii.Number(123),
//   				SnapshotRetentionPeriodInDays: jsii.Number(123),
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html
//
type CfnTableOptimizerMixinProps struct {
	// The catalog ID of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html#cfn-glue-tableoptimizer-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the database.
	//
	// For Hive compatibility, this is folded to lowercase when it is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html#cfn-glue-tableoptimizer-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The table name.
	//
	// For Hive compatibility, this must be entirely lowercase.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html#cfn-glue-tableoptimizer-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Specifies configuration details of a table optimizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html#cfn-glue-tableoptimizer-tableoptimizerconfiguration
	//
	TableOptimizerConfiguration interface{} `field:"optional" json:"tableOptimizerConfiguration" yaml:"tableOptimizerConfiguration"`
	// The type of table optimizer. The valid values are:.
	//
	// - compaction - for managing compaction with a table optimizer.
	// - retention - for managing the retention of snapshot with a table optimizer.
	// - orphan_file_deletion - for managing the deletion of orphan files with a table optimizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html#cfn-glue-tableoptimizer-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

