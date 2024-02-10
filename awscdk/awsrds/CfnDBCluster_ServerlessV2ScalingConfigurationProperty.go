package awsrds


// The `ServerlessV2ScalingConfiguration` property type specifies the scaling configuration of an Aurora Serverless V2 DB cluster.
//
// For more information, see [Using Amazon Aurora Serverless v2](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html) in the *Amazon Aurora User Guide* .
//
// If you have an Aurora cluster, you must set the `ScalingConfigurationInfo` attribute before you add a DB instance that uses the `db.serverless` DB instance class. For more information, see [Clusters that use Aurora Serverless v2 must have a capacity range specified](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.requirements.html#aurora-serverless-v2.requirements.capacity-range) in the *Amazon Aurora User Guide* .
//
// This property is only supported for Aurora Serverless v2. For Aurora Serverless v1, Use the `ScalingConfiguration` property.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-serverlessv2scalingconfiguration.html
//
type CfnDBCluster_ServerlessV2ScalingConfigurationProperty struct {
	// The maximum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster.
	//
	// You can specify ACU values in half-step increments, such as 40, 40.5, 41, and so on. The largest value that you can use is 128.
	//
	// The maximum capacity must be higher than 0.5 ACUs. For more information, see [Choosing the maximum Aurora Serverless v2 capacity setting for a cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.setting-capacity.html#aurora-serverless-v2.max_capacity_considerations) in the *Amazon Aurora User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-serverlessv2scalingconfiguration.html#cfn-rds-dbcluster-serverlessv2scalingconfiguration-maxcapacity
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster.
	//
	// You can specify ACU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value that you can use is 0.5.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-serverlessv2scalingconfiguration.html#cfn-rds-dbcluster-serverlessv2scalingconfiguration-mincapacity
	//
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}

