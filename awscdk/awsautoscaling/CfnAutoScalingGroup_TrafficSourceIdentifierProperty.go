package awsautoscaling


// Identifying information for a traffic source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficSourceIdentifierProperty := &TrafficSourceIdentifierProperty{
//   	Identifier: jsii.String("identifier"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-trafficsourceidentifier.html
//
type CfnAutoScalingGroup_TrafficSourceIdentifierProperty struct {
	// Identifies the traffic source.
	//
	// For Application Load Balancers, Gateway Load Balancers, Network Load Balancers, and VPC Lattice, this will be the Amazon Resource Name (ARN) for a target group in this account and Region. For Classic Load Balancers, this will be the name of the Classic Load Balancer in this account and Region.
	//
	// For example:
	//
	// - Application Load Balancer ARN: `arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/my-targets/1234567890123456`
	// - Classic Load Balancer name: `my-classic-load-balancer`
	// - VPC Lattice ARN: `arn:aws:vpc-lattice:us-west-2:123456789012:targetgroup/tg-1234567890123456`
	//
	// To get the ARN of a target group for a Application Load Balancer, Gateway Load Balancer, or Network Load Balancer, or the name of a Classic Load Balancer, use the Elastic Load Balancing [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) and [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operations.
	//
	// To get the ARN of a target group for VPC Lattice, use the VPC Lattice [GetTargetGroup](https://docs.aws.amazon.com/vpc-lattice/latest/APIReference/API_GetTargetGroup.html) API operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-trafficsourceidentifier.html#cfn-autoscaling-autoscalinggroup-trafficsourceidentifier-identifier
	//
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Provides additional context for the value of `Identifier` .
	//
	// The following lists the valid values:
	//
	// - `elb` if `Identifier` is the name of a Classic Load Balancer.
	// - `elbv2` if `Identifier` is the ARN of an Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target group.
	// - `vpc-lattice` if `Identifier` is the ARN of a VPC Lattice target group.
	//
	// Required if the identifier is the name of a Classic Load Balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-trafficsourceidentifier.html#cfn-autoscaling-autoscalinggroup-trafficsourceidentifier-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

