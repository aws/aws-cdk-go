package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icebergRetentionConfigurationProperty := &IcebergRetentionConfigurationProperty{
//   	CleanExpiredFiles: jsii.Boolean(false),
//   	NumberOfSnapshotsToRetain: jsii.Number(123),
//   	SnapshotRetentionPeriodInDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergretentionconfiguration.html
//
type CfnTableOptimizer_IcebergRetentionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergretentionconfiguration.html#cfn-glue-tableoptimizer-icebergretentionconfiguration-cleanexpiredfiles
	//
	CleanExpiredFiles interface{} `field:"optional" json:"cleanExpiredFiles" yaml:"cleanExpiredFiles"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergretentionconfiguration.html#cfn-glue-tableoptimizer-icebergretentionconfiguration-numberofsnapshotstoretain
	//
	NumberOfSnapshotsToRetain *float64 `field:"optional" json:"numberOfSnapshotsToRetain" yaml:"numberOfSnapshotsToRetain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergretentionconfiguration.html#cfn-glue-tableoptimizer-icebergretentionconfiguration-snapshotretentionperiodindays
	//
	SnapshotRetentionPeriodInDays *float64 `field:"optional" json:"snapshotRetentionPeriodInDays" yaml:"snapshotRetentionPeriodInDays"`
}

