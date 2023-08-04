package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// The properties for the base NetworkMultipleTargetGroupsEc2Service or NetworkMultipleTargetGroupsFargateService service.
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
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var containerImage containerImage
//   var hostedZone hostedZone
//   var logDriver logDriver
//   var namespace iNamespace
//   var role role
//   var secret secret
//   var vpc vpc
//
//   networkMultipleTargetGroupsServiceBaseProps := &NetworkMultipleTargetGroupsServiceBaseProps{
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
//   	DesiredCount: jsii.Number(123),
//   	EnableECSManagedTags: jsii.Boolean(false),
//   	EnableExecuteCommand: jsii.Boolean(false),
//   	HealthCheckGracePeriod: cdk.Duration_*Minutes(jsii.Number(30)),
//   	LoadBalancers: []networkLoadBalancerProps{
//   		&networkLoadBalancerProps{
//   			Listeners: []networkListenerProps{
//   				&networkListenerProps{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Port: jsii.Number(123),
//   				},
//   			},
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			DomainName: jsii.String("domainName"),
//   			DomainZone: hostedZone,
//   			PublicLoadBalancer: jsii.Boolean(false),
//   		},
//   	},
//   	PropagateTags: awscdk.Aws_ecs.PropagatedTagSource_SERVICE,
//   	ServiceName: jsii.String("serviceName"),
//   	TargetGroups: []networkTargetProps{
//   		&networkTargetProps{
//   			ContainerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			Listener: jsii.String("listener"),
//   		},
//   	},
//   	TaskImageOptions: &NetworkLoadBalancedTaskImageProps{
//   		Image: containerImage,
//
//   		// the properties below are optional
//   		ContainerName: jsii.String("containerName"),
//   		ContainerPorts: []*f64{
//   			jsii.Number(123),
//   		},
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
type NetworkMultipleTargetGroupsServiceBaseProps struct {
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
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1.
	// Default: - If the feature flag, ECS_REMOVE_DEFAULT_DESIRED_COUNT is false, the default is 1;
	// if true, the default is 1 for all new services and uses the existing services desired count
	// when updating an existing service.
	//
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
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
	// The network load balancer that will serve traffic to the service.
	// Default: - a new load balancer with a listener will be created.
	//
	LoadBalancers *[]*NetworkLoadBalancerProps `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Default: - none.
	//
	PropagateTags awsecs.PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Name of the service.
	// Default: - CloudFormation-generated name.
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Properties to specify NLB target groups.
	// Default: - default portMapping registered as target group and attached to the first defined listener.
	//
	TargetGroups *[]*NetworkTargetProps `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// The properties required to create a new task definition.
	//
	// Only one of TaskDefinition or TaskImageOptions must be specified.
	// Default: - none.
	//
	TaskImageOptions *NetworkLoadBalancedTaskImageProps `field:"optional" json:"taskImageOptions" yaml:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Default: - uses the VPC defined in the cluster or creates a new VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

