package awselasticloadbalancingv2


// Load balancing protocol for application load balancers.
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
//   		AdvertiseTrustStoreCaNames: jsii.Boolean(true),
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
type ApplicationProtocol string

const (
	// HTTP.
	ApplicationProtocol_HTTP ApplicationProtocol = "HTTP"
	// HTTPS.
	ApplicationProtocol_HTTPS ApplicationProtocol = "HTTPS"
)

