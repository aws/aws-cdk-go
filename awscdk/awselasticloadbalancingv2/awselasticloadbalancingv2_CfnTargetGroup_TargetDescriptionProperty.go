package awselasticloadbalancingv2


// Specifies a target to add to a target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetDescriptionProperty := &targetDescriptionProperty{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	availabilityZone: jsii.String("availabilityZone"),
//   	port: jsii.Number(123),
//   }
//
type CfnTargetGroup_TargetDescriptionProperty struct {
	// The ID of the target.
	//
	// If the target type of the target group is `instance` , specify an instance ID. If the target type is `ip` , specify an IP address. If the target type is `lambda` , specify the ARN of the Lambda function. If the target type is `alb` , specify the ARN of the Application Load Balancer target.
	Id *string `field:"required" json:"id" yaml:"id"`
	// An Availability Zone or `all` .
	//
	// This determines whether the target receives traffic from the load balancer nodes in the specified Availability Zone or from all enabled Availability Zones for the load balancer.
	//
	// This parameter is not supported if the target type of the target group is `instance` or `alb` .
	//
	// If the target type is `ip` and the IP address is in a subnet of the VPC for the target group, the Availability Zone is automatically detected and this parameter is optional. If the IP address is outside the VPC, this parameter is required.
	//
	// With an Application Load Balancer, if the target type is `ip` and the IP address is outside the VPC for the target group, the only supported value is `all` .
	//
	// If the target type is `lambda` , this parameter is optional and the only supported value is `all` .
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The port on which the target is listening.
	//
	// If the target group protocol is GENEVE, the supported port is 6081. If the target type is `alb` , the targeted Application Load Balancer must have at least one listener whose port matches the target group port. Not used if the target is a Lambda function.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

