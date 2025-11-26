package previewawsgluemixins


// The configuration for a snapshot retention optimizer for Apache Iceberg tables.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retentionConfigurationProperty := &RetentionConfigurationProperty{
//   	IcebergConfiguration: &IcebergRetentionConfigurationProperty{
//   		CleanExpiredFiles: jsii.Boolean(false),
//   		NumberOfSnapshotsToRetain: jsii.Number(123),
//   		SnapshotRetentionPeriodInDays: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-retentionconfiguration.html
//
type CfnTableOptimizerPropsMixin_RetentionConfigurationProperty struct {
	// The configuration for an Iceberg snapshot retention optimizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-retentionconfiguration.html#cfn-glue-tableoptimizer-retentionconfiguration-icebergconfiguration
	//
	IcebergConfiguration interface{} `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
}

