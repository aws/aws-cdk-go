package awselasticloadbalancingv2


// Specifies a target to add to a target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetDescriptionProperty := &TargetDescriptionProperty{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html
//
type CfnTargetGroup_TargetDescriptionProperty struct {
	// The ID of the target.
	//
	// If the target type of the target group is `instance` , specify an instance ID. If the target type is `ip` , specify an IP address. If the target type is `lambda` , specify the ARN of the Lambda function. If the target type is `alb` , specify the ARN of the Application Load Balancer target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html#cfn-elasticloadbalancingv2-targetgroup-targetdescription-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// An Availability Zone or `all` .
	//
	// This determines whether the target receives traffic from the load balancer nodes in the specified Availability Zone or from all enabled Availability Zones for the load balancer.
	//
	// For Application Load Balancer target groups, the specified Availability Zone value is only applicable when cross-zone load balancing is off. Otherwise the parameter is ignored and treated as `all` .
	//
	// This parameter is not supported if the target type of the target group is `instance` or `alb` .
	//
	// If the target type is `ip` and the IP address is in a subnet of the VPC for the target group, the Availability Zone is automatically detected and this parameter is optional. If the IP address is outside the VPC, this parameter is required.
	//
	// For Application Load Balancer target groups with cross-zone load balancing off, if the target type is `ip` and the IP address is outside of the VPC for the target group, this should be an Availability Zone inside the VPC for the target group.
	//
	// If the target type is `lambda` , this parameter is optional and the only supported value is `all` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html#cfn-elasticloadbalancingv2-targetgroup-targetdescription-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The port on which the target is listening.
	//
	// If the target group protocol is GENEVE, the supported port is 6081. If the target type is `alb` , the targeted Application Load Balancer must have at least one listener whose port matches the target group port. This parameter is not used if the target is a Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html#cfn-elasticloadbalancingv2-targetgroup-targetdescription-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

