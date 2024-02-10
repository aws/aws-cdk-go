package awsrds


// The `ScalingConfiguration` property type specifies the scaling configuration of an Aurora Serverless DB cluster.
//
// For more information, see [Using Amazon Aurora Serverless](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html) in the *Amazon Aurora User Guide* .
//
// This property is only supported for Aurora Serverless v1. For Aurora Serverless v2, Use the `ServerlessV2ScalingConfiguration` property.
//
// Valid for: Aurora DB clusters only.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingConfigurationProperty := &ScalingConfigurationProperty{
//   	AutoPause: jsii.Boolean(false),
//   	MaxCapacity: jsii.Number(123),
//   	MinCapacity: jsii.Number(123),
//   	SecondsBeforeTimeout: jsii.Number(123),
//   	SecondsUntilAutoPause: jsii.Number(123),
//   	TimeoutAction: jsii.String("timeoutAction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-scalingconfiguration.html
//
type CfnDBCluster_ScalingConfigurationProperty struct {
	// Indicates whether to allow or disallow automatic pause for an Aurora DB cluster in `serverless` DB engine mode.
	//
	// A DB cluster can be paused only when it's idle (it has no connections).
	//
	// > If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot. In this case, the DB cluster is restored when there is a request to connect to it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-scalingconfiguration.html#cfn-rds-dbcluster-scalingconfiguration-autopause
	//
	AutoPause interface{} `field:"optional" json:"autoPause" yaml:"autoPause"`
	// The maximum capacity for an Aurora DB cluster in `serverless` DB engine mode.
	//
	// For Aurora MySQL, valid capacity values are `1` , `2` , `4` , `8` , `16` , `32` , `64` , `128` , and `256` .
	//
	// For Aurora PostgreSQL, valid capacity values are `2` , `4` , `8` , `16` , `32` , `64` , `192` , and `384` .
	//
	// The maximum capacity must be greater than or equal to the minimum capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-scalingconfiguration.html#cfn-rds-dbcluster-scalingconfiguration-maxcapacity
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity for an Aurora DB cluster in `serverless` DB engine mode.
	//
	// For Aurora MySQL, valid capacity values are `1` , `2` , `4` , `8` , `16` , `32` , `64` , `128` , and `256` .
	//
	// For Aurora PostgreSQL, valid capacity values are `2` , `4` , `8` , `16` , `32` , `64` , `192` , and `384` .
	//
	// The minimum capacity must be less than or equal to the maximum capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-scalingconfiguration.html#cfn-rds-dbcluster-scalingconfiguration-mincapacity
	//
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// The amount of time, in seconds, that Aurora Serverless v1 tries to find a scaling point to perform seamless scaling before enforcing the timeout action.
	//
	// The default is 300.
	//
	// Specify a value between 60 and 600 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-scalingconfiguration.html#cfn-rds-dbcluster-scalingconfiguration-secondsbeforetimeout
	//
	SecondsBeforeTimeout *float64 `field:"optional" json:"secondsBeforeTimeout" yaml:"secondsBeforeTimeout"`
	// The time, in seconds, before an Aurora DB cluster in `serverless` mode is paused.
	//
	// Specify a value between 300 and 86,400 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-scalingconfiguration.html#cfn-rds-dbcluster-scalingconfiguration-secondsuntilautopause
	//
	SecondsUntilAutoPause *float64 `field:"optional" json:"secondsUntilAutoPause" yaml:"secondsUntilAutoPause"`
	// The action to take when the timeout is reached, either `ForceApplyCapacityChange` or `RollbackCapacityChange` .
	//
	// `ForceApplyCapacityChange` sets the capacity to the specified value as soon as possible.
	//
	// `RollbackCapacityChange` , the default, ignores the capacity change if a scaling point isn't found in the timeout period.
	//
	// > If you specify `ForceApplyCapacityChange` , connections that prevent Aurora Serverless v1 from finding a scaling point might be dropped.
	//
	// For more information, see [Autoscaling for Aurora Serverless v1](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.auto-scaling) in the *Amazon Aurora User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-scalingconfiguration.html#cfn-rds-dbcluster-scalingconfiguration-timeoutaction
	//
	TimeoutAction *string `field:"optional" json:"timeoutAction" yaml:"timeoutAction"`
}

