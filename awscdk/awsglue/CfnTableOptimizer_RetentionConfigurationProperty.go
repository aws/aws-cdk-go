package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retentionConfigurationProperty := &RetentionConfigurationProperty{
//   	IcebergConfiguration: &IcebergConfigurationProperty{
//   		Location: jsii.String("location"),
//   		OrphanFileRetentionPeriodInDays: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-retentionconfiguration.html
//
type CfnTableOptimizer_RetentionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-retentionconfiguration.html#cfn-glue-tableoptimizer-retentionconfiguration-icebergconfiguration
	//
	IcebergConfiguration interface{} `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
}

