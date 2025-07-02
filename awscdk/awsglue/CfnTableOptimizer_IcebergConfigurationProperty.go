package awsglue


// IcebergConfiguration is a property type within the `AWS::Glue::TableOptimizer` resource in AWS CloudFormation.
//
// This configuration is used when setting up table optimization for Iceberg tables in AWS Glue .
//
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
	// Specifies a directory in which to look for orphan files (defaults to the table's location).
	//
	// You may choose a sub-directory rather than the top-level table location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergconfiguration.html#cfn-glue-tableoptimizer-icebergconfiguration-location
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The specific number of days you want to keep the orphan files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergconfiguration.html#cfn-glue-tableoptimizer-icebergconfiguration-orphanfileretentionperiodindays
	//
	OrphanFileRetentionPeriodInDays *float64 `field:"optional" json:"orphanFileRetentionPeriodInDays" yaml:"orphanFileRetentionPeriodInDays"`
}

