package awsrds


// The `ScalingConfiguration` property type specifies the scaling configuration of an Aurora Serverless DB cluster.
//
// For more information, see [Using Amazon Aurora Serverless](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html) in the *Amazon Aurora User Guide* .
//
// Currently, AWS CloudFormation only supports Aurora Serverless v1. AWS CloudFormation doesn't support Aurora Serverless v2.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingConfigurationProperty := &scalingConfigurationProperty{
//   	autoPause: jsii.Boolean(false),
//   	maxCapacity: jsii.Number(123),
//   	minCapacity: jsii.Number(123),
//   	secondsBeforeTimeout: jsii.Number(123),
//   	secondsUntilAutoPause: jsii.Number(123),
//   	timeoutAction: jsii.String("timeoutAction"),
//   }
//
type CfnDBCluster_ScalingConfigurationProperty struct {
	// A value that indicates whether to allow or disallow automatic pause for an Aurora DB cluster in `serverless` DB engine mode.
	//
	// A DB cluster can be paused only when it's idle (it has no connections).
	//
	// > If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot. In this case, the DB cluster is restored when there is a request to connect to it.
	AutoPause interface{} `field:"optional" json:"autoPause" yaml:"autoPause"`
	// The maximum capacity for an Aurora DB cluster in `serverless` DB engine mode.
	//
	// For Aurora MySQL, valid capacity values are `1` , `2` , `4` , `8` , `16` , `32` , `64` , `128` , and `256` .
	//
	// For Aurora PostgreSQL, valid capacity values are `2` , `4` , `8` , `16` , `32` , `64` , `192` , and `384` .
	//
	// The maximum capacity must be greater than or equal to the minimum capacity.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity for an Aurora DB cluster in `serverless` DB engine mode.
	//
	// For Aurora MySQL, valid capacity values are `1` , `2` , `4` , `8` , `16` , `32` , `64` , `128` , and `256` .
	//
	// For Aurora PostgreSQL, valid capacity values are `2` , `4` , `8` , `16` , `32` , `64` , `192` , and `384` .
	//
	// The minimum capacity must be less than or equal to the maximum capacity.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// `CfnDBCluster.ScalingConfigurationProperty.SecondsBeforeTimeout`.
	SecondsBeforeTimeout *float64 `field:"optional" json:"secondsBeforeTimeout" yaml:"secondsBeforeTimeout"`
	// The time, in seconds, before an Aurora DB cluster in `serverless` mode is paused.
	//
	// Specify a value between 300 and 86,400 seconds.
	SecondsUntilAutoPause *float64 `field:"optional" json:"secondsUntilAutoPause" yaml:"secondsUntilAutoPause"`
	// `CfnDBCluster.ScalingConfigurationProperty.TimeoutAction`.
	TimeoutAction *string `field:"optional" json:"timeoutAction" yaml:"timeoutAction"`
}

