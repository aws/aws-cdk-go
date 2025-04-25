package awscassandra


// The AWS Region specific settings of a multi-Region table.
//
// For a multi-Region table, you can configure the table's read capacity differently per AWS Region. You can do this by configuring the following parameters.
//
// - `region` : The Region where these settings are applied. (Required)
// - `readCapacityUnits` : The provisioned read capacity units. (Optional)
// - `readCapacityAutoScaling` : The read capacity auto scaling settings for the table. (Optional)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicaSpecificationProperty := &ReplicaSpecificationProperty{
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	ReadCapacityAutoScaling: &AutoScalingSettingProperty{
//   		AutoScalingDisabled: jsii.Boolean(false),
//   		MaximumUnits: jsii.Number(123),
//   		MinimumUnits: jsii.Number(123),
//   		ScalingPolicy: &ScalingPolicyProperty{
//   			TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   				TargetValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				DisableScaleIn: jsii.Boolean(false),
//   				ScaleInCooldown: jsii.Number(123),
//   				ScaleOutCooldown: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ReadCapacityUnits: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-replicaspecification.html
//
type CfnTable_ReplicaSpecificationProperty struct {
	// The AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-replicaspecification.html#cfn-cassandra-table-replicaspecification-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	// The read capacity auto scaling settings for the multi-Region table in the specified AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-replicaspecification.html#cfn-cassandra-table-replicaspecification-readcapacityautoscaling
	//
	ReadCapacityAutoScaling interface{} `field:"optional" json:"readCapacityAutoScaling" yaml:"readCapacityAutoScaling"`
	// The provisioned read capacity units for the multi-Region table in the specified AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-replicaspecification.html#cfn-cassandra-table-replicaspecification-readcapacityunits
	//
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
}

