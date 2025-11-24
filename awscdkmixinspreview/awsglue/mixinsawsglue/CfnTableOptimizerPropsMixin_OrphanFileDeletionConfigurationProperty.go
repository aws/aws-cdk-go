package mixinsawsglue


// Configuration for removing files that are are not tracked by the Iceberg table metadata, and are older than your configured age limit.
//
// This configuration helps optimize storage usage and costs by automatically cleaning up files that are no longer needed by the table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   orphanFileDeletionConfigurationProperty := &OrphanFileDeletionConfigurationProperty{
//   	IcebergConfiguration: &IcebergConfigurationProperty{
//   		Location: jsii.String("location"),
//   		OrphanFileRetentionPeriodInDays: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-orphanfiledeletionconfiguration.html
//
type CfnTableOptimizerPropsMixin_OrphanFileDeletionConfigurationProperty struct {
	// The `IcebergConfiguration` property helps optimize your Iceberg tables in AWS Glue by allowing you to specify format-specific settings that control how data is stored, compressed, and managed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-orphanfiledeletionconfiguration.html#cfn-glue-tableoptimizer-orphanfiledeletionconfiguration-icebergconfiguration
	//
	IcebergConfiguration interface{} `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
}

