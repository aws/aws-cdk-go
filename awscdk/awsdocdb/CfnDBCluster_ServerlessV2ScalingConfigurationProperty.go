package awsdocdb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverlessV2ScalingConfigurationProperty := &ServerlessV2ScalingConfigurationProperty{
//   	MaxCapacity: jsii.Number(123),
//   	MinCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-docdb-dbcluster-serverlessv2scalingconfiguration.html
//
type CfnDBCluster_ServerlessV2ScalingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-docdb-dbcluster-serverlessv2scalingconfiguration.html#cfn-docdb-dbcluster-serverlessv2scalingconfiguration-maxcapacity
	//
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-docdb-dbcluster-serverlessv2scalingconfiguration.html#cfn-docdb-dbcluster-serverlessv2scalingconfiguration-mincapacity
	//
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
}

