package awselasticloadbalancingv2


// Backend protocol for network load balancers and health checks.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	MaxAzs: jsii.Number(1),
//   })
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationMultipleTargetGroupsFargateService(this, jsii.String("myService"), &ApplicationMultipleTargetGroupsFargateServiceProps{
//   	Cluster: ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
//   		Vpc: *Vpc,
//   	}),
//   	MemoryLimitMiB: jsii.Number(256),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	EnableExecuteCommand: jsii.Boolean(true),
//   	LoadBalancers: []applicationLoadBalancerProps{
//   		&applicationLoadBalancerProps{
//   			Name: jsii.String("lb"),
//   			IdleTimeout: awscdk.Duration_Seconds(jsii.Number(400)),
//   			DomainName: jsii.String("api.example.com"),
//   			DomainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
//   				ZoneName: jsii.String("example.com"),
//   			}),
//   			Listeners: []applicationListenerProps{
//   				&applicationListenerProps{
//   					Name: jsii.String("listener"),
//   					Protocol: awscdk.ApplicationProtocol_HTTPS,
//   					Certificate: awscdk.Certificate_FromCertificateArn(this, jsii.String("Cert"), jsii.String("helloworld")),
//   					SslPolicy: awscdk.SslPolicy_TLS12_EXT,
//   				},
//   			},
//   		},
//   		&applicationLoadBalancerProps{
//   			Name: jsii.String("lb2"),
//   			IdleTimeout: awscdk.Duration_*Seconds(jsii.Number(120)),
//   			DomainName: jsii.String("frontend.com"),
//   			DomainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
//   				ZoneName: jsii.String("frontend.com"),
//   			}),
//   			Listeners: []*applicationListenerProps{
//   				&applicationListenerProps{
//   					Name: jsii.String("listener2"),
//   					Protocol: awscdk.ApplicationProtocol_HTTPS,
//   					Certificate: awscdk.Certificate_*FromCertificateArn(this, jsii.String("Cert2"), jsii.String("helloworld")),
//   					SslPolicy: awscdk.SslPolicy_TLS12_EXT,
//   				},
//   			},
//   		},
//   	},
//   	TargetGroups: []applicationTargetProps{
//   		&applicationTargetProps{
//   			ContainerPort: jsii.Number(80),
//   			Listener: jsii.String("listener"),
//   		},
//   		&applicationTargetProps{
//   			ContainerPort: jsii.Number(90),
//   			PathPattern: jsii.String("a/b/c"),
//   			Priority: jsii.Number(10),
//   			Listener: jsii.String("listener"),
//   		},
//   		&applicationTargetProps{
//   			ContainerPort: jsii.Number(443),
//   			Listener: jsii.String("listener2"),
//   		},
//   		&applicationTargetProps{
//   			ContainerPort: jsii.Number(80),
//   			PathPattern: jsii.String("a/b/c"),
//   			Priority: jsii.Number(10),
//   			Listener: jsii.String("listener2"),
//   		},
//   	},
//   })
//
//   loadBalancedFargateService.TargetGroups[0].ConfigureHealthCheck(&HealthCheck{
//   	Port: jsii.String("8050"),
//   	Protocol: awscdk.Protocol_HTTP,
//   	HealthyThresholdCount: jsii.Number(2),
//   	UnhealthyThresholdCount: jsii.Number(2),
//   	Timeout: awscdk.Duration_*Seconds(jsii.Number(10)),
//   	Interval: awscdk.Duration_*Seconds(jsii.Number(30)),
//   	HealthyHttpCodes: jsii.String("200"),
//   })
//
//   loadBalancedFargateService.TargetGroups[1].ConfigureHealthCheck(&HealthCheck{
//   	Port: jsii.String("8050"),
//   	Protocol: awscdk.Protocol_HTTP,
//   	HealthyThresholdCount: jsii.Number(2),
//   	UnhealthyThresholdCount: jsii.Number(2),
//   	Timeout: awscdk.Duration_*Seconds(jsii.Number(10)),
//   	Interval: awscdk.Duration_*Seconds(jsii.Number(30)),
//   	HealthyHttpCodes: jsii.String("200"),
//   })
//
type Protocol string

const (
	// HTTP (ALB health checks and NLB health checks).
	Protocol_HTTP Protocol = "HTTP"
	// HTTPS (ALB health checks and NLB health checks).
	Protocol_HTTPS Protocol = "HTTPS"
	// TCP (NLB, NLB health checks).
	Protocol_TCP Protocol = "TCP"
	// TLS (NLB).
	Protocol_TLS Protocol = "TLS"
	// UDP (NLB).
	Protocol_UDP Protocol = "UDP"
	// Listen to both TCP and UDP on the same port (NLB).
	Protocol_TCP_UDP Protocol = "TCP_UDP"
)

