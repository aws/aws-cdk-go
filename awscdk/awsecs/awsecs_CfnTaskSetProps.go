package awsecs


// Properties for defining a `CfnTaskSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTaskSetProps := &cfnTaskSetProps{
//   	cluster: jsii.String("cluster"),
//   	service: jsii.String("service"),
//   	taskDefinition: jsii.String("taskDefinition"),
//
//   	// the properties below are optional
//   	externalId: jsii.String("externalId"),
//   	launchType: jsii.String("launchType"),
//   	loadBalancers: []interface{}{
//   		&loadBalancerProperty{
//   			containerName: jsii.String("containerName"),
//   			containerPort: jsii.Number(123),
//   			loadBalancerName: jsii.String("loadBalancerName"),
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	networkConfiguration: &networkConfigurationProperty{
//   		awsVpcConfiguration: &awsVpcConfigurationProperty{
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//
//   			// the properties below are optional
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   			securityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   		},
//   	},
//   	platformVersion: jsii.String("platformVersion"),
//   	scale: &scaleProperty{
//   		unit: jsii.String("unit"),
//   		value: jsii.Number(123),
//   	},
//   	serviceRegistries: []interface{}{
//   		&serviceRegistryProperty{
//   			containerName: jsii.String("containerName"),
//   			containerPort: jsii.Number(123),
//   			port: jsii.Number(123),
//   			registryArn: jsii.String("registryArn"),
//   		},
//   	},
//   }
//
type CfnTaskSetProps struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
	Service *string `field:"required" json:"service" yaml:"service"`
	// The task definition for the tasks in the task set to use.
	TaskDefinition *string `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// An optional non-unique tag that identifies this task set in external systems.
	//
	// If the task set is associated with a service discovery registry, the tasks in this task set will have the `ECS_TASK_SET_EXTERNAL_ID` AWS Cloud Map attribute set to the provided value.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The launch type that new tasks in the task set uses.
	//
	// For more information, see [Amazon ECS launch types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// If a `launchType` is specified, the `capacityProviderStrategy` parameter must be omitted.
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// A load balancer object representing the load balancer to use with the task set.
	//
	// The supported load balancer types are either an Application Load Balancer or a Network Load Balancer.
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// The network configuration for the task set.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The platform version that the tasks in the task set uses.
	//
	// A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the `LATEST` platform version is used.
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// A floating-point percentage of your desired number of tasks to place and keep running in the task set.
	Scale interface{} `field:"optional" json:"scale" yaml:"scale"`
	// The details of the service discovery registries to assign to this task set.
	//
	// For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html) .
	ServiceRegistries interface{} `field:"optional" json:"serviceRegistries" yaml:"serviceRegistries"`
}

