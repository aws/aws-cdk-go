package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   compactionConfigurationProperty := &CompactionConfigurationProperty{
//   	IcebergConfiguration: &IcebergCompactionConfigurationProperty{
//   		DeleteFileThreshold: jsii.Number(123),
//   		MinInputFiles: jsii.Number(123),
//   		Strategy: jsii.String("strategy"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-compactionconfiguration.html
//
type CfnTableOptimizerPropsMixin_CompactionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-compactionconfiguration.html#cfn-glue-tableoptimizer-compactionconfiguration-icebergconfiguration
	//
	IcebergConfiguration interface{} `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
}

