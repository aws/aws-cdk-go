package awsecs


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayscalingtarget.html#cfn-ecs-expressgatewayservice-expressgatewayscalingtarget-autoscalingmetric
	//
	// Default: - "AVERAGE_CPU".
	//
	AutoScalingMetric *string `field:"optional" json:"autoScalingMetric" yaml:"autoScalingMetric"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayscalingtarget.html#cfn-ecs-expressgatewayservice-expressgatewayscalingtarget-autoscalingtargetvalue
	//
	// Default: - 60.
	//
	AutoScalingTargetValue *float64 `field:"optional" json:"autoScalingTargetValue" yaml:"autoScalingTargetValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayscalingtarget.html#cfn-ecs-expressgatewayservice-expressgatewayscalingtarget-maxtaskcount
	//
	// Default: - 1.
	//
	MaxTaskCount *float64 `field:"optional" json:"maxTaskCount" yaml:"maxTaskCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayscalingtarget.html#cfn-ecs-expressgatewayservice-expressgatewayscalingtarget-mintaskcount
	//
	// Default: - 1.
	//
	MinTaskCount *float64 `field:"optional" json:"minTaskCount" yaml:"minTaskCount"`
}

