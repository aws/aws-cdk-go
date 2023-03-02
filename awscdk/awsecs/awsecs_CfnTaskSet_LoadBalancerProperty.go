package awsecs


// Details on the load balancer or load balancers to use with a task set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerProperty := &loadBalancerProperty{
//   	containerName: jsii.String("containerName"),
//   	containerPort: jsii.Number(123),
//   	loadBalancerName: jsii.String("loadBalancerName"),
//   	targetGroupArn: jsii.String("targetGroupArn"),
//   }
//
type CfnTaskSet_LoadBalancerProperty struct {
	// The name of the container (as it appears in a container definition) to associate with the load balancer.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port on the container to associate with the load balancer.
	//
	// This port must correspond to a `containerPort` in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they're launched on must allow ingress traffic on the `hostPort` of the port mapping.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The name of the load balancer to associate with the Amazon ECS service or task set.
	//
	// A load balancer name is only specified when using a Classic Load Balancer. If you are using an Application Load Balancer or a Network Load Balancer the load balancer name parameter should be omitted.
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set.
	//
	// A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer. If you're using a Classic Load Balancer, omit the target group ARN.
	//
	// For services using the `ECS` deployment controller, you can specify one or multiple target groups. For more information, see [Registering multiple target groups with a service](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// For services using the `CODE_DEPLOY` deployment controller, you're required to define two target groups for the load balancer. For more information, see [Blue/green deployment with CodeDeploy](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > If your service's task definition uses the `awsvpc` network mode, you must choose `ip` as the target type, not `instance` . Do this when creating your target groups because tasks that use the `awsvpc` network mode are associated with an elastic network interface, not an Amazon EC2 instance. This network mode is required for the Fargate launch type.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

