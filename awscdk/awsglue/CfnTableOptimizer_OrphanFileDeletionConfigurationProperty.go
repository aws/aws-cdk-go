package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnTableOptimizer_OrphanFileDeletionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-orphanfiledeletionconfiguration.html#cfn-glue-tableoptimizer-orphanfiledeletionconfiguration-icebergconfiguration
	//
	IcebergConfiguration interface{} `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
}

