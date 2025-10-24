package awselasticloadbalancingv2


// Options for `ListenerAction.fixedResponse()`.
//
// Example:
//   import acm "github.com/aws/aws-cdk-go/awscdk"
//
//   var certificate Certificate
//   var lb ApplicationLoadBalancer
//   var bucket Bucket
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
//   	Certificates: []IListenerCertificate{
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
type FixedResponseOptions struct {
	// Content Type of the response.
	//
	// Valid Values: text/plain | text/css | text/html | application/javascript | application/json.
	// Default: - Automatically determined.
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The response body.
	// Default: - No body.
	//
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
}

