package awselasticloadbalancingv2


// The mutual authentication configuration information.
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
type MutualAuthentication struct {
	// Indicates whether expired client certificates are ignored.
	//
	// Cannot be used with MutualAuthenticationMode.OFF or MutualAuthenticationMode.PASS_THROUGH
	// Default: false.
	//
	IgnoreClientCertificateExpiry *bool `field:"optional" json:"ignoreClientCertificateExpiry" yaml:"ignoreClientCertificateExpiry"`
	// The client certificate handling method.
	// Default: MutualAuthenticationMode.OFF
	//
	MutualAuthenticationMode MutualAuthenticationMode `field:"optional" json:"mutualAuthenticationMode" yaml:"mutualAuthenticationMode"`
	// The trust store.
	//
	// Cannot be used with MutualAuthenticationMode.OFF or MutualAuthenticationMode.PASS_THROUGH
	// Default: - no trust store.
	//
	TrustStore ITrustStore `field:"optional" json:"trustStore" yaml:"trustStore"`
}

