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
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	MaxAzs: jsii.Number(1),
//   })
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
	// TLS 1.3 only with AES 128 and 256 GCM SHA ciphers.
	SslPolicy_FIPS_TLS13_13 SslPolicy = "FIPS_TLS13_13"
	// TLS 1.2 and 1.3 with AES and ECDHE GCM/SHA ciphers.
	SslPolicy_FIPS_TLS13_12_RES SslPolicy = "FIPS_TLS13_12_RES"
	// TLS 1.2 and 1.3 with ECDHE SHA/GCM ciphers, excluding SHA1 ciphers.
	SslPolicy_FIPS_TLS13_12 SslPolicy = "FIPS_TLS13_12"
	// TLS 1.2 and 1.3 with all ECDHE ciphers.
	SslPolicy_FIPS_TLS13_12_EXT0 SslPolicy = "FIPS_TLS13_12_EXT0"
	// TLS 1.2 and 1.3 with all AES and ECDHE ciphers excluding SHA1 ciphers.
	SslPolicy_FIPS_TLS13_12_EXT1 SslPolicy = "FIPS_TLS13_12_EXT1"
	// TLS 1.2 and 1.3 with all ciphers.
	SslPolicy_FIPS_TLS13_12_EXT2 SslPolicy = "FIPS_TLS13_12_EXT2"
	// TLS1.1 through 1.3 with all ciphers.
	SslPolicy_FIPS_TLS13_11 SslPolicy = "FIPS_TLS13_11"
	// TLS1.0 through 1.3 with all ciphers.
	SslPolicy_FIPS_TLS13_10 SslPolicy = "FIPS_TLS13_10"
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

