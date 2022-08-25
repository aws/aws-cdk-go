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
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
//   	maxAzs: jsii.Number(1),
//   })
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationMultipleTargetGroupsFargateService(this, jsii.String("myService"), &applicationMultipleTargetGroupsFargateServiceProps{
//   	cluster: ecs.NewCluster(this, jsii.String("EcsCluster"), &clusterProps{
//   		vpc: vpc,
//   	}),
//   	memoryLimitMiB: jsii.Number(256),
//   	taskImageOptions: &applicationLoadBalancedTaskImageProps{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	enableExecuteCommand: jsii.Boolean(true),
//   	loadBalancers: []applicationLoadBalancerProps{
//   		&applicationLoadBalancerProps{
//   			name: jsii.String("lb"),
//   			idleTimeout: awscdk.Duration.seconds(jsii.Number(400)),
//   			domainName: jsii.String("api.example.com"),
//   			domainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &publicHostedZoneProps{
//   				zoneName: jsii.String("example.com"),
//   			}),
//   			listeners: []applicationListenerProps{
//   				&applicationListenerProps{
//   					name: jsii.String("listener"),
//   					protocol: awscdk.ApplicationProtocol_HTTPS,
//   					certificate: awscdk.Certificate.fromCertificateArn(this, jsii.String("Cert"), jsii.String("helloworld")),
//   					sslPolicy: awscdk.SslPolicy_TLS12_EXT,
//   				},
//   			},
//   		},
//   		&applicationLoadBalancerProps{
//   			name: jsii.String("lb2"),
//   			idleTimeout: awscdk.Duration.seconds(jsii.Number(120)),
//   			domainName: jsii.String("frontend.com"),
//   			domainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &publicHostedZoneProps{
//   				zoneName: jsii.String("frontend.com"),
//   			}),
//   			listeners: []*applicationListenerProps{
//   				&applicationListenerProps{
//   					name: jsii.String("listener2"),
//   					protocol: awscdk.ApplicationProtocol_HTTPS,
//   					certificate: awscdk.Certificate.fromCertificateArn(this, jsii.String("Cert2"), jsii.String("helloworld")),
//   					sslPolicy: awscdk.SslPolicy_TLS12_EXT,
//   				},
//   			},
//   		},
//   	},
//   	targetGroups: []applicationTargetProps{
//   		&applicationTargetProps{
//   			containerPort: jsii.Number(80),
//   			listener: jsii.String("listener"),
//   		},
//   		&applicationTargetProps{
//   			containerPort: jsii.Number(90),
//   			pathPattern: jsii.String("a/b/c"),
//   			priority: jsii.Number(10),
//   			listener: jsii.String("listener"),
//   		},
//   		&applicationTargetProps{
//   			containerPort: jsii.Number(443),
//   			listener: jsii.String("listener2"),
//   		},
//   		&applicationTargetProps{
//   			containerPort: jsii.Number(80),
//   			pathPattern: jsii.String("a/b/c"),
//   			priority: jsii.Number(10),
//   			listener: jsii.String("listener2"),
//   		},
//   	},
//   })
//
//   loadBalancedFargateService.targetGroups[0].configureHealthCheck(&healthCheck{
//   	port: jsii.String("8050"),
//   	protocol: awscdk.Protocol_HTTP,
//   	healthyThresholdCount: jsii.Number(2),
//   	unhealthyThresholdCount: jsii.Number(2),
//   	timeout: awscdk.Duration.seconds(jsii.Number(10)),
//   	interval: awscdk.Duration.seconds(jsii.Number(30)),
//   	healthyHttpCodes: jsii.String("200"),
//   })
//
//   loadBalancedFargateService.targetGroups[1].configureHealthCheck(&healthCheck{
//   	port: jsii.String("8050"),
//   	protocol: awscdk.Protocol_HTTP,
//   	healthyThresholdCount: jsii.Number(2),
//   	unhealthyThresholdCount: jsii.Number(2),
//   	timeout: awscdk.Duration.seconds(jsii.Number(10)),
//   	interval: awscdk.Duration.seconds(jsii.Number(30)),
//   	healthyHttpCodes: jsii.String("200"),
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

