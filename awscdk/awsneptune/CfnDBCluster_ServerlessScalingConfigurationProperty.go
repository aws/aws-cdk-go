package awsneptune


// Contains the scaling configuration of an Neptune Serverless DB cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverlessScalingConfigurationProperty := &ServerlessScalingConfigurationProperty{
//   	MaxCapacity: jsii.Number(123),
//   	MinCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptune-dbcluster-serverlessscalingconfiguration.html
//
type CfnDBCluster_ServerlessScalingConfigurationProperty struct {
	// The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster.
	//
	// You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptune-dbcluster-serverlessscalingconfiguration.html#cfn-neptune-dbcluster-serverlessscalingconfiguration-maxcapacity
	//
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster.
	//
	// You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptune-dbcluster-serverlessscalingconfiguration.html#cfn-neptune-dbcluster-serverlessscalingconfiguration-mincapacity
	//
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
}

