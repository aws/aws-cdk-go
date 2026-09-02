package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icebergCompactionConfigurationProperty := &IcebergCompactionConfigurationProperty{
//   	DeleteFileThreshold: jsii.Number(123),
//   	MinInputFiles: jsii.Number(123),
//   	Strategy: jsii.String("strategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergcompactionconfiguration.html
//
type CfnTableOptimizer_IcebergCompactionConfigurationProperty struct {
	// The minimum number of deletes in a data file to make it eligible for compaction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergcompactionconfiguration.html#cfn-glue-tableoptimizer-icebergcompactionconfiguration-deletefilethreshold
	//
	DeleteFileThreshold *float64 `field:"optional" json:"deleteFileThreshold" yaml:"deleteFileThreshold"`
	// The minimum number of input files before compaction is triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergcompactionconfiguration.html#cfn-glue-tableoptimizer-icebergcompactionconfiguration-mininputfiles
	//
	MinInputFiles *float64 `field:"optional" json:"minInputFiles" yaml:"minInputFiles"`
	// The compaction strategy to use.
	//
	// Valid values are binpack, sort, and z-order.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-icebergcompactionconfiguration.html#cfn-glue-tableoptimizer-icebergcompactionconfiguration-strategy
	//
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

