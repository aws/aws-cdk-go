package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icebergConfigurationProperty := &IcebergConfigurationProperty{
//   	Location: jsii.String("location"),
//   	OrphanFileRetentionPeriodInDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergconfiguration.html
//
type CfnTableOptimizer_IcebergConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergconfiguration.html#cfn-glue-tableoptimizer-icebergconfiguration-location
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergconfiguration.html#cfn-glue-tableoptimizer-icebergconfiguration-orphanfileretentionperiodindays
	//
	OrphanFileRetentionPeriodInDays *float64 `field:"optional" json:"orphanFileRetentionPeriodInDays" yaml:"orphanFileRetentionPeriodInDays"`
}

