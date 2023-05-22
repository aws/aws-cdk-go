package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
)

// The properties for the base ApplicationLoadBalancedEc2Service or ApplicationLoadBalancedFargateService service.
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
//   var applicationLoadBalancer applicationLoadBalancer
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
//   applicationLoadBalancedServiceBaseProps := &ApplicationLoadBalancedServiceBaseProps{
//   	Certificate: certificate,
//   	CircuitBreaker: &DeploymentCircuitBreaker{
//   		Rollback: jsii.Boolean(false),
//   	},
//   	CloudMapOptions: &CloudMapOptions{
//   		CloudMapNamespace: namespace,
//   		Container: containerDefinition,
//   		ContainerPort: jsii.Number(123),
//   		DnsRecordType: awscdk.Aws_servicediscovery.DnsRecordType_A,
//   		DnsTtl: duration,
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
//   	HealthCheckGracePeriod: duration,
//   	ListenerPort: jsii.Number(123),
//   	LoadBalancer: applicationLoadBalancer,
//   	LoadBalancerName: jsii.String("loadBalancerName"),
//   	MaxHealthyPercent: jsii.Number(123),
//   	MinHealthyPercent: jsii.Number(123),
//   	OpenListener: jsii.Boolean(false),
//   	PropagateTags: awscdk.*Aws_ecs.PropagatedTagSource_SERVICE,
//   	Protocol: awscdk.Aws_elasticloadbalancingv2.ApplicationProtocol_HTTP,
//   	ProtocolVersion: awscdk.*Aws_elasticloadbalancingv2.ApplicationProtocolVersion_GRPC,
//   	PublicLoadBalancer: jsii.Boolean(false),
//   	RecordType: awscdk.Aws_ecs_patterns.ApplicationLoadBalancedServiceRecordType_ALIAS,
//   	RedirectHTTP: jsii.Boolean(false),
//   	ServiceName: jsii.String("serviceName"),
//   	SslPolicy: awscdk.*Aws_elasticloadbalancingv2.SslPolicy_RECOMMENDED,
//   	TargetProtocol: awscdk.*Aws_elasticloadbalancingv2.ApplicationProtocol_HTTP,
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
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
// Experimental.
type ApplicationLoadBalancedServiceBaseProps struct {
	// Certificate Manager certificate to associate with the load balancer.
	//
	// Setting this option will set the load balancer protocol to HTTPS.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Experimental.
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	// Experimental.
	CloudMapOptions *awsecs.CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Experimental.
	DeploymentController *awsecs.DeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The minimum value is 1.
	// Experimental.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// The domain name for the service, e.g. "api.example.com.".
	// Experimental.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Experimental.
	DomainZone awsroute53.IHostedZone `field:"optional" json:"domainZone" yaml:"domainZone"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	// Experimental.
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// Listener port of the application load balancer that will serve traffic to the service.
	// Experimental.
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// The application load balancer that will serve traffic to the service.
	//
	// The VPC attribute of a load balancer must be specified for it to be used
	// to create a new service with this pattern.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	LoadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// Name of the load balancer.
	// Experimental.
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Experimental.
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Experimental.
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Determines whether or not the Security Group for the Load Balancer's Listener will be open to all traffic by default.
	// Experimental.
	OpenListener *bool `field:"optional" json:"openListener" yaml:"openListener"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The protocol for connections from clients to the load balancer.
	//
	// The load balancer port is determined from the protocol (port 80 for
	// HTTP, port 443 for HTTPS).  A domain name and zone must be also be
	// specified if using HTTPS.
	// Experimental.
	Protocol awselasticloadbalancingv2.ApplicationProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol version to use.
	// Experimental.
	ProtocolVersion awselasticloadbalancingv2.ApplicationProtocolVersion `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// Determines whether the Load Balancer will be internet-facing.
	// Experimental.
	PublicLoadBalancer *bool `field:"optional" json:"publicLoadBalancer" yaml:"publicLoadBalancer"`
	// Specifies whether the Route53 record should be a CNAME, an A record using the Alias feature or no record at all.
	//
	// This is useful if you need to work with DNS systems that do not support alias records.
	// Experimental.
	RecordType ApplicationLoadBalancedServiceRecordType `field:"optional" json:"recordType" yaml:"recordType"`
	// Specifies whether the load balancer should redirect traffic on port 80 to port 443 to support HTTP->HTTPS redirects This is only valid if the protocol of the ALB is HTTPS.
	// Experimental.
	RedirectHTTP *bool `field:"optional" json:"redirectHTTP" yaml:"redirectHTTP"`
	// The name of the service.
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The security policy that defines which ciphers and protocols are supported by the ALB Listener.
	// Experimental.
	SslPolicy awselasticloadbalancingv2.SslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// The protocol for connections from the load balancer to the ECS tasks.
	//
	// The default target port is determined from the protocol (port 80 for
	// HTTP, port 443 for HTTPS).
	// Experimental.
	TargetProtocol awselasticloadbalancingv2.ApplicationProtocol `field:"optional" json:"targetProtocol" yaml:"targetProtocol"`
	// The properties required to create a new task definition.
	//
	// TaskDefinition or TaskImageOptions must be specified, but not both.
	// Experimental.
	TaskImageOptions *ApplicationLoadBalancedTaskImageOptions `field:"optional" json:"taskImageOptions" yaml:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

