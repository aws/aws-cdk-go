package awsecs


// The load balancer configuration to use with a service or task set.
//
// When you add, update, or remove a load balancer configuration, Amazon ECS starts a new deployment with the updated Elastic Load Balancing configuration. This causes tasks to register to and deregister from load balancers.
//
// We recommend that you verify this on a test environment before you update the Elastic Load Balancing configuration.
//
// A service-linked role is required for services that use multiple target groups. For more information, see [Using service-linked roles](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerProperty := &LoadBalancerProperty{
//   	ContainerName: jsii.String("containerName"),
//   	ContainerPort: jsii.Number(123),
//   	TargetGroupArn: jsii.String("targetGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-loadbalancer.html
//
type CfnTaskSet_LoadBalancerProperty struct {
	// The name of the container (as it appears in a container definition) to associate with the load balancer.
	//
	// You need to specify the container name when configuring the target group for an Amazon ECS load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-loadbalancer.html#cfn-ecs-taskset-loadbalancer-containername
	//
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port on the container to associate with the load balancer.
	//
	// This port must correspond to a `containerPort` in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they're launched on must allow ingress traffic on the `hostPort` of the port mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-loadbalancer.html#cfn-ecs-taskset-loadbalancer-containerport
	//
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set.
	//
	// A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer.
	//
	// For services using the `ECS` deployment controller, you can specify one or multiple target groups. For more information, see [Registering multiple target groups with a service](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// For services using the `CODE_DEPLOY` deployment controller, you're required to define two target groups for the load balancer. For more information, see [Blue/green deployment with CodeDeploy](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > If your service's task definition uses the `awsvpc` network mode, you must choose `ip` as the target type, not `instance` . Do this when creating your target groups because tasks that use the `awsvpc` network mode are associated with an elastic network interface, not an Amazon EC2 instance. This network mode is required for the Fargate launch type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-loadbalancer.html#cfn-ecs-taskset-loadbalancer-targetgrouparn
	//
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

