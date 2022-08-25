package awselasticloadbalancingv2


// Elastic Load Balancing provides the following security policies for Application Load Balancers.
//
// We recommend the Recommended policy for general use. You can
// use the ForwardSecrecy policy if you require Forward Secrecy
// (FS).
//
// You can use one of the TLS policies to meet compliance and security
// standards that require disabling certain TLS protocol versions, or to
// support legacy clients that require deprecated ciphers.
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
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html
//
type SslPolicy string

const (
	// The recommended security policy for TLS listeners.
	//
	// This is the default policy for listeners created using the AWS Management Console.
	SslPolicy_RECOMMENDED_TLS SslPolicy = "RECOMMENDED_TLS"
	// The recommended policy for http listeners.
	//
	// This is the default security policy for listeners created using the AWS CLI.
	SslPolicy_RECOMMENDED SslPolicy = "RECOMMENDED"
	// TLS1.2 and 1.3.
	SslPolicy_TLS13_RES SslPolicy = "TLS13_RES"
	// TLS1.2 and 1.3 and no SHA ciphers.
	SslPolicy_TLS13_EXT1 SslPolicy = "TLS13_EXT1"
	// TLS1.2 and 1.3 with all ciphers.
	SslPolicy_TLS13_EXT2 SslPolicy = "TLS13_EXT2"
	// TLS1.0 through 1.3 with all ciphers.
	SslPolicy_TLS13_10 SslPolicy = "TLS13_10"
	// TLS1.1 through 1.3 with all ciphers.
	SslPolicy_TLS13_11 SslPolicy = "TLS13_11"
	// TLS1.3 only.
	SslPolicy_TLS13_13 SslPolicy = "TLS13_13"
	// Strong foward secrecy ciphers and TLV1.2 only (2020 edition). Same as FORWARD_SECRECY_TLS12_RES, but only supports GCM versions of the TLS ciphers.
	SslPolicy_FORWARD_SECRECY_TLS12_RES_GCM SslPolicy = "FORWARD_SECRECY_TLS12_RES_GCM"
	// Strong forward secrecy ciphers and TLS1.2 only.
	SslPolicy_FORWARD_SECRECY_TLS12_RES SslPolicy = "FORWARD_SECRECY_TLS12_RES"
	// Forward secrecy ciphers and TLS1.2 only.
	SslPolicy_FORWARD_SECRECY_TLS12 SslPolicy = "FORWARD_SECRECY_TLS12"
	// Forward secrecy ciphers only with TLS1.1 and 1.2.
	SslPolicy_FORWARD_SECRECY_TLS11 SslPolicy = "FORWARD_SECRECY_TLS11"
	// Forward secrecy ciphers only.
	SslPolicy_FORWARD_SECRECY SslPolicy = "FORWARD_SECRECY"
	// TLS1.2 only and no SHA ciphers.
	SslPolicy_TLS12 SslPolicy = "TLS12"
	// TLS1.2 only with all ciphers.
	SslPolicy_TLS12_EXT SslPolicy = "TLS12_EXT"
	// TLS1.1 and 1.2 with all ciphers.
	SslPolicy_TLS11 SslPolicy = "TLS11"
	// Support for DES-CBC3-SHA.
	//
	// Do not use this security policy unless you must support a legacy client
	// that requires the DES-CBC3-SHA cipher, which is a weak cipher.
	SslPolicy_LEGACY SslPolicy = "LEGACY"
)

