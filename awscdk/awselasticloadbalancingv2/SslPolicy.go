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
//
//   var vpc vpc
//   var cluster cluster
//
//
//   domainZone := awscdk.HostedZone_FromLookup(this, jsii.String("Zone"), &HostedZoneProviderProps{
//   	DomainName: jsii.String("example.com"),
//   })
//   certificate := awscdk.Certificate_FromCertificateArn(this, jsii.String("Cert"), jsii.String("arn:aws:acm:us-east-1:123456:certificate/abcdefg"))
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
//   	Vpc: Vpc,
//   	Cluster: Cluster,
//   	Certificate: Certificate,
//   	SslPolicy: awscdk.SslPolicy_RECOMMENDED,
//   	DomainName: jsii.String("api.example.com"),
//   	DomainZone: DomainZone,
//   	RedirectHTTP: jsii.Boolean(true),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html
//
// Experimental.
type SslPolicy string

const (
	// The recommended security policy.
	// Experimental.
	SslPolicy_RECOMMENDED SslPolicy = "RECOMMENDED"
	// Strong foward secrecy ciphers and TLV1.2 only (2020 edition). Same as FORWARD_SECRECY_TLS12_RES, but only supports GCM versions of the TLS ciphers.
	// Experimental.
	SslPolicy_FORWARD_SECRECY_TLS12_RES_GCM SslPolicy = "FORWARD_SECRECY_TLS12_RES_GCM"
	// Strong forward secrecy ciphers and TLS1.2 only.
	// Experimental.
	SslPolicy_FORWARD_SECRECY_TLS12_RES SslPolicy = "FORWARD_SECRECY_TLS12_RES"
	// Forward secrecy ciphers and TLS1.2 only.
	// Experimental.
	SslPolicy_FORWARD_SECRECY_TLS12 SslPolicy = "FORWARD_SECRECY_TLS12"
	// Forward secrecy ciphers only with TLS1.1 and higher.
	// Experimental.
	SslPolicy_FORWARD_SECRECY_TLS11 SslPolicy = "FORWARD_SECRECY_TLS11"
	// Forward secrecy ciphers only.
	// Experimental.
	SslPolicy_FORWARD_SECRECY SslPolicy = "FORWARD_SECRECY"
	// TLS1.2 only and no SHA ciphers.
	// Experimental.
	SslPolicy_TLS12 SslPolicy = "TLS12"
	// TLS1.2 only with all ciphers.
	// Experimental.
	SslPolicy_TLS12_EXT SslPolicy = "TLS12_EXT"
	// TLS1.1 and higher with all ciphers.
	// Experimental.
	SslPolicy_TLS11 SslPolicy = "TLS11"
	// Support for DES-CBC3-SHA.
	//
	// Do not use this security policy unless you must support a legacy client
	// that requires the DES-CBC3-SHA cipher, which is a weak cipher.
	// Experimental.
	SslPolicy_LEGACY SslPolicy = "LEGACY"
)

