package awsecs


// Properties for defining a `CfnTaskSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTaskSetProps := &CfnTaskSetProps{
//   	Cluster: jsii.String("cluster"),
//   	Service: jsii.String("service"),
//   	TaskDefinition: jsii.String("taskDefinition"),
//
//   	// the properties below are optional
//   	ExternalId: jsii.String("externalId"),
//   	LaunchType: jsii.String("launchType"),
//   	LoadBalancers: []interface{}{
//   		&LoadBalancerProperty{
//   			ContainerName: jsii.String("containerName"),
//   			ContainerPort: jsii.Number(123),
//   			LoadBalancerName: jsii.String("loadBalancerName"),
//   			TargetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		AwsVpcConfiguration: &AwsVpcConfigurationProperty{
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//
//   			// the properties below are optional
//   			AssignPublicIp: jsii.String("assignPublicIp"),
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   		},
//   	},
//   	PlatformVersion: jsii.String("platformVersion"),
//   	Scale: &ScaleProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   	ServiceRegistries: []interface{}{
//   		&ServiceRegistryProperty{
//   			ContainerName: jsii.String("containerName"),
//   			ContainerPort: jsii.Number(123),
//   			Port: jsii.Number(123),
//   			RegistryArn: jsii.String("registryArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html
//
type CfnTaskSetProps struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-cluster
	//
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-service
	//
	Service *string `field:"required" json:"service" yaml:"service"`
	// The task definition for the tasks in the task set to use.
	//
	// If a revision isn't specified, the latest `ACTIVE` revision is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-taskdefinition
	//
	TaskDefinition *string `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// An optional non-unique tag that identifies this task set in external systems.
	//
	// If the task set is associated with a service discovery registry, the tasks in this task set will have the `ECS_TASK_SET_EXTERNAL_ID` AWS Cloud Map attribute set to the provided value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The launch type that new tasks in the task set uses.
	//
	// For more information, see [Amazon ECS launch types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// If a `launchType` is specified, the `capacityProviderStrategy` parameter must be omitted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-launchtype
	//
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// A load balancer object representing the load balancer to use with the task set.
	//
	// The supported load balancer types are either an Application Load Balancer or a Network Load Balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-loadbalancers
	//
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// The network configuration for the task set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The platform version that the tasks in the task set uses.
	//
	// A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the `LATEST` platform version is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-platformversion
	//
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// A floating-point percentage of your desired number of tasks to place and keep running in the task set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-scale
	//
	Scale interface{} `field:"optional" json:"scale" yaml:"scale"`
	// The details of the service discovery registries to assign to this task set.
	//
	// For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html#cfn-ecs-taskset-serviceregistries
	//
	ServiceRegistries interface{} `field:"optional" json:"serviceRegistries" yaml:"serviceRegistries"`
}

