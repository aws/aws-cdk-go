package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// The properties for the base NetworkLoadBalancedEc2Service or NetworkLoadBalancedFargateService service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var containerImage containerImage
//   var hostedZone hostedZone
//   var logDriver logDriver
//   var namespace iNamespace
//   var networkLoadBalancer networkLoadBalancer
//   var role role
//   var secret secret
//   var vpc vpc
//
//   networkLoadBalancedServiceBaseProps := &NetworkLoadBalancedServiceBaseProps{
//   	CapacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			CapacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			Base: jsii.Number(123),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	CircuitBreaker: &DeploymentCircuitBreaker{
//   		Rollback: jsii.Boolean(false),
//   	},
//   	CloudMapOptions: &CloudMapOptions{
//   		CloudMapNamespace: namespace,
//   		Container: containerDefinition,
//   		ContainerPort: jsii.Number(123),
//   		DnsRecordType: awscdk.Aws_servicediscovery.DnsRecordType_A,
//   		DnsTtl: cdk.Duration_Minutes(jsii.Number(30)),
//   		FailureThreshold: jsii.Number(123),
//   		Name: jsii.String("name"),
//   	},
//   	Cluster: cluster,
//   	DeploymentController: &DeploymentController{
//   		Type: awscdk.Aws_ecs.DeploymentControllerType_ECS,
//   	},
//   	DesiredCount: jsii.Number(123),
//   	DomainName: jsii.String("domainName"),
//   	DomainZone: hostedZone,
//   	EnableECSManagedTags: jsii.Boolean(false),
//   	EnableExecuteCommand: jsii.Boolean(false),
//   	HealthCheckGracePeriod: cdk.Duration_*Minutes(jsii.Number(30)),
//   	ListenerPort: jsii.Number(123),
//   	LoadBalancer: networkLoadBalancer,
//   	MaxHealthyPercent: jsii.Number(123),
//   	MinHealthyPercent: jsii.Number(123),
//   	PropagateTags: awscdk.*Aws_ecs.PropagatedTagSource_SERVICE,
//   	PublicLoadBalancer: jsii.Boolean(false),
//   	RecordType: awscdk.Aws_ecs_patterns.NetworkLoadBalancedServiceRecordType_ALIAS,
//   	ServiceName: jsii.String("serviceName"),
//   	TaskImageOptions: &NetworkLoadBalancedTaskImageOptions{
//   		Image: containerImage,
//
//   		// the properties below are optional
//   		ContainerName: jsii.String("containerName"),
//   		ContainerPort: jsii.Number(123),
//   		DockerLabels: map[string]*string{
//   			"dockerLabelsKey": jsii.String("dockerLabels"),
//   		},
//   		EnableLogging: jsii.Boolean(false),
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		ExecutionRole: role,
//   		Family: jsii.String("family"),
//   		LogDriver: logDriver,
//   		Secrets: map[string]*secret{
//   			"secretsKey": secret,
//   		},
//   		TaskRole: role,
//   	},
//   	Vpc: vpc,
//   }
//
type NetworkLoadBalancedServiceBaseProps struct {
	// A list of Capacity Provider strategies used to place a service.
	// Default: - undefined.
	//
	CapacityProviderStrategies *[]*awsecs.CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategies" yaml:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Default: - disabled.
	//
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	// Default: - AWS Cloud Map service discovery is not enabled.
	//
	CloudMapOptions *awsecs.CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Default: - create a new cluster; if both cluster and vpc are omitted, a new VPC will be created for you.
	//
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Default: - Rolling update (ECS).
	//
	DeploymentController *awsecs.DeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1.
	// Default: - If the feature flag, ECS_REMOVE_DEFAULT_DESIRED_COUNT is false, the default is 1;
	// if true, the default is 1 for all new services and uses the existing services desired count
	// when updating an existing service.
	//
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// The domain name for the service, e.g. "api.example.com.".
	// Default: - No domain name.
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Default: - No Route53 hosted domain zone.
	//
	DomainZone awsroute53.IHostedZone `field:"optional" json:"domainZone" yaml:"domainZone"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Default: false.
	//
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Whether ECS Exec should be enabled.
	// Default: - false.
	//
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Default: - defaults to 60 seconds if at least one load balancer is in-use and it is not already set.
	//
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// Listener port of the network load balancer that will serve traffic to the service.
	// Default: 80.
	//
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// The network load balancer that will serve traffic to the service.
	//
	// If the load balancer has been imported, the vpc attribute must be specified
	// in the call to fromNetworkLoadBalancerAttributes().
	//
	// [disable-awslint:ref-via-interface].
	// Default: - a new load balancer will be created.
	//
	LoadBalancer awselasticloadbalancingv2.INetworkLoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Default: - 100 if daemon, otherwise 200.
	//
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Default: - 0 if daemon, otherwise 50.
	//
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Default: - none.
	//
	PropagateTags awsecs.PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Determines whether the Load Balancer will be internet-facing.
	// Default: true.
	//
	PublicLoadBalancer *bool `field:"optional" json:"publicLoadBalancer" yaml:"publicLoadBalancer"`
	// Specifies whether the Route53 record should be a CNAME, an A record using the Alias feature or no record at all.
	//
	// This is useful if you need to work with DNS systems that do not support alias records.
	// Default: NetworkLoadBalancedServiceRecordType.ALIAS
	//
	RecordType NetworkLoadBalancedServiceRecordType `field:"optional" json:"recordType" yaml:"recordType"`
	// The name of the service.
	// Default: - CloudFormation-generated name.
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The properties required to create a new task definition.
	//
	// One of taskImageOptions or taskDefinition must be specified.
	// Default: - none.
	//
	TaskImageOptions *NetworkLoadBalancedTaskImageOptions `field:"optional" json:"taskImageOptions" yaml:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Default: - uses the VPC defined in the cluster or creates a new VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

