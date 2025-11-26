package previewawsgluemixins


// Specifies configuration details of a table optimizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableOptimizerConfigurationProperty := &TableOptimizerConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	OrphanFileDeletionConfiguration: &OrphanFileDeletionConfigurationProperty{
//   		IcebergConfiguration: &IcebergConfigurationProperty{
//   			Location: jsii.String("location"),
//   			OrphanFileRetentionPeriodInDays: jsii.Number(123),
//   		},
//   	},
//   	RetentionConfiguration: &RetentionConfigurationProperty{
//   		IcebergConfiguration: &IcebergRetentionConfigurationProperty{
//   			CleanExpiredFiles: jsii.Boolean(false),
//   			NumberOfSnapshotsToRetain: jsii.Number(123),
//   			SnapshotRetentionPeriodInDays: jsii.Number(123),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		GlueConnectionName: jsii.String("glueConnectionName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-tableoptimizerconfiguration.html
//
type CfnTableOptimizerPropsMixin_TableOptimizerConfigurationProperty struct {
	// Whether the table optimization is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-tableoptimizerconfiguration.html#cfn-glue-tableoptimizer-tableoptimizerconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// `OrphanFileDeletionConfiguration` is a property that can be included within the TableOptimizer resource.
	//
	// It controls the automatic deletion of orphaned files - files that are not tracked by the table metadata, and older than the configured age limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-tableoptimizerconfiguration.html#cfn-glue-tableoptimizer-tableoptimizerconfiguration-orphanfiledeletionconfiguration
	//
	OrphanFileDeletionConfiguration interface{} `field:"optional" json:"orphanFileDeletionConfiguration" yaml:"orphanFileDeletionConfiguration"`
	// The configuration for a snapshot retention optimizer for Apache Iceberg tables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-tableoptimizerconfiguration.html#cfn-glue-tableoptimizer-tableoptimizerconfiguration-retentionconfiguration
	//
	RetentionConfiguration interface{} `field:"optional" json:"retentionConfiguration" yaml:"retentionConfiguration"`
	// A role passed by the caller which gives the service permission to update the resources associated with the optimizer on the caller's behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-tableoptimizerconfiguration.html#cfn-glue-tableoptimizer-tableoptimizerconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// An object that describes the VPC configuration for a table optimizer.
	//
	// This configuration is necessary to perform optimization on tables that are in a customer VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-tableoptimizerconfiguration.html#cfn-glue-tableoptimizer-tableoptimizerconfiguration-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

