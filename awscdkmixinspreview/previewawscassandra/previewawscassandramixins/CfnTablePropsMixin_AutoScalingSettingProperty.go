package previewawscassandramixins


// The optional auto scaling settings for a table with provisioned throughput capacity.
//
// To turn on auto scaling for a table in `throughputMode:PROVISIONED` , you must specify the following parameters.
//
// Configure the minimum and maximum capacity units. The auto scaling policy ensures that capacity never goes below the minimum or above the maximum range.
//
// - `minimumUnits` : The minimum level of throughput the table should always be ready to support. The value must be between 1 and the max throughput per second quota for your account (40,000 by default).
// - `maximumUnits` : The maximum level of throughput the table should always be ready to support. The value must be between 1 and the max throughput per second quota for your account (40,000 by default).
// - `scalingPolicy` : Amazon Keyspaces supports the `target tracking` scaling policy. The auto scaling target is a percentage of the provisioned capacity of the table.
//
// For more information, see [Managing throughput capacity automatically with Amazon Keyspaces auto scaling](https://docs.aws.amazon.com/keyspaces/latest/devguide/autoscaling.html) in the *Amazon Keyspaces Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autoScalingSettingProperty := &AutoScalingSettingProperty{
//   	AutoScalingDisabled: jsii.Boolean(false),
//   	MaximumUnits: jsii.Number(123),
//   	MinimumUnits: jsii.Number(123),
//   	ScalingPolicy: &ScalingPolicyProperty{
//   		TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   			DisableScaleIn: jsii.Boolean(false),
//   			ScaleInCooldown: jsii.Number(123),
//   			ScaleOutCooldown: jsii.Number(123),
//   			TargetValue: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingsetting.html
//
type CfnTablePropsMixin_AutoScalingSettingProperty struct {
	// This optional parameter enables auto scaling for the table if set to `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingsetting.html#cfn-cassandra-table-autoscalingsetting-autoscalingdisabled
	//
	// Default: - false.
	//
	AutoScalingDisabled interface{} `field:"optional" json:"autoScalingDisabled" yaml:"autoScalingDisabled"`
	// Manage costs by specifying the maximum amount of throughput to provision.
	//
	// The value must be between 1 and the max throughput per second quota for your account (40,000 by default).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingsetting.html#cfn-cassandra-table-autoscalingsetting-maximumunits
	//
	MaximumUnits *float64 `field:"optional" json:"maximumUnits" yaml:"maximumUnits"`
	// The minimum level of throughput the table should always be ready to support.
	//
	// The value must be between 1 and the max throughput per second quota for your account (40,000 by default).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingsetting.html#cfn-cassandra-table-autoscalingsetting-minimumunits
	//
	MinimumUnits *float64 `field:"optional" json:"minimumUnits" yaml:"minimumUnits"`
	// Amazon Keyspaces supports the `target tracking` auto scaling policy.
	//
	// With this policy, Amazon Keyspaces auto scaling ensures that the table's ratio of consumed to provisioned capacity stays at or near the target value that you specify. You define the target value as a percentage between 20 and 90.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingsetting.html#cfn-cassandra-table-autoscalingsetting-scalingpolicy
	//
	ScalingPolicy interface{} `field:"optional" json:"scalingPolicy" yaml:"scalingPolicy"`
}

