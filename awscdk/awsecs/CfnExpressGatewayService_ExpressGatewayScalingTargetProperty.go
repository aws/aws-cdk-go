package awsecs


// Defines the auto-scaling configuration for an Express service.
//
// This determines how the service automatically adjusts the number of running tasks based on demand metrics such as CPU utilization, memory utilization, or request count per target.
//
// Auto-scaling helps ensure your application can handle varying levels of traffic while optimizing costs by scaling down during low-demand periods. You can specify the minimum and maximum number of tasks, the scaling metric, and the target value for that metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expressGatewayScalingTargetProperty := &ExpressGatewayScalingTargetProperty{
//   	AutoScalingMetric: jsii.String("autoScalingMetric"),
//   	AutoScalingTargetValue: jsii.Number(123),
//   	MaxTaskCount: jsii.Number(123),
//   	MinTaskCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayscalingtarget.html
//
type CfnExpressGatewayService_ExpressGatewayScalingTargetProperty struct {
	// The metric used for auto-scaling decisions.
	//
	// The default metric used for an Express service is `CPUUtilization` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayscalingtarget.html#cfn-ecs-expressgatewayservice-expressgatewayscalingtarget-autoscalingmetric
	//
	// Default: - "AVERAGE_CPU".
	//
	AutoScalingMetric *string `field:"optional" json:"autoScalingMetric" yaml:"autoScalingMetric"`
	// The target value for the auto-scaling metric.
	//
	// The default value for an Express service is 60.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayscalingtarget.html#cfn-ecs-expressgatewayservice-expressgatewayscalingtarget-autoscalingtargetvalue
	//
	// Default: - 60.
	//
	AutoScalingTargetValue *float64 `field:"optional" json:"autoScalingTargetValue" yaml:"autoScalingTargetValue"`
	// The maximum number of tasks to run in the Express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayscalingtarget.html#cfn-ecs-expressgatewayservice-expressgatewayscalingtarget-maxtaskcount
	//
	// Default: - 1.
	//
	MaxTaskCount *float64 `field:"optional" json:"maxTaskCount" yaml:"maxTaskCount"`
	// The minimum number of tasks to run in the Express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayscalingtarget.html#cfn-ecs-expressgatewayservice-expressgatewayscalingtarget-mintaskcount
	//
	// Default: - 1.
	//
	MinTaskCount *float64 `field:"optional" json:"minTaskCount" yaml:"minTaskCount"`
}

