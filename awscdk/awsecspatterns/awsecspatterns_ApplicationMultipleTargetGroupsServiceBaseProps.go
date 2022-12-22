package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
)

// The properties for the base ApplicationMultipleTargetGroupsEc2Service or ApplicationMultipleTargetGroupsFargateService service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var certificate certificate
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var containerImage containerImage
//   var duration duration
//   var hostedZone hostedZone
//   var logDriver logDriver
//   var namespace iNamespace
//   var role role
//   var secret secret
//   var vpc vpc
//
//   applicationMultipleTargetGroupsServiceBaseProps := &applicationMultipleTargetGroupsServiceBaseProps{
//   	cloudMapOptions: &cloudMapOptions{
//   		cloudMapNamespace: namespace,
//   		container: containerDefinition,
//   		containerPort: jsii.Number(123),
//   		dnsRecordType: awscdk.Aws_servicediscovery.dnsRecordType_A,
//   		dnsTtl: duration,
//   		failureThreshold: jsii.Number(123),
//   		name: jsii.String("name"),
//   	},
//   	cluster: cluster,
//   	desiredCount: jsii.Number(123),
//   	enableECSManagedTags: jsii.Boolean(false),
//   	healthCheckGracePeriod: duration,
//   	loadBalancers: []applicationLoadBalancerProps{
//   		&applicationLoadBalancerProps{
//   			listeners: []applicationListenerProps{
//   				&applicationListenerProps{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					certificate: certificate,
//   					port: jsii.Number(123),
//   					protocol: awscdk.Aws_elasticloadbalancingv2.applicationProtocol_HTTP,
//   					sslPolicy: awscdk.*Aws_elasticloadbalancingv2.sslPolicy_RECOMMENDED,
//   				},
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			domainName: jsii.String("domainName"),
//   			domainZone: hostedZone,
//   			publicLoadBalancer: jsii.Boolean(false),
//   		},
//   	},
//   	propagateTags: awscdk.Aws_ecs.propagatedTagSource_SERVICE,
//   	serviceName: jsii.String("serviceName"),
//   	targetGroups: []applicationTargetProps{
//   		&applicationTargetProps{
//   			containerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			hostHeader: jsii.String("hostHeader"),
//   			listener: jsii.String("listener"),
//   			pathPattern: jsii.String("pathPattern"),
//   			priority: jsii.Number(123),
//   			protocol: awscdk.*Aws_ecs.protocol_TCP,
//   		},
//   	},
//   	taskImageOptions: &applicationLoadBalancedTaskImageProps{
//   		image: containerImage,
//
//   		// the properties below are optional
//   		containerName: jsii.String("containerName"),
//   		containerPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		dockerLabels: map[string]*string{
//   			"dockerLabelsKey": jsii.String("dockerLabels"),
//   		},
//   		enableLogging: jsii.Boolean(false),
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		executionRole: role,
//   		family: jsii.String("family"),
//   		logDriver: logDriver,
//   		secrets: map[string]*secret{
//   			"secretsKey": secret,
//   		},
//   		taskRole: role,
//   	},
//   	vpc: vpc,
//   }
//
// Experimental.
type ApplicationMultipleTargetGroupsServiceBaseProps struct {
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Experimental.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The application load balancer that will serve traffic to the service.
	// Experimental.
	LoadBalancers *[]*ApplicationLoadBalancerProps `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The name of the service.
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Properties to specify ALB target groups.
	// Experimental.
	TargetGroups *[]*ApplicationTargetProps `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// The properties required to create a new task definition.
	//
	// Only one of TaskDefinition or TaskImageOptions must be specified.
	// Experimental.
	TaskImageOptions *ApplicationLoadBalancedTaskImageProps `field:"optional" json:"taskImageOptions" yaml:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

