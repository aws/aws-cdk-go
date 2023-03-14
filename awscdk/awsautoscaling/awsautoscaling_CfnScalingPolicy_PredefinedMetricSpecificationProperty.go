package awsautoscaling


// Contains predefined metric specification information for a target tracking scaling policy for Amazon EC2 Auto Scaling.
//
// `PredefinedMetricSpecification` is a property of the [AWS::AutoScaling::ScalingPolicy TargetTrackingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predefinedMetricSpecificationProperty := &predefinedMetricSpecificationProperty{
//   	predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   	// the properties below are optional
//   	resourceLabel: jsii.String("resourceLabel"),
//   }
//
type CfnScalingPolicy_PredefinedMetricSpecificationProperty struct {
	// The metric type. The following predefined metrics are available:.
	//
	// - `ASGAverageCPUUtilization` - Average CPU utilization of the Auto Scaling group.
	// - `ASGAverageNetworkIn` - Average number of bytes received on all network interfaces by the Auto Scaling group.
	// - `ASGAverageNetworkOut` - Average number of bytes sent out on all network interfaces by the Auto Scaling group.
	// - `ALBRequestCountPerTarget` - Average Application Load Balancer request count per target for your Auto Scaling group.
	PredefinedMetricType *string `field:"required" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// A label that uniquely identifies a specific Application Load Balancer target group from which to determine the average request count served by your Auto Scaling group.
	//
	// You can't specify a resource label unless the target group is attached to the Auto Scaling group.
	//
	// You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). The format of the resource label is:
	//
	// `app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff` .
	//
	// Where:
	//
	// - app/<load-balancer-name>/<load-balancer-id> is the final portion of the load balancer ARN
	// - targetgroup/<target-group-name>/<target-group-id> is the final portion of the target group ARN.
	//
	// To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operation. To find the ARN for the target group, use the [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) API operation.
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

