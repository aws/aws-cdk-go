package awselasticloadbalancingv2


// The client certificate handling method.
//
// Example:
//   import acm "github.com/aws/aws-cdk-go/awscdk"
//
//   var certificate certificate
//   var lb applicationLoadBalancer
//   var bucket bucket
//
//
//   trustStore := elbv2.NewTrustStore(this, jsii.String("Store"), &TrustStoreProps{
//   	Bucket: Bucket,
//   	Key: jsii.String("rootCA_cert.pem"),
//   })
//
//   lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(443),
//   	Protocol: elbv2.ApplicationProtocol_HTTPS,
//   	Certificates: []iListenerCertificate{
//   		certificate,
//   	},
//   	// mTLS settings
//   	MutualAuthentication: &MutualAuthentication{
//   		IgnoreClientCertificateExpiry: jsii.Boolean(false),
//   		MutualAuthenticationMode: elbv2.MutualAuthenticationMode_VERIFY,
//   		TrustStore: *TrustStore,
//   	},
//   	DefaultAction: elbv2.ListenerAction_FixedResponse(jsii.Number(200), &FixedResponseOptions{
//   		ContentType: jsii.String("text/plain"),
//   		MessageBody: jsii.String("Success mTLS"),
//   	}),
//   })
//
type MutualAuthenticationMode string

const (
	// Off.
	MutualAuthenticationMode_OFF MutualAuthenticationMode = "OFF"
	// Application Load Balancer sends the whole client certificate chain to the target using HTTP headers.
	MutualAuthenticationMode_PASS_THROUGH MutualAuthenticationMode = "PASS_THROUGH"
	// Application Load Balancer performs X.509 client certificate authentication for clients when a load balancer negotiates TLS connections.
	MutualAuthenticationMode_VERIFY MutualAuthenticationMode = "VERIFY"
)

